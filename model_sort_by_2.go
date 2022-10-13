/*
Lob

The Lob API is organized around REST. Our API is designed to have predictable, resource-oriented URLs and uses HTTP response codes to indicate any API errors. <p> Looking for our [previous documentation](https://lob.github.io/legacy-docs/)? 

API version: 1.3.0
Contact: lob-openapi@lob.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lob.go

import (
	"encoding/json"
)

// SortBy2 struct for SortBy2
type SortBy2 struct {
	DateCreated *string `json:"date_created,omitempty"`
	SendDate *string `json:"send_date,omitempty"`
}

// NewSortBy2 instantiates a new SortBy2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortBy2() *SortBy2 {
	this := SortBy2{}
	return &this
}

// NewSortBy2WithDefaults instantiates a new SortBy2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortBy2WithDefaults() *SortBy2 {
	this := SortBy2{}
	return &this
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SortBy2) GetDateCreated() string {
	if o == nil || o.DateCreated == nil {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortBy2) GetDateCreatedOk() (*string, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SortBy2) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *SortBy2) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *SortBy2) GetSendDate() string {
	if o == nil || o.SendDate == nil {
		var ret string
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortBy2) GetSendDateOk() (*string, bool) {
	if o == nil || o.SendDate == nil {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *SortBy2) HasSendDate() bool {
	if o != nil && o.SendDate != nil {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given string and assigns it to the SendDate field.
func (o *SortBy2) SetSendDate(v string) {
	o.SendDate = &v
}

func (o SortBy2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateCreated != nil {
		toSerialize["date_created"] = o.DateCreated
	}
	if o.SendDate != nil {
		toSerialize["send_date"] = o.SendDate
	}
	return json.Marshal(toSerialize)
}

type NullableSortBy2 struct {
	value *SortBy2
	isSet bool
}

func (v NullableSortBy2) Get() *SortBy2 {
	return v.value
}

func (v *NullableSortBy2) Set(val *SortBy2) {
	v.value = val
	v.isSet = true
}

func (v NullableSortBy2) IsSet() bool {
	return v.isSet
}

func (v *NullableSortBy2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortBy2(val *SortBy2) *NullableSortBy2 {
	return &NullableSortBy2{value: val, isSet: true}
}

func (v NullableSortBy2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortBy2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


