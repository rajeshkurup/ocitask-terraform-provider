package ocitaskclient

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * @brief Interface for HTTP Client Adaptor
 */
type OciTaskHttpInterface interface {
	SendRequest(apiRequest *http.Request) (*http.Response, error)
	IoRead(buffer io.Reader) ([]byte, error)
}

/**
 * Helper to perform HTTP Operations
 */
type OciTaskHttp struct {
	httpClient *http.Client
}

/**
 * @brief Constructor for OciTaskHttp
 * @return Instance of OciTaskHttp
 */
func MakeOciTaskHttp() OciTaskHttp {
	return OciTaskHttp{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

/**
 * @brief Make HTTP call with given request
 * @param apiRequest Instance of HTTP Request
 * @return API Reponse if succeeded
 * @return Instance of error if failed
 */
func (ociTaskHttp *OciTaskHttp) SendRequest(apiRequest *http.Request) (*http.Response, error) {
	return ociTaskHttp.httpClient.Do(apiRequest)
}

/**
 * @brief Read API response content from HTTP Response buffer
 * @param buffer API response buffer
 * @return API Reponse as byte array if succeeded
 * @return Instance of error if failed
 */
func (ociTaskHttp *OciTaskHttp) IoRead(buffer io.Reader) ([]byte, error) {
	return ioutil.ReadAll(buffer)
}
