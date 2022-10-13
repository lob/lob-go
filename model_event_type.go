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

// EventType struct for EventType
type EventType struct {
	Id string `json:"id"`
	// Value is `true` if the event type is enabled in both the test and live environments.
	EnabledForTest bool `json:"enabled_for_test"`
	Resource string `json:"resource"`
	// Value is resource type.
	Object string `json:"object"`
}

// NewEventType instantiates a new EventType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventType(id string, enabledForTest bool, resource string, object string) *EventType {
	this := EventType{}
	this.Id = id
	this.EnabledForTest = enabledForTest
	this.Resource = resource
	this.Object = object
	return &this
}

// NewEventTypeWithDefaults instantiates a new EventType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeWithDefaults() *EventType {
	this := EventType{}
	var object string = "event_type"
	this.Object = object
	return &this
}

// GetId returns the Id field value
func (o *EventType) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventType) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventType) SetId(v string) {
	o.Id = v
}

// GetEnabledForTest returns the EnabledForTest field value
func (o *EventType) GetEnabledForTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnabledForTest
}

// GetEnabledForTestOk returns a tuple with the EnabledForTest field value
// and a boolean to check if the value has been set.
func (o *EventType) GetEnabledForTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledForTest, true
}

// SetEnabledForTest sets field value
func (o *EventType) SetEnabledForTest(v bool) {
	o.EnabledForTest = v
}

// GetResource returns the Resource field value
func (o *EventType) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *EventType) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *EventType) SetResource(v string) {
	o.Resource = v
}

// GetObject returns the Object field value
func (o *EventType) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *EventType) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *EventType) SetObject(v string) {
	o.Object = v
}

func (o EventType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["enabled_for_test"] = o.EnabledForTest
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	if true {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableEventType struct {
	value *EventType
	isSet bool
}

func (v NullableEventType) Get() *EventType {
	return v.value
}

func (v *NullableEventType) Set(val *EventType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventType(val *EventType) *NullableEventType {
	return &NullableEventType{value: val, isSet: true}
}

func (v NullableEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

