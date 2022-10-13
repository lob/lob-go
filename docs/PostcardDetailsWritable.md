# PostcardDetailsWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to FIRST_CLASS]
**Size** | Pointer to [**PostcardSize**](PostcardSize.md) |  | [optional] [default to _4X6]

## Methods

### NewPostcardDetailsWritable

`func NewPostcardDetailsWritable() *PostcardDetailsWritable`

NewPostcardDetailsWritable instantiates a new PostcardDetailsWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostcardDetailsWritableWithDefaults

`func NewPostcardDetailsWritableWithDefaults() *PostcardDetailsWritable`

NewPostcardDetailsWritableWithDefaults instantiates a new PostcardDetailsWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailType

`func (o *PostcardDetailsWritable) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *PostcardDetailsWritable) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *PostcardDetailsWritable) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *PostcardDetailsWritable) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetSize

`func (o *PostcardDetailsWritable) GetSize() PostcardSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostcardDetailsWritable) GetSizeOk() (*PostcardSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostcardDetailsWritable) SetSize(v PostcardSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *PostcardDetailsWritable) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


