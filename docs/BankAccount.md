# BankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**RoutingNumber** | **string** | Must be a [valid US routing number](https://www.frbservices.org/index.html). | 
**AccountNumber** | **string** |  | 
**AccountType** | **string** | The type of entity that holds the account. | 
**Signatory** | **string** | The signatory associated with your account. This name will be printed on checks created with this bank account. If you prefer to use a custom signature image on your checks instead, please create your bank account from the [Dashboard](https://dashboard.lob.com/#/login). | 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**Id** | **string** | Unique identifier prefixed with &#x60;bank_&#x60;. | 
**SignatureUrl** | Pointer to **NullableString** | A signed link to the signature image. will be generated. | [optional] 
**BankName** | Pointer to **string** | The name of the bank based on the provided routing number, e.g. &#x60;JPMORGAN CHASE BANK&#x60;. | [optional] 
**Verified** | Pointer to **bool** | A bank account must be verified before a check can be created. | [optional] [default to false]
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Object** | **string** |  | [default to "bank_account"]

## Methods

### NewBankAccount

`func NewBankAccount(routingNumber string, accountNumber string, accountType string, signatory string, id string, dateCreated time.Time, dateModified time.Time, object string, ) *BankAccount`

NewBankAccount instantiates a new BankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountWithDefaults

`func NewBankAccountWithDefaults() *BankAccount`

NewBankAccountWithDefaults instantiates a new BankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BankAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BankAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BankAccount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BankAccount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRoutingNumber

`func (o *BankAccount) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankAccount) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankAccount) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetAccountNumber

`func (o *BankAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountType

`func (o *BankAccount) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccount) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccount) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetSignatory

`func (o *BankAccount) GetSignatory() string`

GetSignatory returns the Signatory field if non-nil, zero value otherwise.

### GetSignatoryOk

`func (o *BankAccount) GetSignatoryOk() (*string, bool)`

GetSignatoryOk returns a tuple with the Signatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatory

`func (o *BankAccount) SetSignatory(v string)`

SetSignatory sets Signatory field to given value.


### GetMetadata

`func (o *BankAccount) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BankAccount) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BankAccount) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BankAccount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *BankAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankAccount) SetId(v string)`

SetId sets Id field to given value.


### GetSignatureUrl

`func (o *BankAccount) GetSignatureUrl() string`

GetSignatureUrl returns the SignatureUrl field if non-nil, zero value otherwise.

### GetSignatureUrlOk

`func (o *BankAccount) GetSignatureUrlOk() (*string, bool)`

GetSignatureUrlOk returns a tuple with the SignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureUrl

`func (o *BankAccount) SetSignatureUrl(v string)`

SetSignatureUrl sets SignatureUrl field to given value.

### HasSignatureUrl

`func (o *BankAccount) HasSignatureUrl() bool`

HasSignatureUrl returns a boolean if a field has been set.

### SetSignatureUrlNil

`func (o *BankAccount) SetSignatureUrlNil(b bool)`

 SetSignatureUrlNil sets the value for SignatureUrl to be an explicit nil

### UnsetSignatureUrl
`func (o *BankAccount) UnsetSignatureUrl()`

UnsetSignatureUrl ensures that no value is present for SignatureUrl, not even an explicit nil
### GetBankName

`func (o *BankAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *BankAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *BankAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *BankAccount) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetVerified

`func (o *BankAccount) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *BankAccount) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *BankAccount) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *BankAccount) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetDateCreated

`func (o *BankAccount) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BankAccount) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BankAccount) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *BankAccount) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *BankAccount) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *BankAccount) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetDeleted

`func (o *BankAccount) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *BankAccount) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *BankAccount) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *BankAccount) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetObject

`func (o *BankAccount) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BankAccount) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BankAccount) SetObject(v string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


