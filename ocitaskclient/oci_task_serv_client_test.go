package ocitaskclient

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const HostUrl string = "http://localhost"

func TestCreateTaskSuccess(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 201,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	ociTaskServReq := OciTaskServRequest{}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return([]byte(strOciTaskServResp), nil).Once()

	apiResp, err := ociTaskServClient.CreateTask(&ociTaskServReq)

	httpClientMock.AssertExpectations(test)

	assert.NoError(test, err, "TestCreateTaskSuccess Failed: No error expected")
	assert.NotNil(test, apiResp, "TestCreateTaskSuccess Failed: Valid api response expected")
	assert.Equal(test, int64(1001), *apiResp.TaskId, "TestCreateTaskSuccess Failed: Task Id doesn't match with expected value")
}

func TestCreateTaskFailedBadTask(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	apiResp, err := ociTaskServClient.CreateTask(nil)

	assert.Error(test, err, "TestCreateTaskFailedBadTask Failed: Error expected")
	assert.Nil(test, apiResp, "TestCreateTaskFailedBadTask Failed: Invalid api response expected")
}

func TestCreateTaskFailedBadStatus(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 500,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	ociTaskServReq := OciTaskServRequest{}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return([]byte(strOciTaskServResp), nil).Once()

	apiResp, err := ociTaskServClient.CreateTask(&ociTaskServReq)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestCreateTaskFailedBadStatus Failed: Error expected")
	assert.Nil(test, apiResp, "TestCreateTaskFailedBadStatus Failed: Invalid api response expected")
}

func TestCreateTaskFailedSendRequest(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	ociTaskServReq := OciTaskServRequest{}

	httpClientMock.On("SendRequest", mock.Anything).Return(nil, errors.New("SendRequest Failed")).Once()

	apiResp, err := ociTaskServClient.CreateTask(&ociTaskServReq)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestCreateTaskFailedSendRequest Failed: Error expected")
	assert.Nil(test, apiResp, "TestCreateTaskFailedSendRequest Failed: No api response expected")
}

func TestCreateTaskFailedIoRead(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 201,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	ociTaskServReq := OciTaskServRequest{}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return(nil, errors.New("IoRead Failed")).Once()

	apiResp, err := ociTaskServClient.CreateTask(&ociTaskServReq)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestCreateTaskFailedIoRead Failed: Error expected")
	assert.Nil(test, apiResp, "TestCreateTaskFailedIoRead Failed: Invalid api response expected")
}

func TestUpdateTaskSuccess(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	ociTaskServReq := OciTaskServRequest{}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return([]byte(strOciTaskServResp), nil).Once()

	apiResp, err := ociTaskServClient.UpdateTask(&taskId, &ociTaskServReq)

	httpClientMock.AssertExpectations(test)

	assert.NoError(test, err, "TestUpdateTaskSuccess Failed: No error expected")
	assert.NotNil(test, apiResp, "TestUpdateTaskSuccess Failed: Valid api response expected")
	assert.Equal(test, int64(1001), *apiResp.TaskId, "TestUpdateTaskSuccess Failed: Task Id doesn't match with expected value")
}

func TestUpdateTaskFailedBadTask(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)

	apiResp, err := ociTaskServClient.UpdateTask(&taskId, nil)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestUpdateTaskSuccess Failed: Error expected")
	assert.Nil(test, apiResp, "TestUpdateTaskSuccess Failed: Invalid api response expected")
}

func TestUpdateTaskFailedBadId(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	apiResp, err := ociTaskServClient.UpdateTask(nil, nil)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestUpdateTaskFailedBadId Failed: Error expected")
	assert.Nil(test, apiResp, "TestUpdateTaskFailedBadId Failed: Invalid api response expected")
}

func TestUpdateTaskFailedBadStatus(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 500,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	ociTaskServReq := OciTaskServRequest{}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return([]byte(strOciTaskServResp), nil).Once()

	apiResp, err := ociTaskServClient.UpdateTask(&taskId, &ociTaskServReq)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestUpdateTaskFailedBadStatus Failed: Error expected")
	assert.Nil(test, apiResp, "TestUpdateTaskFailedBadStatus Failed: Invalid api response expected")
}

func TestUpdateTaskFailedSendRequest(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServReq := OciTaskServRequest{}

	httpClientMock.On("SendRequest", mock.Anything).Return(nil, errors.New("SendRequest Failed")).Once()

	apiResp, err := ociTaskServClient.UpdateTask(&taskId, &ociTaskServReq)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestUpdateTaskFailedSendRequest Failed: Error expected")
	assert.Nil(test, apiResp, "TestUpdateTaskFailedSendRequest Failed: No api response expected")
}

func TestUpdateTaskFailedIoRead(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 201,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	ociTaskServReq := OciTaskServRequest{}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return(nil, errors.New("IoRead Failed")).Once()

	apiResp, err := ociTaskServClient.UpdateTask(&taskId, &ociTaskServReq)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestUpdateTaskFailedIoRead Failed: Error expected")
	assert.Nil(test, apiResp, "TestUpdateTaskFailedIoRead Failed: Invalid api response expected")
}

