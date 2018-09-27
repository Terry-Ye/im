package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"

	"runtime"
)



func main() {
	flag.Parse()

	if err := InitConfig(); err != nil {
		log.Errorf("Fatal error config file: %s \n", err)
	}



	// 设置cpu 核数
	runtime.GOMAXPROCS(Conf.Base.MaxProc)

	// 初始化redis
	if err := InitRedis(); err != nil {
		log.Panic(fmt.Errorf("InitRedis() fatal error : %s \n", err))
	}


	// 通过rpc初始化comet对应的 server bucket等
	if err := InitComets(Conf.CometConf); err != nil {
		log.Panic(fmt.Errorf("InitRPC() fatal error : %s \n", err))
	}


	// log.Info("111 noteworthy happened!")

	// 加入监控 后补




	if err := InitHTTP(); err != nil {
		log.Panic(fmt.Errorf("InitHttp() fatal error : %s \n", err))
	}



}
