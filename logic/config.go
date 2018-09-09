package main

import (
	"runtime"
	"time"
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
	flag.StringVar(&confPath, "d", "./", " set logic config file path")
}

type Config struct {
	Base      BaseConf      `mapstructure:"base"`
	Bucket    BucketConf    `mapstructure:"bucket"`
}


// 基础的配置信息
type BaseConf struct {
	Pidfile         string `mapstructure:"pidfile"`
	MaxProc         int
	PprofAddrs       []string `mapstructure:"pprofBind"` // 性能监控的域名端口

}

func InitConfig() (err error) {
	viper.SetConfigName("comet")
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
			Pidfile:         "/tmp/logic.pid",

			MaxProc:        runtime.NumCPU(),
			PprofAddrs:     []string{"localhost:6971"},
			// RedisAddrs:     []string{"localhost:6973"},
		},

	}
}

// 重新加载配置
// func ReloadConfig() {
//
// }