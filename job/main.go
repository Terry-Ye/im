package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"runtime"
)

func main() {
	flag.Parse()

	if err := InitConfig(); err != nil {
		log.Panicf("InitConfig() Fatal error config file:  %s \n", err)
	}

	// 设置cpu 核数
	runtime.GOMAXPROCS(Conf.Base.MaxProc)

	// 初始化redis
	if err := InitRedis(); err != nil {
		log.Panicf("InitRedis() fatal error :  %s \n", err)
	}

	// 通过rpc初始化comet对应的 server bucket等
	if err := InitComets(); err != nil {
		log.Panicf("InitRPC() fatal error :  %s \n", err)
	}

	InitPush()
	select {}

	// 加入监控 后补

}
