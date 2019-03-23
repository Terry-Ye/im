package main

import (
	"context"
	metrics "github.com/rcrowley/go-metrics"
	log "github.com/sirupsen/logrus"
	"im/libs/define"
	inet "im/libs/net"
	"im/libs/proto"
	// example "github.com/rpcx-ecosystem/rpcx-examples3"
	"flag"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

type PushRpc int

func InitPushRpc(addrs []RpcPushAddrs) (err error) {
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
	flag.Parse()
	s := server.NewServer()

	addRegistryPlugin(s)

	s.RegisterName(define.RPC_COMET_SERVER_PATH, new(PushRpc), "1")
	log.Infof("createServer addr %s", addr)
	s.Serve(network, addr)
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@0.0.0.0:6912",
		ZooKeeperServers: []string{"127.0.0.1:2181"},
		BasePath:         "/im_logic_rpc_server",
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}

// func (rpc *PushRpc) MPushMsg(ctx context.Context, args *proto.PushMsgArg, noReply *proto.NoReply) (err error) {
//
// 	log.Info("rpc PushMsg :%v ", args)
// 	if args == nil {
// 		log.Errorf("rpc PushRpc() error(%v)", err)
// 		return
// 	}
//
// 	return
// }

func (rpc *PushRpc) PushSingleMsg(ctx context.Context, args *proto.PushMsgArg, SuccessReply *proto.SuccessReply) (err error) {
	var (
		bucket  *Bucket
		channel *Channel
	)

	log.Info("rpc PushMsg :%v ", args)
	if args == nil {
		log.Errorf("rpc PushRpc() error(%v)", err)
		return
	}
	bucket = DefaultServer.Bucket(args.Uid)
	if channel = bucket.Channel(args.Uid); channel != nil {
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
	for _, bucket := range DefaultServer.Buckets {
		bucket.BroadcastRoom(args)
		// room.next

	}
	return
}

/**
广播房间人数
*/
func (rpc *PushRpc) PushRoomCount(ctx context.Context, args *proto.RoomMsgArg, SuccessReply *proto.SuccessReply) (err error) {
	SuccessReply.Code = define.SUCCESS_REPLY
	SuccessReply.Msg = define.SUCCESS_REPLY_MSG
	log.Infof("PushRoomCount count %v", args)
	for _, bucket := range DefaultServer.Buckets {
		bucket.BroadcastRoom(args)
		// room.next
	}
	return
}

/**
广播房间信息
*/
func (rpc *PushRpc) PushRoomInfo(ctx context.Context, args *proto.RoomMsgArg, SuccessReply *proto.SuccessReply) (err error) {
	log.Infof("PushRoomInfo  %v", args)
	SuccessReply.Code = define.SUCCESS_REPLY
	SuccessReply.Msg = define.SUCCESS_REPLY_MSG

	for _, bucket := range DefaultServer.Buckets {
		bucket.BroadcastRoom(args)
		// room.next
	}
	return
}
