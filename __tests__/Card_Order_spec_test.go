package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CardOrderTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	addressEditableList [7]lob.AddressEditable
	cardOrderEditable   lob.CardOrderEditable
	mockCard            lob.Card
}

// Before all
func (suite *CardOrderTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.addressEditableList = CreateAddressesEditableList() // AddressEditable

	var cardEditable = *lob.NewCardEditable(GetFileLocation())
	resp, _, err := suite.apiClient.CardsApi.CardCreate(suite.ctx).CardEditable(cardEditable).Execute()

	if err != nil {
		panic(err)
	}
	suite.mockCard = *resp
	suite.cardOrderEditable = *lob.NewCardOrderEditable(10000)

}

func (suite *CardOrderTestSuite) TestCardOrderCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.CardOrdersApi.CardOrderCreate(suite.ctx, suite.mockCard.GetId()).CardOrderEditable(suite.cardOrderEditable).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
		assert.Equal(t, resp.GetCardId(), suite.mockCard.GetId())
	}
}

func (suite *CardOrderTestSuite) TestCardOrderCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.CardOrdersApi.CardOrderCreate(suite.badctx, suite.mockCard.GetId()).CardOrderEditable(suite.cardOrderEditable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *CardOrderTestSuite) TestCardOrderRetrieve() {
	t := suite.T()
	createdCardOrder, _, _ := suite.apiClient.CardOrdersApi.CardOrderCreate(suite.ctx, suite.mockCard.GetId()).CardOrderEditable(suite.cardOrderEditable).Execute()

	resp, _, err := suite.apiClient.CardOrdersApi.CardOrdersRetrieve(suite.ctx, createdCardOrder.GetCardId()).Execute()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

// // Runs all of the tests
func TestCardOrderTestSuite(t *testing.T) {
	suite.Run(t, new(CardOrderTestSuite))
}
