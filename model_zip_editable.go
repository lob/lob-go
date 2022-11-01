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

// ZipEditable struct for ZipEditable
type ZipEditable struct {
	// A 5-digit ZIP code.
	ZipCode *string `json:"zip_code,omitempty"`
}

// NewZipEditable instantiates a new ZipEditable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZipEditable() *ZipEditable {
	this := ZipEditable{}
	return &this
}

// NewZipEditableWithDefaults instantiates a new ZipEditable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZipEditableWithDefaults() *ZipEditable {
	this := ZipEditable{}
	return &this
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *ZipEditable) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZipEditable) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *ZipEditable) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *ZipEditable) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o ZipEditable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZipCode != nil {
		toSerialize["zip_code"] = o.ZipCode
	}
	return json.Marshal(toSerialize)
}

type NullableZipEditable struct {
	value *ZipEditable
	isSet bool
}

func (v NullableZipEditable) Get() *ZipEditable {
	return v.value
}

func (v *NullableZipEditable) Set(val *ZipEditable) {
	v.value = val
	v.isSet = true
}

func (v NullableZipEditable) IsSet() bool {
	return v.isSet
}

func (v *NullableZipEditable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZipEditable(val *ZipEditable) *NullableZipEditable {
	return &NullableZipEditable{value: val, isSet: true}
}

func (v NullableZipEditable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZipEditable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


