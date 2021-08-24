package pkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type suitePRDetector struct {
	suite.Suite
	stubbedPRDetector *PRDetector
}

func (s *suitePRDetector) SetupTest() {
	s.stubbedPRDetector = &PRDetector{
		crawler:   nil,
		formatter: nil,
	}
}

func TestSuitePRDetector(t *testing.T) {
	suite.Run(t, new(suitePRDetector))
}
