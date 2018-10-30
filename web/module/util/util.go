package util

import (


	"github.com/satori/go.uuid"
	"strings"
)

func GenUuid() string {
	uuidStr := uuid.Must(uuid.NewV4()).String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	uuidByt := []rune(uuidStr)
	return string(uuidByt[8:24])
}