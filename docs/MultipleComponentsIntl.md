# MultipleComponentsIntl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | Pointer to **NullableString** | The intended recipient, typically a person&#39;s or firm&#39;s name. | [optional] 
**PrimaryLine** | **string** | The primary delivery line (usually the street address) of the address.  | 
**SecondaryLine** | Pointer to **string** | The secondary delivery line of the address. This field is typically empty but may contain information if &#x60;primary_line&#x60; is too long.  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The name of the state. | [optional] 
**PostalCode** | Pointer to **string** | The postal code. | [optional] 
**Country** | [**CountryExtended**](CountryExtended.md) |  | 

## Methods

### NewMultipleComponentsIntl

`func NewMultipleComponentsIntl(primaryLine string, country CountryExtended, ) *MultipleComponentsIntl`

NewMultipleComponentsIntl instantiates a new MultipleComponentsIntl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleComponentsIntlWithDefaults

`func NewMultipleComponentsIntlWithDefaults() *MultipleComponentsIntl`

NewMultipleComponentsIntlWithDefaults instantiates a new MultipleComponentsIntl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *MultipleComponentsIntl) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *MultipleComponentsIntl) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *MultipleComponentsIntl) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *MultipleComponentsIntl) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *MultipleComponentsIntl) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *MultipleComponentsIntl) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetPrimaryLine

`func (o *MultipleComponentsIntl) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *MultipleComponentsIntl) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *MultipleComponentsIntl) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.


### GetSecondaryLine

`func (o *MultipleComponentsIntl) GetSecondaryLine() string`

GetSecondaryLine returns the SecondaryLine field if non-nil, zero value otherwise.

### GetSecondaryLineOk

`func (o *MultipleComponentsIntl) GetSecondaryLineOk() (*string, bool)`

GetSecondaryLineOk returns a tuple with the SecondaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLine

`func (o *MultipleComponentsIntl) SetSecondaryLine(v string)`

SetSecondaryLine sets SecondaryLine field to given value.

### HasSecondaryLine

`func (o *MultipleComponentsIntl) HasSecondaryLine() bool`

HasSecondaryLine returns a boolean if a field has been set.

### GetCity

`func (o *MultipleComponentsIntl) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MultipleComponentsIntl) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MultipleComponentsIntl) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MultipleComponentsIntl) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *MultipleComponentsIntl) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MultipleComponentsIntl) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MultipleComponentsIntl) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MultipleComponentsIntl) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *MultipleComponentsIntl) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *MultipleComponentsIntl) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *MultipleComponentsIntl) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *MultipleComponentsIntl) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *MultipleComponentsIntl) GetCountry() CountryExtended`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MultipleComponentsIntl) GetCountryOk() (*CountryExtended, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MultipleComponentsIntl) SetCountry(v CountryExtended)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


