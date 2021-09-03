package pkg

import (
	"bytes"
	"github.com/knlambert/stale-pr-detector/pkg/git"
	"github.com/knlambert/stale-pr-detector/pkg/models"
	"github.com/stretchr/testify/assert"
	"time"
)

func (s *suitePRDetector) TestParseLastActivityThreeDays() {
	date, err := s.stubbedPRDetector.parseLastActivity("3d")
	expected := s.now.Add(3 * 24 * time.Hour)
	assert.NoError(s.T(), err, "date format is correct")
	assert.Equal(s.T(), expected.Format("YY-MM-DD"), date.Format("YY-MM-DD"))
}

func (s *suitePRDetector) TestParseLastActivityFiveMonth() {
	date, err := s.stubbedPRDetector.parseLastActivity("5m")
	expected := s.now.Add(5 * 30 * 24 * time.Hour)
	assert.NoError(s.T(), err, "date format is correct")
	assert.Equal(s.T(), expected.Format("YY-MM-DD"), date.Format("YY-MM-DD"))
}

func (s *suitePRDetector) TestParseLastActivityOneYear() {
	date, err := s.stubbedPRDetector.parseLastActivity("1y")
	expected := s.now.Add(365 * 24 * time.Hour)
	assert.NoError(s.T(), err, "date format is correct")
	assert.Equal(s.T(), expected.Format("YY-MM-DD"), date.Format("YY-MM-DD"))
}

func (s *suitePRDetector) TestStaleListTestBase() {
	//Tear up
	states := []string{"open"}
	labels := []string{"high"}
	lastActivity := s.now.AddDate(0, 0, -30)

	number := "1"
	result := []models.PullRequest{{Number: &number}}
	s.gitClientMock.EXPECT().PullRequestsList(
		"https://github.com/kubernetes/kubernetes",
		&git.PullRequestsListFilters{
			States:       &states,
			Labels:       &labels,
			LastActivity: &lastActivity,
		}).Return(result, nil)

	s.formatterMock.EXPECT().PrettyPrint(result).Times(1).Return(
		[]byte("some pr pretty print"), nil,
	)

	//Assertions
	err := s.stubbedPRDetector.StaleList(
		[]string{"https://github.com/kubernetes/kubernetes"},
		[]string{"high"},
		"30d",
	)

	assert.NoError(s.T(), err, "Should not fail")
	assert.Equal(s.T(), "some pr pretty print", s.output.(*bytes.Buffer).String())
	s.ctrl.Finish()
}
