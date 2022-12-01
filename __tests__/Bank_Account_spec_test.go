package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BankAccountTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	addressEditableList [7]lob.AddressEditable
	bankAccountWritable lob.BankAccountWritable
}

// Before all
func (suite *BankAccountTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.addressEditableList = CreateAddressesEditableList() // AddressEditable

	suite.bankAccountWritable = *lob.NewBankAccountWritable("322271627", "123456789", lob.BANKTYPEENUM_INDIVIDUAL, "Sinead Connor")
}

func (suite *BankAccountTestSuite) TestBankAccountCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.BankAccountsApi.BankAccountCreate(suite.ctx).BankAccountWritable(suite.bankAccountWritable).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
		assert.Equal(t, suite.bankAccountWritable.AccountNumber, resp.AccountNumber)
	}
}

func (suite *BankAccountTestSuite) TestBankAccountCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.BankAccountsApi.BankAccountCreate(suite.badctx).BankAccountWritable(suite.bankAccountWritable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *BankAccountTestSuite) TestBankAccountRetrieve() {
	t := suite.T()
	createdBA, _, _ := suite.apiClient.BankAccountsApi.BankAccountCreate(suite.ctx).BankAccountWritable(suite.bankAccountWritable).Execute()

	resp, _, err := suite.apiClient.BankAccountsApi.BankAccountRetrieve(suite.ctx, createdBA.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, createdBA.Id)
	}
}

func (suite *BankAccountTestSuite) TestBankAccountRetrieveBadApiKey() {
	t := suite.T()
	createdBA, _, _ := suite.apiClient.BankAccountsApi.BankAccountCreate(suite.ctx).BankAccountWritable(suite.bankAccountWritable).Execute()

	_, _, err := suite.apiClient.CardsApi.CardRetrieve(suite.badctx, createdBA.Id).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *BankAccountTestSuite) TestBankAccountList() {
	t := suite.T()
	resp, _, err := suite.apiClient.BankAccountsApi.BankAccountsList(suite.ctx).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *BankAccountTestSuite) TestBankAccountListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.BankAccountsApi.BankAccountsList(suite.ctx).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(3))
	}
}

func (suite *BankAccountTestSuite) TestBankAccountsListWithNextPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.BankAccountsApi.BankAccountsList(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, err := suite.apiClient.BankAccountsApi.BankAccountsList(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responeAfter) {
		assert.Greater(t, responeAfter.GetCount(), int32(0))
	}
}

func (suite *BankAccountTestSuite) TestBankAccountsListWithPrevPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.BankAccountsApi.BankAccountsList(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, errAfter := suite.apiClient.BankAccountsApi.BankAccountsList(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, errAfter)
	responseBefore, _, err := suite.apiClient.BankAccountsApi.BankAccountsList(suite.ctx).Before(responeAfter.GetPrevPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responseBefore) {
		assert.Greater(t, responseBefore.GetCount(), int32(0))
	}
}

func (suite *BankAccountTestSuite) TestBankAccountVerify() {
	t := suite.T()
	createdBA, _, _ := suite.apiClient.BankAccountsApi.BankAccountCreate(suite.ctx).BankAccountWritable(suite.bankAccountWritable).Execute()

	verifyAmounts := []int32{11, 35}
	mockVerify := *lob.NewBankAccountVerify(verifyAmounts)

	resp, _, err := suite.apiClient.BankAccountsApi.BankAccountVerify(suite.ctx, createdBA.Id).BankAccountVerify(mockVerify).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetVerified(), true)
	}
}

func (suite *BankAccountTestSuite) TestBankAccountDelete() {
	t := suite.T()
	createdBA, _, _ := suite.apiClient.BankAccountsApi.BankAccountCreate(suite.ctx).BankAccountWritable(suite.bankAccountWritable).Execute()

	resp, _, err := suite.apiClient.BankAccountsApi.BankAccountDelete(suite.ctx, createdBA.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetId(), createdBA.Id)
	}
}

// Runs all of the tests
func TestBankAccountTestSuite(t *testing.T) {
	suite.Run(t, new(BankAccountTestSuite))
}
