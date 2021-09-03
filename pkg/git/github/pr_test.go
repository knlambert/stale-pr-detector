package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v38/github"
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"github.com/stretchr/testify/assert"
	"time"
)

func searchIssueSample(sample, startNumber, count int) []*github.Issue {
	creationDate := time.Date(2021, 01, 01, 0, 0, 0, 0, time.UTC)
	lastActivity := creationDate.AddDate(0, 0, 5)

	var result []*github.Issue
	for i := 0; i < count; i++ {
		number := startNumber + i
		result = append(result, map[int]*github.Issue{
			1: {
				Number: github.Int(number),
				State:  github.String("open"),
				Title:  github.String(fmt.Sprintf("my pr from mclaren %d", number)),
				User: &github.User{
					Login: github.String("bstinson"),
				},
				Labels: []*github.Label{
					{Name: github.String("stale")},
					{Name: github.String("high")},
				},
				CreatedAt:     &creationDate,
				UpdatedAt:     &lastActivity,
				RepositoryURL: github.String("https://github.com/google/go-github.git"),
				Reactions:     nil,
				HTMLURL: github.String("some-url"),
			},
			2: {
				Number: github.Int(number),
				State:  github.String("open"),
				Title:  github.String(fmt.Sprintf("My other PR %d", number)),
				User: &github.User{
					Login: github.String("meriksen"),
				},
				Labels: []*github.Label{
					{Name: github.String("ignore")},
				},
				CreatedAt:     &creationDate,
				UpdatedAt:     &lastActivity,
				RepositoryURL: github.String("https://github.com/google/go-github.git"),
				Reactions:     nil,
				HTMLURL: github.String("some-url"),
			},
		}[sample])
	}
	return result
}

func (s *suiteClient) TestPullRequestListWithFilters() {
	//Tear up
	states := []string{"open"}
	labels := []string{"stale", "high"}
	authors := []string{"bstinson"}
	creationDate := time.Date(2021, 01, 01, 0, 0, 0, 0, time.UTC)
	lastActivity := creationDate.AddDate(0, 0, 5)
	expectedIssues := searchIssueSample(1, 1, 1)

	s.githubSearchMock.EXPECT().Issues(
		context.Background(),
		"is:pr repo:google/go-github author:bstinson label:stale label:high state:open updated:<=2021-01-06",
		&github.SearchOptions{
			Sort: "updated",
			ListOptions: github.ListOptions{
				PerPage: 5,
			},
		}).Return(
		&github.IssuesSearchResult{
			Total:  github.Int(1),
			Issues: expectedIssues,
		}, &github.Response{
			NextPage: 0,
		}, nil)

	//Assertions
	prs, err := s.client.PullRequestsList(
		"https://github.com/google/go-github.git",
		&git.PullRequestsListFilters{
			LastActivity: &lastActivity,
			States:       &states,
			Labels:       &labels,
			Authors:      &authors,
		},
	)

	assert.NoError(s.T(), err, "should not raise errors")
	assert.Equal(s.T(), prs, []models.PullRequest{{
		Number:    github.String("1"),
		State:     github.String("open"),
		Title:     github.String("my pr from mclaren 1"),
		Author:    github.String("bstinson"),
		Labels:    []string{"stale", "high"},
		CreatedAt: expectedIssues[0].CreatedAt,
		UpdatedAt: expectedIssues[0].UpdatedAt,
		Repository: &models.Repository{
			URL:   expectedIssues[0].RepositoryURL,
			Owner: github.String("google"),
			Name:  github.String("go-github"),
		},
		Link: github.String("some-url"),
	}})
	s.ctrl.Finish()
}

func (s *suiteClient) TestPullRequestListWithPagination() {
	//Tear up
	batch1 := searchIssueSample(1, 1, 5)
	batch2 := searchIssueSample(1, 6, 2)

	//page 1
	s.githubSearchMock.EXPECT().Issues(
		context.Background(),
		"is:pr repo:google/go-github",
		&github.SearchOptions{
			Sort: "updated",
			ListOptions: github.ListOptions{
				PerPage: 5,
			},
		}).Return(
		&github.IssuesSearchResult{
			Total:  github.Int(5),
			Issues: batch1,
		}, &github.Response{
			NextPage: 1,
		}, nil)
	//page 2
	s.githubSearchMock.EXPECT().Issues(
		context.Background(),
		"is:pr repo:google/go-github",
		&github.SearchOptions{
			Sort: "updated",
			ListOptions: github.ListOptions{
				PerPage: 5,
				Page:    1,
			},
		}).Return(
		&github.IssuesSearchResult{
			Total:  github.Int(2),
			Issues: batch2,
		}, &github.Response{
			NextPage: 0,
		}, nil)

	//Assertions
	prs, err := s.client.PullRequestsList(
		"https://github.com/google/go-github.git",
		&git.PullRequestsListFilters{},
	)

	assert.NoError(s.T(), err, "should not raise errors")
	assert.Equal(s.T(), 7, len(prs))
	s.ctrl.Finish()
}

func (s *suiteClient) TestPullRequestListWithRetry() {
	//Tear up

	//First requests fails.
	s.githubSearchMock.EXPECT().Issues(
		context.Background(),
		"is:pr repo:google/go-github",
		&github.SearchOptions{
			Sort: "updated",
			ListOptions: github.ListOptions{
				PerPage: 5,
			},
		}).Return(
		&github.IssuesSearchResult{},
		&github.Response{},
		&github.RateLimitError{
			Rate:     github.Rate{},
			Response: nil,
			Message:  "Some rate limit error",
		})

	s.time.EXPECT().Sleep(10 * time.Second).Times(1)

	//Second works.
	s.githubSearchMock.EXPECT().Issues(
		context.Background(),
		"is:pr repo:google/go-github",
		&github.SearchOptions{
			Sort: "updated",
			ListOptions: github.ListOptions{
				PerPage: 5,
			},
		}).Return(
		&github.IssuesSearchResult{
			Total:  github.Int(5),
			Issues: searchIssueSample(1, 1, 1),
		}, &github.Response{}, nil)

	//Assertions
	prs, err := s.client.PullRequestsList(
		"https://github.com/google/go-github.git",
		&git.PullRequestsListFilters{},
	)

	assert.NoError(s.T(), err, "should not raise errors")
	assert.Equal(s.T(), 1, len(prs))

	s.ctrl.Finish()
}
