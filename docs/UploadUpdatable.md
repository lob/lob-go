# UploadUpdatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnMapping** | Pointer to **map[string]interface{}** | The mapping of column headers in your file to Lob-required fields for the resource created. See our &lt;a href&#x3D;\&quot;https://help.lob.com/best-practices/campaign-audience-guide\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Campaign Audience Guide&lt;/a&gt; for additional details. | [optional] 
**State** | Pointer to [**UploadState**](UploadState.md) |  | [optional] [default to DRAFT]
**OriginalFilename** | Pointer to **string** | Original filename provided when the upload is created. | [optional] 
**OverwriteColumnMapping** | Pointer to **bool** | Properties in &#x60;columnMapping&#x60; will be appended to the existing &#x60;columnMapping&#x60; object if set to &#x60;false&#x60;. If set to &#x60;true&#x60;, the existing &#x60;columnMapping&#x60; object will be overwritten with the data supplied in &#x60;columnMapping&#x60; property.  | [optional] 

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

### GetColumnMapping

`func (o *UploadUpdatable) GetColumnMapping() map[string]interface{}`

GetColumnMapping returns the ColumnMapping field if non-nil, zero value otherwise.

### GetColumnMappingOk

`func (o *UploadUpdatable) GetColumnMappingOk() (*map[string]interface{}, bool)`

GetColumnMappingOk returns a tuple with the ColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnMapping

`func (o *UploadUpdatable) SetColumnMapping(v map[string]interface{})`

SetColumnMapping sets ColumnMapping field to given value.

### HasColumnMapping

`func (o *UploadUpdatable) HasColumnMapping() bool`

HasColumnMapping returns a boolean if a field has been set.

### GetState

`func (o *UploadUpdatable) GetState() UploadState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UploadUpdatable) GetStateOk() (*UploadState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UploadUpdatable) SetState(v UploadState)`

SetState sets State field to given value.

### HasState

`func (o *UploadUpdatable) HasState() bool`

HasState returns a boolean if a field has been set.

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

### GetOverwriteColumnMapping

`func (o *UploadUpdatable) GetOverwriteColumnMapping() bool`

GetOverwriteColumnMapping returns the OverwriteColumnMapping field if non-nil, zero value otherwise.

### GetOverwriteColumnMappingOk

`func (o *UploadUpdatable) GetOverwriteColumnMappingOk() (*bool, bool)`

GetOverwriteColumnMappingOk returns a tuple with the OverwriteColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteColumnMapping

`func (o *UploadUpdatable) SetOverwriteColumnMapping(v bool)`

SetOverwriteColumnMapping sets OverwriteColumnMapping field to given value.

### HasOverwriteColumnMapping

`func (o *UploadUpdatable) HasOverwriteColumnMapping() bool`

HasOverwriteColumnMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


