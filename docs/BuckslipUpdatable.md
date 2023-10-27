# BuckslipUpdatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description of the buckslip. | [optional] 
**AutoReorder** | Pointer to **bool** | Allows for auto reordering | [optional] 
**ReorderQuantity** | Pointer to **float32** | The quantity of items to be reordered (only required when auto_reorder is true). | [optional] 

## Methods

### NewBuckslipUpdatable

`func NewBuckslipUpdatable() *BuckslipUpdatable`

NewBuckslipUpdatable instantiates a new BuckslipUpdatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuckslipUpdatableWithDefaults

`func NewBuckslipUpdatableWithDefaults() *BuckslipUpdatable`

NewBuckslipUpdatableWithDefaults instantiates a new BuckslipUpdatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BuckslipUpdatable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuckslipUpdatable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuckslipUpdatable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuckslipUpdatable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BuckslipUpdatable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BuckslipUpdatable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAutoReorder

`func (o *BuckslipUpdatable) GetAutoReorder() bool`

GetAutoReorder returns the AutoReorder field if non-nil, zero value otherwise.

### GetAutoReorderOk

`func (o *BuckslipUpdatable) GetAutoReorderOk() (*bool, bool)`

GetAutoReorderOk returns a tuple with the AutoReorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoReorder

`func (o *BuckslipUpdatable) SetAutoReorder(v bool)`

SetAutoReorder sets AutoReorder field to given value.

### HasAutoReorder

`func (o *BuckslipUpdatable) HasAutoReorder() bool`

HasAutoReorder returns a boolean if a field has been set.

### GetReorderQuantity

`func (o *BuckslipUpdatable) GetReorderQuantity() float32`

GetReorderQuantity returns the ReorderQuantity field if non-nil, zero value otherwise.

### GetReorderQuantityOk

`func (o *BuckslipUpdatable) GetReorderQuantityOk() (*float32, bool)`

GetReorderQuantityOk returns a tuple with the ReorderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderQuantity

`func (o *BuckslipUpdatable) SetReorderQuantity(v float32)`

SetReorderQuantity sets ReorderQuantity field to given value.

### HasReorderQuantity

`func (o *BuckslipUpdatable) HasReorderQuantity() bool`

HasReorderQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


