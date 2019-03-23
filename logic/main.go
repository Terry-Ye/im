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

	// 加入监控 后补
	if err := InitRedis(); err != nil {
		log.Panic(fmt.Errorf("InitRedis() fatal error : %s \n", err))
	}

	// http

	if err := InitHTTP(); err != nil {
		log.Panic(fmt.Errorf("InitHttp() fatal error : %s \n", err))
	}

	// https
	/**
	You need to configure certPath and keyPath in logic.toml.
	*/
	// if err := InitHTTPS(); err != nil {
	// 	log.Panicf("Please check the certPath and keyPath of wss or other, error:  %s \n", err)
	// }

}
