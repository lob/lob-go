# Upload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;upl_&#x60;. | 
**AccountId** | **string** | Account ID that made the request | 
**CampaignId** | **string** | Unique identifier prefixed with &#x60;cmp_&#x60;. | 
**ColumnMapping** | **map[string]interface{}** | The mapping of column headers in your file to Lob-required fields for the resource created. See our &lt;a href&#x3D;\&quot;https://help.lob.com/best-practices/campaign-audience-guide\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Campaign Audience Guide&lt;/a&gt; for additional details. | 
**Mode** | **string** | The environment in which the mailpieces were created. Today, will only be &#x60;live&#x60;. | 
**FailuresUrl** | Pointer to **string** | Url where your campaign mailpiece failures can be retrieved | [optional] 
**OriginalFilename** | Pointer to **string** | Filename of the upload | [optional] 
**State** | [**UploadState**](UploadState.md) |  | [default to DRAFT]
**TotalMailpieces** | **int32** | Total number of recipients for the campaign | 
**FailedMailpieces** | **int32** | Number of mailpieces that failed to create | 
**ValidatedMailpieces** | **int32** | Number of mailpieces that were successfully created | 
**BytesProcessed** | **int32** | Number of bytes processed in your CSV | 
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the upload was created | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the upload was last modified | 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 

## Methods

### NewUpload

`func NewUpload(id string, accountId string, campaignId string, columnMapping map[string]interface{}, mode string, state UploadState, totalMailpieces int32, failedMailpieces int32, validatedMailpieces int32, bytesProcessed int32, dateCreated time.Time, dateModified time.Time, ) *Upload`

NewUpload instantiates a new Upload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadWithDefaults

`func NewUploadWithDefaults() *Upload`

NewUploadWithDefaults instantiates a new Upload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Upload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Upload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Upload) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Upload) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Upload) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Upload) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCampaignId

`func (o *Upload) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Upload) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *Upload) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetColumnMapping

`func (o *Upload) GetColumnMapping() map[string]interface{}`

GetColumnMapping returns the ColumnMapping field if non-nil, zero value otherwise.

### GetColumnMappingOk

`func (o *Upload) GetColumnMappingOk() (*map[string]interface{}, bool)`

GetColumnMappingOk returns a tuple with the ColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnMapping

`func (o *Upload) SetColumnMapping(v map[string]interface{})`

SetColumnMapping sets ColumnMapping field to given value.


### GetMode

`func (o *Upload) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Upload) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Upload) SetMode(v string)`

SetMode sets Mode field to given value.


### GetFailuresUrl

`func (o *Upload) GetFailuresUrl() string`

GetFailuresUrl returns the FailuresUrl field if non-nil, zero value otherwise.

### GetFailuresUrlOk

`func (o *Upload) GetFailuresUrlOk() (*string, bool)`

GetFailuresUrlOk returns a tuple with the FailuresUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailuresUrl

`func (o *Upload) SetFailuresUrl(v string)`

SetFailuresUrl sets FailuresUrl field to given value.

### HasFailuresUrl

`func (o *Upload) HasFailuresUrl() bool`

HasFailuresUrl returns a boolean if a field has been set.

### GetOriginalFilename

`func (o *Upload) GetOriginalFilename() string`

GetOriginalFilename returns the OriginalFilename field if non-nil, zero value otherwise.

### GetOriginalFilenameOk

`func (o *Upload) GetOriginalFilenameOk() (*string, bool)`

GetOriginalFilenameOk returns a tuple with the OriginalFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFilename

`func (o *Upload) SetOriginalFilename(v string)`

SetOriginalFilename sets OriginalFilename field to given value.

### HasOriginalFilename

`func (o *Upload) HasOriginalFilename() bool`

HasOriginalFilename returns a boolean if a field has been set.

### GetState

`func (o *Upload) GetState() UploadState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Upload) GetStateOk() (*UploadState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Upload) SetState(v UploadState)`

SetState sets State field to given value.


### GetTotalMailpieces

`func (o *Upload) GetTotalMailpieces() int32`

GetTotalMailpieces returns the TotalMailpieces field if non-nil, zero value otherwise.

### GetTotalMailpiecesOk

`func (o *Upload) GetTotalMailpiecesOk() (*int32, bool)`

GetTotalMailpiecesOk returns a tuple with the TotalMailpieces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMailpieces

`func (o *Upload) SetTotalMailpieces(v int32)`

SetTotalMailpieces sets TotalMailpieces field to given value.


### GetFailedMailpieces

`func (o *Upload) GetFailedMailpieces() int32`

GetFailedMailpieces returns the FailedMailpieces field if non-nil, zero value otherwise.

### GetFailedMailpiecesOk

`func (o *Upload) GetFailedMailpiecesOk() (*int32, bool)`

GetFailedMailpiecesOk returns a tuple with the FailedMailpieces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMailpieces

`func (o *Upload) SetFailedMailpieces(v int32)`

SetFailedMailpieces sets FailedMailpieces field to given value.


### GetValidatedMailpieces

`func (o *Upload) GetValidatedMailpieces() int32`

GetValidatedMailpieces returns the ValidatedMailpieces field if non-nil, zero value otherwise.

### GetValidatedMailpiecesOk

`func (o *Upload) GetValidatedMailpiecesOk() (*int32, bool)`

GetValidatedMailpiecesOk returns a tuple with the ValidatedMailpieces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedMailpieces

`func (o *Upload) SetValidatedMailpieces(v int32)`

SetValidatedMailpieces sets ValidatedMailpieces field to given value.


### GetBytesProcessed

`func (o *Upload) GetBytesProcessed() int32`

GetBytesProcessed returns the BytesProcessed field if non-nil, zero value otherwise.

### GetBytesProcessedOk

`func (o *Upload) GetBytesProcessedOk() (*int32, bool)`

GetBytesProcessedOk returns a tuple with the BytesProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesProcessed

`func (o *Upload) SetBytesProcessed(v int32)`

SetBytesProcessed sets BytesProcessed field to given value.


### GetDateCreated

`func (o *Upload) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Upload) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Upload) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *Upload) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *Upload) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *Upload) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetDeleted

`func (o *Upload) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Upload) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Upload) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Upload) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


