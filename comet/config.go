package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/spf13/viper"
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
	Base      BaseConf      `mapstructure:"base"`
	Websocket WebsocketConf `mapstructure:"websocket"`
	Bucket    BucketConf    `mapstructure:"bucket"`
}

// 基础的配置信息
type BaseConf struct {
	Pidfile         string `mapstructure:"pidfile"`
	MaxProc         int
	PprofBind       []string `mapstructure:"pprofBind"` // 性能监控的域名端口
	Logfile         string   `mapstructure:"logfile"`   // log 文件
	WriteWait       time.Duration
	PongWait        time.Duration
	PingPeriod      time.Duration
	MaxMessageSize  int64
	BroadcastSize int
	ReadBufferSize  int
	WriteBufferSize int
}

type BucketConf struct {
	Num      int `mapstructure:"num"`
	Channel  int `mapstructure:"channel"`
	Room     int `mapstructure:"room"`
	SvrProto int `mapstructure:"svrProto"`
}

type WebsocketConf struct {
	Bind string `mapstructure:"bind"` // 性能监控的域名端口
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
			Pidfile:         "/tmp/comet.pid",
			Logfile:         "/Users/AT/go/src/im/logs/comet/comet.log",
			MaxProc:         runtime.NumCPU(),
			PprofBind:       []string{"localhost:7911"},
			WriteWait:       10 * time.Second,
			PongWait:        60 * time.Second,
			PingPeriod:      54 * time.Second,
			MaxMessageSize:  512,
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			BroadcastSize: 512,
		},
		Bucket: BucketConf{
			Num:      256,
			Channel:  1024,
			Room:     1024,
			SvrProto: 80,
		},
		Websocket: WebsocketConf{
			Bind: ":7911",
		},
	}
}
