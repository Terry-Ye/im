package main

import (
	"im/libs/proto"
	"bytes"
	"strconv"
	"im/libs/define"
)


func getRouter(auth string) (router *proto.Router, err error) {
	var key bytes.Buffer
	key.WriteString(define.REDIS_AUTH_PREFIX)
	key.WriteString(auth)


	userInfo, err := RedisCli.HGetAll(key.String()).Result()
	if err != nil {
		return
	}

	uid, err := strconv.ParseInt(userInfo["UserId"], 10, 64)
	if err != nil {
		return
	}
	rid, err := strconv.ParseInt(userInfo["RoomId"], 10, 32)
	if err != nil {
		return
	}
	sid, err := strconv.ParseInt(userInfo["ServerId"], 10, 16)
	if err != nil {
		return
	}
	router = &proto.Router{ServerId: int8(sid), RoomId: int32(rid), UserId: uid}
	return

}
