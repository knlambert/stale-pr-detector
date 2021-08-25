package pkg

import (
	"github.com/knlambert/stale-pr-detector/pkg/output"
	"github.com/pkg/errors"
	"io"
	"os"
)

func CreatePRDetector(
	gitVendor GitClientVendor,
	formatType OutputFormat,
	fileOutput string,
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

	var selectedOutput io.Writer = os.Stdout

	if fileOutput != "" {
		selectedOutput = output.CreateFile(fileOutput)
	}

	return &PRDetector{
		formatter: formatter,
		gitClient: gitClient,
		output:    selectedOutput,
	}, nil
}

type PRDetector struct {
	formatter Formatter
	gitClient Client
	output    io.Writer
}
