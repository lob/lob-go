# Thumbnail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Small** | Pointer to **string** | A [signed link](#section/Asset-URLs) served over HTTPS. The link returned will expire in 30 days to prevent mis-sharing. Each time a GET request is initiated, a new signed URL will be generated. | [optional] 
**Medium** | Pointer to **string** | A [signed link](#section/Asset-URLs) served over HTTPS. The link returned will expire in 30 days to prevent mis-sharing. Each time a GET request is initiated, a new signed URL will be generated. | [optional] 
**Large** | Pointer to **string** | A [signed link](#section/Asset-URLs) served over HTTPS. The link returned will expire in 30 days to prevent mis-sharing. Each time a GET request is initiated, a new signed URL will be generated. | [optional] 

## Methods

### NewThumbnail

`func NewThumbnail() *Thumbnail`

NewThumbnail instantiates a new Thumbnail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbnailWithDefaults

`func NewThumbnailWithDefaults() *Thumbnail`

NewThumbnailWithDefaults instantiates a new Thumbnail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmall

`func (o *Thumbnail) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *Thumbnail) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *Thumbnail) SetSmall(v string)`

SetSmall sets Small field to given value.

### HasSmall

`func (o *Thumbnail) HasSmall() bool`

HasSmall returns a boolean if a field has been set.

### GetMedium

`func (o *Thumbnail) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *Thumbnail) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *Thumbnail) SetMedium(v string)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *Thumbnail) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetLarge

`func (o *Thumbnail) GetLarge() string`

GetLarge returns the Large field if non-nil, zero value otherwise.

### GetLargeOk

`func (o *Thumbnail) GetLargeOk() (*string, bool)`

GetLargeOk returns a tuple with the Large field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLarge

`func (o *Thumbnail) SetLarge(v string)`

SetLarge sets Large field to given value.

### HasLarge

`func (o *Thumbnail) HasLarge() bool`

HasLarge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


