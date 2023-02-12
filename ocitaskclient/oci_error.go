package ocitaskclient

import (
	"encoding/json"
	"fmt"
	"log"
)

type OciError struct {
	ErrorCode    *int    `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

func (ociError *OciError) Serialize() (string, error) {
	result := ""
	data, err := json.Marshal(ociError)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to serialize OciError - error=%s", err))
	} else {
		result = string(data)
	}

	return result, err
}

func (ociError *OciError) Deserialize(data []byte) error {
	err := json.Unmarshal(data, ociError)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to deserialize OciError - error=%s", err))
	}

	return err
}
