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
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 设置cpu 核数
	runtime.GOMAXPROCS(Conf.Base.MaxProc)
	// 使用logrus包

	// log.Info("111 noteworthy happened!")
	// 加入性能监控
	perf.Init(Conf.Base.PprofBind)

	if err := InitLogicRpc(); err != nil {
		log.Panicf(fmt.Errorf("InitLogicRpc Fatal error: %s \n", err))
	}

	// new server
	Buckets := make([]*Bucket, Conf.Bucket.Num)

	for i := 0; i < Conf.Bucket.Num; i++ {
		Buckets[i] = NewBucket(BucketOptions{
			ChannelSize: Conf.Bucket.Channel,
			RoomSize:    Conf.Bucket.Room,
		})
	}
	DefaultServer = NewServer(Buckets, ServerOptions{
		WriteWait:       Conf.Base.WriteWait,
		PongWait:        Conf.Base.PongWait,
		PingPeriod:      Conf.Base.PingPeriod,
		MaxMessageSize:  Conf.Base.MaxMessageSize,
		ReadBufferSize:  Conf.Base.ReadBufferSize,
		WriteBufferSize: Conf.Base.WriteBufferSize,
		BroadcastSize: Conf.Base.BroadcastSize,
	})

	// log.Infof("server %v", DefaultServer)
	// log.Panicf("buckets :%v", buckets)

	if err := InitWebsocket(Conf.Websocket.Bind); err != nil {
		log.Fatal(err)
	}
	log.Infof("size2: %d",DefaultServer.Options.ReadBufferSize)


}
