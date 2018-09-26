package main

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"im/libs/define"
	"im/libs/proto"
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
func RedisPublishCh(serverId int8, uid string , msg []byte ) (err error) {
	var redisMsg = &proto.RedisMsg{Op:define.REDIS_MESSAGE_BROADCAST, ServerId:serverId, UserId:uid,Msg:msg}
	err = RedisCli.Publish(define.REDIS_SUB_CHANNEL, redisMsg).Err()
	return
}







