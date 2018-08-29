package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
	"im/libs/perf"
	"runtime"
)

func main() {
	flag.Parse()

	if err := InitConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 设置cpu 核数
	runtime.GOMAXPROCS(Conf.Base.MaxProc)
	// 使用logrus包

	// log.Info("111 noteworthy happened!")
	// 加入性能监控
	perf.Init(Conf.Base.PprofBind)
	// new server
	buckets := make([]*Bucket, Conf.Bucket.Num)
	for i := 0; i < Conf.Bucket.Num; i++ {
		buckets[i] = NewBucket(BucketOptions{
			ChannelSize: Conf.Bucket.Channel,
			RoomSize:    Conf.Bucket.Room,
		})
	}

	if err := InitWebsocket(Conf.Websocket.Bind); err != nil {
		log.Fatal(err)
	}

}
