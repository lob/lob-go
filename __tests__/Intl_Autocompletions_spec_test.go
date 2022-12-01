package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IntlAutocompletionsSuite struct {
	suite.Suite
	ctx                         context.Context
	badctx                      context.Context
	configuration               lob.Configuration
	apiClient                   lob.APIClient
	createdAddress              lob.Address
	intlAutocomplationsWritable lob.IntlAutocompletionsWritable
}

// Before all
func (suite *IntlAutocompletionsSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.intlAutocomplationsWritable = *lob.NewIntlAutocompletionsWritable("340 Wat", lob.COUNTRYEXTENDED_CA)
}

func (suite *IntlAutocompletionsSuite) TestIntlAutocompletions() {
	t := suite.T()
	resp, _, err := suite.apiClient.IntlAutocompletionsApi.Autocomplete(suite.ctx).IntlAutocompletionsWritable(suite.intlAutocomplationsWritable).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, len(resp.GetSuggestions()), 0)
		suite.createdAddress.Id = resp.Id
	}
}

func (suite *IntlAutocompletionsSuite) TestIntlAutocompletionsBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.IntlAutocompletionsApi.Autocomplete(suite.badctx).IntlAutocompletionsWritable(suite.intlAutocomplationsWritable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

// Runs all of the tests
func TestIntlAutocompletionTestSuite(t *testing.T) {
	suite.Run(t, new(IntlAutocompletionsSuite))
}
