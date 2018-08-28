package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLog(logFile string) (err error) {

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err == nil {
		Log.Out = file
	} else {
		return err
	}

	return nil

}
