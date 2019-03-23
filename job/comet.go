package main

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/client"
	"im/libs/define"
	"im/libs/proto"
	"strconv"
	"strings"
)

// type CometRpc int

var (
	logicRpcClient client.XClient
	RpcClientList  map[int8]client.XClient
)

func InitComets() (err error) {

	d := client.NewZookeeperDiscovery(Conf.ZookeeperInfo.BasePath, Conf.ZookeeperInfo.ServerPathComet, []string{Conf.ZookeeperInfo.Host}, nil)
	RpcClientList = make(map[int8]client.XClient, len(d.GetServices()))
	// Get comet service configuration from zookeeper
	for _, cometConf := range d.GetServices() {

		cometConf.Value = strings.Replace(cometConf.Value, "=&tps=0", "", 1)

		serverId, error := strconv.ParseInt(cometConf.Value, 10, 8)
		if error != nil {
			log.Panicf("InitComets err，Can't find serverId. error: %s", error)
		}
		d := client.NewPeer2PeerDiscovery(cometConf.Key, "")
		RpcClientList[int8(serverId)] = client.NewXClient(Conf.ZookeeperInfo.ServerPathComet, client.Failtry, client.RandomSelect, d, client.DefaultOption)
		log.Infof("RpcClientList addr %s, v %v", cometConf.Key, RpcClientList[int8(serverId)])

	}
	logicRpcClient = client.NewXClient(Conf.ZookeeperInfo.ServerPathComet, client.Failtry, client.RandomSelect, d, client.DefaultOption)

	return
}

/**
广播消息到单个用户
*/
func PushSingleToComet(serverId int8, userId string, msg []byte) {
	log.Infof("PushSingleToComet Body %s", msg)
	pushMsgArg := &proto.PushMsgArg{Uid: userId, P: proto.Proto{Ver: 1, Operation: define.OP_SINGLE_SEND, Body: msg}}
	reply := &proto.SuccessReply{}
	err := RpcClientList[serverId].Call(context.Background(), "PushSingleMsg", pushMsgArg, reply)
	if err != nil {
		log.Infof(" PushSingleToComet Call err %v", err)
	}
	log.Infof("reply %s", reply.Msg)
}

/**
广播消息到房间
*/
func broadcastRoomToComet(RoomId int32, msg []byte) {
	pushMsgArg := &proto.RoomMsgArg{
		RoomId: RoomId, P: proto.Proto{
			Ver:       1,
			Operation: define.OP_ROOM_SEND,
			Body:      msg,
		},
	}
	reply := &proto.SuccessReply{}
	log.Infof("broadcastRoomToComet roomid %d", RoomId)
	for _, rpc := range RpcClientList {
		log.Infof("broadcastRoomToComet rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomMsg", pushMsgArg, reply)
	}
}

/**
广播在线人数到房间
*/
func broadcastRoomCountToComet(RoomId int32, count int) {

	var (
		body []byte
		err  error
	)
	msg := &proto.RedisRoomCountMsg{
		Count: count,
		Op:    define.OP_ROOM_COUNT_SEND,
	}

	if body, err = json.Marshal(msg); err != nil {
		log.Warnf("broadcastRoomCountToComet  json.Marshal err :%s", err)
		return
	}

	pushMsgArg := &proto.RoomMsgArg{
		RoomId: RoomId, P: proto.Proto{
			Ver:       1,
			Operation: define.OP_ROOM_SEND,
			Body:      body,
		},
	}

	reply := &proto.SuccessReply{}
	for _, rpc := range RpcClientList {
		log.Infof("broadcastRoomToComet rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomCount", pushMsgArg, reply)
	}
}

/**
广播房间信息到房间
*/
func broadcastRoomInfoToComet(RoomId int32, RoomUserInfo map[string]string) {

	var (
		body []byte
		err  error
	)
	msg := &proto.RedisRoomInfo{
		Count:        len(RoomUserInfo),
		Op:           define.OP_ROOM_COUNT_SEND,
		RoomUserInfo: RoomUserInfo,
		RoomId:       RoomId,
	}

	if body, err = json.Marshal(msg); err != nil {
		log.Warnf("broadcastRoomInfoToComet  json.Marshal err :%s", err)
		return
	}

	pushMsgArg := &proto.RoomMsgArg{
		RoomId: RoomId, P: proto.Proto{
			Ver:       1,
			Operation: define.OP_ROOM_SEND,
			Body:      body,
		},
	}

	reply := &proto.SuccessReply{}
	for _, rpc := range RpcClientList {
		log.Infof("broadcastRoomInfoToComet rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomInfo", pushMsgArg, reply)
		log.Infof("broadcastRoomInfoToComet rpc  reply %v", reply)
	}
}
