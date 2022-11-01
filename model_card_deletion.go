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

// CardDeletion Lob uses RESTful HTTP response codes to indicate success or failure of an API request. In general, 2xx indicates success, 4xx indicate an input error, and 5xx indicates an error on Lob's end.
type CardDeletion struct {
	// Unique identifier prefixed with `card_`.
	Id *string `json:"id,omitempty"`
	// Only returned if the resource has been successfully deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// Value is type of resource.
	Object *string `json:"object,omitempty"`
}

// NewCardDeletion instantiates a new CardDeletion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardDeletion() *CardDeletion {
	this := CardDeletion{}
	var object string = "card_deleted"
	this.Object = &object
	return &this
}

// NewCardDeletionWithDefaults instantiates a new CardDeletion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardDeletionWithDefaults() *CardDeletion {
	this := CardDeletion{}
	var object string = "card_deleted"
	this.Object = &object
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CardDeletion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDeletion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CardDeletion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CardDeletion) SetId(v string) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CardDeletion) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDeletion) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CardDeletion) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *CardDeletion) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CardDeletion) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDeletion) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CardDeletion) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CardDeletion) SetObject(v string) {
	o.Object = &v
}

func (o CardDeletion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableCardDeletion struct {
	value *CardDeletion
	isSet bool
}

func (v NullableCardDeletion) Get() *CardDeletion {
	return v.value
}

func (v *NullableCardDeletion) Set(val *CardDeletion) {
	v.value = val
	v.isSet = true
}

func (v NullableCardDeletion) IsSet() bool {
	return v.isSet
}

func (v *NullableCardDeletion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardDeletion(val *CardDeletion) *NullableCardDeletion {
	return &NullableCardDeletion{value: val, isSet: true}
}

func (v NullableCardDeletion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardDeletion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


