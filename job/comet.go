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

)


func InitComets(cometConf []CometConf) (err error)  {
	log.Infof("len : %d ", len(cometConf))
	LogicAddrs := make([]*client.KVPair, len(cometConf))
	log.Infof("cometConf : %v ", cometConf)

	for i, bind := range cometConf {
		log.Infof("bind key %d", bind.Key)
		log.Infof("bind key %s", bind.Addr)
		b := new(client.KVPair)
		b.Key = bind.Addr
		// 需要转int 类型
		LogicAddrs[i] = b
		d := client.NewPeer2PeerDiscovery(bind.Addr, "")

		RpcClientList[bind.Key] = client.NewXClient(define.RPC_LOGIC_SERVER_PATH, client.Failtry, client.RandomSelect, d, client.DefaultOption)
		defer RpcClientList[bind.Key].Close()
	}

	// servers
	log.Infof("comet InitLogicRpc Server : %v ", RpcClientList)

	return
}

func PushSingleToComet(serverId int8, userId string, msg []byte) {

	pushMsgArg := &proto.PushMsgArg{Uid:userId, P:proto.Proto{Ver:1, Operation:define.REDIS_MESSAGE_SINGLE,Body:msg}}

	RpcClientList[serverId].Call(context.Background(), "pushSingle", pushMsgArg, proto.NoReply{})





}




