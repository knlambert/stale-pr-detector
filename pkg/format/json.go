package format

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

//CreateJSONFormatter creates an instance of JSONFormatter.
func CreateJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

//JSONFormatter implements methods to pretty print golang objects into JSON objects.
type JSONFormatter struct{}

//PrettyPrint encode the golang struct into bytes.
func (j *JSONFormatter) PrettyPrint(
	values interface{},
) ([]byte, error) {
	if encoded, err := json.Marshal(values); err != nil {
		return nil, errors.Wrap(err, "failed to encode to JSON (Marshal)")
	} else {
		return encoded, nil
	}
}

//flatMap flattens the deep maps structures, resulting in a one level map.
func flatMap(obj map[string]interface{}, prefix []string) map[string]interface{} {
	var result = map[string]interface{}{}

	for key, val := range obj {

		if convertedVal, ok := val.(map[string]interface{}); ok {
			result = mergeMap(result, flatMap(convertedVal, append(prefix, key)))
		} else if _, ok := val.([]map[string]interface{}); ok {

		} else {
			result[strings.Join(append(prefix, key), ".")] = val
		}
	}

	return result
}

//mergeMap merges two maps together.
func mergeMap(obj1, obj2 map[string]interface{}) map[string]interface{} {
	for k := range obj2 {
		obj1[k] = obj2[k]
	}
	return obj1
}
