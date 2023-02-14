package ocitaskclient

import (
	"encoding/json"
	"fmt"
	"log"
)

/**
 * @brief Container for OCI Task Service API Response
 */
type OciTaskServResponse struct {
	TaskId *int64    `json:"taskId,omitempty"`
	Task   *OciTask  `json:"task,omitempty"`
	Err    *OciError `json:"error,omitempty"`
}

/**
 * @brief Convert OciTaskServResponse object into JSON String
 * @return JSON String equivalent to OciTaskServResponse object if succeeded
 * @return Instance of error if failed
 */
func (ociTaskServResponse *OciTaskServResponse) Serialize() (string, error) {
	result := ""
	data, err := json.Marshal(ociTaskServResponse)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to serialize OciTaskServResponse - error=%s", err))
	} else {
		result = string(data)
	}

	return result, err
}

/**
 * @brief Convert JSON String into OciTaskServResponse object
 * @param data JSON String equivalent to OciTaskServResponse object
 * @return Instance of error if failed
 */
func (ociTaskServResponse *OciTaskServResponse) Deserialize(data []byte) error {
	err := json.Unmarshal(data, ociTaskServResponse)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to deserialize OciTaskServResponse - error=%s", err))
	}

	return err
}
