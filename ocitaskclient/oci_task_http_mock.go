package ocitaskclient

import (
	"io"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type OciTaskHttpMock struct {
	mock.Mock
}

func (ociTaskHttpMock *OciTaskHttpMock) SendRequest(apiRequest *http.Request) (*http.Response, error) {
	args := ociTaskHttpMock.Called(apiRequest)
	if args.Get(0) != nil {
		return args.Get(0).(*http.Response), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (ociTaskHttpMock *OciTaskHttpMock) IoRead(buffer io.Reader) ([]byte, error) {
	args := ociTaskHttpMock.Called(buffer)
	if args.Get(0) != nil {
		return args.Get(0).([]byte), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}
