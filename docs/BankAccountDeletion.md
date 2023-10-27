# BankAccountDeletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;bank_&#x60;. | [optional] 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Object** | Pointer to **string** | Value is type of resource. | [optional] [default to "bank_account_deleted"]

## Methods

### NewBankAccountDeletion

`func NewBankAccountDeletion() *BankAccountDeletion`

NewBankAccountDeletion instantiates a new BankAccountDeletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountDeletionWithDefaults

`func NewBankAccountDeletionWithDefaults() *BankAccountDeletion`

NewBankAccountDeletionWithDefaults instantiates a new BankAccountDeletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankAccountDeletion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankAccountDeletion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankAccountDeletion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankAccountDeletion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *BankAccountDeletion) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *BankAccountDeletion) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *BankAccountDeletion) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *BankAccountDeletion) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetObject

`func (o *BankAccountDeletion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BankAccountDeletion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BankAccountDeletion) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BankAccountDeletion) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


