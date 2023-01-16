package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SelfMailersTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	addressEditableList [7]lob.AddressEditable
	selfMailerEditable  lob.SelfMailerEditable
}

// Before all
func (suite *SelfMailersTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.addressEditableList = CreateAddressesEditableList() // AddressEditable

	useType := *lob.NewNullableSfmUseType(lob.SFMUSETYPE_MARKETING.Ptr())
	suite.selfMailerEditable = *lob.NewSelfMailerEditable(suite.addressEditableList[0], GetFileLocation6x18(), GetFileLocation6x18(), useType)
}

func (suite *SelfMailersTestSuite) TestSelfMailersCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.SelfMailersApi.Create(suite.ctx).SelfMailerEditable(suite.selfMailerEditable).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
	}
}

func (suite *SelfMailersTestSuite) TestSelfMailersCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.SelfMailersApi.Create(suite.badctx).SelfMailerEditable(suite.selfMailerEditable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *SelfMailersTestSuite) TestSelfMailersRetrieve() {
	t := suite.T()
	createdSM, _, _ := suite.apiClient.SelfMailersApi.Create(suite.ctx).SelfMailerEditable(suite.selfMailerEditable).Execute()

	resp, _, err := suite.apiClient.SelfMailersApi.Get(suite.ctx, createdSM.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, createdSM.Id)
	}
}

func (suite *SelfMailersTestSuite) TestSelfMailerRetrieveBadApiKey() {
	t := suite.T()
	createdSM, _, _ := suite.apiClient.SelfMailersApi.Create(suite.ctx).SelfMailerEditable(suite.selfMailerEditable).Execute()

	_, _, err := suite.apiClient.SelfMailersApi.Get(suite.badctx, createdSM.Id).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *SelfMailersTestSuite) TestSelfMailerList() {
	t := suite.T()
	resp, _, err := suite.apiClient.SelfMailersApi.List(suite.ctx).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *SelfMailersTestSuite) TestPostcardListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.SelfMailersApi.List(suite.ctx).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(3))
	}
}

func (suite *SelfMailersTestSuite) TestSelfMailersListWithNextPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.SelfMailersApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, err := suite.apiClient.SelfMailersApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responeAfter) {
		assert.Greater(t, responeAfter.GetCount(), int32(0))
	}
}

func (suite *SelfMailersTestSuite) TestSelfMailersListWithPrevPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.SelfMailersApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, errAfter := suite.apiClient.SelfMailersApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, errAfter)
	responseBefore, _, err := suite.apiClient.SelfMailersApi.List(suite.ctx).Before(responeAfter.GetPrevPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responseBefore) {
		assert.Greater(t, responseBefore.GetCount(), int32(0))
	}
}

func (suite *SelfMailersTestSuite) TestPostcardDelete() {
	t := suite.T()
	createdSM, _, _ := suite.apiClient.SelfMailersApi.Create(suite.ctx).SelfMailerEditable(suite.selfMailerEditable).Execute()

	resp, _, err := suite.apiClient.SelfMailersApi.Delete(suite.ctx, createdSM.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetId(), createdSM.Id)
	}
}

// Runs all of the tests
func TestSelfMailersTestSuite(t *testing.T) {
	suite.Run(t, new(SelfMailersTestSuite))
}
