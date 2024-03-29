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

// CampaignWritable struct for CampaignWritable
type CampaignWritable struct {
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
	UseType NullableCmpUseType `json:"use_type"`
	// Whether or not a mail piece should be automatically canceled and not sent if the address is updated via NCOA.
	AutoCancelIfNcoa *bool `json:"auto_cancel_if_ncoa,omitempty"`
}

// NewCampaignWritable instantiates a new CampaignWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignWritable(name string, scheduleType CmpScheduleType, useType NullableCmpUseType) *CampaignWritable {
	this := CampaignWritable{}
	this.Name = name
	this.ScheduleType = scheduleType
	this.UseType = useType
	return &this
}

// NewCampaignWritableWithDefaults instantiates a new CampaignWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignWritableWithDefaults() *CampaignWritable {
	this := CampaignWritable{}
	return &this
}

// GetBillingGroupId returns the BillingGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignWritable) GetBillingGroupId() string {
	if o == nil || o.BillingGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingGroupId.Get()
}

// GetBillingGroupIdOk returns a tuple with the BillingGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignWritable) GetBillingGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingGroupId.Get(), o.BillingGroupId.IsSet()
}

// HasBillingGroupId returns a boolean if a field has been set.
func (o *CampaignWritable) HasBillingGroupId() bool {
	if o != nil && o.BillingGroupId.IsSet() {
		return true
	}

	return false
}

// SetBillingGroupId gets a reference to the given NullableString and assigns it to the BillingGroupId field.
func (o *CampaignWritable) SetBillingGroupId(v string) {
	o.BillingGroupId.Set(&v)
}
// SetBillingGroupIdNil sets the value for BillingGroupId to be an explicit nil
func (o *CampaignWritable) SetBillingGroupIdNil() {
	o.BillingGroupId.Set(nil)
}

// UnsetBillingGroupId ensures that no value is present for BillingGroupId, not even an explicit nil
func (o *CampaignWritable) UnsetBillingGroupId() {
	o.BillingGroupId.Unset()
}

// GetName returns the Name field value
func (o *CampaignWritable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignWritable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignWritable) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignWritable) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignWritable) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignWritable) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CampaignWritable) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CampaignWritable) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CampaignWritable) UnsetDescription() {
	o.Description.Unset()
}

// GetScheduleType returns the ScheduleType field value
func (o *CampaignWritable) GetScheduleType() CmpScheduleType {
	if o == nil {
		var ret CmpScheduleType
		return ret
	}

	return o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value
// and a boolean to check if the value has been set.
func (o *CampaignWritable) GetScheduleTypeOk() (*CmpScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleType, true
}

// SetScheduleType sets field value
func (o *CampaignWritable) SetScheduleType(v CmpScheduleType) {
	o.ScheduleType = v
}

// GetTargetDeliveryDate returns the TargetDeliveryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignWritable) GetTargetDeliveryDate() time.Time {
	if o == nil || o.TargetDeliveryDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.TargetDeliveryDate.Get()
}

// GetTargetDeliveryDateOk returns a tuple with the TargetDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignWritable) GetTargetDeliveryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetDeliveryDate.Get(), o.TargetDeliveryDate.IsSet()
}

// HasTargetDeliveryDate returns a boolean if a field has been set.
func (o *CampaignWritable) HasTargetDeliveryDate() bool {
	if o != nil && o.TargetDeliveryDate.IsSet() {
		return true
	}

	return false
}

// SetTargetDeliveryDate gets a reference to the given NullableTime and assigns it to the TargetDeliveryDate field.
func (o *CampaignWritable) SetTargetDeliveryDate(v time.Time) {
	o.TargetDeliveryDate.Set(&v)
}
// SetTargetDeliveryDateNil sets the value for TargetDeliveryDate to be an explicit nil
func (o *CampaignWritable) SetTargetDeliveryDateNil() {
	o.TargetDeliveryDate.Set(nil)
}

// UnsetTargetDeliveryDate ensures that no value is present for TargetDeliveryDate, not even an explicit nil
func (o *CampaignWritable) UnsetTargetDeliveryDate() {
	o.TargetDeliveryDate.Unset()
}

// GetSendDate returns the SendDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignWritable) GetSendDate() time.Time {
	if o == nil || o.SendDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SendDate.Get()
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignWritable) GetSendDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendDate.Get(), o.SendDate.IsSet()
}

