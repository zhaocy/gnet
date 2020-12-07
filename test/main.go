package main

import (
	"github.com/zhaocy/gnet"
	"github.com/zhaocy/gnet/test/pb"
	"github.com/golang/protobuf/proto"
)

var gamerMap = map[string]gnet.IMsgQue{}

func C2SHandlerFunc(msgque gnet.IMsgQue, msg *gnet.Message) bool {
	ppb := msg.C2S().(*pb.Pb)
	gnet.LogInfo("%v", ppb.GamerLoginC2S)
	gnet.LogInfo("%v", ppb.GamerGlobalChatC2S)

	if ppb.GamerLoginC2S != nil{
		gamerMap[ppb.GamerLoginC2S.GetId()] = msgque
		s2c := &pb.Pb {
			GamerLoginS2C: &pb.Pb_GamerLoginS2C{Json: proto.String("登陆成功"),},
		}
		msgque.Send(gnet.NewDataMsg(gnet.PbData(s2c)))
		msgque.SetUser(ppb.GamerLoginC2S.GetId())
	} else if ppb.GamerGlobalChatC2S != nil {
		s2c := &pb.Pb {
			GamerGlobalChatS2C: &pb.Pb_GamerGlobalChatS2C{Data: ppb.GamerGlobalChatC2S.Data,},
		}
		gnet.Send(gnet.NewDataMsg(gnet.PbData(s2c)), func(msgque gnet.IMsgQue) bool {
			return true
		})
	} else if ppb.GamerChatC2S != nil {
		s2c := &pb.Pb{
			GamerChatS2C: &pb.Pb_GamerChatS2C{
				FromId: proto.String(msgque.GetUser().(string)),
				Data: proto.String("玩家不在线"),
			},
		}
		if mq, ok := gamerMap[ppb.GamerChatC2S.GetId()]; ok && mq.Available(){
			s2c.GamerChatS2C.Data = proto.String("发送成功")
			msgque.Send(gnet.NewDataMsg(gnet.PbData(s2c)))

			s2c.GamerChatS2C.Data = ppb.GamerChatC2S.Data
			mq.Send(gnet.NewDataMsg(gnet.PbData(s2c)))
		} else {
			msgque.Send(gnet.NewDataMsg(gnet.PbData(s2c)))
		}
	}

	return true
}

func main() {
	ExtNetHandler := &gnet.DefMsgHandler{}
	PbParser := &gnet.Parser{Type: gnet.ParserTypePB}
	PbParser.RegisterMsg(&pb.Pb{}, nil)
	ExtNetHandler.RegisterMsg(&pb.Pb{}, C2SHandlerFunc)
	gnet.StartServer("tcp://:5000", gnet.MsgTypeCmd, ExtNetHandler, PbParser)
	gnet.WaitForSystemExit(nil)
}
