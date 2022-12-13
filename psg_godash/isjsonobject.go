package psggodash

import (
	"encoding/json"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&isJSONObjectFn{})
}

type isJSONObjectFn struct {
}

// Name returns the name of the function
func (isJSONObjectFn) Name() string {
	return "isJSONObject"
}

// Sig returns the function signature
func (isJSONObjectFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny}, false
}

var isJSONObjectFnLogger = log.RootLogger()

/* Eval executes the function
* This function returns true, if input is a json object.
* It can validate a json string and json object.
 */
func (isJSONObjectFn) Eval(params ...interface{}) (interface{}, error) {
	if isJSONObjectFnLogger.DebugEnabled() {
		isJSONObjectFnLogger.Debugf("Entering function isJSONObject (eval) with 1 parameter")
		isJSONObjectFnLogger.Debugf("Parameter 1 json object = << %+v >>", params[0])
		isJSONObjectFnLogger.Debugf("Parameter 2 treat null as json object = << %+v >>", params[1])
	}

	inputParamJson := params[0]
	inputParamIsNullAValidJsonObject := params[1]

	var isNullAValidJsonObject bool
	var isInputAJsonObject bool = false
	var err error

	if isJSONObjectFnLogger.DebugEnabled() {
		isJSONObjectFnLogger.Debugf("Print value after variable assignment of parameters. No type conversion.")
		isJSONObjectFnLogger.Debugf("Input to validate is = << %+v >> with type = << %T >> ", inputParamJson, inputParamJson)
		isJSONObjectFnLogger.Debugf("Is Null A Valid Json Object = << %+v >> with type = << %T >>", inputParamIsNullAValidJsonObject, inputParamIsNullAValidJsonObject)
		isJSONObjectFnLogger.Debugf("Start Validation")
	}

	isNullAValidJsonObject, err = coerce.ToBool(inputParamIsNullAValidJsonObject)

	if err != nil {
		isNullAValidJsonObject = false
		if isJSONObjectFnLogger.DebugEnabled() {
			isJSONObjectFnLogger.Debugf("Error coercing second param to boolean. Defaulting to false.")
		}
	}

	switch t := inputParamJson.(type) {
	case nil:
		//If the input object is nil, return true or false, depending on the user input.
		if isJSONObjectFnLogger.DebugEnabled() {
			isJSONObjectFnLogger.Debugf("Value is null.")
			isJSONObjectFnLogger.Debugf("Return %+v.", isNullAValidJsonObject)
		}
		isInputAJsonObject = isNullAValidJsonObject
	case string:
		if isJSONObjectFnLogger.DebugEnabled() {
			isJSONObjectFnLogger.Debugf("Input is string.")
		}
		//check if first parameter is a json object
		var inputJsonObject map[string]interface{}
		err := json.Unmarshal([]byte(inputParamJson.(string)), &inputJsonObject)
		if err != nil {
			isJSONObjectFnLogger.Debugf("Unable to coerce input string to json object. Error is %+v", err)
			isJSONObjectFnLogger.Debugf("Return false.")
		} else {
			isInputAJsonObject = true
		}
	case map[string]interface{}:
		if isJSONObjectFnLogger.DebugEnabled() {
			isJSONObjectFnLogger.Debugf("Input is a json object.")
			isJSONObjectFnLogger.Debugf("Return true.")
			isInputAJsonObject = true
		}
	default:
		if isJSONObjectFnLogger.DebugEnabled() {
			isJSONObjectFnLogger.Debugf("Default value in switch statement. Input is not a json object. Type is << %T >>", t)
			isJSONObjectFnLogger.Debugf("Return false.")
		}
	}

	if isJSONObjectFnLogger.DebugEnabled() {
		isJSONObjectFnLogger.Debugf("Final output value = [%+v]", isInputAJsonObject)
	}

	if isJSONObjectFnLogger.DebugEnabled() {
		isJSONObjectFnLogger.Debugf("Exiting function isJSONObject (eval)")
	}

	return isInputAJsonObject, nil
}
