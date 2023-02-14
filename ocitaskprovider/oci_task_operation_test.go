package ocitaskprovider

import (
	"errors"
	"ocitaskclient"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func TestCreateTaskOperationFailedEmptyItems(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	srcTasks := make([]interface{}, 0)

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})
	testData["items"] = srcTasks

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)

	diags := ociTaskOperation.OciTaskCreate(nil, rd, &ociTaskServClientMock)

	assert.Equal(test, 1, len(diags), "TestCreateTaskOperationFailedEmptyItems Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestCreateTaskOperationFailedEmptyItems Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Invalid Argument - No tasks found", diags[0].Summary, "TestCreateTaskOperationFailedEmptyItems Failed: Wrong Diagnostic Summary expected")
	assert.Equal(test, "No Items found in incoming Resource data", diags[0].Detail, "TestCreateTaskOperationFailedEmptyItems Failed: Wrong Diagnostic Detail expected")
}

func TestCreateTaskOperationFailedCreateTask(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	title := "Test Task 1"
	desc := "Test Task 1 Desc"
	priority := 5
	completed := true
	startDate := time.Now().Unix() * 1000
	dueDate := startDate
	timeCreated := startDate
	timeUpdated := startDate

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

	ociTaskServClientMock.On("CreateTask", mock.Anything).Return(nil, errors.New("Create Task Failed")).Once()

	diags := ociTaskOperation.OciTaskCreate(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 1, len(diags), "TestCreateTaskOperationFailedCreateTask Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestCreateTaskOperationFailedCreateTask Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to create task", diags[0].Summary, "TestCreateTaskOperationFailedCreateTask Failed: Wrong Diagnostic Summary expected")
	assert.Equal(test, "Create Task Failed", diags[0].Detail, "TestCreateTaskOperationFailedCreateTask Failed: Wrong Diagnostic Detail expected")
}

func TestCreateTaskOperationFailedBadResponse(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	errCode := 501
	errMsg := "Internal Error"
	ociErr := ocitaskclient.OciError{}
	ociErr.ErrorCode = &errCode
	ociErr.ErrorMessage = &errMsg

	createResponse := ocitaskclient.OciTaskServResponse{}
	createResponse.Err = &ociErr

	title := "Test Task 1"
	desc := "Test Task 1 Desc"
	priority := 5
	completed := true
	startDate := time.Now().Unix() * 1000
	dueDate := startDate
	timeCreated := startDate
	timeUpdated := startDate

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

	diags := ociTaskOperation.OciTaskCreate(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 1, len(diags), "TestCreateTaskOperationFailedBadResponse Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestCreateTaskOperationFailedBadResponse Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to create task", diags[0].Summary, "TestCreateTaskOperationFailedBadResponse Failed: Wrong Diagnostic Summary expected")

	apiErr, _ := ociErr.Serialize()

	assert.Equal(test, apiErr, diags[0].Detail, "TestCreateTaskOperationFailedBadResponse Failed: Wrong Diagnostic Detail expected")
}

func TestUpdateTaskOperationSuccess(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)
	updateResponse := ocitaskclient.OciTaskServResponse{}
	updateResponse.TaskId = &taskId

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
	srcTask["id"] = 1001
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
	rd.SetId("1001")

	ociTaskServClientMock.On("UpdateTask", updateResponse.TaskId, mock.Anything).Return(&updateResponse, nil).Once()
	ociTaskServClientMock.On("GetTask", updateResponse.TaskId).Return(&readResponse, nil).Once()

	diags := ociTaskOperation.OciTaskUpdate(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 0, len(diags), "TestUpdateTaskOperationSuccess Failed: No Diagnostics expected")
	assert.Equal(test, "1001", rd.Id(), "TestUpdateTaskOperationSuccess Failed: Task Id doesn't match with Task Id in Resource Data")

	items := rd.Get("items").([]interface{})

	assert.NotNil(test, items, "TestUpdateTaskOperationSuccess Failed: Valid list of Tasks expected")

	destTask := items[0].(map[string]interface{})

	assert.NotNil(test, destTask, "TestUpdateTaskOperationSuccess Failed: Valid Task expected")

	assert.Equal(test, 1001, destTask["id"].(int), "TestUpdateTaskOperationSuccess Failed: Task Id doesn't match with Task Id after Read")
	assert.Equal(test, title, destTask["title"].(string), "TestUpdateTaskOperationSuccess Failed: Task Title doesn't match with Task Title after Read")
	assert.Equal(test, desc, destTask["description"].(string), "TestUpdateTaskOperationSuccess Failed: Task Description doesn't match with Task Description after Read")
	assert.Equal(test, priority, destTask["priority"].(int), "TestUpdateTaskOperationSuccess Failed: Task Priority doesn't match with Task Priority after Read")
	assert.Equal(test, completed, destTask["completed"].(bool), "TestUpdateTaskOperationSuccess Failed: Task Completed doesn't match with Task Completed after Read")
	assert.Equal(test, srcTask["start_date"], destTask["start_date"].(string), "TestUpdateTaskOperationSuccess Failed: Task StartDate doesn't match with Task StartDate after Read")
	assert.Equal(test, srcTask["due_date"], destTask["due_date"].(string), "TestUpdateTaskOperationSuccess Failed: Task DueDate doesn't match with Task DueDate after Read")
	assert.Equal(test, srcTask["time_updated"], destTask["time_updated"].(string), "TestUpdateTaskOperationSuccess Failed: Task TimeUpdated doesn't match with Task TimeUpdated after Read")
	assert.Equal(test, srcTask["time_created"], destTask["time_created"].(string), "TestUpdateTaskOperationSuccess Failed: Task TimeCreated doesn't match with Task TimeCreated after Read")
}

