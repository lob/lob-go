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

// LetterDetailsWritable Properties that the letters in your Creative should have.
type LetterDetailsWritable struct {
	// Specifies the location of the address information that will show through the double-window envelope. 
	AddressPlacement *string `json:"address_placement,omitempty"`
	// A single-element array containing an existing card id in a string format. See [cards](#tag/Cards) for more information.
	Cards []string `json:"cards"`
	// Set this key to `true` if you would like to print in color. Set to `false` if you would like to print in black and white.
	Color bool `json:"color"`
	// Set this attribute to `true` for double sided printing, or `false` for for single sided printing. Defaults to `true`.
	DoubleSided *bool `json:"double_sided,omitempty"`
	// Add an extra service to your letter.
	ExtraService *string `json:"extra_service,omitempty"`
	MailType *MailType `json:"mail_type,omitempty"`
	ReturnEnvelope *bool `json:"return_envelope,omitempty"`
	// Accepts an envelope ID for any customized envelope with available inventory.
	CustomEnvelope NullableString `json:"custom_envelope,omitempty"`
}

// NewLetterDetailsWritable instantiates a new LetterDetailsWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLetterDetailsWritable(cards []string, color bool) *LetterDetailsWritable {
	this := LetterDetailsWritable{}
	var addressPlacement string = "top_first_page"
	this.AddressPlacement = &addressPlacement
	this.Cards = cards
	this.Color = color
	var doubleSided bool = true
	this.DoubleSided = &doubleSided
	var mailType MailType = MAILTYPE_FIRST_CLASS
	this.MailType = &mailType
	var returnEnvelope bool = false
	this.ReturnEnvelope = &returnEnvelope
	return &this
}

// NewLetterDetailsWritableWithDefaults instantiates a new LetterDetailsWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLetterDetailsWritableWithDefaults() *LetterDetailsWritable {
	this := LetterDetailsWritable{}
	var addressPlacement string = "top_first_page"
	this.AddressPlacement = &addressPlacement
	var doubleSided bool = true
	this.DoubleSided = &doubleSided
	var mailType MailType = MAILTYPE_FIRST_CLASS
	this.MailType = &mailType
	var returnEnvelope bool = false
	this.ReturnEnvelope = &returnEnvelope
	return &this
}

// GetAddressPlacement returns the AddressPlacement field value if set, zero value otherwise.
func (o *LetterDetailsWritable) GetAddressPlacement() string {
	if o == nil || o.AddressPlacement == nil {
		var ret string
		return ret
	}
	return *o.AddressPlacement
}

// GetAddressPlacementOk returns a tuple with the AddressPlacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LetterDetailsWritable) GetAddressPlacementOk() (*string, bool) {
	if o == nil || o.AddressPlacement == nil {
		return nil, false
	}
	return o.AddressPlacement, true
}

// HasAddressPlacement returns a boolean if a field has been set.
func (o *LetterDetailsWritable) HasAddressPlacement() bool {
	if o != nil && o.AddressPlacement != nil {
		return true
	}

	return false
}

// SetAddressPlacement gets a reference to the given string and assigns it to the AddressPlacement field.
func (o *LetterDetailsWritable) SetAddressPlacement(v string) {
	o.AddressPlacement = &v
}

// GetCards returns the Cards field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *LetterDetailsWritable) GetCards() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cards
}

// GetCardsOk returns a tuple with the Cards field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LetterDetailsWritable) GetCardsOk() ([]string, bool) {
	if o == nil || o.Cards == nil {
		return nil, false
	}
	return o.Cards, true
}

// SetCards sets field value
func (o *LetterDetailsWritable) SetCards(v []string) {
	o.Cards = v
}

// GetColor returns the Color field value
func (o *LetterDetailsWritable) GetColor() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *LetterDetailsWritable) GetColorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *LetterDetailsWritable) SetColor(v bool) {
	o.Color = v
}

// GetDoubleSided returns the DoubleSided field value if set, zero value otherwise.
func (o *LetterDetailsWritable) GetDoubleSided() bool {
	if o == nil || o.DoubleSided == nil {
		var ret bool
		return ret
	}
	return *o.DoubleSided
}

// GetDoubleSidedOk returns a tuple with the DoubleSided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LetterDetailsWritable) GetDoubleSidedOk() (*bool, bool) {
	if o == nil || o.DoubleSided == nil {
		return nil, false
	}
	return o.DoubleSided, true
}

// HasDoubleSided returns a boolean if a field has been set.
func (o *LetterDetailsWritable) HasDoubleSided() bool {
	if o != nil && o.DoubleSided != nil {
		return true
	}

	return false
}

// SetDoubleSided gets a reference to the given bool and assigns it to the DoubleSided field.
func (o *LetterDetailsWritable) SetDoubleSided(v bool) {
	o.DoubleSided = &v
}

// GetExtraService returns the ExtraService field value if set, zero value otherwise.
func (o *LetterDetailsWritable) GetExtraService() string {
	if o == nil || o.ExtraService == nil {
		var ret string
		return ret
	}
	return *o.ExtraService
}

