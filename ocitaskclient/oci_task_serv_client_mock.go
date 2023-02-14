package ocitaskclient

import (
	"github.com/stretchr/testify/mock"
)

type OciTaskServClientMock struct {
	mock.Mock
}

func (ociTaskServClientMock *OciTaskServClientMock) CreateTask(ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error) {
	args := ociTaskServClientMock.Called(ociTaskServRequest)
	if args.Get(0) != nil {
		return args.Get(0).(*OciTaskServResponse), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (ociTaskServClientMock *OciTaskServClientMock) UpdateTask(taskId *int64, ociTaskServRequest *OciTaskServRequest) (*OciTaskServResponse, error) {
	args := ociTaskServClientMock.Called(taskId, ociTaskServRequest)
	if args.Get(0) != nil {
		return args.Get(0).(*OciTaskServResponse), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (ociTaskServClientMock *OciTaskServClientMock) GetTask(taskId *int64) (*OciTaskServResponse, error) {
	args := ociTaskServClientMock.Called(taskId)
	if args.Get(0) != nil {
		return args.Get(0).(*OciTaskServResponse), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (ociTaskServClientMock *OciTaskServClientMock) DeleteTask(taskId *int64) (*OciTaskServResponse, error) {
	args := ociTaskServClientMock.Called(taskId)
	if args.Get(0) != nil {
		return args.Get(0).(*OciTaskServResponse), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}
