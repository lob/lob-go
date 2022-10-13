# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingGroupId** | Pointer to **NullableString** | Unique identifier prefixed with &#x60;bg_&#x60;. | [optional] 
**Name** | **string** | Name of the campaign. | 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**ScheduleType** | [**CmpScheduleType**](CmpScheduleType.md) |  | 
**TargetDeliveryDate** | Pointer to **NullableTime** | If &#x60;schedule_type&#x60; is &#x60;target_delivery_date&#x60;, provide a targeted delivery date for mail pieces in this campaign. | [optional] 
**SendDate** | Pointer to **NullableTime** | If &#x60;schedule_type&#x60; is &#x60;scheduled_send_date&#x60;, provide a date to send this campaign. | [optional] 
**CancelWindowCampaignMinutes** | Pointer to **NullableInt32** | A window, in minutes, within which the campaign can be canceled. | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**UseType** | Pointer to [**NullableCmpUseType**](CmpUseType.md) |  | [optional] 
**AutoCancelIfNcoa** | **bool** | Whether or not a mail piece should be automatically canceled and not sent if the address is updated via NCOA. | 
**Id** | **string** | Unique identifier prefixed with &#x60;cmp_&#x60;. | 
**AccountId** | Pointer to **string** | Account ID that this campaign is associated with. | [optional] 
**IsDraft** | **bool** | Whether or not the campaign is still a draft. | [default to true]
**Creatives** | [**[]CampaignCreative**](CampaignCreative.md) | An array of creatives that have been associated with this campaign. | 
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Object** | **string** | Value is resource type. | [default to "campaign"]

## Methods

### NewCampaign

`func NewCampaign(name string, scheduleType CmpScheduleType, autoCancelIfNcoa bool, id string, isDraft bool, creatives []CampaignCreative, dateCreated time.Time, dateModified time.Time, object string, ) *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingGroupId

`func (o *Campaign) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *Campaign) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *Campaign) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *Campaign) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### SetBillingGroupIdNil

`func (o *Campaign) SetBillingGroupIdNil(b bool)`

 SetBillingGroupIdNil sets the value for BillingGroupId to be an explicit nil

### UnsetBillingGroupId
`func (o *Campaign) UnsetBillingGroupId()`

UnsetBillingGroupId ensures that no value is present for BillingGroupId, not even an explicit nil
### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Campaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Campaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Campaign) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Campaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Campaign) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Campaign) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduleType

`func (o *Campaign) GetScheduleType() CmpScheduleType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *Campaign) GetScheduleTypeOk() (*CmpScheduleType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *Campaign) SetScheduleType(v CmpScheduleType)`

SetScheduleType sets ScheduleType field to given value.


### GetTargetDeliveryDate

`func (o *Campaign) GetTargetDeliveryDate() time.Time`

GetTargetDeliveryDate returns the TargetDeliveryDate field if non-nil, zero value otherwise.

### GetTargetDeliveryDateOk

`func (o *Campaign) GetTargetDeliveryDateOk() (*time.Time, bool)`

GetTargetDeliveryDateOk returns a tuple with the TargetDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDeliveryDate

`func (o *Campaign) SetTargetDeliveryDate(v time.Time)`

SetTargetDeliveryDate sets TargetDeliveryDate field to given value.

### HasTargetDeliveryDate

`func (o *Campaign) HasTargetDeliveryDate() bool`

HasTargetDeliveryDate returns a boolean if a field has been set.

### SetTargetDeliveryDateNil

`func (o *Campaign) SetTargetDeliveryDateNil(b bool)`

 SetTargetDeliveryDateNil sets the value for TargetDeliveryDate to be an explicit nil

### UnsetTargetDeliveryDate
`func (o *Campaign) UnsetTargetDeliveryDate()`

UnsetTargetDeliveryDate ensures that no value is present for TargetDeliveryDate, not even an explicit nil
### GetSendDate

