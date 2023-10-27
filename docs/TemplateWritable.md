# TemplateWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Html** | **string** | An HTML string of less than 100,000 characters to be used as the &#x60;published_version&#x60; of this template. See [here](#section/HTML-Examples) for guidance on designing HTML templates. Please see endpoint specific documentation for any other product-specific HTML details: - [Postcards](https://docs.lob.com/#tag/Postcards/operation/postcard_create) - &#x60;front&#x60; and &#x60;back&#x60; - [Self Mailers](https://docs.lob.com/#tag/Self-Mailers/operation/self_mailer_create) - &#x60;inside&#x60; and &#x60;outside&#x60; - [Letters](https://docs.lob.com/#tag/Letters/operation/letter_create) - &#x60;file&#x60; - [Checks](https://docs.lob.com/#tag/Checks/operation/check_create) - &#x60;check_bottom&#x60; and &#x60;attachment&#x60; - [Cards](https://docs.lob.com/#tag/Cards/operation/card_create) - &#x60;front&#x60; and &#x60;back&#x60;  If there is a syntax error with your variable names within your HTML, then an error will be thrown, e.g. using a &#x60;{{#users}}&#x60; opening tag without the corresponding closing tag &#x60;{{/users}}&#x60;.  | 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**Engine** | Pointer to [**NullableEngineHtml**](EngineHtml.md) |  | [optional] [default to ENGINEHTML_LEGACY]

## Methods

### NewTemplateWritable

`func NewTemplateWritable(html string, ) *TemplateWritable`

NewTemplateWritable instantiates a new TemplateWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWritableWithDefaults

`func NewTemplateWritableWithDefaults() *TemplateWritable`

NewTemplateWritableWithDefaults instantiates a new TemplateWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TemplateWritable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateWritable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateWritable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateWritable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateWritable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateWritable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHtml

`func (o *TemplateWritable) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *TemplateWritable) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *TemplateWritable) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetMetadata

`func (o *TemplateWritable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TemplateWritable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TemplateWritable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TemplateWritable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEngine

`func (o *TemplateWritable) GetEngine() EngineHtml`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *TemplateWritable) GetEngineOk() (*EngineHtml, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *TemplateWritable) SetEngine(v EngineHtml)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *TemplateWritable) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### SetEngineNil

`func (o *TemplateWritable) SetEngineNil(b bool)`

 SetEngineNil sets the value for Engine to be an explicit nil

### UnsetEngine
`func (o *TemplateWritable) UnsetEngine()`

UnsetEngine ensures that no value is present for Engine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


