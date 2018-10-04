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
	msgStr := []byte(msg)
	if err := json.Unmarshal(msgStr, m); err != nil {
		log.Infof(" json.Unmarshal err:%v ", err)
	}
	switch m.Op {
	case define.REDIS_MESSAGE_SINGLE:
		pushChs[rand.Int()%Conf.Base.PushChan] <- &pushArg{ServerId: m.ServerId, UserId: m.UserId, Msg: m.Msg, RoomId: m.RoomId}
		break

	}
	return
}
