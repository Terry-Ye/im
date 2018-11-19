package proto

type ConnArg struct {
	Auth   string
	RoomId int32
	ServerId int8
}

type ConnReply struct {
	Uid string

}
