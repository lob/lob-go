# CampaignUpdatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**ScheduleType** | Pointer to [**CmpScheduleType**](CmpScheduleType.md) |  | [optional] 
**TargetDeliveryDate** | Pointer to **time.Time** | If &#x60;schedule_type&#x60; is &#x60;target_delivery_date&#x60;, provide a targeted delivery date for mail pieces in this campaign. | [optional] 
**SendDate** | Pointer to **time.Time** | If &#x60;schedule_type&#x60; is &#x60;scheduled_send_date&#x60;, provide a date to send this campaign. | [optional] 
**CancelWindowCampaignMinutes** | Pointer to **int32** | A window, in minutes, within which the campaign can be canceled. | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**IsDraft** | Pointer to **bool** | Whether or not the campaign is still a draft. | [optional] 
**UseType** | Pointer to [**NullableCmpUseType**](CmpUseType.md) |  | [optional] 
**AutoCancelIfNcoa** | Pointer to **bool** | Whether or not a mail piece should be automatically canceled and not sent if the address is updated via NCOA. | [optional] 

## Methods

### NewCampaignUpdatable

`func NewCampaignUpdatable() *CampaignUpdatable`

NewCampaignUpdatable instantiates a new CampaignUpdatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignUpdatableWithDefaults

`func NewCampaignUpdatableWithDefaults() *CampaignUpdatable`

NewCampaignUpdatableWithDefaults instantiates a new CampaignUpdatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignUpdatable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignUpdatable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignUpdatable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignUpdatable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CampaignUpdatable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignUpdatable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignUpdatable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignUpdatable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CampaignUpdatable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CampaignUpdatable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduleType

`func (o *CampaignUpdatable) GetScheduleType() CmpScheduleType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *CampaignUpdatable) GetScheduleTypeOk() (*CmpScheduleType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *CampaignUpdatable) SetScheduleType(v CmpScheduleType)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *CampaignUpdatable) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetTargetDeliveryDate

`func (o *CampaignUpdatable) GetTargetDeliveryDate() time.Time`

GetTargetDeliveryDate returns the TargetDeliveryDate field if non-nil, zero value otherwise.

### GetTargetDeliveryDateOk

`func (o *CampaignUpdatable) GetTargetDeliveryDateOk() (*time.Time, bool)`

GetTargetDeliveryDateOk returns a tuple with the TargetDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDeliveryDate

`func (o *CampaignUpdatable) SetTargetDeliveryDate(v time.Time)`

SetTargetDeliveryDate sets TargetDeliveryDate field to given value.

### HasTargetDeliveryDate

`func (o *CampaignUpdatable) HasTargetDeliveryDate() bool`

HasTargetDeliveryDate returns a boolean if a field has been set.

### GetSendDate

`func (o *CampaignUpdatable) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *CampaignUpdatable) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *CampaignUpdatable) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *CampaignUpdatable) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### GetCancelWindowCampaignMinutes

`func (o *CampaignUpdatable) GetCancelWindowCampaignMinutes() int32`

GetCancelWindowCampaignMinutes returns the CancelWindowCampaignMinutes field if non-nil, zero value otherwise.

### GetCancelWindowCampaignMinutesOk

`func (o *CampaignUpdatable) GetCancelWindowCampaignMinutesOk() (*int32, bool)`

GetCancelWindowCampaignMinutesOk returns a tuple with the CancelWindowCampaignMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelWindowCampaignMinutes

`func (o *CampaignUpdatable) SetCancelWindowCampaignMinutes(v int32)`

SetCancelWindowCampaignMinutes sets CancelWindowCampaignMinutes field to given value.

### HasCancelWindowCampaignMinutes

`func (o *CampaignUpdatable) HasCancelWindowCampaignMinutes() bool`

HasCancelWindowCampaignMinutes returns a boolean if a field has been set.

### GetMetadata

`func (o *CampaignUpdatable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CampaignUpdatable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CampaignUpdatable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CampaignUpdatable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIsDraft

`func (o *CampaignUpdatable) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *CampaignUpdatable) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *CampaignUpdatable) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *CampaignUpdatable) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetUseType

`func (o *CampaignUpdatable) GetUseType() CmpUseType`

GetUseType returns the UseType field if non-nil, zero value otherwise.

### GetUseTypeOk

`func (o *CampaignUpdatable) GetUseTypeOk() (*CmpUseType, bool)`

GetUseTypeOk returns a tuple with the UseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseType

`func (o *CampaignUpdatable) SetUseType(v CmpUseType)`

SetUseType sets UseType field to given value.

### HasUseType

`func (o *CampaignUpdatable) HasUseType() bool`

HasUseType returns a boolean if a field has been set.

### SetUseTypeNil

`func (o *CampaignUpdatable) SetUseTypeNil(b bool)`

 SetUseTypeNil sets the value for UseType to be an explicit nil

### UnsetUseType
`func (o *CampaignUpdatable) UnsetUseType()`

UnsetUseType ensures that no value is present for UseType, not even an explicit nil
### GetAutoCancelIfNcoa

`func (o *CampaignUpdatable) GetAutoCancelIfNcoa() bool`

GetAutoCancelIfNcoa returns the AutoCancelIfNcoa field if non-nil, zero value otherwise.

### GetAutoCancelIfNcoaOk

`func (o *CampaignUpdatable) GetAutoCancelIfNcoaOk() (*bool, bool)`

GetAutoCancelIfNcoaOk returns a tuple with the AutoCancelIfNcoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelIfNcoa

`func (o *CampaignUpdatable) SetAutoCancelIfNcoa(v bool)`

SetAutoCancelIfNcoa sets AutoCancelIfNcoa field to given value.

### HasAutoCancelIfNcoa

`func (o *CampaignUpdatable) HasAutoCancelIfNcoa() bool`

HasAutoCancelIfNcoa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


