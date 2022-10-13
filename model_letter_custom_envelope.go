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

// LetterCustomEnvelope A nested custom envelope object containing more information about the custom envelope used or `null` if a custom envelope was not used.  Accepts an envelope ID for any customized envelope with available inventory. If no inventory is available for the specified ID, the letter will not be sent, and an error will be returned. If the letter has more than 6 sheets, it will be sent in a blank flat envelope. Custom envelopes may be created and ordered from the dashboard. This feature is exclusive to certain customers. Upgrade to the appropriate [Print & Mail Edition](https://dashboard.lob.com/#/settings/editions) to gain access.
type LetterCustomEnvelope struct {
	// The unique identifier of the custom envelope used.
	Id *string `json:"id,omitempty"`
	// The url of the envelope asset used.
	Url *string `json:"url,omitempty"`
	Object *string `json:"object,omitempty"`
}

// NewLetterCustomEnvelope instantiates a new LetterCustomEnvelope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLetterCustomEnvelope() *LetterCustomEnvelope {
	this := LetterCustomEnvelope{}
	var object string = "envelope"
	this.Object = &object
	return &this
}

// NewLetterCustomEnvelopeWithDefaults instantiates a new LetterCustomEnvelope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLetterCustomEnvelopeWithDefaults() *LetterCustomEnvelope {
	this := LetterCustomEnvelope{}
	var object string = "envelope"
	this.Object = &object
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LetterCustomEnvelope) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LetterCustomEnvelope) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LetterCustomEnvelope) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LetterCustomEnvelope) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LetterCustomEnvelope) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LetterCustomEnvelope) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LetterCustomEnvelope) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LetterCustomEnvelope) SetUrl(v string) {
	o.Url = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *LetterCustomEnvelope) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LetterCustomEnvelope) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *LetterCustomEnvelope) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *LetterCustomEnvelope) SetObject(v string) {
	o.Object = &v
}

func (o LetterCustomEnvelope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableLetterCustomEnvelope struct {
	value *LetterCustomEnvelope
	isSet bool
}

func (v NullableLetterCustomEnvelope) Get() *LetterCustomEnvelope {
	return v.value
}

func (v *NullableLetterCustomEnvelope) Set(val *LetterCustomEnvelope) {
	v.value = val
	v.isSet = true
}

func (v NullableLetterCustomEnvelope) IsSet() bool {
	return v.isSet
}

func (v *NullableLetterCustomEnvelope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLetterCustomEnvelope(val *LetterCustomEnvelope) *NullableLetterCustomEnvelope {
	return &NullableLetterCustomEnvelope{value: val, isSet: true}
}

func (v NullableLetterCustomEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLetterCustomEnvelope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

