# LobError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | A human-readable message with more details about the error | [optional] 
**StatusCode** | Pointer to **int32** | A conventional HTTP status code. | [optional] 
**Code** | Pointer to **string** | A pre-defined string identifying an error. | [optional] 

## Methods

### NewLobError

`func NewLobError() *LobError`

NewLobError instantiates a new LobError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobErrorWithDefaults

`func NewLobErrorWithDefaults() *LobError`

NewLobErrorWithDefaults instantiates a new LobError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *LobError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LobError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LobError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LobError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *LobError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *LobError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *LobError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *LobError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetCode

`func (o *LobError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LobError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LobError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LobError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


