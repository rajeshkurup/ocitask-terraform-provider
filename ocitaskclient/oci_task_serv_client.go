package ocitaskclient

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type OciTaskServClientInterface interface {
	CreateTask(ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error)
	UpdateTask(taskId *int64, ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error)
	GetTask(taskId *int64) (*OciTaskServResponse, error)
	DeleteTask(taskId *int64) (*OciTaskServResponse, error)
}

type OciTaskServClient struct {
	httpClient *http.Client
	hostUrl    *string
}

func MakeOciTaskServClient(hostUrl *string) *OciTaskServClient {
	return &OciTaskServClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		hostUrl:    hostUrl,
	}
}

func (ociTaskServClient *OciTaskServClient) CreateTask(ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error) {
	apiRequest, err := ociTaskServClient.buildRequest("POST", fmt.Sprintf("%s/tasks", *ociTaskServClient.hostUrl), ociTaskServRequest)
	if err != nil {
		return nil, err
	}

	apiResp, body, err := ociTaskServClient.sendRequest(apiRequest)
	if err != nil {
		return nil, err
	}

	if apiResp.StatusCode != http.StatusCreated {
		errMsg := fmt.Sprintf("Create Task failed - status: %d, body: %s", apiResp.StatusCode, string(body))
		log.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	ociTaskServResponse := OciTaskServResponse{}
	errResp := ociTaskServResponse.Deserialize(body)
	return &ociTaskServResponse, errResp
}

func (ociTaskServClient *OciTaskServClient) UpdateTask(taskId *int64, ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error) {
	apiRequest, err := ociTaskServClient.buildRequest("PUT", fmt.Sprintf("%s/tasks/%d", *ociTaskServClient.hostUrl, *taskId), ociTaskServRequest)
	if err != nil {
		return nil, err
	}

	apiResp, body, err := ociTaskServClient.sendRequest(apiRequest)
	if err != nil {
		return nil, err
	}

	if apiResp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Update Task failed - status: %d, body: %s", apiResp.StatusCode, string(body))
		log.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	ociTaskServResponse := OciTaskServResponse{}
	errResp := ociTaskServResponse.Deserialize(body)
	return &ociTaskServResponse, errResp
}

func (ociTaskServClient *OciTaskServClient) GetTask(taskId *int64) (*OciTaskServResponse, error) {
	apiRequest, err := ociTaskServClient.buildRequest("GET", fmt.Sprintf("%s/tasks/%d", *ociTaskServClient.hostUrl, *taskId), nil)
	if err != nil {
		return nil, err
	}

	apiResp, body, err := ociTaskServClient.sendRequest(apiRequest)
	if err != nil {
		return nil, err
	}

	if apiResp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Get Task failed - status: %d, body: %s", apiResp.StatusCode, string(body))
		log.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	ociTaskServResponse := OciTaskServResponse{}
	errResp := ociTaskServResponse.Deserialize(body)
	return &ociTaskServResponse, errResp
}

func (ociTaskServClient *OciTaskServClient) DeleteTask(taskId *int64) (*OciTaskServResponse, error) {
	apiRequest, err := ociTaskServClient.buildRequest("DELETE", fmt.Sprintf("%s/tasks/%d", *ociTaskServClient.hostUrl, *taskId), nil)
	if err != nil {
		return nil, err
	}

	apiResp, body, err := ociTaskServClient.sendRequest(apiRequest)
	if err != nil {
		return nil, err
	}

	if apiResp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Delete Task failed - status: %d, body: %s", apiResp.StatusCode, string(body))
		log.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	ociTaskServResponse := OciTaskServResponse{}
	errResp := ociTaskServResponse.Deserialize(body)
	return &ociTaskServResponse, errResp
}

func (ociTaskServClient *OciTaskServClient) buildRequest(method string, url string, ociRequest *OciTaskServRequest) (*http.Request, error) {
	var body io.Reader = nil
	if ociRequest != nil {
		strReq, err := ociRequest.Serialize()
		if err != nil {
			return nil, err
		}

		body = strings.NewReader(strReq)
	}

	apiRequest, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to build request to OCI Task Management Service - error=%s", err))
		return nil, err
	}

	return apiRequest, nil
}

func (ociTaskServClient *OciTaskServClient) sendRequest(apiRequest *http.Request) (*http.Response, []byte, error) {
	apiRequest.Header.Set("Content-Type", "application/json")
	apiRequest.Header.Set("Accept", "application/json")

	apiResp, err := ociTaskServClient.httpClient.Do(apiRequest)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to send request to OCI Task Management Service - error=%s", err))
		return nil, nil, err
	}

	defer apiResp.Body.Close()

	body, err := ioutil.ReadAll(apiResp.Body)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to read response from OCI Task Management Service - error=%s", err))
		return apiResp, nil, err
	}

	return apiResp, body, nil
}