func TestGetTaskSuccess(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	task := OciTask{Id: &taskId}
	ociTaskServResp := OciTaskServResponse{Task: &task}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return([]byte(strOciTaskServResp), nil).Once()

	apiResp, err := ociTaskServClient.GetTask(&taskId)

	httpClientMock.AssertExpectations(test)

	assert.NoError(test, err, "TestGetTaskSuccess Failed: No error expected")
	assert.NotNil(test, apiResp, "TestGetTaskSuccess Failed: Valid api response expected")
	assert.Equal(test, int64(1001), *apiResp.Task.Id, "TestGetTaskSuccess Failed: Task Id doesn't match with expected value")
}

func TestGetTaskFailedBadId(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	apiResp, err := ociTaskServClient.GetTask(nil)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestGetTaskFailedBadId Failed: Error expected")
	assert.Nil(test, apiResp, "TestGetTaskFailedBadId Failed: Invalid api response expected")
}

func TestGetTaskFailedBadStatus(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 500,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return([]byte(strOciTaskServResp), nil).Once()

	apiResp, err := ociTaskServClient.GetTask(&taskId)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestGetTaskFailedBadStatus Failed: Error expected")
	assert.Nil(test, apiResp, "TestGetTaskFailedBadStatus Failed: Invalid api response expected")
}

func TestGetTaskFailedSendRequest(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)

	httpClientMock.On("SendRequest", mock.Anything).Return(nil, errors.New("SendRequest Failed")).Once()

	apiResp, err := ociTaskServClient.GetTask(&taskId)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestGetTaskFailedSendRequest Failed: Error expected")
	assert.Nil(test, apiResp, "TestGetTaskFailedSendRequest Failed: No api response expected")
}

func TestGetTaskFailedIoRead(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 201,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return(nil, errors.New("IoRead Failed")).Once()

	apiResp, err := ociTaskServClient.GetTask(&taskId)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestGetTaskFailedIoRead Failed: Error expected")
	assert.Nil(test, apiResp, "TestGetTaskFailedIoRead Failed: Invalid api response expected")
}

func TestDeleteTaskSuccess(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	task := OciTask{Id: &taskId}
	ociTaskServResp := OciTaskServResponse{Task: &task}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return([]byte(strOciTaskServResp), nil).Once()

	apiResp, err := ociTaskServClient.DeleteTask(&taskId)

	httpClientMock.AssertExpectations(test)

	assert.NoError(test, err, "TestDeleteTaskSuccess Failed: No error expected")
	assert.NotNil(test, apiResp, "TestDeleteTaskSuccess Failed: Valid api response expected")
	assert.Equal(test, int64(1001), *apiResp.Task.Id, "TestDeleteTaskSuccess Failed: Task Id doesn't match with expected value")
}

func TestDeleteTaskFailedBadId(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	apiResp, err := ociTaskServClient.DeleteTask(nil)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestDeleteTaskFailedBadId Failed: Error expected")
	assert.Nil(test, apiResp, "TestDeleteTaskFailedBadId Failed: Invalid api response expected")
}

func TestDeleteTaskFailedBadStatus(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 500,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return([]byte(strOciTaskServResp), nil).Once()

	apiResp, err := ociTaskServClient.DeleteTask(&taskId)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestDeleteTaskFailedBadStatus Failed: Error expected")
	assert.Nil(test, apiResp, "TestDeleteTaskFailedBadStatus Failed: Invalid api response expected")
}

func TestDeleteTaskFailedSendRequest(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)

	httpClientMock.On("SendRequest", mock.Anything).Return(nil, errors.New("SendRequest Failed")).Once()

	apiResp, err := ociTaskServClient.DeleteTask(&taskId)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestDeleteTaskFailedSendRequest Failed: Error expected")
	assert.Nil(test, apiResp, "TestDeleteTaskFailedSendRequest Failed: No api response expected")
}

func TestDeleteTaskFailedIoRead(test *testing.T) {
	httpClientMock := OciTaskHttpMock{}
	url := HostUrl
	ociTaskServClient := OciTaskServClient{&httpClientMock, &url}

	taskId := int64(1001)
	ociTaskServResp := OciTaskServResponse{TaskId: &taskId}
	strOciTaskServResp, _ := ociTaskServResp.Serialize()

	httpResp := http.Response{
		StatusCode: 201,
		Body:       ioutil.NopCloser(strings.NewReader(strOciTaskServResp)),
	}

	httpClientMock.On("SendRequest", mock.Anything).Return(&httpResp, nil).Once()
	httpClientMock.On("IoRead", mock.Anything).Return(nil, errors.New("IoRead Failed")).Once()

	apiResp, err := ociTaskServClient.DeleteTask(&taskId)

	httpClientMock.AssertExpectations(test)

	assert.Error(test, err, "TestDeleteTaskFailedIoRead Failed: Error expected")
	assert.Nil(test, apiResp, "TestDeleteTaskFailedIoRead Failed: Invalid api response expected")
}
