# SelfMailerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SelfMailer**](SelfMailer.md) | list of self-mailers | [optional] 
**Object** | Pointer to **string** | Value is type of resource. | [optional] 
**NextUrl** | Pointer to **NullableString** | url of next page of items in list. | [optional] 
**PreviousUrl** | Pointer to **NullableString** | url of previous page of items in list. | [optional] 
**Count** | Pointer to **int32** | number of resources in a set | [optional] 
**TotalCount** | Pointer to **int32** | indicates the total number of records. Provided when the request specifies an \&quot;include\&quot; query parameter | [optional] 

## Methods

### NewSelfMailerList

`func NewSelfMailerList() *SelfMailerList`

NewSelfMailerList instantiates a new SelfMailerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfMailerListWithDefaults

`func NewSelfMailerListWithDefaults() *SelfMailerList`

NewSelfMailerListWithDefaults instantiates a new SelfMailerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SelfMailerList) GetData() []SelfMailer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SelfMailerList) GetDataOk() (*[]SelfMailer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SelfMailerList) SetData(v []SelfMailer)`

SetData sets Data field to given value.

### HasData

`func (o *SelfMailerList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetObject

`func (o *SelfMailerList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SelfMailerList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SelfMailerList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SelfMailerList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetNextUrl

`func (o *SelfMailerList) GetNextUrl() string`

GetNextUrl returns the NextUrl field if non-nil, zero value otherwise.

### GetNextUrlOk

`func (o *SelfMailerList) GetNextUrlOk() (*string, bool)`

GetNextUrlOk returns a tuple with the NextUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUrl

`func (o *SelfMailerList) SetNextUrl(v string)`

SetNextUrl sets NextUrl field to given value.

### HasNextUrl

`func (o *SelfMailerList) HasNextUrl() bool`

HasNextUrl returns a boolean if a field has been set.

### SetNextUrlNil

`func (o *SelfMailerList) SetNextUrlNil(b bool)`

 SetNextUrlNil sets the value for NextUrl to be an explicit nil

### UnsetNextUrl
`func (o *SelfMailerList) UnsetNextUrl()`

UnsetNextUrl ensures that no value is present for NextUrl, not even an explicit nil
### GetPreviousUrl

`func (o *SelfMailerList) GetPreviousUrl() string`

GetPreviousUrl returns the PreviousUrl field if non-nil, zero value otherwise.

### GetPreviousUrlOk

`func (o *SelfMailerList) GetPreviousUrlOk() (*string, bool)`

GetPreviousUrlOk returns a tuple with the PreviousUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousUrl

`func (o *SelfMailerList) SetPreviousUrl(v string)`

SetPreviousUrl sets PreviousUrl field to given value.

### HasPreviousUrl

`func (o *SelfMailerList) HasPreviousUrl() bool`

HasPreviousUrl returns a boolean if a field has been set.

### SetPreviousUrlNil

`func (o *SelfMailerList) SetPreviousUrlNil(b bool)`

 SetPreviousUrlNil sets the value for PreviousUrl to be an explicit nil

### UnsetPreviousUrl
`func (o *SelfMailerList) UnsetPreviousUrl()`

UnsetPreviousUrl ensures that no value is present for PreviousUrl, not even an explicit nil
### GetCount

`func (o *SelfMailerList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SelfMailerList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SelfMailerList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SelfMailerList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *SelfMailerList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SelfMailerList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SelfMailerList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SelfMailerList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


