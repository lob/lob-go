package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CheckTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	mockBankWritable    lob.BankAccountWritable
	mockAccount         lob.BankAccount
	mockVerify          lob.BankAccountVerify
	checkEditable       lob.CheckEditable
	addressEditableList [7]lob.AddressEditable
}

// Before all
func (suite *CheckTestSuite) SetupTest() {
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

	useType := *lob.NewNullableChkUseType(lob.CHKUSETYPE_MARKETING.Ptr())
	suite.addressEditableList = CreateAddressesEditableList()
	suite.checkEditable = *lob.NewCheckEditable(suite.addressEditableList[0], suite.addressEditableList[1], *lob.NewNullableString(&suite.mockAccount.Id), 100, useType)
}

func (suite *CheckTestSuite) TestCheckCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.ChecksApi.Create(suite.ctx).CheckEditable(suite.checkEditable).Execute()

	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
		assert.Equal(t, suite.checkEditable.Amount, resp.Amount)
	}
}

func (suite *CheckTestSuite) TestCheckCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.ChecksApi.Create(suite.badctx).CheckEditable(suite.checkEditable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *CheckTestSuite) TestCheckRetrieve() {
	t := suite.T()
	createdCheck, _, _ := suite.apiClient.ChecksApi.Create(suite.ctx).CheckEditable(suite.checkEditable).Execute()

	resp, _, err := suite.apiClient.ChecksApi.Get(suite.ctx, createdCheck.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, createdCheck.Id)
	}
}

func (suite *CheckTestSuite) TestCheckRetrieveBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.ChecksApi.Get(suite.badctx, "fake id").Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *CheckTestSuite) TestCheckList() {
	t := suite.T()
	resp, _, err := suite.apiClient.ChecksApi.List(suite.ctx).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *CheckTestSuite) TestCheckListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.ChecksApi.List(suite.ctx).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(3))
	}
}

func (suite *CheckTestSuite) TestCheckListWithNextPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.ChecksApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, err := suite.apiClient.ChecksApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responeAfter) {
		assert.Greater(t, responeAfter.GetCount(), int32(0))
	}
}

func (suite *CheckTestSuite) TestCheckListWithPrevPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.ChecksApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, errAfter := suite.apiClient.ChecksApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, errAfter)
	responseBefore, _, err := suite.apiClient.ChecksApi.List(suite.ctx).Before(responeAfter.GetPrevPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responseBefore) {
		assert.Greater(t, responseBefore.GetCount(), int32(0))
	}
}

func (suite *CheckTestSuite) TestCheckCancel() {
	t := suite.T()
	createdCheck, _, _ := suite.apiClient.ChecksApi.Create(suite.ctx).CheckEditable(suite.checkEditable).Execute()

	resp, _, err := suite.apiClient.ChecksApi.Cancel(suite.ctx, createdCheck.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetId(), createdCheck.Id)
	}
}

// Runs all of the tests
func TestCheckTestSuite(t *testing.T) {
	suite.Run(t, new(CheckTestSuite))
}
