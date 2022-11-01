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

// IntlComponents A nested object containing a breakdown of each component of an address.
type IntlComponents struct {
	// The numeric or alphanumeric part of an address preceding the street name. Often the house, building, or PO Box number.
	PrimaryNumber *string `json:"primary_number,omitempty"`
	// The name of the street.
	StreetName *string `json:"street_name,omitempty"`
	City *string `json:"city,omitempty"`
	// The [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) two letter code for the state. 
	State *string `json:"state,omitempty"`
	// The postal code.
	PostalCode *string `json:"postal_code,omitempty"`
}

// NewIntlComponents instantiates a new IntlComponents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntlComponents() *IntlComponents {
	this := IntlComponents{}
	return &this
}

// NewIntlComponentsWithDefaults instantiates a new IntlComponents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntlComponentsWithDefaults() *IntlComponents {
	this := IntlComponents{}
	return &this
}

// GetPrimaryNumber returns the PrimaryNumber field value if set, zero value otherwise.
func (o *IntlComponents) GetPrimaryNumber() string {
	if o == nil || o.PrimaryNumber == nil {
		var ret string
		return ret
	}
	return *o.PrimaryNumber
}

// GetPrimaryNumberOk returns a tuple with the PrimaryNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlComponents) GetPrimaryNumberOk() (*string, bool) {
	if o == nil || o.PrimaryNumber == nil {
		return nil, false
	}
	return o.PrimaryNumber, true
}

// HasPrimaryNumber returns a boolean if a field has been set.
func (o *IntlComponents) HasPrimaryNumber() bool {
	if o != nil && o.PrimaryNumber != nil {
		return true
	}

	return false
}

// SetPrimaryNumber gets a reference to the given string and assigns it to the PrimaryNumber field.
func (o *IntlComponents) SetPrimaryNumber(v string) {
	o.PrimaryNumber = &v
}

// GetStreetName returns the StreetName field value if set, zero value otherwise.
func (o *IntlComponents) GetStreetName() string {
	if o == nil || o.StreetName == nil {
		var ret string
		return ret
	}
	return *o.StreetName
}

// GetStreetNameOk returns a tuple with the StreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlComponents) GetStreetNameOk() (*string, bool) {
	if o == nil || o.StreetName == nil {
		return nil, false
	}
	return o.StreetName, true
}

// HasStreetName returns a boolean if a field has been set.
func (o *IntlComponents) HasStreetName() bool {
	if o != nil && o.StreetName != nil {
		return true
	}

	return false
}

// SetStreetName gets a reference to the given string and assigns it to the StreetName field.
func (o *IntlComponents) SetStreetName(v string) {
	o.StreetName = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *IntlComponents) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlComponents) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *IntlComponents) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *IntlComponents) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IntlComponents) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlComponents) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IntlComponents) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IntlComponents) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *IntlComponents) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlComponents) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *IntlComponents) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *IntlComponents) SetPostalCode(v string) {
	o.PostalCode = &v
}

func (o IntlComponents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrimaryNumber != nil {
		toSerialize["primary_number"] = o.PrimaryNumber
	}
	if o.StreetName != nil {
		toSerialize["street_name"] = o.StreetName
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	return json.Marshal(toSerialize)
}

type NullableIntlComponents struct {
	value *IntlComponents
	isSet bool
}

func (v NullableIntlComponents) Get() *IntlComponents {
	return v.value
}

func (v *NullableIntlComponents) Set(val *IntlComponents) {
	v.value = val
	v.isSet = true
}

func (v NullableIntlComponents) IsSet() bool {
	return v.isSet
}

func (v *NullableIntlComponents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntlComponents(val *IntlComponents) *NullableIntlComponents {
	return &NullableIntlComponents{value: val, isSet: true}
}

func (v NullableIntlComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntlComponents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


