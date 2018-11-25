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
	beego.Debug("redis get key %s", GetKey(key))
	return RedisCli.Get(GetKey(key)).Val()
}

func HGet(key string, field string ) string{
	// beego.Debug("redis get key %s", GetKey(key))
	return RedisCli.HGet(GetKey(key), field).Val()

}

func HGetAll(key string) (map[string]string, error)  {
	return RedisCli.HGetAll(GetKey(key)).Result()
}

func Set(key string, val interface{}, expiration time.Duration) ( err error) {
	err = RedisCli.Set(GetKey(key), val, expiration).Err()
	return
}


func Delete(key string) error {
	return RedisCli.Del(GetKey(key)).Err()

}

func GetKey(key string) string {
	beego.Debug("redis get Prefix %s",Prefix)
	return Prefix + key
}

func HMSet(key string, mapData map[string]interface{}) (err error){
	err = RedisCli.HMSet(GetKey(key), mapData).Err()
	return

}







