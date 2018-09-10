package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
	"im/libs/perf"
	"runtime"
)

var (
	DefaultServer *Server
	// Debug bool
)

func main() {
	flag.Parse()

	if err := InitConfig(); err != nil {
		log.Panic(fmt.Errorf("Fatal error config file: %s \n", err))

	}
	// 设置cpu 核数
	runtime.GOMAXPROCS(Conf.Base.MaxProc)
	// 使用logrus包

	// log.Info("111 noteworthy happened!")
	// 加入监控 后补
	if err := InitRedis(); err != nil {
		log.Panic(fmt.Errorf("InitRedis() fatal error : %s \n", err))
	}

	if err := InitHttp(); err != nil {
		log.Panic(fmt.Errorf("InitHttp() fatal error : %s \n", err))
	}
	


}
