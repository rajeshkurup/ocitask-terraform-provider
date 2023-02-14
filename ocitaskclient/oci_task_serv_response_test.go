package ocitaskclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOciTaskServResponseSerializeSuccess(test *testing.T) {
	taskId := int64(1001)
	ociTask := OciTask{}
	ociTask.Id = &taskId

	errorCode := 2001
	errorMessage := "Test Error Message"

	ociError := OciError{}
	ociError.ErrorCode = &errorCode
	ociError.ErrorMessage = &errorMessage

	ociTaskServResponse := OciTaskServResponse{}
	ociTaskServResponse.TaskId = &taskId
	ociTaskServResponse.Task = &ociTask
	ociTaskServResponse.Err = &ociError

	dataJson, err := ociTaskServResponse.Serialize()

	assert.NoError(test, err, "TestOciTaskServResponseSerializeSuccess Failed: Unable to serialize OciError")

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(dataJson), &data)

	assert.NoError(test, err, "TestOciTaskServResponseSerializeSuccess Failed: Unable to deserialize OciError")
	assert.Equal(test, 1001, int(data["taskId"].(float64)), "TestOciTaskServResponseSerializeSuccess Failed: Wrong Task Id")

	ociErr := data["error"].(map[string]interface{})
	assert.Equal(test, 2001, int(ociErr["errorCode"].(float64)), "TestOciTaskServResponseSerializeSuccess Failed: Wrong Error Code")
	assert.Equal(test, "Test Error Message", ociErr["errorMessage"].(string), "TestOciTaskServResponseSerializeSuccess Failed: Wrong Error Message")

	task := data["task"].(map[string]interface{})
	assert.Equal(test, 1001, int(task["id"].(float64)), "TestOciTaskServResponseSerializeSuccess Failed: Wrong Task Id in Task")
}

func TestOciTaskServResponseDeserializeSuccess(test *testing.T) {
	dataErr := make(map[string]interface{})
	dataErr["errorCode"] = 2001
	dataErr["errorMessage"] = "Test Error Message"

	dataTask := make(map[string]interface{})
	dataTask["id"] = 1001

	dataResp := make(map[string]interface{})
	dataResp["error"] = dataErr
	dataResp["taskId"] = 1001
	dataResp["task"] = dataTask

	dataJson, _ := json.Marshal(dataResp)
	ociTaskServResponse := OciTaskServResponse{}

	err := ociTaskServResponse.Deserialize(dataJson)

	assert.NoError(test, err, "TestOciTaskServResponseDeserializeSuccess Failed: Unable to Deserialize OciTaskServResponse")
	assert.Equal(test, int64(1001), *ociTaskServResponse.TaskId, "TestOciTaskServResponseDeserializeSuccess Failed: Wrong Task Id")
	assert.Equal(test, int64(1001), *ociTaskServResponse.Task.Id, "TestOciTaskServResponseDeserializeSuccess Failed: Wrong Task Id in Task")
	assert.Equal(test, 2001, *ociTaskServResponse.Err.ErrorCode, "TestOciTaskServResponseDeserializeSuccess Failed: Wrong Error Code")
	assert.Equal(test, "Test Error Message", *ociTaskServResponse.Err.ErrorMessage, "TestOciTaskServResponseDeserializeSuccess Failed: Wrong Error Message")
}

func TestOciTaskServResponseDeserializeFailed(test *testing.T) {
	data := "Test Error Message"

	dataJson, _ := json.Marshal(data)
	ociTaskServResponse := OciTaskServResponse{}

	err := ociTaskServResponse.Deserialize(dataJson)

	assert.Error(test, err, "TestOciTaskServResponseDeserializeFailed Failed")
}
