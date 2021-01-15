package gnet

import (
	"net"
	"reflect"
	"strings"
	"sync"
	"time"
)

type HandlerFunc func(msgque IMsgQue, msg *Message) bool

type MsgType int //消息类型
const (
	MsgTypeMsg MsgType = iota //消息基于确定的消息头
	MsgTypeCmd                //消息没有消息头，以\n分割
)

type NetType int

const (
	NetTypeTcp NetType = iota //TCP类型
	NetTypeUdp                //UDP类型
	NetTypeWs                 //websocket
)

type ConnType int

const (
	ConnTypeListen ConnType = iota //监听
	ConnTypeConn                   //连接产生的
	ConnTypeAccept                 //Accept产生的
)

//消息队列基类
type msgQue struct {
	id      uint32        //唯一标识
	cwrite  chan *Message //写入通道
	stop    int32         //停止标记
	msgType MsgType       //消息类型
	connTyp ConnType      //通道类型

	handler       IMsgHandler    //处理器
	parser        IParser        //解析器
	parserFactory IParserFactory //解析器工厂
	timeout       int            //传输超时
	lastTick      int64

	init           bool
	available      bool
	sendFast       bool
	multiplex      bool
	callback       map[int]chan *Message
	group          map[string]int
	user           interface{}
	callbackLock   sync.Mutex
	gmsgId         uint16
	realRemoteAddr string //当使用代理是，需要特殊设置客户端真实IP
}

func (r *msgQue) SetSendFast() {
	r.sendFast = true
}

func (r *msgQue) SetUser(user interface{}) {
	r.user = user
}

func (r *msgQue) getGMsg(add bool) *gMsg {
	if add {
		r.gmsgId++
	}
	gm := gmsgArray[r.gmsgId]
	return gm
}
func (r *msgQue) SetCmdReadRaw() {

}
func (r *msgQue) Available() bool {
	return r.available
}

func (r *msgQue) GetUser() interface{} {
	return r.user
}

func (r *msgQue) GetHandler() IMsgHandler {
	return r.handler
}

func (r *msgQue) GetMsgType() MsgType {
	return r.msgType
}

func (r *msgQue) GetConnType() ConnType {
	return r.connTyp
}

func (r *msgQue) Id() uint32 {
	return r.id
}

func (r *msgQue) SetTimeout(t int) {
	if t >= 0 {
		r.timeout = t
	}
}

func (r *msgQue) isTimeout(tick *time.Timer) bool {
	left := int(Timestamp - r.lastTick)
	if left < r.timeout || r.timeout == 0 {
		if r.timeout == 0 {
			tick.Reset(time.Second * time.Duration(DefMsgQueTimeout))
		} else {
			tick.Reset(time.Second * time.Duration(r.timeout-left))
		}
		return false
	}
	LogInfo("msgque close because timeout id:%v wait:%v timeout:%v", r.id, left, r.timeout)
	return true
}

func (r *msgQue) GetTimeout() int {
	return r.timeout
}

func (r *msgQue) Reconnect(t int) {

}

//是否代理
func (r *msgQue) IsProxy() bool {
	return r.realRemoteAddr != ""
}

//设置实际远程地址，直接相连的远程地址是通过连接conn获取
func (r *msgQue) SetRealRemoteAddr(addr string) {
	r.realRemoteAddr = addr
}

//设置分组
func (r *msgQue) SetGroupId(group string) {
	r.callbackLock.Lock()
	if r.group == nil {
		r.group = make(map[string]int)
	}
	r.group[group] = 0
	r.callbackLock.Unlock()
}

//删除分组
func (r *msgQue) DelGroupId(group string) {
	r.callbackLock.Lock()
	if r.group != nil {
		delete(r.group, group)
	}
	r.callbackLock.Unlock()
}

//清除所有分组
func (r *msgQue) ClearGroupId(group string) {
	r.callbackLock.Lock()
	r.group = nil
	r.callbackLock.Unlock()
}

//是否在分组中
func (r *msgQue) IsInGroup(group string) bool {
	re := false
	r.callbackLock.Lock()
	if r.group != nil {
		_, re = r.group[group]
	}
	r.callbackLock.Unlock()
	return re
}

func (r *msgQue) SetMultiplex(multiplex bool, cwriteCnt int) bool {
	t := r.multiplex
	r.multiplex = multiplex
	if cwriteCnt > 0 {
		r.cwrite = make(chan *Message, cwriteCnt)
	}
	return t
}

//发送消息
func (r *msgQue) Send(m *Message) (re bool) {
	if m == nil || r.stop == 1 {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			re = false
		}
	}()
	if Config.AutoCompressLen > 0 && m.Head != nil && m.Head.Len >= Config.AutoCompressLen && (m.Head.Flags&FlagCompress) == 0 {
		m.Head.Flags |= FlagCompress
		m.Data = GZipCompress(m.Data)
		m.Head.Len = uint32(len(m.Data))
	}
	select {
	case r.cwrite <- m:
	default:
		LogWarn("msgque write channel full msgque:%v", r.id)
		r.cwrite <- m
	}

	return true
}

