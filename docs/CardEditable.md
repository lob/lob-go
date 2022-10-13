# CardEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Front** | **string** | A PDF template for the front of the card | 
**Back** | Pointer to **string** | A PDF template for the back of the card | [optional] [default to "https://s3.us-west-2.amazonaws.com/public.lob.com/assets/card_blank_horizontal.pdf"]
**Size** | Pointer to **string** | The size of the card | [optional] [default to "2.125x3.375"]
**Description** | Pointer to **NullableString** | Description of the card. | [optional] 

## Methods

### NewCardEditable

`func NewCardEditable(front string, ) *CardEditable`

NewCardEditable instantiates a new CardEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardEditableWithDefaults

`func NewCardEditableWithDefaults() *CardEditable`

NewCardEditableWithDefaults instantiates a new CardEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFront

`func (o *CardEditable) GetFront() string`

GetFront returns the Front field if non-nil, zero value otherwise.

### GetFrontOk

`func (o *CardEditable) GetFrontOk() (*string, bool)`

GetFrontOk returns a tuple with the Front field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFront

`func (o *CardEditable) SetFront(v string)`

SetFront sets Front field to given value.


### GetBack

`func (o *CardEditable) GetBack() string`

GetBack returns the Back field if non-nil, zero value otherwise.

### GetBackOk

`func (o *CardEditable) GetBackOk() (*string, bool)`

GetBackOk returns a tuple with the Back field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBack

`func (o *CardEditable) SetBack(v string)`

SetBack sets Back field to given value.

### HasBack

`func (o *CardEditable) HasBack() bool`

HasBack returns a boolean if a field has been set.

### GetSize

`func (o *CardEditable) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CardEditable) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CardEditable) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CardEditable) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDescription

`func (o *CardEditable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CardEditable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CardEditable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CardEditable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CardEditable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CardEditable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


