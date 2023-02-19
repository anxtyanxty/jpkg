package jpkg

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strconv"
)

var (
	query_expr, _ = regexp.Compile(`\[(.*?)\]`)
)

func LoadJObject(json_bytes []byte) (*JObject, error) {

	object := &JObject{
		data: nil,
	}
	if err := json.Unmarshal(json_bytes, &object.data); err != nil {
		return nil, err
	}
	return object, nil
}

func QueryJSON(data map[interface{}]interface{}, jpath_query string) (interface{}, error) {
	var err error
	var interface_value interface{}
	var interface_kind reflect.Kind
	query := parse_query(jpath_query)
	for index, item := range query {
		if item.IsIndex {
			number, _ := strconv.Atoi(item.Data)
			interface_arr := interface_value.([]interface{})
			if len(interface_arr)-1 < number {
				return nil, fmt.Errorf("could not progress any further, trying to index out of bounds, array last index is %d and your index is %d", len(interface_arr)-1, number)
			}
			interface_value = interface_value.([]interface{})[number]

		} else {
			interface_value, err = match_value(data, item.Data)
			if err != nil {
				return nil, err
			}

			interface_type := reflect.TypeOf(interface_value)
			if interface_type == nil && index == len(query)-1 {
				break
			} else if interface_type == nil {
				return nil, fmt.Errorf("could not progress any further, type is null, last item evaluated: %s", item.Data)
			}

			interface_kind = interface_type.Kind()
			if !is_indexable(interface_kind) && index != len(query)-1 {
				return nil, fmt.Errorf("could not progress any further, last item evaluated: %s", item.Data)
			}
		}

		if interface_kind == reflect.Map {
			data = CreateInterfaceMap(interface_value.(map[string]interface{}))
		}
	}

	return interface_value, nil
}

func QueryJSONString(data string, jpath_query string) (interface{}, error) {
	var jmap interface{}
	err := json.Unmarshal([]byte(data), &jmap)
	if err != nil {
		return nil, err
	}
	var m map[interface{}]interface{}

	switch jmap := jmap.(type) {
	case map[string]interface{}:
		m = CreateInterfaceMap(jmap)
	case []interface{}:
		m = CreateInterfaceMapSlice(jmap)
	}
	return QueryJSON(m, jpath_query)
}

func QueryJSONInterface(data interface{}, jpath_query string) (interface{}, error) {
	var m map[interface{}]interface{}

	switch jmap := data.(type) {
	case map[string]interface{}:
		m = CreateInterfaceMap(jmap)
	case []interface{}:
		m = CreateInterfaceMapSlice(jmap)
	default:
		return nil, fmt.Errorf("the type %#v is not supported", reflect.TypeOf(jmap))
	}
	return QueryJSON(m, jpath_query)
}

func QueryJSONReader(data io.Reader, jpath_query string) (interface{}, error) {
	body, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}
	return QueryJSONString(string(body), jpath_query)
}

func ParseToRawString(data interface{}) (string, error) {
	buffer, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

func CreateInterfaceMap(data map[string]interface{}) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	for str, value := range data {
		var key interface{} = str
		m[key] = value
	}

	return m
}

func CreateInterfaceMapSlice(data []interface{}) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	for index, value := range data {
		var key interface{} = index
		m[key] = value
	}

	return m
}
