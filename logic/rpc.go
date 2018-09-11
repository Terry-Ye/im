package main

import (
	"github.com/smallnest/rpcx/server"
	log "github.com/sirupsen/logrus"
	inet "im/libs/net"
)

func InitRPC() (err error) {
	var (
		bind          string
		network, addr string
	)
	for _, bind = range Conf.Base.RPCAddrs {
		if network, addr, err = inet.ParseNetwork(bind); err != nil {
			log.Panicf("InitLogicRpc ParseNetwork error : %s", err)
		}
		go createServer(network, addr)
	}

}

func createServer(network string, addr string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve(network, addr)
}



