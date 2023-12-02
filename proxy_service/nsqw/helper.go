package nsqw

import (
	"encoding/json"
	"log"
)

func MarshalMessage(message interface{}) []byte {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("[NSQ] Error marshaling message: %v", err)
		return nil
	}
	return messageBytes
}
