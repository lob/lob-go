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
	"fmt"
)

// SelfMailerSize Specifies the size of the self mailer.
type SelfMailerSize string

// List of self_mailer_size
const (
	_6X18_BIFOLD SelfMailerSize = "6x18_bifold"
	_11X9_BIFOLD SelfMailerSize = "11x9_bifold"
	_12X9_BIFOLD SelfMailerSize = "12x9_bifold"
)

// All allowed values of SelfMailerSize enum
var AllowedSelfMailerSizeEnumValues = []SelfMailerSize{
	"6x18_bifold",
	"11x9_bifold",
	"12x9_bifold",
}

func (v *SelfMailerSize) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelfMailerSize(value)
	for _, existing := range AllowedSelfMailerSizeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelfMailerSize", value)
}

// NewSelfMailerSizeFromValue returns a pointer to a valid SelfMailerSize
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelfMailerSizeFromValue(v string) (*SelfMailerSize, error) {
	ev := SelfMailerSize(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SelfMailerSize: valid values are %v", v, AllowedSelfMailerSizeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelfMailerSize) IsValid() bool {
	for _, existing := range AllowedSelfMailerSizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to self_mailer_size value
func (v SelfMailerSize) Ptr() *SelfMailerSize {
	return &v
}

type NullableSelfMailerSize struct {
	value *SelfMailerSize
	isSet bool
}

func (v NullableSelfMailerSize) Get() *SelfMailerSize {
	return v.value
}

func (v *NullableSelfMailerSize) Set(val *SelfMailerSize) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfMailerSize) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfMailerSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfMailerSize(val *SelfMailerSize) *NullableSelfMailerSize {
	return &NullableSelfMailerSize{value: val, isSet: true}
}

func (v NullableSelfMailerSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfMailerSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

