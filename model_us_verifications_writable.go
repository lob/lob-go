/*
Lob

The Lob API is organized around REST. Our API is designed to have predictable, resource-oriented URLs and uses HTTP response codes to indicate any API errors. <p> Looking for our [previous documentation](https://lob.github.io/legacy-docs/)? 

API version: 1.3.0
Contact: lob-openapi@lob.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lob.go

import (
	"encoding/json"
)

// UsVerificationsWritable struct for UsVerificationsWritable
type UsVerificationsWritable struct {
	// The entire address in one string (e.g., \"210 King Street 94107\"). _Does not support a recipient and will error when other payload parameters are provided._ 
	Address *string `json:"address,omitempty"`
	// The intended recipient, typically a person's or firm's name.
	Recipient NullableString `json:"recipient,omitempty"`
	// The primary delivery line (usually the street address) of the address. Combination of the following applicable `components`: * `primary_number` * `street_predirection` * `street_name` * `street_suffix` * `street_postdirection` * `secondary_designator` * `secondary_number` * `pmb_designator` * `pmb_number` 
	PrimaryLine *string `json:"primary_line,omitempty"`
	// The secondary delivery line of the address. This field is typically empty but may contain information if `primary_line` is too long. 
	SecondaryLine *string `json:"secondary_line,omitempty"`
	// Only present for addresses in Puerto Rico. An urbanization refers to an area, sector, or development within a city. See [USPS documentation](https://pe.usps.com/text/pub28/28api_008.htm#:~:text=I51.,-4%20Urbanizations&text=In%20Puerto%20Rico%2C%20identical%20street,placed%20before%20the%20urbanization%20name.) for clarification. 
	Urbanization *string `json:"urbanization,omitempty"`
	City *string `json:"city,omitempty"`
	// The [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2:US) two letter code or subdivision name for the state. `city` and `state` are required if no `zip_code` is passed.
	State *string `json:"state,omitempty"`
	// Required if `city` and `state` are not passed in. If included, must be formatted as a US ZIP or ZIP+4 (e.g. `94107`, `941072282`, `94107-2282`).
	ZipCode *string `json:"zip_code,omitempty"`
}

// NewUsVerificationsWritable instantiates a new UsVerificationsWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsVerificationsWritable() *UsVerificationsWritable {
	this := UsVerificationsWritable{}
	return &this
}

// NewUsVerificationsWritableWithDefaults instantiates a new UsVerificationsWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsVerificationsWritableWithDefaults() *UsVerificationsWritable {
	this := UsVerificationsWritable{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UsVerificationsWritable) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerificationsWritable) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UsVerificationsWritable) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UsVerificationsWritable) SetAddress(v string) {
	o.Address = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsVerificationsWritable) GetRecipient() string {
	if o == nil || o.Recipient.Get() == nil {
		var ret string
		return ret
	}
	return *o.Recipient.Get()
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsVerificationsWritable) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipient.Get(), o.Recipient.IsSet()
}

// HasRecipient returns a boolean if a field has been set.
func (o *UsVerificationsWritable) HasRecipient() bool {
	if o != nil && o.Recipient.IsSet() {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given NullableString and assigns it to the Recipient field.
func (o *UsVerificationsWritable) SetRecipient(v string) {
	o.Recipient.Set(&v)
}
// SetRecipientNil sets the value for Recipient to be an explicit nil
func (o *UsVerificationsWritable) SetRecipientNil() {
	o.Recipient.Set(nil)
}

// UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
func (o *UsVerificationsWritable) UnsetRecipient() {
	o.Recipient.Unset()
}

// GetPrimaryLine returns the PrimaryLine field value if set, zero value otherwise.
func (o *UsVerificationsWritable) GetPrimaryLine() string {
	if o == nil || o.PrimaryLine == nil {
		var ret string
		return ret
	}
	return *o.PrimaryLine
}

// GetPrimaryLineOk returns a tuple with the PrimaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerificationsWritable) GetPrimaryLineOk() (*string, bool) {
	if o == nil || o.PrimaryLine == nil {
		return nil, false
	}
	return o.PrimaryLine, true
}

// HasPrimaryLine returns a boolean if a field has been set.
func (o *UsVerificationsWritable) HasPrimaryLine() bool {
	if o != nil && o.PrimaryLine != nil {
		return true
	}

	return false
}

// SetPrimaryLine gets a reference to the given string and assigns it to the PrimaryLine field.
func (o *UsVerificationsWritable) SetPrimaryLine(v string) {
	o.PrimaryLine = &v
}

// GetSecondaryLine returns the SecondaryLine field value if set, zero value otherwise.
func (o *UsVerificationsWritable) GetSecondaryLine() string {
	if o == nil || o.SecondaryLine == nil {
		var ret string
		return ret
	}
	return *o.SecondaryLine
}

// GetSecondaryLineOk returns a tuple with the SecondaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerificationsWritable) GetSecondaryLineOk() (*string, bool) {
	if o == nil || o.SecondaryLine == nil {
		return nil, false
	}
	return o.SecondaryLine, true
}

// HasSecondaryLine returns a boolean if a field has been set.
func (o *UsVerificationsWritable) HasSecondaryLine() bool {
	if o != nil && o.SecondaryLine != nil {
		return true
	}

	return false
}

// SetSecondaryLine gets a reference to the given string and assigns it to the SecondaryLine field.
func (o *UsVerificationsWritable) SetSecondaryLine(v string) {
	o.SecondaryLine = &v
}

// GetUrbanization returns the Urbanization field value if set, zero value otherwise.
func (o *UsVerificationsWritable) GetUrbanization() string {
	if o == nil || o.Urbanization == nil {
		var ret string
		return ret
	}
	return *o.Urbanization
}

// GetUrbanizationOk returns a tuple with the Urbanization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerificationsWritable) GetUrbanizationOk() (*string, bool) {
	if o == nil || o.Urbanization == nil {
		return nil, false
	}
	return o.Urbanization, true
}

// HasUrbanization returns a boolean if a field has been set.
func (o *UsVerificationsWritable) HasUrbanization() bool {
	if o != nil && o.Urbanization != nil {
		return true
	}

	return false
}

// SetUrbanization gets a reference to the given string and assigns it to the Urbanization field.
func (o *UsVerificationsWritable) SetUrbanization(v string) {
	o.Urbanization = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *UsVerificationsWritable) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerificationsWritable) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UsVerificationsWritable) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *UsVerificationsWritable) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UsVerificationsWritable) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerificationsWritable) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UsVerificationsWritable) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UsVerificationsWritable) SetState(v string) {
	o.State = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *UsVerificationsWritable) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerificationsWritable) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *UsVerificationsWritable) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *UsVerificationsWritable) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o UsVerificationsWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Recipient.IsSet() {
		toSerialize["recipient"] = o.Recipient.Get()
	}
	if o.PrimaryLine != nil {
		toSerialize["primary_line"] = o.PrimaryLine
	}
	if o.SecondaryLine != nil {
		toSerialize["secondary_line"] = o.SecondaryLine
	}
	if o.Urbanization != nil {
		toSerialize["urbanization"] = o.Urbanization
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
	return json.Marshal(toSerialize)
}

type NullableUsVerificationsWritable struct {
	value *UsVerificationsWritable
	isSet bool
}

func (v NullableUsVerificationsWritable) Get() *UsVerificationsWritable {
	return v.value
}

func (v *NullableUsVerificationsWritable) Set(val *UsVerificationsWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableUsVerificationsWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableUsVerificationsWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsVerificationsWritable(val *UsVerificationsWritable) *NullableUsVerificationsWritable {
	return &NullableUsVerificationsWritable{value: val, isSet: true}
}

func (v NullableUsVerificationsWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsVerificationsWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


