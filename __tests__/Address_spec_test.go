package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AddressTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	addressEditableList [7]lob.AddressEditable
	createdAddress      lob.Address
}

// Before all
func (suite *AddressTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.addressEditableList = CreateAddressesEditableList() // AddressEditable
}

func (suite *AddressTestSuite) Testcreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.AddressesApi.Create(suite.ctx).AddressEditable(suite.addressEditableList[0]).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, suite.addressEditableList[0].AddressLine1, resp.AddressLine1)
		suite.createdAddress.Id = resp.Id
	}
}

func (suite *AddressTestSuite) TestcreateBadParameter() {
	t := suite.T()
	_, _, err := suite.apiClient.AddressesApi.Create(suite.ctx).AddressEditable(suite.addressEditableList[6]).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "422 Unprocessable Entity", err.Error())
	}
}

func (suite *AddressTestSuite) TestcreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.AddressesApi.Create(suite.badctx).AddressEditable(suite.addressEditableList[0]).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *AddressTestSuite) TestAddressRetrieve() {
	t := suite.T()
	resp, _, err := suite.apiClient.AddressesApi.Get(suite.ctx, *suite.createdAddress.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, suite.createdAddress.Id)
	}
}

func (suite *AddressTestSuite) TestAddressRetrieveBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.AddressesApi.Get(suite.badctx, *suite.createdAddress.Id).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *AddressTestSuite) TestAddressList() {
	t := suite.T()
	resp, _, err := suite.apiClient.AddressesApi.List(suite.ctx).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *AddressTestSuite) TestAddressListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.AddressesApi.List(suite.ctx).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(3))
	}
}

func (suite *AddressTestSuite) TestAddressListWithNextPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.AddressesApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, err := suite.apiClient.AddressesApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responeAfter) {
		assert.Greater(t, responeAfter.GetCount(), int32(0))
	}
}

func (suite *AddressTestSuite) TestAddressListWithPrevPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.AddressesApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, errAfter := suite.apiClient.AddressesApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, errAfter)
	responseBefore, _, err := suite.apiClient.AddressesApi.List(suite.ctx).Before(responeAfter.GetPrevPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responseBefore) {
		assert.Greater(t, responseBefore.GetCount(), int32(0))
	}
}

func (suite *AddressTestSuite) TestAddressDelete() {
	t := suite.T()
	resp, _, err := suite.apiClient.AddressesApi.Delete(suite.ctx, *suite.createdAddress.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, suite.createdAddress.Id)
	}
}

// Runs all of the tests
func TestAddressTestSuite(t *testing.T) {
	suite.Run(t, new(AddressTestSuite))
}
