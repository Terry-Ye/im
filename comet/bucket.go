package main

import (
	"sync"
)

func InitWebsocket() {
	websocket.BinaryMessage()
}

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
	rooms    map[int32]*Room // bucket room channels
	// routines    []chan *proto.BoardcastRoomArg
	// routinesNum uint64
}