//发送回调消息，设置回调消息通道，
func (r *msgQue) SendCallback(m *Message, c chan *Message) (re bool) {
	if c == nil || cap(c) < 1 {
		LogError("try send callback but chan is null or no buffer")
		return
	}
	if r.Send(m) {
		r.setCallback(m.Tag(), c)
	} else {
		c <- nil
		return
	}
	return true
}

func (r *msgQue) DelCallback(m *Message) {
	if r.callback == nil {
		return
	}
	r.callbackLock.Lock()
	delete(r.callback, m.Tag())
	r.callbackLock.Unlock()
}

func (r *msgQue) SendString(str string) (re bool) {
	return r.Send(&Message{Data: []byte(str)})
}

func (r *msgQue) SendStringLn(str string) (re bool) {
	return r.SendString(str + "\n")
}

func (r *msgQue) SendByteStr(str []byte) (re bool) {
	return r.SendString(string(str))
}

func (r *msgQue) SendByteStrLn(str []byte) (re bool) {
	return r.SendString(string(str) + "\n")
}

func (r *msgQue) tryCallback(msg *Message) (re bool) {
	if r.callback == nil {
		return false
	}
	defer func() {
		if err := recover(); err != nil {

		}
		r.callbackLock.Unlock()
	}()
	r.callbackLock.Lock()
	if r.callback != nil {
		tag := msg.Tag()
		if c, ok := r.callback[tag]; ok {
			delete(r.callback, tag)
			c <- msg
			re = true
		}
	}
	return
}

func (r *msgQue) setCallback(tag int, c chan *Message) {
	defer func() {
		if err := recover(); err != nil {

		}
		r.callback[tag] = c
		r.callbackLock.Unlock()
	}()

	r.callbackLock.Lock()
	if r.callback == nil {
		r.callback = make(map[int]chan *Message)
	}
	oc, ok := r.callback[tag]
	if ok { //可能已经关闭
		oc <- nil
	}
}

func (r *msgQue) baseStop() {
	if r.cwrite != nil {
		close(r.cwrite)
	}

	for k, v := range r.callback {
		Try(func() {
			v <- nil
		}, func(i interface{}) {

		})
		delete(r.callback, k)
	}
	msgqueMapSync.Lock()
	delete(msgqueMap, r.id)
	msgqueMapSync.Unlock()
	LogInfo("msgque close id:%d", r.id)
}

func (r *msgQue) processMsg(msgque IMsgQue, msg *Message) bool {
	if r.multiplex {
		Go(func() {
			r.processMsgTrue(msgque, msg)
		})
	} else {
		return r.processMsgTrue(msgque, msg)
	}
	return true
}

//处理原始消息
func (r *msgQue) processMsgTrue(msgque IMsgQue, msg *Message) bool {
	if msg.Head != nil && msg.Head.Flags&FlagCompress > 0 && msg.Data != nil {
		data, err := GZipUnCompress(msg.Data)
		if err != nil {
			LogError("msgque uncompress failed msgque:%v cmd:%v act:%v len:%v err:%v", msgque.Id(), msg.Head.Cmd, msg.Head.Act, msg.Head.Len, err)
			return false
		}
		msg.Data = data
		msg.Head.Flags -= FlagCompress
		msg.Head.Len = uint32(len(msg.Data))
	}
	//消息解析器解析消息得到可读数据
	if r.parser != nil {
		mp, err := r.parser.ParseC2S(msg)
		if err == nil {
			msg.IMsgParser = mp
		} else {
			if r.parser.GetErrType() == ParseErrTypeSendRemind {
				if msg.Head != nil {
					r.Send(r.parser.GetRemindMsg(err, r.msgType).CopyTag(msg))
				} else {
					r.Send(r.parser.GetRemindMsg(err, r.msgType))
				}
				return true
			} else if r.parser.GetErrType() == ParseErrTypeClose {
				return false
			} else if r.parser.GetErrType() == ParseErrTypeContinue {
				return true
			}
		}
	}
	//如果发送的是回调消息，优先处理回调消息
	if msgque.tryCallback(msg) {
		return true
	}
	//获取处理消息函数
	f := r.handler.GetHandlerFunc(msgque, msg)
	if f == nil {
		f = r.handler.OnProcessMsg
	}
	return f(msgque, msg)
}

//消息处理接口
type IMsgHandler interface {
	OnNewMsgQue(msgque IMsgQue) bool                         //新的消息队列
	OnDelMsgQue(msgque IMsgQue)                              //消息队列关闭
	OnProcessMsg(msgque IMsgQue, msg *Message) bool          //默认的消息处理函数
	OnConnectComplete(msgque IMsgQue, ok bool) bool          //连接成功完成
	GetHandlerFunc(msgque IMsgQue, msg *Message) HandlerFunc //根据消息获得处理函数
}

type DefMsgHandler struct {
	msgMap  map[int]HandlerFunc
	typeMap map[reflect.Type]HandlerFunc
}

