# PostcardDetailsReturned

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to MAILTYPE_FIRST_CLASS]
**Size** | Pointer to [**PostcardSize**](PostcardSize.md) |  | [optional] [default to POSTCARDSIZE__4X6]
**Setting** | Pointer to **int32** |  | [optional] [default to 1001]
**FrontOriginalUrl** | Pointer to **string** | The original URL of the front template. | [optional] 
**BackOriginalUrl** | Pointer to **string** | The original URL of the back template. | [optional] 

## Methods

### NewPostcardDetailsReturned

`func NewPostcardDetailsReturned() *PostcardDetailsReturned`

NewPostcardDetailsReturned instantiates a new PostcardDetailsReturned object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostcardDetailsReturnedWithDefaults

`func NewPostcardDetailsReturnedWithDefaults() *PostcardDetailsReturned`

NewPostcardDetailsReturnedWithDefaults instantiates a new PostcardDetailsReturned object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailType

`func (o *PostcardDetailsReturned) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *PostcardDetailsReturned) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *PostcardDetailsReturned) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *PostcardDetailsReturned) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetSize

`func (o *PostcardDetailsReturned) GetSize() PostcardSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostcardDetailsReturned) GetSizeOk() (*PostcardSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostcardDetailsReturned) SetSize(v PostcardSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *PostcardDetailsReturned) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSetting

`func (o *PostcardDetailsReturned) GetSetting() int32`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *PostcardDetailsReturned) GetSettingOk() (*int32, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *PostcardDetailsReturned) SetSetting(v int32)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *PostcardDetailsReturned) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### GetFrontOriginalUrl

`func (o *PostcardDetailsReturned) GetFrontOriginalUrl() string`

GetFrontOriginalUrl returns the FrontOriginalUrl field if non-nil, zero value otherwise.

### GetFrontOriginalUrlOk

`func (o *PostcardDetailsReturned) GetFrontOriginalUrlOk() (*string, bool)`

GetFrontOriginalUrlOk returns a tuple with the FrontOriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontOriginalUrl

`func (o *PostcardDetailsReturned) SetFrontOriginalUrl(v string)`

SetFrontOriginalUrl sets FrontOriginalUrl field to given value.

### HasFrontOriginalUrl

`func (o *PostcardDetailsReturned) HasFrontOriginalUrl() bool`

HasFrontOriginalUrl returns a boolean if a field has been set.

### GetBackOriginalUrl

`func (o *PostcardDetailsReturned) GetBackOriginalUrl() string`

GetBackOriginalUrl returns the BackOriginalUrl field if non-nil, zero value otherwise.

### GetBackOriginalUrlOk

`func (o *PostcardDetailsReturned) GetBackOriginalUrlOk() (*string, bool)`

GetBackOriginalUrlOk returns a tuple with the BackOriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackOriginalUrl

`func (o *PostcardDetailsReturned) SetBackOriginalUrl(v string)`

SetBackOriginalUrl sets BackOriginalUrl field to given value.

### HasBackOriginalUrl

`func (o *PostcardDetailsReturned) HasBackOriginalUrl() bool`

HasBackOriginalUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


