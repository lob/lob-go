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

// CmpUseType The usage type of this campaign. Can be one of `marketing` or `operational`.
type CmpUseType string

// List of cmp_use_type
const (
	MARKETING CmpUseType = "marketing"
	OPERATIONAL CmpUseType = "operational"
	NULL CmpUseType = "null"
)

// All allowed values of CmpUseType enum
var AllowedCmpUseTypeEnumValues = []CmpUseType{
	"marketing",
	"operational",
	"null",
}

func (v *CmpUseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CmpUseType(value)
	for _, existing := range AllowedCmpUseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CmpUseType", value)
}

// NewCmpUseTypeFromValue returns a pointer to a valid CmpUseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCmpUseTypeFromValue(v string) (*CmpUseType, error) {
	ev := CmpUseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CmpUseType: valid values are %v", v, AllowedCmpUseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CmpUseType) IsValid() bool {
	for _, existing := range AllowedCmpUseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cmp_use_type value
func (v CmpUseType) Ptr() *CmpUseType {
	return &v
}

type NullableCmpUseType struct {
	value *CmpUseType
	isSet bool
}

func (v NullableCmpUseType) Get() *CmpUseType {
	return v.value
}

func (v *NullableCmpUseType) Set(val *CmpUseType) {
	v.value = val
	v.isSet = true
}

func (v NullableCmpUseType) IsSet() bool {
	return v.isSet
}

func (v *NullableCmpUseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmpUseType(val *CmpUseType) *NullableCmpUseType {
	return &NullableCmpUseType{value: val, isSet: true}
}

func (v NullableCmpUseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmpUseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

