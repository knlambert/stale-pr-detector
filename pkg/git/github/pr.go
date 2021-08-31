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

	//Base filters
	var queryFilters = []string{
		"is:pr",
		fmt.Sprintf("repo:%s/%s", owner, repo),
	}

	buildFilter("label", filters.Labels, &queryFilters)
	buildFilter("state", filters.States, &queryFilters)

	if filters.LastActivity != nil {
		queryFilters = append(
			queryFilters, fmt.Sprintf("updated:<=%s", filters.LastActivity.Format("2006-01-02")),
		)
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

	var decoded = make([]models.PullRequest, 0)
	for _, pr := range results {
		var labels []string

		for i := range pr.Labels {
			labels = append(labels, *pr.Labels[i].Name)
		}

		number := strconv.Itoa(*pr.Number)

		owner, repoName, _ := c.ParseRepositoryURL(*pr.RepositoryURL)
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
				Name: &repoName,
				Owner: &owner,
			},
		})
	}

	return decoded
}
