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
	Base      BaseConf  `mapstructure:"base"`
	CometConf CometConf `mapstructure:"cometsAddrs"`
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
}
type CometConf struct {
	Addrs map[string]string `mapstructure:"addr"` // rpc comet层的配置
}

func InitConfig() (err error) {
	Conf = NewConfig()
	viper.SetConfigName("job")
	viper.SetConfigType("toml")
	viper.AddConfigPath(confPath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}

	return nil
}

func NewConfig() *Config {
	return &Config{
		Base: BaseConf{
			Pidfile:    "/tmp/job.pid",
			MaxProc:    runtime.NumCPU(),
			PprofAddrs: []string{"localhost:6922"},

			RedisAddr:      "127.0.0.1:6379",
			RedisPw:        "",
			RedisDefaultDB: 0,
		},
		CometConf: CometConf{
			Addrs: make(map[string]string),
		},
	}
}

// 重新加载配置
// func ReloadConfig() {
//
// }
