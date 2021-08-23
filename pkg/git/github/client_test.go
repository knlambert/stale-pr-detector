package github

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientTestSuite struct {
	suite.Suite
	client *Client
}

func (c *ClientTestSuite) SetupTest() {
	c.client = &Client{}
}

func (c *ClientTestSuite) TestParseRepositoryURL() {
	for _, url := range []string{
		"https://github.com/google/go-github.git",
		"git@github.com:google/go-github.git",
		"github.com/google/go-github.git",
	} {
		owner, repo, err := c.client.parseRepositoryURL(url)
		assert.Equal(c.T(), "google", owner)
		assert.Equal(c.T(), "go-github", repo)
		assert.NoErrorf(c.T(), err, "the regex should match")
	}
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
