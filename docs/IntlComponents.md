# IntlComponents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryNumber** | Pointer to **string** | The numeric or alphanumeric part of an address preceding the street name. Often the house, building, or PO Box number. | [optional] 
**StreetName** | Pointer to **string** | The name of the street. | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) two letter code for the state.  | [optional] 
**PostalCode** | Pointer to **string** | The postal code. | [optional] 

## Methods

### NewIntlComponents

`func NewIntlComponents() *IntlComponents`

NewIntlComponents instantiates a new IntlComponents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntlComponentsWithDefaults

`func NewIntlComponentsWithDefaults() *IntlComponents`

NewIntlComponentsWithDefaults instantiates a new IntlComponents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryNumber

`func (o *IntlComponents) GetPrimaryNumber() string`

GetPrimaryNumber returns the PrimaryNumber field if non-nil, zero value otherwise.

### GetPrimaryNumberOk

`func (o *IntlComponents) GetPrimaryNumberOk() (*string, bool)`

GetPrimaryNumberOk returns a tuple with the PrimaryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNumber

`func (o *IntlComponents) SetPrimaryNumber(v string)`

SetPrimaryNumber sets PrimaryNumber field to given value.

### HasPrimaryNumber

`func (o *IntlComponents) HasPrimaryNumber() bool`

HasPrimaryNumber returns a boolean if a field has been set.

### GetStreetName

`func (o *IntlComponents) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *IntlComponents) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *IntlComponents) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *IntlComponents) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetCity

`func (o *IntlComponents) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *IntlComponents) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *IntlComponents) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *IntlComponents) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *IntlComponents) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IntlComponents) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IntlComponents) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IntlComponents) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *IntlComponents) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *IntlComponents) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *IntlComponents) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *IntlComponents) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


