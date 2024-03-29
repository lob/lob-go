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

// CardUpdatable struct for CardUpdatable
type CardUpdatable struct {
	// Description of the card.
	Description NullableString `json:"description,omitempty"`
	// Allows for auto reordering
	AutoReorder *bool `json:"auto_reorder,omitempty"`
	// The quantity of items to be reordered (only required when auto_reorder is true).
	ReorderQuantity *float32 `json:"reorder_quantity,omitempty"`
}

// NewCardUpdatable instantiates a new CardUpdatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardUpdatable() *CardUpdatable {
	this := CardUpdatable{}
	return &this
}

// NewCardUpdatableWithDefaults instantiates a new CardUpdatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardUpdatableWithDefaults() *CardUpdatable {
	this := CardUpdatable{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardUpdatable) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardUpdatable) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CardUpdatable) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CardUpdatable) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CardUpdatable) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CardUpdatable) UnsetDescription() {
	o.Description.Unset()
}

// GetAutoReorder returns the AutoReorder field value if set, zero value otherwise.
func (o *CardUpdatable) GetAutoReorder() bool {
	if o == nil || o.AutoReorder == nil {
		var ret bool
		return ret
	}
	return *o.AutoReorder
}

// GetAutoReorderOk returns a tuple with the AutoReorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardUpdatable) GetAutoReorderOk() (*bool, bool) {
	if o == nil || o.AutoReorder == nil {
		return nil, false
	}
	return o.AutoReorder, true
}

// HasAutoReorder returns a boolean if a field has been set.
func (o *CardUpdatable) HasAutoReorder() bool {
	if o != nil && o.AutoReorder != nil {
		return true
	}

	return false
}

// SetAutoReorder gets a reference to the given bool and assigns it to the AutoReorder field.
func (o *CardUpdatable) SetAutoReorder(v bool) {
	o.AutoReorder = &v
}

// GetReorderQuantity returns the ReorderQuantity field value if set, zero value otherwise.
func (o *CardUpdatable) GetReorderQuantity() float32 {
	if o == nil || o.ReorderQuantity == nil {
		var ret float32
		return ret
	}
	return *o.ReorderQuantity
}

// GetReorderQuantityOk returns a tuple with the ReorderQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardUpdatable) GetReorderQuantityOk() (*float32, bool) {
	if o == nil || o.ReorderQuantity == nil {
		return nil, false
	}
	return o.ReorderQuantity, true
}

// HasReorderQuantity returns a boolean if a field has been set.
func (o *CardUpdatable) HasReorderQuantity() bool {
	if o != nil && o.ReorderQuantity != nil {
		return true
	}

	return false
}

// SetReorderQuantity gets a reference to the given float32 and assigns it to the ReorderQuantity field.
func (o *CardUpdatable) SetReorderQuantity(v float32) {
	o.ReorderQuantity = &v
}

func (o CardUpdatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.AutoReorder != nil {
		toSerialize["auto_reorder"] = o.AutoReorder
	}
	if o.ReorderQuantity != nil {
		toSerialize["reorder_quantity"] = o.ReorderQuantity
	}
	return json.Marshal(toSerialize)
}

type NullableCardUpdatable struct {
	value *CardUpdatable
	isSet bool
}

func (v NullableCardUpdatable) Get() *CardUpdatable {
	return v.value
}

func (v *NullableCardUpdatable) Set(val *CardUpdatable) {
	v.value = val
	v.isSet = true
}

func (v NullableCardUpdatable) IsSet() bool {
	return v.isSet
}

func (v *NullableCardUpdatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardUpdatable(val *CardUpdatable) *NullableCardUpdatable {
	return &NullableCardUpdatable{value: val, isSet: true}
}

func (v NullableCardUpdatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardUpdatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


