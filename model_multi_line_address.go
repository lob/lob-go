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

// MultiLineAddress struct for MultiLineAddress
type MultiLineAddress struct {
	// The intended recipient, typically a person's or firm's name.
	Recipient NullableString `json:"recipient,omitempty"`
	// Either `name` or `company` is required, you may also add both.
	Company NullableString `json:"company,omitempty"`
	// The primary delivery line (usually the street address) of the address. Combination of the following applicable `components`: * `primary_number` * `street_predirection` * `street_name` * `street_suffix` * `street_postdirection` * `secondary_designator` * `secondary_number` * `pmb_designator` * `pmb_number` 
	PrimaryLine string `json:"primary_line"`
	// The secondary delivery line of the address. This field is typically empty but may contain information if `primary_line` is too long. 
	SecondaryLine *string `json:"secondary_line,omitempty"`
	// Only present for addresses in Puerto Rico. An urbanization refers to an area, sector, or development within a city. See [USPS documentation](https://pe.usps.com/text/pub28/28api_008.htm#:~:text=I51.,-4%20Urbanizations&text=In%20Puerto%20Rico%2C%20identical%20street,placed%20before%20the%20urbanization%20name.) for clarification. 
	Urbanization *string `json:"urbanization,omitempty"`
	City *string `json:"city,omitempty"`
	// The <a href=\"https://en.wikipedia.org/wiki/ISO_3166-2:US\" target=\"_blank\">ISO 3166-2</a> two letter code or subdivision name for the state. `city` and `state` are required if no `zip_code` is passed.
	State *string `json:"state,omitempty"`
	// Required if `city` and `state` are not passed in. If included, must be formatted as a US ZIP or ZIP+4 (e.g. `94107`, `941072282`, `94107-2282`).
	ZipCode *string `json:"zip_code,omitempty"`
}

// NewMultiLineAddress instantiates a new MultiLineAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiLineAddress(primaryLine string) *MultiLineAddress {
	this := MultiLineAddress{}
	this.PrimaryLine = primaryLine
	return &this
}

// NewMultiLineAddressWithDefaults instantiates a new MultiLineAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiLineAddressWithDefaults() *MultiLineAddress {
	this := MultiLineAddress{}
	return &this
}

// GetRecipient returns the Recipient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultiLineAddress) GetRecipient() string {
	if o == nil || o.Recipient.Get() == nil {
		var ret string
		return ret
	}
	return *o.Recipient.Get()
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultiLineAddress) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipient.Get(), o.Recipient.IsSet()
}

// HasRecipient returns a boolean if a field has been set.
func (o *MultiLineAddress) HasRecipient() bool {
	if o != nil && o.Recipient.IsSet() {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given NullableString and assigns it to the Recipient field.
func (o *MultiLineAddress) SetRecipient(v string) {
	o.Recipient.Set(&v)
}
// SetRecipientNil sets the value for Recipient to be an explicit nil
func (o *MultiLineAddress) SetRecipientNil() {
	o.Recipient.Set(nil)
}

// UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
func (o *MultiLineAddress) UnsetRecipient() {
	o.Recipient.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultiLineAddress) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultiLineAddress) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *MultiLineAddress) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *MultiLineAddress) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *MultiLineAddress) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *MultiLineAddress) UnsetCompany() {
	o.Company.Unset()
}

// GetPrimaryLine returns the PrimaryLine field value
func (o *MultiLineAddress) GetPrimaryLine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryLine
}

// GetPrimaryLineOk returns a tuple with the PrimaryLine field value
// and a boolean to check if the value has been set.
func (o *MultiLineAddress) GetPrimaryLineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryLine, true
}

// SetPrimaryLine sets field value
func (o *MultiLineAddress) SetPrimaryLine(v string) {
	o.PrimaryLine = v
}

// GetSecondaryLine returns the SecondaryLine field value if set, zero value otherwise.
func (o *MultiLineAddress) GetSecondaryLine() string {
	if o == nil || o.SecondaryLine == nil {
		var ret string
		return ret
	}
	return *o.SecondaryLine
}

// GetSecondaryLineOk returns a tuple with the SecondaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiLineAddress) GetSecondaryLineOk() (*string, bool) {
	if o == nil || o.SecondaryLine == nil {
		return nil, false
	}
	return o.SecondaryLine, true
}

// HasSecondaryLine returns a boolean if a field has been set.
func (o *MultiLineAddress) HasSecondaryLine() bool {
	if o != nil && o.SecondaryLine != nil {
		return true
	}

	return false
}

// SetSecondaryLine gets a reference to the given string and assigns it to the SecondaryLine field.
func (o *MultiLineAddress) SetSecondaryLine(v string) {
	o.SecondaryLine = &v
}

// GetUrbanization returns the Urbanization field value if set, zero value otherwise.
func (o *MultiLineAddress) GetUrbanization() string {
	if o == nil || o.Urbanization == nil {
		var ret string
		return ret
	}
	return *o.Urbanization
}

// GetUrbanizationOk returns a tuple with the Urbanization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiLineAddress) GetUrbanizationOk() (*string, bool) {
	if o == nil || o.Urbanization == nil {
		return nil, false
	}
	return o.Urbanization, true
}

// HasUrbanization returns a boolean if a field has been set.
func (o *MultiLineAddress) HasUrbanization() bool {
	if o != nil && o.Urbanization != nil {
		return true
	}

	return false
}

// SetUrbanization gets a reference to the given string and assigns it to the Urbanization field.
func (o *MultiLineAddress) SetUrbanization(v string) {
	o.Urbanization = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *MultiLineAddress) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiLineAddress) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *MultiLineAddress) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *MultiLineAddress) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MultiLineAddress) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiLineAddress) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MultiLineAddress) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MultiLineAddress) SetState(v string) {
	o.State = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *MultiLineAddress) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiLineAddress) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *MultiLineAddress) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *MultiLineAddress) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o MultiLineAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Recipient.IsSet() {
		toSerialize["recipient"] = o.Recipient.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if true {
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

type NullableMultiLineAddress struct {
	value *MultiLineAddress
	isSet bool
}

func (v NullableMultiLineAddress) Get() *MultiLineAddress {
	return v.value
}

func (v *NullableMultiLineAddress) Set(val *MultiLineAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiLineAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiLineAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiLineAddress(val *MultiLineAddress) *NullableMultiLineAddress {
	return &NullableMultiLineAddress{value: val, isSet: true}
}

func (v NullableMultiLineAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiLineAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


