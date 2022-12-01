package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ZipLookupsTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	addressEditableList [7]lob.AddressEditable
	zipEditable         lob.ZipEditable
}

// Before all
func (suite *ZipLookupsTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.addressEditableList = CreateAddressesEditableList() // AddressEditable

	suite.zipEditable = *lob.NewZipEditable()
	suite.zipEditable.SetZipCode("94107")
}

func (suite *ZipLookupsTestSuite) TestZipLookup() {
	t := suite.T()
	resp, _, err := suite.apiClient.ZipLookupsApi.Lookup(suite.ctx).ZipEditable(suite.zipEditable).Execute()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func (suite *ZipLookupsTestSuite) TestZipLookupBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.ZipLookupsApi.Lookup(suite.badctx).ZipEditable(suite.zipEditable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

// Runs all of the tests
func TestZipLookupsTestSuite(t *testing.T) {
	suite.Run(t, new(ZipLookupsTestSuite))
}
