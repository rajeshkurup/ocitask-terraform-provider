package ocitaskclient

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type OciTask struct {
	Id          *int64  `json:"Id,omitempty"`
	Title       *string `json:"Title,omitempty"`
	Description *string `json:"Description,omitempty"`
	Priority    *int    `json:"Priority,omitempty"`
	Completed   *bool   `json:"Completed,omitempty"`
	StartDate   *int64  `json:"StartDate,omitempty"`
	DueDate     *int64  `json:"DueDate,omitempty"`
	TimeUpdated *int64  `json:"TimeUpdated,omitempty"`
	TimeCreated *int64  `json:"TimeCreated,omitempty"`
}

func FlattenOciTask(srcTask *OciTask, diags diag.Diagnostics) []interface{} {
	items := make([]interface{}, 0)
	if srcTask != nil {
		destTask := make(map[string]interface{})
		destTask["id"] = srcTask.Id
		destTask["title"] = srcTask.Title
		destTask["description"] = srcTask.Description
		destTask["priority"] = srcTask.Priority
		destTask["completed"] = srcTask.Completed

		startDate := time.UnixMilli(*srcTask.StartDate)
		destTask["start_date"] = startDate.Format("yyyy-MM-dd")

		dueDate := time.UnixMilli(*srcTask.DueDate)
		destTask["due_date"] = dueDate.Format("yyyy-MM-dd")

		timeUpdated := time.UnixMilli(*srcTask.TimeUpdated)
		destTask["time_updated"] = timeUpdated.Format("yyyy-MM-dd")

		timeCreated := time.UnixMilli(*srcTask.TimeCreated)
		destTask["time_created"] = timeCreated.Format("yyyy-MM-dd")

		items = append(items, destTask)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to read task",
			Detail:   "Empty response from OCI Task Service",
		})
	}

	return items
}

func (ociTask *OciTask) Serialize() (string, error) {
	result := ""
	data, err := json.Marshal(ociTask)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to serialize OciTask - error=%s", err))
	} else {
		result = string(data)
	}

	return result, err
}

func (ociTask *OciTask) Deserialize(data []byte) error {
	err := json.Unmarshal(data, ociTask)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to deserialize OciTask - error=%s", err))
	}

	return err
}
