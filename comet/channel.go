package main

import (
	"im/libs/proto"
	"github.com/gorilla/websocket"
)

type Channel struct {
	Room   *Room
	signal chan *proto.Proto
	broadcast chan []byte
	conn *websocket.Conn
}

func NewChannel(svr int) *Channel {
	c := new(Channel)
	c.signal = make(chan *proto.Proto, svr)
	return c
}
