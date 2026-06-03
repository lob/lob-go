# BankAccountVerify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amounts** | **[]int32** | In live mode, an array containing the two micro deposits (in cents) placed in the bank account. In test mode, no micro deposits will be placed, so any two integers between &#x60;1&#x60; and &#x60;100&#x60; will work. Required when microdeposit_type is &#x60;amounts&#x60;. | [optional]
**DescriptorCode** | Pointer to **string** | The 6-character code (beginning with &#x60;SM&#x60;) from the bank statement descriptor of the single $0.01 microdeposit. Required when microdeposit_type is &#x60;descriptor_code&#x60;. Must match &#x60;^SM[a-zA-Z0-9]{4}$&#x60;. | [optional]

Exactly one of `Amounts` or `DescriptorCode` must be provided.

## Methods

### NewBankAccountVerify

`func NewBankAccountVerify(amounts []int32) *BankAccountVerify`

NewBankAccountVerify instantiates a new BankAccountVerify object using the amounts verification path.

### NewBankAccountVerifyWithDescriptorCode

`func NewBankAccountVerifyWithDescriptorCode(descriptorCode string) *BankAccountVerify`

NewBankAccountVerifyWithDescriptorCode instantiates a new BankAccountVerify object using the descriptor code verification path.

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

### HasAmounts

`func (o *BankAccountVerify) HasAmounts() bool`

HasAmounts returns a boolean if a field has been set.

### SetAmounts

`func (o *BankAccountVerify) SetAmounts(v []int32)`

SetAmounts sets Amounts field to given value.

### GetDescriptorCode

`func (o *BankAccountVerify) GetDescriptorCode() string`

GetDescriptorCode returns the DescriptorCode field if non-nil, zero value otherwise.

### GetDescriptorCodeOk

`func (o *BankAccountVerify) GetDescriptorCodeOk() (*string, bool)`

GetDescriptorCodeOk returns a tuple with the DescriptorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescriptorCode

`func (o *BankAccountVerify) HasDescriptorCode() bool`

HasDescriptorCode returns a boolean if a field has been set.

### SetDescriptorCode

`func (o *BankAccountVerify) SetDescriptorCode(v string) error`

SetDescriptorCode sets DescriptorCode field to given value. Returns an error if the code does not match `^SM[a-zA-Z0-9]{4}$`.

### Validate

`func (o *BankAccountVerify) Validate() error`

Validate checks that exactly one of Amounts or DescriptorCode is set and that constraints are met.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


