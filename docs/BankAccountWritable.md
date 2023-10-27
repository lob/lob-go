# BankAccountWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**RoutingNumber** | **string** | Must be a [valid US routing number](https://www.frbservices.org/index.html). | 
**AccountNumber** | **string** |  | 
**AccountType** | [**BankTypeEnum**](BankTypeEnum.md) |  | 
**Signatory** | **string** | The signatory associated with your account. This name will be printed on checks created with this bank account. If you prefer to use a custom signature image on your checks instead, please create your bank account from the [Dashboard](https://dashboard.lob.com/#/login). | 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 

## Methods

### NewBankAccountWritable

`func NewBankAccountWritable(routingNumber string, accountNumber string, accountType BankTypeEnum, signatory string, ) *BankAccountWritable`

NewBankAccountWritable instantiates a new BankAccountWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountWritableWithDefaults

`func NewBankAccountWritableWithDefaults() *BankAccountWritable`

NewBankAccountWritableWithDefaults instantiates a new BankAccountWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BankAccountWritable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankAccountWritable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankAccountWritable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BankAccountWritable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BankAccountWritable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BankAccountWritable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRoutingNumber

`func (o *BankAccountWritable) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankAccountWritable) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankAccountWritable) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetAccountNumber

`func (o *BankAccountWritable) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankAccountWritable) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankAccountWritable) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountType

`func (o *BankAccountWritable) GetAccountType() BankTypeEnum`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountWritable) GetAccountTypeOk() (*BankTypeEnum, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountWritable) SetAccountType(v BankTypeEnum)`

SetAccountType sets AccountType field to given value.


### GetSignatory

`func (o *BankAccountWritable) GetSignatory() string`

GetSignatory returns the Signatory field if non-nil, zero value otherwise.

### GetSignatoryOk

`func (o *BankAccountWritable) GetSignatoryOk() (*string, bool)`

GetSignatoryOk returns a tuple with the Signatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatory

`func (o *BankAccountWritable) SetSignatory(v string)`

SetSignatory sets Signatory field to given value.


### GetMetadata

`func (o *BankAccountWritable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BankAccountWritable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BankAccountWritable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BankAccountWritable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


