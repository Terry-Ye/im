package main

import (
	"github.com/smallnest/rpcx/server"
	"context"
	inet "im/libs/net"

	"im/libs/proto"
	"github.com/smallnest/rpcx/log"
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
	s.RegisterName("LogicRpc", new(LogicRpc), "")
	s.Serve(network, addr)

}

func (rpc *LogicRpc) Connect(ctx context.Context, args *proto.ConnArg, reply *proto.ConnReply) (err error) {

	if args == nil {
		err = ErrConnectArgs
		log.Error("Connect() error(%v)", err)
		return
	}

	reply.Uid = "333333"
	return
}

//

