package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v38/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"regexp"
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

	return &Client{
		client: github.NewClient(httpClient),
	}
}

//Client wrapper around a third party library to interact with Github.
type Client struct {
	client *github.Client
}

//ParseRepositoryURL takes a git repo URL and extracts its owner and repository.
func (c *Client) ParseRepositoryURL(url string) (owner string, repo string, err error) {
	var regex = regexp.MustCompile(`(?:git@|https?://)?[\w.@]+[/:]?(?:repos/)?(\S+)/([\w-]+)(?:\.git)?`)

	if groups := regex.FindStringSubmatch(url); groups != nil {
		owner, repo = groups[1], groups[2]
	} else {
		err = errors.New(fmt.Sprintf("'%s' is not a repository URL", url))
	}

	return
}
