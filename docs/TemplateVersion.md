# TemplateVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;vrsn_&#x60;. | 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Html** | **string** | An HTML string of less than 100,000 characters to be used as the &#x60;published_version&#x60; of this template. See [here](#section/HTML-Examples) for guidance on designing HTML templates. Please see endpoint specific documentation for any other product-specific HTML details: - [Postcards](https://docs.lob.com/#tag/Postcards/operation/postcard_create) - &#x60;front&#x60; and &#x60;back&#x60; - [Self Mailers](https://docs.lob.com/#tag/Self-Mailers/operation/self_mailer_create) - &#x60;inside&#x60; and &#x60;outside&#x60; - [Letters](https://docs.lob.com/#tag/Letters/operation/letter_create) - &#x60;file&#x60; - [Checks](https://docs.lob.com/#tag/Checks/operation/check_create) - &#x60;check_bottom&#x60; and &#x60;attachment&#x60; - [Cards](https://docs.lob.com/#tag/Cards/operation/card_create) - &#x60;front&#x60; and &#x60;back&#x60;  If there is a syntax error with your variable names within your HTML, then an error will be thrown, e.g. using a &#x60;{{#users}}&#x60; opening tag without the corresponding closing tag &#x60;{{/users}}&#x60;.  | 
**Engine** | Pointer to [**NullableEngineHtml**](EngineHtml.md) |  | [optional] [default to LEGACY]
**SuggestJsonEditor** | Pointer to **bool** | Used by frontend, true if the template uses advanced features.  | [optional] 
**MergeVariables** | Pointer to **map[string]interface{}** | Used by frontend, an object representing the keys of every merge variable present in the template. It has one key named &#39;keys&#39;, and its value is an array of strings.  | [optional] 
**DateCreated** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | [optional] 
**DateModified** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | [optional] 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Object** | Pointer to **string** | Value is resource type. | [optional] [default to "version"]

## Methods

### NewTemplateVersion

`func NewTemplateVersion(id string, html string, ) *TemplateVersion`

NewTemplateVersion instantiates a new TemplateVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVersionWithDefaults

`func NewTemplateVersionWithDefaults() *TemplateVersion`

NewTemplateVersionWithDefaults instantiates a new TemplateVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateVersion) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *TemplateVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateVersion) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateVersion) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHtml

`func (o *TemplateVersion) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *TemplateVersion) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *TemplateVersion) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetEngine

`func (o *TemplateVersion) GetEngine() EngineHtml`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *TemplateVersion) GetEngineOk() (*EngineHtml, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *TemplateVersion) SetEngine(v EngineHtml)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *TemplateVersion) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### SetEngineNil

`func (o *TemplateVersion) SetEngineNil(b bool)`

 SetEngineNil sets the value for Engine to be an explicit nil

### UnsetEngine
`func (o *TemplateVersion) UnsetEngine()`

UnsetEngine ensures that no value is present for Engine, not even an explicit nil
### GetSuggestJsonEditor

`func (o *TemplateVersion) GetSuggestJsonEditor() bool`

GetSuggestJsonEditor returns the SuggestJsonEditor field if non-nil, zero value otherwise.

### GetSuggestJsonEditorOk

`func (o *TemplateVersion) GetSuggestJsonEditorOk() (*bool, bool)`

GetSuggestJsonEditorOk returns a tuple with the SuggestJsonEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestJsonEditor

`func (o *TemplateVersion) SetSuggestJsonEditor(v bool)`

SetSuggestJsonEditor sets SuggestJsonEditor field to given value.

### HasSuggestJsonEditor

`func (o *TemplateVersion) HasSuggestJsonEditor() bool`

HasSuggestJsonEditor returns a boolean if a field has been set.

### GetMergeVariables

`func (o *TemplateVersion) GetMergeVariables() map[string]interface{}`

GetMergeVariables returns the MergeVariables field if non-nil, zero value otherwise.

### GetMergeVariablesOk

`func (o *TemplateVersion) GetMergeVariablesOk() (*map[string]interface{}, bool)`

GetMergeVariablesOk returns a tuple with the MergeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariables

`func (o *TemplateVersion) SetMergeVariables(v map[string]interface{})`

SetMergeVariables sets MergeVariables field to given value.

### HasMergeVariables

`func (o *TemplateVersion) HasMergeVariables() bool`

HasMergeVariables returns a boolean if a field has been set.

### GetDateCreated

`func (o *TemplateVersion) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TemplateVersion) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TemplateVersion) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TemplateVersion) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateModified

`func (o *TemplateVersion) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *TemplateVersion) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *TemplateVersion) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.

### HasDateModified

`func (o *TemplateVersion) HasDateModified() bool`

HasDateModified returns a boolean if a field has been set.

### GetDeleted

`func (o *TemplateVersion) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *TemplateVersion) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *TemplateVersion) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *TemplateVersion) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetObject

`func (o *TemplateVersion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TemplateVersion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TemplateVersion) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *TemplateVersion) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


