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
	
	"time"
)

// BuckslipOrder struct for BuckslipOrder
type BuckslipOrder struct {
	// A timestamp in ISO 8601 format of the date the resource was created.
	DateCreated time.Time `json:"date_created"`
	// A timestamp in ISO 8601 format of the date the resource was last modified.
	DateModified time.Time `json:"date_modified"`
	// Only returned if the resource has been successfully deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// Value is type of resource.
	Object string `json:"object"`
	// Unique identifier prefixed with `bo_`.
	Id *string `json:"id,omitempty"`
	// Unique identifier prefixed with `bck_`.
	BuckslipId *string `json:"buckslip_id,omitempty"`
	// The status of the buckslip order.
	Status *string `json:"status,omitempty"`
	// The quantity of buckslips ordered.
	QuantityOrdered *float32 `json:"quantity_ordered,omitempty"`
	// The unit price for the buckslip order.
	UnitPrice *float32 `json:"unit_price,omitempty"`
	// The inventory of the buckslip order.
	Inventory *float32 `json:"inventory,omitempty"`
	// The reason for cancellation.
	CancelledReason *string `json:"cancelled_reason,omitempty"`
	// A timestamp in ISO 8601 format of the date the resource was created.
	AvailabilityDate *time.Time `json:"availability_date,omitempty"`
	// The fixed deadline for the buckslips to be printed.
	ExpectedAvailabilityDate *time.Time `json:"expected_availability_date,omitempty"`
}

// NewBuckslipOrder instantiates a new BuckslipOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuckslipOrder(dateCreated time.Time, dateModified time.Time, object string) *BuckslipOrder {
	this := BuckslipOrder{}
	this.DateCreated = dateCreated
	this.DateModified = dateModified
	this.Object = object
	var quantityOrdered float32 = 0
	this.QuantityOrdered = &quantityOrdered
	var unitPrice float32 = 0
	this.UnitPrice = &unitPrice
	var inventory float32 = 0
	this.Inventory = &inventory
	return &this
}

// NewBuckslipOrderWithDefaults instantiates a new BuckslipOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuckslipOrderWithDefaults() *BuckslipOrder {
	this := BuckslipOrder{}
	var quantityOrdered float32 = 0
	this.QuantityOrdered = &quantityOrdered
	var unitPrice float32 = 0
	this.UnitPrice = &unitPrice
	var inventory float32 = 0
	this.Inventory = &inventory
	return &this
}

// GetDateCreated returns the DateCreated field value
func (o *BuckslipOrder) GetDateCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *BuckslipOrder) SetDateCreated(v time.Time) {
	o.DateCreated = v
}

// GetDateModified returns the DateModified field value
func (o *BuckslipOrder) GetDateModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetDateModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateModified, true
}

// SetDateModified sets field value
func (o *BuckslipOrder) SetDateModified(v time.Time) {
	o.DateModified = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *BuckslipOrder) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *BuckslipOrder) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *BuckslipOrder) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetObject returns the Object field value
func (o *BuckslipOrder) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *BuckslipOrder) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BuckslipOrder) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BuckslipOrder) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BuckslipOrder) SetId(v string) {
	o.Id = &v
}

// GetBuckslipId returns the BuckslipId field value if set, zero value otherwise.
func (o *BuckslipOrder) GetBuckslipId() string {
	if o == nil || o.BuckslipId == nil {
		var ret string
		return ret
	}
	return *o.BuckslipId
}

// GetBuckslipIdOk returns a tuple with the BuckslipId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetBuckslipIdOk() (*string, bool) {
	if o == nil || o.BuckslipId == nil {
		return nil, false
	}
	return o.BuckslipId, true
}

// HasBuckslipId returns a boolean if a field has been set.
func (o *BuckslipOrder) HasBuckslipId() bool {
	if o != nil && o.BuckslipId != nil {
		return true
	}

	return false
}

// SetBuckslipId gets a reference to the given string and assigns it to the BuckslipId field.
func (o *BuckslipOrder) SetBuckslipId(v string) {
	o.BuckslipId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BuckslipOrder) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BuckslipOrder) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BuckslipOrder) SetStatus(v string) {
	o.Status = &v
}

// GetQuantityOrdered returns the QuantityOrdered field value if set, zero value otherwise.
func (o *BuckslipOrder) GetQuantityOrdered() float32 {
	if o == nil || o.QuantityOrdered == nil {
		var ret float32
		return ret
	}
	return *o.QuantityOrdered
}

// GetQuantityOrderedOk returns a tuple with the QuantityOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetQuantityOrderedOk() (*float32, bool) {
	if o == nil || o.QuantityOrdered == nil {
		return nil, false
	}
	return o.QuantityOrdered, true
}

// HasQuantityOrdered returns a boolean if a field has been set.
func (o *BuckslipOrder) HasQuantityOrdered() bool {
	if o != nil && o.QuantityOrdered != nil {
		return true
	}

	return false
}

