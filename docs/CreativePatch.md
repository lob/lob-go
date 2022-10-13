# CreativePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **interface{}** | Must either be an address ID or an inline object with correct address parameters. | [optional] 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 

## Methods

### NewCreativePatch

`func NewCreativePatch() *CreativePatch`

NewCreativePatch instantiates a new CreativePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativePatchWithDefaults

`func NewCreativePatchWithDefaults() *CreativePatch`

NewCreativePatchWithDefaults instantiates a new CreativePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CreativePatch) GetFrom() interface{}`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreativePatch) GetFromOk() (*interface{}, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreativePatch) SetFrom(v interface{})`

SetFrom sets From field to given value.

### HasFrom

`func (o *CreativePatch) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *CreativePatch) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *CreativePatch) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetDescription

`func (o *CreativePatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreativePatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreativePatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreativePatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreativePatch) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreativePatch) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *CreativePatch) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreativePatch) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreativePatch) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreativePatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


