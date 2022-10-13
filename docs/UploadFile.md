# UploadFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | [default to "File uploaded successfully"]
**Filename** | **string** |  | 

## Methods

### NewUploadFile

`func NewUploadFile(message string, filename string, ) *UploadFile`

NewUploadFile instantiates a new UploadFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFileWithDefaults

`func NewUploadFileWithDefaults() *UploadFile`

NewUploadFileWithDefaults instantiates a new UploadFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UploadFile) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UploadFile) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UploadFile) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetFilename

`func (o *UploadFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadFile) SetFilename(v string)`

SetFilename sets Filename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


