package main

import (
	"flag"
	"github.com/spf13/viper"
	"runtime"
	// "github.com/go-redis/redis"
	"fmt"
	"time"
)

var (
	Conf     *Config
	confPath string
)

func init() {
	flag.StringVar(&confPath, "d", "./", " set logic config file path")
}

type Config struct {
	Base          BaseConf  `mapstructure:"base"`
	ZookeeperInfo Zookeeper `mapstructure:"zookeeper"`
	Redis         Redis     `mapstructure:"redis"`
	// Bucket BucketConf `mapstructure:"bucket"`
}
type Zookeeper struct {
	Host            string `mapstructure:"host"`
	BasePath        string `mapstructure:"basePath"`
	ServerId        string `mapstructure:"serverId"`
	ServerPathLogic string `mapstructure:"serverPathLogic"`
}

type Redis struct {
	RedisAddr      string `mapstructure:"RedisAddr"` //
	RedisPw        string `mapstructure:"redisPw"`
	RedisDefaultDB int    `mapstructure:"redisDefaultDB"`
}

// 基础的配置信息
type BaseConf struct {
	Pidfile          string `mapstructure:"pidfile"`
	MaxProc          int
	PprofAddrs       []string      `mapstructure:"pprofBind"` //
	HttpAddrs        []string      `mapstructure:"httpAddr"`  //
	RPCAddrs         []string      `mapstructure:"RPCAddrs"`  //
	HTTPReadTimeout  time.Duration `mapstructure:"HTTPReadTimeout"`
	HTTPWriteTimeout time.Duration `mapstructure:"HTTPWriteTimeout"`
	CertPath         string        `mapstructure:"certPath"`
	KeyPath          string        `mapstructure:"keyPath"`
}

func InitConfig() (err error) {
	Conf = NewConfig()
	viper.SetConfigName("logic")
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
			Pidfile:          "/tmp/logic.pid",
			MaxProc:          runtime.NumCPU(),
			PprofAddrs:       []string{"localhost:6922"},
			HttpAddrs:        []string{"tcp@0.0.0.0:6921"},
			RPCAddrs:         []string{"tcp@localhost:6923"},
			HTTPReadTimeout:  10 * time.Second,
			HTTPWriteTimeout: 20 * time.Second,
		},
	}
}

// 重新加载配置
// func ReloadConfig() {
//
// }
