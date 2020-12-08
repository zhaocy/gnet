package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/zhaocy/gnet"
	pb "github.com/zhaocy/gnet/test/pb"
	book "github.com/zhaocy/gnet/test/pb/book"
	"net/http"
 _ "net/http/pprof"
)

var gamerMap = map[string]gnet.IMsgQue{}

func C2SBookHandlerFunc(msgque gnet.IMsgQue, msg *gnet.Message) bool{
	fmt.Printf("book-> %v", msg)
	ab := msg.C2S().(*book.AddressBook)
	gnet.LogInfo("-> %v", ab.People)
	s2c := &pb.Pb {
		GamerLoginS2C: &pb.Pb_GamerLoginS2C{Json: proto.String("登陆成功1"),},
	}

	msgque.SetGroupId("001")

	msgque.Send(gnet.NewDataMsg(gnet.PbData(s2c)))
	return true
}

func C2SHandlerFunc(msgque gnet.IMsgQue, msg *gnet.Message) bool {
	fmt.Printf("pb-> :%v\n", msg)
	ppb := msg.C2S().(*pb.Pb)
	gnet.LogInfo("-> %v", ppb.GamerLoginC2S)
	gnet.LogInfo("%v", ppb.GamerGlobalChatC2S)

	if ppb.GamerLoginC2S != nil{
		gnet.SendGroup("001",gnet.NewDataMsg(gnet.PbData(ppb)))
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
	} /*else if ab.People!=nil{
		fmt.Println("---------")
	}*/

	return true
}

func main() {
	ExtNetHandler := &gnet.DefMsgHandler{}
	PbParser := &gnet.Parser{Type: gnet.ParserTypePB}
	//PbParser.RegisterMsg(&pb.Pb{}, nil)
	//ExtNetHandler.RegisterMsg(&pb.Pb{}, C2SHandlerFunc)
	PbParser.RegisterMsg(&book.AddressBook{},nil)

	ExtNetHandler.RegisterMsg(&book.AddressBook{},C2SBookHandlerFunc)
	gnet.StartServer("tcp://:5000", gnet.MsgTypeCmd, ExtNetHandler, PbParser)

	go func() {
		http.ListenAndServe(":3999",nil)
	}()
	gnet.WaitForSystemExit(nil)
}
