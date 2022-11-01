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

// ExportModel struct for ExportModel
type ExportModel struct {
	Type *string `json:"type,omitempty"`
}

// NewExportModel instantiates a new ExportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportModel() *ExportModel {
	this := ExportModel{}
	return &this
}

// NewExportModelWithDefaults instantiates a new ExportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportModelWithDefaults() *ExportModel {
	this := ExportModel{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExportModel) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportModel) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExportModel) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExportModel) SetType(v string) {
	o.Type = &v
}

func (o ExportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableExportModel struct {
	value *ExportModel
	isSet bool
}

func (v NullableExportModel) Get() *ExportModel {
	return v.value
}

func (v *NullableExportModel) Set(val *ExportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableExportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableExportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportModel(val *ExportModel) *NullableExportModel {
	return &NullableExportModel{value: val, isSet: true}
}

func (v NullableExportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