func TestUpdateTaskOperationFailedEmptyItems(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	srcTasks := make([]interface{}, 0)

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})
	testData["items"] = srcTasks

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)
	rd.SetId("1001")

	diags := ociTaskOperation.OciTaskUpdate(nil, rd, &ociTaskServClientMock)

	assert.Equal(test, 1, len(diags), "TestUpdateTaskOperationFailedEmptyItems Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestUpdateTaskOperationFailedEmptyItems Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Invalid Argument - No tasks found", diags[0].Summary, "TestUpdateTaskOperationFailedEmptyItems Failed: Wrong Diagnostic Summary expected")
	assert.Equal(test, "No Items found in incoming Resource data", diags[0].Detail, "TestUpdateTaskOperationFailedEmptyItems Failed: Wrong Diagnostic Detail expected")
}

func TestUpdateTaskOperationFailedNoId(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	srcTasks := make([]interface{}, 0)

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})
	testData["items"] = srcTasks

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)

	diags := ociTaskOperation.OciTaskUpdate(nil, rd, &ociTaskServClientMock)

	assert.Equal(test, 1, len(diags), "TestUpdateTaskOperationFailedNoId Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestUpdateTaskOperationFailedNoId Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to get Id from resource data", diags[0].Summary, "TestUpdateTaskOperationFailedNoId Failed: Wrong Diagnostic Summary expected")
}

func TestCreateTaskOperationFailedUpdateTask(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)
	title := "Test Task 1"
	desc := "Test Task 1 Desc"
	priority := 5
	completed := true
	startDate := time.Now().Unix() * 1000
	dueDate := startDate
	timeCreated := startDate
	timeUpdated := startDate

	srcTask := make(map[string]interface{})
	srcTask["id"] = 1001
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
	rd.SetId("1001")

	ociTaskServClientMock.On("UpdateTask", &taskId, mock.Anything).Return(nil, errors.New("Update Task Failed")).Once()

	diags := ociTaskOperation.OciTaskUpdate(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 1, len(diags), "TestCreateTaskOperationFailedUpdateTask Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestCreateTaskOperationFailedUpdateTask Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to update task", diags[0].Summary, "TestCreateTaskOperationFailedUpdateTask Failed: Wrong Diagnostic Summary expected")
	assert.Equal(test, "Update Task Failed", diags[0].Detail, "TestCreateTaskOperationFailedUpdateTask Failed: Wrong Diagnostic Detail expected")
}

func TestUpdateTaskOperationFailedBadResponse(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	errCode := 501
	errMsg := "Internal Error"
	ociErr := ocitaskclient.OciError{}
	ociErr.ErrorCode = &errCode
	ociErr.ErrorMessage = &errMsg

	createResponse := ocitaskclient.OciTaskServResponse{}
	createResponse.Err = &ociErr

	taskId := int64(1001)
	title := "Test Task 1"
	desc := "Test Task 1 Desc"
	priority := 5
	completed := true
	startDate := time.Now().Unix() * 1000
	dueDate := startDate
	timeCreated := startDate
	timeUpdated := startDate

	srcTask := make(map[string]interface{})
	srcTask["id"] = 1001
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
	rd.SetId("1001")

	ociTaskServClientMock.On("UpdateTask", &taskId, mock.Anything).Return(&createResponse, nil).Once()

	diags := ociTaskOperation.OciTaskUpdate(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 1, len(diags), "TestUpdateTaskOperationFailedBadResponse Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestUpdateTaskOperationFailedBadResponse Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to update task", diags[0].Summary, "TestUpdateTaskOperationFailedBadResponse Failed: Wrong Diagnostic Summary expected")

	apiErr, _ := ociErr.Serialize()

	assert.Equal(test, apiErr, diags[0].Detail, "TestUpdateTaskOperationFailedBadResponse Failed: Wrong Diagnostic Detail expected")
}

