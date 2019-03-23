package main

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"im/libs/define"
)

var (
	RedisCli *redis.Client
)

func InitRedis() (err error) {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     Conf.Redis.RedisAddr,
		Password: Conf.Redis.RedisPw,        // no password set
		DB:       Conf.Redis.RedisDefaultDB, // use default DB
	})
	if pong, err := RedisCli.Ping().Result(); err != nil {
		log.Infof("RedisCli Ping Result pong: %s,  err: %s", pong, err)
	}

	go func() {
		redisSub := RedisCli.Subscribe(define.REDIS_SUB)
		ch := redisSub.Channel()
		for {
			msg, ok := <-ch
			if !ok {
				log.Debugf("redisSub Channel !ok: %v", ok)
				break
			}

			push(msg.Payload)
			if Conf.Base.IsDebug == true {
				log.Infof("redisSub Subscribe msg : %s", msg.Payload)
			}

		}

	}()

	return
}
