package main

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"im/libs/define"
	"im/libs/proto"
	"encoding/json"
	"bytes"
)

var (
	RedisCli *redis.Client
)

func InitRedis() (err error) {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     Conf.Base.RedisAddr,
		Password: Conf.Base.RedisPw,        // no password set
		DB:       Conf.Base.RedisDefaultDB, // use default DB
	})
	if pong, err := RedisCli.Ping().Result(); err != nil {
		log.Infof("RedisCli Ping Result pong: %s,  err: %s", pong, err)
	}

	return
}

// 发布订阅消息
func RedisPublishCh(serverId int8, uid string, msg []byte, formUserInfo *proto.Router) (err error) {
	var redisMsg = &proto.RedisMsg{
		Op: define.REDIS_MESSAGE_SINGLE,
		ServerId: serverId,
		UserId: uid,
		Msg: msg,
		FormUserId: formUserInfo.UserId,
		FormServerId: formUserInfo.ServerId,
	}

	redisMsgStr, err := json.Marshal(redisMsg)

	log.Infof("redisMsg info : %s", redisMsgStr)

	err = RedisCli.Publish(define.REDIS_SUB, redisMsgStr).Err()
	return
}

func RedisPublishRoom(rid int32, msg []byte, formUserId string) (err error) {
	var redisMsg = &proto.RedisMsg{
		Op: define.REDIS_MESSAGE_ROOM,
		RoomId: rid,
		Msg: msg,
		FormUserId:formUserId,
	}
	redisMsgStr, err := json.Marshal(redisMsg)
	log.Infof("redisMsg info : %s", redisMsgStr)
	err = RedisCli.Publish(define.REDIS_SUB, redisMsgStr).Err()
	return
}

func getAuthKey(auth string) (string){

	var key bytes.Buffer
	key.WriteString(define.REDIS_AUTH_PREFIX)
	key.WriteString(auth)
	return key.String()


}
