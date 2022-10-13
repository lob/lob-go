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

// AddressEditable struct for AddressEditable
type AddressEditable struct {
	// The building number, street name, street suffix, and any street directionals. For US addresses, the max length is 64 characters.
	AddressLine1 *string `json:"address_line1,omitempty"`
	// The suite or apartment number of the recipient address, if applicable. For US addresses, the max length is 64 characters.
	AddressLine2 NullableString `json:"address_line2,omitempty"`
	AddressCity NullableString `json:"address_city,omitempty"`
	AddressState NullableString `json:"address_state,omitempty"`
	// Optional postal code. For US addresses, must be either 5 or 9 digits.
	AddressZip NullableString `json:"address_zip,omitempty"`
	AddressCountry *CountryExtended `json:"address_country,omitempty"`
	// An internal description that identifies this resource. Must be no longer than 255 characters. 
	Description NullableString `json:"description,omitempty"`
	// name associated with address.
	Name NullableString `json:"name,omitempty"`
	// Either `name` or `company` is required, you may also add both.
	Company NullableString `json:"company,omitempty"`
	// Must be no longer than 40 characters.
	Phone NullableString `json:"phone,omitempty"`
	// Must be no longer than 100 characters.
	Email NullableString `json:"email,omitempty"`
	// Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters `\"` and `\\`. i.e. '{\"customer_id\" : \"NEWYORK2015\"}' Nested objects are not supported.  See [Metadata](#section/Metadata) for more information.
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewAddressEditable instantiates a new AddressEditable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressEditable() *AddressEditable {
	this := AddressEditable{}
	return &this
}

// NewAddressEditableWithDefaults instantiates a new AddressEditable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressEditableWithDefaults() *AddressEditable {
	this := AddressEditable{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *AddressEditable) GetAddressLine1() string {
	if o == nil || o.AddressLine1 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressEditable) GetAddressLine1Ok() (*string, bool) {
	if o == nil || o.AddressLine1 == nil {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *AddressEditable) HasAddressLine1() bool {
	if o != nil && o.AddressLine1 != nil {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *AddressEditable) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressEditable) GetAddressLine2() string {
	if o == nil || o.AddressLine2.Get() == nil {
		var ret string
		return ret
	}
	return *o.AddressLine2.Get()
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressEditable) GetAddressLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressLine2.Get(), o.AddressLine2.IsSet()
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *AddressEditable) HasAddressLine2() bool {
	if o != nil && o.AddressLine2.IsSet() {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given NullableString and assigns it to the AddressLine2 field.
func (o *AddressEditable) SetAddressLine2(v string) {
	o.AddressLine2.Set(&v)
}
// SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil
func (o *AddressEditable) SetAddressLine2Nil() {
	o.AddressLine2.Set(nil)
}

// UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
func (o *AddressEditable) UnsetAddressLine2() {
	o.AddressLine2.Unset()
}

// GetAddressCity returns the AddressCity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressEditable) GetAddressCity() string {
	if o == nil || o.AddressCity.Get() == nil {
		var ret string
		return ret
	}
	return *o.AddressCity.Get()
}

// GetAddressCityOk returns a tuple with the AddressCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressEditable) GetAddressCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressCity.Get(), o.AddressCity.IsSet()
}

// HasAddressCity returns a boolean if a field has been set.
func (o *AddressEditable) HasAddressCity() bool {
	if o != nil && o.AddressCity.IsSet() {
		return true
	}

	return false
}

// SetAddressCity gets a reference to the given NullableString and assigns it to the AddressCity field.
func (o *AddressEditable) SetAddressCity(v string) {
	o.AddressCity.Set(&v)
}
// SetAddressCityNil sets the value for AddressCity to be an explicit nil
func (o *AddressEditable) SetAddressCityNil() {
	o.AddressCity.Set(nil)
}

// UnsetAddressCity ensures that no value is present for AddressCity, not even an explicit nil
func (o *AddressEditable) UnsetAddressCity() {
	o.AddressCity.Unset()
}

// GetAddressState returns the AddressState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressEditable) GetAddressState() string {
	if o == nil || o.AddressState.Get() == nil {
		var ret string
		return ret
	}
	return *o.AddressState.Get()
}

// GetAddressStateOk returns a tuple with the AddressState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressEditable) GetAddressStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressState.Get(), o.AddressState.IsSet()
}

// HasAddressState returns a boolean if a field has been set.
func (o *AddressEditable) HasAddressState() bool {
	if o != nil && o.AddressState.IsSet() {
		return true
	}

	return false
}

