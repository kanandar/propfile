package psgjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsonData interface{}
var pathfnRef = &fnPath{}

func init() {
	data := `{
        "lead": {
            "email": "ab@test.com"
        },
        "eventType": "test"
    }`

	json.Unmarshal([]byte(data), &jsonData)
}

func Test_path_JsonPathLookup_1(t *testing.T) {

	// Retrieve email value.
	res, err := pathfnRef.Eval("$.lead.email", jsonData)
	assert.Nil(t, err)
	assert.Equal(t, "ab@test.com", res)

	// Retrieve bicycle height, which does not exist.
	res, err = pathfnRef.Eval("$.store.bicycle.height", jsonData)
	assert.Nil(t, err)
	assert.Empty(t, res)

	// retrieve event type from root.
	res, _ = pathfnRef.Eval("$.eventType", jsonData)
	if resv, ok := res.(string); ok != true || resv != "test" {
		t.Errorf("Event Type should be test")
	}
}

func Test_path_JsonPathLookup_2(t *testing.T) {

	data := `null`

	json.Unmarshal([]byte(data), &jsonData)

	// Retrieve email value.
	res, err := pathfnRef.Eval("$.lead.email", jsonData)
	assert.Nil(t, err)
	assert.Equal(t, "", res)
}
