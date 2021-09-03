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
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: c.defaultPageSize,
		},
	}
	var prs []models.PullRequest

	for {
		if results, response, err := c.search.Issues(context.Background(), query, opts); err == nil {
			prs = append(prs, c.issuesToPullRequests(results.Issues)...)
			//Next page = 0 when there is no more pages.
			if response.NextPage == 0 {
				break
			}
			opts.Page = response.NextPage
		} else {
			return nil, errors.Wrapf(err, "failed to list PRs for %s/%s", owner, repo)
		}
	}

	return prs, nil
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
				URL:   pr.RepositoryURL,
				Name:  &repoName,
				Owner: &owner,
			},
		})
	}

	return decoded
}
