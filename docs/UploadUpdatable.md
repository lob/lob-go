# UploadUpdatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalFilename** | Pointer to **string** | Original filename provided when the upload is created. | [optional] 
**RequiredAddressColumnMapping** | Pointer to [**RequiredAddressColumnMapping**](RequiredAddressColumnMapping.md) |  | [optional] 
**OptionalAddressColumnMapping** | Pointer to [**OptionalAddressColumnMapping**](OptionalAddressColumnMapping.md) |  | [optional] 
**Metadata** | Pointer to [**UploadsMetadata**](UploadsMetadata.md) |  | [optional] 
**MergeVariableColumnMapping** | Pointer to **map[string]interface{}** | The mapping of column headers in your file to the merge variables present in your creative. See our &lt;a href&#x3D;\&quot;https://help.lob.com/print-and-mail/building-a-mail-strategy/campaign-or-triggered-sends/campaign-audience-guide#step-3-map-merge-variable-data-if-applicable-7\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Campaign Audience Guide&lt;/a&gt; for additional details. &lt;br /&gt;If a merge variable has the same \&quot;name\&quot; as a \&quot;key\&quot; in the &#x60;requiredAddressColumnMapping&#x60; or &#x60;optionalAddressColumnMapping&#x60; objects, then they **CANNOT** have a different value in this object. If a different value is provided, then when the campaign is processing it will get overwritten with the mapped value present in the &#x60;requiredAddressColumnMapping&#x60; or &#x60;optionalAddressColumnMapping&#x60; objects. | [optional] 

## Methods

### NewUploadUpdatable

`func NewUploadUpdatable() *UploadUpdatable`

NewUploadUpdatable instantiates a new UploadUpdatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadUpdatableWithDefaults

`func NewUploadUpdatableWithDefaults() *UploadUpdatable`

NewUploadUpdatableWithDefaults instantiates a new UploadUpdatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalFilename

`func (o *UploadUpdatable) GetOriginalFilename() string`

GetOriginalFilename returns the OriginalFilename field if non-nil, zero value otherwise.

### GetOriginalFilenameOk

`func (o *UploadUpdatable) GetOriginalFilenameOk() (*string, bool)`

GetOriginalFilenameOk returns a tuple with the OriginalFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFilename

`func (o *UploadUpdatable) SetOriginalFilename(v string)`

SetOriginalFilename sets OriginalFilename field to given value.

### HasOriginalFilename

`func (o *UploadUpdatable) HasOriginalFilename() bool`

HasOriginalFilename returns a boolean if a field has been set.

### GetRequiredAddressColumnMapping

`func (o *UploadUpdatable) GetRequiredAddressColumnMapping() RequiredAddressColumnMapping`

GetRequiredAddressColumnMapping returns the RequiredAddressColumnMapping field if non-nil, zero value otherwise.

### GetRequiredAddressColumnMappingOk

`func (o *UploadUpdatable) GetRequiredAddressColumnMappingOk() (*RequiredAddressColumnMapping, bool)`

GetRequiredAddressColumnMappingOk returns a tuple with the RequiredAddressColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAddressColumnMapping

`func (o *UploadUpdatable) SetRequiredAddressColumnMapping(v RequiredAddressColumnMapping)`

SetRequiredAddressColumnMapping sets RequiredAddressColumnMapping field to given value.

### HasRequiredAddressColumnMapping

`func (o *UploadUpdatable) HasRequiredAddressColumnMapping() bool`

HasRequiredAddressColumnMapping returns a boolean if a field has been set.

### GetOptionalAddressColumnMapping

`func (o *UploadUpdatable) GetOptionalAddressColumnMapping() OptionalAddressColumnMapping`

GetOptionalAddressColumnMapping returns the OptionalAddressColumnMapping field if non-nil, zero value otherwise.

### GetOptionalAddressColumnMappingOk

`func (o *UploadUpdatable) GetOptionalAddressColumnMappingOk() (*OptionalAddressColumnMapping, bool)`

GetOptionalAddressColumnMappingOk returns a tuple with the OptionalAddressColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAddressColumnMapping

`func (o *UploadUpdatable) SetOptionalAddressColumnMapping(v OptionalAddressColumnMapping)`

SetOptionalAddressColumnMapping sets OptionalAddressColumnMapping field to given value.

### HasOptionalAddressColumnMapping

`func (o *UploadUpdatable) HasOptionalAddressColumnMapping() bool`

HasOptionalAddressColumnMapping returns a boolean if a field has been set.

### GetMetadata

`func (o *UploadUpdatable) GetMetadata() UploadsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UploadUpdatable) GetMetadataOk() (*UploadsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UploadUpdatable) SetMetadata(v UploadsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UploadUpdatable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMergeVariableColumnMapping

`func (o *UploadUpdatable) GetMergeVariableColumnMapping() map[string]interface{}`

GetMergeVariableColumnMapping returns the MergeVariableColumnMapping field if non-nil, zero value otherwise.

### GetMergeVariableColumnMappingOk

`func (o *UploadUpdatable) GetMergeVariableColumnMappingOk() (*map[string]interface{}, bool)`

GetMergeVariableColumnMappingOk returns a tuple with the MergeVariableColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariableColumnMapping

`func (o *UploadUpdatable) SetMergeVariableColumnMapping(v map[string]interface{})`

SetMergeVariableColumnMapping sets MergeVariableColumnMapping field to given value.

### HasMergeVariableColumnMapping

`func (o *UploadUpdatable) HasMergeVariableColumnMapping() bool`

HasMergeVariableColumnMapping returns a boolean if a field has been set.

### SetMergeVariableColumnMappingNil

`func (o *UploadUpdatable) SetMergeVariableColumnMappingNil(b bool)`

 SetMergeVariableColumnMappingNil sets the value for MergeVariableColumnMapping to be an explicit nil

### UnsetMergeVariableColumnMapping
`func (o *UploadUpdatable) UnsetMergeVariableColumnMapping()`

UnsetMergeVariableColumnMapping ensures that no value is present for MergeVariableColumnMapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


