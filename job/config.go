package main

import (
	"runtime"
	"flag"
	"github.com/spf13/viper"
	// "github.com/go-redis/redis"
	"fmt"
)

var (
	Conf     *Config
	confPath string
)

func init() {
	flag.StringVar(&confPath, "d", "./", " set job config file path")
}

type Config struct {
	Base      BaseConf    `mapstructure:"base"`
	CometConf []CometConf `mapstructure:"cometsAddrs"`
	// Bucket BucketConf `mapstructure:"bucket"`
}

// 基础的配置信息
type BaseConf struct {
	Pidfile    string   `mapstructure:"pidfile"`
	MaxProc    int
	PprofAddrs []string `mapstructure:"pprofBind"` //

	RedisAddr      string `mapstructure:"redisAddr"` //
	RedisPw        string `mapstructure:"redisPw"`
	RedisDefaultDB int    `mapstructure:"redisDefaultDB"`
	PushChan       int    `mapstructure:"pushChan"`
	PushChanSize   int    `mapstructure:"pushChanSize"`
	IsDebug		bool
}
type CometConf struct {
	Key  int8   `mapstructure:"key"`
	Addr string `mapstructure:"addr"`
}

func InitConfig() (err error) {
	Conf = NewConfig()
	viper.SetConfigName("job")
	viper.SetConfigType("toml")
	viper.AddConfigPath(confPath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	// log.Infof("conf %v", Conf)
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}

	return nil
}

func NewConfig() *Config {
	return &Config{
		Base: BaseConf{
			Pidfile:        "/tmp/job.pid",
			MaxProc:        runtime.NumCPU(),
			PprofAddrs:     []string{"localhost:6922"},
			RedisAddr:      "127.0.0.1:6379",
			RedisPw:        "",
			RedisDefaultDB: 0,
			PushChan:       2,
			PushChanSize:   50,
			IsDebug: true,
		},
		CometConf: []CometConf{
			{Key: 1, Addr: "tcp@0.0.0.0:6912"},
		},
	}
}

// 重新加载配置
// func ReloadConfig() {
//
// }
