package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PostcardsTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	addressEditableList [7]lob.AddressEditable
	postcardEditable    lob.PostcardEditable
}

// Before all
func (suite *PostcardsTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.addressEditableList = CreateAddressesEditableList() // AddressEditable

	suite.postcardEditable = *lob.NewPostcardEditable(suite.addressEditableList[0], GetFileLocation4x6(), GetFileLocation4x6())
}

func (suite *PostcardsTestSuite) TestPostcardsCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.PostcardsApi.PostcardCreate(suite.ctx).PostcardEditable(suite.postcardEditable).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
		v := suite.postcardEditable.GetTo().(lob.AddressEditable)
		assert.Equal(t, v.AddressLine1, resp.GetTo().AddressLine1)
	}
}

func (suite *PostcardsTestSuite) TestPostcardCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.PostcardsApi.PostcardCreate(suite.badctx).PostcardEditable(suite.postcardEditable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *PostcardsTestSuite) TestPostcardRetrieve() {
	t := suite.T()
	createdPostcard, _, _ := suite.apiClient.PostcardsApi.PostcardCreate(suite.ctx).PostcardEditable(suite.postcardEditable).Execute()

	resp, _, err := suite.apiClient.PostcardsApi.PostcardRetrieve(suite.ctx, createdPostcard.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, createdPostcard.Id)
	}
}

func (suite *PostcardsTestSuite) TestPostcardRetrieveBadApiKey() {
	t := suite.T()
	createdPostcard, _, _ := suite.apiClient.PostcardsApi.PostcardCreate(suite.ctx).PostcardEditable(suite.postcardEditable).Execute()

	_, _, err := suite.apiClient.CardsApi.CardRetrieve(suite.badctx, createdPostcard.Id).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *PostcardsTestSuite) TestPostcardList() {
	t := suite.T()
	resp, _, err := suite.apiClient.PostcardsApi.PostcardsList(suite.ctx).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *PostcardsTestSuite) TestPostcardListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.PostcardsApi.PostcardsList(suite.ctx).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(3))
	}
}

func (suite *PostcardsTestSuite) TestPostcardsListWithNextPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.PostcardsApi.PostcardsList(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, err := suite.apiClient.PostcardsApi.PostcardsList(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responeAfter) {
		assert.Greater(t, responeAfter.GetCount(), int32(0))
	}
}

func (suite *PostcardsTestSuite) TestPostcardsListWithPrevPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.PostcardsApi.PostcardsList(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, errAfter := suite.apiClient.PostcardsApi.PostcardsList(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, errAfter)
	responseBefore, _, err := suite.apiClient.PostcardsApi.PostcardsList(suite.ctx).Before(responeAfter.GetPrevPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responseBefore) {
		assert.Greater(t, responseBefore.GetCount(), int32(0))
	}
}

func (suite *PostcardsTestSuite) TestPostcardDelete() {
	t := suite.T()
	createdPostcard, _, _ := suite.apiClient.PostcardsApi.PostcardCreate(suite.ctx).PostcardEditable(suite.postcardEditable).Execute()

	resp, _, err := suite.apiClient.PostcardsApi.PostcardDelete(suite.ctx, createdPostcard.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetId(), createdPostcard.Id)
	}
}

// Runs all of the tests
func TestPostcardsTestSuite(t *testing.T) {
	suite.Run(t, new(PostcardsTestSuite))
}
