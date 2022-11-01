/*
Lob

The Lob API is organized around REST. Our API is designed to have predictable, resource-oriented URLs and uses HTTP response codes to indicate any API errors. <p> Looking for our [previous documentation](https://lob.github.io/legacy-docs/)? 

API version: 1.3.0
Contact: lob-openapi@lob.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lob

import (
	"encoding/json"
	
	"time"
)

// Campaign struct for Campaign
type Campaign struct {
	// Unique identifier prefixed with `bg_`.
	BillingGroupId NullableString `json:"billing_group_id,omitempty"`
	// Name of the campaign.
	Name string `json:"name"`
	// An internal description that identifies this resource. Must be no longer than 255 characters. 
	Description NullableString `json:"description,omitempty"`
	ScheduleType CmpScheduleType `json:"schedule_type"`
	// If `schedule_type` is `target_delivery_date`, provide a targeted delivery date for mail pieces in this campaign.
	TargetDeliveryDate NullableTime `json:"target_delivery_date,omitempty"`
	// If `schedule_type` is `scheduled_send_date`, provide a date to send this campaign.
	SendDate NullableTime `json:"send_date,omitempty"`
	// A window, in minutes, within which the campaign can be canceled.
	CancelWindowCampaignMinutes NullableInt32 `json:"cancel_window_campaign_minutes,omitempty"`
	// Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters `\"` and `\\`. i.e. '{\"customer_id\" : \"NEWYORK2015\"}' Nested objects are not supported.  See [Metadata](#section/Metadata) for more information.
	Metadata *map[string]string `json:"metadata,omitempty"`
	UseType NullableCmpUseType `json:"use_type,omitempty"`
	// Whether or not a mail piece should be automatically canceled and not sent if the address is updated via NCOA.
	AutoCancelIfNcoa bool `json:"auto_cancel_if_ncoa"`
	// Unique identifier prefixed with `cmp_`.
	Id string `json:"id"`
	// Account ID that this campaign is associated with.
	AccountId *string `json:"account_id,omitempty"`
	// Whether or not the campaign is still a draft.
	IsDraft bool `json:"is_draft"`
	// An array of creatives that have been associated with this campaign.
	Creatives []CampaignCreative `json:"creatives"`
	// A timestamp in ISO 8601 format of the date the resource was created.
	DateCreated time.Time `json:"date_created"`
	// A timestamp in ISO 8601 format of the date the resource was last modified.
	DateModified time.Time `json:"date_modified"`
	// Only returned if the resource has been successfully deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// Value is resource type.
	Object string `json:"object"`
}

// NewCampaign instantiates a new Campaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaign(name string, scheduleType CmpScheduleType, autoCancelIfNcoa bool, id string, isDraft bool, creatives []CampaignCreative, dateCreated time.Time, dateModified time.Time, object string) *Campaign {
	this := Campaign{}
	this.Name = name
	this.ScheduleType = scheduleType
	this.AutoCancelIfNcoa = autoCancelIfNcoa
	this.Id = id
	this.IsDraft = isDraft
	this.Creatives = creatives
	this.DateCreated = dateCreated
	this.DateModified = dateModified
	this.Object = object
	return &this
}

// NewCampaignWithDefaults instantiates a new Campaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignWithDefaults() *Campaign {
	this := Campaign{}
	var isDraft bool = true
	this.IsDraft = isDraft
	var object string = "campaign"
	this.Object = object
	return &this
}

// GetBillingGroupId returns the BillingGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Campaign) GetBillingGroupId() string {
	if o == nil || o.BillingGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingGroupId.Get()
}

// GetBillingGroupIdOk returns a tuple with the BillingGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Campaign) GetBillingGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingGroupId.Get(), o.BillingGroupId.IsSet()
}

// HasBillingGroupId returns a boolean if a field has been set.
func (o *Campaign) HasBillingGroupId() bool {
	if o != nil && o.BillingGroupId.IsSet() {
		return true
	}

	return false
}

// SetBillingGroupId gets a reference to the given NullableString and assigns it to the BillingGroupId field.
func (o *Campaign) SetBillingGroupId(v string) {
	o.BillingGroupId.Set(&v)
}
// SetBillingGroupIdNil sets the value for BillingGroupId to be an explicit nil
func (o *Campaign) SetBillingGroupIdNil() {
	o.BillingGroupId.Set(nil)
}

// UnsetBillingGroupId ensures that no value is present for BillingGroupId, not even an explicit nil
func (o *Campaign) UnsetBillingGroupId() {
	o.BillingGroupId.Unset()
}

// GetName returns the Name field value
func (o *Campaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Campaign) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Campaign) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Campaign) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Campaign) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Campaign) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Campaign) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Campaign) UnsetDescription() {
	o.Description.Unset()
}

// GetScheduleType returns the ScheduleType field value
func (o *Campaign) GetScheduleType() CmpScheduleType {
	if o == nil {
		var ret CmpScheduleType
		return ret
	}

	return o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetScheduleTypeOk() (*CmpScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleType, true
}

// SetScheduleType sets field value
func (o *Campaign) SetScheduleType(v CmpScheduleType) {
	o.ScheduleType = v
}

// GetTargetDeliveryDate returns the TargetDeliveryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Campaign) GetTargetDeliveryDate() time.Time {
	if o == nil || o.TargetDeliveryDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.TargetDeliveryDate.Get()
}

// GetTargetDeliveryDateOk returns a tuple with the TargetDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Campaign) GetTargetDeliveryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetDeliveryDate.Get(), o.TargetDeliveryDate.IsSet()
}

// HasTargetDeliveryDate returns a boolean if a field has been set.
func (o *Campaign) HasTargetDeliveryDate() bool {
	if o != nil && o.TargetDeliveryDate.IsSet() {
		return true
	}

	return false
}

// SetTargetDeliveryDate gets a reference to the given NullableTime and assigns it to the TargetDeliveryDate field.
func (o *Campaign) SetTargetDeliveryDate(v time.Time) {
	o.TargetDeliveryDate.Set(&v)
}
// SetTargetDeliveryDateNil sets the value for TargetDeliveryDate to be an explicit nil
func (o *Campaign) SetTargetDeliveryDateNil() {
	o.TargetDeliveryDate.Set(nil)
}

// UnsetTargetDeliveryDate ensures that no value is present for TargetDeliveryDate, not even an explicit nil
func (o *Campaign) UnsetTargetDeliveryDate() {
	o.TargetDeliveryDate.Unset()
}

// GetSendDate returns the SendDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Campaign) GetSendDate() time.Time {
	if o == nil || o.SendDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SendDate.Get()
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Campaign) GetSendDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendDate.Get(), o.SendDate.IsSet()
}

// HasSendDate returns a boolean if a field has been set.
func (o *Campaign) HasSendDate() bool {
	if o != nil && o.SendDate.IsSet() {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given NullableTime and assigns it to the SendDate field.
func (o *Campaign) SetSendDate(v time.Time) {
	o.SendDate.Set(&v)
}
// SetSendDateNil sets the value for SendDate to be an explicit nil
func (o *Campaign) SetSendDateNil() {
	o.SendDate.Set(nil)
}

// UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
func (o *Campaign) UnsetSendDate() {
	o.SendDate.Unset()
}

// GetCancelWindowCampaignMinutes returns the CancelWindowCampaignMinutes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Campaign) GetCancelWindowCampaignMinutes() int32 {
	if o == nil || o.CancelWindowCampaignMinutes.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CancelWindowCampaignMinutes.Get()
}

// GetCancelWindowCampaignMinutesOk returns a tuple with the CancelWindowCampaignMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Campaign) GetCancelWindowCampaignMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelWindowCampaignMinutes.Get(), o.CancelWindowCampaignMinutes.IsSet()
}

// HasCancelWindowCampaignMinutes returns a boolean if a field has been set.
func (o *Campaign) HasCancelWindowCampaignMinutes() bool {
	if o != nil && o.CancelWindowCampaignMinutes.IsSet() {
		return true
	}

	return false
}

// SetCancelWindowCampaignMinutes gets a reference to the given NullableInt32 and assigns it to the CancelWindowCampaignMinutes field.
func (o *Campaign) SetCancelWindowCampaignMinutes(v int32) {
	o.CancelWindowCampaignMinutes.Set(&v)
}
// SetCancelWindowCampaignMinutesNil sets the value for CancelWindowCampaignMinutes to be an explicit nil
func (o *Campaign) SetCancelWindowCampaignMinutesNil() {
	o.CancelWindowCampaignMinutes.Set(nil)
}

// UnsetCancelWindowCampaignMinutes ensures that no value is present for CancelWindowCampaignMinutes, not even an explicit nil
func (o *Campaign) UnsetCancelWindowCampaignMinutes() {
	o.CancelWindowCampaignMinutes.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Campaign) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Campaign) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Campaign) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetUseType returns the UseType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Campaign) GetUseType() CmpUseType {
	if o == nil || o.UseType.Get() == nil {
		var ret CmpUseType
		return ret
	}
	return *o.UseType.Get()
}

// GetUseTypeOk returns a tuple with the UseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Campaign) GetUseTypeOk() (*CmpUseType, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseType.Get(), o.UseType.IsSet()
}

// HasUseType returns a boolean if a field has been set.
func (o *Campaign) HasUseType() bool {
	if o != nil && o.UseType.IsSet() {
		return true
	}

	return false
}

// SetUseType gets a reference to the given NullableCmpUseType and assigns it to the UseType field.
func (o *Campaign) SetUseType(v CmpUseType) {
	o.UseType.Set(&v)
}
// SetUseTypeNil sets the value for UseType to be an explicit nil
func (o *Campaign) SetUseTypeNil() {
	o.UseType.Set(nil)
}

// UnsetUseType ensures that no value is present for UseType, not even an explicit nil
func (o *Campaign) UnsetUseType() {
	o.UseType.Unset()
}

// GetAutoCancelIfNcoa returns the AutoCancelIfNcoa field value
func (o *Campaign) GetAutoCancelIfNcoa() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCancelIfNcoa
}

// GetAutoCancelIfNcoaOk returns a tuple with the AutoCancelIfNcoa field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetAutoCancelIfNcoaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCancelIfNcoa, true
}

// SetAutoCancelIfNcoa sets field value
func (o *Campaign) SetAutoCancelIfNcoa(v bool) {
	o.AutoCancelIfNcoa = v
}

// GetId returns the Id field value
func (o *Campaign) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Campaign) SetId(v string) {
	o.Id = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Campaign) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Campaign) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Campaign) SetAccountId(v string) {
	o.AccountId = &v
}

// GetIsDraft returns the IsDraft field value
func (o *Campaign) GetIsDraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDraft
}

// GetIsDraftOk returns a tuple with the IsDraft field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetIsDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDraft, true
}

// SetIsDraft sets field value
func (o *Campaign) SetIsDraft(v bool) {
	o.IsDraft = v
}

// GetCreatives returns the Creatives field value
func (o *Campaign) GetCreatives() []CampaignCreative {
	if o == nil {
		var ret []CampaignCreative
		return ret
	}

	return o.Creatives
}

// GetCreativesOk returns a tuple with the Creatives field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetCreativesOk() ([]CampaignCreative, bool) {
	if o == nil {
		return nil, false
	}
	return o.Creatives, true
}

// SetCreatives sets field value
func (o *Campaign) SetCreatives(v []CampaignCreative) {
	o.Creatives = v
}

// GetDateCreated returns the DateCreated field value
func (o *Campaign) GetDateCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *Campaign) SetDateCreated(v time.Time) {
	o.DateCreated = v
}

// GetDateModified returns the DateModified field value
func (o *Campaign) GetDateModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetDateModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateModified, true
}

// SetDateModified sets field value
func (o *Campaign) SetDateModified(v time.Time) {
	o.DateModified = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Campaign) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Campaign) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Campaign) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetObject returns the Object field value
func (o *Campaign) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Campaign) SetObject(v string) {
	o.Object = v
}

func (o Campaign) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingGroupId.IsSet() {
		toSerialize["billing_group_id"] = o.BillingGroupId.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["schedule_type"] = o.ScheduleType
	}
	if o.TargetDeliveryDate.IsSet() {
		toSerialize["target_delivery_date"] = o.TargetDeliveryDate.Get()
	}
	if o.SendDate.IsSet() {
		toSerialize["send_date"] = o.SendDate.Get()
	}
	if o.CancelWindowCampaignMinutes.IsSet() {
		toSerialize["cancel_window_campaign_minutes"] = o.CancelWindowCampaignMinutes.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.UseType.IsSet() {
		toSerialize["use_type"] = o.UseType.Get()
	}
	if true {
		toSerialize["auto_cancel_if_ncoa"] = o.AutoCancelIfNcoa
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["is_draft"] = o.IsDraft
	}
	if true {
		toSerialize["creatives"] = o.Creatives
	}
	if true {
		toSerialize["date_created"] = o.DateCreated
	}
	if true {
		toSerialize["date_modified"] = o.DateModified
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if true {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableCampaign struct {
	value *Campaign
	isSet bool
}

func (v NullableCampaign) Get() *Campaign {
	return v.value
}

func (v *NullableCampaign) Set(val *Campaign) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaign(val *Campaign) *NullableCampaign {
	return &NullableCampaign{value: val, isSet: true}
}

func (v NullableCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


