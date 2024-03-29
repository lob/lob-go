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

// BankAccountVerify struct for BankAccountVerify
type BankAccountVerify struct {
	// In live mode, an array containing the two micro deposits (in cents) placed in the bank account. In test mode, no micro deposits will be placed, so any two integers between `1` and `100` will work.
	Amounts []int32 `json:"amounts"`
}

// NewBankAccountVerify instantiates a new BankAccountVerify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountVerify(amounts []int32) *BankAccountVerify {
	this := BankAccountVerify{}
	this.Amounts = amounts
	return &this
}

// NewBankAccountVerifyWithDefaults instantiates a new BankAccountVerify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountVerifyWithDefaults() *BankAccountVerify {
	this := BankAccountVerify{}
	return &this
}

// GetAmounts returns the Amounts field value
func (o *BankAccountVerify) GetAmounts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Amounts
}

// GetAmountsOk returns a tuple with the Amounts field value
// and a boolean to check if the value has been set.
func (o *BankAccountVerify) GetAmountsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amounts, true
}

// SetAmounts sets field value
func (o *BankAccountVerify) SetAmounts(v []int32) {
	o.Amounts = v
}

func (o BankAccountVerify) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amounts"] = o.Amounts
	}
	return json.Marshal(toSerialize)
}

type NullableBankAccountVerify struct {
	value *BankAccountVerify
	isSet bool
}

func (v NullableBankAccountVerify) Get() *BankAccountVerify {
	return v.value
}

func (v *NullableBankAccountVerify) Set(val *BankAccountVerify) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountVerify) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountVerify(val *BankAccountVerify) *NullableBankAccountVerify {
	return &NullableBankAccountVerify{value: val, isSet: true}
}

func (v NullableBankAccountVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


