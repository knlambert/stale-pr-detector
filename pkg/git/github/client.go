package github

import (
	"fmt"
	"github.com/google/go-github/v38/github"
	"github.com/pkg/errors"
	"regexp"
)

func CreateClient() *Client {
	return &Client{
		client: github.NewClient(nil),
	}
}

type Client struct {
	client *github.Client
}

//parseRepositoryURL takes a git repo URL and extracts its owner and repository.
func (c *Client) parseRepositoryURL(url string) (owner string, repo string, err error) {
	var regex = regexp.MustCompile(`(?:git@|https?://)?[\w.@]+[/:]?(\S+)/(\S+)\.git`)

	if groups := regex.FindStringSubmatch(url); groups != nil {
		owner, repo = groups[1], groups[2]
	} else {
		err = errors.New(fmt.Sprintf("'%s' is not a repository URL", url))
	}

	return
}

