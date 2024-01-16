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

// OptionalAddressColumnMapping The mapping of column headers in your file to Lob-optional fields for the resource created. See our <a href=\"https://help.lob.com/print-and-mail/building-a-mail-strategy/campaign-or-triggered-sends/campaign-audience-guide#optional-columns-3\" target=\"_blank\">Campaign Audience Guide</a> for additional details.
type OptionalAddressColumnMapping struct {
	// The column header from the csv file that should be mapped to the optional field \"address_line2\"
	AddressLine2 NullableString `json:"address_line2"`
	// The column header from the csv file that should be mapped to the optional field \"company\"
	Company NullableString `json:"company"`
	// The column header from the csv file that should be mapped to the optional field \"address_country\"
	AddressCountry NullableString `json:"address_country"`
}

// NewOptionalAddressColumnMapping instantiates a new OptionalAddressColumnMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionalAddressColumnMapping(addressLine2 NullableString, company NullableString, addressCountry NullableString) *OptionalAddressColumnMapping {
	this := OptionalAddressColumnMapping{}
	this.AddressLine2 = addressLine2
	this.Company = company
	this.AddressCountry = addressCountry
	return &this
}

// NewOptionalAddressColumnMappingWithDefaults instantiates a new OptionalAddressColumnMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionalAddressColumnMappingWithDefaults() *OptionalAddressColumnMapping {
	this := OptionalAddressColumnMapping{}
	var addressLine2 string = "null"
	this.AddressLine2 = *NewNullableString(&addressLine2)
	var company string = "null"
	this.Company = *NewNullableString(&company)
	var addressCountry string = "null"
	this.AddressCountry = *NewNullableString(&addressCountry)
	return &this
}

// GetAddressLine2 returns the AddressLine2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OptionalAddressColumnMapping) GetAddressLine2() string {
	if o == nil || o.AddressLine2.Get() == nil {
		var ret string
		return ret
	}

	return *o.AddressLine2.Get()
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OptionalAddressColumnMapping) GetAddressLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressLine2.Get(), o.AddressLine2.IsSet()
}

// SetAddressLine2 sets field value
func (o *OptionalAddressColumnMapping) SetAddressLine2(v string) {
	o.AddressLine2.Set(&v)
}

// GetCompany returns the Company field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OptionalAddressColumnMapping) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}

	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OptionalAddressColumnMapping) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// SetCompany sets field value
func (o *OptionalAddressColumnMapping) SetCompany(v string) {
	o.Company.Set(&v)
}

// GetAddressCountry returns the AddressCountry field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OptionalAddressColumnMapping) GetAddressCountry() string {
	if o == nil || o.AddressCountry.Get() == nil {
		var ret string
		return ret
	}

	return *o.AddressCountry.Get()
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OptionalAddressColumnMapping) GetAddressCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressCountry.Get(), o.AddressCountry.IsSet()
}

// SetAddressCountry sets field value
func (o *OptionalAddressColumnMapping) SetAddressCountry(v string) {
	o.AddressCountry.Set(&v)
}

func (o OptionalAddressColumnMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address_line2"] = o.AddressLine2.Get()
	}
	if true {
		toSerialize["company"] = o.Company.Get()
	}
	if true {
		toSerialize["address_country"] = o.AddressCountry.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOptionalAddressColumnMapping struct {
	value *OptionalAddressColumnMapping
	isSet bool
}

func (v NullableOptionalAddressColumnMapping) Get() *OptionalAddressColumnMapping {
	return v.value
}

func (v *NullableOptionalAddressColumnMapping) Set(val *OptionalAddressColumnMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionalAddressColumnMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionalAddressColumnMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionalAddressColumnMapping(val *OptionalAddressColumnMapping) *NullableOptionalAddressColumnMapping {
	return &NullableOptionalAddressColumnMapping{value: val, isSet: true}
}

func (v NullableOptionalAddressColumnMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionalAddressColumnMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

