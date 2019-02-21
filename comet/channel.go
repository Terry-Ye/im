package main

import(
	"github.com/gorilla/websocket"
	"im/libs/proto"
)


type Channel struct {
	Room      *Room
	broadcast chan *proto.Proto
	uid       string
	conn      *websocket.Conn
	Next      *Channel
	Prev      *Channel
}

func NewChannel(svr int) *Channel {
	c := new(Channel)
	c.broadcast = make(chan *proto.Proto, svr)
	c.Next = nil
	c.Prev = nil
	return c
}


func (ch *Channel) Push(p *proto.Proto) (err error){
	select {
		case ch.broadcast <- p:
		default:
	}

	return
}


