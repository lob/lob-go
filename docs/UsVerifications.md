# UsVerifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]UsVerificationOrError**](UsVerificationOrError.md) |  | 
**Errors** | **bool** | Indicates whether any errors occurred during the verification process. | 

## Methods

### NewUsVerifications

`func NewUsVerifications(addresses []UsVerificationOrError, errors bool, ) *UsVerifications`

NewUsVerifications instantiates a new UsVerifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsVerificationsWithDefaults

`func NewUsVerificationsWithDefaults() *UsVerifications`

NewUsVerificationsWithDefaults instantiates a new UsVerifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *UsVerifications) GetAddresses() []UsVerificationOrError`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *UsVerifications) GetAddressesOk() (*[]UsVerificationOrError, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *UsVerifications) SetAddresses(v []UsVerificationOrError)`

SetAddresses sets Addresses field to given value.


### GetErrors

`func (o *UsVerifications) GetErrors() bool`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UsVerifications) GetErrorsOk() (*bool, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UsVerifications) SetErrors(v bool)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


