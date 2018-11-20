package proto

type RedisMsg struct {
	Op       int32 `json:"op"`
	ServerId int8   `json:"serverId,omitempty"`
	RoomId   int32  `json:"roomId,omitempty"`
	UserId   string `json:"userId,omitempty"`
	Msg      []byte `json:"msg"`

}


// msg: "132"
// op: "redis_message_single"
// rid: 1

type NoReply struct {
}

type SuccessReply struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
