package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v38/github"
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func (c *Client) PullRequestsList(
	repositoryURL string,
	filters *git.PullRequestsListFilters,
) (
	[]models.PullRequest, error,
) {
	//TODO: Pagination
	owner, repo, err := c.ParseRepositoryURL(repositoryURL)

	if err != nil {
		return nil, err
	}

	var queryFilters = []string{
		"is:pr",
		fmt.Sprintf("repo:%s/%s", owner, repo),
	}

	if filters.Labels != nil {
		var labelsFilters []string
		for _, label := range *filters.Labels {
			labelsFilters = append(labelsFilters, fmt.Sprintf("label:%s", label))
		}

		queryFilters = append(queryFilters, strings.Join(labelsFilters, " "))
	}

	query := strings.Join(queryFilters, " ")

	if results, _, err := c.client.Search.Issues(context.Background(), query, &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}); err == nil {
		return c.issuesToPullRequests(results.Issues), nil
	} else {
		return nil, errors.Wrapf(err, "failed to list PRs for %s/%s", owner, repo)
	}
}

func (c *Client) issuesToPullRequests(results []*github.Issue) []models.PullRequest {
	var decoded []models.PullRequest
	for _, pr := range results {
		var labels []string

		for i := range pr.Labels {
			labels = append(labels, *pr.Labels[i].Name)
		}

		number := strconv.Itoa(*pr.Number)

		decoded = append(decoded, models.PullRequest{
			Number:    &number,
			State:     pr.State,
			Title:     pr.Title,
			Author:    pr.User.Login,
			Labels:    labels,
			CreatedAt: pr.CreatedAt,
			UpdatedAt: pr.UpdatedAt,
			Repository: &models.Repository{
				URL: pr.RepositoryURL,
			},
		})
	}

	return decoded
}
