# Export

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;ex_&#x60;. | 
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the export was created | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the export was last modified | 
**Deleted** | **bool** | Returns as &#x60;true&#x60; if the resource has been successfully deleted. | 
**S3Url** | Pointer to **string** | The URL for the generated export file. | [optional] 
**State** | **string** | The state of the export file, which can be &#x60;in_progress&#x60;, &#x60;failed&#x60; or &#x60;succeeded&#x60;. | 
**Type** | **string** | The export file type, which can be &#x60;all&#x60;, &#x60;failures&#x60; or &#x60;successes&#x60;. | 
**UploadId** | **string** | Unique identifier prefixed with &#x60;upl_&#x60;. | 

## Methods

### NewExport

`func NewExport(id string, dateCreated time.Time, dateModified time.Time, deleted bool, state string, type_ string, uploadId string, ) *Export`

NewExport instantiates a new Export object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportWithDefaults

`func NewExportWithDefaults() *Export`

NewExportWithDefaults instantiates a new Export object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Export) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Export) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Export) SetId(v string)`

SetId sets Id field to given value.


### GetDateCreated

`func (o *Export) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Export) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Export) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *Export) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *Export) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *Export) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetDeleted

`func (o *Export) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Export) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Export) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetS3Url

`func (o *Export) GetS3Url() string`

GetS3Url returns the S3Url field if non-nil, zero value otherwise.

### GetS3UrlOk

`func (o *Export) GetS3UrlOk() (*string, bool)`

GetS3UrlOk returns a tuple with the S3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Url

`func (o *Export) SetS3Url(v string)`

SetS3Url sets S3Url field to given value.

### HasS3Url

`func (o *Export) HasS3Url() bool`

HasS3Url returns a boolean if a field has been set.

### GetState

`func (o *Export) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Export) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Export) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *Export) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Export) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Export) SetType(v string)`

SetType sets Type field to given value.


### GetUploadId

`func (o *Export) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *Export) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *Export) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


