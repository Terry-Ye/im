package main

import (
	"im/libs/proto"
)

type Operator interface {
	Connect(arg *proto.ConnArg) (string, error)
}

type DefaultOperator struct {
}

func (operator *DefaultOperator) Connect(connArg proto.ConnArg) (uid string, err error) {
	uid, err = connect(connArg)
	return
}
