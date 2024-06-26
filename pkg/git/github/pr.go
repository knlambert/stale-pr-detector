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

//PullRequestsList lists pull requests from Github.
func (c *Client) PullRequestsList(
	repositoryURL string,
	filters *git.PullRequestsListFilters,
) (
	[]models.PullRequest, error,
) {
	owner, repo, err := c.parseRepositoryURL(repositoryURL)

	if err != nil {
		return nil, err
	}
	//Base filters
	var queryFilters = []string{
		"is:pr",
		fmt.Sprintf("repo:%s/%s", owner, repo),
	}

	buildFilter("author", filters.Authors, &queryFilters)
	buildFilter("label", filters.Labels, &queryFilters)
	buildFilter("state", filters.States, &queryFilters)

	if filters.LastActivity != nil {
		queryFilters = append(
			queryFilters, fmt.Sprintf("updated:<=%s", filters.LastActivity.Format("2006-01-02")),
		)
	}

	query := strings.Join(queryFilters, " ")
	opts := &github.SearchOptions{
		Sort: "updated",
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
			//Rate limit errors can happen. Wait and retry.
			if retry := c.waitIfRateLimitError(err); retry {
				continue
			}
			return nil, errors.Wrapf(err, "failed to list PRs for %s/%s", owner, repo)
		}
	}

	return prs, nil
}

//issuesToPullRequests converts github pull requests structs into an equivalent domain object.
func (c *Client) issuesToPullRequests(results []*github.Issue) []models.PullRequest {

	var decoded = make([]models.PullRequest, 0)
	for _, pr := range results {
		var labels []string

		for i := range pr.Labels {
			labels = append(labels, *pr.Labels[i].Name)
		}

		number := strconv.Itoa(*pr.Number)

		owner, repoName, _ := c.parseRepositoryURL(*pr.RepositoryURL)

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
			Link: pr.HTMLURL,
		})
	}

	return decoded
}
