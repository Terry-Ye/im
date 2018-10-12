package proto

type RedisMsg struct {
	Op       string `json:"op"`
	ServerId int8   `json:"serverId,omitempty"`
	RoomId   int32  `json:"roomId,omitempty"`
	UserId   string `json:"userId,omitempty"`
	Msg      []byte `json:"msg"`
}

type NoReply struct {
}

type SuccessReply struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
