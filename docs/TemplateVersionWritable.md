# TemplateVersionWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Html** | **string** | An HTML string of less than 100,000 characters to be used as the &#x60;published_version&#x60; of this template. See [here](#section/HTML-Examples) for guidance on designing HTML templates. Please see endpoint specific documentation for any other product-specific HTML details: - [Postcards](https://docs.lob.com/#tag/Postcards/operation/postcard_create) - &#x60;front&#x60; and &#x60;back&#x60; - [Self Mailers](https://docs.lob.com/#tag/Self-Mailers/operation/self_mailer_create) - &#x60;inside&#x60; and &#x60;outside&#x60; - [Letters](https://docs.lob.com/#tag/Letters/operation/letter_create) - &#x60;file&#x60; - [Checks](https://docs.lob.com/#tag/Checks/operation/check_create) - &#x60;check_bottom&#x60; and &#x60;attachment&#x60; - [Cards](https://docs.lob.com/#tag/Cards/operation/card_create) - &#x60;front&#x60; and &#x60;back&#x60;  If there is a syntax error with your variable names within your HTML, then an error will be thrown, e.g. using a &#x60;{{#users}}&#x60; opening tag without the corresponding closing tag &#x60;{{/users}}&#x60;.  | 
**Engine** | Pointer to [**NullableEngineHtml**](EngineHtml.md) |  | [optional] [default to LEGACY]

## Methods

### NewTemplateVersionWritable

`func NewTemplateVersionWritable(html string, ) *TemplateVersionWritable`

NewTemplateVersionWritable instantiates a new TemplateVersionWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVersionWritableWithDefaults

`func NewTemplateVersionWritableWithDefaults() *TemplateVersionWritable`

NewTemplateVersionWritableWithDefaults instantiates a new TemplateVersionWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TemplateVersionWritable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateVersionWritable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateVersionWritable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateVersionWritable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateVersionWritable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateVersionWritable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHtml

`func (o *TemplateVersionWritable) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *TemplateVersionWritable) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *TemplateVersionWritable) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetEngine

`func (o *TemplateVersionWritable) GetEngine() EngineHtml`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *TemplateVersionWritable) GetEngineOk() (*EngineHtml, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *TemplateVersionWritable) SetEngine(v EngineHtml)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *TemplateVersionWritable) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### SetEngineNil

`func (o *TemplateVersionWritable) SetEngineNil(b bool)`

 SetEngineNil sets the value for Engine to be an explicit nil

### UnsetEngine
`func (o *TemplateVersionWritable) UnsetEngine()`

UnsetEngine ensures that no value is present for Engine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


