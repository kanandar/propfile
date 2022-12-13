package list

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Sort struct {
}

func init() {
	function.Register(&Sort{})
}

func (s *Sort) Name() string {
	return "sort"
}

func (s *Sort) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeBool}, true
}

func (s *Sort) Eval(params ...interface{}) (interface{}, error) {
	items := s.InterfaceToArray(params[0])
	desc := params[1].(bool)
	field:=""

	if len(params) > 2{
		field = params[2].(string)
	}

	if items == nil || len(items) == 0 {
		return []interface{}{}, nil
	}

	//If the field is empty then is a flat array
	if len(strings.TrimSpace(field)) == 0 {
		switch reflect.ValueOf(items[0]).Kind() {
		case reflect.String:
			return s.SortStrings(items, desc), nil
		default:
			return s.SortNumbers(items, desc), nil
		}
	}

	//Sorting by field
	item := items[0].(map[string]interface{})
	fieldValue :=  item[field]
	reflect.ValueOf(fieldValue)
	switch reflect.ValueOf(fieldValue).Kind() {
	case reflect.String:
		s.SortByStringField(items, field, desc)
	default:
		s.SortByNumberField(items, field, desc)
	}

	return items, nil
}

func (s *Sort) InterfaceToArray(something interface{}) []interface{} {
	var items []interface{}
	value := reflect.ValueOf(something)
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i).Interface()
		items = append(items, item)
	}

	return items
}

func (s *Sort) ToFloat(something interface{}) float64 {
	number, _ := strconv.ParseFloat(s.ToString(something), 64)

	return number
}

func (s *Sort) ToString(something interface{}) string {
	return fmt.Sprintf("%v", something)
}

func (s *Sort) SortStrings(list []interface{}, desc bool) []string{
	var stringList []string
	for i := 0; i < len(list); i++ {
		stringList = append(stringList, list[i].(string))
	}
	sort.Strings(stringList)
	if desc {
		stringList = s.ReverseStrings(stringList)
	}

	return stringList
}

func (s *Sort) SortNumbers(list []interface{}, desc bool) []float64 {
	var numberList []float64
	for i := 0; i < len(list); i++ {
		numberList = append(numberList, s.ToFloat(list[i]))
	}
	sort.Float64s(numberList)
	if desc {
		numberList = s.ReverseNumbers(numberList)
	}

	return numberList
}

func (s *Sort) ReverseNumbers(numbers []float64) []float64 {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}

func (s *Sort) ReverseStrings(strings []string) []string {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}

	return strings
}

func (s *Sort) SortByNumberField(list []interface{}, field string, desc bool){
	sort.Slice(list, func(i, j int) bool {
		result := s.ToFloat(list[i].(map[string]interface{})[field]) > s.ToFloat(list[j].(map[string]interface{})[field])
		if desc {
			return result
		}else{
			return !result
		}
	})
}

func (s *Sort) SortByStringField(list []interface{}, field string, desc bool){
	sort.Slice(list, func(i, j int) bool {
		result := s.ToString(list[i].(map[string]interface{})[field]) > s.ToString(list[j].(map[string]interface{})[field])
		if desc {
			return result
		} else {
			return !result
		}
	})
}
