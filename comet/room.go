package main

import (
	"sync"
	"im/libs/proto"
)

type Room struct {
	Id     int32 // 房间号
	rlock  sync.RWMutex
	next   *Channel // 该房间的所有客户端的Channel
	drop   bool // 标示房间是否存活
	Online int  // 在线用户数量
}

func NewRoom(Id int32) (r *Room) {
	r = new(Room)
	r.Id = Id
	r.drop = false
	r.next = nil
	r.Online = 0
	return

}

func (r *Room) Put(ch *Channel) (err error) {
	if !r.drop {
		if r.next != nil {
			r.next.Prev = ch

		}
		ch.Next = r.next
		ch.Prev = nil
		r.next = ch
		r.Online++
	}else {
		err = ErrRoomDroped
	}
	return
}

func (r *Room) Push(p *proto.Proto) {
	r.rlock.RLock()
	for ch := r.next; ch != nil; ch = ch.Next{
		ch.Push(p)
	}
	r.rlock.RUnlock()
	return
}
