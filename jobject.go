package jpkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func (obj *JObject) QueryString(path string) (string, error) {
	undefined_data, err := QueryJSONInterface(obj.data, path)
	if err != nil {
		return "", err
	}
	undefined_type := reflect.TypeOf(undefined_data)
	if undefined_type.Kind() != reflect.String {
		return "", fmt.Errorf("the type is invalid, you are asking for string but the type is %s", undefined_type)
	}
	return undefined_data.(string), nil
}

func (obj *JObject) QueryInt(path string) (int, error) {

	decimal, err := obj.QueryDecimal(path)
	if err != nil {
		return 0, err
	}
	return int(decimal), nil
}

func (obj *JObject) QueryDecimal(path string) (float64, error) {
	undefined_data, err := QueryJSONInterface(obj.data, path)
	if err != nil {
		return 0, err
	}
	undefined_type := reflect.TypeOf(undefined_data)
	if undefined_type.Kind() != reflect.Float64 {
		return 0, fmt.Errorf("the type is invalid, you are asking for float64 but the type is %s", undefined_type)
	}
	return undefined_data.(float64), nil
}

func (obj *JObject) QueryBoolean(path string) (bool, error) {
	undefined_data, err := QueryJSONInterface(obj.data, path)
	if err != nil {
		return false, err
	}
	undefined_type := reflect.TypeOf(undefined_data)
	if undefined_type.Kind() != reflect.Bool {
		return false, fmt.Errorf("the type is invalid, you are asking for boolean but the type is %s", undefined_type)
	}
	return undefined_data.(bool), nil
}

func (obj *JObject) QueryArray(path string) ([]*JObject, error) {

	undefined_data, err := QueryJSONInterface(obj.data, path)
	if err != nil {
		return nil, err
	}
	undefined_type := reflect.TypeOf(undefined_data)
	if undefined_type.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the type is invalid, you are asking for a []*JObject but the type is %s", undefined_type)
	}

	var objects []*JObject
	for _, undefined_object := range undefined_data.([]interface{}) {
		objects = append(objects, &JObject{
			data: undefined_object,
		})
	}

	return objects, nil
}

func (obj *JObject) QueryStringArray(path string) ([]string, error) {
	undefined_data, err := QueryJSONInterface(obj.data, path)
	if err != nil {
		return nil, err
	}
	undefined_type := reflect.TypeOf(undefined_data)
	if undefined_type.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the type is invalid, you are asking for a []string but the type is %s", undefined_type)
	}

	if undefined_type.Elem().Kind() != reflect.String {
		return nil, fmt.Errorf("the type is invalid, you are asking for a element type string but the element type is %s", undefined_type.Elem())
	}

	return undefined_data.([]string), nil
}

func (obj *JObject) QueryDecimalArray(path string) ([]float64, error) {
	undefined_data, err := QueryJSONInterface(obj.data, path)
	if err != nil {
		return nil, err
	}
	undefined_type := reflect.TypeOf(undefined_data)
	if undefined_type.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the type is invalid, you are asking for a []float64 but the type is %s", undefined_type)
	}

	if undefined_type.Elem().Kind() != reflect.Float64 {
		return nil, fmt.Errorf("the type is invalid, you are asking for a element type float64 but the element type is %s", undefined_type.Elem())
	}

	return undefined_data.([]float64), nil
}

func (obj *JObject) QueryIntArray(path string) ([]int, error) {

	decimal_array, err := obj.QueryDecimalArray(path)
	if err != nil {
		return nil, err
	}

	var numbers []int
	for _, decimal := range decimal_array {
		numbers = append(numbers, int(decimal))
	}

	return numbers, nil
}

func (obj *JObject) QueryBooleanArray(path string) ([]bool, error) {

	undefined_data, err := QueryJSONInterface(obj.data, path)
	if err != nil {
		return nil, err
	}
	undefined_type := reflect.TypeOf(undefined_data)
	if undefined_type.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the type is invalid, you are asking for a []bool but the type is %s", undefined_type)
	}

	if undefined_type.Elem().Kind() != reflect.Float64 {
		return nil, fmt.Errorf("the type is invalid, you are asking for a element type bool but the element type is %s", undefined_type.Elem())
	}

	return undefined_data.([]bool), nil
}

func (obj *JObject) QueryObject(path string) (*JObject, error) {
	undefined_data, err := QueryJSONInterface(obj.data, path)
	if err != nil {
		return nil, err
	}
	return &JObject{
		data: undefined_data,
	}, nil
}

func (obj *JObject) Keys() ([]string, error) {
	object_type := reflect.TypeOf(obj.data)
	if object_type.Kind() != reflect.Map {
		return nil, fmt.Errorf("you cannot get the keys for type: %s", object_type)
	}

	var keys []string
	switch obj.data.(type) {
	case map[string]interface{}:
		for key := range obj.data.(map[string]interface{}) {
			keys = append(keys, key)
		}
	case map[interface{}]interface{}:
		for key := range obj.data.(map[interface{}]interface{}) {
			keys = append(keys, fmt.Sprint(key))
		}
	}
	return keys, nil
}

func (obj *JObject) Values() ([]string, error) {
	object_type := reflect.TypeOf(obj.data)
	if object_type.Kind() != reflect.Map {
		return nil, fmt.Errorf("you cannot get the value for type: %s", object_type)
	}

	var values []string
	switch obj.data.(type) {
	case map[string]interface{}:
		for _, value := range obj.data.(map[string]interface{}) {
			values = append(values, fmt.Sprint(value))
		}
	case map[interface{}]interface{}:
		for _, value := range obj.data.(map[interface{}]interface{}) {
			values = append(values, fmt.Sprint(value))
		}
	}
	return values, nil
}

func (obj *JObject) ContainsKey(key string) bool {
	result := false
	object_type := reflect.TypeOf(obj.data)
	if object_type.Kind() != reflect.Map {
		return result
	}

	switch obj.data.(type) {
	case map[string]interface{}:
		for k := range obj.data.(map[string]interface{}) {
			if k == key {
				result = true
				return result
			}
		}
	case map[interface{}]interface{}:
		for k := range obj.data.(map[interface{}]interface{}) {
			if k == key {
				result = true
				return result
			}
		}
	}
	return result
}

func (obj *JObject) Raw() interface{} {
	return obj.data
}

func QueryGeneric[T any](object *JObject, path string) (*T, error) {
	undefined_data, err := QueryJSONInterface(object.data, path)
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(undefined_data)
	if err != nil {
		return nil, err
	}

	var generic_value T
	if err := json.Unmarshal(body, &generic_value); err != nil {
		return nil, err
	}
	return &generic_value, nil
}

func QueryGenericSlice[T any](object *JObject, path string) ([]*T, error) {
	undefined_data, err := QueryJSONInterface(object.data, path)
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(undefined_data)
	if err != nil {
		return nil, err
	}

	var generic_value []*T
	if err := json.Unmarshal(body, &generic_value); err != nil {
		return nil, err
	}
	return generic_value, nil
}
