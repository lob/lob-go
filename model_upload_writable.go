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
	
)

// UploadWritable struct for UploadWritable
type UploadWritable struct {
	CampaignId string `json:"campaignId"`
	RequiredAddressColumnMapping *RequiredAddressColumnMapping `json:"requiredAddressColumnMapping,omitempty"`
	OptionalAddressColumnMapping *OptionalAddressColumnMapping `json:"optionalAddressColumnMapping,omitempty"`
	Metadata *UploadsMetadata `json:"metadata,omitempty"`
	// The mapping of column headers in your file to the merge variables present in your creative. See our <a href=\"https://help.lob.com/print-and-mail/building-a-mail-strategy/campaign-or-triggered-sends/campaign-audience-guide#step-3-map-merge-variable-data-if-applicable-7\" target=\"_blank\">Campaign Audience Guide</a> for additional details. <br />If a merge variable has the same \"name\" as a \"key\" in the `requiredAddressColumnMapping` or `optionalAddressColumnMapping` objects, then they **CANNOT** have a different value in this object. If a different value is provided, then when the campaign is processing it will get overwritten with the mapped value present in the `requiredAddressColumnMapping` or `optionalAddressColumnMapping` objects.
	MergeVariableColumnMapping map[string]interface{} `json:"mergeVariableColumnMapping,omitempty"`
}

// NewUploadWritable instantiates a new UploadWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadWritable(campaignId string) *UploadWritable {
	this := UploadWritable{}
	this.CampaignId = campaignId
	return &this
}

// NewUploadWritableWithDefaults instantiates a new UploadWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadWritableWithDefaults() *UploadWritable {
	this := UploadWritable{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *UploadWritable) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *UploadWritable) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *UploadWritable) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetRequiredAddressColumnMapping returns the RequiredAddressColumnMapping field value if set, zero value otherwise.
func (o *UploadWritable) GetRequiredAddressColumnMapping() RequiredAddressColumnMapping {
	if o == nil || o.RequiredAddressColumnMapping == nil {
		var ret RequiredAddressColumnMapping
		return ret
	}
	return *o.RequiredAddressColumnMapping
}

// GetRequiredAddressColumnMappingOk returns a tuple with the RequiredAddressColumnMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadWritable) GetRequiredAddressColumnMappingOk() (*RequiredAddressColumnMapping, bool) {
	if o == nil || o.RequiredAddressColumnMapping == nil {
		return nil, false
	}
	return o.RequiredAddressColumnMapping, true
}

// HasRequiredAddressColumnMapping returns a boolean if a field has been set.
func (o *UploadWritable) HasRequiredAddressColumnMapping() bool {
	if o != nil && o.RequiredAddressColumnMapping != nil {
		return true
	}

	return false
}

// SetRequiredAddressColumnMapping gets a reference to the given RequiredAddressColumnMapping and assigns it to the RequiredAddressColumnMapping field.
func (o *UploadWritable) SetRequiredAddressColumnMapping(v RequiredAddressColumnMapping) {
	o.RequiredAddressColumnMapping = &v
}

// GetOptionalAddressColumnMapping returns the OptionalAddressColumnMapping field value if set, zero value otherwise.
func (o *UploadWritable) GetOptionalAddressColumnMapping() OptionalAddressColumnMapping {
	if o == nil || o.OptionalAddressColumnMapping == nil {
		var ret OptionalAddressColumnMapping
		return ret
	}
	return *o.OptionalAddressColumnMapping
}

// GetOptionalAddressColumnMappingOk returns a tuple with the OptionalAddressColumnMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadWritable) GetOptionalAddressColumnMappingOk() (*OptionalAddressColumnMapping, bool) {
	if o == nil || o.OptionalAddressColumnMapping == nil {
		return nil, false
	}
	return o.OptionalAddressColumnMapping, true
}

// HasOptionalAddressColumnMapping returns a boolean if a field has been set.
func (o *UploadWritable) HasOptionalAddressColumnMapping() bool {
	if o != nil && o.OptionalAddressColumnMapping != nil {
		return true
	}

	return false
}

// SetOptionalAddressColumnMapping gets a reference to the given OptionalAddressColumnMapping and assigns it to the OptionalAddressColumnMapping field.
func (o *UploadWritable) SetOptionalAddressColumnMapping(v OptionalAddressColumnMapping) {
	o.OptionalAddressColumnMapping = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UploadWritable) GetMetadata() UploadsMetadata {
	if o == nil || o.Metadata == nil {
		var ret UploadsMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadWritable) GetMetadataOk() (*UploadsMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UploadWritable) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given UploadsMetadata and assigns it to the Metadata field.
func (o *UploadWritable) SetMetadata(v UploadsMetadata) {
	o.Metadata = &v
}

// GetMergeVariableColumnMapping returns the MergeVariableColumnMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadWritable) GetMergeVariableColumnMapping() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MergeVariableColumnMapping
}

// GetMergeVariableColumnMappingOk returns a tuple with the MergeVariableColumnMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadWritable) GetMergeVariableColumnMappingOk() (map[string]interface{}, bool) {
	if o == nil || o.MergeVariableColumnMapping == nil {
		return nil, false
	}
	return o.MergeVariableColumnMapping, true
}

// HasMergeVariableColumnMapping returns a boolean if a field has been set.
func (o *UploadWritable) HasMergeVariableColumnMapping() bool {
	if o != nil && o.MergeVariableColumnMapping != nil {
		return true
	}

	return false
}

// SetMergeVariableColumnMapping gets a reference to the given map[string]interface{} and assigns it to the MergeVariableColumnMapping field.
func (o *UploadWritable) SetMergeVariableColumnMapping(v map[string]interface{}) {
	o.MergeVariableColumnMapping = v
}

func (o UploadWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["campaignId"] = o.CampaignId
	}
	if o.RequiredAddressColumnMapping != nil {
		toSerialize["requiredAddressColumnMapping"] = o.RequiredAddressColumnMapping
	}
	if o.OptionalAddressColumnMapping != nil {
		toSerialize["optionalAddressColumnMapping"] = o.OptionalAddressColumnMapping
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MergeVariableColumnMapping != nil {
		toSerialize["mergeVariableColumnMapping"] = o.MergeVariableColumnMapping
	}
	return json.Marshal(toSerialize)
}

type NullableUploadWritable struct {
	value *UploadWritable
	isSet bool
}

func (v NullableUploadWritable) Get() *UploadWritable {
	return v.value
}

func (v *NullableUploadWritable) Set(val *UploadWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadWritable(val *UploadWritable) *NullableUploadWritable {
	return &NullableUploadWritable{value: val, isSet: true}
}

func (v NullableUploadWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


