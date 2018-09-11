package main

import (
	"github.com/smallnest/rpcx/client"
	"log"
)

type Arith int

type Addr struct {
	Key string
}

func InitLogicRpc() (err error) {
	// var (
	// 	rpcLogicAddrs []*addr
	//
	// )
	// rpcLogicAddrs = make([]*addr, len(Conf.Base.RpcLogicAddr))
	// for i, bind := range Conf.Base.RpcLogicAddr {
	// 	Addr := new(Addr)
	// 	Addr.Key = bind
	//
	// 	rpcLogicAddrs[i] = Addr
	// }



	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	xclient := client.NewXClient("Arith", "Mul", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	d := client.NewMultipleServersDiscovery([]*client.KVPair{rpcLogicAddrs})

}
