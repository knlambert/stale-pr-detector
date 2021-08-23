package crawler

import (
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"time"
)

type parallelizedImpl struct {
	gitClient git.Client
}

func (g *parallelizedImpl) PullRequestsList(
	repositories []models.Repository,
	lastActivity time.Time,
) ([]models.PullRequest, error) {
	return nil, nil
}
