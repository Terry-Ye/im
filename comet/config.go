package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Base BaseConf `mapstructure:"base"`
}

type BaseConf struct {
	pidfile string `mapstructure:"pidfile"`
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
			pidfile: "/tmp/comet.pid",
		},
	}
}
