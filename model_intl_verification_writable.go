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

// IntlVerificationWritable struct for IntlVerificationWritable
type IntlVerificationWritable struct {
	// The intended recipient, typically a person's or firm's name.
	Recipient NullableString `json:"recipient,omitempty"`
	// The primary delivery line (usually the street address) of the address. 
	PrimaryLine *string `json:"primary_line,omitempty"`
	// The secondary delivery line of the address. This field is typically empty but may contain information if `primary_line` is too long. 
	SecondaryLine *string `json:"secondary_line,omitempty"`
	City *string `json:"city,omitempty"`
	// The name of the state.
	State *string `json:"state,omitempty"`
	// The postal code.
	PostalCode *string `json:"postal_code,omitempty"`
	Country *CountryExtended `json:"country,omitempty"`
	// The entire address in one string (e.g., \"370 Water St C1N 1C4\"). 
	Address *string `json:"address,omitempty"`
}

// NewIntlVerificationWritable instantiates a new IntlVerificationWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntlVerificationWritable() *IntlVerificationWritable {
	this := IntlVerificationWritable{}
	return &this
}

// NewIntlVerificationWritableWithDefaults instantiates a new IntlVerificationWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntlVerificationWritableWithDefaults() *IntlVerificationWritable {
	this := IntlVerificationWritable{}
	return &this
}

// GetRecipient returns the Recipient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntlVerificationWritable) GetRecipient() string {
	if o == nil || o.Recipient.Get() == nil {
		var ret string
		return ret
	}
	return *o.Recipient.Get()
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntlVerificationWritable) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipient.Get(), o.Recipient.IsSet()
}

// HasRecipient returns a boolean if a field has been set.
func (o *IntlVerificationWritable) HasRecipient() bool {
	if o != nil && o.Recipient.IsSet() {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given NullableString and assigns it to the Recipient field.
func (o *IntlVerificationWritable) SetRecipient(v string) {
	o.Recipient.Set(&v)
}
// SetRecipientNil sets the value for Recipient to be an explicit nil
func (o *IntlVerificationWritable) SetRecipientNil() {
	o.Recipient.Set(nil)
}

// UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
func (o *IntlVerificationWritable) UnsetRecipient() {
	o.Recipient.Unset()
}

// GetPrimaryLine returns the PrimaryLine field value if set, zero value otherwise.
func (o *IntlVerificationWritable) GetPrimaryLine() string {
	if o == nil || o.PrimaryLine == nil {
		var ret string
		return ret
	}
	return *o.PrimaryLine
}

// GetPrimaryLineOk returns a tuple with the PrimaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationWritable) GetPrimaryLineOk() (*string, bool) {
	if o == nil || o.PrimaryLine == nil {
		return nil, false
	}
	return o.PrimaryLine, true
}

// HasPrimaryLine returns a boolean if a field has been set.
func (o *IntlVerificationWritable) HasPrimaryLine() bool {
	if o != nil && o.PrimaryLine != nil {
		return true
	}

	return false
}

// SetPrimaryLine gets a reference to the given string and assigns it to the PrimaryLine field.
func (o *IntlVerificationWritable) SetPrimaryLine(v string) {
	o.PrimaryLine = &v
}

// GetSecondaryLine returns the SecondaryLine field value if set, zero value otherwise.
func (o *IntlVerificationWritable) GetSecondaryLine() string {
	if o == nil || o.SecondaryLine == nil {
		var ret string
		return ret
	}
	return *o.SecondaryLine
}

// GetSecondaryLineOk returns a tuple with the SecondaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationWritable) GetSecondaryLineOk() (*string, bool) {
	if o == nil || o.SecondaryLine == nil {
		return nil, false
	}
	return o.SecondaryLine, true
}

// HasSecondaryLine returns a boolean if a field has been set.
func (o *IntlVerificationWritable) HasSecondaryLine() bool {
	if o != nil && o.SecondaryLine != nil {
		return true
	}

	return false
}

// SetSecondaryLine gets a reference to the given string and assigns it to the SecondaryLine field.
func (o *IntlVerificationWritable) SetSecondaryLine(v string) {
	o.SecondaryLine = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *IntlVerificationWritable) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationWritable) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *IntlVerificationWritable) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *IntlVerificationWritable) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IntlVerificationWritable) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationWritable) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IntlVerificationWritable) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IntlVerificationWritable) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *IntlVerificationWritable) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationWritable) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *IntlVerificationWritable) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *IntlVerificationWritable) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *IntlVerificationWritable) GetCountry() CountryExtended {
	if o == nil || o.Country == nil {
		var ret CountryExtended
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationWritable) GetCountryOk() (*CountryExtended, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *IntlVerificationWritable) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given CountryExtended and assigns it to the Country field.
func (o *IntlVerificationWritable) SetCountry(v CountryExtended) {
	o.Country = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IntlVerificationWritable) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationWritable) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IntlVerificationWritable) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *IntlVerificationWritable) SetAddress(v string) {
	o.Address = &v
}

func (o IntlVerificationWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Recipient.IsSet() {
		toSerialize["recipient"] = o.Recipient.Get()
	}
	if o.PrimaryLine != nil {
		toSerialize["primary_line"] = o.PrimaryLine
	}
	if o.SecondaryLine != nil {
		toSerialize["secondary_line"] = o.SecondaryLine
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
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableIntlVerificationWritable struct {
	value *IntlVerificationWritable
	isSet bool
}

func (v NullableIntlVerificationWritable) Get() *IntlVerificationWritable {
	return v.value
}

func (v *NullableIntlVerificationWritable) Set(val *IntlVerificationWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableIntlVerificationWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableIntlVerificationWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntlVerificationWritable(val *IntlVerificationWritable) *NullableIntlVerificationWritable {
	return &NullableIntlVerificationWritable{value: val, isSet: true}
}

func (v NullableIntlVerificationWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntlVerificationWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


