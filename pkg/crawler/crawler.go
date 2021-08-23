package crawler

import (
	"fmt"
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"github.com/pkg/errors"
	"time"
)

type Crawler interface {
	PullRequestsList(
		repositories []models.Repository,
		lastActivity time.Time,
	) ([]models.PullRequest, error)
}

type Type string

const (
	Parallel Type = "parallel"
)

func CreateCrawler(
	gitClient git.Client,
	crawlerType Type,
) (Crawler, error) {
	switch crawlerType {
	case Parallel:
		return &parallelizedImpl{
			gitClient: gitClient,
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("crawler %s not supported", crawlerType))
	}
}
