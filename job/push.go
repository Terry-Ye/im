package main

import (
	"im/libs/define"
	"encoding/json"
	"im/libs/proto"
	"math/rand"
	"github.com/smallnest/rpcx/log"
)

type pushArg struct {
	ServerId int8
	UserId   string
	Msg      []byte
	RoomId   int32
	// FormUserId string
	// FormServerId int8
	// FormUserName string
}

var pushChs []chan *pushArg

func InitPush() {
	pushChs = make([]chan *pushArg, Conf.Base.PushChan)
	for i := 0; i < len(pushChs); i++ {

		pushChs[i] = make(chan *pushArg, Conf.Base.PushChanSize)
		go processPush(pushChs[i])
	}
}

func processPush(ch chan *pushArg) {
	var arg *pushArg
	for {
		arg = <-ch
		PushSingleToComet(arg.ServerId, arg.UserId, arg.Msg)


	}
}
func push(msg string) (err error) {
	m := &proto.RedisMsg{}
	msgByte := []byte(msg)
	if err := json.Unmarshal(msgByte, m); err != nil {
		log.Infof(" json.Unmarshal err:%v ", err)
	}
	log.Infof("push m info %s", m)

	switch m.Op {
	case define.OP_SINGLE_SEND:
		pushChs[rand.Int()%Conf.Base.PushChan] <- &pushArg{
			ServerId: m.ServerId,
			UserId: m.UserId,
			Msg: m.Msg,

		}
		break
	case define.OP_ROOM_SEND:
		broadcastRoomToComet(m.RoomId, m.Msg)
		break;
	case define.OP_ROOM_COUNT_SEND:
		broadcastRoomCountToComet(m.RoomId, m.Count)
	case define.OP_ROOM_INFO_SEND:
		broadcastRoomInfoToComet(m.RoomId, m.RoomUserInfo)
	}


	return
}
