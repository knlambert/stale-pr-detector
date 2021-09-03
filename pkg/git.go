package pkg

import (
	"fmt"
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/knlambert/stale-pr-detector/pkg/git/github"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"github.com/pkg/errors"
)

//GitClientVendor is an enum describing the available git client depending on the chosen vendor.
type GitClientVendor string

const (
	GitClientVendorGithub GitClientVendor = "github"
)

//GitClient describes how a git client should be implemented depending on the vendor in use.
type GitClient interface {
	PullRequestsList(
		repositoryURL string,
		filters *git.PullRequestsListFilters,
	) ([]models.PullRequest, error)
}

//CreateGitClient creates a git client instance depending on the chosen vendor.
func CreateGitClient(vendor GitClientVendor) (GitClient, error) {
	switch vendor {
	case GitClientVendorGithub:
		return github.CreateClient(), nil
	default:
		return nil, errors.New(fmt.Sprintf("vendor %s not supported", vendor))
	}
}
