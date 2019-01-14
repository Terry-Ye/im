package proto

type ConnArg struct {
	Auth   string
	RoomId int32
	ServerId int8
}

type ConnReply struct {
	Uid string
}


type DisconnArg struct {
	RoomId int32
	Uid string
}

type DisconnReply struct {
	Has bool
}


type Send struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	FormUserId string `json:"fuid"`
	FormUserName string `json:"fname"`
	Op int32 `json:"op"`
}
