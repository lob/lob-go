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

// IdentityValidation struct for IdentityValidation
type IdentityValidation struct {
	// Unique identifier prefixed with `id_validation_`.
	Id *string `json:"id,omitempty"`
	// The intended recipient, typically a person's or firm's name.
	Recipient NullableString `json:"recipient,omitempty"`
	// The primary delivery line (usually the street address) of the address. Combination of the following applicable `components`: * `primary_number` * `street_predirection` * `street_name` * `street_suffix` * `street_postdirection` * `secondary_designator` * `secondary_number` * `pmb_designator` * `pmb_number` 
	PrimaryLine *string `json:"primary_line,omitempty"`
	// The secondary delivery line of the address. This field is typically empty but may contain information if `primary_line` is too long. 
	SecondaryLine *string `json:"secondary_line,omitempty"`
	// Only present for addresses in Puerto Rico. An urbanization refers to an area, sector, or development within a city. See [USPS documentation](https://pe.usps.com/text/pub28/28api_008.htm#:~:text=I51.,-4%20Urbanizations&text=In%20Puerto%20Rico%2C%20identical%20street,placed%20before%20the%20urbanization%20name.) for clarification. 
	Urbanization *string `json:"urbanization,omitempty"`
	// Combination of the following applicable `components`: * City (`city`) * State (`state`) * ZIP code (`zip_code`) * ZIP+4 (`zip_code_plus_4`) 
	LastLine *string `json:"last_line,omitempty"`
	// A numerical score between 0 and 100 that represents the likelihood the provided name is associated with a physical address. 
	Score NullableFloat32 `json:"score,omitempty"`
	// Indicates the likelihood the recipient name and address match based on our custom internal calculation. Possible values are: - `high` — Has a Lob confidence score greater than 70. - `medium` — Has a Lob confidence score between 40 and 70. - `low` — Has a Lob confidence score less than 40. - `\"\"` — No tracking data exists for this address. 
	Confidence *string `json:"confidence,omitempty"`
	// Value is resource type.
	Object *string `json:"object,omitempty"`
}

// NewIdentityValidation instantiates a new IdentityValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityValidation() *IdentityValidation {
	this := IdentityValidation{}
	var object string = "id_validation"
	this.Object = &object
	return &this
}

// NewIdentityValidationWithDefaults instantiates a new IdentityValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityValidationWithDefaults() *IdentityValidation {
	this := IdentityValidation{}
	var object string = "id_validation"
	this.Object = &object
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityValidation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityValidation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityValidation) SetId(v string) {
	o.Id = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityValidation) GetRecipient() string {
	if o == nil || o.Recipient.Get() == nil {
		var ret string
		return ret
	}
	return *o.Recipient.Get()
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityValidation) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipient.Get(), o.Recipient.IsSet()
}

// HasRecipient returns a boolean if a field has been set.
func (o *IdentityValidation) HasRecipient() bool {
	if o != nil && o.Recipient.IsSet() {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given NullableString and assigns it to the Recipient field.
func (o *IdentityValidation) SetRecipient(v string) {
	o.Recipient.Set(&v)
}
// SetRecipientNil sets the value for Recipient to be an explicit nil
func (o *IdentityValidation) SetRecipientNil() {
	o.Recipient.Set(nil)
}

// UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
func (o *IdentityValidation) UnsetRecipient() {
	o.Recipient.Unset()
}

// GetPrimaryLine returns the PrimaryLine field value if set, zero value otherwise.
func (o *IdentityValidation) GetPrimaryLine() string {
	if o == nil || o.PrimaryLine == nil {
		var ret string
		return ret
	}
	return *o.PrimaryLine
}

// GetPrimaryLineOk returns a tuple with the PrimaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidation) GetPrimaryLineOk() (*string, bool) {
	if o == nil || o.PrimaryLine == nil {
		return nil, false
	}
	return o.PrimaryLine, true
}

// HasPrimaryLine returns a boolean if a field has been set.
func (o *IdentityValidation) HasPrimaryLine() bool {
	if o != nil && o.PrimaryLine != nil {
		return true
	}

	return false
}

// SetPrimaryLine gets a reference to the given string and assigns it to the PrimaryLine field.
func (o *IdentityValidation) SetPrimaryLine(v string) {
	o.PrimaryLine = &v
}

// GetSecondaryLine returns the SecondaryLine field value if set, zero value otherwise.
func (o *IdentityValidation) GetSecondaryLine() string {
	if o == nil || o.SecondaryLine == nil {
		var ret string
		return ret
	}
	return *o.SecondaryLine
}

// GetSecondaryLineOk returns a tuple with the SecondaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidation) GetSecondaryLineOk() (*string, bool) {
	if o == nil || o.SecondaryLine == nil {
		return nil, false
	}
	return o.SecondaryLine, true
}

