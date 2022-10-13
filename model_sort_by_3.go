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

// SortBy3 struct for SortBy3
type SortBy3 struct {
	DateCreated *string `json:"date_created,omitempty"`
	SendDate *string `json:"send_date,omitempty"`
}

// NewSortBy3 instantiates a new SortBy3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortBy3() *SortBy3 {
	this := SortBy3{}
	return &this
}

// NewSortBy3WithDefaults instantiates a new SortBy3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortBy3WithDefaults() *SortBy3 {
	this := SortBy3{}
	return &this
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SortBy3) GetDateCreated() string {
	if o == nil || o.DateCreated == nil {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortBy3) GetDateCreatedOk() (*string, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SortBy3) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *SortBy3) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *SortBy3) GetSendDate() string {
	if o == nil || o.SendDate == nil {
		var ret string
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortBy3) GetSendDateOk() (*string, bool) {
	if o == nil || o.SendDate == nil {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *SortBy3) HasSendDate() bool {
	if o != nil && o.SendDate != nil {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given string and assigns it to the SendDate field.
func (o *SortBy3) SetSendDate(v string) {
	o.SendDate = &v
}

func (o SortBy3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateCreated != nil {
		toSerialize["date_created"] = o.DateCreated
	}
	if o.SendDate != nil {
		toSerialize["send_date"] = o.SendDate
	}
	return json.Marshal(toSerialize)
}

type NullableSortBy3 struct {
	value *SortBy3
	isSet bool
}

func (v NullableSortBy3) Get() *SortBy3 {
	return v.value
}

func (v *NullableSortBy3) Set(val *SortBy3) {
	v.value = val
	v.isSet = true
}

func (v NullableSortBy3) IsSet() bool {
	return v.isSet
}

func (v *NullableSortBy3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortBy3(val *SortBy3) *NullableSortBy3 {
	return &NullableSortBy3{value: val, isSet: true}
}

func (v NullableSortBy3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortBy3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

