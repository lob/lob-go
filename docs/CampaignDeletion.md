# CampaignDeletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;cmp_&#x60;. | [optional] 
**Deleted** | Pointer to **bool** | True if the resource has been successfully deleted. | [optional] 

## Methods

### NewCampaignDeletion

`func NewCampaignDeletion() *CampaignDeletion`

NewCampaignDeletion instantiates a new CampaignDeletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDeletionWithDefaults

`func NewCampaignDeletionWithDefaults() *CampaignDeletion`

NewCampaignDeletionWithDefaults instantiates a new CampaignDeletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignDeletion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignDeletion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignDeletion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignDeletion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *CampaignDeletion) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CampaignDeletion) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CampaignDeletion) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CampaignDeletion) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


