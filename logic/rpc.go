package main

import (
	"github.com/smallnest/rpcx/server"
	inet "im/libs/net"
	"context"
	log "github.com/sirupsen/logrus"
	"im/libs/proto"

	"im/libs/define"
	"time"
)

type LogicRpc int

func InitRPC() (err error) {

	var (
		network, addr string
	)
	for _, bind := range Conf.Base.RPCAddrs {
		log.Infof("RPCAddrs :%v", bind)
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

	if args == nil {
		log.Errorf("Connect() error(%v)", err)
		return
	}

	key := getKey(args.Auth)
	log.Infof("logic rpc key:%s", key)
	reply.Uid = RedisCli.HGet(key, "UserId").Val()

	if reply.Uid == "" {
		reply.Uid = define.NO_AUTH
	}else {
		userKey := getKey(reply.Uid)

		log.Infof("logic redis set uid serverId:%s, serverId : %s", userKey, args.ServerId)
		validTime := define.REDIS_BASE_VALID_TIME * time.Second
		err  = RedisCli.Set(userKey, args.ServerId,  validTime).Err()
		if err != nil {
			log.Infof("logic set err:%s", err)
		}
	}



	log.Infof("logic rpc uid:%s", reply.Uid)

	return
}

//