`func (o *Campaign) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *Campaign) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *Campaign) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *Campaign) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### SetSendDateNil

`func (o *Campaign) SetSendDateNil(b bool)`

 SetSendDateNil sets the value for SendDate to be an explicit nil

### UnsetSendDate
`func (o *Campaign) UnsetSendDate()`

UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
### GetCancelWindowCampaignMinutes

`func (o *Campaign) GetCancelWindowCampaignMinutes() int32`

GetCancelWindowCampaignMinutes returns the CancelWindowCampaignMinutes field if non-nil, zero value otherwise.

### GetCancelWindowCampaignMinutesOk

`func (o *Campaign) GetCancelWindowCampaignMinutesOk() (*int32, bool)`

GetCancelWindowCampaignMinutesOk returns a tuple with the CancelWindowCampaignMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelWindowCampaignMinutes

`func (o *Campaign) SetCancelWindowCampaignMinutes(v int32)`

SetCancelWindowCampaignMinutes sets CancelWindowCampaignMinutes field to given value.

### HasCancelWindowCampaignMinutes

`func (o *Campaign) HasCancelWindowCampaignMinutes() bool`

HasCancelWindowCampaignMinutes returns a boolean if a field has been set.

### SetCancelWindowCampaignMinutesNil

`func (o *Campaign) SetCancelWindowCampaignMinutesNil(b bool)`

 SetCancelWindowCampaignMinutesNil sets the value for CancelWindowCampaignMinutes to be an explicit nil

### UnsetCancelWindowCampaignMinutes
`func (o *Campaign) UnsetCancelWindowCampaignMinutes()`

UnsetCancelWindowCampaignMinutes ensures that no value is present for CancelWindowCampaignMinutes, not even an explicit nil
### GetMetadata

`func (o *Campaign) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Campaign) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Campaign) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Campaign) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUseType

`func (o *Campaign) GetUseType() CmpUseType`

GetUseType returns the UseType field if non-nil, zero value otherwise.

### GetUseTypeOk

`func (o *Campaign) GetUseTypeOk() (*CmpUseType, bool)`

GetUseTypeOk returns a tuple with the UseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseType

`func (o *Campaign) SetUseType(v CmpUseType)`

SetUseType sets UseType field to given value.

### HasUseType

`func (o *Campaign) HasUseType() bool`

HasUseType returns a boolean if a field has been set.

### SetUseTypeNil

`func (o *Campaign) SetUseTypeNil(b bool)`

 SetUseTypeNil sets the value for UseType to be an explicit nil

### UnsetUseType
`func (o *Campaign) UnsetUseType()`

UnsetUseType ensures that no value is present for UseType, not even an explicit nil
### GetAutoCancelIfNcoa

`func (o *Campaign) GetAutoCancelIfNcoa() bool`

GetAutoCancelIfNcoa returns the AutoCancelIfNcoa field if non-nil, zero value otherwise.

### GetAutoCancelIfNcoaOk

`func (o *Campaign) GetAutoCancelIfNcoaOk() (*bool, bool)`

GetAutoCancelIfNcoaOk returns a tuple with the AutoCancelIfNcoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelIfNcoa

`func (o *Campaign) SetAutoCancelIfNcoa(v bool)`

SetAutoCancelIfNcoa sets AutoCancelIfNcoa field to given value.


### GetId

`func (o *Campaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Campaign) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Campaign) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Campaign) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Campaign) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetIsDraft

`func (o *Campaign) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *Campaign) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *Campaign) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.


### GetCreatives

`func (o *Campaign) GetCreatives() []CampaignCreative`

GetCreatives returns the Creatives field if non-nil, zero value otherwise.

### GetCreativesOk

`func (o *Campaign) GetCreativesOk() (*[]CampaignCreative, bool)`

GetCreativesOk returns a tuple with the Creatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatives

`func (o *Campaign) SetCreatives(v []CampaignCreative)`

SetCreatives sets Creatives field to given value.


### GetDateCreated

`func (o *Campaign) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Campaign) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Campaign) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *Campaign) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *Campaign) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *Campaign) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetDeleted

`func (o *Campaign) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Campaign) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Campaign) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Campaign) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetObject

`func (o *Campaign) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Campaign) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Campaign) SetObject(v string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


