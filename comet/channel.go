package main

import "goim/libs/proto"

type Channel struct {
	Room   *Room
	signal chan *proto.Proto
}

func NewChannel(svr int) *Channel {
	c := new(Channel)
	c.signal = make(chan *proto.Proto, svr)
	return c
}
