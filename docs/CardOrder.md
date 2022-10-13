# CardOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;co_&#x60;. | [optional] 
**CardId** | Pointer to **string** | Unique identifier prefixed with &#x60;card_&#x60;. | [optional] 
**Status** | Pointer to **string** | The status of the card order. | [optional] 
**Inventory** | Pointer to **float32** | The inventory of the card order. | [optional] [default to 0]
**QuantityOrdered** | Pointer to **float32** | The quantity of cards ordered | [optional] [default to 0]
**UnitPrice** | Pointer to **float32** | The unit price for the card order. | [optional] [default to 0]
**CancelledReason** | Pointer to **string** | The reason for cancellation. | [optional] 
**AvailabilityDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | [optional] 
**ExpectedAvailabilityDate** | Pointer to **time.Time** | The fixed deadline for the cards to be printed. | [optional] 
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Object** | **string** | Value is type of resource. | 

## Methods

### NewCardOrder

`func NewCardOrder(dateCreated time.Time, dateModified time.Time, object string, ) *CardOrder`

NewCardOrder instantiates a new CardOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardOrderWithDefaults

`func NewCardOrderWithDefaults() *CardOrder`

NewCardOrderWithDefaults instantiates a new CardOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCardId

`func (o *CardOrder) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *CardOrder) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *CardOrder) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *CardOrder) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetStatus

`func (o *CardOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CardOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInventory

`func (o *CardOrder) GetInventory() float32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *CardOrder) GetInventoryOk() (*float32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *CardOrder) SetInventory(v float32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *CardOrder) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *CardOrder) GetQuantityOrdered() float32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *CardOrder) GetQuantityOrderedOk() (*float32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *CardOrder) SetQuantityOrdered(v float32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *CardOrder) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetUnitPrice

`func (o *CardOrder) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CardOrder) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CardOrder) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *CardOrder) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetCancelledReason

`func (o *CardOrder) GetCancelledReason() string`

GetCancelledReason returns the CancelledReason field if non-nil, zero value otherwise.

### GetCancelledReasonOk

`func (o *CardOrder) GetCancelledReasonOk() (*string, bool)`

GetCancelledReasonOk returns a tuple with the CancelledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledReason

`func (o *CardOrder) SetCancelledReason(v string)`

SetCancelledReason sets CancelledReason field to given value.

### HasCancelledReason

`func (o *CardOrder) HasCancelledReason() bool`

HasCancelledReason returns a boolean if a field has been set.

### GetAvailabilityDate

`func (o *CardOrder) GetAvailabilityDate() time.Time`

GetAvailabilityDate returns the AvailabilityDate field if non-nil, zero value otherwise.

### GetAvailabilityDateOk

`func (o *CardOrder) GetAvailabilityDateOk() (*time.Time, bool)`

GetAvailabilityDateOk returns a tuple with the AvailabilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityDate

`func (o *CardOrder) SetAvailabilityDate(v time.Time)`

SetAvailabilityDate sets AvailabilityDate field to given value.

### HasAvailabilityDate

`func (o *CardOrder) HasAvailabilityDate() bool`

HasAvailabilityDate returns a boolean if a field has been set.

### GetExpectedAvailabilityDate

`func (o *CardOrder) GetExpectedAvailabilityDate() time.Time`

GetExpectedAvailabilityDate returns the ExpectedAvailabilityDate field if non-nil, zero value otherwise.

### GetExpectedAvailabilityDateOk

`func (o *CardOrder) GetExpectedAvailabilityDateOk() (*time.Time, bool)`

GetExpectedAvailabilityDateOk returns a tuple with the ExpectedAvailabilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedAvailabilityDate

`func (o *CardOrder) SetExpectedAvailabilityDate(v time.Time)`

SetExpectedAvailabilityDate sets ExpectedAvailabilityDate field to given value.

### HasExpectedAvailabilityDate

`func (o *CardOrder) HasExpectedAvailabilityDate() bool`

HasExpectedAvailabilityDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *CardOrder) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CardOrder) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CardOrder) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *CardOrder) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *CardOrder) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *CardOrder) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetDeleted

`func (o *CardOrder) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CardOrder) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CardOrder) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CardOrder) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetObject

`func (o *CardOrder) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CardOrder) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CardOrder) SetObject(v string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


