package main

import (
	"github.com/smallnest/rpcx/client"
	log "github.com/sirupsen/logrus"
	"im/libs/proto"
	"context"
	"im/libs/define"
)

var (
	logicRpcClient client.XClient
)

func InitLogicRpc(rpcLogicAddrs []RpcLogicAddrs) (err error) {
	LogicAddrs := make([]*client.KVPair, len(rpcLogicAddrs))
	for i, bind := range rpcLogicAddrs {

		b := new(client.KVPair)
		b.Key = bind.Addr
		LogicAddrs[i] = b

	}

	d := client.NewMultipleServersDiscovery(LogicAddrs)
	logicRpcClient = client.NewXClient(define.RPC_LOGIC_SERVER_PATH, client.Failover, client.RoundRobin, d, client.DefaultOption)
	return
}

func connect(connArg *proto.ConnArg) (uid string, err error) {
	reply := &proto.ConnReply{}
	err = logicRpcClient.Call(context.Background(), "Connect", connArg, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	uid = reply.Uid
	log.Infof("comet logic uid :%s", reply.Uid)

	return
}

func disconnect(disconnArg *proto.DisconnArg) (err error){
	reply := &proto.DisconnReply{}
	if err = logicRpcClient.Call(context.Background(), "Disconnect", disconnArg, reply); err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	return
}
