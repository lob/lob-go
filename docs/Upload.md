# Upload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;upl_&#x60;. | 
**AccountId** | **string** | Account ID that made the request | 
**Mode** | **string** | The environment in which the mailpieces were created. Today, will only be &#x60;live&#x60;. | 
**CampaignId** | **string** | Campaign ID associated with the upload | 
**FailuresUrl** | Pointer to **string** | Url where your campaign mailpiece failures can be retrieved | [optional] 
**OriginalFilename** | Pointer to **string** | Filename of the upload | [optional] 
**State** | [**UploadState**](UploadState.md) |  | [default to UPLOADSTATE_DRAFT]
**TotalMailpieces** | **int32** | Total number of recipients for the campaign | 
**FailedMailpieces** | **int32** | Number of mailpieces that failed to create | 
**ValidatedMailpieces** | **int32** | Number of mailpieces that were successfully created | 
**BytesProcessed** | **int32** | Number of bytes processed in your CSV | 
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the upload was created | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the upload was last modified | 
**RequiredAddressColumnMapping** | [**RequiredAddressColumnMapping**](RequiredAddressColumnMapping.md) |  | 
**OptionalAddressColumnMapping** | [**OptionalAddressColumnMapping**](OptionalAddressColumnMapping.md) |  | 
**Metadata** | [**UploadsMetadata**](UploadsMetadata.md) |  | 
**MergeVariableColumnMapping** | **map[string]interface{}** | The mapping of column headers in your file to the merge variables present in your creative. See our &lt;a href&#x3D;\&quot;https://help.lob.com/print-and-mail/building-a-mail-strategy/campaign-or-triggered-sends/campaign-audience-guide#step-3-map-merge-variable-data-if-applicable-7\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Campaign Audience Guide&lt;/a&gt; for additional details. &lt;br /&gt;If a merge variable has the same \&quot;name\&quot; as a \&quot;key\&quot; in the &#x60;requiredAddressColumnMapping&#x60; or &#x60;optionalAddressColumnMapping&#x60; objects, then they **CANNOT** have a different value in this object. If a different value is provided, then when the campaign is processing it will get overwritten with the mapped value present in the &#x60;requiredAddressColumnMapping&#x60; or &#x60;optionalAddressColumnMapping&#x60; objects. | 

## Methods

### NewUpload

`func NewUpload(id string, accountId string, mode string, campaignId string, state UploadState, totalMailpieces int32, failedMailpieces int32, validatedMailpieces int32, bytesProcessed int32, dateCreated time.Time, dateModified time.Time, requiredAddressColumnMapping RequiredAddressColumnMapping, optionalAddressColumnMapping OptionalAddressColumnMapping, metadata UploadsMetadata, mergeVariableColumnMapping map[string]interface{}, ) *Upload`

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


### GetRequiredAddressColumnMapping

`func (o *Upload) GetRequiredAddressColumnMapping() RequiredAddressColumnMapping`

GetRequiredAddressColumnMapping returns the RequiredAddressColumnMapping field if non-nil, zero value otherwise.

### GetRequiredAddressColumnMappingOk

`func (o *Upload) GetRequiredAddressColumnMappingOk() (*RequiredAddressColumnMapping, bool)`

GetRequiredAddressColumnMappingOk returns a tuple with the RequiredAddressColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAddressColumnMapping

`func (o *Upload) SetRequiredAddressColumnMapping(v RequiredAddressColumnMapping)`

SetRequiredAddressColumnMapping sets RequiredAddressColumnMapping field to given value.


### GetOptionalAddressColumnMapping

`func (o *Upload) GetOptionalAddressColumnMapping() OptionalAddressColumnMapping`

GetOptionalAddressColumnMapping returns the OptionalAddressColumnMapping field if non-nil, zero value otherwise.

### GetOptionalAddressColumnMappingOk

`func (o *Upload) GetOptionalAddressColumnMappingOk() (*OptionalAddressColumnMapping, bool)`

GetOptionalAddressColumnMappingOk returns a tuple with the OptionalAddressColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAddressColumnMapping

`func (o *Upload) SetOptionalAddressColumnMapping(v OptionalAddressColumnMapping)`

SetOptionalAddressColumnMapping sets OptionalAddressColumnMapping field to given value.


### GetMetadata

`func (o *Upload) GetMetadata() UploadsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Upload) GetMetadataOk() (*UploadsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Upload) SetMetadata(v UploadsMetadata)`

SetMetadata sets Metadata field to given value.


### GetMergeVariableColumnMapping

`func (o *Upload) GetMergeVariableColumnMapping() map[string]interface{}`

GetMergeVariableColumnMapping returns the MergeVariableColumnMapping field if non-nil, zero value otherwise.

### GetMergeVariableColumnMappingOk

`func (o *Upload) GetMergeVariableColumnMappingOk() (*map[string]interface{}, bool)`

GetMergeVariableColumnMappingOk returns a tuple with the MergeVariableColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariableColumnMapping

`func (o *Upload) SetMergeVariableColumnMapping(v map[string]interface{})`

SetMergeVariableColumnMapping sets MergeVariableColumnMapping field to given value.


### SetMergeVariableColumnMappingNil

`func (o *Upload) SetMergeVariableColumnMappingNil(b bool)`

 SetMergeVariableColumnMappingNil sets the value for MergeVariableColumnMapping to be an explicit nil

### UnsetMergeVariableColumnMapping
`func (o *Upload) UnsetMergeVariableColumnMapping()`

UnsetMergeVariableColumnMapping ensures that no value is present for MergeVariableColumnMapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


