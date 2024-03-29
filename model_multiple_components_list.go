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

// MultipleComponentsList struct for MultipleComponentsList
type MultipleComponentsList struct {
	Addresses []MultipleComponents `json:"addresses"`
}

// NewMultipleComponentsList instantiates a new MultipleComponentsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleComponentsList(addresses []MultipleComponents) *MultipleComponentsList {
	this := MultipleComponentsList{}
	this.Addresses = addresses
	return &this
}

// NewMultipleComponentsListWithDefaults instantiates a new MultipleComponentsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleComponentsListWithDefaults() *MultipleComponentsList {
	this := MultipleComponentsList{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *MultipleComponentsList) GetAddresses() []MultipleComponents {
	if o == nil {
		var ret []MultipleComponents
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *MultipleComponentsList) GetAddressesOk() ([]MultipleComponents, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *MultipleComponentsList) SetAddresses(v []MultipleComponents) {
	o.Addresses = v
}

func (o MultipleComponentsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleComponentsList struct {
	value *MultipleComponentsList
	isSet bool
}

func (v NullableMultipleComponentsList) Get() *MultipleComponentsList {
	return v.value
}

func (v *NullableMultipleComponentsList) Set(val *MultipleComponentsList) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleComponentsList) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleComponentsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleComponentsList(val *MultipleComponentsList) *NullableMultipleComponentsList {
	return &NullableMultipleComponentsList{value: val, isSet: true}
}

func (v NullableMultipleComponentsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleComponentsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


