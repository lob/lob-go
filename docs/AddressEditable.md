# AddressEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The building number, street name, street suffix, and any street directionals. For US addresses, the max length is 64 characters. | [optional] 
**AddressLine2** | Pointer to **NullableString** | The suite or apartment number of the recipient address, if applicable. For US addresses, the max length is 64 characters. | [optional] 
**AddressCity** | Pointer to **NullableString** |  | [optional] 
**AddressState** | Pointer to **NullableString** |  | [optional] 
**AddressZip** | Pointer to **NullableString** | Optional postal code. For US addresses, must be either 5 or 9 digits. | [optional] 
**AddressCountry** | Pointer to [**CountryExtended**](CountryExtended.md) |  | [optional] 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Name** | Pointer to **NullableString** | name associated with address. | [optional] 
**Company** | Pointer to **NullableString** | Either &#x60;name&#x60; or &#x60;company&#x60; is required, you may also add both. | [optional] 
**Phone** | Pointer to **NullableString** | Must be no longer than 40 characters. | [optional] 
**Email** | Pointer to **NullableString** | Must be no longer than 100 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 

## Methods

### NewAddressEditable

`func NewAddressEditable() *AddressEditable`

NewAddressEditable instantiates a new AddressEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressEditableWithDefaults

`func NewAddressEditableWithDefaults() *AddressEditable`

NewAddressEditableWithDefaults instantiates a new AddressEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *AddressEditable) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AddressEditable) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AddressEditable) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *AddressEditable) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *AddressEditable) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressEditable) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressEditable) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AddressEditable) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *AddressEditable) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *AddressEditable) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetAddressCity

`func (o *AddressEditable) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *AddressEditable) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *AddressEditable) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *AddressEditable) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### SetAddressCityNil

`func (o *AddressEditable) SetAddressCityNil(b bool)`

 SetAddressCityNil sets the value for AddressCity to be an explicit nil

### UnsetAddressCity
`func (o *AddressEditable) UnsetAddressCity()`

UnsetAddressCity ensures that no value is present for AddressCity, not even an explicit nil
### GetAddressState

`func (o *AddressEditable) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *AddressEditable) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *AddressEditable) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *AddressEditable) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### SetAddressStateNil

`func (o *AddressEditable) SetAddressStateNil(b bool)`

 SetAddressStateNil sets the value for AddressState to be an explicit nil

### UnsetAddressState
`func (o *AddressEditable) UnsetAddressState()`

UnsetAddressState ensures that no value is present for AddressState, not even an explicit nil
### GetAddressZip

`func (o *AddressEditable) GetAddressZip() string`

GetAddressZip returns the AddressZip field if non-nil, zero value otherwise.

### GetAddressZipOk

`func (o *AddressEditable) GetAddressZipOk() (*string, bool)`

GetAddressZipOk returns a tuple with the AddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZip

`func (o *AddressEditable) SetAddressZip(v string)`

SetAddressZip sets AddressZip field to given value.

### HasAddressZip

`func (o *AddressEditable) HasAddressZip() bool`

HasAddressZip returns a boolean if a field has been set.

### SetAddressZipNil

`func (o *AddressEditable) SetAddressZipNil(b bool)`

 SetAddressZipNil sets the value for AddressZip to be an explicit nil

### UnsetAddressZip
`func (o *AddressEditable) UnsetAddressZip()`

UnsetAddressZip ensures that no value is present for AddressZip, not even an explicit nil
### GetAddressCountry

`func (o *AddressEditable) GetAddressCountry() CountryExtended`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *AddressEditable) GetAddressCountryOk() (*CountryExtended, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *AddressEditable) SetAddressCountry(v CountryExtended)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *AddressEditable) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetDescription

`func (o *AddressEditable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddressEditable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddressEditable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddressEditable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddressEditable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddressEditable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *AddressEditable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressEditable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressEditable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressEditable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AddressEditable) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddressEditable) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompany

`func (o *AddressEditable) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressEditable) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressEditable) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressEditable) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *AddressEditable) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *AddressEditable) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetPhone

`func (o *AddressEditable) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressEditable) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressEditable) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressEditable) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *AddressEditable) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *AddressEditable) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *AddressEditable) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressEditable) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressEditable) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressEditable) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AddressEditable) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AddressEditable) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMetadata

`func (o *AddressEditable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddressEditable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddressEditable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddressEditable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


