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
	"strings"
)

// AddressList struct for AddressList
type AddressList struct {
	// list of addresses
	Data []Address `json:"data,omitempty"`
	// Value is type of resource.
	Object *string `json:"object,omitempty"`
	// url of next page of items in list.
	NextUrl NullableString `json:"next_url,omitempty"`
	// url of previous page of items in list.
	PreviousUrl NullableString `json:"previous_url,omitempty"`
	// number of resources in a set
	Count *int32 `json:"count,omitempty"`
	// indicates the total number of records. Provided when the request specifies an \"include\" query parameter
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewAddressList instantiates a new AddressList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressList() *AddressList {
	this := AddressList{}
	return &this
}

// NewAddressListWithDefaults instantiates a new AddressList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressListWithDefaults() *AddressList {
	this := AddressList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddressList) GetData() []Address {
	if o == nil || o.Data == nil {
		var ret []Address
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressList) GetDataOk() ([]Address, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddressList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Address and assigns it to the Data field.
func (o *AddressList) SetData(v []Address) {
	o.Data = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *AddressList) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}


func (o *AddressList) GetNextPageToken() string {
    if *o.NextUrl.value == "" {
        return "";
    }

	queryPartitionArray := strings.Split(*o.NextUrl.value, "?")

    if len(queryPartitionArray) < 2 {
        return "";
    }

    paramPartitionArray := strings.Split(queryPartitionArray[1], ("&"))

	var beforeParamString string
	
	for _, partition := range paramPartitionArray {
		if strings.Contains(partition, "after=") {
			beforeParamString = partition
			break
		}
	}

    if (beforeParamString == "") {
        return "";
    }
    return strings.Split(beforeParamString,"after=")[1];

  }

  func (o *AddressList) GetPrevPageToken() string {
	if *o.PreviousUrl.value == "" {
        return "";
    }

	queryPartitionArray := strings.Split(*o.PreviousUrl.value, "?")

    if len(queryPartitionArray) < 2 {
        return "";
    }

    paramPartitionArray := strings.Split(queryPartitionArray[1], ("&"))

	var afterParamString string
	
	for _, partition := range paramPartitionArray {
		if strings.Contains(partition, "before=") {
			afterParamString = partition
			break
		}
	}

    if (afterParamString == "") {
        return "";
    }
    return strings.Split(afterParamString,"before=")[1];

  }

  
// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressList) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *AddressList) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *AddressList) SetObject(v string) {
	o.Object = &v
}

// GetNextUrl returns the NextUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressList) GetNextUrl() string {
	if o == nil || o.NextUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextUrl.Get()
}

// GetNextUrlOk returns a tuple with the NextUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressList) GetNextUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextUrl.Get(), o.NextUrl.IsSet()
}

// HasNextUrl returns a boolean if a field has been set.
func (o *AddressList) HasNextUrl() bool {
	if o != nil && o.NextUrl.IsSet() {
		return true
	}

	return false
}

// SetNextUrl gets a reference to the given NullableString and assigns it to the NextUrl field.
func (o *AddressList) SetNextUrl(v string) {
	o.NextUrl.Set(&v)
}
// SetNextUrlNil sets the value for NextUrl to be an explicit nil
func (o *AddressList) SetNextUrlNil() {
	o.NextUrl.Set(nil)
}

// UnsetNextUrl ensures that no value is present for NextUrl, not even an explicit nil
func (o *AddressList) UnsetNextUrl() {
	o.NextUrl.Unset()
}

// GetPreviousUrl returns the PreviousUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressList) GetPreviousUrl() string {
	if o == nil || o.PreviousUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviousUrl.Get()
}

// GetPreviousUrlOk returns a tuple with the PreviousUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressList) GetPreviousUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousUrl.Get(), o.PreviousUrl.IsSet()
}

// HasPreviousUrl returns a boolean if a field has been set.
func (o *AddressList) HasPreviousUrl() bool {
	if o != nil && o.PreviousUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviousUrl gets a reference to the given NullableString and assigns it to the PreviousUrl field.
func (o *AddressList) SetPreviousUrl(v string) {
	o.PreviousUrl.Set(&v)
}
// SetPreviousUrlNil sets the value for PreviousUrl to be an explicit nil
func (o *AddressList) SetPreviousUrlNil() {
	o.PreviousUrl.Set(nil)
}

// UnsetPreviousUrl ensures that no value is present for PreviousUrl, not even an explicit nil
func (o *AddressList) UnsetPreviousUrl() {
	o.PreviousUrl.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AddressList) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressList) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AddressList) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AddressList) SetCount(v int32) {
	o.Count = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AddressList) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressList) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AddressList) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AddressList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AddressList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.NextUrl.IsSet() {
		toSerialize["next_url"] = o.NextUrl.Get()
	}
	if o.PreviousUrl.IsSet() {
		toSerialize["previous_url"] = o.PreviousUrl.Get()
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableAddressList struct {
	value *AddressList
	isSet bool
}

func (v NullableAddressList) Get() *AddressList {
	return v.value
}

func (v *NullableAddressList) Set(val *AddressList) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressList) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressList(val *AddressList) *NullableAddressList {
	return &NullableAddressList{value: val, isSet: true}
}

func (v NullableAddressList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


