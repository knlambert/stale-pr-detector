package format

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type suiteJSON struct {
	suite.Suite
	stubbedJSONFormatter *JSONFormatter
}

func (s *suiteJSON) SetupTest() {
	s.stubbedJSONFormatter = &JSONFormatter{}
}

func (s *suiteJSON) TestPrettyPrint() {
	result, err := s.stubbedJSONFormatter.PrettyPrint([]*someUser{{
		Email: "john.doe@gmail.com",
		Address: someAddress{
			Street: "Tree street",
			Town:   "Montreal",
		},
	}})
	assert.NoError(s.T(), err, "should not fail")
	expected := `[{"email":"john.doe@gmail.com","address":{"street":"Tree street","town":"Montreal"}}]`
	assert.Equal(s.T(), string(result), expected)
}

func TestSuiteJSON(t *testing.T) {
	suite.Run(t, new(suiteJSON))
}
