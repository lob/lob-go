# BuckslipEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Front** | **string** | A PDF template for the front of the buckslip | 
**Back** | Pointer to **string** | A PDF template for the back of the buckslip | [optional] 
**Description** | Pointer to **NullableString** | Description of the buckslip. | [optional] 
**Size** | Pointer to **string** | The size of the buckslip | [optional] [default to "8.75x3.75"]

## Methods

### NewBuckslipEditable

`func NewBuckslipEditable(front string, ) *BuckslipEditable`

NewBuckslipEditable instantiates a new BuckslipEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuckslipEditableWithDefaults

`func NewBuckslipEditableWithDefaults() *BuckslipEditable`

NewBuckslipEditableWithDefaults instantiates a new BuckslipEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFront

`func (o *BuckslipEditable) GetFront() string`

GetFront returns the Front field if non-nil, zero value otherwise.

### GetFrontOk

`func (o *BuckslipEditable) GetFrontOk() (*string, bool)`

GetFrontOk returns a tuple with the Front field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFront

`func (o *BuckslipEditable) SetFront(v string)`

SetFront sets Front field to given value.


### GetBack

`func (o *BuckslipEditable) GetBack() string`

GetBack returns the Back field if non-nil, zero value otherwise.

### GetBackOk

`func (o *BuckslipEditable) GetBackOk() (*string, bool)`

GetBackOk returns a tuple with the Back field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBack

`func (o *BuckslipEditable) SetBack(v string)`

SetBack sets Back field to given value.

### HasBack

`func (o *BuckslipEditable) HasBack() bool`

HasBack returns a boolean if a field has been set.

### GetDescription

`func (o *BuckslipEditable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuckslipEditable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuckslipEditable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuckslipEditable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BuckslipEditable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BuckslipEditable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSize

`func (o *BuckslipEditable) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BuckslipEditable) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BuckslipEditable) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *BuckslipEditable) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


