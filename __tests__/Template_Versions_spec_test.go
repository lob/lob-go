package tests

import (
	"context"
	"os"
	"testing"

	lob "github.com/lob/lob-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TemplateVersionTestSuite struct {
	suite.Suite
	ctx                     context.Context
	badctx                  context.Context
	configuration           lob.Configuration
	apiClient               lob.APIClient
	mockBankWritable        lob.BankAccountWritable
	mockAccount             lob.BankAccount
	mockVerify              lob.BankAccountVerify
	templateVersionWritable lob.TemplateVersionWritable
	template                lob.Template
	addressEditableList     [7]lob.AddressEditable
}

// Before all
func (suite *TemplateVersionTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

	suite.badctx = context.Background()
	suite.badctx = context.WithValue(suite.badctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: "bad api key"})

	suite.configuration = *lob.NewConfiguration()

	suite.apiClient = *lob.NewAPIClient(&suite.configuration)

	suite.mockBankWritable = *lob.NewBankAccountWritable("322271627", "123456789", lob.BANKTYPEENUM_INDIVIDUAL, "Sinead Connor")

	resp, _, err := suite.apiClient.BankAccountsApi.BankAccountCreate(suite.ctx).BankAccountWritable(suite.mockBankWritable).Execute()

	if err != nil {
		panic(err)
	}
	suite.mockAccount = *resp

	verifyAmounts := []int32{11, 35}
	suite.mockVerify = *lob.NewBankAccountVerify(verifyAmounts)
	suite.apiClient.BankAccountsApi.BankAccountVerify(suite.ctx, suite.mockAccount.Id).BankAccountVerify(suite.mockVerify).Execute()

	suite.addressEditableList = CreateAddressesEditableList()
	var mockTemplateWritable = *lob.NewTemplateWritable("<html>Updated HTML for Template 1/html>")

	createdTemplate, _, _ := suite.apiClient.TemplatesApi.CreateTemplate(suite.ctx).TemplateWritable(mockTemplateWritable).Execute()
	suite.template = *createdTemplate
	suite.templateVersionWritable = *lob.NewTemplateVersionWritable("<html>Updated HTML for Template 1/html>")
}

func (suite *TemplateVersionTestSuite) TestTemplateCreate() {
	t := suite.T()
	resp, _, err := suite.apiClient.TemplateVersionsApi.CreateTemplateVersion(suite.ctx, suite.template.Id).TemplateVersionWritable(suite.templateVersionWritable).Execute()

	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Id)
	}
}

func (suite *TemplateVersionTestSuite) TestTemplateCreateBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.TemplateVersionsApi.CreateTemplateVersion(suite.badctx, suite.template.Id).TemplateVersionWritable(suite.templateVersionWritable).Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *TemplateVersionTestSuite) TestTemplateVersionsRetrieve() {
	t := suite.T()
	createdTV, _, _ := suite.apiClient.TemplateVersionsApi.CreateTemplateVersion(suite.ctx, suite.template.Id).TemplateVersionWritable(suite.templateVersionWritable).Execute()

	resp, _, err := suite.apiClient.TemplateVersionsApi.TemplateVersionRetrieve(suite.ctx, suite.template.Id, createdTV.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Id, createdTV.Id)
	}
}

func (suite *TemplateVersionTestSuite) TestTemplateVersionsRetrieveBadApiKey() {
	t := suite.T()
	_, _, err := suite.apiClient.TemplateVersionsApi.TemplateVersionRetrieve(suite.badctx, "fake id", "fake version id").Execute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "401 Unauthorized", err.Error())
	}
}

func (suite *TemplateVersionTestSuite) TestVersionTemplateList() {
	t := suite.T()
	resp, _, err := suite.apiClient.TemplateVersionsApi.TemplateVersionsList(suite.ctx, suite.template.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Greater(t, resp.GetCount(), int32(0))
	}
}

func (suite *TemplateVersionTestSuite) TestTemplateVersionListWithIncludeParameter() {
	t := suite.T()
	resp, _, err := suite.apiClient.TemplateVersionsApi.TemplateVersionsList(suite.ctx, suite.template.Id).Limit(3).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetCount(), int32(1))
	}
}

func (suite *TemplateVersionTestSuite) TestTemplateDelete() {
	t := suite.T()
	createdTV, _, _ := suite.apiClient.TemplateVersionsApi.CreateTemplateVersion(suite.ctx, suite.template.Id).TemplateVersionWritable(suite.templateVersionWritable).Execute()

	resp, _, err := suite.apiClient.TemplateVersionsApi.TemplateVersionDelete(suite.ctx, suite.template.Id, createdTV.Id).Execute()
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.GetId(), createdTV.Id)
	}
}

// Runs all of the tests
func TestTemplateVersionsTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateVersionTestSuite))
}
