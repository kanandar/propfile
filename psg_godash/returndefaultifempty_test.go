package psggodash

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var returnDefaultIfEmptyFnRef = &returnDefaultIfEmptyFn{}
var returnDefaultIfEmptyFnTestLogger log.Logger
var actualOutput interface{}
var expectedOutput interface{}
var input interface{}
var err error

func init() {
	returnDefaultIfEmptyFnTestLogger = log.RootLogger()
	log.SetLogLevel(returnDefaultIfEmptyFnTestLogger, log.DebugLevel)
}

func Test_returnDefaultIfEmpty_1(t *testing.T) {
	input = 1.3
	expectedOutput = 1.3

	actualOutput, err = returnDefaultIfEmptyFnRef.Eval(input, 0.0)
	actualOutput = actualOutput.(float64)

	returnDefaultIfEmptyFnTestLogger.Debug("In tester: Output of function call = ", expectedOutput)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedOutput, actualOutput)
}

func Test_returnDefaultIfEmpty_2(t *testing.T) {
	input = ""
	expectedOutput = 0.0

	actualOutput, err = returnDefaultIfEmptyFnRef.Eval(input, 0.0)
	actualOutput = actualOutput.(float64)

	returnDefaultIfEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", actualOutput)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedOutput, actualOutput)
}

func Test_returnDefaultIfEmpty_3(t *testing.T) {
	input = nil
	expectedOutput = 3

	actualOutput, err = returnDefaultIfEmptyFnRef.Eval(input, 3)
	actualOutput = actualOutput.(int)

	returnDefaultIfEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", actualOutput)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedOutput, actualOutput)
}

func Test_returnDefaultIfEmpty_4(t *testing.T) {

	type Config struct {
		host string
		port float64
	}
	config := Config{
		host: "myhost.com",
		port: 22,
	} // not nil

	var input Config
	var expectedOutput Config

	actualOutput, err = returnDefaultIfEmptyFnRef.Eval(input, config)

	returnDefaultIfEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", actualOutput)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedOutput, actualOutput)
}
func Test_returnDefaultIfEmpty_5(t *testing.T) {

	var input []interface{}
	var expectedOutput []interface{}
	var defaultValue []interface{}
	data := `[
				{"test1":[{"test":"123"}]},
				{"test1":[{"test":"456"}]}
	]`

	ok := json.Unmarshal([]byte(data), &expectedOutput)
	fmt.Println(ok)
	ok = json.Unmarshal([]byte(data), &defaultValue)
	fmt.Println(ok)

	actualOutput, err = returnDefaultIfEmptyFnRef.Eval(input, defaultValue)

	returnDefaultIfEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", actualOutput)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedOutput, actualOutput)
}
func Test_returnDefaultIfEmpty_6(t *testing.T) {

	var input map[string]interface{}
	var expectedOutput map[string]interface{}
	var defaultValue map[string]interface{}
	data := `{
  "name":"John",
  "age":30,
  "cars": [
    { "name":"Ford", "models":[ "Fiesta", "Focus", "Mustang" ] },
    { "name":"BMW", "models":[ "320", "X3", "X5" ] },
    { "name":"Fiat", "models":[ "500", "Panda" ] }
  ]
 }`

	ok := json.Unmarshal([]byte(data), &expectedOutput)
	fmt.Println(ok)
	ok = json.Unmarshal([]byte(data), &defaultValue)
	fmt.Println(ok)

	actualOutput, err = returnDefaultIfEmptyFnRef.Eval(input, defaultValue)

	returnDefaultIfEmptyFnTestLogger.Debugf("In tester: Output of function call = %+v", actualOutput)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedOutput, actualOutput)
}
