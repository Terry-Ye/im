package main

import (
	"github.com/smallnest/rpcx/server"
	log "github.com/sirupsen/logrus"
	inet "im/libs/net"
	"im/libs/proto"
)

type Rpc int

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

}

func createServer(network string, addr string) {
	s := server.NewServer()
	s.RegisterName("Logic", new(Rpc), "")
	s.Serve(network, addr)
}


func (rpc *Rpc) Connect(args *proto.ConnArg, reply *proto.ConnReply) (err error) {
	// test
	reply.Uid = "23333333"
	return

}