// HasSendDate returns a boolean if a field has been set.
func (o *CampaignWritable) HasSendDate() bool {
	if o != nil && o.SendDate.IsSet() {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given NullableTime and assigns it to the SendDate field.
func (o *CampaignWritable) SetSendDate(v time.Time) {
	o.SendDate.Set(&v)
}
// SetSendDateNil sets the value for SendDate to be an explicit nil
func (o *CampaignWritable) SetSendDateNil() {
	o.SendDate.Set(nil)
}

// UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
func (o *CampaignWritable) UnsetSendDate() {
	o.SendDate.Unset()
}

// GetCancelWindowCampaignMinutes returns the CancelWindowCampaignMinutes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignWritable) GetCancelWindowCampaignMinutes() int32 {
	if o == nil || o.CancelWindowCampaignMinutes.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CancelWindowCampaignMinutes.Get()
}

// GetCancelWindowCampaignMinutesOk returns a tuple with the CancelWindowCampaignMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignWritable) GetCancelWindowCampaignMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelWindowCampaignMinutes.Get(), o.CancelWindowCampaignMinutes.IsSet()
}

// HasCancelWindowCampaignMinutes returns a boolean if a field has been set.
func (o *CampaignWritable) HasCancelWindowCampaignMinutes() bool {
	if o != nil && o.CancelWindowCampaignMinutes.IsSet() {
		return true
	}

	return false
}

// SetCancelWindowCampaignMinutes gets a reference to the given NullableInt32 and assigns it to the CancelWindowCampaignMinutes field.
func (o *CampaignWritable) SetCancelWindowCampaignMinutes(v int32) {
	o.CancelWindowCampaignMinutes.Set(&v)
}
// SetCancelWindowCampaignMinutesNil sets the value for CancelWindowCampaignMinutes to be an explicit nil
func (o *CampaignWritable) SetCancelWindowCampaignMinutesNil() {
	o.CancelWindowCampaignMinutes.Set(nil)
}

// UnsetCancelWindowCampaignMinutes ensures that no value is present for CancelWindowCampaignMinutes, not even an explicit nil
func (o *CampaignWritable) UnsetCancelWindowCampaignMinutes() {
	o.CancelWindowCampaignMinutes.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CampaignWritable) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignWritable) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CampaignWritable) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CampaignWritable) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetUseType returns the UseType field value
// If the value is explicit nil, the zero value for CmpUseType will be returned
func (o *CampaignWritable) GetUseType() CmpUseType {
	if o == nil || o.UseType.Get() == nil {
		var ret CmpUseType
		return ret
	}

	return *o.UseType.Get()
}

// GetUseTypeOk returns a tuple with the UseType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignWritable) GetUseTypeOk() (*CmpUseType, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseType.Get(), o.UseType.IsSet()
}

// SetUseType sets field value
func (o *CampaignWritable) SetUseType(v CmpUseType) {
	o.UseType.Set(&v)
}

// GetAutoCancelIfNcoa returns the AutoCancelIfNcoa field value if set, zero value otherwise.
func (o *CampaignWritable) GetAutoCancelIfNcoa() bool {
	if o == nil || o.AutoCancelIfNcoa == nil {
		var ret bool
		return ret
	}
	return *o.AutoCancelIfNcoa
}

// GetAutoCancelIfNcoaOk returns a tuple with the AutoCancelIfNcoa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignWritable) GetAutoCancelIfNcoaOk() (*bool, bool) {
	if o == nil || o.AutoCancelIfNcoa == nil {
		return nil, false
	}
	return o.AutoCancelIfNcoa, true
}

// HasAutoCancelIfNcoa returns a boolean if a field has been set.
func (o *CampaignWritable) HasAutoCancelIfNcoa() bool {
	if o != nil && o.AutoCancelIfNcoa != nil {
		return true
	}

	return false
}

// SetAutoCancelIfNcoa gets a reference to the given bool and assigns it to the AutoCancelIfNcoa field.
func (o *CampaignWritable) SetAutoCancelIfNcoa(v bool) {
	o.AutoCancelIfNcoa = &v
}

func (o CampaignWritable) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["use_type"] = o.UseType.Get()
	}
	if o.AutoCancelIfNcoa != nil {
		toSerialize["auto_cancel_if_ncoa"] = o.AutoCancelIfNcoa
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignWritable struct {
	value *CampaignWritable
	isSet bool
}

func (v NullableCampaignWritable) Get() *CampaignWritable {
	return v.value
}

func (v *NullableCampaignWritable) Set(val *CampaignWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignWritable(val *CampaignWritable) *NullableCampaignWritable {
	return &NullableCampaignWritable{value: val, isSet: true}
}

func (v NullableCampaignWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


