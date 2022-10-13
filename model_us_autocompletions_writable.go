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

// UsAutocompletionsWritable struct for UsAutocompletionsWritable
type UsAutocompletionsWritable struct {
	// Only accepts numbers and street names in an alphanumeric format. 
	AddressPrefix string `json:"address_prefix"`
	// An optional city input used to filter suggestions. Case insensitive and does not match partial abbreviations. 
	City *string `json:"city,omitempty"`
	// An optional state input used to filter suggestions. Case insensitive and does not match partial abbreviations. 
	State *string `json:"state,omitempty"`
	// An optional ZIP Code input used to filter suggestions. Matches partial entries. 
	ZipCode *string `json:"zip_code,omitempty"`
	// If `true`, sort suggestions by proximity to the IP set in the `X-Forwarded-For` header. 
	GeoIpSort *bool `json:"geo_ip_sort,omitempty"`
}

// NewUsAutocompletionsWritable instantiates a new UsAutocompletionsWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsAutocompletionsWritable(addressPrefix string) *UsAutocompletionsWritable {
	this := UsAutocompletionsWritable{}
	this.AddressPrefix = addressPrefix
	return &this
}

// NewUsAutocompletionsWritableWithDefaults instantiates a new UsAutocompletionsWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsAutocompletionsWritableWithDefaults() *UsAutocompletionsWritable {
	this := UsAutocompletionsWritable{}
	return &this
}

// GetAddressPrefix returns the AddressPrefix field value
func (o *UsAutocompletionsWritable) GetAddressPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressPrefix
}

// GetAddressPrefixOk returns a tuple with the AddressPrefix field value
// and a boolean to check if the value has been set.
func (o *UsAutocompletionsWritable) GetAddressPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressPrefix, true
}

// SetAddressPrefix sets field value
func (o *UsAutocompletionsWritable) SetAddressPrefix(v string) {
	o.AddressPrefix = v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *UsAutocompletionsWritable) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsAutocompletionsWritable) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UsAutocompletionsWritable) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *UsAutocompletionsWritable) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UsAutocompletionsWritable) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsAutocompletionsWritable) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UsAutocompletionsWritable) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UsAutocompletionsWritable) SetState(v string) {
	o.State = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *UsAutocompletionsWritable) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsAutocompletionsWritable) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *UsAutocompletionsWritable) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *UsAutocompletionsWritable) SetZipCode(v string) {
	o.ZipCode = &v
}

// GetGeoIpSort returns the GeoIpSort field value if set, zero value otherwise.
func (o *UsAutocompletionsWritable) GetGeoIpSort() bool {
	if o == nil || o.GeoIpSort == nil {
		var ret bool
		return ret
	}
	return *o.GeoIpSort
}

// GetGeoIpSortOk returns a tuple with the GeoIpSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsAutocompletionsWritable) GetGeoIpSortOk() (*bool, bool) {
	if o == nil || o.GeoIpSort == nil {
		return nil, false
	}
	return o.GeoIpSort, true
}

// HasGeoIpSort returns a boolean if a field has been set.
func (o *UsAutocompletionsWritable) HasGeoIpSort() bool {
	if o != nil && o.GeoIpSort != nil {
		return true
	}

	return false
}

// SetGeoIpSort gets a reference to the given bool and assigns it to the GeoIpSort field.
func (o *UsAutocompletionsWritable) SetGeoIpSort(v bool) {
	o.GeoIpSort = &v
}

func (o UsAutocompletionsWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address_prefix"] = o.AddressPrefix
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ZipCode != nil {
		toSerialize["zip_code"] = o.ZipCode
	}
	if o.GeoIpSort != nil {
		toSerialize["geo_ip_sort"] = o.GeoIpSort
	}
	return json.Marshal(toSerialize)
}

type NullableUsAutocompletionsWritable struct {
	value *UsAutocompletionsWritable
	isSet bool
}

func (v NullableUsAutocompletionsWritable) Get() *UsAutocompletionsWritable {
	return v.value
}

func (v *NullableUsAutocompletionsWritable) Set(val *UsAutocompletionsWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableUsAutocompletionsWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableUsAutocompletionsWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsAutocompletionsWritable(val *UsAutocompletionsWritable) *NullableUsAutocompletionsWritable {
	return &NullableUsAutocompletionsWritable{value: val, isSet: true}
}

func (v NullableUsAutocompletionsWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsAutocompletionsWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

