package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BillingGroupTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	addressEditableList [7]lob.AddressEditable
	bgEditable          lob.BillingGroupEditable
}

// Before all
func (suite *BillingGroupTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.addressEditableList = CreateAddressesEditableList() // AddressEditable

	suite.bgEditable = *lob.NewBillingGroupEditable("Test Billing group")
}

func (suite *BillingGroupTestSuite) TestBillingGroupCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.BillingGroupsApi.BillingGroupCreate(suite.ctx).BillingGroupEditable(suite.bgEditable).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
		assert.Equal(t, suite.bgEditable.Name, resp.GetName())
	}
}

func (suite *BillingGroupTestSuite) TestBillingGroupCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.BillingGroupsApi.BillingGroupCreate(suite.badctx).BillingGroupEditable(suite.bgEditable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *BillingGroupTestSuite) TestBillingGroupRetrieve() {
	t := suite.T()
	createdBG, _, _ := suite.apiClient.BillingGroupsApi.BillingGroupCreate(suite.ctx).BillingGroupEditable(suite.bgEditable).Execute()

	resp, _, err := suite.apiClient.BillingGroupsApi.BillingGroupRetrieve(suite.ctx, createdBG.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, createdBG.Id)
	}
}

func (suite *BillingGroupTestSuite) TestBillingGroupRetrieveBadApiKey() {
	t := suite.T()
	createdBG, _, _ := suite.apiClient.BillingGroupsApi.BillingGroupCreate(suite.ctx).BillingGroupEditable(suite.bgEditable).Execute()

	_, _, err := suite.apiClient.BillingGroupsApi.BillingGroupRetrieve(suite.badctx, createdBG.Id).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *BillingGroupTestSuite) TestBillingGroupList() {
	t := suite.T()
	resp, _, err := suite.apiClient.BillingGroupsApi.BillingGroupsList(suite.ctx).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *BillingGroupTestSuite) TestBillingGroupListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.BillingGroupsApi.BillingGroupsList(suite.ctx).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(3))
	}
}

func (suite *BillingGroupTestSuite) TestBillingGroupUpdate() {
	t := suite.T()
	createdBG, _, _ := suite.apiClient.BillingGroupsApi.BillingGroupCreate(suite.ctx).BillingGroupEditable(suite.bgEditable).Execute()

	suite.bgEditable.SetDescription("BG updated")
	resp, _, err := suite.apiClient.BillingGroupsApi.BillingGroupUpdate(suite.ctx, createdBG.Id).BillingGroupEditable(suite.bgEditable).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetDescription(), "BG updated")
	}
}

// Runs all of the tests
func TestBillingGroupTestSuite(t *testing.T) {
	suite.Run(t, new(BillingGroupTestSuite))
}
