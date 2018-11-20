package proto

type ConnArg struct {
	Auth   string
	RoomId int32
	ServerId int8
}

type ConnReply struct {
	Uid string

}

type Send struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	FormUserId string `json:"formUserId"`
	FormUserName string `json:"formUserName"`
}