// HasSecondaryLine returns a boolean if a field has been set.
func (o *IdentityValidation) HasSecondaryLine() bool {
	if o != nil && o.SecondaryLine != nil {
		return true
	}

	return false
}

// SetSecondaryLine gets a reference to the given string and assigns it to the SecondaryLine field.
func (o *IdentityValidation) SetSecondaryLine(v string) {
	o.SecondaryLine = &v
}

// GetUrbanization returns the Urbanization field value if set, zero value otherwise.
func (o *IdentityValidation) GetUrbanization() string {
	if o == nil || o.Urbanization == nil {
		var ret string
		return ret
	}
	return *o.Urbanization
}

// GetUrbanizationOk returns a tuple with the Urbanization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidation) GetUrbanizationOk() (*string, bool) {
	if o == nil || o.Urbanization == nil {
		return nil, false
	}
	return o.Urbanization, true
}

// HasUrbanization returns a boolean if a field has been set.
func (o *IdentityValidation) HasUrbanization() bool {
	if o != nil && o.Urbanization != nil {
		return true
	}

	return false
}

// SetUrbanization gets a reference to the given string and assigns it to the Urbanization field.
func (o *IdentityValidation) SetUrbanization(v string) {
	o.Urbanization = &v
}

// GetLastLine returns the LastLine field value if set, zero value otherwise.
func (o *IdentityValidation) GetLastLine() string {
	if o == nil || o.LastLine == nil {
		var ret string
		return ret
	}
	return *o.LastLine
}

// GetLastLineOk returns a tuple with the LastLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidation) GetLastLineOk() (*string, bool) {
	if o == nil || o.LastLine == nil {
		return nil, false
	}
	return o.LastLine, true
}

// HasLastLine returns a boolean if a field has been set.
func (o *IdentityValidation) HasLastLine() bool {
	if o != nil && o.LastLine != nil {
		return true
	}

	return false
}

// SetLastLine gets a reference to the given string and assigns it to the LastLine field.
func (o *IdentityValidation) SetLastLine(v string) {
	o.LastLine = &v
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityValidation) GetScore() float32 {
	if o == nil || o.Score.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityValidation) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *IdentityValidation) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableFloat32 and assigns it to the Score field.
func (o *IdentityValidation) SetScore(v float32) {
	o.Score.Set(&v)
}
// SetScoreNil sets the value for Score to be an explicit nil
func (o *IdentityValidation) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *IdentityValidation) UnsetScore() {
	o.Score.Unset()
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *IdentityValidation) GetConfidence() string {
	if o == nil || o.Confidence == nil {
		var ret string
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidation) GetConfidenceOk() (*string, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *IdentityValidation) HasConfidence() bool {
	if o != nil && o.Confidence != nil {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given string and assigns it to the Confidence field.
func (o *IdentityValidation) SetConfidence(v string) {
	o.Confidence = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *IdentityValidation) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidation) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *IdentityValidation) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *IdentityValidation) SetObject(v string) {
	o.Object = &v
}

func (o IdentityValidation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.LastLine != nil {
		toSerialize["last_line"] = o.LastLine
	}
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityValidation struct {
	value *IdentityValidation
	isSet bool
}

func (v NullableIdentityValidation) Get() *IdentityValidation {
	return v.value
}

func (v *NullableIdentityValidation) Set(val *IdentityValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityValidation(val *IdentityValidation) *NullableIdentityValidation {
	return &NullableIdentityValidation{value: val, isSet: true}
}

func (v NullableIdentityValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


