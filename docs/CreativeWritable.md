# CreativeWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **interface{}** | Must either be an address ID or an inline object with correct address parameters. | 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**ResourceType** | **string** | Mailpiece type for the creative | 
**CampaignId** | **string** | Unique identifier prefixed with &#x60;cmp_&#x60;. | 
**Details** | Pointer to **interface{}** | Either PostcardDetailsWritable or LetterDetailsWritable | [optional] 
**File** | Pointer to **string** | PDF file containing the letter&#39;s formatting. Do not include for resource_type &#x3D; postcard. | [optional] 
**Front** | Pointer to **string** | The artwork to use as the front of your postcard. Do not include for resource_type &#x3D; letter.  | [optional] 
**Back** | Pointer to **string** | The artwork to use as the back of your postcard. Do not include for resource_type &#x3D; letter.  | [optional] 

## Methods

### NewCreativeWritable

`func NewCreativeWritable(from interface{}, resourceType string, campaignId string, ) *CreativeWritable`

NewCreativeWritable instantiates a new CreativeWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeWritableWithDefaults

`func NewCreativeWritableWithDefaults() *CreativeWritable`

NewCreativeWritableWithDefaults instantiates a new CreativeWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CreativeWritable) GetFrom() interface{}`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreativeWritable) GetFromOk() (*interface{}, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreativeWritable) SetFrom(v interface{})`

SetFrom sets From field to given value.


### SetFromNil

`func (o *CreativeWritable) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *CreativeWritable) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetDescription

`func (o *CreativeWritable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreativeWritable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreativeWritable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreativeWritable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreativeWritable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreativeWritable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *CreativeWritable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreativeWritable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreativeWritable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreativeWritable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResourceType

`func (o *CreativeWritable) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreativeWritable) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreativeWritable) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetCampaignId

`func (o *CreativeWritable) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CreativeWritable) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CreativeWritable) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetDetails

`func (o *CreativeWritable) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CreativeWritable) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CreativeWritable) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CreativeWritable) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *CreativeWritable) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *CreativeWritable) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetFile

`func (o *CreativeWritable) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreativeWritable) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreativeWritable) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *CreativeWritable) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFront

`func (o *CreativeWritable) GetFront() string`

GetFront returns the Front field if non-nil, zero value otherwise.

### GetFrontOk

`func (o *CreativeWritable) GetFrontOk() (*string, bool)`

GetFrontOk returns a tuple with the Front field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFront

`func (o *CreativeWritable) SetFront(v string)`

SetFront sets Front field to given value.

### HasFront

`func (o *CreativeWritable) HasFront() bool`

HasFront returns a boolean if a field has been set.

### GetBack

`func (o *CreativeWritable) GetBack() string`

GetBack returns the Back field if non-nil, zero value otherwise.

### GetBackOk

`func (o *CreativeWritable) GetBackOk() (*string, bool)`

GetBackOk returns a tuple with the Back field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBack

`func (o *CreativeWritable) SetBack(v string)`

SetBack sets Back field to given value.

### HasBack

`func (o *CreativeWritable) HasBack() bool`

HasBack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


