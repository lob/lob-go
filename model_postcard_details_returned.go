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
)

// PostcardDetailsReturned Properties that the postcards in your Creative should have.
type PostcardDetailsReturned struct {
	MailType *MailType `json:"mail_type,omitempty"`
	Size *PostcardSize `json:"size,omitempty"`
	Setting *int32 `json:"setting,omitempty"`
	// The original URL of the front template.
	FrontOriginalUrl *string `json:"front_original_url,omitempty"`
	// The original URL of the back template.
	BackOriginalUrl *string `json:"back_original_url,omitempty"`
}

// NewPostcardDetailsReturned instantiates a new PostcardDetailsReturned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostcardDetailsReturned() *PostcardDetailsReturned {
	this := PostcardDetailsReturned{}
	var mailType MailType = FIRST_CLASS
	this.MailType = &mailType
	var size PostcardSize = _4X6
	this.Size = &size
	var setting int32 = 1001
	this.Setting = &setting
	return &this
}

// NewPostcardDetailsReturnedWithDefaults instantiates a new PostcardDetailsReturned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostcardDetailsReturnedWithDefaults() *PostcardDetailsReturned {
	this := PostcardDetailsReturned{}
	var mailType MailType = FIRST_CLASS
	this.MailType = &mailType
	var size PostcardSize = _4X6
	this.Size = &size
	var setting int32 = 1001
	this.Setting = &setting
	return &this
}

// GetMailType returns the MailType field value if set, zero value otherwise.
func (o *PostcardDetailsReturned) GetMailType() MailType {
	if o == nil || o.MailType == nil {
		var ret MailType
		return ret
	}
	return *o.MailType
}

// GetMailTypeOk returns a tuple with the MailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostcardDetailsReturned) GetMailTypeOk() (*MailType, bool) {
	if o == nil || o.MailType == nil {
		return nil, false
	}
	return o.MailType, true
}

// HasMailType returns a boolean if a field has been set.
func (o *PostcardDetailsReturned) HasMailType() bool {
	if o != nil && o.MailType != nil {
		return true
	}

	return false
}

// SetMailType gets a reference to the given MailType and assigns it to the MailType field.
func (o *PostcardDetailsReturned) SetMailType(v MailType) {
	o.MailType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PostcardDetailsReturned) GetSize() PostcardSize {
	if o == nil || o.Size == nil {
		var ret PostcardSize
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostcardDetailsReturned) GetSizeOk() (*PostcardSize, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PostcardDetailsReturned) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given PostcardSize and assigns it to the Size field.
func (o *PostcardDetailsReturned) SetSize(v PostcardSize) {
	o.Size = &v
}

// GetSetting returns the Setting field value if set, zero value otherwise.
func (o *PostcardDetailsReturned) GetSetting() int32 {
	if o == nil || o.Setting == nil {
		var ret int32
		return ret
	}
	return *o.Setting
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostcardDetailsReturned) GetSettingOk() (*int32, bool) {
	if o == nil || o.Setting == nil {
		return nil, false
	}
	return o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *PostcardDetailsReturned) HasSetting() bool {
	if o != nil && o.Setting != nil {
		return true
	}

	return false
}

// SetSetting gets a reference to the given int32 and assigns it to the Setting field.
func (o *PostcardDetailsReturned) SetSetting(v int32) {
	o.Setting = &v
}

// GetFrontOriginalUrl returns the FrontOriginalUrl field value if set, zero value otherwise.
func (o *PostcardDetailsReturned) GetFrontOriginalUrl() string {
	if o == nil || o.FrontOriginalUrl == nil {
		var ret string
		return ret
	}
	return *o.FrontOriginalUrl
}

// GetFrontOriginalUrlOk returns a tuple with the FrontOriginalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostcardDetailsReturned) GetFrontOriginalUrlOk() (*string, bool) {
	if o == nil || o.FrontOriginalUrl == nil {
		return nil, false
	}
	return o.FrontOriginalUrl, true
}

// HasFrontOriginalUrl returns a boolean if a field has been set.
func (o *PostcardDetailsReturned) HasFrontOriginalUrl() bool {
	if o != nil && o.FrontOriginalUrl != nil {
		return true
	}

	return false
}

// SetFrontOriginalUrl gets a reference to the given string and assigns it to the FrontOriginalUrl field.
func (o *PostcardDetailsReturned) SetFrontOriginalUrl(v string) {
	o.FrontOriginalUrl = &v
}

// GetBackOriginalUrl returns the BackOriginalUrl field value if set, zero value otherwise.
func (o *PostcardDetailsReturned) GetBackOriginalUrl() string {
	if o == nil || o.BackOriginalUrl == nil {
		var ret string
		return ret
	}
	return *o.BackOriginalUrl
}

// GetBackOriginalUrlOk returns a tuple with the BackOriginalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostcardDetailsReturned) GetBackOriginalUrlOk() (*string, bool) {
	if o == nil || o.BackOriginalUrl == nil {
		return nil, false
	}
	return o.BackOriginalUrl, true
}

// HasBackOriginalUrl returns a boolean if a field has been set.
func (o *PostcardDetailsReturned) HasBackOriginalUrl() bool {
	if o != nil && o.BackOriginalUrl != nil {
		return true
	}

	return false
}

// SetBackOriginalUrl gets a reference to the given string and assigns it to the BackOriginalUrl field.
func (o *PostcardDetailsReturned) SetBackOriginalUrl(v string) {
	o.BackOriginalUrl = &v
}

func (o PostcardDetailsReturned) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MailType != nil {
		toSerialize["mail_type"] = o.MailType
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Setting != nil {
		toSerialize["setting"] = o.Setting
	}
	if o.FrontOriginalUrl != nil {
		toSerialize["front_original_url"] = o.FrontOriginalUrl
	}
	if o.BackOriginalUrl != nil {
		toSerialize["back_original_url"] = o.BackOriginalUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePostcardDetailsReturned struct {
	value *PostcardDetailsReturned
	isSet bool
}

func (v NullablePostcardDetailsReturned) Get() *PostcardDetailsReturned {
	return v.value
}

func (v *NullablePostcardDetailsReturned) Set(val *PostcardDetailsReturned) {
	v.value = val
	v.isSet = true
}

func (v NullablePostcardDetailsReturned) IsSet() bool {
	return v.isSet
}

func (v *NullablePostcardDetailsReturned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostcardDetailsReturned(val *PostcardDetailsReturned) *NullablePostcardDetailsReturned {
	return &NullablePostcardDetailsReturned{value: val, isSet: true}
}

func (v NullablePostcardDetailsReturned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostcardDetailsReturned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


