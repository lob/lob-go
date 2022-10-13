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

// SortByDateModified struct for SortByDateModified
type SortByDateModified struct {
	DateCreated *string `json:"date_created,omitempty"`
	DateModified *string `json:"date_modified,omitempty"`
}

// NewSortByDateModified instantiates a new SortByDateModified object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortByDateModified() *SortByDateModified {
	this := SortByDateModified{}
	return &this
}

// NewSortByDateModifiedWithDefaults instantiates a new SortByDateModified object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortByDateModifiedWithDefaults() *SortByDateModified {
	this := SortByDateModified{}
	return &this
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SortByDateModified) GetDateCreated() string {
	if o == nil || o.DateCreated == nil {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortByDateModified) GetDateCreatedOk() (*string, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SortByDateModified) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *SortByDateModified) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetDateModified returns the DateModified field value if set, zero value otherwise.
func (o *SortByDateModified) GetDateModified() string {
	if o == nil || o.DateModified == nil {
		var ret string
		return ret
	}
	return *o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortByDateModified) GetDateModifiedOk() (*string, bool) {
	if o == nil || o.DateModified == nil {
		return nil, false
	}
	return o.DateModified, true
}

// HasDateModified returns a boolean if a field has been set.
func (o *SortByDateModified) HasDateModified() bool {
	if o != nil && o.DateModified != nil {
		return true
	}

	return false
}

// SetDateModified gets a reference to the given string and assigns it to the DateModified field.
func (o *SortByDateModified) SetDateModified(v string) {
	o.DateModified = &v
}

func (o SortByDateModified) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateCreated != nil {
		toSerialize["date_created"] = o.DateCreated
	}
	if o.DateModified != nil {
		toSerialize["date_modified"] = o.DateModified
	}
	return json.Marshal(toSerialize)
}

type NullableSortByDateModified struct {
	value *SortByDateModified
	isSet bool
}

func (v NullableSortByDateModified) Get() *SortByDateModified {
	return v.value
}

func (v *NullableSortByDateModified) Set(val *SortByDateModified) {
	v.value = val
	v.isSet = true
}

func (v NullableSortByDateModified) IsSet() bool {
	return v.isSet
}

func (v *NullableSortByDateModified) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortByDateModified(val *SortByDateModified) *NullableSortByDateModified {
	return &NullableSortByDateModified{value: val, isSet: true}
}

func (v NullableSortByDateModified) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortByDateModified) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

