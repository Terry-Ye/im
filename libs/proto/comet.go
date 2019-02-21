package proto



type PushMsgArg struct {
	Uid string
	P   Proto
}


type RoomMsgArg struct {
	RoomId int32
	P   Proto
}



type RoomCountArg struct {
	RoomId int32
	Count  int
}


