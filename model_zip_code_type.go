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

// ZipCodeType A description of the ZIP code type. For more detailed information about each ZIP code type, see [US Verification Details](#tag/US-Verification-Types). 
type ZipCodeType string

// List of zip_code_type
const (
	ZIPCODETYPE_STANDARD ZipCodeType = "standard"
	ZIPCODETYPE_PO_BOX ZipCodeType = "po_box"
	ZIPCODETYPE_UNIQUE ZipCodeType = "unique"
	ZIPCODETYPE_MILITARY ZipCodeType = "military"
	ZIPCODETYPE_EMPTY ZipCodeType = ""
)

// All allowed values of ZipCodeType enum
var AllowedZipCodeTypeEnumValues = []ZipCodeType{
	"standard",
	"po_box",
	"unique",
	"military",
	"",
}

func (v *ZipCodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ZipCodeType(value)
	for _, existing := range AllowedZipCodeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ZipCodeType", value)
}

// NewZipCodeTypeFromValue returns a pointer to a valid ZipCodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewZipCodeTypeFromValue(v string) (*ZipCodeType, error) {
	ev := ZipCodeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ZipCodeType: valid values are %v", v, AllowedZipCodeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ZipCodeType) IsValid() bool {
	for _, existing := range AllowedZipCodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to zip_code_type value
func (v ZipCodeType) Ptr() *ZipCodeType {
	return &v
}

type NullableZipCodeType struct {
	value *ZipCodeType
	isSet bool
}

func (v NullableZipCodeType) Get() *ZipCodeType {
	return v.value
}

func (v *NullableZipCodeType) Set(val *ZipCodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableZipCodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableZipCodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZipCodeType(val *ZipCodeType) *NullableZipCodeType {
	return &NullableZipCodeType{value: val, isSet: true}
}

func (v NullableZipCodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZipCodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

