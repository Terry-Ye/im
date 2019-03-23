package main

import (
	"flag"
	"github.com/spf13/viper"
	"runtime"
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
	Base          BaseConf  `mapstructure:"base"`
	ZookeeperInfo Zookeeper `mapstructure:"zookeeper"`
	Redis         Redis     `mapstructure:"redis"`
	// CometConf []CometConf `mapstructure:"cometsAddrs"`
	// Bucket BucketConf `mapstructure:"bucket"`
}

type Zookeeper struct {
	Host            string `mapstructure:"host"`
	BasePath        string `mapstructure:"basePath"`
	ServerPathComet string `mapstructure:"serverPathComet"`
	ServerId        string `mapstructure:"ServerId"`
}

type Redis struct {
	RedisAddr      string `mapstructure:"RedisAddr"` //
	RedisPw        string `mapstructure:"redisPw"`
	RedisDefaultDB int    `mapstructure:"redisDefaultDB"`
}

// 基础的配置信息
type BaseConf struct {
	Pidfile      string `mapstructure:"pidfile"`
	MaxProc      int
	PprofAddrs   []string `mapstructure:"pprofBind"` //
	PushChan     int      `mapstructure:"pushChan"`
	PushChanSize int      `mapstructure:"pushChanSize"`
	IsDebug      bool
}

// type CometConf struct {
// 	Key  int8   `mapstructure:"key"`
// 	Addr string `mapstructure:"addr"`
// }

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
			Pidfile:      "/tmp/job.pid",
			MaxProc:      runtime.NumCPU(),
			PprofAddrs:   []string{"localhost:6922"},
			PushChan:     2,
			PushChanSize: 50,
			IsDebug:      true,
		},
		// CometConf: []CometConf{
		// 	{Key: 1, Addr: "tcp@0.0.0.0:6912"},
		// },
	}
}