func TestReadTaskOperationSuccess(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)
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

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)
	rd.SetId("1001")

	ociTaskServClientMock.On("GetTask", &taskId).Return(&readResponse, nil).Once()

	diags := ociTaskOperation.OciTaskRead(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 0, len(diags), "TestReadTaskOperationSuccess Failed: No Diagnostics expected")
	assert.Equal(test, "1001", rd.Id(), "TestReadTaskOperationSuccess Failed: Task Id doesn't match with Task Id in Resource Data")

	items := rd.Get("items").([]interface{})

	assert.NotNil(test, items, "TestReadTaskOperationSuccess Failed: Valid list of Tasks expected")

	destTask := items[0].(map[string]interface{})

	assert.NotNil(test, destTask, "TestReadTaskOperationSuccess Failed: Valid Task expected")

	assert.Equal(test, 1001, destTask["id"].(int), "TestReadTaskOperationSuccess Failed: Task Id doesn't match with Task Id after Read")
	assert.Equal(test, title, destTask["title"].(string), "TestReadTaskOperationSuccess Failed: Task Title doesn't match with Task Title after Read")
	assert.Equal(test, desc, destTask["description"].(string), "TestReadTaskOperationSuccess Failed: Task Description doesn't match with Task Description after Read")
	assert.Equal(test, priority, destTask["priority"].(int), "TestReadTaskOperationSuccess Failed: Task Priority doesn't match with Task Priority after Read")
	assert.Equal(test, completed, destTask["completed"].(bool), "TestReadTaskOperationSuccess Failed: Task Completed doesn't match with Task Completed after Read")
	assert.Equal(test, time.UnixMilli(startDate).Format("yyyy-MM-dd"), destTask["start_date"].(string), "TestReadTaskOperationSuccess Failed: Task StartDate doesn't match with Task StartDate after Read")
	assert.Equal(test, time.UnixMilli(dueDate).Format("yyyy-MM-dd"), destTask["due_date"].(string), "TestReadTaskOperationSuccess Failed: Task DueDate doesn't match with Task DueDate after Read")
	assert.Equal(test, time.UnixMilli(timeUpdated).Format("yyyy-MM-dd"), destTask["time_updated"].(string), "TestReadTaskOperationSuccess Failed: Task TimeUpdated doesn't match with Task TimeUpdated after Read")
	assert.Equal(test, time.UnixMilli(timeCreated).Format("yyyy-MM-dd"), destTask["time_created"].(string), "TestReadTaskOperationSuccess Failed: Task TimeCreated doesn't match with Task TimeCreated after Read")
}

func TestReadTaskOperationFailedBadId(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)

	diags := ociTaskOperation.OciTaskRead(nil, rd, &ociTaskServClientMock)

	assert.Equal(test, 1, len(diags), "TestReadTaskOperationFailedBadId Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestReadTaskOperationFailedBadId Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to get Id from resource data", diags[0].Summary, "TestReadTaskOperationFailedBadId Failed: Wrong Diagnostic Summary expected")
}

func TestReadTaskOperationFailedGetTask(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)
	rd.SetId("1001")

	ociTaskServClientMock.On("GetTask", &taskId).Return(nil, errors.New("Get Task Failed")).Once()

	diags := ociTaskOperation.OciTaskRead(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 1, len(diags), "TestReadTaskOperationFailedGetTask Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestReadTaskOperationFailedGetTask Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to read task", diags[0].Summary, "TestReadTaskOperationFailedGetTask Failed: Wrong Diagnostic Summary expected")
	assert.Equal(test, "Get Task Failed", diags[0].Detail, "TestReadTaskOperationFailedGetTask Failed: Wrong Diagnostic Detail expected")
}

func TestReadTaskOperationFailedBadResponse(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)
	errCode := 501
	errMsg := "Internal Error"
	ociErr := ocitaskclient.OciError{}
	ociErr.ErrorCode = &errCode
	ociErr.ErrorMessage = &errMsg

	readResponse := ocitaskclient.OciTaskServResponse{}
	readResponse.Err = &ociErr

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)
	rd.SetId("1001")

	ociTaskServClientMock.On("GetTask", &taskId).Return(&readResponse, nil).Once()

	diags := ociTaskOperation.OciTaskRead(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 1, len(diags), "TestReadTaskOperationFailedBadResponse Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestReadTaskOperationFailedBadResponse Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to read task", diags[0].Summary, "TestReadTaskOperationFailedBadResponse Failed: Wrong Diagnostic Summary expected")

	apiErr, _ := ociErr.Serialize()

	assert.Equal(test, apiErr, diags[0].Detail, "TestReadTaskOperationFailedBadResponse Failed: Wrong Diagnostic Detail expected")
}

