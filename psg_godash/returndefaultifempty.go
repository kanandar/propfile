package psggodash

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&returnDefaultIfEmptyFn{})
}

type returnDefaultIfEmptyFn struct {
}

// Name returns the name of the function
func (returnDefaultIfEmptyFn) Name() string {
	return "returnDefaultIfEmpty"
}

// Sig returns the function signature
func (returnDefaultIfEmptyFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny}, false
}

var returnDefaultIfEmptyFnLogger = log.RootLogger()

// Eval executes the function
func (returnDefaultIfEmptyFn) Eval(params ...interface{}) (interface{}, error) {
	if returnDefaultIfEmptyFnLogger.DebugEnabled() {
		returnDefaultIfEmptyFnLogger.Debugf("Entering function returnDefaultIfEmpty (eval) with param: %+v", params[0])
	}

	inputParamValue := params[0]
	inputParamDefaultValue := params[1]
	returnDefault := false
	inputString := ""
	var outputValue interface{}
	var err error

	switch inputParamValue.(type) {
	case nil:
		if returnDefaultIfEmptyFnLogger.DebugEnabled() {
			returnDefaultIfEmptyFnLogger.Debugf("Processing value of type nil.")
			returnDefaultIfEmptyFnLogger.Debugf("Input Parameter is nil. Returning default value.")
		}
		returnDefault = true
	case []interface{}:
		if returnDefaultIfEmptyFnLogger.DebugEnabled() {
			returnDefaultIfEmptyFnLogger.Debugf("Processing an array of interfaces.")
		}
		if inputParamValue.([]interface{}) == nil {
			returnDefault = true
			returnDefaultIfEmptyFnLogger.Debugf("Input Parameter is nil. Returning default value.")
		}
	case map[string]interface{}:
		if returnDefaultIfEmptyFnLogger.DebugEnabled() {
			returnDefaultIfEmptyFnLogger.Debug("Processing a map of interface (json object).")
		}
		if inputParamValue.(map[string]interface{}) == nil {
			returnDefault = true
			returnDefaultIfEmptyFnLogger.Debugf("Input Parameter is nil. Returning default value.")
		}
	default:
		if returnDefaultIfEmptyFnLogger.DebugEnabled() {
			returnDefaultIfEmptyFnLogger.Debugf("Applying default processing logic for input param value of type %+T.", inputParamValue)
		}
		inputString, err = coerce.ToString(inputParamValue)
		if err != nil {
			return nil, fmt.Errorf("Unable to coerece input value to a string. value = %+v.", inputParamValue)
		}

		if returnDefaultIfEmptyFnLogger.DebugEnabled() {
			returnDefaultIfEmptyFnLogger.Debugf("Input Parameter's string length is %+v.", len(inputString))
		}

		if len(inputString) <= 0 {
			returnDefault = true
			returnDefaultIfEmptyFnLogger.Debugf("Input Parameter is empty or nil. Returning default value.")
		}

	} //ENDS - switch statement.

	if returnDefault {
		outputValue = inputParamDefaultValue
	} else {
		outputValue = inputParamValue
	}

	if returnDefaultIfEmptyFnLogger.DebugEnabled() {
		returnDefaultIfEmptyFnLogger.Debugf("Final output value = %+v", outputValue)
	}

	if returnDefaultIfEmptyFnLogger.DebugEnabled() {
		returnDefaultIfEmptyFnLogger.Debugf("Exiting function returnDefaultIfEmpty (eval)")
	}

	return outputValue, nil
}
