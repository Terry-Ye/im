package main

import (
	"github.com/smallnest/rpcx/log"
	inet "im/libs/net"

	"im/libs/define"
	"github.com/smallnest/rpcx/server"
)

type CometRpc int

func InitComets(cometConf map[string]string) (err error)  {
	var (
		network, addr string
	)
	for key, bind := range cometConf {
		if network, addr, err = inet.ParseNetwork(bind); err != nil {
			log.Error("inet.ParseNetwork() error(%v)", err)
			return
		}
		go createServer(network, addr)
	}
}
func createServer(network string, addr string) {

	s := server.NewServer()
	s.RegisterName(define.RPC_COMET_SERVER_PATH, new(CometRpc), "")
	s.Serve(network, addr)

}