# UploadUpdatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**UploadState**](UploadState.md) |  | [optional] [default to UPLOADSTATE_DRAFT]
**OriginalFilename** | Pointer to **string** | Original filename provided when the upload is created. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


