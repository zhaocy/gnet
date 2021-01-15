package gnet

import (
	"fmt"
	"unsafe"
)

const (
	MsgHeadSize      = 12
	MsgShortHeadSize = 6
)
const (
	FlagEncrypt  = 1 << 0 //数据是经过加密的
	FlagCompress = 1 << 1 //数据是经过压缩的
	FlagContinue = 1 << 2 //消息还有后续
	FlagNeedAck  = 1 << 3 //消息需要确认
	FlagAck      = 1 << 4 //确认消息
	FlagReSend   = 1 << 5 //重发消息
	FlagClient   = 1 << 6 //消息来自客户端，用于判断index来之服务器还是其他玩家
)

var MaxMsgDataSize uint32 = 1024 * 1024

//定义消息头
type MessageHead struct {
	Len   uint32 //数据长度
	Error uint16 //错误码
	Cmd   uint8  //命令
	Act   uint8  //动作
	Index uint16 //序号
	Flags uint16 //标记

	forever bool
	data    []byte //数据
}

//定义短消息头
type MessageShortHead struct {
	Cmd uint16
	Act uint16
	Len uint16 //数据长度
}

func (r *MessageHead) Bytes() []byte {
	if r.forever && r.data != nil {
		return r.data
	}
	r.data = make([]byte, MsgHeadSize)
	phead := (*MessageHead)(unsafe.Pointer(&r.data[0]))
	phead.Len = r.Len
	phead.Error = r.Error
	phead.Cmd = r.Cmd
	phead.Act = r.Act
	phead.Index = r.Index
	phead.Flags = r.Flags
	return r.data
}

func (r *MessageHead) FastBytes(data []byte) []byte {
	phead := (*MessageHead)(unsafe.Pointer(&data[0]))
	phead.Len = r.Len
	phead.Error = r.Error
	phead.Cmd = r.Cmd
	phead.Act = r.Act
	phead.Index = r.Index
	phead.Flags = r.Flags
	return data
}

func (r *MessageHead) BytesWithData(wdata []byte) []byte {
	if r.forever && r.data != nil {
		return r.data
	}
	r.Len = uint32(len(wdata))
	r.data = make([]byte, MsgHeadSize+r.Len)
	phead := (*MessageHead)(unsafe.Pointer(&r.data[0]))
	phead.Len = r.Len
	phead.Error = r.Error
	phead.Cmd = r.Cmd
	phead.Act = r.Act
	phead.Index = r.Index
	phead.Flags = r.Flags
	if wdata != nil {
		copy(r.data[MsgHeadSize:], wdata)
	}
	return r.data
}

func (r *MessageHead) FromBytes(data []byte) error {
	if len(data) < MsgHeadSize {
		return ErrMsgLenTooShort
	}
	phead := (*MessageHead)(unsafe.Pointer(&data[0]))
	r.Len = phead.Len
	r.Error = phead.Error
	r.Cmd = phead.Cmd
	r.Act = phead.Act
	r.Index = phead.Index
	r.Flags = phead.Flags
	if r.Len > MaxMsgDataSize {
		return ErrMsgLenTooLong
	}
	return nil
}

func (r *MessageHead) CmdAct() int {
	return CmdAct(r.Cmd, r.Act)
}

func (r *MessageHead) Tag() int {
	return Tag(r.Cmd, r.Act, r.Index)
}

func (r *MessageHead) String() string {
	return fmt.Sprintf("Len:%v Error:%v Cmd:%v Act:%v Index:%v Flags:%v", r.Len, r.Error, r.Cmd, r.Act, r.Index, r.Flags)
}

func NewMessageHead(data []byte) *MessageHead {
	head := &MessageHead{}
	if err := head.FromBytes(data); err != nil {
		return nil
	}
	return head
}

func (s *MessageShortHead) CmdAct() int {
	return CmdActUint16(s.Cmd, s.Act)
}

func (s *MessageShortHead) FromBytes(data []byte) error {
	if len(data) < MsgShortHeadSize {
		return ErrMsgLenTooShort
	}
	phead := (*MessageShortHead)(unsafe.Pointer(&data[0]))
	s.Len = phead.Len
	s.Cmd = phead.Cmd
	s.Act = phead.Act
	if uint32(s.Len)> MaxMsgDataSize{
		return ErrMsgLenTooLong
	}
	return nil
}

func NewMessageShortHead(data []byte) *MessageShortHead {
	shortHead := &MessageShortHead{}
	if err:= shortHead.FromBytes(data); err!=nil{
		return nil
	}
	return shortHead
}

//处理器
type Message struct {
	ShortHead *MessageShortHead
	Head       *MessageHead //消息头，可能为nil
	Data       []byte       //消息数据
	IMsgParser              //解析器
	User       interface{}  //用户自定义数据
}

func (r *Message) CmdAct() int {
	if r.Head != nil {
		return CmdAct(r.Head.Cmd, r.Head.Act)
	}
	return 0
}

func (r *Message) Len() uint32 {
	if r.Head != nil {
		return r.Head.Len
	}
	return 0
}

func (r *Message) Error() uint16 {
	if r.Head != nil {
		return r.Head.Error
	}
	return 0
}

func (r *Message) Cmd() uint8 {
	if r.Head != nil {
		return r.Head.Cmd
	}
	return 0
}

func (r *Message) Act() uint8 {
	if r.Head != nil {
		return r.Head.Act
	}
	return 0
}

func (r *Message) Index() uint16 {
	if r.Head != nil {
		return r.Head.Index
	}
	return 0
}

func (r *Message) Flags() uint16 {
	if r.Head != nil {
		return r.Head.Flags
	}
	return 0
}

func (r *Message) Tag() int {
	if r.Head != nil {
		return Tag(r.Head.Cmd, r.Head.Act, r.Head.Index)
	}
	return 0
}

func (r *Message) Bytes() []byte {
	if r.Head != nil {
		if r.Data != nil {
			return r.Head.BytesWithData(r.Data)
		}
		return r.Head.Bytes()
	}
	return r.Data
}

func (r *Message) CopyTag(old *Message) *Message {
	if r.Head != nil && old.Head != nil {
		r.Head.Cmd = old.Head.Cmd
		r.Head.Act = old.Head.Act
		r.Head.Index = old.Head.Index
	}
	return r
}

func NewErrMsg(err error) *Message {
	errcode, ok := errIdMap.Load(err)
	if !ok {
		errcode, _ = errIdMap.Load(ErrErrIdNotFound)
	}
	return &Message{
		Head: &MessageHead{
			Error: errcode.(uint16),
		},
	}
}

func NewStrMsg(str string) *Message {
	return &Message{
		Data: []byte(str),
	}
}

func NewDataMsg(data []byte) *Message {
	return &Message{
		Head: &MessageHead{
			Len: uint32(len(data)),
		},
		Data: data,
	}
}
