package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CardsTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	addressEditableList [7]lob.AddressEditable
	cardEditable        lob.CardEditable
}

// Before all
func (suite *CardsTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.addressEditableList = CreateAddressesEditableList() // AddressEditable

	suite.cardEditable = *lob.NewCardEditable(GetFileLocation())
}

func (suite *CardsTestSuite) TestCardsCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.CardsApi.Create(suite.ctx).CardEditable(suite.cardEditable).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
		assert.Equal(t, suite.cardEditable.Description, resp.Description)
	}
}

func (suite *CardsTestSuite) TestCardCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.CardsApi.Create(suite.badctx).CardEditable(suite.cardEditable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *CardsTestSuite) TestCardRetrieve() {
	t := suite.T()
	createdCard, _, _ := suite.apiClient.CardsApi.Create(suite.ctx).CardEditable(suite.cardEditable).Execute()

	resp, _, err := suite.apiClient.CardsApi.Get(suite.ctx, createdCard.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, createdCard.Id)
	}
}

func (suite *CardsTestSuite) TestCardRetrieveBadApiKey() {
	t := suite.T()
	createdCard, _, _ := suite.apiClient.CardsApi.Create(suite.ctx).CardEditable(suite.cardEditable).Execute()

	_, _, err := suite.apiClient.CardsApi.Get(suite.badctx, createdCard.Id).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *CardsTestSuite) TestCardList() {
	t := suite.T()
	resp, _, err := suite.apiClient.CardsApi.List(suite.ctx).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *CardsTestSuite) TestCardListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.CardsApi.List(suite.ctx).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(3))
	}
}

func (suite *CardsTestSuite) TestCardsListWithNextPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.CardsApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, err := suite.apiClient.CardsApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responeAfter) {
		assert.Greater(t, responeAfter.GetCount(), int32(0))
	}
}

func (suite *CardsTestSuite) TestCardsListWithPrevPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.CardsApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, errAfter := suite.apiClient.CardsApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, errAfter)
	responseBefore, _, err := suite.apiClient.CardsApi.List(suite.ctx).Before(responeAfter.GetPrevPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responseBefore) {
		assert.Greater(t, responseBefore.GetCount(), int32(0))
	}
}

func (suite *CardsTestSuite) TestCardUpdate() {
	t := suite.T()
	createdCard, _, _ := suite.apiClient.CardsApi.Create(suite.ctx).CardEditable(suite.cardEditable).Execute()
	var updatedCard = *lob.NewCardUpdatable()
	updatedCard.SetDescription("Card was updated")

	resp, _, err := suite.apiClient.CardsApi.Update(suite.ctx, createdCard.Id).CardUpdatable(updatedCard).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetDescription(), updatedCard.GetDescription())
	}
}

func (suite *CardsTestSuite) TestCardDelete() {
	t := suite.T()
	createdCard, _, _ := suite.apiClient.CardsApi.Create(suite.ctx).CardEditable(suite.cardEditable).Execute()

	resp, _, err := suite.apiClient.CardsApi.Delete(suite.ctx, createdCard.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetId(), createdCard.Id)
	}
}

// Runs all of the tests
func TestCardsTestSuite(t *testing.T) {
	suite.Run(t, new(CardsTestSuite))
}
