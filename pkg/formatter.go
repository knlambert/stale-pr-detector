package pkg

import (
	"fmt"
	"github.com/knlambert/stale-pr-detector/pkg/format"
	"github.com/pkg/errors"
)

//Formatter describes how an object should be printed depending on the required implementation.
type Formatter interface {
	PrettyPrint(values interface{}) ([]byte, error)
}

//OutputFormat is an enum describing the available formats of printing.
type OutputFormat string

const (
	OutputFormatJSON OutputFormat = "json"
	OutputFormatText OutputFormat = "text"
)

//CreateFormatter creates a formatter depending on the required format.
func CreateFormatter(f OutputFormat) (Formatter, error) {
	switch f {
	case OutputFormatJSON:
		return format.CreateJSONFormatter(), nil
	case OutputFormatText:
		return format.CreateTextFormatter(), nil
	default:
		return nil, errors.New(fmt.Sprintf("format %s not supported", f))
	}
}
