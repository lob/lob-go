# TemplateVersionUpdatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Engine** | Pointer to [**NullableEngineHtml**](EngineHtml.md) |  | [optional] [default to ENGINEHTML_LEGACY]

## Methods

### NewTemplateVersionUpdatable

`func NewTemplateVersionUpdatable() *TemplateVersionUpdatable`

NewTemplateVersionUpdatable instantiates a new TemplateVersionUpdatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVersionUpdatableWithDefaults

`func NewTemplateVersionUpdatableWithDefaults() *TemplateVersionUpdatable`

NewTemplateVersionUpdatableWithDefaults instantiates a new TemplateVersionUpdatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TemplateVersionUpdatable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateVersionUpdatable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateVersionUpdatable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateVersionUpdatable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateVersionUpdatable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateVersionUpdatable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEngine

`func (o *TemplateVersionUpdatable) GetEngine() EngineHtml`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *TemplateVersionUpdatable) GetEngineOk() (*EngineHtml, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *TemplateVersionUpdatable) SetEngine(v EngineHtml)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *TemplateVersionUpdatable) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### SetEngineNil

`func (o *TemplateVersionUpdatable) SetEngineNil(b bool)`

 SetEngineNil sets the value for Engine to be an explicit nil

### UnsetEngine
`func (o *TemplateVersionUpdatable) UnsetEngine()`

UnsetEngine ensures that no value is present for Engine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


