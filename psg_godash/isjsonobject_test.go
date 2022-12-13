package psggodash

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var isJSONObjectFnRef = &isJSONObjectFn{}
var isJSONObjectFnTestLogger log.Logger
var isJSONObjectFnActualOutput interface{}
var isJSONObjectFnExpectedOutput interface{}
var isJSONObjectFnInputParam1 interface{}
var isJSONObjectFnInputParam2 interface{}
var isJSONObjectFnErr error

func init() {
	isJSONObjectFnTestLogger = log.RootLogger()
	log.SetLogLevel(isJSONObjectFnTestLogger, log.DebugLevel)
}

//Expected input values
//isJSONObjectFnInputParam1 -> json string or object
//isJSONObjectFnInputParam2 -> treat null value as valid json object.

//sunny path case
func Test_isJSONObject_1(t *testing.T) {
	data := `{"test":{"test1":"123"},"test2":"123"}`
	err := json.Unmarshal([]byte(data), &isJSONObjectFnInputParam1)
	isJSONObjectFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)
	isJSONObjectFnInputParam2 = true

	isJSONObjectFnExpectedOutput = true

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

func Test_isJSONObject_2(t *testing.T) {
	data := `{"test":{"test1":"123"},"test2":"123"}`
	err := json.Unmarshal([]byte(data), &isJSONObjectFnInputParam1)
	isJSONObjectFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)
	isJSONObjectFnInputParam2 = false

	isJSONObjectFnExpectedOutput = true

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

func Test_isJSONObject_3(t *testing.T) {
	data := "{\"test\":{\"test1\":\"123\"},\"test2\":\"123\"}"

	isJSONObjectFnInputParam1 = data
	isJSONObjectFnInputParam2 = false

	isJSONObjectFnExpectedOutput = true

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

//empty json object (as string)
func Test_isJSONObject_4(t *testing.T) {
	data := "{}"

	isJSONObjectFnInputParam1 = data
	isJSONObjectFnInputParam2 = false

	isJSONObjectFnExpectedOutput = true

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

//empty json object
func Test_isJSONObject_5(t *testing.T) {
	data := `{}`
	err := json.Unmarshal([]byte(data), &isJSONObjectFnInputParam1)
	isJSONObjectFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)
	isJSONObjectFnInputParam2 = false

	isJSONObjectFnExpectedOutput = true

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

//null value, with treat null value as valid json object = true
func Test_isJSONObject_6(t *testing.T) {
	data := `null`
	err := json.Unmarshal([]byte(data), &isJSONObjectFnInputParam1)
	isJSONObjectFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)
	isJSONObjectFnInputParam2 = true

	isJSONObjectFnExpectedOutput = true

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

//null value, with treat null value as valid json object = false
func Test_isJSONObject_7(t *testing.T) {
	data := `null`
	err := json.Unmarshal([]byte(data), &isJSONObjectFnInputParam1)
	isJSONObjectFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)
	isJSONObjectFnInputParam2 = false

	isJSONObjectFnExpectedOutput = false

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

//null value, with treat null value as valid json object = false
func Test_isJSONObject_8(t *testing.T) {

	isJSONObjectFnInputParam1 = nil
	isJSONObjectFnInputParam2 = false

	isJSONObjectFnExpectedOutput = false

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

//null value, with treat null value as valid json object = false
func Test_isJSONObject_9(t *testing.T) {

	isJSONObjectFnInputParam1 = nil
	isJSONObjectFnInputParam2 = true

	isJSONObjectFnExpectedOutput = true

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

//negative test cases
//empty string
func Test_isJSONObject_10(t *testing.T) {
	data := ""

	isJSONObjectFnInputParam1 = data
	isJSONObjectFnInputParam2 = true

	isJSONObjectFnExpectedOutput = false

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

//array
func Test_isJSONObject_11(t *testing.T) {
	data := `[0,1,2,3,4,5,6,7]`
	err := json.Unmarshal([]byte(data), &isJSONObjectFnInputParam1)
	isJSONObjectFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)
	isJSONObjectFnInputParam2 = true

	isJSONObjectFnExpectedOutput = false

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}

//empty array
func Test_isJSONObject_12(t *testing.T) {
	data := `[]`
	err := json.Unmarshal([]byte(data), &isJSONObjectFnInputParam1)
	isJSONObjectFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)
	isJSONObjectFnInputParam2 = true

	isJSONObjectFnExpectedOutput = false

	isJSONObjectFnActualOutput, isJSONObjectFnErr = isJSONObjectFnRef.Eval(isJSONObjectFnInputParam1, isJSONObjectFnInputParam2)

	isJSONObjectFnTestLogger.Debug("In tester: Output of function call = ", isJSONObjectFnActualOutput)
	assert.Nil(t, isJSONObjectFnErr)
	assert.EqualValues(t, isJSONObjectFnExpectedOutput, isJSONObjectFnActualOutput)
}
