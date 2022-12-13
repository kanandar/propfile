package list

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sortF = &Sort{}

func GenerateList() []map[string]interface{} {
	f1 := map[string]interface{}{
		"name": "B",
		"age": 3,
		"height": 1.73,
	}
	f2 := map[string]interface{}{
		"name": "D",
		"age":1,
		"height": 1.63,
	}
	f3 := map[string]interface{}{
		"name": "A",
		"age":4,
		"height": 1.88,
	}
	f4 := map[string]interface{}{
		"name": "C",
		"age":2,
		"height": 1.79,
	}

	var persons []map[string]interface{}
	persons = append(persons, f1)
	persons = append(persons, f2)
	persons = append(persons, f3)
	persons = append(persons, f4)

	return persons
}

func TestSortNumber(t *testing.T) {
	numbers := [4]int{3, 4, 1, 2}
	sorted, _ := sortF.Eval(numbers, true)
	sortedList := sortF.InterfaceToArray(sorted)
	assert.Equal(t, sortF.ToFloat(sortedList[0]), float64(4))
	assert.Equal(t, sortF.ToFloat(sortedList[1]), float64(3))
	assert.Equal(t, sortF.ToFloat(sortedList[2]), float64(2))
	assert.Equal(t, sortF.ToFloat(sortedList[3]), float64(1))
}

func TestSortString(t *testing.T) {
	numbers := [4]string{"C", "D", "A", "B"}
	sorted, _ := sortF.Eval(numbers, false)
	sortedList := sortF.InterfaceToArray(sorted)
	assert.Equal(t, reflect.ValueOf(sortedList[0]).String(), "A")
	assert.Equal(t, reflect.ValueOf(sortedList[1]).String(), "B")
	assert.Equal(t, reflect.ValueOf(sortedList[2]).String(), "C")
	assert.Equal(t, reflect.ValueOf(sortedList[3]).String(), "D")
}

func TestSortNumberByField(t *testing.T) {
	field := "age"
	items := GenerateList()
	sorted, _ := sortF.Eval(items, false, field)
	sortedList := sorted.([]interface{})
	assert.Equal(t, sortF.ToFloat(sortedList[0].(map[string]interface{})[field]), float64(1))
	assert.Equal(t, sortF.ToFloat(sortedList[1].(map[string]interface{})[field]), float64(2))
	assert.Equal(t, sortF.ToFloat(sortedList[2].(map[string]interface{})[field]), float64(3))
	assert.Equal(t, sortF.ToFloat(sortedList[3].(map[string]interface{})[field]), float64(4))
}

func TestSortStringByField(t *testing.T) {
	field := "name"
	items := GenerateList()
	sorted, _ := sortF.Eval(items, true, field)
	sortedList := sorted.([]interface{})
	assert.Equal(t, sortF.ToString(sortedList[0].(map[string]interface{})[field]), "D")
	assert.Equal(t, sortF.ToString(sortedList[1].(map[string]interface{})[field]), "C")
	assert.Equal(t, sortF.ToString(sortedList[2].(map[string]interface{})[field]), "B")
	assert.Equal(t, sortF.ToString(sortedList[3].(map[string]interface{})[field]), "A")
}

func TestEmptyArray(t *testing.T) {
	var items []string
	sorted, _ := sortF.Eval(items, true)
	sortedList := sorted.([]interface{})
	assert.Equal(t, len(sortedList), 0)
}
