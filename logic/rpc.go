package main

import (
	"github.com/smallnest/rpcx/server"
	inet "im/libs/net"
	"context"
	log "github.com/sirupsen/logrus"
	"im/libs/proto"

	"im/libs/define"
)

type LogicRpc int

func InitRPC() (err error) {

	var (
		network, addr string
	)
	for _, bind := range Conf.Base.RPCAddrs {
		if network, addr, err = inet.ParseNetwork(bind); err != nil {
			log.Panicf("InitLogicRpc ParseNetwork error : %s", err)
		}
		go createServer(network, addr)
	}
	// select {}
	return
}

func createServer(network string, addr string) {

	s := server.NewServer()
	s.RegisterName(define.RPC_LOGIC_SERVER_PATH, new(LogicRpc), "")
	s.Serve(network, addr)

}

func (rpc *LogicRpc) Connect(ctx context.Context, args *proto.ConnArg, reply *proto.ConnReply) (err error) {
	log.Info("rpc logic 2  rpc uid ")
	if args == nil {
		log.Errorf("Connect() error(%v)", err)
		return
	}
	reply.Uid = "555"
	log.Infof("logic rpc uid:%s", reply.Uid)

	return
}




//

