package proto

import "encoding/json"

type Proto struct {
	Ver       int16           `json:"ver"`  // protocol version
	Operation int32           `json:"op"`   // operation for request
	// SeqId     int32           `json:"seq"`  // sequence number chosen by client
	Body      json.RawMessage `json:"body"` // binary body bytes(json.RawMessage is []byte)


}

