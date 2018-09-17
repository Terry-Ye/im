package main

import (
	"github.com/smallnest/rpcx/client"
	log "github.com/sirupsen/logrus"
	"im/libs/proto"
	"context"
)

var (
	logicRpcClient client.XClient
)

func InitLogicRpc() (err error) {

	LogicAddrs := make([]*client.KVPair, len(Conf.Base.RpcLogicAddr))

	for i, bind := range Conf.Base.RpcLogicAddr {
		log.Infof("bind %s", bind)
		b := new(client.KVPair)
		b.Key = bind
		LogicAddrs[i] = b

	}
	log.Infof("server :%v", LogicAddrs)
	d := client.NewMultipleServersDiscovery(LogicAddrs)
	// d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: "tcp@0.0.0.0:6923"}})

	logicRpcClient = client.NewXClient("LogicRpc", client.Failover, client.RoundRobin, d, client.DefaultOption)
	log.Infof("comet InitLogicRpc Server : %v ", logicRpcClient)
	return
}

func connect(connArg *proto.ConnArg) (uid string, err error) {

	log.Infof("comet logic rpc logicRpcClient %s:", logicRpcClient)
	reply := &proto.ConnReply{}
	err = logicRpcClient.Call(context.Background(), "Connect", connArg, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	uid = reply.Uid
	log.Infof("comet logic uid :%s", reply.Uid)

	return
}
