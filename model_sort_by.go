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

// SortBy struct for SortBy
type SortBy struct {
	DateCreated *string `json:"date_created,omitempty"`
	SendDate *string `json:"send_date,omitempty"`
}

// NewSortBy instantiates a new SortBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortBy() *SortBy {
	this := SortBy{}
	return &this
}

// NewSortByWithDefaults instantiates a new SortBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortByWithDefaults() *SortBy {
	this := SortBy{}
	return &this
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SortBy) GetDateCreated() string {
	if o == nil || o.DateCreated == nil {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortBy) GetDateCreatedOk() (*string, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SortBy) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *SortBy) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *SortBy) GetSendDate() string {
	if o == nil || o.SendDate == nil {
		var ret string
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortBy) GetSendDateOk() (*string, bool) {
	if o == nil || o.SendDate == nil {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *SortBy) HasSendDate() bool {
	if o != nil && o.SendDate != nil {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given string and assigns it to the SendDate field.
func (o *SortBy) SetSendDate(v string) {
	o.SendDate = &v
}

func (o SortBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateCreated != nil {
		toSerialize["date_created"] = o.DateCreated
	}
	if o.SendDate != nil {
		toSerialize["send_date"] = o.SendDate
	}
	return json.Marshal(toSerialize)
}

type NullableSortBy struct {
	value *SortBy
	isSet bool
}

func (v NullableSortBy) Get() *SortBy {
	return v.value
}

func (v *NullableSortBy) Set(val *SortBy) {
	v.value = val
	v.isSet = true
}

func (v NullableSortBy) IsSet() bool {
	return v.isSet
}

func (v *NullableSortBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortBy(val *SortBy) *NullableSortBy {
	return &NullableSortBy{value: val, isSet: true}
}

func (v NullableSortBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


