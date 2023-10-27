# BulkErrorProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | A human-readable message with more details about the error | [optional] 
**StatusCode** | Pointer to **int32** | A conventional HTTP status code. | [optional] 

## Methods

### NewBulkErrorProperties

`func NewBulkErrorProperties() *BulkErrorProperties`

NewBulkErrorProperties instantiates a new BulkErrorProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkErrorPropertiesWithDefaults

`func NewBulkErrorPropertiesWithDefaults() *BulkErrorProperties`

NewBulkErrorPropertiesWithDefaults instantiates a new BulkErrorProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BulkErrorProperties) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BulkErrorProperties) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BulkErrorProperties) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BulkErrorProperties) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *BulkErrorProperties) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BulkErrorProperties) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BulkErrorProperties) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BulkErrorProperties) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


