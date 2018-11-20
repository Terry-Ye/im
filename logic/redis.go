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
func RedisPublishCh(serverId int8, uid string, msg []byte) (err error) {
	var redisMsg = &proto.RedisMsg{
		Op: define.OP_SINGLE_SEND,
		ServerId: serverId,
		UserId: uid,
		Msg: msg,
	}

	redisMsgStr, err := json.Marshal(redisMsg)

	log.Infof("redisMsg info : %s", redisMsgStr)

	err = RedisCli.Publish(define.REDIS_SUB, redisMsgStr).Err()
	return
}

func RedisPublishRoom(rid int32, msg []byte) (err error) {
	var redisMsg = &proto.RedisMsg{
		Op: define.OP_ROOM_SEND,
		RoomId: rid,
		Msg: msg,
	}
	redisMsgStr, err := json.Marshal(redisMsg)
	log.Infof("redisMsg info : %s", redisMsgStr)
	err = RedisCli.Publish(define.REDIS_SUB, redisMsgStr).Err()
	return
}

func getKey(key string) (string){

	var returnKey bytes.Buffer
	returnKey.WriteString(define.REDIS_PREFIX)
	returnKey.WriteString(key)
	return returnKey.String()


}
