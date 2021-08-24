package pkg

import (
	"fmt"
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/knlambert/stale-pr-detector/pkg/git/github"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"github.com/pkg/errors"
)

type GitClientVendor string

const (
	GitClientVendorGithub GitClientVendor = "github"
)

type Client interface {
	PullRequestsList(
		repositoryURL string,
		filters *git.PullRequestsListFilters,
	) ([]models.PullRequest, error)
}

func CreateGitClient(vendor GitClientVendor) (Client, error) {
	switch vendor {
	case GitClientVendorGithub:
		return github.CreateClient(), nil
	default:
		return nil, errors.New(fmt.Sprintf("vendor %s not supported", vendor))
	}
}
