package main

import (
	"im/libs/proto"
	"bytes"
	"strconv"
	"im/libs/define"
	"github.com/smallnest/rpcx/log"
)


func getRouter(auth string) (router *proto.Router, err error) {
	var key bytes.Buffer
	key.WriteString(define.REDIS_AUTH_PREFIX)
	key.WriteString(auth)

	log.Infof("key %s", key.String())
	userInfo, err := RedisCli.HGetAll(key.String()).Result()


	log.Infof("userInfo %v", userInfo)

	if err != nil {
		log.Infof("router err %v", err)
		return
	}



	// rid, err := strconv.ParseInt(userInfo["RoomId"], 10, 32)
	// if err != nil {
	// 	return
	// }
	sid, err := strconv.ParseInt(userInfo["ServerId"], 10, 8)
	if err != nil {
		log.Infof("router err %v", err)
		return
	}
	// router = &proto.Router{ServerId: int8(sid), RoomId: int32(rid), UserId: userInfo["UserId"]}
	router = &proto.Router{ServerId: int8(sid),  UserId: userInfo["UserId"]}
	return

}
