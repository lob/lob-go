# OptionalAddressColumnMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine2** | **NullableString** | The column header from the csv file that should be mapped to the optional field \&quot;address_line2\&quot; | [default to "null"]
**Company** | **NullableString** | The column header from the csv file that should be mapped to the optional field \&quot;company\&quot; | [default to "null"]
**AddressCountry** | **NullableString** | The column header from the csv file that should be mapped to the optional field \&quot;address_country\&quot; | [default to "null"]

## Methods

### NewOptionalAddressColumnMapping

`func NewOptionalAddressColumnMapping(addressLine2 NullableString, company NullableString, addressCountry NullableString, ) *OptionalAddressColumnMapping`

NewOptionalAddressColumnMapping instantiates a new OptionalAddressColumnMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionalAddressColumnMappingWithDefaults

`func NewOptionalAddressColumnMappingWithDefaults() *OptionalAddressColumnMapping`

NewOptionalAddressColumnMappingWithDefaults instantiates a new OptionalAddressColumnMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine2

`func (o *OptionalAddressColumnMapping) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OptionalAddressColumnMapping) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OptionalAddressColumnMapping) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.


### SetAddressLine2Nil

`func (o *OptionalAddressColumnMapping) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *OptionalAddressColumnMapping) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetCompany

`func (o *OptionalAddressColumnMapping) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *OptionalAddressColumnMapping) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *OptionalAddressColumnMapping) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *OptionalAddressColumnMapping) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *OptionalAddressColumnMapping) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetAddressCountry

`func (o *OptionalAddressColumnMapping) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *OptionalAddressColumnMapping) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *OptionalAddressColumnMapping) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.


### SetAddressCountryNil

`func (o *OptionalAddressColumnMapping) SetAddressCountryNil(b bool)`

 SetAddressCountryNil sets the value for AddressCountry to be an explicit nil

### UnsetAddressCountry
`func (o *OptionalAddressColumnMapping) UnsetAddressCountry()`

UnsetAddressCountry ensures that no value is present for AddressCountry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


