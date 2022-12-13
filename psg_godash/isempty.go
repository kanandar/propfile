package psggodash

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&isEmptyFn{})
}

type isEmptyFn struct {
}

// Name returns the name of the function
func (isEmptyFn) Name() string {
	return "isEmpty"
}

// Sig returns the function signature
func (isEmptyFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

var isEmptyFnLogger = log.RootLogger()

// Eval executes the function
// This function returns true, if the input value is either empty or nil  (json null)
func (isEmptyFn) Eval(params ...interface{}) (interface{}, error) {
	if isEmptyFnLogger.DebugEnabled() {
		isEmptyFnLogger.Debugf("Entering function isEmpty (eval) with param: %+v", params[0])
	}

	inputParamValue := params[0]
	inputString := ""
	var outputValue bool
	var err error

	switch inputParamValue.(type) {
	case nil:
		if isEmptyFnLogger.DebugEnabled() {
			isEmptyFnLogger.Debugf("Processing value of type nil.")
			isEmptyFnLogger.Debugf("Input Parameter is nil. Returning true.")
		}
		outputValue = true
	case []interface{}:
		if isEmptyFnLogger.DebugEnabled() {
			isEmptyFnLogger.Debugf("Processing an array of interfaces.")
		}
		if inputParamValue.([]interface{}) == nil {
			outputValue = true
			isEmptyFnLogger.Debugf("Input Parameter is nil. Returning true.")
		}
	case map[string]interface{}:
		if isEmptyFnLogger.DebugEnabled() {
			isEmptyFnLogger.Debug("Processing a map of interface (json object).")
		}
		if inputParamValue.(map[string]interface{}) == nil {
			outputValue = true
			isEmptyFnLogger.Debugf("Input Parameter is nil. Returning true.")
		}
	default:
		if isEmptyFnLogger.DebugEnabled() {
			isEmptyFnLogger.Debugf("Applying default processing logic for input param value of type %+T.", inputParamValue)
		}
		inputString, err = coerce.ToString(inputParamValue)
		if err != nil {
			return nil, fmt.Errorf("Unable to coerece input value to a string. value = %+v.", inputParamValue)
		}

		if isEmptyFnLogger.DebugEnabled() {
			isEmptyFnLogger.Debugf("Input Parameter's string length is %+v.", len(inputString))
		}

		if len(inputString) <= 0 {
			outputValue = true
			isEmptyFnLogger.Debugf("Input Parameter is empty or nil. Returning default value.")
		}

	} //ENDS - switch statement.

	if isEmptyFnLogger.DebugEnabled() {
		isEmptyFnLogger.Debugf("Final output value = %+v", outputValue)
	}

	if isEmptyFnLogger.DebugEnabled() {
		isEmptyFnLogger.Debugf("Exiting function isEmpty (eval)")
	}

	return outputValue, nil
}
