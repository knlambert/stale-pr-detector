package github

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type suiteClient struct {
	suite.Suite
	client *Client
}

func (s *suiteClient) SetupTest() {
	s.client = &Client{}
}

func (s *suiteClient) TestParseRepositoryURL() {
	for _, url := range []string{
		"https://github.com/google/go-github.git",
		"git@github.com:google/go-github.git",
		"github.com/google/go-github.git",
		"github.com/google/go-github",
	} {
		owner, repo, err := s.client.ParseRepositoryURL(url)
		assert.Equal(s.T(), "google", owner)
		assert.Equal(s.T(), "go-github", repo)
		assert.NoErrorf(s.T(), err, "the regex should match")
	}
}

func TestSuiteClient(t *testing.T) {
	suite.Run(t, new(suiteClient))
}
