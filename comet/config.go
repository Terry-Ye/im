package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Base      BaseConf      `mapstructure:"base"`
	Websocket WebsocketConf `mapstructure:"websocket"`
	Bucket    BucketConf    `mapstructure:"bucket"`
}

// 基础的配置信息
type BaseConf struct {
	Pidfile   string `mapstructure:"pidfile"`
	MaxProc   int
	PprofBind []string `mapstructure:"pprofBind"` // 性能监控的域名端口
	Logfile   string   `mapstructure:"logfile"`   // log 文件
}

type BucketConf struct {
	Num     int `mapstructure:"num"`
	Channel int `mapstructure:"channel"`
	Room    int `mapstructure:"room"`
}

type WebsocketConf struct {
	Bind []string `mapstructure:"bind"` // 性能监控的域名端口
}

var (
	Conf     *Config
	confPath string
)

func init() {
	flag.StringVar(&confPath, "d", "./", " set logic config file path")
}

func InitConfig() (err error) {
	Conf = NewConfig()
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
			Pidfile:   "/tmp/comet.pid",
			Logfile:   "/Users/AT/go/src/im/logs/comet/comet.log",
			MaxProc:   runtime.NumCPU(),
			PprofBind: []string{"localhost:7911"},
		},
		Bucket: BucketConf{
			Num:     256,
			Channel: 1024,
			Room:    1024,
		},
	}
}
