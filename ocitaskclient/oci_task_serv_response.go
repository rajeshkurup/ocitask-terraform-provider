package ocitaskclient

import (
	"encoding/json"
	"fmt"
	"log"
)

type OciTaskServResponse struct {
	TaskId *int64    `json:"taskId,omitempty"`
	Task   *OciTask  `json:"task,omitempty"`
	Err    *OciError `json:"error,omitempty"`
}

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

func (ociTaskServResponse *OciTaskServResponse) Deserialize(data []byte) error {
	err := json.Unmarshal(data, ociTaskServResponse)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to deserialize OciTaskServResponse - error=%s", err))
	}

	return err
}
