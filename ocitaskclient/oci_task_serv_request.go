package ocitaskclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

/**
 * @brief Request container for OCI Task Service
 */
type OciTaskServRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
	StartDate   *string `json:"startDate,omitempty"`
	DueDate     *string `json:"dueDate,omitempty"`
}

/**
 * @brief Constructor for OciTaskServRequest
 * @param srcOciTask Instance of OciTask
 * @return Instnace of OciTaskServRequest if succeeded
 * @return Instance of error if failed
 */
func MakeOciTaskServRequest(srcOciTask *interface{}) (*OciTaskServRequest, error) {
	if srcOciTask == nil {
		errMsg := "Invalid Argment: Invalid OCI Task passed"
		log.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	ociTask := (*srcOciTask).(map[string]interface{})
	title := ociTask["title"].(string)
	description := ociTask["description"].(string)
	priority := ociTask["priority"].(int)
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

/**
 * @brief Convert OciTaskServRequest object into JSON String
 * @return JSON String equivalent to OciTaskServRequest object if succeeded
 * @return Instance of error if failed
 */
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

/**
 * @brief Convert JSON String into OciTaskServRequest object
 * @param data JSON String equivalent to OciTaskServRequest object
 * @return Instance of error if failed
 */
func (ociTaskServRequest *OciTaskServRequest) Deserialize(data []byte) error {
	err := json.Unmarshal(data, ociTaskServRequest)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to deserialize OciTaskServRequest - error=%s", err))
	}

	return err
}
