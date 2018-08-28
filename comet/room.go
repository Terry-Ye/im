package main

import (
	"sync"
)

type Room struct {
	Id     int32 // 房间号
	rlock  sync.RWMutex
	chs    []*Channel
	drop   bool // 标示房间是否存活
	Online int  // dirty read is ok  // 房间的channel数量，即房间的在线用户的多少
}

func NewRoom(Id int32) (r *Room) {
	r = new(Room)
	r.Id = Id
	r.drop = false
	r.chs = nil
	r.Online = 0
	return

}
