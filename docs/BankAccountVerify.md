# BankAccountVerify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amounts** | **[]int32** | In live mode, an array containing the two micro deposits (in cents) placed in the bank account. In test mode, no micro deposits will be placed, so any two integers between &#x60;1&#x60; and &#x60;100&#x60; will work. | 

## Methods

### NewBankAccountVerify

`func NewBankAccountVerify(amounts []int32, ) *BankAccountVerify`

NewBankAccountVerify instantiates a new BankAccountVerify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountVerifyWithDefaults

`func NewBankAccountVerifyWithDefaults() *BankAccountVerify`

NewBankAccountVerifyWithDefaults instantiates a new BankAccountVerify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmounts

`func (o *BankAccountVerify) GetAmounts() []int32`

GetAmounts returns the Amounts field if non-nil, zero value otherwise.

### GetAmountsOk

`func (o *BankAccountVerify) GetAmountsOk() (*[]int32, bool)`

GetAmountsOk returns a tuple with the Amounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmounts

`func (o *BankAccountVerify) SetAmounts(v []int32)`

SetAmounts sets Amounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


