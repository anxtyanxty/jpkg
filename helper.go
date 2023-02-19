package jpkg

import (
	"fmt"
	"reflect"
	"strings"
)

func parse_query(query string) []item {
	submatches := query_expr.FindAllStringSubmatch(query, -1)
	matches := []string{}
	for _, m := range submatches {
		matches = append(matches, m[1])
	}

	items := []item{}
	for _, m := range matches {
		data, is_index := remove_quotes(m)
		items = append(items, item{Data: data, IsIndex: is_index})
	}

	return items
}

func remove_quotes(item string) (string, bool) {
	old_item := item
	item = strings.ReplaceAll(item, "\"", "")
	item = strings.ReplaceAll(item, "'", "")
	return item, len(old_item) == len(item)
}

func match_value(data map[interface{}]interface{}, item interface{}) (interface{}, error) {
	for key, value := range data {
		if key == item {
			return value, nil
		}
	}
	return nil, fmt.Errorf("the key for %+v could not be found in the json", item)
}

func is_indexable(kind reflect.Kind) bool {
	return kind == reflect.Map || kind == reflect.Array
}
