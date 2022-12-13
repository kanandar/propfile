package psgjson

import (
	"regexp"
	"strings"

	"github.com/oliveagle/jsonpath"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&fnPath{})
}

type fnPath struct {
}

// Name returns the name of the function
func (fnPath) Name() string {
	return "path"
}

// Sig returns the function signature
func (fnPath) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeAny}, false
}

var pathFnLogger = log.RootLogger()

// Eval executes the function
func (fnPath) Eval(params ...interface{}) (interface{}, error) {
	if pathFnLogger.DebugEnabled() {
		pathFnLogger.Debugf("Entering function ylopo_json.path (eval) with params: %+v and %+v", params[0], params[1])
	}

	inputExpressionString := params[0].(string)
	inputObject := params[1]
	//tmp fix to take $loop as $. for now
	if strings.HasPrefix(strings.TrimSpace(inputExpressionString), "$loop.") {
		inputExpressionString = strings.Replace(inputExpressionString, "$loop", "$", -1)
	}

	if pathFnLogger.DebugEnabled() {
		pathFnLogger.Debugf("Json path string is: ", inputExpressionString)
		pathFnLogger.Debugf("Json object is of type %T", inputObject)
	}

	if inputObject == nil {
		if pathFnLogger.DebugEnabled() {
			pathFnLogger.Debugf("JSON Data value is nil. Return empty string")
		}
		return "", nil
	}

	out, mainerr := jsonpath.JsonPathLookup(inputObject, inputExpressionString)

	if mainerr != nil {
		if pathFnLogger.DebugEnabled() {
			pathFnLogger.Debugf("Error when calling json path lookup. Error is: %+v", mainerr)
		}
		matched, _ := regexp.MatchString("key error:.*not found in object", mainerr.Error())
		if matched {
			if pathFnLogger.DebugEnabled() {
				pathFnLogger.Debugf("Requested key is not found in the object. Return empty string")
			}
			return "", nil
		} else {
			if pathFnLogger.DebugEnabled() {
				pathFnLogger.Debugf("Some other error occurred. Return error.")
			}
			return nil, mainerr
		}
	} else {
		if out == nil {
			if pathFnLogger.DebugEnabled() {
				pathFnLogger.Debugf("Value of the requested key is nil. Returning empty string.")
			}
			return "", mainerr
		}

		if pathFnLogger.DebugEnabled() {
			pathFnLogger.Debugf("Output value is: ##%+v##", out)
		}

		if pathFnLogger.DebugEnabled() {
			pathFnLogger.Debugf("Exiting function ylopo_json.path (eval)")
		}
		return out, mainerr
	}
}
