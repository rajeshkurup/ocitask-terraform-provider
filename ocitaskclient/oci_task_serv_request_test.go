package ocitaskclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOciTaskServRequestSerializeSuccess(test *testing.T) {
	taskTitle := "Test Task"
	taskDesc := "Test Task Desc"
	taskPriority := 2
	taskCompleted := true
	taskStartDate := "2023-02-11"
	taskDueDate := "2023-02-12"

	ociTaskServRequest := OciTaskServRequest{}
	ociTaskServRequest.Completed = &taskCompleted
	ociTaskServRequest.Description = &taskDesc
	ociTaskServRequest.DueDate = &taskDueDate
	ociTaskServRequest.Priority = &taskPriority
	ociTaskServRequest.StartDate = &taskStartDate
	ociTaskServRequest.Title = &taskTitle

	dataJson, err := ociTaskServRequest.Serialize()

	assert.NoError(test, err, "TestOciTaskServRequestSerializeSuccess Failed: Unable to serialize OciTaskServRequest")

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(dataJson), &data)

	assert.NoError(test, err, "TestOciTaskServRequestSerializeSuccess Failed: Unable to deserialize OciTaskServRequest")
	assert.Equal(test, "Test Task", data["title"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Title")
	assert.Equal(test, "Test Task Desc", data["description"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Description")
	assert.Equal(test, 2, int(data["priority"].(float64)), "TestOciTaskSerializeSuccess Failed: Wrong Task Priority")
	assert.Equal(test, true, data["completed"].(bool), "TestOciTaskSerializeSuccess Failed: Wrong Task Completed")
	assert.Equal(test, "2023-02-11", data["startDate"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Start Date")
	assert.Equal(test, "2023-02-12", data["dueDate"].(string), "TestOciTaskSerializeSuccess Failed: Wrong Task Due Date")
}

func TestOciTaskServRequestDeserializeSuccess(test *testing.T) {
	data := make(map[string]interface{})
	data["title"] = "Test Task"
	data["description"] = "Test Task Desc"
	data["priority"] = 2
	data["completed"] = true
	data["startDate"] = "2023-02-11"
	data["dueDate"] = "2023-02-12"

	dataJson, _ := json.Marshal(data)
	ociTaskServRequest := OciTaskServRequest{}

	err := ociTaskServRequest.Deserialize(dataJson)

	assert.NoError(test, err, "TestOciTaskServRequestDeserializeSuccess Failed: Unable to Deserialize OciTaskServRequest")
	assert.Equal(test, "Test Task", *ociTaskServRequest.Title, "TestOciTaskDeserializeSuccess Failed: Wrong Task Title")
	assert.Equal(test, "Test Task Desc", *ociTaskServRequest.Description, "TestOciTaskDeserializeSuccess Failed: Wrong Task Description")
	assert.Equal(test, 2, *ociTaskServRequest.Priority, "TestOciTaskDeserializeSuccess Failed: Wrong Task Priority")
	assert.Equal(test, true, *ociTaskServRequest.Completed, "TestOciTaskDeserializeSuccess Failed: Wrong Task Completed")
	assert.Equal(test, "2023-02-11", *ociTaskServRequest.StartDate, "TestOciTaskDeserializeSuccess Failed: Wrong Task Start Date")
	assert.Equal(test, "2023-02-12", *ociTaskServRequest.DueDate, "TestOciTaskDeserializeSuccess Failed: Wrong Task Due Date")
}

func TestOciTaskServRequestDeserializeFailed(test *testing.T) {
	data := "Test Error Message"

	dataJson, _ := json.Marshal(data)
	ociTaskServRequest := OciTaskServRequest{}

	err := ociTaskServRequest.Deserialize(dataJson)

	assert.Error(test, err, "TestOciTaskServRequestDeserializeFailed Failed")
}

func TestMakeOciTaskServRequestSuccess(test *testing.T) {
	data := make(map[string]interface{})
	data["title"] = "Test Task"
	data["description"] = "Test Task Desc"
	data["priority"] = 2
	data["completed"] = true
	data["start_date"] = "2023-02-11"
	data["due_date"] = "2023-02-12"

	var iData interface{} = data

	ociTaskServRequest, err := MakeOciTaskServRequest(&iData)

	assert.NoError(test, err, "TestMakeOciTaskServRequestSuccess Failed: Failed to create OciTaskServRequest")
	assert.NotNil(test, ociTaskServRequest, "TestMakeOciTaskServRequestSuccess Failed: Unable to create OciTaskServRequest")
	assert.Equal(test, "Test Task", *ociTaskServRequest.Title, "TestMakeOciTaskServRequestSuccess Failed: Wrong Task Title")
	assert.Equal(test, "Test Task Desc", *ociTaskServRequest.Description, "TestMakeOciTaskServRequestSuccess Failed: Wrong Task Description")
	assert.Equal(test, 2, *ociTaskServRequest.Priority, "TestMakeOciTaskServRequestSuccess Failed: Wrong Task Priority")
	assert.Equal(test, true, *ociTaskServRequest.Completed, "TestMakeOciTaskServRequestSuccess Failed: Wrong Task Completed")
	assert.Equal(test, "2023-02-11", *ociTaskServRequest.StartDate, "TestMakeOciTaskServRequestSuccess Failed: Wrong Task Start Date")
	assert.Equal(test, "2023-02-12", *ociTaskServRequest.DueDate, "TestMakeOciTaskServRequestSuccess Failed: Wrong Task Due Date")
}

func TestMakeOciTaskServRequestFailedIvalidArgument(test *testing.T) {
	ociTaskServRequest, err := MakeOciTaskServRequest(nil)

	assert.Error(test, err, "TestMakeOciTaskServRequestFailedIvalidArgument Failed: Error expected")
	assert.Nil(test, ociTaskServRequest, "TestMakeOciTaskServRequestFailedIvalidArgument Failed: No response expected")
}
