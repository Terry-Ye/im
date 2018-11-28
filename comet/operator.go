package main

import (
	"im/libs/proto"
)

type Operator interface {
	Connect(*proto.ConnArg) (string, error)
	Disconnect(*proto.DisconnArg) (error)
}

type DefaultOperator struct {
}

func (operator *DefaultOperator) Connect(connArg *proto.ConnArg) (uid string, err error) {
	// var connReply *proto.ConnReply
	uid, err = connect(connArg)
	return
}


func (operator *DefaultOperator) Disconnect(disconnArg *proto.DisconnArg) ( err error) {


	if err = disconnect(disconnArg); err != nil {
		return
	}

	return
}