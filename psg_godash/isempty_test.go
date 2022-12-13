package psggodash

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var isEmptyFnRef = &isEmptyFn{}
var isEmptyFnTestLogger log.Logger
var isEmptyFnActualOutput interface{}
var isEmptyFnExpectedOutput interface{}
var isEmptyFnInput interface{}
var isEmptyFnErr error

func init() {
	isEmptyFnTestLogger = log.RootLogger()
	log.SetLogLevel(isEmptyFnTestLogger, log.DebugLevel)
}

func Test_isEmpty_1(t *testing.T) {
	isEmptyFnInput = 1.3
	isEmptyFnExpectedOutput = false

	isEmptyFnActualOutput, isEmptyFnErr = isEmptyFnRef.Eval(isEmptyFnInput)
	isEmptyFnActualOutput = isEmptyFnActualOutput.(bool)

	isEmptyFnTestLogger.Debug("In tester: Output of function call = ", isEmptyFnExpectedOutput)
	assert.Nil(t, isEmptyFnErr)
	assert.EqualValues(t, isEmptyFnExpectedOutput, isEmptyFnActualOutput)
}

func Test_isEmpty_2(t *testing.T) {
	isEmptyFnInput = ""
	isEmptyFnExpectedOutput = true

	isEmptyFnActualOutput, isEmptyFnErr = isEmptyFnRef.Eval(isEmptyFnInput)
	isEmptyFnActualOutput = isEmptyFnActualOutput.(bool)

	isEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", isEmptyFnActualOutput)
	assert.Nil(t, isEmptyFnErr)
	assert.EqualValues(t, isEmptyFnExpectedOutput, isEmptyFnActualOutput)
}

func Test_isEmpty_3(t *testing.T) {
	isEmptyFnInput = nil
	isEmptyFnExpectedOutput = true

	isEmptyFnActualOutput, isEmptyFnErr = isEmptyFnRef.Eval(isEmptyFnInput, 3)
	isEmptyFnActualOutput = isEmptyFnActualOutput.(bool)

	isEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", isEmptyFnActualOutput)
	assert.Nil(t, isEmptyFnErr)
	assert.EqualValues(t, isEmptyFnExpectedOutput, isEmptyFnActualOutput)
}

func Test_isEmpty_4(t *testing.T) {

	type Config struct {
		host string
		port float64
	}

	var input4 Config
	isEmptyFnExpectedOutput = false

	isEmptyFnActualOutput, isEmptyFnErr = isEmptyFnRef.Eval(input4)
	isEmptyFnActualOutput = isEmptyFnActualOutput.(bool)

	isEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", isEmptyFnActualOutput)
	assert.Nil(t, isEmptyFnErr)
	assert.EqualValues(t, isEmptyFnExpectedOutput, isEmptyFnActualOutput)
}
func Test_isEmpty_5(t *testing.T) {

	var input []interface{}
	isEmptyFnExpectedOutput = false

	data := `[
				{"test1":[{"test":"123"}]},
				{"test1":[{"test":"456"}]}
	]`

	ok := json.Unmarshal([]byte(data), &input)
	fmt.Println(ok)

	isEmptyFnActualOutput, isEmptyFnErr = isEmptyFnRef.Eval(input)
	isEmptyFnActualOutput = isEmptyFnActualOutput.(bool)

	isEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", isEmptyFnActualOutput)
	assert.Nil(t, isEmptyFnErr)
	assert.EqualValues(t, isEmptyFnExpectedOutput, isEmptyFnActualOutput)
}
func Test_isEmpty_6(t *testing.T) {

	var input map[string]interface{}
	isEmptyFnExpectedOutput = false
	data := `{
  "name":"John",
  "age":30,
  "cars": [
    { "name":"Ford", "models":[ "Fiesta", "Focus", "Mustang" ] },
    { "name":"BMW", "models":[ "320", "X3", "X5" ] },
    { "name":"Fiat", "models":[ "500", "Panda" ] }
  ]
 }`

	ok := json.Unmarshal([]byte(data), &input)
	fmt.Println(ok)

	isEmptyFnActualOutput, isEmptyFnErr = isEmptyFnRef.Eval(input)
	isEmptyFnActualOutput = isEmptyFnActualOutput.(bool)

	isEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", isEmptyFnActualOutput)
	assert.Nil(t, isEmptyFnErr)
	assert.EqualValues(t, isEmptyFnExpectedOutput, isEmptyFnActualOutput)
}
