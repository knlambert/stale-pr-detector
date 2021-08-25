package pkg

import (
	"fmt"
	"github.com/knlambert/stale-pr-detector/pkg/format"
	"github.com/pkg/errors"
)

type Formatter interface {
	PrettyPrint(value interface{}) ([]byte, error)
}

type OutputFormat string

const (
	OutputFormatJSON OutputFormat = "json"
	OutputFormatText OutputFormat = "text"
)

func CreateFormatter(f OutputFormat) (Formatter, error) {
	switch f {
	case OutputFormatJSON:
		return format.CreateJSONFormatter(), nil
	default:
		return nil, errors.New(fmt.Sprintf("format %s not supported", f))
	}
}
