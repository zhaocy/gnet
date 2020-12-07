package gnet

import (
	"encoding/json"
	"reflect"
)

type ParserType int //解析类型
const (
	ParserTypePB      ParserType = iota //protobuf类型，用于和客户端交互
	ParserTypeCmd                       //cmd类型，类似telnet指令，用于直接和程序交互
	ParserTypeJson                      //json类型，可以用于客户端或者服务器之间交互
	ParserTypeMsgpack                   //msgpack类型，可以用于客户端或者服务器之间交互
	ParserTypeCustom                    //自定义类型
	ParserTypeRaw                       //不做任何解析
)

type ParseErrType int

const (
	ParseErrTypeSendRemind ParseErrType = iota //消息解析失败发送提醒消息
	ParseErrTypeContinue                       //消息解析失败则跳过本条消息
	ParseErrTypeAlways                         //消息解析失败依然处理
	ParseErrTypeClose                          //消息解析失败则关闭连接
)

type ParseFunc func() interface{}

//解析器
type IParser interface {
	GetType() ParserType      //解析类型
	GetErrType() ParseErrType //解析错误类型
	ParseC2S(msg *Message) (IMsgParser, error)
	PackMsg(v interface{}) []byte
	GetRemindMsg(err error, t MsgType) *Message //获取自定义消息
}

type Parser struct {
	Type    ParserType
	ErrType ParseErrType

	msgMap  map[int]MsgParser
	typMap  map[reflect.Type]MsgParser
	cmdRoot *cmdParseNode
	parser  IParser
}

func (r *Parser) Get() IParser {
	switch r.Type {
	case ParserTypePB:
		if r.parser == nil {
			r.parser = &PBParser{Parser: r}
		}
	case ParserTypeCmd:
		return &CmdParser{Parser: r}
	case ParserTypeRaw:
		return nil
	}

	return r.parser
}

func (r *Parser) GetType() ParserType {
	return r.Type
}

func (r *Parser) GetErrType() ParseErrType {
	return r.ErrType
}

func (r *Parser) RegisterFunc(cmd, act uint8, c2sFunc ParseFunc, s2cFunc ParseFunc) {
	if r.msgMap == nil {
		r.msgMap = map[int]MsgParser{}
	}

	r.msgMap[CmdAct(cmd, act)] = MsgParser{c2sFunc: c2sFunc, s2cFunc: s2cFunc}
}

func (r *Parser) Register(cmd, act uint8, c2s interface{}, s2c interface{}) {
	var c2sFunc ParseFunc = nil
	var s2cFunc ParseFunc = nil

	if c2s != nil {
		c2sType := reflect.TypeOf(c2s).Elem()
		c2sFunc = func() interface{} {
			return reflect.New(c2sType).Interface()
		}
	}
	if s2c != nil {
		s2cType := reflect.TypeOf(s2c).Elem()
		s2cFunc = func() interface{} {
			return reflect.New(s2cType).Interface()
		}
	}
	r.RegisterFunc(cmd, act, c2sFunc, s2cFunc)
}

func (r *Parser) RegisterMsgFunc(c2sFunc ParseFunc, s2cFunc ParseFunc) {
	if r.Type == ParserTypeCmd {
		if r.cmdRoot == nil {
			r.cmdRoot = &cmdParseNode{}
		}
		registerCmdParser(r.cmdRoot, c2sFunc, s2cFunc)
	} else {
		if r.typMap == nil {
			r.typMap = map[reflect.Type]MsgParser{}
		}
		r.typMap[reflect.TypeOf(c2sFunc())] = MsgParser{c2sFunc: c2sFunc, s2cFunc: s2cFunc}
	}
}

func (r *Parser) RegisterMsg(c2s interface{}, s2c interface{}) {
	var c2sFunc ParseFunc = nil
	var s2cFunc ParseFunc = nil
	if c2s != nil {
		c2sType := reflect.TypeOf(c2s).Elem()
		c2sFunc = func() interface{} {
			return reflect.New(c2sType).Interface()
		}
	}
	if s2c != nil {
		s2cType := reflect.TypeOf(s2c).Elem()
		s2cFunc = func() interface{} {
			return reflect.New(s2cType).Interface()
		}
	}

	r.RegisterMsgFunc(c2sFunc, s2cFunc)
}

//消息解析器
type IMsgParser interface {
	C2S() interface{}  //parser Client To Server 反序列化得到的数据类型
	S2C() interface{}  //parser Server to Client
	C2SData() []byte   //序列化得到的[]byte数据
	S2CData() []byte   //Server to Client字节数据
	C2SString() string //序列化Client to Server 字符串
	S2CString() string //返回Server to Client 字符串
}

type MsgParser struct {
	s2c     interface{}
	c2s     interface{}
	c2sFunc ParseFunc
	s2cFunc ParseFunc
	parser  IParser
}

func (r *MsgParser) C2S() interface{} {
	if r.c2s == nil && r.c2sFunc != nil {
		r.c2s = r.c2sFunc()
	}
	return r.c2s
}

func (r *MsgParser) S2C() interface{} {
	if r.s2c == nil && r.s2cFunc != nil {
		r.s2c = r.s2cFunc()
	}
	return r.s2c
}

func (r *MsgParser) C2SData() []byte {
	return r.parser.PackMsg(r.C2S())
}

func (r *MsgParser) S2CData() []byte {
	return r.parser.PackMsg(r.S2C())
}

func (r *MsgParser) C2SString() string {
	return string(r.C2SData())
}

func (r *MsgParser) S2CString() string {
	return string(r.S2CData())
}

type IParserFactory interface {
	Get() IParser
}

func JsonUnPack(data []byte, msg interface{}) error {
	if data == nil || msg == nil {
		return ErrJsonUnPack
	}

	err := json.Unmarshal(data, msg)
	if err != nil {
		return ErrJsonUnPack
	}
	return nil
}

func JsonPack(msg interface{}) ([]byte, error) {
	if msg == nil {
		return nil, ErrJsonPack
	}

	data, err := json.Marshal(msg)
	if err != nil {
		LogInfo("")
		return nil, ErrJsonPack
	}

	return data, nil
}
