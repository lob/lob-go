package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TemplateTestSuite struct {
	suite.Suite
	ctx                 context.Context
	badctx              context.Context
	configuration       lob.Configuration
	apiClient           lob.APIClient
	mockBankWritable    lob.BankAccountWritable
	mockAccount         lob.BankAccount
	mockVerify          lob.BankAccountVerify
	templateWritable    lob.TemplateWritable
	addressEditableList [7]lob.AddressEditable
}

// Before all
func (suite *TemplateTestSuite) SetupTest() {
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
	suite.templateWritable = *lob.NewTemplateWritable("<html>Updated HTML for Template 1/html>")
}

func (suite *TemplateTestSuite) TestTemplateCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.TemplatesApi.Create(suite.ctx).TemplateWritable(suite.templateWritable).Execute()

	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
	}
}

func (suite *TemplateTestSuite) TestTemplateCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.TemplatesApi.Create(suite.badctx).TemplateWritable(suite.templateWritable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *TemplateTestSuite) TestTemplateRetrieve() {
	t := suite.T()
	createdTemplate, _, _ := suite.apiClient.TemplatesApi.Create(suite.ctx).TemplateWritable(suite.templateWritable).Execute()

	resp, _, err := suite.apiClient.TemplatesApi.Get(suite.ctx, createdTemplate.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, createdTemplate.Id)
	}
}

func (suite *TemplateTestSuite) TestTemplateRetrieveBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.TemplatesApi.Get(suite.badctx, "fake id").Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *TemplateTestSuite) TestTemplateList() {
	t := suite.T()
	resp, _, err := suite.apiClient.TemplatesApi.List(suite.ctx).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *TemplateTestSuite) TestTemplateListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.TemplatesApi.List(suite.ctx).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(3))
	}
}

func (suite *TemplateTestSuite) TestTemplateListWithNextPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.TemplatesApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, err := suite.apiClient.TemplatesApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responeAfter) {
		assert.Greater(t, responeAfter.GetCount(), int32(0))
	}
}

func (suite *TemplateTestSuite) TestTemplateListWithPrevPageToken() {
	t := suite.T()
	firstResponse, _, firstErr := suite.apiClient.TemplatesApi.List(suite.ctx).Limit(1).Execute()
	assert.Nil(t, firstErr)
	responeAfter, _, errAfter := suite.apiClient.TemplatesApi.List(suite.ctx).After(firstResponse.GetNextPageToken()).Execute()
	assert.Nil(t, errAfter)
	responseBefore, _, err := suite.apiClient.TemplatesApi.List(suite.ctx).Before(responeAfter.GetPrevPageToken()).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, responseBefore) {
		assert.Greater(t, responseBefore.GetCount(), int32(0))
	}
}

func (suite *TemplateTestSuite) TestTemplateDelete() {
	t := suite.T()
	createdTemplate, _, _ := suite.apiClient.TemplatesApi.Create(suite.ctx).TemplateWritable(suite.templateWritable).Execute()

	resp, _, err := suite.apiClient.TemplatesApi.Delete(suite.ctx, createdTemplate.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetId(), createdTemplate.Id)
	}
}

// Runs all of the tests
func TestTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateTestSuite))
}
