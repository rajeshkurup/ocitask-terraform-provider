package ocitaskclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOciErrorSerializeSuccess(test *testing.T) {
	errorCode := 1001
	errorMessage := "Test Error Message"

	ociError := OciError{}
	ociError.ErrorCode = &errorCode
	ociError.ErrorMessage = &errorMessage

	dataJson, err := ociError.Serialize()

	assert.NoError(test, err, "TestOciErrorSerializeSuccess Failed: Unable to serialize OciError")

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(dataJson), &data)

	assert.NoError(test, err, "TestOciErrorSerializeSuccess Failed: Unable to deserialize OciError")
	assert.Equal(test, 1001, int(data["errorCode"].(float64)), "TestOciErrorSerializeSuccess Failed: Wrong Error Code")
	assert.Equal(test, "Test Error Message", data["errorMessage"].(string), "TestOciErrorSerializeSuccess Failed: Wrong Error Message")
}

func TestOciErrorDeserializeSuccess(test *testing.T) {
	data := make(map[string]interface{})
	data["errorCode"] = 1001
	data["errorMessage"] = "Test Error Message"

	dataJson, _ := json.Marshal(data)
	ociError := OciError{}

	err := ociError.Deserialize(dataJson)

	assert.NoError(test, err, "TestOciErrorDeserializeSuccess Failed: Unable to Deserialize OciError")
	assert.Equal(test, 1001, *ociError.ErrorCode, "TestOciErrorSerializeSuccess Failed: Wrong Error Code")
	assert.Equal(test, "Test Error Message", *ociError.ErrorMessage, "TestOciErrorSerializeSuccess Failed: Wrong Error Message")
}

func TestOciErrorDeserializeFailed(test *testing.T) {
	data := "Test Error Message"

	dataJson, _ := json.Marshal(data)
	ociError := OciError{}

	err := ociError.Deserialize(dataJson)

	assert.Error(test, err, "TestOciErrorDeserializeFailed Failed")
}
