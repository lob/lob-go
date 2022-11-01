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

// UploadFile struct for UploadFile
type UploadFile struct {
	Message string `json:"message"`
	Filename string `json:"filename"`
}

// NewUploadFile instantiates a new UploadFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadFile(message string, filename string) *UploadFile {
	this := UploadFile{}
	this.Message = message
	this.Filename = filename
	return &this
}

// NewUploadFileWithDefaults instantiates a new UploadFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadFileWithDefaults() *UploadFile {
	this := UploadFile{}
	var message string = "File uploaded successfully"
	this.Message = message
	return &this
}

// GetMessage returns the Message field value
func (o *UploadFile) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UploadFile) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UploadFile) SetMessage(v string) {
	o.Message = v
}

// GetFilename returns the Filename field value
func (o *UploadFile) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *UploadFile) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *UploadFile) SetFilename(v string) {
	o.Filename = v
}

func (o UploadFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["filename"] = o.Filename
	}
	return json.Marshal(toSerialize)
}

type NullableUploadFile struct {
	value *UploadFile
	isSet bool
}

func (v NullableUploadFile) Get() *UploadFile {
	return v.value
}

func (v *NullableUploadFile) Set(val *UploadFile) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadFile) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadFile(val *UploadFile) *NullableUploadFile {
	return &NullableUploadFile{value: val, isSet: true}
}

func (v NullableUploadFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


