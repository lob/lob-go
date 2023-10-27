# BuckslipOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Object** | **string** | Value is type of resource. | 
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;bo_&#x60;. | [optional] 
**BuckslipId** | Pointer to **string** | Unique identifier prefixed with &#x60;bck_&#x60;. | [optional] 
**Status** | Pointer to **string** | The status of the buckslip order. | [optional] 
**QuantityOrdered** | Pointer to **float32** | The quantity of buckslips ordered. | [optional] [default to 0]
**UnitPrice** | Pointer to **float32** | The unit price for the buckslip order. | [optional] [default to 0]
**Inventory** | Pointer to **float32** | The inventory of the buckslip order. | [optional] [default to 0]
**CancelledReason** | Pointer to **string** | The reason for cancellation. | [optional] 
**AvailabilityDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | [optional] 
**ExpectedAvailabilityDate** | Pointer to **time.Time** | The fixed deadline for the buckslips to be printed. | [optional] 

## Methods

### NewBuckslipOrder

`func NewBuckslipOrder(dateCreated time.Time, dateModified time.Time, object string, ) *BuckslipOrder`

NewBuckslipOrder instantiates a new BuckslipOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuckslipOrderWithDefaults

`func NewBuckslipOrderWithDefaults() *BuckslipOrder`

NewBuckslipOrderWithDefaults instantiates a new BuckslipOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *BuckslipOrder) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BuckslipOrder) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BuckslipOrder) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *BuckslipOrder) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *BuckslipOrder) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *BuckslipOrder) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetDeleted

`func (o *BuckslipOrder) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *BuckslipOrder) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *BuckslipOrder) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *BuckslipOrder) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetObject

`func (o *BuckslipOrder) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BuckslipOrder) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BuckslipOrder) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *BuckslipOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuckslipOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuckslipOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BuckslipOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBuckslipId

`func (o *BuckslipOrder) GetBuckslipId() string`

GetBuckslipId returns the BuckslipId field if non-nil, zero value otherwise.

### GetBuckslipIdOk

`func (o *BuckslipOrder) GetBuckslipIdOk() (*string, bool)`

GetBuckslipIdOk returns a tuple with the BuckslipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckslipId

`func (o *BuckslipOrder) SetBuckslipId(v string)`

SetBuckslipId sets BuckslipId field to given value.

### HasBuckslipId

`func (o *BuckslipOrder) HasBuckslipId() bool`

HasBuckslipId returns a boolean if a field has been set.

### GetStatus

`func (o *BuckslipOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuckslipOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuckslipOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BuckslipOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *BuckslipOrder) GetQuantityOrdered() float32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *BuckslipOrder) GetQuantityOrderedOk() (*float32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *BuckslipOrder) SetQuantityOrdered(v float32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *BuckslipOrder) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetUnitPrice

`func (o *BuckslipOrder) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *BuckslipOrder) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *BuckslipOrder) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *BuckslipOrder) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetInventory

`func (o *BuckslipOrder) GetInventory() float32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *BuckslipOrder) GetInventoryOk() (*float32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *BuckslipOrder) SetInventory(v float32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *BuckslipOrder) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetCancelledReason

`func (o *BuckslipOrder) GetCancelledReason() string`

GetCancelledReason returns the CancelledReason field if non-nil, zero value otherwise.

### GetCancelledReasonOk

`func (o *BuckslipOrder) GetCancelledReasonOk() (*string, bool)`

GetCancelledReasonOk returns a tuple with the CancelledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledReason

`func (o *BuckslipOrder) SetCancelledReason(v string)`

SetCancelledReason sets CancelledReason field to given value.

### HasCancelledReason

`func (o *BuckslipOrder) HasCancelledReason() bool`

HasCancelledReason returns a boolean if a field has been set.

### GetAvailabilityDate

`func (o *BuckslipOrder) GetAvailabilityDate() time.Time`

GetAvailabilityDate returns the AvailabilityDate field if non-nil, zero value otherwise.

### GetAvailabilityDateOk

`func (o *BuckslipOrder) GetAvailabilityDateOk() (*time.Time, bool)`

GetAvailabilityDateOk returns a tuple with the AvailabilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityDate

`func (o *BuckslipOrder) SetAvailabilityDate(v time.Time)`

SetAvailabilityDate sets AvailabilityDate field to given value.

### HasAvailabilityDate

`func (o *BuckslipOrder) HasAvailabilityDate() bool`

HasAvailabilityDate returns a boolean if a field has been set.

### GetExpectedAvailabilityDate

`func (o *BuckslipOrder) GetExpectedAvailabilityDate() time.Time`

GetExpectedAvailabilityDate returns the ExpectedAvailabilityDate field if non-nil, zero value otherwise.

### GetExpectedAvailabilityDateOk

`func (o *BuckslipOrder) GetExpectedAvailabilityDateOk() (*time.Time, bool)`

GetExpectedAvailabilityDateOk returns a tuple with the ExpectedAvailabilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedAvailabilityDate

`func (o *BuckslipOrder) SetExpectedAvailabilityDate(v time.Time)`

SetExpectedAvailabilityDate sets ExpectedAvailabilityDate field to given value.

### HasExpectedAvailabilityDate

`func (o *BuckslipOrder) HasExpectedAvailabilityDate() bool`

HasExpectedAvailabilityDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


