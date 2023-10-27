# CampaignWritable

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
**UseType** | [**NullableCmpUseType**](CmpUseType.md) |  | 
**AutoCancelIfNcoa** | Pointer to **bool** | Whether or not a mail piece should be automatically canceled and not sent if the address is updated via NCOA. | [optional] 

## Methods

### NewCampaignWritable

`func NewCampaignWritable(name string, scheduleType CmpScheduleType, useType NullableCmpUseType, ) *CampaignWritable`

NewCampaignWritable instantiates a new CampaignWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWritableWithDefaults

`func NewCampaignWritableWithDefaults() *CampaignWritable`

NewCampaignWritableWithDefaults instantiates a new CampaignWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingGroupId

`func (o *CampaignWritable) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *CampaignWritable) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *CampaignWritable) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *CampaignWritable) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### SetBillingGroupIdNil

`func (o *CampaignWritable) SetBillingGroupIdNil(b bool)`

 SetBillingGroupIdNil sets the value for BillingGroupId to be an explicit nil

### UnsetBillingGroupId
`func (o *CampaignWritable) UnsetBillingGroupId()`

UnsetBillingGroupId ensures that no value is present for BillingGroupId, not even an explicit nil
### GetName

`func (o *CampaignWritable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignWritable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignWritable) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignWritable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignWritable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignWritable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignWritable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CampaignWritable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CampaignWritable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduleType

`func (o *CampaignWritable) GetScheduleType() CmpScheduleType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *CampaignWritable) GetScheduleTypeOk() (*CmpScheduleType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *CampaignWritable) SetScheduleType(v CmpScheduleType)`

SetScheduleType sets ScheduleType field to given value.


### GetTargetDeliveryDate

`func (o *CampaignWritable) GetTargetDeliveryDate() time.Time`

GetTargetDeliveryDate returns the TargetDeliveryDate field if non-nil, zero value otherwise.

### GetTargetDeliveryDateOk

`func (o *CampaignWritable) GetTargetDeliveryDateOk() (*time.Time, bool)`

GetTargetDeliveryDateOk returns a tuple with the TargetDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDeliveryDate

`func (o *CampaignWritable) SetTargetDeliveryDate(v time.Time)`

SetTargetDeliveryDate sets TargetDeliveryDate field to given value.

### HasTargetDeliveryDate

`func (o *CampaignWritable) HasTargetDeliveryDate() bool`

HasTargetDeliveryDate returns a boolean if a field has been set.

### SetTargetDeliveryDateNil

`func (o *CampaignWritable) SetTargetDeliveryDateNil(b bool)`

 SetTargetDeliveryDateNil sets the value for TargetDeliveryDate to be an explicit nil

### UnsetTargetDeliveryDate
`func (o *CampaignWritable) UnsetTargetDeliveryDate()`

UnsetTargetDeliveryDate ensures that no value is present for TargetDeliveryDate, not even an explicit nil
### GetSendDate

`func (o *CampaignWritable) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *CampaignWritable) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *CampaignWritable) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *CampaignWritable) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### SetSendDateNil

`func (o *CampaignWritable) SetSendDateNil(b bool)`

 SetSendDateNil sets the value for SendDate to be an explicit nil

### UnsetSendDate
`func (o *CampaignWritable) UnsetSendDate()`

UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
### GetCancelWindowCampaignMinutes

`func (o *CampaignWritable) GetCancelWindowCampaignMinutes() int32`

GetCancelWindowCampaignMinutes returns the CancelWindowCampaignMinutes field if non-nil, zero value otherwise.

### GetCancelWindowCampaignMinutesOk

`func (o *CampaignWritable) GetCancelWindowCampaignMinutesOk() (*int32, bool)`

GetCancelWindowCampaignMinutesOk returns a tuple with the CancelWindowCampaignMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelWindowCampaignMinutes

`func (o *CampaignWritable) SetCancelWindowCampaignMinutes(v int32)`

SetCancelWindowCampaignMinutes sets CancelWindowCampaignMinutes field to given value.

### HasCancelWindowCampaignMinutes

`func (o *CampaignWritable) HasCancelWindowCampaignMinutes() bool`

HasCancelWindowCampaignMinutes returns a boolean if a field has been set.

### SetCancelWindowCampaignMinutesNil

`func (o *CampaignWritable) SetCancelWindowCampaignMinutesNil(b bool)`

 SetCancelWindowCampaignMinutesNil sets the value for CancelWindowCampaignMinutes to be an explicit nil

### UnsetCancelWindowCampaignMinutes
`func (o *CampaignWritable) UnsetCancelWindowCampaignMinutes()`

UnsetCancelWindowCampaignMinutes ensures that no value is present for CancelWindowCampaignMinutes, not even an explicit nil
### GetMetadata

`func (o *CampaignWritable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CampaignWritable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CampaignWritable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CampaignWritable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUseType

`func (o *CampaignWritable) GetUseType() CmpUseType`

GetUseType returns the UseType field if non-nil, zero value otherwise.

### GetUseTypeOk

`func (o *CampaignWritable) GetUseTypeOk() (*CmpUseType, bool)`

GetUseTypeOk returns a tuple with the UseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseType

`func (o *CampaignWritable) SetUseType(v CmpUseType)`

SetUseType sets UseType field to given value.


### SetUseTypeNil

`func (o *CampaignWritable) SetUseTypeNil(b bool)`

 SetUseTypeNil sets the value for UseType to be an explicit nil

### UnsetUseType
`func (o *CampaignWritable) UnsetUseType()`

UnsetUseType ensures that no value is present for UseType, not even an explicit nil
### GetAutoCancelIfNcoa

`func (o *CampaignWritable) GetAutoCancelIfNcoa() bool`

GetAutoCancelIfNcoa returns the AutoCancelIfNcoa field if non-nil, zero value otherwise.

### GetAutoCancelIfNcoaOk

`func (o *CampaignWritable) GetAutoCancelIfNcoaOk() (*bool, bool)`

GetAutoCancelIfNcoaOk returns a tuple with the AutoCancelIfNcoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelIfNcoa

`func (o *CampaignWritable) SetAutoCancelIfNcoa(v bool)`

SetAutoCancelIfNcoa sets AutoCancelIfNcoa field to given value.

### HasAutoCancelIfNcoa

`func (o *CampaignWritable) HasAutoCancelIfNcoa() bool`

HasAutoCancelIfNcoa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


