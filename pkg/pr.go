package pkg

import (
	"github.com/knlambert/stale-pr-detector/pkg/output"
	"github.com/knlambert/stale-pr-detector/pkg/std"
	"github.com/pkg/errors"
	"io"
	"os"
)

func CreatePRDetector(
	gitVendor GitClientVendor,
	formatType OutputFormat,
	fileOutput string,
) (*PRDetector, error) {
	var gitClient GitClient
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
		time:      std.CreateTime(),
	}, nil
}

type PRDetector struct {
	formatter Formatter
	gitClient GitClient
	output    io.Writer
	time      Time
}
