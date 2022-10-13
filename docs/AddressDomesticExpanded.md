# AddressDomesticExpanded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The building number, street name, street suffix, and any street directionals. For US addresses, the max length is 64 characters. | [optional] 
**AddressLine2** | Pointer to **NullableString** | The suite or apartment number of the recipient address, if applicable. For US addresses, the max length is 64 characters. | [optional] 
**AddressCity** | Pointer to **NullableString** |  | [optional] 
**AddressState** | Pointer to **NullableString** |  | [optional] 
**AddressZip** | Pointer to **NullableString** | Optional postal code. For US addresses, must be either 5 or 9 digits. | [optional] 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Name** | Pointer to **NullableString** | Either &#x60;name&#x60; or &#x60;company&#x60; is required, you may also add both. Must be no longer than 40 characters. If both &#x60;name&#x60; and &#x60;company&#x60; are provided, they will be printed on two separate lines above the rest of the address.  | [optional] 
**Company** | Pointer to **NullableString** | Either &#x60;name&#x60; or &#x60;company&#x60; is required, you may also add both. | [optional] 
**Phone** | Pointer to **NullableString** | Must be no longer than 40 characters. | [optional] 
**Email** | Pointer to **NullableString** | Must be no longer than 100 characters. | [optional] 
**AddressCountry** | Pointer to **string** | The country associated with this address. | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 

## Methods

### NewAddressDomesticExpanded

`func NewAddressDomesticExpanded() *AddressDomesticExpanded`

NewAddressDomesticExpanded instantiates a new AddressDomesticExpanded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressDomesticExpandedWithDefaults

`func NewAddressDomesticExpandedWithDefaults() *AddressDomesticExpanded`

NewAddressDomesticExpandedWithDefaults instantiates a new AddressDomesticExpanded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *AddressDomesticExpanded) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AddressDomesticExpanded) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AddressDomesticExpanded) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *AddressDomesticExpanded) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *AddressDomesticExpanded) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressDomesticExpanded) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressDomesticExpanded) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AddressDomesticExpanded) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *AddressDomesticExpanded) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *AddressDomesticExpanded) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetAddressCity

`func (o *AddressDomesticExpanded) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *AddressDomesticExpanded) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *AddressDomesticExpanded) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *AddressDomesticExpanded) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### SetAddressCityNil

`func (o *AddressDomesticExpanded) SetAddressCityNil(b bool)`

 SetAddressCityNil sets the value for AddressCity to be an explicit nil

### UnsetAddressCity
`func (o *AddressDomesticExpanded) UnsetAddressCity()`

UnsetAddressCity ensures that no value is present for AddressCity, not even an explicit nil
### GetAddressState

`func (o *AddressDomesticExpanded) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *AddressDomesticExpanded) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *AddressDomesticExpanded) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *AddressDomesticExpanded) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### SetAddressStateNil

`func (o *AddressDomesticExpanded) SetAddressStateNil(b bool)`

 SetAddressStateNil sets the value for AddressState to be an explicit nil

### UnsetAddressState
`func (o *AddressDomesticExpanded) UnsetAddressState()`

UnsetAddressState ensures that no value is present for AddressState, not even an explicit nil
### GetAddressZip

`func (o *AddressDomesticExpanded) GetAddressZip() string`

GetAddressZip returns the AddressZip field if non-nil, zero value otherwise.

### GetAddressZipOk

`func (o *AddressDomesticExpanded) GetAddressZipOk() (*string, bool)`

GetAddressZipOk returns a tuple with the AddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZip

`func (o *AddressDomesticExpanded) SetAddressZip(v string)`

SetAddressZip sets AddressZip field to given value.

### HasAddressZip

`func (o *AddressDomesticExpanded) HasAddressZip() bool`

HasAddressZip returns a boolean if a field has been set.

### SetAddressZipNil

`func (o *AddressDomesticExpanded) SetAddressZipNil(b bool)`

 SetAddressZipNil sets the value for AddressZip to be an explicit nil

### UnsetAddressZip
`func (o *AddressDomesticExpanded) UnsetAddressZip()`

UnsetAddressZip ensures that no value is present for AddressZip, not even an explicit nil
### GetDescription

`func (o *AddressDomesticExpanded) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddressDomesticExpanded) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddressDomesticExpanded) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddressDomesticExpanded) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddressDomesticExpanded) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddressDomesticExpanded) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *AddressDomesticExpanded) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressDomesticExpanded) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressDomesticExpanded) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressDomesticExpanded) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AddressDomesticExpanded) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddressDomesticExpanded) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompany

`func (o *AddressDomesticExpanded) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressDomesticExpanded) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressDomesticExpanded) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressDomesticExpanded) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *AddressDomesticExpanded) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *AddressDomesticExpanded) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetPhone

`func (o *AddressDomesticExpanded) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressDomesticExpanded) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressDomesticExpanded) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressDomesticExpanded) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *AddressDomesticExpanded) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *AddressDomesticExpanded) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *AddressDomesticExpanded) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressDomesticExpanded) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressDomesticExpanded) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressDomesticExpanded) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AddressDomesticExpanded) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AddressDomesticExpanded) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddressCountry

`func (o *AddressDomesticExpanded) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *AddressDomesticExpanded) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *AddressDomesticExpanded) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *AddressDomesticExpanded) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetMetadata

`func (o *AddressDomesticExpanded) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddressDomesticExpanded) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddressDomesticExpanded) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddressDomesticExpanded) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


