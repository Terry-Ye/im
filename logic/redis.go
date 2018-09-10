package main

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)
var (
	RedisCli *redis.Client

)
const (
	REDIS_PUSH_CODE = "redis_push"
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
