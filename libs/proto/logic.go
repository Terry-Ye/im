package proto



type ConnArg struct {
	Auth  string
	Server int32
}

type ConnReply struct {
	Uid    string

}
