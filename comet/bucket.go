package main

import (
	"sync"
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
}

func NewBucket(boptions BucketOptions) (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[string]*Channel, boptions.ChannelSize)

	b.boptions = boptions

	b.rooms = make(map[int32]*Room, boptions.RoomSize)
	return
}


func (b *Bucket) Put(key string, ch *Channel) (err error){
	b.cLock.Lock()
	b.chs[key] = ch
	b.cLock.Unlock()
	return
}

