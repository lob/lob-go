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

// IntlVerifications struct for IntlVerifications
type IntlVerifications struct {
	Addresses []IntlVerificationOrError `json:"addresses"`
	// Indicates whether any errors occurred during the verification process.
	Errors bool `json:"errors"`
}

// NewIntlVerifications instantiates a new IntlVerifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntlVerifications(addresses []IntlVerificationOrError, errors bool) *IntlVerifications {
	this := IntlVerifications{}
	this.Addresses = addresses
	this.Errors = errors
	return &this
}

// NewIntlVerificationsWithDefaults instantiates a new IntlVerifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntlVerificationsWithDefaults() *IntlVerifications {
	this := IntlVerifications{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *IntlVerifications) GetAddresses() []IntlVerificationOrError {
	if o == nil {
		var ret []IntlVerificationOrError
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *IntlVerifications) GetAddressesOk() ([]IntlVerificationOrError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *IntlVerifications) SetAddresses(v []IntlVerificationOrError) {
	o.Addresses = v
}

// GetErrors returns the Errors field value
func (o *IntlVerifications) GetErrors() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *IntlVerifications) GetErrorsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *IntlVerifications) SetErrors(v bool) {
	o.Errors = v
}

func (o IntlVerifications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableIntlVerifications struct {
	value *IntlVerifications
	isSet bool
}

func (v NullableIntlVerifications) Get() *IntlVerifications {
	return v.value
}

func (v *NullableIntlVerifications) Set(val *IntlVerifications) {
	v.value = val
	v.isSet = true
}

func (v NullableIntlVerifications) IsSet() bool {
	return v.isSet
}

func (v *NullableIntlVerifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntlVerifications(val *IntlVerifications) *NullableIntlVerifications {
	return &NullableIntlVerifications{value: val, isSet: true}
}

func (v NullableIntlVerifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntlVerifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

