package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UsAutocompletionsSuite struct {
	suite.Suite
	ctx                       context.Context
	badctx                    context.Context
	configuration             lob.Configuration
	apiClient                 lob.APIClient
	createdAddress            lob.Address
	usAutocomplationsWritable lob.UsAutocompletionsWritable
}

// Before all
func (suite *UsAutocompletionsSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.usAutocomplationsWritable = *lob.NewUsAutocompletionsWritable("340 Wat")
}

func (suite *UsAutocompletionsSuite) TestIntlAutocompletions() {
	t := suite.T()
	resp, _, err := suite.apiClient.UsAutocompletionsApi.UsAutocompletion(suite.ctx).UsAutocompletionsWritable(suite.usAutocomplationsWritable).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, len(resp.GetSuggestions()), 0)
		suite.createdAddress.Id = resp.Id
	}
}

func (suite *UsAutocompletionsSuite) TestIntlAutocompletionsBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.UsAutocompletionsApi.UsAutocompletion(suite.badctx).UsAutocompletionsWritable(suite.usAutocomplationsWritable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

// Runs all of the tests
func TestUsAutocompletionTestSuite(t *testing.T) {
	suite.Run(t, new(UsAutocompletionsSuite))
}
