# CardList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Card**](Card.md) | list of cards | [optional] 
**Object** | Pointer to **string** | Value is type of resource. | [optional] 
**NextUrl** | Pointer to **NullableString** | url of next page of items in list. | [optional] 
**PreviousUrl** | Pointer to **NullableString** | url of previous page of items in list. | [optional] 
**Count** | Pointer to **int32** | number of resources in a set | [optional] 
**TotalCount** | Pointer to **int32** | indicates the total number of records. Provided when the request specifies an \&quot;include\&quot; query parameter | [optional] 

## Methods

### NewCardList

`func NewCardList() *CardList`

NewCardList instantiates a new CardList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardListWithDefaults

`func NewCardListWithDefaults() *CardList`

NewCardListWithDefaults instantiates a new CardList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CardList) GetData() []Card`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CardList) GetDataOk() (*[]Card, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CardList) SetData(v []Card)`

SetData sets Data field to given value.

### HasData

`func (o *CardList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetObject

`func (o *CardList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CardList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CardList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CardList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetNextUrl

`func (o *CardList) GetNextUrl() string`

GetNextUrl returns the NextUrl field if non-nil, zero value otherwise.

### GetNextUrlOk

`func (o *CardList) GetNextUrlOk() (*string, bool)`

GetNextUrlOk returns a tuple with the NextUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUrl

`func (o *CardList) SetNextUrl(v string)`

SetNextUrl sets NextUrl field to given value.

### HasNextUrl

`func (o *CardList) HasNextUrl() bool`

HasNextUrl returns a boolean if a field has been set.

### SetNextUrlNil

`func (o *CardList) SetNextUrlNil(b bool)`

 SetNextUrlNil sets the value for NextUrl to be an explicit nil

### UnsetNextUrl
`func (o *CardList) UnsetNextUrl()`

UnsetNextUrl ensures that no value is present for NextUrl, not even an explicit nil
### GetPreviousUrl

`func (o *CardList) GetPreviousUrl() string`

GetPreviousUrl returns the PreviousUrl field if non-nil, zero value otherwise.

### GetPreviousUrlOk

`func (o *CardList) GetPreviousUrlOk() (*string, bool)`

GetPreviousUrlOk returns a tuple with the PreviousUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousUrl

`func (o *CardList) SetPreviousUrl(v string)`

SetPreviousUrl sets PreviousUrl field to given value.

### HasPreviousUrl

`func (o *CardList) HasPreviousUrl() bool`

HasPreviousUrl returns a boolean if a field has been set.

### SetPreviousUrlNil

`func (o *CardList) SetPreviousUrlNil(b bool)`

 SetPreviousUrlNil sets the value for PreviousUrl to be an explicit nil

### UnsetPreviousUrl
`func (o *CardList) UnsetPreviousUrl()`

UnsetPreviousUrl ensures that no value is present for PreviousUrl, not even an explicit nil
### GetCount

`func (o *CardList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CardList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CardList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CardList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CardList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CardList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CardList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CardList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


