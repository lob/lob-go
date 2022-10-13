# UploadWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **string** | Unique identifier prefixed with &#x60;cmp_&#x60;. | 
**ColumnMapping** | **map[string]interface{}** | The mapping of column headers in your file to Lob-required fields for the resource created. See our &lt;a href&#x3D;\&quot;https://help.lob.com/best-practices/campaign-audience-guide\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Campaign Audience Guide&lt;/a&gt; for additional details. | 

## Methods

### NewUploadWritable

`func NewUploadWritable(campaignId string, columnMapping map[string]interface{}, ) *UploadWritable`

NewUploadWritable instantiates a new UploadWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadWritableWithDefaults

`func NewUploadWritableWithDefaults() *UploadWritable`

NewUploadWritableWithDefaults instantiates a new UploadWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *UploadWritable) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *UploadWritable) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *UploadWritable) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetColumnMapping

`func (o *UploadWritable) GetColumnMapping() map[string]interface{}`

GetColumnMapping returns the ColumnMapping field if non-nil, zero value otherwise.

### GetColumnMappingOk

`func (o *UploadWritable) GetColumnMappingOk() (*map[string]interface{}, bool)`

GetColumnMappingOk returns a tuple with the ColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnMapping

`func (o *UploadWritable) SetColumnMapping(v map[string]interface{})`

SetColumnMapping sets ColumnMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


