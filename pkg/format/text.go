package format

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"strings"
)

//CreateTextFormatter creates a new TextFormatter.
func CreateTextFormatter() *TextFormatter {
	return &TextFormatter{
		headers: []string{
			"number", "title", "author", "labels", "created_at", "updated_at", "repository.owner", "repository.name",
		},
	}
}

//TextFormatter implements method to pretty print golang objects into text.
type TextFormatter struct {
	headers []string
}

//PrettyPrint pretty prints objects to tables.
func (t *TextFormatter) PrettyPrint(
	values interface{},
) ([]byte, error) {
	var objs []map[string]interface{}

	encoded, err := json.Marshal(values)

	if err != nil {
		return nil, errors.Wrap(err, "failed to encode to map")
	}

	if err := json.Unmarshal(encoded, &objs); err != nil {
		return nil, errors.Wrap(err, "failed to decode to map")
	}

	var flat []map[string]interface{}

	for _, obj := range objs {
		flatObj := flatMap(obj, nil)
		flat = append(flat, flatObj)
	}

	return t.toTable(flat)
}

//toTable converts a list of map into a 2 dimensions table.
func (t *TextFormatter) toTable(
	items []map[string]interface{},
) ([]byte, error) {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(t.headers)

	for i := range items {
		var line []string

		for _, c := range t.headers {

			switch v := items[i][c].(type) {
			case string:
				line = append(line, v)
			case []interface{}:
				line = append(line, strings.Join(sliceInterfaceToString(v), ", "))
			default:
				panic("unsupported case")
			}

		}

		table.Append(line)
	}

	table.Render()

	return []byte(tableString.String()), nil
}

//sliceInterfaceToString converts a list of interfaces to a list of strings.
func sliceInterfaceToString(s []interface{}) []string {
	var result []string

	for i := range s {
		result = append(result, s[i].(string))
	}

	return result
}
