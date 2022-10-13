# IntlVerificationWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | Pointer to **NullableString** | The intended recipient, typically a person&#39;s or firm&#39;s name. | [optional] 
**PrimaryLine** | Pointer to **string** | The primary delivery line (usually the street address) of the address.  | [optional] 
**SecondaryLine** | Pointer to **string** | The secondary delivery line of the address. This field is typically empty but may contain information if &#x60;primary_line&#x60; is too long.  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The name of the state. | [optional] 
**PostalCode** | Pointer to **string** | The postal code. | [optional] 
**Country** | Pointer to [**CountryExtended**](CountryExtended.md) |  | [optional] 
**Address** | Pointer to **string** | The entire address in one string (e.g., \&quot;370 Water St C1N 1C4\&quot;).  | [optional] 

## Methods

### NewIntlVerificationWritable

`func NewIntlVerificationWritable() *IntlVerificationWritable`

NewIntlVerificationWritable instantiates a new IntlVerificationWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntlVerificationWritableWithDefaults

`func NewIntlVerificationWritableWithDefaults() *IntlVerificationWritable`

NewIntlVerificationWritableWithDefaults instantiates a new IntlVerificationWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *IntlVerificationWritable) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *IntlVerificationWritable) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *IntlVerificationWritable) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *IntlVerificationWritable) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *IntlVerificationWritable) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *IntlVerificationWritable) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetPrimaryLine

`func (o *IntlVerificationWritable) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *IntlVerificationWritable) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *IntlVerificationWritable) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.

### HasPrimaryLine

`func (o *IntlVerificationWritable) HasPrimaryLine() bool`

HasPrimaryLine returns a boolean if a field has been set.

### GetSecondaryLine

`func (o *IntlVerificationWritable) GetSecondaryLine() string`

GetSecondaryLine returns the SecondaryLine field if non-nil, zero value otherwise.

### GetSecondaryLineOk

`func (o *IntlVerificationWritable) GetSecondaryLineOk() (*string, bool)`

GetSecondaryLineOk returns a tuple with the SecondaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLine

`func (o *IntlVerificationWritable) SetSecondaryLine(v string)`

SetSecondaryLine sets SecondaryLine field to given value.

### HasSecondaryLine

`func (o *IntlVerificationWritable) HasSecondaryLine() bool`

HasSecondaryLine returns a boolean if a field has been set.

### GetCity

`func (o *IntlVerificationWritable) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *IntlVerificationWritable) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *IntlVerificationWritable) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *IntlVerificationWritable) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *IntlVerificationWritable) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IntlVerificationWritable) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IntlVerificationWritable) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IntlVerificationWritable) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *IntlVerificationWritable) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *IntlVerificationWritable) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *IntlVerificationWritable) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *IntlVerificationWritable) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *IntlVerificationWritable) GetCountry() CountryExtended`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IntlVerificationWritable) GetCountryOk() (*CountryExtended, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IntlVerificationWritable) SetCountry(v CountryExtended)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *IntlVerificationWritable) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAddress

`func (o *IntlVerificationWritable) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IntlVerificationWritable) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IntlVerificationWritable) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IntlVerificationWritable) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


