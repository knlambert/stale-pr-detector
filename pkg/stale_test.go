package pkg

import (
	"github.com/stretchr/testify/assert"
	"time"
)

func (s *suitePRDetector) TestParseLastActivityThreeDays() {
	date, err := s.stubbedPRDetector.parseLastActivity("3d")
	expected := time.Now().Add(3 * 24 * time.Hour)
	assert.NoError(s.T(), err, "date format is correct")
	assert.Equal(s.T(), expected.Format("YY-MM-DD"), date.Format("YY-MM-DD"))
}

func (s *suitePRDetector) TestParseLastActivityFiveMonth() {
	date, err := s.stubbedPRDetector.parseLastActivity("5m")
	expected := time.Now().Add(5 * 30 * 24 * time.Hour)
	assert.NoError(s.T(), err, "date format is correct")
	assert.Equal(s.T(), expected.Format("YY-MM-DD"), date.Format("YY-MM-DD"))
}

func (s *suitePRDetector) TestParseLastActivityOneYear() {
	date, err := s.stubbedPRDetector.parseLastActivity("1y")
	expected := time.Now().Add(365 * 24 * time.Hour)
	assert.NoError(s.T(), err, "date format is correct")
	assert.Equal(s.T(), expected.Format("YY-MM-DD"), date.Format("YY-MM-DD"))
}
