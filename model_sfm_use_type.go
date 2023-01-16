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
	
	"fmt"
)

// SfmUseType The use type for each mailpiece. Can be one of marketing, operational, or null. Null use_type is only allowed if an account default use_type is selected in Account Settings. For more information on use_type, see our  [Help Center article](https://help.lob.com/print-and-mail/building-a-mail-strategy/managing-mail-settings/declaring-mail-use-type).
type SfmUseType string

// List of sfm_use_type
const (
	SFMUSETYPE_MARKETING SfmUseType = "marketing"
	SFMUSETYPE_OPERATIONAL SfmUseType = "operational"
	SFMUSETYPE_NULL SfmUseType = "null"
)

// All allowed values of SfmUseType enum
var AllowedSfmUseTypeEnumValues = []SfmUseType{
	"marketing",
	"operational",
	"null",
}

func (v *SfmUseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SfmUseType(value)
	for _, existing := range AllowedSfmUseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SfmUseType", value)
}

// NewSfmUseTypeFromValue returns a pointer to a valid SfmUseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSfmUseTypeFromValue(v string) (*SfmUseType, error) {
	ev := SfmUseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SfmUseType: valid values are %v", v, AllowedSfmUseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SfmUseType) IsValid() bool {
	for _, existing := range AllowedSfmUseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sfm_use_type value
func (v SfmUseType) Ptr() *SfmUseType {
	return &v
}

type NullableSfmUseType struct {
	value *SfmUseType
	isSet bool
}

func (v NullableSfmUseType) Get() *SfmUseType {
	return v.value
}

func (v *NullableSfmUseType) Set(val *SfmUseType) {
	v.value = val
	v.isSet = true
}

func (v NullableSfmUseType) IsSet() bool {
	return v.isSet
}

func (v *NullableSfmUseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSfmUseType(val *SfmUseType) *NullableSfmUseType {
	return &NullableSfmUseType{value: val, isSet: true}
}

func (v NullableSfmUseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSfmUseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

