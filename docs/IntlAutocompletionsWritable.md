# IntlAutocompletionsWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressPrefix** | **string** | Only accepts numbers and street names in an alphanumeric format.  | 
**City** | Pointer to **string** | An optional city input used to filter suggestions. Case insensitive and does not match partial abbreviations.  | [optional] 
**State** | Pointer to **string** | An optional state input used to filter suggestions. Case insensitive and does not match partial abbreviations.  | [optional] 
**ZipCode** | Pointer to **string** | An optional Zip Code input used to filter suggestions. Matches partial entries.  | [optional] 
**Country** | [**CountryExtended**](CountryExtended.md) |  | 

## Methods

### NewIntlAutocompletionsWritable

`func NewIntlAutocompletionsWritable(addressPrefix string, country CountryExtended, ) *IntlAutocompletionsWritable`

NewIntlAutocompletionsWritable instantiates a new IntlAutocompletionsWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntlAutocompletionsWritableWithDefaults

`func NewIntlAutocompletionsWritableWithDefaults() *IntlAutocompletionsWritable`

NewIntlAutocompletionsWritableWithDefaults instantiates a new IntlAutocompletionsWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressPrefix

`func (o *IntlAutocompletionsWritable) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *IntlAutocompletionsWritable) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *IntlAutocompletionsWritable) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.


### GetCity

`func (o *IntlAutocompletionsWritable) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *IntlAutocompletionsWritable) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *IntlAutocompletionsWritable) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *IntlAutocompletionsWritable) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *IntlAutocompletionsWritable) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IntlAutocompletionsWritable) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IntlAutocompletionsWritable) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IntlAutocompletionsWritable) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *IntlAutocompletionsWritable) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *IntlAutocompletionsWritable) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *IntlAutocompletionsWritable) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *IntlAutocompletionsWritable) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCountry

`func (o *IntlAutocompletionsWritable) GetCountry() CountryExtended`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IntlAutocompletionsWritable) GetCountryOk() (*CountryExtended, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IntlAutocompletionsWritable) SetCountry(v CountryExtended)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


