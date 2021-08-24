package format

import (
	"encoding/json"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"github.com/pkg/errors"
)

func CreateJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}
type JSONFormatter struct {}

func (j *JSONFormatter) PrettyPrintPullRequests(prs []models.PullRequest) ([]byte, error) {
	if encoded, err := json.Marshal(prs); err != nil {
		return nil, errors.Wrap(err, "failed to encode to JSON")
	} else {
		return encoded, nil
	}
}