package github

import (
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"time"
)

func (c *Client) PullRequestsList(
	repositoryURL string,
	lastActivity time.Time,
) (
	[]models.Repository, error,
) {
	//prs, err := c.client.PullRequests.List(
	//	context.Context(),
	//
	//	)
	return nil, nil
}
