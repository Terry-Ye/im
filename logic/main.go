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



	if err := InitRPC(); err != nil {
		log.Panic(fmt.Errorf("InitRPC() fatal error : %s \n", err))
	}


	// log.Info("111 noteworthy happened!")
	// 加入监控 后补
	if err := InitRedis(); err != nil {
		log.Panic(fmt.Errorf("InitRedis() fatal error : %s \n", err))
	}
	// UserId, _ := RedisCli.HGet("im_auth_555","UserId").Result()
	//
	// log.Info("UserId ", UserId)
	if err := InitHTTP(); err != nil {
		log.Panic(fmt.Errorf("InitHttp() fatal error : %s \n", err))
	}



}
