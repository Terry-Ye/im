package main

import (
	"im/libs/proto"
	log "github.com/sirupsen/logrus"
)

type Operator interface {
	Connect(*proto.ConnArg) (string, error)
}

type DefaultOperator struct {
}

func (operator *DefaultOperator) Connect(connArg *proto.ConnArg) (uid string, err error) {
	log.Infof("Operator uid %s:", uid)
	// var connReply *proto.ConnReply
	uid, err = connect(connArg)
	return
}
