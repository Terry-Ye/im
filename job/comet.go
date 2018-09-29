package main

import (
	log "github.com/sirupsen/logrus"
	"im/libs/define"

	"github.com/smallnest/rpcx/client"
)

type CometRpc int

var (
	logicRpcClient client.XClient
)

func InitComets(cometConf []CometConf) (err error)  {
	log.Infof("len : %d ", len(cometConf))
	LogicAddrs := make([]*client.KVPair, len(cometConf))
	log.Infof("cometConf : %v ", cometConf)

	for i, bind := range cometConf {
		log.Infof("bind key %d", bind.Key)
		log.Infof("bind key %s", bind.Addr)
		b := new(client.KVPair)
		b.Key = bind.Addr
		// 需要转int 类型
		LogicAddrs[i] = b

	}
	log.Infof("LogicAddrs %v", LogicAddrs)
	d := client.NewMultipleServersDiscovery(LogicAddrs)

	logicRpcClient = client.NewXClient(define.RPC_LOGIC_SERVER_PATH, client.Failover, client.RoundRobin, d, client.DefaultOption)
	log.Infof("comet InitLogicRpc Server : %v ", logicRpcClient)
	return

}
