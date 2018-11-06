package redis

import (
	"github.com/astaxie/beego"
	"time"
	"github.com/go-redis/redis"
	"im/web/module/config"
)


var (
	Prefix string
	RedisCli *redis.Client
)

func InitRedis() (err error) {
	config, err := config.Reader("database.conf")
	Prefix = config.String("redis::RedisPrefix")
	if err != nil  {
		beego.Error("config reader err: %v", err)
	}
	db, err := config.Int("redis::RedisDefaultDB")
	if err != nil {
		beego.Error("Redis get db err: %s",  err)
		return
	}

	RedisCli = redis.NewClient(&redis.Options{
		Addr:     config.String("redis::RedisAddr"),
		Password: config.String("redis::RedisPw"),        // no password set
		DB:        db, // use default DB
	})
	beego.Debug("redis db %d", db)
	if pong, err := RedisCli.Ping().Result(); err != nil {
		beego.Error("RedisCli Ping Result pong: %s,  err: %s", pong, err)

	}
	return
}




func Get(key string) string{
	return RedisCli.Get(GetKey(key)).String()
}

func Set(key string, val interface{}, expiration time.Duration) ( err error) {
	beego.Debug("RedisCli %v", RedisCli)
	err = RedisCli.Set(GetKey(key), val, expiration).Err()
	return
}

func GetKey(key string) string {
	return Prefix + key
}







