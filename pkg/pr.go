package pkg

import (
	"github.com/pkg/errors"
)

func CreatePRDetector(
	gitVendor GitClientVendor,
	formatType OutputFormat,
) (*PRDetector, error) {
	var gitClient Client
	var formatter Formatter
	var err error

	if gitClient, err = CreateGitClient(gitVendor); err != nil {
		return nil, errors.Wrapf(err, "failed to initialize PR detector")
	}

	if formatter, err = CreateFormatter(formatType); err != nil {
		return nil, errors.Wrapf(err, "failed to initialize PR detector")
	}

	return &PRDetector{
		formatter: formatter,
		gitClient: gitClient,
	}, nil
}

type PRDetector struct {
	formatter Formatter
	gitClient Client
}
