package ocitaskclient

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

/**
 * @brief Container for Task resource in OCI Task System
 */
type OciTask struct {
	Id          *int64  `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
	StartDate   *int64  `json:"startDate,omitempty"`
	DueDate     *int64  `json:"dueDate,omitempty"`
	TimeUpdated *int64  `json:"timeUpdated,omitempty"`
	TimeCreated *int64  `json:"timeCreated,omitempty"`
}

/**
 * @brief Convert OciTask instance into generic Interface object
 * @param srcTask Instance of OciTask to be converted
 * @param diags Instance of diag.Diagnostics array to add error details if any
 * @return Instnace of Interface array. Would contain one element equivalent to OciTask if succeeded, empty otherwise.
 * @return Instance of diag.Diagnostics array with error details if failed
 */
func FlattenOciTask(srcTask *OciTask) ([]interface{}, diag.Diagnostics) {
	items := make([]interface{}, 0)
	var diags diag.Diagnostics
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
			Summary:  "Invalid Argument",
			Detail:   "Invalid OciTask instance passed",
		})
	}

	return items, diags
}

/**
 * @brief Convert OciTask object into JSON String
 * @return JSON String equivalent to OciTask object if succeeded
 * @return Instance of error if failed
 */
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

/**
 * @brief Convert JSON String into OciTask object
 * @param data JSON String equivalent to OciTask object
 * @return Instance of error if failed
 */
func (ociTask *OciTask) Deserialize(data []byte) error {
	err := json.Unmarshal(data, ociTask)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to deserialize OciTask - error=%s", err))
	}

	return err
}
