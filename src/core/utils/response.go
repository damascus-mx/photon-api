package core

import (
	"encoding/json"
)

// responseModel Model for HTTP responses
type responseModel struct {
	Message string `json:"message"`
}

// BuildMessage Builds a message for generic messages
func BuildMessage(message []byte) ([]byte, error) {
	message, err := json.Marshal(responseModel{Message: string(message)})

	if err != nil {
		return nil, err
	}

	return message, nil
}
