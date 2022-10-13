# UsAutocompletionsWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressPrefix** | **string** | Only accepts numbers and street names in an alphanumeric format.  | 
**City** | Pointer to **string** | An optional city input used to filter suggestions. Case insensitive and does not match partial abbreviations.  | [optional] 
**State** | Pointer to **string** | An optional state input used to filter suggestions. Case insensitive and does not match partial abbreviations.  | [optional] 
**ZipCode** | Pointer to **string** | An optional ZIP Code input used to filter suggestions. Matches partial entries.  | [optional] 
**GeoIpSort** | Pointer to **bool** | If &#x60;true&#x60;, sort suggestions by proximity to the IP set in the &#x60;X-Forwarded-For&#x60; header.  | [optional] 

## Methods

### NewUsAutocompletionsWritable

`func NewUsAutocompletionsWritable(addressPrefix string, ) *UsAutocompletionsWritable`

NewUsAutocompletionsWritable instantiates a new UsAutocompletionsWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsAutocompletionsWritableWithDefaults

`func NewUsAutocompletionsWritableWithDefaults() *UsAutocompletionsWritable`

NewUsAutocompletionsWritableWithDefaults instantiates a new UsAutocompletionsWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressPrefix

`func (o *UsAutocompletionsWritable) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *UsAutocompletionsWritable) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *UsAutocompletionsWritable) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.


### GetCity

`func (o *UsAutocompletionsWritable) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UsAutocompletionsWritable) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UsAutocompletionsWritable) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UsAutocompletionsWritable) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *UsAutocompletionsWritable) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UsAutocompletionsWritable) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UsAutocompletionsWritable) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UsAutocompletionsWritable) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *UsAutocompletionsWritable) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UsAutocompletionsWritable) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UsAutocompletionsWritable) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UsAutocompletionsWritable) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetGeoIpSort

`func (o *UsAutocompletionsWritable) GetGeoIpSort() bool`

GetGeoIpSort returns the GeoIpSort field if non-nil, zero value otherwise.

### GetGeoIpSortOk

`func (o *UsAutocompletionsWritable) GetGeoIpSortOk() (*bool, bool)`

GetGeoIpSortOk returns a tuple with the GeoIpSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoIpSort

`func (o *UsAutocompletionsWritable) SetGeoIpSort(v bool)`

SetGeoIpSort sets GeoIpSort field to given value.

### HasGeoIpSort

`func (o *UsAutocompletionsWritable) HasGeoIpSort() bool`

HasGeoIpSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


