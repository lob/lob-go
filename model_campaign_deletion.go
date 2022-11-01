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

// CampaignDeletion Lob uses RESTful HTTP response codes to indicate success or failure of an API request. In general, 2xx indicates success, 4xx indicate an input error, and 5xx indicates an error on Lob's end.
type CampaignDeletion struct {
	// Unique identifier prefixed with `cmp_`.
	Id *string `json:"id,omitempty"`
	// True if the resource has been successfully deleted.
	Deleted *bool `json:"deleted,omitempty"`
}

// NewCampaignDeletion instantiates a new CampaignDeletion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignDeletion() *CampaignDeletion {
	this := CampaignDeletion{}
	return &this
}

// NewCampaignDeletionWithDefaults instantiates a new CampaignDeletion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignDeletionWithDefaults() *CampaignDeletion {
	this := CampaignDeletion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CampaignDeletion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDeletion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CampaignDeletion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CampaignDeletion) SetId(v string) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CampaignDeletion) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDeletion) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CampaignDeletion) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *CampaignDeletion) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o CampaignDeletion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignDeletion struct {
	value *CampaignDeletion
	isSet bool
}

func (v NullableCampaignDeletion) Get() *CampaignDeletion {
	return v.value
}

func (v *NullableCampaignDeletion) Set(val *CampaignDeletion) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignDeletion) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignDeletion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignDeletion(val *CampaignDeletion) *NullableCampaignDeletion {
	return &NullableCampaignDeletion{value: val, isSet: true}
}

func (v NullableCampaignDeletion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignDeletion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


