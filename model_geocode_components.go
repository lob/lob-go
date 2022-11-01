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

// GeocodeComponents A nested object containing a breakdown of each component of a reverse geocoded response.
type GeocodeComponents struct {
	// The 5-digit ZIP code
	ZipCode string `json:"zip_code"`
	ZipCodePlus4 string `json:"zip_code_plus_4"`
}

// NewGeocodeComponents instantiates a new GeocodeComponents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeocodeComponents(zipCode string, zipCodePlus4 string) *GeocodeComponents {
	this := GeocodeComponents{}
	this.ZipCode = zipCode
	this.ZipCodePlus4 = zipCodePlus4
	return &this
}

// NewGeocodeComponentsWithDefaults instantiates a new GeocodeComponents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeocodeComponentsWithDefaults() *GeocodeComponents {
	this := GeocodeComponents{}
	return &this
}

// GetZipCode returns the ZipCode field value
func (o *GeocodeComponents) GetZipCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value
// and a boolean to check if the value has been set.
func (o *GeocodeComponents) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZipCode, true
}

// SetZipCode sets field value
func (o *GeocodeComponents) SetZipCode(v string) {
	o.ZipCode = v
}

// GetZipCodePlus4 returns the ZipCodePlus4 field value
func (o *GeocodeComponents) GetZipCodePlus4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZipCodePlus4
}

// GetZipCodePlus4Ok returns a tuple with the ZipCodePlus4 field value
// and a boolean to check if the value has been set.
func (o *GeocodeComponents) GetZipCodePlus4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZipCodePlus4, true
}

// SetZipCodePlus4 sets field value
func (o *GeocodeComponents) SetZipCodePlus4(v string) {
	o.ZipCodePlus4 = v
}

func (o GeocodeComponents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["zip_code"] = o.ZipCode
	}
	if true {
		toSerialize["zip_code_plus_4"] = o.ZipCodePlus4
	}
	return json.Marshal(toSerialize)
}

type NullableGeocodeComponents struct {
	value *GeocodeComponents
	isSet bool
}

func (v NullableGeocodeComponents) Get() *GeocodeComponents {
	return v.value
}

func (v *NullableGeocodeComponents) Set(val *GeocodeComponents) {
	v.value = val
	v.isSet = true
}

func (v NullableGeocodeComponents) IsSet() bool {
	return v.isSet
}

func (v *NullableGeocodeComponents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeocodeComponents(val *GeocodeComponents) *NullableGeocodeComponents {
	return &NullableGeocodeComponents{value: val, isSet: true}
}

func (v NullableGeocodeComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeocodeComponents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


