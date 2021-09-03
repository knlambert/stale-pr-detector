package pkg

import (
	"bytes"
	"github.com/golang/mock/gomock"
	pkg_mock "github.com/knlambert/stale-pr-detector/pkg/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"testing"
	"time"
)

type suitePRDetector struct {
	suite.Suite
	ctrl              *gomock.Controller
	stubbedPRDetector *PRDetector
	formatterMock     *pkg_mock.MockFormatter
	gitClientMock     *pkg_mock.MockGitClient
	output            io.Writer
	timeMock          *pkg_mock.MockTimeWrapper

	now time.Time
}

func (s *suitePRDetector) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.formatterMock = pkg_mock.NewMockFormatter(s.ctrl)
	s.gitClientMock = pkg_mock.NewMockGitClient(s.ctrl)
	s.timeMock = pkg_mock.NewMockTimeWrapper(s.ctrl)

	s.now = time.Now()
	s.timeMock.EXPECT().Now().Return(s.now).AnyTimes()
	s.output = &bytes.Buffer{}

	s.stubbedPRDetector = &PRDetector{
		gitClient: s.gitClientMock,
		formatter: s.formatterMock,
		time:      s.timeMock,
		output:    s.output,
	}
}

func TestSuitePRDetector(t *testing.T) {
	suite.Run(t, new(suitePRDetector))
}
