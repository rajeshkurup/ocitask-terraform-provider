package ocitaskclient

import (
	"encoding/json"
	"fmt"
	"log"
)

/**
 * @brief Container for error details
 */
type OciError struct {
	ErrorCode    *int    `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

/**
 * @brief Convert OciError object into JSON String
 * @return JSON String equivalent to OciError object if succeeded
 * @return Instance of error if failed
 */
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

/**
 * @brief Convert JSON String into OciError object
 * @param data JSON String equivalent to OciError object
 * @return Instance of error if failed
 */
func (ociError *OciError) Deserialize(data []byte) error {
	err := json.Unmarshal(data, ociError)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to deserialize OciError - error=%s", err))
	}

	return err
}
