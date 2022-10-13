# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;adr_&#x60;. | [optional] 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Name** | Pointer to **NullableString** | name associated with address | [optional] 
**Company** | Pointer to **NullableString** | Either &#x60;name&#x60; or &#x60;company&#x60; is required, you may also add both. | [optional] 
**Phone** | Pointer to **NullableString** | Must be no longer than 40 characters. | [optional] 
**Email** | Pointer to **NullableString** | Must be no longer than 100 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **NullableString** |  | [optional] 
**AddressCity** | Pointer to **string** |  | [optional] 
**AddressState** | Pointer to **string** | 2 letter state short-name code | [optional] 
**AddressZip** | Pointer to **string** | Must follow the ZIP format of &#x60;12345&#x60; or ZIP+4 format of &#x60;12345-1234&#x60;.  | [optional] 
**AddressCountry** | Pointer to [**CountryExtendedExpanded**](CountryExtendedExpanded.md) |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "address"]
**DateCreated** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | [optional] 
**DateModified** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | [optional] 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**RecipientMoved** | Pointer to **NullableBool** | Only returned for accounts on certain &lt;a href&#x3D;\&quot;https://dashboard.lob.com/#/settings/editions\&quot;&gt;Print &amp;amp; Mail Editions&lt;/a&gt;. Value is &#x60;true&#x60; if the address was altered because the recipient filed for a &lt;a href&#x3D;\&quot;#ncoa\&quot;&gt;National Change of Address (NCOA)&lt;/a&gt;, &#x60;false&#x60; if the NCOA check was run but no altered address was found, and &#x60;null&#x60; if the NCOA check was not run. The NCOA check does not happen for non-US addresses, for non-deliverable US addresses, or for addresses created before the NCOA feature was added to your account.  | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Address) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Address) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Address) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Address) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *Address) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Address) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Address) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Address) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Address) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Address) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *Address) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Address) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Address) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Address) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Address) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Address) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompany

`func (o *Address) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Address) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Address) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Address) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Address) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Address) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetPhone

`func (o *Address) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Address) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Address) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Address) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *Address) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *Address) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *Address) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Address) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Address) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Address) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Address) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Address) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMetadata

`func (o *Address) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Address) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Address) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Address) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAddressLine1

`func (o *Address) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Address) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Address) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *Address) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *Address) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Address) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Address) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Address) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *Address) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *Address) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetAddressCity

`func (o *Address) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *Address) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *Address) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *Address) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressState

`func (o *Address) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *Address) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *Address) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *Address) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### GetAddressZip

`func (o *Address) GetAddressZip() string`

GetAddressZip returns the AddressZip field if non-nil, zero value otherwise.

### GetAddressZipOk

`func (o *Address) GetAddressZipOk() (*string, bool)`

GetAddressZipOk returns a tuple with the AddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZip

`func (o *Address) SetAddressZip(v string)`

SetAddressZip sets AddressZip field to given value.

### HasAddressZip

`func (o *Address) HasAddressZip() bool`

HasAddressZip returns a boolean if a field has been set.

### GetAddressCountry

`func (o *Address) GetAddressCountry() CountryExtendedExpanded`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *Address) GetAddressCountryOk() (*CountryExtendedExpanded, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *Address) SetAddressCountry(v CountryExtendedExpanded)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *Address) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetObject

`func (o *Address) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Address) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Address) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Address) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDateCreated

`func (o *Address) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Address) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Address) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Address) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateModified

`func (o *Address) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *Address) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *Address) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.

### HasDateModified

`func (o *Address) HasDateModified() bool`

HasDateModified returns a boolean if a field has been set.

### GetDeleted

`func (o *Address) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Address) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Address) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Address) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetRecipientMoved

`func (o *Address) GetRecipientMoved() bool`

GetRecipientMoved returns the RecipientMoved field if non-nil, zero value otherwise.

### GetRecipientMovedOk

`func (o *Address) GetRecipientMovedOk() (*bool, bool)`

GetRecipientMovedOk returns a tuple with the RecipientMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientMoved

`func (o *Address) SetRecipientMoved(v bool)`

SetRecipientMoved sets RecipientMoved field to given value.

### HasRecipientMoved

`func (o *Address) HasRecipientMoved() bool`

HasRecipientMoved returns a boolean if a field has been set.

### SetRecipientMovedNil

`func (o *Address) SetRecipientMovedNil(b bool)`

 SetRecipientMovedNil sets the value for RecipientMoved to be an explicit nil

### UnsetRecipientMoved
`func (o *Address) UnsetRecipientMoved()`

UnsetRecipientMoved ensures that no value is present for RecipientMoved, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


