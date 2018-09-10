package main

import "github.com/gorilla/websocket"

type Channel struct {
	Room      *Room
	broadcast chan []byte
	uid       string
	conn      *websocket.Conn
	Next      *Channel
	Prev      *Channel
}

func NewChannel(svr int) *Channel {
	c := new(Channel)
	c.broadcast = make(chan []byte, svr)
	c.Next = nil
	c.Prev = nil
	return c
}
