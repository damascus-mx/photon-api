package core

import (
	"encoding/json"
)

// ResponseModel Model for HTTP responses
type ResponseModel struct {
	Message string `json:"message"`
}

// BuildMessage Builds a message for system messages
func BuildMessage(message []byte) ([]byte, error) {
	message, err := json.Marshal(ResponseModel{Message: string(message)})

	if err != nil {
		return nil, err
	}

	return message, nil
}
