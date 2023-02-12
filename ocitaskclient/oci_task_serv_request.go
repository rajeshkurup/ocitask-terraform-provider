package ocitaskclient

import (
	"encoding/json"
	"fmt"
	"log"
)

type OciTaskServRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
	StartDate   *string `json:"startDate,omitempty"`
	DueDate     *string `json:"dueDate,omitempty"`
}

func MakeOciTaskServRequest(srcOciTask *interface{}) (*OciTaskServRequest, error) {
	data, err := json.Marshal(srcOciTask)
	if err != nil {
		log.Println(fmt.Sprintf("Invalid Argment: Failed to serialize OCI Task - error=%s", err))
		return nil, err
	}

	ociTask := make(map[string]interface{})
	err = json.Unmarshal(data, &ociTask)
	if err != nil {
		log.Println(fmt.Sprintf("Invalid Argment: Failed to deserialize OCI Task - error=%s", err))
		return nil, err
	}

	title := ociTask["title"].(string)
	description := ociTask["description"].(string)
	priority := int(ociTask["priority"].(float64))
	completed := ociTask["completed"].(bool)
	startDate := ociTask["start_date"].(string)
	dueDate := ociTask["due_date"].(string)

	return &OciTaskServRequest{
		Title:       &title,
		Description: &description,
		Priority:    &priority,
		Completed:   &completed,
		StartDate:   &startDate,
		DueDate:     &dueDate,
	}, nil
}

func (ociTaskServRequest *OciTaskServRequest) Serialize() (string, error) {
	result := ""
	data, err := json.Marshal(ociTaskServRequest)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to serialize OciTaskServRequest - error=%s", err))
	} else {
		result = string(data)
	}

	return result, err
}

func (ociTaskServRequest *OciTaskServRequest) Deserialize(data []byte) error {
	err := json.Unmarshal(data, ociTaskServRequest)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to deserialize OciTaskServRequest - error=%s", err))
	}

	return err
}
