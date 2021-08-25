package format

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func CreateJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}
type JSONFormatter struct {}

func (j *JSONFormatter) PrettyPrint(value interface{}) ([]byte, error) {
	if encoded, err := json.Marshal(value); err != nil {
		return nil, errors.Wrap(err, "failed to encode to JSON")
	} else {
		return encoded, nil
	}
}