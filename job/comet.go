package main

import (
	log "github.com/sirupsen/logrus"
	"im/libs/define"
	"im/libs/proto"
	"context"
	"github.com/smallnest/rpcx/client"
)

type CometRpc int

var (
	logicRpcClient client.XClient
	RpcClientList map[int8]client.XClient
	// RpcClientList map[int8]client.XClient

)


func InitComets(cometConf []CometConf) (err error)  {
	log.Infof("len : %d ", len(cometConf))
	LogicAddrs := make([]*client.KVPair, len(cometConf))

	RpcClientList = make(map[int8]client.XClient, len(cometConf))

	log.Infof("cometConf : %v ", cometConf)

	for i, bind := range cometConf {
		// log.Infof("bind key %d", bind.Key)
		// log.Infof("bind Addr %s", bind.Addr)
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
	pushMsgArg := &proto.PushMsgArg{Uid:userId, P:proto.Proto{Ver:1, Operation:define.REDIS_MESSAGE_SINGLE,Body:msg}}
	// log.Infof("PushSingleToComet serverId %d", serverId)
	log.Infof("PushSingleToComet RpcClientList %v", RpcClientList[serverId])
	err := RpcClientList[serverId].Call(context.Background(), "PushSingleMsg", pushMsgArg, proto.NoReply{})
	if err != nil {
		log.Infof(" PushSingleToComet Call err %v", err)
	}
}




