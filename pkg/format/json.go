package format

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

func CreateJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

type JSONFormatter struct{}

func (j *JSONFormatter) PrettyPrint(
	values interface{},
) ([]byte, error) {
	var decoded []map[string]interface{}

	if encoded, err := json.Marshal(values); err != nil {
		return nil, errors.Wrap(err, "failed to encode to JSON (Marshal)")
	} else {
		if err := json.Unmarshal(encoded, &decoded); err != nil {
			return nil, errors.Wrap(err, "failed to encode to JSON (Unmarshal)")
		}

		var result []map[string]interface{}
		for i := range decoded {
			decoded[i] = flatMap(decoded[i], []string{})

			obj := map[string]interface{}{}

			for k, v := range decoded[i] {
				obj[k] = v
			}

			result = append(result, obj)
		}

		if encodedResult, err := json.Marshal(result); err != nil {
			return nil, errors.Wrap(err, "failed to encode to JSON (Marshal)")
		} else {
			return encodedResult, nil
		}
	}
}

func containsString(s []string, v string) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}
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

func mergeMap(obj1, obj2 map[string]interface{}) map[string]interface{} {
	for k := range obj2 {
		obj1[k] = obj2[k]
	}
	return obj1
}
