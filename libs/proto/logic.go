package proto

type ConnArg struct {
	Auth   string
	RoomId int32
	Server int32
}

type ConnReply struct {
	Uid string

}
