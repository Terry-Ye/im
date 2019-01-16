package main

import (
	"github.com/smallnest/rpcx/server"
	inet "im/libs/net"
	"context"
	log "github.com/sirupsen/logrus"
	"im/libs/proto"
	"strconv"
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
	userInfo, err := RedisCli.HGetAll(key).Result()
	if err != nil {
		log.Infof("RedisCli HGetAll key :%s , err:%s",key,  err)
	}

	reply.Uid = userInfo["UserId"]
	roomUserKey := getRoomUserKey(strconv.Itoa(int(args.RoomId)))
	// UserId
	if reply.Uid == "" {
		reply.Uid = define.NO_AUTH
	}else {
		userKey := getKey(reply.Uid)
		log.Infof("logic redis set uid serverId:%s, serverId : %s", userKey, args.ServerId)
		validTime := define.REDIS_BASE_VALID_TIME * time.Second
		err  = RedisCli.Set(userKey, args.ServerId,  validTime).Err()
		if err != nil {
			log.Warnf("logic set err:%s", err)
		}
		// write redis roomUserList
		RedisCli.HSet(roomUserKey, reply.Uid, userInfo["UserName"])
	}

	roomUserInfo, err := RedisCli.HGetAll(roomUserKey).Result()
	if err != nil {
		log.Warnf("RedisCli HGetAll roomUserInfo key:%s, err: %s", roomUserKey, err)
	}

	if err = RedisPublishRoomInfo(int32(args.RoomId), len(roomUserInfo), roomUserInfo); err != nil {
		log.Warnf("Count redis RedisPublishRoomCount err: %s", err)
		return
	}

	log.Infof("logic rpc uid:%s", reply.Uid)

	return
}

func (rpc *LogicRpc) Disconnect(ctx context.Context, args *proto.DisconnArg, reply *proto.DisconnReply) (err error) {
	// 房间人数减少
	roomUserKey := getRoomUserKey(strconv.Itoa(int(args.RoomId)))

	err = RedisCli.HDel(roomUserKey, args.Uid).Err()
	if err != nil {
		log.Warnf("HDel getRoomUserKey err : %s", err)
	}
	roomUserInfo, err := RedisCli.HGetAll(roomUserKey).Result()
	if err != nil {
		log.Warnf("RedisCli HGetAll roomUserInfo key:%s, err: %s", roomUserKey, err)
	}

	if err = RedisPublishRoomInfo(args.RoomId, len(roomUserInfo), roomUserInfo); err != nil {
		log.Warnf("Count redis RedisPublishRoomCount err: %s", err)
		return
	}
	return
}

//
