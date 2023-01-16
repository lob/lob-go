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

// ChkUseType TThe use type for each mailpiece. Can be one of marketing, operational, or null. Null use_type is only allowed if an account default use_type is selected in Account Settings. For more information on use_type, see our  [Help Center article](https://help.lob.com/print-and-mail/building-a-mail-strategy/managing-mail-settings/declaring-mail-use-type).
type ChkUseType string

// List of chk_use_type
const (
	CHKUSETYPE_MARKETING ChkUseType = "marketing"
	CHKUSETYPE_OPERATIONAL ChkUseType = "operational"
	CHKUSETYPE_NULL ChkUseType = "null"
)

// All allowed values of ChkUseType enum
var AllowedChkUseTypeEnumValues = []ChkUseType{
	"marketing",
	"operational",
	"null",
}

func (v *ChkUseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChkUseType(value)
	for _, existing := range AllowedChkUseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChkUseType", value)
}

// NewChkUseTypeFromValue returns a pointer to a valid ChkUseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChkUseTypeFromValue(v string) (*ChkUseType, error) {
	ev := ChkUseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChkUseType: valid values are %v", v, AllowedChkUseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChkUseType) IsValid() bool {
	for _, existing := range AllowedChkUseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to chk_use_type value
func (v ChkUseType) Ptr() *ChkUseType {
	return &v
}

type NullableChkUseType struct {
	value *ChkUseType
	isSet bool
}

func (v NullableChkUseType) Get() *ChkUseType {
	return v.value
}

func (v *NullableChkUseType) Set(val *ChkUseType) {
	v.value = val
	v.isSet = true
}

func (v NullableChkUseType) IsSet() bool {
	return v.isSet
}

func (v *NullableChkUseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChkUseType(val *ChkUseType) *NullableChkUseType {
	return &NullableChkUseType{value: val, isSet: true}
}

func (v NullableChkUseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChkUseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

