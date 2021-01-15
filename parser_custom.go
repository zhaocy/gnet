package gnet

import (
	"github.com/zhaocy/codec"
)

type CustomParser struct {
	*Parser
}

func (r *CustomParser) ParseC2S(msg *Message) (IMsgParser, error) {
	if msg == nil {
		return nil, ErrCustomUnPack
	}

	if len(msg.Data) == 0 {
		return nil, ErrCustomUnPack
	}

	headData := make([]byte, MsgShortHeadSize)
	var data []byte
	var head *MessageShortHead

	headData = msg.Data[:MsgShortHeadSize]
	if head = NewMessageShortHead(headData); head == nil {
		LogError("short read msg head failed")
		return nil, ErrCustomUnPack
	}

	LogDebug("head %v", headData)
	msg.ShortHead = head
	if head.Len > 0 {
		data = make([]byte, head.Len)
		data = msg.Data[MsgShortHeadSize:]
		msg.Data = data
		LogDebug("data len: %v %v",head.Len, msg.Data)
	}else{
		return nil, ErrCustomUnPack
	}

	for _, p := range r.typeMap {
		if p.C2S() != nil {

			err := CustomUnPack(msg.Data, p.C2S())
			if err != nil {
				continue
			}
			p.parser = r
			return &p, nil
		}
	}

	return nil, ErrCustomUnPack
}

//解码接受的数据
func CustomUnPack(data []byte, msg interface{}) error{
	if data == nil || msg == nil {
		return ErrCustomUnPack
	}

	err := codec.Decode(data, msg)
	if err != nil {
		return ErrCustomUnPack
	}
	return nil
}

func (r *CustomParser) PackMsg(v interface{}) []byte {
	data, _ := CustomPack(v)
	return data
}

//编码要发送的数据
func CustomPack(msg interface{}) ([]byte, error) {
	if msg == nil {
		return nil, ErrPBPack
	}

	data, err := codec.Encode(msg)
	if err != nil {
		LogError("custom encode err")
	}

	return data, nil
}

func (r *CustomParser) GetRemindMsg(err error, t MsgType) *Message {
	if t == MsgTypeMsg {
		return nil
	} else {
		return nil
	}
}