// GetExtraServiceOk returns a tuple with the ExtraService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LetterDetailsWritable) GetExtraServiceOk() (*string, bool) {
	if o == nil || o.ExtraService == nil {
		return nil, false
	}
	return o.ExtraService, true
}

// HasExtraService returns a boolean if a field has been set.
func (o *LetterDetailsWritable) HasExtraService() bool {
	if o != nil && o.ExtraService != nil {
		return true
	}

	return false
}

// SetExtraService gets a reference to the given string and assigns it to the ExtraService field.
func (o *LetterDetailsWritable) SetExtraService(v string) {
	o.ExtraService = &v
}

// GetMailType returns the MailType field value if set, zero value otherwise.
func (o *LetterDetailsWritable) GetMailType() MailType {
	if o == nil || o.MailType == nil {
		var ret MailType
		return ret
	}
	return *o.MailType
}

// GetMailTypeOk returns a tuple with the MailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LetterDetailsWritable) GetMailTypeOk() (*MailType, bool) {
	if o == nil || o.MailType == nil {
		return nil, false
	}
	return o.MailType, true
}

// HasMailType returns a boolean if a field has been set.
func (o *LetterDetailsWritable) HasMailType() bool {
	if o != nil && o.MailType != nil {
		return true
	}

	return false
}

// SetMailType gets a reference to the given MailType and assigns it to the MailType field.
func (o *LetterDetailsWritable) SetMailType(v MailType) {
	o.MailType = &v
}

// GetReturnEnvelope returns the ReturnEnvelope field value if set, zero value otherwise.
func (o *LetterDetailsWritable) GetReturnEnvelope() bool {
	if o == nil || o.ReturnEnvelope == nil {
		var ret bool
		return ret
	}
	return *o.ReturnEnvelope
}

// GetReturnEnvelopeOk returns a tuple with the ReturnEnvelope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LetterDetailsWritable) GetReturnEnvelopeOk() (*bool, bool) {
	if o == nil || o.ReturnEnvelope == nil {
		return nil, false
	}
	return o.ReturnEnvelope, true
}

// HasReturnEnvelope returns a boolean if a field has been set.
func (o *LetterDetailsWritable) HasReturnEnvelope() bool {
	if o != nil && o.ReturnEnvelope != nil {
		return true
	}

	return false
}

// SetReturnEnvelope gets a reference to the given bool and assigns it to the ReturnEnvelope field.
func (o *LetterDetailsWritable) SetReturnEnvelope(v bool) {
	o.ReturnEnvelope = &v
}

// GetCustomEnvelope returns the CustomEnvelope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LetterDetailsWritable) GetCustomEnvelope() string {
	if o == nil || o.CustomEnvelope.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomEnvelope.Get()
}

// GetCustomEnvelopeOk returns a tuple with the CustomEnvelope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LetterDetailsWritable) GetCustomEnvelopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomEnvelope.Get(), o.CustomEnvelope.IsSet()
}

// HasCustomEnvelope returns a boolean if a field has been set.
func (o *LetterDetailsWritable) HasCustomEnvelope() bool {
	if o != nil && o.CustomEnvelope.IsSet() {
		return true
	}

	return false
}

// SetCustomEnvelope gets a reference to the given NullableString and assigns it to the CustomEnvelope field.
func (o *LetterDetailsWritable) SetCustomEnvelope(v string) {
	o.CustomEnvelope.Set(&v)
}
// SetCustomEnvelopeNil sets the value for CustomEnvelope to be an explicit nil
func (o *LetterDetailsWritable) SetCustomEnvelopeNil() {
	o.CustomEnvelope.Set(nil)
}

// UnsetCustomEnvelope ensures that no value is present for CustomEnvelope, not even an explicit nil
func (o *LetterDetailsWritable) UnsetCustomEnvelope() {
	o.CustomEnvelope.Unset()
}

func (o LetterDetailsWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressPlacement != nil {
		toSerialize["address_placement"] = o.AddressPlacement
	}
	if o.Cards != nil {
		toSerialize["cards"] = o.Cards
	}
	if true {
		toSerialize["color"] = o.Color
	}
	if o.DoubleSided != nil {
		toSerialize["double_sided"] = o.DoubleSided
	}
	if o.ExtraService != nil {
		toSerialize["extra_service"] = o.ExtraService
	}
	if o.MailType != nil {
		toSerialize["mail_type"] = o.MailType
	}
	if o.ReturnEnvelope != nil {
		toSerialize["return_envelope"] = o.ReturnEnvelope
	}
	if o.CustomEnvelope.IsSet() {
		toSerialize["custom_envelope"] = o.CustomEnvelope.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLetterDetailsWritable struct {
	value *LetterDetailsWritable
	isSet bool
}

func (v NullableLetterDetailsWritable) Get() *LetterDetailsWritable {
	return v.value
}

func (v *NullableLetterDetailsWritable) Set(val *LetterDetailsWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableLetterDetailsWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableLetterDetailsWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLetterDetailsWritable(val *LetterDetailsWritable) *NullableLetterDetailsWritable {
	return &NullableLetterDetailsWritable{value: val, isSet: true}
}

func (v NullableLetterDetailsWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLetterDetailsWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


