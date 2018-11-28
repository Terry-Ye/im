package main

import (
	log "github.com/sirupsen/logrus"
	"im/libs/define"
	"im/libs/proto"
	"context"
	"github.com/smallnest/rpcx/client"
	"encoding/json"
)

type CometRpc int

var (
	logicRpcClient client.XClient
	RpcClientList map[int8]client.XClient
	// RpcClientList map[int8]client.XClient

)


func InitComets(cometConf []CometConf) (err error)  {
	LogicAddrs := make([]*client.KVPair, len(cometConf))
	RpcClientList = make(map[int8]client.XClient, len(cometConf))

	for i, bind := range cometConf {
		// log.Infof("bind key %d", bind.Key)
		b := new(client.KVPair)
		b.Key = bind.Addr
		// 需要转int 类型
		LogicAddrs[i] = b
		d := client.NewPeer2PeerDiscovery(bind.Addr, "")
		RpcClientList[bind.Key] = client.NewXClient(define.RPC_PUSH_SERVER_PATH, client.Failtry, client.RandomSelect, d, client.DefaultOption)

		log.Infof("RpcClientList addr %s, v %v", bind.Addr, RpcClientList[bind.Key])

	}

	// servers
	log.Infof("comet InitLogicRpc Server : %v ", RpcClientList)

	return
}

func PushSingleToComet(serverId int8, userId string, msg []byte)  {
	log.Infof("PushSingleToComet Body %s", msg)
	pushMsgArg := &proto.PushMsgArg{Uid:userId, P:proto.Proto{Ver:1, Operation:define.OP_SINGLE_SEND,Body:msg}}
	// log.Infof("PushSingleToComet serverId %d", serverId)
	// log.Infof("PushSingleToComet RpcClientList %v", RpcClientList[serverId])
	reply := &proto.SuccessReply{}
	err := RpcClientList[serverId].Call(context.Background(), "PushSingleMsg", pushMsgArg, reply)
	if err != nil {
		log.Infof(" PushSingleToComet Call err %v", err)
	}
	log.Infof("reply %s", reply.Msg)
}


func broadcastRoomToComet(RoomId int32, msg []byte) {
	pushMsgArg := &proto.RoomMsgArg{
		RoomId:RoomId, P:proto.Proto{
			Ver:1,
			Operation:define.OP_ROOM_SEND,
			Body:msg,
		},
	}
	reply := &proto.SuccessReply{}
	log.Infof("broadcastRoomToComet roomid %d", RoomId)
	for _, rpc :=  range RpcClientList {
		log.Infof("broadcastRoomToComet rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomMsg", pushMsgArg, reply)
	}
}


func broadcastRoomCountToComet(RoomId int32, count int) {

	var (
		body []byte
		err error
	)
	msg :=	&proto.RedisRoomCountMsg{
		Count : count,
		Op : define.OP_ROOM_COUNT_SEND,
	}

	if body, err = json.Marshal(msg); err != nil {
		log.Warnf("broadcastRoomCountToComet  json.Marshal err :%s", err)
		return
	}

	pushMsgArg := &proto.RoomMsgArg{
		RoomId:RoomId, P:proto.Proto{
			Ver:1,
			Operation:define.OP_ROOM_SEND,
			Body: body,
		},
	}


	reply := &proto.SuccessReply{}
	for _, rpc :=  range RpcClientList {
		log.Infof("broadcastRoomToComet rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomCount", pushMsgArg, reply)
	}
}


