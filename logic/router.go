package main

import (
	"bytes"
	"im/libs/define"
	"github.com/smallnest/rpcx/log"
	"im/libs/proto"
)
var (
	userInfo map[string]string
)


func getRouterUserId(auth string) (userId string) {
	var key bytes.Buffer
	key.WriteString(define.REDIS_PREFIX)
	key.WriteString(auth)

	log.Infof("key %s", key.String())
	userId = RedisCli.HGet(key.String(), "UserId").Val()
	return
}



func getRouter(auth string) (router *proto.Router, err error) {
	userInfo, err = RedisCli.HGetAll(getKey(auth)).Result()
	if err != nil {
		return
	}
	log.Infof("getRouter auth :%s, userId:%s", auth, userInfo["UserId"])
	router = &proto.Router{UserId: userInfo["UserId"], UserName: userInfo["UserName"]}
	return

}

