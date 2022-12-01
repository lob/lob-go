package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ReverseGeocodeLookupsTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	addressEditableList [7]lob.AddressEditable
	location            lob.Location
}

// Before all
func (suite *ReverseGeocodeLookupsTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.addressEditableList = CreateAddressesEditableList() // AddressEditable

	suite.location = *lob.NewLocationWithDefaults()
	suite.location.SetLatitude(37.777456)
	suite.location.SetLongitude(122.393039)
}

func (suite *ReverseGeocodeLookupsTestSuite) TestReverseGeocodeLookup() {
	t := suite.T()
	resp, _, err := suite.apiClient.ReverseGeocodeLookupsApi.ReverseGeocodeLookup(suite.ctx).Location(suite.location).Execute()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func (suite *ReverseGeocodeLookupsTestSuite) TestReverseGeocodeLookupBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.ReverseGeocodeLookupsApi.ReverseGeocodeLookup(suite.badctx).Location(suite.location).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

// Runs all of the tests
func TestReverseGeocodeLookupsTestSuite(t *testing.T) {
	suite.Run(t, new(ReverseGeocodeLookupsTestSuite))
}
