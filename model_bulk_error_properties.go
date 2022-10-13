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

// BulkErrorProperties Bulk error properties
type BulkErrorProperties struct {
	// A human-readable message with more details about the error
	Message *string `json:"message,omitempty"`
	// A conventional HTTP status code.
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewBulkErrorProperties instantiates a new BulkErrorProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkErrorProperties() *BulkErrorProperties {
	this := BulkErrorProperties{}
	return &this
}

// NewBulkErrorPropertiesWithDefaults instantiates a new BulkErrorProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkErrorPropertiesWithDefaults() *BulkErrorProperties {
	this := BulkErrorProperties{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BulkErrorProperties) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkErrorProperties) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BulkErrorProperties) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BulkErrorProperties) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *BulkErrorProperties) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkErrorProperties) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *BulkErrorProperties) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *BulkErrorProperties) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o BulkErrorProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableBulkErrorProperties struct {
	value *BulkErrorProperties
	isSet bool
}

func (v NullableBulkErrorProperties) Get() *BulkErrorProperties {
	return v.value
}

func (v *NullableBulkErrorProperties) Set(val *BulkErrorProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkErrorProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkErrorProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkErrorProperties(val *BulkErrorProperties) *NullableBulkErrorProperties {
	return &NullableBulkErrorProperties{value: val, isSet: true}
}

func (v NullableBulkErrorProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkErrorProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