func TestReadTaskOperationFailedFlatten(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)

	readResponse := ocitaskclient.OciTaskServResponse{}

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)
	rd.SetId("1001")

	ociTaskServClientMock.On("GetTask", &taskId).Return(&readResponse, nil).Once()

	diags := ociTaskOperation.OciTaskRead(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 1, len(diags), "TestReadTaskOperationFailedFlatten Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestReadTaskOperationFailedFlatten Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Invalid Argument", diags[0].Summary, "TestReadTaskOperationFailedFlatten Failed: Wrong Diagnostic Summary expected")
	assert.Equal(test, "Invalid OciTask instance passed", diags[0].Detail, "TestReadTaskOperationFailedFlatten Failed: Wrong Diagnostic Detail expected")
}

func TestDeleteTaskOperationSuccess(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)

	deleteResponse := ocitaskclient.OciTaskServResponse{}

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)
	rd.SetId("1001")

	ociTaskServClientMock.On("DeleteTask", &taskId).Return(&deleteResponse, nil).Once()

	diags := ociTaskOperation.OciTaskDelete(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 0, len(diags), "TestDeleteTaskOperationSuccess Failed: No Diagnostics expected")
}

func TestDeleteTaskOperationFailedBadId(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)

	diags := ociTaskOperation.OciTaskDelete(nil, rd, &ociTaskServClientMock)

	assert.Equal(test, 1, len(diags), "TestDeleteTaskOperationFailedBadId Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestDeleteTaskOperationFailedBadId Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to get Id from resource data", diags[0].Summary, "TestDeleteTaskOperationFailedBadId Failed: Wrong Diagnostic Summary expected")
}

func TestDeleteTaskOperationFailedDeleteTask(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)
	rd.SetId("1001")

	ociTaskServClientMock.On("DeleteTask", &taskId).Return(nil, errors.New("Failed to delete task")).Once()

	diags := ociTaskOperation.OciTaskDelete(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 1, len(diags), "TestDeleteTaskOperationFailedDeleteTask Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestDeleteTaskOperationFailedDeleteTask Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to delete task", diags[0].Summary, "TestDeleteTaskOperationFailedDeleteTask Failed: Wrong Diagnostic Summary expected")
	assert.Equal(test, "Failed to delete task", diags[0].Detail, "TestDeleteTaskOperationFailedDeleteTask Failed: Wrong Diagnostic Detail expected")
}

func TestDeleteTaskOperationFailedBadResponse(test *testing.T) {
	ociTaskServClientMock := ocitaskclient.OciTaskServClientMock{}
	ociTaskOperation := OciTaskOperation{}

	taskId := int64(1001)
	errCode := 501
	errMsg := "Internal Error"
	ociErr := ocitaskclient.OciError{}
	ociErr.ErrorCode = &errCode
	ociErr.ErrorMessage = &errMsg

	deleteResponse := ocitaskclient.OciTaskServResponse{}
	deleteResponse.Err = &ociErr

	ociTaskResource := MakeOciTaskResource()
	testSchema := ociTaskResource.ResourceOciTask()

	testData := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema.Schema, testData)
	rd.SetId("1001")

	ociTaskServClientMock.On("DeleteTask", &taskId).Return(&deleteResponse, nil).Once()

	diags := ociTaskOperation.OciTaskDelete(nil, rd, &ociTaskServClientMock)

	ociTaskServClientMock.AssertExpectations(test)

	assert.Equal(test, 1, len(diags), "TestDeleteTaskOperationFailedBadResponse Failed: One Diagnostic instance expected")
	assert.Equal(test, diag.Error, diags[0].Severity, "TestDeleteTaskOperationFailedBadResponse Failed: Wrong Diagnostic Severity expected")
	assert.Equal(test, "Failed to delete task", diags[0].Summary, "TestDeleteTaskOperationFailedBadResponse Failed: Wrong Diagnostic Summary expected")

	apiErr, _ := ociErr.Serialize()

	assert.Equal(test, apiErr, diags[0].Detail, "TestDeleteTaskOperationFailedBadResponse Failed: Wrong Diagnostic Detail expected")
}
