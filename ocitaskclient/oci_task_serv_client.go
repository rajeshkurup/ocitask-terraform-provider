package ocitaskclient

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

/**
 * @brief Interface for OCI Task Service
 */
type OciTaskServClientInterface interface {
	CreateTask(ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error)
	UpdateTask(taskId *int64, ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error)
	GetTask(taskId *int64) (*OciTaskServResponse, error)
	DeleteTask(taskId *int64) (*OciTaskServResponse, error)
}

/**
 * @brief Client for OCI Task Service
 */
type OciTaskServClient struct {
	httpClient OciTaskHttpInterface
	hostUrl    *string
}

/**
 * @brief Constructor to create instance of OciTaskServClient
 * @param hostUrl Host URL to OCI Task Service
 * @return Instance of OciTaskServClient
 */
func MakeOciTaskServClient(hostUrl *string) *OciTaskServClient {
	client := MakeOciTaskHttp()
	return &OciTaskServClient{
		httpClient: &client,
		hostUrl:    hostUrl,
	}
}

/**
 * @brief Getter function for host URL
 * @return Host URL
 */
func (ociTaskServClient *OciTaskServClient) GetUrl() string {
	return *ociTaskServClient.hostUrl
}

/**
 * @brief Public method to cretae Task using OCI Task Service.
 *			Returns Task Idetifier if succeeded.
 *			Returns instance of OciError if failed.
 * @param ociTaskServRequest Request to OCI Task Service
 * @return Instance of OciTaskServResponse
 * @return Instance of error if failed
 */
func (ociTaskServClient *OciTaskServClient) CreateTask(ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error) {
	if ociTaskServRequest == nil {
		return nil, errors.New("Invalid Argument - please check Api Request")
	}

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

/**
 * @brief Public method to update Task using OCI Task Service.
 *			Returns Task Idetifier if succeeded.
 *			Returns instance of OciError if failed.
 * @param taskId Identifier of the Task
 * @param ociTaskServRequest Request to OCI Task Service
 * @return Instance of OciTaskServResponse
 * @return Instance of error if failed
 */
func (ociTaskServClient *OciTaskServClient) UpdateTask(taskId *int64, ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error) {
	if taskId == nil || ociTaskServRequest == nil {
		return nil, errors.New("Invalid Argument - please check Id or Api Request")
	}

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

/**
 * @brief Public method to read Task using OCI Task Service.
 *			Returns OciTask instance if succeeded.
 *			Returns instance of OciError if failed.
 * @param taskId Identifier of the Task
 * @return Instance of OciTaskServResponse
 * @return Instance of error if failed
 */
func (ociTaskServClient *OciTaskServClient) GetTask(taskId *int64) (*OciTaskServResponse, error) {
	if taskId == nil {
		return nil, errors.New("Invalid Argument - please check Id")
	}

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

/**
 * @brief Public method to delete Task using OCI Task Service.
 *			Returns nothing if succeeded.
 *			Returns instance of OciError if failed.
 * @param taskId Identifier of the Task
 * @return Instance of OciTaskServResponse
 * @return Instance of error if failed
 */
func (ociTaskServClient *OciTaskServClient) DeleteTask(taskId *int64) (*OciTaskServResponse, error) {
	if taskId == nil {
		return nil, errors.New("Invalid Argument - please check Id")
	}

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

/**
 * @brief Private method to build OCI Task Service HTTP request.
 * @param method HTTP Method (GET, POST, PUT or DELETE)
 * @param url HTTP URL to OCI Task Service
 * @param ociRequest Instance of OciTaskServRequest. This is optional.
 * @return Instance of http.Request if succeeded
 * @return Instance of error if failed
 */
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

/**
 * @brief Private method to send HTTP request OCI Task Service.
 * @param apiRequest Instance of http.Request
 * @return Instance of http.Response if succeeded
 * @return Instance of http.Response Body if succeeded
 * @return Instance of error if failed
 */
func (ociTaskServClient *OciTaskServClient) sendRequest(apiRequest *http.Request) (*http.Response, []byte, error) {
	apiRequest.Header.Set("Content-Type", "application/json")
	apiRequest.Header.Set("Accept", "application/json")

	apiResp, err := ociTaskServClient.httpClient.SendRequest(apiRequest)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to send request to OCI Task Management Service - error=%s", err))
		return nil, nil, err
	}

	defer apiResp.Body.Close()

	body, err := ociTaskServClient.httpClient.IoRead(apiResp.Body)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to read response from OCI Task Management Service - error=%s", err))
		return apiResp, nil, err
	}

	return apiResp, body, nil
}
