package main

import (
	"sync"
	"im/libs/define"
)

type BucketOptions struct {
	ChannelSize int
	RoomSize    int
	// RoutineAmount uint64
	// RoutineSize   int
}
type Bucket struct {
	cLock    sync.RWMutex        // protect the channels for chs
	chs      map[string]*Channel // map sub key to a channel
	boptions BucketOptions
	// room
	rooms map[int32]*Room // bucket room channels
	// routines    []chan *proto.BoardcastRoomArg
	routinesNum uint64
	broadcast chan []byte
}

func NewBucket(boptions BucketOptions) (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[string]*Channel, boptions.ChannelSize)

	b.boptions = boptions

	b.rooms = make(map[int32]*Room, boptions.RoomSize)
	// tmp
	b.broadcast = make(chan []byte, 256)
	return
}


func (b *Bucket) Put(key string, rid int32, ch *Channel) (err error){
	var (
		room *Room
		ok   bool
	)
	b.cLock.Lock()

	if rid != define.NO_ROOM {
		if  room, ok = b.rooms[rid]; !ok {
			room = NewRoom(rid)
			b.rooms[rid] = room
		}
		ch.Room = room
	}

	b.chs[key] = ch
	b.cLock.Unlock()

	if room != nil {
		err = room.Put(ch)
	}
	return
}

