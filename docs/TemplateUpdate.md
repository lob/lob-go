# TemplateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**PublishedVersion** | Pointer to **string** | Unique identifier prefixed with &#x60;vrsn_&#x60;. | [optional] 

## Methods

### NewTemplateUpdate

`func NewTemplateUpdate() *TemplateUpdate`

NewTemplateUpdate instantiates a new TemplateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateUpdateWithDefaults

`func NewTemplateUpdateWithDefaults() *TemplateUpdate`

NewTemplateUpdateWithDefaults instantiates a new TemplateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TemplateUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPublishedVersion

`func (o *TemplateUpdate) GetPublishedVersion() string`

GetPublishedVersion returns the PublishedVersion field if non-nil, zero value otherwise.

### GetPublishedVersionOk

`func (o *TemplateUpdate) GetPublishedVersionOk() (*string, bool)`

GetPublishedVersionOk returns a tuple with the PublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersion

`func (o *TemplateUpdate) SetPublishedVersion(v string)`

SetPublishedVersion sets PublishedVersion field to given value.

### HasPublishedVersion

`func (o *TemplateUpdate) HasPublishedVersion() bool`

HasPublishedVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


