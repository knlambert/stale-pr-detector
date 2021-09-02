package format

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type suiteText struct {
	suite.Suite
	stubbedTextFormatter *TextFormatter
}

func (s *suiteText) SetupTest() {
	s.stubbedTextFormatter = &TextFormatter{
		headers: []string{
			"email", "address.street", "address.town",
		},
	}
}

type someUserContainer struct {
	Users []someUser `json:"users"`
}

type someUser struct {
	Email   string      `json:"email"`
	Address someAddress `json:"address"`
}

type someAddress struct {
	Street string `json:"street"`
	Town   string `json:"town"`
}

func (s *suiteText) TestPrettyPrint() {
	result, err := s.stubbedTextFormatter.PrettyPrint([]*someUser{{
		Email: "john.doe@gmail.com",
		Address: someAddress{
			Street: "Tree street",
			Town:   "Montreal",
		},
	}})
	assert.NoError(s.T(), err, "should not fail")
	expected := `+--------------------+----------------+--------------+
|       EMAIL        | ADDRESS STREET | ADDRESS TOWN |
+--------------------+----------------+--------------+
| john.doe@gmail.com | Tree street    | Montreal     |
+--------------------+----------------+--------------+
`
	assert.Equal(s.T(), string(result), expected)
}

func TestSuiteText(t *testing.T) {
	suite.Run(t, new(suiteText))
}
