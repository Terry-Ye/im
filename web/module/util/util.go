package util

import (


	"github.com/satori/go.uuid"
	"strings"
	"encoding/json"
)

func GenUuid() string {
	uuidStr := uuid.Must(uuid.NewV4()).String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	uuidByt := []rune(uuidStr)
	return string(uuidByt[8:24])
}

func JsonEncode(structModel interface{}) (string, error) {
	jsonStr, err := json.Marshal(structModel)
	return string(jsonStr), err
}
func JsonDecode(jsonStr string, structModel interface{}) error {
	decode := json.NewDecoder(strings.NewReader(jsonStr))
	err := decode.Decode(structModel)
	return err
}