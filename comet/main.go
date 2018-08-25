package main

import (
	"flag"
	"fmt"
	"runtime"
)

// func init() {

// }

func main() {
	flag.Parse()
	if err := InitConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 使用logrus包
	if err := InitLog(Conf.Base.Logfile); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	runtime.GOMAXPROCS(runtime.NumCPU()) // 后续写入配置文件

}
