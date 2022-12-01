package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LetterTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	mockBankWritable    lob.BankAccountWritable
	mockAccount         lob.BankAccount
	mockVerify          lob.BankAccountVerify
	addressEditableList [7]lob.AddressEditable
	letterEditable      lob.LetterEditable
}

// // Before all
func (suite *LetterTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.mockBankWritable = *lob.NewBankAccountWritable("322271627", "123456789", lob.BANKTYPEENUM_INDIVIDUAL, "Sinead Connor")

	resp, _, err := suite.apiClient.BankAccountsApi.Create(suite.ctx).BankAccountWritable(suite.mockBankWritable).Execute()

	if err != nil {
		panic(err)
	}
	suite.mockAccount = *resp

	verifyAmounts := []int32{11, 35}
	suite.mockVerify = *lob.NewBankAccountVerify(verifyAmounts)
	suite.apiClient.BankAccountsApi.Verify(suite.ctx, suite.mockAccount.Id).BankAccountVerify(suite.mockVerify).Execute()

	suite.addressEditableList = CreateAddressesEditableList()
	suite.letterEditable = *lob.NewLetterEditable(false, suite.addressEditableList[0], suite.addressEditableList[1], GetFileLocation8x11())
}

func (suite *LetterTestSuite) TestLetterCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.LettersApi.Create(suite.ctx).LetterEditable(suite.letterEditable).Execute()

	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
	}
}

func (suite *LetterTestSuite) TestLetterCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.LettersApi.Create(suite.badctx).LetterEditable(suite.letterEditable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *LetterTestSuite) TestLetterRetrieve() {
	t := suite.T()
	createdLetter, _, _ := suite.apiClient.LettersApi.Create(suite.ctx).LetterEditable(suite.letterEditable).Execute()

	resp, _, err := suite.apiClient.LettersApi.Get(suite.ctx, createdLetter.GetId()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetId(), createdLetter.GetId())
	}
}

func (suite *LetterTestSuite) TestLetterRetrieveBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.LettersApi.Get(suite.badctx, "fake id").Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *LetterTestSuite) TestLetterList() {
	t := suite.T()
	resp, _, err := suite.apiClient.LettersApi.List(suite.ctx).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *LetterTestSuite) TestLetterListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.LettersApi.List(suite.ctx).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(3))
	}
}

func (suite *LetterTestSuite) TestLetterListWithNextPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.LettersApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, err := suite.apiClient.LettersApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responeAfter) {
		assert.Greater(t, responeAfter.GetCount(), int32(0))
	}
}

func (suite *LetterTestSuite) TestLetterListWithPrevPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.LettersApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, errAfter := suite.apiClient.LettersApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, errAfter)
	responseBefore, _, err := suite.apiClient.LettersApi.List(suite.ctx).Before(responeAfter.GetPrevPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responseBefore) {
		assert.Greater(t, responseBefore.GetCount(), int32(0))
	}
}

func (suite *LetterTestSuite) TestLetterCancel() {
	t := suite.T()
	createdLetter, _, _ := suite.apiClient.LettersApi.Create(suite.ctx).LetterEditable(suite.letterEditable).Execute()

	resp, _, err := suite.apiClient.LettersApi.Cancel(suite.ctx, createdLetter.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetId(), createdLetter.Id)
	}
}

// // Runs all of the tests
func TestLetterTestSuite(t *testing.T) {
	suite.Run(t, new(LetterTestSuite))
}