func (r *DefMsgHandler) OnNewMsgQue(msgque IMsgQue) bool                { return true }
func (r *DefMsgHandler) OnDelMsgQue(msgque IMsgQue)                     {}
func (r *DefMsgHandler) OnProcessMsg(msgque IMsgQue, msg *Message) bool { return true }
func (r *DefMsgHandler) OnConnectComplete(msgque IMsgQue, ok bool) bool { return true }
func (r *DefMsgHandler) GetHandlerFunc(msgque IMsgQue, msg *Message) HandlerFunc {
	if msg.Head == nil {
		if r.typeMap != nil {
			if f, ok := r.typeMap[reflect.TypeOf(msg.C2S())]; ok {
				return f
			}
		}
	} else if r.msgMap != nil {
		if f, ok := r.msgMap[msg.CmdAct()]; ok {
			return f
		}
	}

	return nil
}

func (r *DefMsgHandler) RegisterMsg(v interface{}, fun HandlerFunc) {
	msgType := reflect.TypeOf(v)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		LogFatal("message pointer required")
		return
	}
	if r.typeMap == nil {
		r.typeMap = map[reflect.Type]HandlerFunc{}
	}
	r.typeMap[msgType] = fun
}

func (r *DefMsgHandler) Register(cmd, act uint8, fun HandlerFunc) {
	if r.msgMap == nil {
		r.msgMap = map[int]HandlerFunc{}
	}
	r.msgMap[CmdAct(cmd, act)] = fun
}

//消息队列接口
type IMsgQue interface {
	Id() uint32            //消息队列ID
	GetMsgType() MsgType   //获取消息类型
	GetConnType() ConnType //获取通道类型
	GetNetType() NetType   //获取网络类型

	LocalAddr() string             //本地地址
	RemoteAddr() string            //远程地址
	SetRealRemoteAddr(addr string) //设置实际的远程地址（代理实现）

	Stop()           //停止
	IsStop() bool    //是否停机
	Available() bool //消息队列是否可用
	IsProxy() bool   //是否代理

	Send(m *Message) (re bool)                          //发送消息
	SendString(str string) (re bool)                    //发送字符串消息
	SendStringLn(str string) (re bool)                  //发送有换行的字符串
	SendByteStr(str []byte) (re bool)                   //发送byte数组消息
	SendByteStrLn(str []byte) (re bool)                 //发送有换行符的byte数组消息
	SendCallback(m *Message, c chan *Message) (re bool) //发送回调消息,c接收回调消息
	DelCallback(m *Message)                             //删除回调
	SetSendFast()                                       //设置快速发送消息
	SetTimeout(t int)                                   //设置超时时间
	SetCmdReadRaw()                                     //cmd读取原始数据
	GetTimeout() int                                    //获取超时时间
	Reconnect(t int)                                    //重连间隔  最小1s，此函数仅能连接关闭是调用

	GetHandler() IMsgHandler //获取消息处理器

	SetUser(user interface{}) //设置User，服务器连接是serverID,客户端连接是gamerId
	GetUser() interface{}     //获取User

	SetGroupId(group string)     //设置分组,用作组播消息发送
	DelGroupId(group string)     //删除分组
	ClearGroupId(group string)   //清除分组
	IsInGroup(group string) bool //是否在某个组内
	//服务器内部通讯时提升效率，比如战斗服发送消息到网关服，应该在连接建立时使用，cwriteCnt大于0表示重新设置cwrite缓存长度，内网一般发送较快，不用考虑
	SetMultiplex(multiplex bool, cwriteCnt int) bool //设置多goroutine处理消息

	tryCallback(msg *Message) (re bool) //尝试执行消息回调
}

func StartServer(addr string, typ MsgType, handler IMsgHandler, parser IParserFactory) error {
	addrs := strings.Split(addr, "://")
	if addrs[0] == "tcp" || addrs[0] == "all" {
		listen, err := net.Listen("tcp", addrs[1])
		if err == nil {
			msgque := newTcpListen(listen, typ, handler, parser, addr)
			Go(func() {
				LogDebug("process listen for tcp msgque:%d", msgque.id)
				msgque.listen()
				LogDebug("process listen end for tcp msgque:%d", msgque.id)
			})
		} else {
			LogError("listen on %s failed, errstr:%s", addr, err)
			return err
		}
	}

	if addrs[0] == "ws" || addrs[0] == "wss" {
		naddr := strings.SplitN(addrs[1], "/", 2)
		url := "/"
		if len(naddr) > 1 {
			url = "/" + naddr[1]
		}
		if addrs[0] == "wss" {
			Config.EnableWss = true
		}
		if typ != MsgTypeCmd {
			LogInfo("ws type msgque noly support MsgTypeCmd now auto set to MsgTypeCmd")
		}
		msgque := newWsListen(naddr[0], url, MsgTypeCmd, handler, parser)
		Go(func() {
			LogDebug("process listen for ws msgque:%d", msgque.id)
			msgque.listen()
			LogDebug("process listen end for ws msgque:%d", msgque.id)
		})
	}
	return nil
}
