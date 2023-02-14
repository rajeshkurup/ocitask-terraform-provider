package ocitaskclient

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stretchr/testify/assert"
)

func TestOciTaskSerializeSuccess(test *testing.T) {
	taskId := int64(1001)
	taskTitle := "Test Task"
	taskDesc := "Test Task Desc"
	taskPriority := 2
	taskCompleted := true
	taskStartDate := int64(2002)
	taskDueDate := int64(2003)
	taskTimeCreated := int64(2004)
	taskTimeUpdated := int64(2005)

	ociTask := OciTask{}
	ociTask.Completed = &taskCompleted
	ociTask.Description = &taskDesc
	ociTask.DueDate = &taskDueDate
	ociTask.Id = &taskId
	ociTask.Priority = &taskPriority
	ociTask.StartDate = &taskStartDate
	ociTask.TimeCreated = &taskTimeCreated
	ociTask.TimeUpdated = &taskTimeUpdated
	ociTask.Title = &taskTitle

	dataJson, err := ociTask.Serialize()

	assert.NoError(test, err, "TestOciTaskSerializeSuccess Failed: Unable to serialize OciTask")

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(dataJson), &data)

	assert.NoError(test, err, "TestOciTaskSerializeSuccess Failed: Unable to deserialize OciError")
	assert.Equal(test, 1001, int(data["id"].(float64)), "TestOciTaskSerializeSuccess Failed: Wrong Task Id")
	assert.Equal(test, "Test Task", data["title"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Title")
	assert.Equal(test, "Test Task Desc", data["description"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Description")
	assert.Equal(test, 2, int(data["priority"].(float64)), "TestOciTaskSerializeSuccess Failed: Wrong Task Priority")
	assert.Equal(test, true, data["completed"].(bool), "TestOciTaskSerializeSuccess Failed: Wrong Task Completed")
	assert.Equal(test, 2002, int(data["startDate"].(float64)), "TestOciTaskSerializeSuccess Failed: Wrong Task Start Date")
	assert.Equal(test, 2003, int(data["dueDate"].(float64)), "TestOciTaskSerializeSuccess Failed: Wrong Task Due Date")
	assert.Equal(test, 2004, int(data["timeCreated"].(float64)), "TestOciTaskSerializeSuccess Failed: Wrong Task Time Created")
	assert.Equal(test, 2005, int(data["timeUpdated"].(float64)), "TestOciTaskSerializeSuccess Failed: Wrong Task Time Updated")
}

func TestOciTaskDeserializeSuccess(test *testing.T) {
	data := make(map[string]interface{})
	data["id"] = 1001
	data["title"] = "Test Task"
	data["description"] = "Test Task Desc"
	data["priority"] = 2
	data["completed"] = true
	data["startDate"] = 2002
	data["dueDate"] = 2003
	data["timeCreated"] = 2004
	data["timeUpdated"] = 2005

	dataJson, _ := json.Marshal(data)
	ociTask := OciTask{}

	err := ociTask.Deserialize(dataJson)

	assert.NoError(test, err, "TestOciTaskDeserializeSuccess Failed: Unable to Deserialize OciTask")
	assert.Equal(test, int64(1001), *ociTask.Id, "TestOciTaskDeserializeSuccess Failed: Wrong Task Id")
	assert.Equal(test, "Test Task", *ociTask.Title, "TestOciTaskDeserializeSuccess Failed: Wrong Task Title")
	assert.Equal(test, "Test Task Desc", *ociTask.Description, "TestOciTaskDeserializeSuccess Failed: Wrong Task Description")
	assert.Equal(test, 2, *ociTask.Priority, "TestOciTaskDeserializeSuccess Failed: Wrong Task Priority")
	assert.Equal(test, true, *ociTask.Completed, "TestOciTaskDeserializeSuccess Failed: Wrong Task Completed")
	assert.Equal(test, int64(2002), *ociTask.StartDate, "TestOciTaskDeserializeSuccess Failed: Wrong Task Start Date")
	assert.Equal(test, int64(2003), *ociTask.DueDate, "TestOciTaskDeserializeSuccess Failed: Wrong Task Due Date")
	assert.Equal(test, int64(2004), *ociTask.TimeCreated, "TestOciTaskDeserializeSuccess Failed: Wrong Task Time Created")
	assert.Equal(test, int64(2005), *ociTask.TimeUpdated, "TestOciTaskDeserializeSuccess Failed: Wrong Task Time Updated")
}

func TestOciTaskDeserializeFailed(test *testing.T) {
	data := "Test Error Message"

	dataJson, _ := json.Marshal(data)
	ociTask := OciTask{}

	err := ociTask.Deserialize(dataJson)

	assert.Error(test, err, "TestOciTaskDeserializeFailed Failed")
}

func TestFlattenOciTaskSuccess(test *testing.T) {
	taskId := int64(1001)
	taskTitle := "Test Task"
	taskDesc := "Test Task Desc"
	taskPriority := 2
	taskCompleted := true

	currTime := time.Now()
	taskStartDate := currTime.Unix() * 1000
	taskDueDate := taskStartDate
	taskTimeCreated := taskStartDate
	taskTimeUpdated := taskStartDate

	ociTask := OciTask{}
	ociTask.Completed = &taskCompleted
	ociTask.Description = &taskDesc
	ociTask.DueDate = &taskDueDate
	ociTask.Id = &taskId
	ociTask.Priority = &taskPriority
	ociTask.StartDate = &taskStartDate
	ociTask.TimeCreated = &taskTimeCreated
	ociTask.TimeUpdated = &taskTimeUpdated
	ociTask.Title = &taskTitle

	genericOciTasks, diags := FlattenOciTask(&ociTask)

	assert.Equal(test, 1, len(genericOciTasks), "TestFlattenOciTaskSuccess Failed: Wrong Response")
	assert.Equal(test, 0, len(diags), "TestFlattenOciTaskSuccess Failed: Wrong Diag Response")

	genericOciTask := genericOciTasks[0]
	dataJson, err := json.Marshal(genericOciTask)

	assert.NoError(test, err, "TestFlattenOciTaskSuccess Failed: Unable to serialize OciTask")

	data := make(map[string]interface{})
	err = json.Unmarshal(dataJson, &data)

	assert.NoError(test, err, "TestFlattenOciTaskSuccess Failed: Unable to deserialize OciTask")

	assert.Equal(test, 1001, int(data["id"].(float64)), "TestOciTaskSerializeSuccess Failed: Wrong Task Id")
	assert.Equal(test, "Test Task", data["title"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Title")
	assert.Equal(test, "Test Task Desc", data["description"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Description")
	assert.Equal(test, 2, int(data["priority"].(float64)), "TestOciTaskSerializeSuccess Failed: Wrong Task Priority")
	assert.Equal(test, true, data["completed"].(bool), "TestOciTaskSerializeSuccess Failed: Wrong Task Completed")

	taskDate := currTime.Format("yyyy-MM-dd")
	assert.Equal(test, taskDate, data["start_date"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Start Date")
	assert.Equal(test, taskDate, data["due_date"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Due Date")
	assert.Equal(test, taskDate, data["time_created"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Time Created")
	assert.Equal(test, taskDate, data["time_updated"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Time Updated")

}

func TestFlattenOciTaskFailed(test *testing.T) {

	genericOciTasks, diags := FlattenOciTask(nil)

	assert.Equal(test, 0, len(genericOciTasks), "TestFlattenOciTaskFailed Failed: Wrong Response")
	assert.Equal(test, 1, len(diags), "TestFlattenOciTaskFailed Failed: Wrong Diag Response")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestFlattenOciTaskFailed Failed: Wrong Diag Severity")
	assert.Equal(test, "Invalid Argument", diags[0].Summary, "TestFlattenOciTaskFailed Failed: Wrong Diag Summary")
	assert.Equal(test, "Invalid OciTask instance passed", diags[0].Detail, "TestFlattenOciTaskFailed Failed: Wrong Diag Detail")
}
