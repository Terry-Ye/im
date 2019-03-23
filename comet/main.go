package main

import (
	"flag"
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
		log.Panicf("Fatal error config file: %s \n", err)
	}
	// 设置cpu 核数
	runtime.GOMAXPROCS(Conf.Base.MaxProc)

	// 加入性能监控
	perf.Init(Conf.Base.PprofBind)

	if err := InitLogicRpcClient(); err != nil {

		log.Panicf("InitLogicRpc Fatal error: %s \n", err)
	}

	// new server
	Buckets := make([]*Bucket, Conf.Bucket.Num)

	for i := 0; i < Conf.Bucket.Num; i++ {
		Buckets[i] = NewBucket(BucketOptions{
			ChannelSize:   Conf.Bucket.Channel,
			RoomSize:      Conf.Bucket.Room,
			RoutineAmount: Conf.Bucket.RoutineAmount,
			RoutineSize:   Conf.Bucket.RoutineSize,
		})
	}
	operator := new(DefaultOperator)
	DefaultServer = NewServer(Buckets, operator, ServerOptions{
		WriteWait:       Conf.Base.WriteWait,
		PongWait:        Conf.Base.PongWait,
		PingPeriod:      Conf.Base.PingPeriod,
		MaxMessageSize:  Conf.Base.MaxMessageSize,
		ReadBufferSize:  Conf.Base.ReadBufferSize,
		WriteBufferSize: Conf.Base.WriteBufferSize,
		BroadcastSize:   Conf.Base.BroadcastSize,
	})

	log.Info("start InitPushRpc")
	if err := InitLogicRpcServer(); err != nil {
		log.Panicf("InitPushRpc Fatal error: %s \n", err)
	}
	/**
	ws
	*/
	if err := InitWebsocket(); err != nil {
		log.Panicf("InitWebsocket() error:  %s \n", err)
	}

	/**
	wss
	You need to configure certPath and keyPath in comet.toml.
	*/
	// if err := InitWebsocketWss(); err != nil {
	// 	log.Panicf("Please check the certPath and keyPath of wss or other, error:  %s \n", err)
	// }

}
