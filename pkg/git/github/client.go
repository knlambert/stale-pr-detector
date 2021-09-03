package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v38/github"
	"github.com/knlambert/stale-pr-detector/pkg/std"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"regexp"
	"time"
)

//CreateClient creates a new client instance.
func CreateClient() *Client {
	var httpClient *http.Client
	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	if githubAccessToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubAccessToken},
		)
		ctx := context.Background()
		httpClient = oauth2.NewClient(ctx, ts)
	}

	client := github.NewClient(httpClient)

	return &Client{
		search:          client.Search,
		defaultPageSize: 100,
		time:            std.CreateTimeWrapper(),
	}
}

//Client wrapper around a third party library to interact with Github.
type Client struct {
	search          goGithubSearch
	defaultPageSize int
	time            timeWrapper
}

//timeWrapper describes a wrapper around the time package.
type timeWrapper interface {
	Sleep(d time.Duration)
}

//goGithubSearch describes the external go-github search API.
// https://pkg.go.dev/github.com/google/go-github/v38/github#SearchService
type goGithubSearch interface {
	Issues(ctx context.Context, query string, opts *github.SearchOptions) (
		*github.IssuesSearchResult, *github.Response, error,
	)
}

//parseRepositoryURL takes a git repo URL and extracts its owner and repository information.
func (c *Client) parseRepositoryURL(url string) (owner string, repo string, err error) {
	var regex = regexp.MustCompile(`(?:git@|https?://)?[\w.@]+[/:]?(?:repos/)?(\S+)/([\w-]+)(?:\.git)?`)

	if groups := regex.FindStringSubmatch(url); groups != nil {
		owner, repo = groups[1], groups[2]
	} else {
		err = errors.New(fmt.Sprintf("'%s' is not a repository URL", url))
	}

	return
}

//waitIfRateLimitError waits if there is a rate error.
func (c *Client) waitIfRateLimitError(err error) bool {
	if _, ok := err.(*github.RateLimitError); ok {
		c.time.Sleep(10 * time.Second)
		return true
	}
	return false
}
