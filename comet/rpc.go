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
	log.Info("InitPushRpc")
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
	log.Infof("createServer addr %s", addr)
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

func (rpc *PushRpc) PushSingleMsg(ctx context.Context, args *proto.PushMsgArg, SuccessReply *proto.SuccessReply) (err error) {
	var(
		bucket  *Bucket
		channel *Channel
	)

	log.Info("rpc PushMsg :%v ", args)
	if args == nil {
		log.Errorf("rpc PushRpc() error(%v)", err)
		return
	}
	bucket = DefaultServer.Bucket(args.Uid)
	if channel = bucket.Channel(args.Uid); channel !=  nil {
		err = channel.Push(&args.P)

		log.Infof("DefaultServer Channel err nil : %v", err)
		return
	}

	SuccessReply.Code = define.SUCCESS_REPLY
	SuccessReply.Msg = define.SUCCESS_REPLY_MSG
	log.Infof("SuccessReply v :%v", SuccessReply)
	return
}

func (rpc *PushRpc) PushRoomMsg(ctx context.Context, args *proto.RoomMsgArg, SuccessReply *proto.SuccessReply) (err error) {

	SuccessReply.Code = define.SUCCESS_REPLY
	SuccessReply.Msg = define.SUCCESS_REPLY_MSG
	log.Infof("PushRoomMsg msg %v", args)
	for _, bucket :=range DefaultServer.Buckets {
		bucket.BroadcastRoom(args)
		// room.next

	}
	return
}




