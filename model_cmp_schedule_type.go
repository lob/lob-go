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

// CmpScheduleType How the campaign should be scheduled. Only value available today is `immediate`.
type CmpScheduleType string

// List of cmp_schedule_type
const (
	IMMEDIATE CmpScheduleType = "immediate"
)

// All allowed values of CmpScheduleType enum
var AllowedCmpScheduleTypeEnumValues = []CmpScheduleType{
	"immediate",
}

func (v *CmpScheduleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CmpScheduleType(value)
	for _, existing := range AllowedCmpScheduleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CmpScheduleType", value)
}

// NewCmpScheduleTypeFromValue returns a pointer to a valid CmpScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCmpScheduleTypeFromValue(v string) (*CmpScheduleType, error) {
	ev := CmpScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CmpScheduleType: valid values are %v", v, AllowedCmpScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CmpScheduleType) IsValid() bool {
	for _, existing := range AllowedCmpScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cmp_schedule_type value
func (v CmpScheduleType) Ptr() *CmpScheduleType {
	return &v
}

type NullableCmpScheduleType struct {
	value *CmpScheduleType
	isSet bool
}

func (v NullableCmpScheduleType) Get() *CmpScheduleType {
	return v.value
}

func (v *NullableCmpScheduleType) Set(val *CmpScheduleType) {
	v.value = val
	v.isSet = true
}

func (v NullableCmpScheduleType) IsSet() bool {
	return v.isSet
}

func (v *NullableCmpScheduleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmpScheduleType(val *CmpScheduleType) *NullableCmpScheduleType {
	return &NullableCmpScheduleType{value: val, isSet: true}
}

func (v NullableCmpScheduleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmpScheduleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

