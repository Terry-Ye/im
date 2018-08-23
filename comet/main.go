package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU()) // 后续写入配置文件
	if err := InitConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Printf("%v\n", Conf)

}
