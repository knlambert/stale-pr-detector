package git

import (
	"fmt"
	"github.com/knlambert/stale-pr-detector/pkg/git/github"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"github.com/pkg/errors"
	"time"
)

type Vendor string

const (
	VendorGithub Vendor = "github"
)

type Client interface {
	PullRequestsList(
		repositoryURL string,
		lastActivity time.Time,
	) (
		[]models.Repository, error,
	)
}

func CreateClient(vendor Vendor) (Client, error) {
	switch vendor {
	case VendorGithub:
		return github.CreateClient(), nil
	default:
		return nil, errors.New(fmt.Sprintf("vendor %s not supported", vendor))
	}
}
