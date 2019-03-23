package main

import (
	"im/libs/define"
	"im/libs/proto"
	"sync"
	"sync/atomic"
)

type BucketOptions struct {
	ChannelSize   int
	RoomSize      int
	RoutineAmount uint64
	RoutineSize   int
}
type Bucket struct {
	cLock    sync.RWMutex        // protect the channels for chs
	chs      map[string]*Channel // map sub key to a channel
	boptions BucketOptions
	// room
	rooms       map[int32]*Room // bucket room channels
	routines    []chan *proto.RoomMsgArg
	routinesNum uint64
	broadcast   chan []byte
}

func NewBucket(boptions BucketOptions) (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[string]*Channel, boptions.ChannelSize)

	b.boptions = boptions
	b.routines = make([]chan *proto.RoomMsgArg, boptions.RoutineAmount)
	b.rooms = make(map[int32]*Room, boptions.RoomSize)
	for i := uint64(0); i < b.boptions.RoutineAmount; i++ {
		c := make(chan *proto.RoomMsgArg, boptions.RoutineSize)
		b.routines[i] = c
		go b.PushRoom(c)
	}
	return
}

func (b *Bucket) Put(uid string, rid int32, ch *Channel) (err error) {
	var (
		room *Room
		ok   bool
	)
	b.cLock.Lock()

	if rid != define.NO_ROOM {
		if room, ok = b.rooms[rid]; !ok {
			room = NewRoom(rid)
			b.rooms[rid] = room
		}
		ch.Room = room
	}
	ch.uid = uid
	b.chs[uid] = ch
	b.cLock.Unlock()

	if room != nil {
		err = room.Put(ch)
	}
	return
}

func (b *Bucket) Channel(key string) (ch *Channel) {
	// 读操作的锁定和解锁
	b.cLock.RLock()
	ch = b.chs[key]
	b.cLock.RUnlock()
	return
}

func (b *Bucket) PushRoom(c chan *proto.RoomMsgArg) {
	for {
		var (
			arg  *proto.RoomMsgArg
			room *Room
		)
		arg = <-c

		if room = b.Room(arg.RoomId); room != nil {
			room.Push(&arg.P)
		}

	}

}

func (b *Bucket) delCh(ch *Channel) {
	var (
		ok   bool
		room *Room
	)
	b.cLock.RLock()

	if ch, ok = b.chs[ch.uid]; ok {
		room = b.chs[ch.uid].Room
		delete(b.chs, ch.uid)

	}
	if room != nil && room.Del(ch) {
		// if room empty delete
		room.Del(ch)
	}

	b.cLock.RUnlock()

}

// Room get a room by roomid.
func (b *Bucket) Room(rid int32) (room *Room) {
	b.cLock.RLock()
	room, _ = b.rooms[rid]
	b.cLock.RUnlock()
	return
}

func (b *Bucket) BroadcastRoom(arg *proto.RoomMsgArg) {
	// 广播消息递增id
	num := atomic.AddUint64(&b.routinesNum, 1) % b.boptions.RoutineAmount
	// log.Infof("BroadcastRoom RoomMsgArg :%s", arg)
	// log.Infof("bucket routinesNum :%d", b.routinesNum)
	b.routines[num] <- arg

}
