package main

import (

	inet "im/libs/net"
	"context"
	log "github.com/sirupsen/logrus"
	"im/libs/proto"
	"im/libs/define"
	"github.com/smallnest/rpcx/server"
)

type PushRpc int

func InitPushRpc(addrs []RpcPushAddrs) (err error){
	var (
		network, addr string
	)
	for _, bind := range addrs {
		if network, addr, err = inet.ParseNetwork(bind.Addr); err != nil {
			log.Panicf("InitLogicRpc ParseNetwork error : %s", err)
		}
		go createServer(network, addr)
	}
	return
}


func createServer(network string, addr string) {
	s := server.NewServer()
	s.RegisterName(define.RPC_PUSH_SERVER_PATH, new(PushRpc), "")
	s.Serve(network, addr)
}


func (rpc *PushRpc) MPushMsg(ctx context.Context, args *proto.PushMsgArg, noReply *proto.NoReply) (err error) {

	log.Info("rpc PushMsg :%v ", args)
	if args == nil {
		log.Errorf("rpc PushRpc() error(%v)", err)
		return
	}

	return
}

func (rpc *PushRpc) PushSingleMsg(ctx context.Context, args *proto.PushMsgArg, noReply *proto.NoReply) (err error) {

	log.Info("rpc PushMsg :%v ", args)
	if args == nil {
		log.Errorf("rpc PushRpc() error(%v)", err)
		return
	}

	return
}







