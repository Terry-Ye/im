package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/client"
	"im/libs/proto"
)

var (
	logicRpcClient client.XClient
)

func InitLogicRpcClient() (err error) {

	d := client.NewZookeeperDiscovery(Conf.ZookeeperInfo.BasePath,
		Conf.ZookeeperInfo.ServerPathLogic,
		[]string{Conf.ZookeeperInfo.Host},
		nil)
	logicRpcClient = client.NewXClient(Conf.ZookeeperInfo.ServerPathLogic, client.Failtry, client.RandomSelect, d, client.DefaultOption)
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

func disconnect(disconnArg *proto.DisconnArg) (err error) {

	reply := &proto.DisconnReply{}
	if err = logicRpcClient.Call(context.Background(), "Disconnect", disconnArg, reply); err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	return
}
