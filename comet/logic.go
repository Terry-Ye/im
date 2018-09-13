package main

import (
	"github.com/smallnest/rpcx/client"
	log "github.com/sirupsen/logrus"

	"context"
	"im/libs/proto"
)

var (
	logicRpcClient client.XClient
)

func InitLogicRpc() (err error) {
	var (
		Server []*client.KVPair
		KVPair *client.KVPair
	)

	// rpcLogicAddrs = make([]*addr, len(Conf.Base.RpcLogicAddr))
	for _, bind := range Conf.Base.RpcLogicAddr {
		KVPair.Key = bind
		Server = append(Server, KVPair)
	}
	d := client.NewMultipleServersDiscovery(Server)
	logicRpcClient = client.NewXClient("Logic", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	log.Debugf("comet InitLogicRpc Server : %v ", Server)
	return
}

func connect(connArg proto.ConnArg) (uid string, err error){

	var reply proto.ConnReply
	err = logicRpcClient.Call(context.Background(), "Logic.Connect", connArg, &reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	uid = reply.Uid
	log.Debugf("%d * %d = %d", reply.Uid)

}