// SetAddressState gets a reference to the given NullableString and assigns it to the AddressState field.
func (o *AddressEditable) SetAddressState(v string) {
	o.AddressState.Set(&v)
}
// SetAddressStateNil sets the value for AddressState to be an explicit nil
func (o *AddressEditable) SetAddressStateNil() {
	o.AddressState.Set(nil)
}

// UnsetAddressState ensures that no value is present for AddressState, not even an explicit nil
func (o *AddressEditable) UnsetAddressState() {
	o.AddressState.Unset()
}

// GetAddressZip returns the AddressZip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressEditable) GetAddressZip() string {
	if o == nil || o.AddressZip.Get() == nil {
		var ret string
		return ret
	}
	return *o.AddressZip.Get()
}

// GetAddressZipOk returns a tuple with the AddressZip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressEditable) GetAddressZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressZip.Get(), o.AddressZip.IsSet()
}

// HasAddressZip returns a boolean if a field has been set.
func (o *AddressEditable) HasAddressZip() bool {
	if o != nil && o.AddressZip.IsSet() {
		return true
	}

	return false
}

// SetAddressZip gets a reference to the given NullableString and assigns it to the AddressZip field.
func (o *AddressEditable) SetAddressZip(v string) {
	o.AddressZip.Set(&v)
}
// SetAddressZipNil sets the value for AddressZip to be an explicit nil
func (o *AddressEditable) SetAddressZipNil() {
	o.AddressZip.Set(nil)
}

// UnsetAddressZip ensures that no value is present for AddressZip, not even an explicit nil
func (o *AddressEditable) UnsetAddressZip() {
	o.AddressZip.Unset()
}

// GetAddressCountry returns the AddressCountry field value if set, zero value otherwise.
func (o *AddressEditable) GetAddressCountry() CountryExtended {
	if o == nil || o.AddressCountry == nil {
		var ret CountryExtended
		return ret
	}
	return *o.AddressCountry
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressEditable) GetAddressCountryOk() (*CountryExtended, bool) {
	if o == nil || o.AddressCountry == nil {
		return nil, false
	}
	return o.AddressCountry, true
}

// HasAddressCountry returns a boolean if a field has been set.
func (o *AddressEditable) HasAddressCountry() bool {
	if o != nil && o.AddressCountry != nil {
		return true
	}

	return false
}

// SetAddressCountry gets a reference to the given CountryExtended and assigns it to the AddressCountry field.
func (o *AddressEditable) SetAddressCountry(v CountryExtended) {
	o.AddressCountry = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressEditable) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressEditable) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AddressEditable) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AddressEditable) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AddressEditable) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AddressEditable) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressEditable) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressEditable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AddressEditable) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AddressEditable) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AddressEditable) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AddressEditable) UnsetName() {
	o.Name.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressEditable) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressEditable) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *AddressEditable) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *AddressEditable) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *AddressEditable) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *AddressEditable) UnsetCompany() {
	o.Company.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressEditable) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressEditable) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *AddressEditable) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *AddressEditable) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *AddressEditable) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *AddressEditable) UnsetPhone() {
	o.Phone.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressEditable) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressEditable) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *AddressEditable) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *AddressEditable) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *AddressEditable) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *AddressEditable) UnsetEmail() {
	o.Email.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AddressEditable) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressEditable) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AddressEditable) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *AddressEditable) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o AddressEditable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressLine1 != nil {
		toSerialize["address_line1"] = o.AddressLine1
	}
	if o.AddressLine2.IsSet() {
		toSerialize["address_line2"] = o.AddressLine2.Get()
	}
	if o.AddressCity.IsSet() {
		toSerialize["address_city"] = o.AddressCity.Get()
	}
	if o.AddressState.IsSet() {
		toSerialize["address_state"] = o.AddressState.Get()
	}
	if o.AddressZip.IsSet() {
		toSerialize["address_zip"] = o.AddressZip.Get()
	}
	if o.AddressCountry != nil {
		toSerialize["address_country"] = o.AddressCountry
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableAddressEditable struct {
	value *AddressEditable
	isSet bool
}

func (v NullableAddressEditable) Get() *AddressEditable {
	return v.value
}

func (v *NullableAddressEditable) Set(val *AddressEditable) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressEditable) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressEditable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressEditable(val *AddressEditable) *NullableAddressEditable {
	return &NullableAddressEditable{value: val, isSet: true}
}

func (v NullableAddressEditable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressEditable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