// SetQuantityOrdered gets a reference to the given float32 and assigns it to the QuantityOrdered field.
func (o *BuckslipOrder) SetQuantityOrdered(v float32) {
	o.QuantityOrdered = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *BuckslipOrder) GetUnitPrice() float32 {
	if o == nil || o.UnitPrice == nil {
		var ret float32
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetUnitPriceOk() (*float32, bool) {
	if o == nil || o.UnitPrice == nil {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *BuckslipOrder) HasUnitPrice() bool {
	if o != nil && o.UnitPrice != nil {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given float32 and assigns it to the UnitPrice field.
func (o *BuckslipOrder) SetUnitPrice(v float32) {
	o.UnitPrice = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *BuckslipOrder) GetInventory() float32 {
	if o == nil || o.Inventory == nil {
		var ret float32
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetInventoryOk() (*float32, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *BuckslipOrder) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given float32 and assigns it to the Inventory field.
func (o *BuckslipOrder) SetInventory(v float32) {
	o.Inventory = &v
}

// GetCancelledReason returns the CancelledReason field value if set, zero value otherwise.
func (o *BuckslipOrder) GetCancelledReason() string {
	if o == nil || o.CancelledReason == nil {
		var ret string
		return ret
	}
	return *o.CancelledReason
}

// GetCancelledReasonOk returns a tuple with the CancelledReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetCancelledReasonOk() (*string, bool) {
	if o == nil || o.CancelledReason == nil {
		return nil, false
	}
	return o.CancelledReason, true
}

// HasCancelledReason returns a boolean if a field has been set.
func (o *BuckslipOrder) HasCancelledReason() bool {
	if o != nil && o.CancelledReason != nil {
		return true
	}

	return false
}

// SetCancelledReason gets a reference to the given string and assigns it to the CancelledReason field.
func (o *BuckslipOrder) SetCancelledReason(v string) {
	o.CancelledReason = &v
}

// GetAvailabilityDate returns the AvailabilityDate field value if set, zero value otherwise.
func (o *BuckslipOrder) GetAvailabilityDate() time.Time {
	if o == nil || o.AvailabilityDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AvailabilityDate
}

// GetAvailabilityDateOk returns a tuple with the AvailabilityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetAvailabilityDateOk() (*time.Time, bool) {
	if o == nil || o.AvailabilityDate == nil {
		return nil, false
	}
	return o.AvailabilityDate, true
}

// HasAvailabilityDate returns a boolean if a field has been set.
func (o *BuckslipOrder) HasAvailabilityDate() bool {
	if o != nil && o.AvailabilityDate != nil {
		return true
	}

	return false
}

// SetAvailabilityDate gets a reference to the given time.Time and assigns it to the AvailabilityDate field.
func (o *BuckslipOrder) SetAvailabilityDate(v time.Time) {
	o.AvailabilityDate = &v
}

// GetExpectedAvailabilityDate returns the ExpectedAvailabilityDate field value if set, zero value otherwise.
func (o *BuckslipOrder) GetExpectedAvailabilityDate() time.Time {
	if o == nil || o.ExpectedAvailabilityDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpectedAvailabilityDate
}

// GetExpectedAvailabilityDateOk returns a tuple with the ExpectedAvailabilityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuckslipOrder) GetExpectedAvailabilityDateOk() (*time.Time, bool) {
	if o == nil || o.ExpectedAvailabilityDate == nil {
		return nil, false
	}
	return o.ExpectedAvailabilityDate, true
}

// HasExpectedAvailabilityDate returns a boolean if a field has been set.
func (o *BuckslipOrder) HasExpectedAvailabilityDate() bool {
	if o != nil && o.ExpectedAvailabilityDate != nil {
		return true
	}

	return false
}

// SetExpectedAvailabilityDate gets a reference to the given time.Time and assigns it to the ExpectedAvailabilityDate field.
func (o *BuckslipOrder) SetExpectedAvailabilityDate(v time.Time) {
	o.ExpectedAvailabilityDate = &v
}

func (o BuckslipOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date_created"] = o.DateCreated
	}
	if true {
		toSerialize["date_modified"] = o.DateModified
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if true {
		toSerialize["object"] = o.Object
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.BuckslipId != nil {
		toSerialize["buckslip_id"] = o.BuckslipId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.QuantityOrdered != nil {
		toSerialize["quantity_ordered"] = o.QuantityOrdered
	}
	if o.UnitPrice != nil {
		toSerialize["unit_price"] = o.UnitPrice
	}
	if o.Inventory != nil {
		toSerialize["inventory"] = o.Inventory
	}
	if o.CancelledReason != nil {
		toSerialize["cancelled_reason"] = o.CancelledReason
	}
	if o.AvailabilityDate != nil {
		toSerialize["availability_date"] = o.AvailabilityDate
	}
	if o.ExpectedAvailabilityDate != nil {
		toSerialize["expected_availability_date"] = o.ExpectedAvailabilityDate
	}
	return json.Marshal(toSerialize)
}

type NullableBuckslipOrder struct {
	value *BuckslipOrder
	isSet bool
}

func (v NullableBuckslipOrder) Get() *BuckslipOrder {
	return v.value
}

func (v *NullableBuckslipOrder) Set(val *BuckslipOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableBuckslipOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableBuckslipOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuckslipOrder(val *BuckslipOrder) *NullableBuckslipOrder {
	return &NullableBuckslipOrder{value: val, isSet: true}
}

func (v NullableBuckslipOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuckslipOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


