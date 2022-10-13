# CampaignsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Campaign**](Campaign.md) | list of campaigns | [optional] 
**Object** | Pointer to **string** | Value is type of resource. | [optional] 
**NextUrl** | Pointer to **NullableString** | url of next page of items in list. | [optional] 
**PreviousUrl** | Pointer to **NullableString** | url of previous page of items in list. | [optional] 
**Count** | Pointer to **int32** | number of resources in a set | [optional] 
**TotalCount** | Pointer to **int32** | indicates the total number of records. Provided when the request specifies an \&quot;include\&quot; query parameter | [optional] 

## Methods

### NewCampaignsList

`func NewCampaignsList() *CampaignsList`

NewCampaignsList instantiates a new CampaignsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignsListWithDefaults

`func NewCampaignsListWithDefaults() *CampaignsList`

NewCampaignsListWithDefaults instantiates a new CampaignsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CampaignsList) GetData() []Campaign`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CampaignsList) GetDataOk() (*[]Campaign, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CampaignsList) SetData(v []Campaign)`

SetData sets Data field to given value.

### HasData

`func (o *CampaignsList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetObject

`func (o *CampaignsList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CampaignsList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CampaignsList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CampaignsList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetNextUrl

`func (o *CampaignsList) GetNextUrl() string`

GetNextUrl returns the NextUrl field if non-nil, zero value otherwise.

### GetNextUrlOk

`func (o *CampaignsList) GetNextUrlOk() (*string, bool)`

GetNextUrlOk returns a tuple with the NextUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUrl

`func (o *CampaignsList) SetNextUrl(v string)`

SetNextUrl sets NextUrl field to given value.

### HasNextUrl

`func (o *CampaignsList) HasNextUrl() bool`

HasNextUrl returns a boolean if a field has been set.

### SetNextUrlNil

`func (o *CampaignsList) SetNextUrlNil(b bool)`

 SetNextUrlNil sets the value for NextUrl to be an explicit nil

### UnsetNextUrl
`func (o *CampaignsList) UnsetNextUrl()`

UnsetNextUrl ensures that no value is present for NextUrl, not even an explicit nil
### GetPreviousUrl

`func (o *CampaignsList) GetPreviousUrl() string`

GetPreviousUrl returns the PreviousUrl field if non-nil, zero value otherwise.

### GetPreviousUrlOk

`func (o *CampaignsList) GetPreviousUrlOk() (*string, bool)`

GetPreviousUrlOk returns a tuple with the PreviousUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousUrl

`func (o *CampaignsList) SetPreviousUrl(v string)`

SetPreviousUrl sets PreviousUrl field to given value.

### HasPreviousUrl

`func (o *CampaignsList) HasPreviousUrl() bool`

HasPreviousUrl returns a boolean if a field has been set.

### SetPreviousUrlNil

`func (o *CampaignsList) SetPreviousUrlNil(b bool)`

 SetPreviousUrlNil sets the value for PreviousUrl to be an explicit nil

### UnsetPreviousUrl
`func (o *CampaignsList) UnsetPreviousUrl()`

UnsetPreviousUrl ensures that no value is present for PreviousUrl, not even an explicit nil
### GetCount

`func (o *CampaignsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CampaignsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CampaignsList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CampaignsList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CampaignsList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CampaignsList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CampaignsList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CampaignsList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


