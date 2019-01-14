package proto

type RedisMsg struct {
	Op       int32  `json:"op"`
	ServerId int8   `json:"serverId,omitempty"`
	RoomId   int32  `json:"roomId,omitempty"`
	UserId   string `json:"userId,omitempty"`
	Msg      []byte `json:"msg"`
	Count    int    `json:"count"`
	RoomUserInfo map[string]string `json:"RoomUserInfo"`
}

type RedisRoomCount struct {
	Op     int32 `json:"op"`
	RoomId int32 `json:"roomId,omitempty"`
	Count  int   `json:"count,omitempty"`
}

type RedisRoomInfo struct {
	Op           int32  `json:"op"`
	RoomId       int32  `json:"roomId,omitempty"`
	Count        int    `json:"count,omitempty"`
	RoomUserInfo map[string]string `json:"roomUserInfo"`
}

type RedisRoomCountMsg struct {
	Count int   `json:"count,omitempty"`
	Op    int32 `json:"op"`
}

// msg: "132"
// op: "redis_message_single"
// rid: 1

type NoReply struct {
}

type SuccessReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
