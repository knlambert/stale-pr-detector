package pkg

import (
	"github.com/knlambert/stale-pr-detector/pkg/crawler"
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/pkg/errors"
)

func CreatePRDetector(
	gitVendor git.Vendor,
	crawlerType crawler.Type,
) (*PRDetector, error) {
	var gitClient git.Client
	var crawl crawler.Crawler
	var err error

	if gitClient, err = git.CreateClient(gitVendor); err != nil {
		return nil, errors.Wrapf(err, "failed to initialize PR detector")
	}

	if crawl, err = crawler.CreateCrawler(gitClient, crawlerType); err != nil {
		return nil, errors.Wrapf(err, "failed to initialize PR detector")
	}

	return &PRDetector{
		crawler: crawl,
	}, nil
}

type PRDetector struct {
	crawler crawler.Crawler
}
