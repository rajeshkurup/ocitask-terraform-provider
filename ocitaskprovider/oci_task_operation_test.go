package ocitaskprovider

import (
	"ocitaskclient"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTaskOperationSuccess(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)
	createResponse := ocitaskclient.OciTaskServResponse{}
	createResponse.TaskId = &taskId

	title := "Test Task 1"
	desc := "Test Task 1 Desc"
	priority := 5
	completed := true
	startDate := time.Now().Unix() * 1000
	dueDate := startDate
	timeCreated := startDate
	timeUpdated := startDate

	task := ocitaskclient.OciTask{}
	task.Completed = &completed
	task.Description = &desc
	task.DueDate = &dueDate
	task.Id = &taskId
	task.Priority = &priority
	task.StartDate = &startDate
	task.TimeCreated = &timeCreated
	task.TimeUpdated = &timeUpdated
	task.Title = &title

	readResponse := ocitaskclient.OciTaskServResponse{}
	readResponse.Task = &task

	srcTask := make(map[string]interface{})
	srcTask["title"] = title
	srcTask["description"] = desc
	srcTask["priority"] = priority
	srcTask["completed"] = completed
	srcTask["start_date"] = time.UnixMilli(startDate).Format("yyyy-MM-dd")
	srcTask["due_date"] = time.UnixMilli(dueDate).Format("yyyy-MM-dd")
	srcTask["time_created"] = time.UnixMilli(timeCreated).Format("yyyy-MM-dd")
	srcTask["time_updated"] = time.UnixMilli(timeUpdated).Format("yyyy-MM-dd")

	srcTasks := make([]interface{}, 0)
	srcTasks = append(srcTasks, srcTask)

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})
	testData["items"] = srcTasks

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)

	ociTaskServClientMock.On("CreateTask", mock.Anything).Return(&createResponse, nil).Once()
	ociTaskServClientMock.On("GetTask", createResponse.TaskId).Return(&readResponse, nil).Once()

	diags := ociTaskOperation.OciTaskCreate(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 0, len(diags), "TestCreateTaskOperationSuccess Failed: No Diagnostics expected")
	assert.Equal(test, "1001", rd.Id(), "TestCreateTaskOperationSuccess Failed: Task Id doesn't match with Task Id in Resource Data")

	items := rd.Get("items").([]interface{})

	assert.NotNil(test, items, "TestCreateTaskOperationSuccess Failed: Valid list of Tasks expected")

	destTask := items[0].(map[string]interface{})

	assert.NotNil(test, destTask, "TestCreateTaskOperationSuccess Failed: Valid Task expected")

	assert.Equal(test, 1001, destTask["id"].(int), "TestCreateTaskOperationSuccess Failed: Task Id doesn't match with Task Id after Read")
	assert.Equal(test, title, destTask["title"].(string), "TestCreateTaskOperationSuccess Failed: Task Title doesn't match with Task Title after Read")
	assert.Equal(test, desc, destTask["description"].(string), "TestCreateTaskOperationSuccess Failed: Task Description doesn't match with Task Description after Read")
	assert.Equal(test, priority, destTask["priority"].(int), "TestCreateTaskOperationSuccess Failed: Task Priority doesn't match with Task Priority after Read")
	assert.Equal(test, completed, destTask["completed"].(bool), "TestCreateTaskOperationSuccess Failed: Task Completed doesn't match with Task Completed after Read")
	assert.Equal(test, srcTask["start_date"], destTask["start_date"].(string), "TestCreateTaskOperationSuccess Failed: Task StartDate doesn't match with Task StartDate after Read")
	assert.Equal(test, srcTask["due_date"], destTask["due_date"].(string), "TestCreateTaskOperationSuccess Failed: Task DueDate doesn't match with Task DueDate after Read")
	assert.Equal(test, srcTask["time_updated"], destTask["time_updated"].(string), "TestCreateTaskOperationSuccess Failed: Task TimeUpdated doesn't match with Task TimeUpdated after Read")
	assert.Equal(test, srcTask["time_created"], destTask["time_created"].(string), "TestCreateTaskOperationSuccess Failed: Task TimeCreated doesn't match with Task TimeCreated after Read")
}
