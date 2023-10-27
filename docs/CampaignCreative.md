# CampaignCreative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;crv_&#x60;. | [optional] 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**From** | Pointer to **interface{}** | Must either be an address ID or an inline object with correct address parameters. | [optional] 
**ResourceType** | Pointer to **string** | Mailpiece type for the creative | [optional] 
**Details** | Pointer to **interface{}** | Either PostcardDetailsReturned or LetterDetailsReturned | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**TemplatePreviewUrls** | Pointer to **map[string]interface{}** | Preview URLs associated with a creative&#39;s artwork asset(s) if the creative uses HTML templates as assets. | [optional] 
**TemplatePreviews** | Pointer to **[]map[string]interface{}** | A list of template preview objects if the creative uses HTML template(s) as artwork asset(s). | [optional] 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Campaigns** | Pointer to [**[]Campaign**](Campaign.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | [optional] 
**DateModified** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | [optional] 
**Object** | Pointer to **string** | Value is resource type. | [optional] [default to "creative"]

## Methods

### NewCampaignCreative

`func NewCampaignCreative() *CampaignCreative`

NewCampaignCreative instantiates a new CampaignCreative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCreativeWithDefaults

`func NewCampaignCreativeWithDefaults() *CampaignCreative`

NewCampaignCreativeWithDefaults instantiates a new CampaignCreative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignCreative) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignCreative) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignCreative) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignCreative) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *CampaignCreative) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignCreative) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignCreative) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignCreative) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CampaignCreative) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CampaignCreative) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFrom

`func (o *CampaignCreative) GetFrom() interface{}`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CampaignCreative) GetFromOk() (*interface{}, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CampaignCreative) SetFrom(v interface{})`

SetFrom sets From field to given value.

### HasFrom

`func (o *CampaignCreative) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *CampaignCreative) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *CampaignCreative) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetResourceType

`func (o *CampaignCreative) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CampaignCreative) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CampaignCreative) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *CampaignCreative) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetDetails

`func (o *CampaignCreative) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CampaignCreative) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CampaignCreative) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CampaignCreative) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *CampaignCreative) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *CampaignCreative) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetMetadata

`func (o *CampaignCreative) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CampaignCreative) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CampaignCreative) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CampaignCreative) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTemplatePreviewUrls

`func (o *CampaignCreative) GetTemplatePreviewUrls() map[string]interface{}`

GetTemplatePreviewUrls returns the TemplatePreviewUrls field if non-nil, zero value otherwise.

### GetTemplatePreviewUrlsOk

`func (o *CampaignCreative) GetTemplatePreviewUrlsOk() (*map[string]interface{}, bool)`

GetTemplatePreviewUrlsOk returns a tuple with the TemplatePreviewUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePreviewUrls

`func (o *CampaignCreative) SetTemplatePreviewUrls(v map[string]interface{})`

SetTemplatePreviewUrls sets TemplatePreviewUrls field to given value.

### HasTemplatePreviewUrls

`func (o *CampaignCreative) HasTemplatePreviewUrls() bool`

HasTemplatePreviewUrls returns a boolean if a field has been set.

### GetTemplatePreviews

`func (o *CampaignCreative) GetTemplatePreviews() []map[string]interface{}`

GetTemplatePreviews returns the TemplatePreviews field if non-nil, zero value otherwise.

### GetTemplatePreviewsOk

`func (o *CampaignCreative) GetTemplatePreviewsOk() (*[]map[string]interface{}, bool)`

GetTemplatePreviewsOk returns a tuple with the TemplatePreviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePreviews

`func (o *CampaignCreative) SetTemplatePreviews(v []map[string]interface{})`

SetTemplatePreviews sets TemplatePreviews field to given value.

### HasTemplatePreviews

`func (o *CampaignCreative) HasTemplatePreviews() bool`

HasTemplatePreviews returns a boolean if a field has been set.

### GetDeleted

`func (o *CampaignCreative) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CampaignCreative) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CampaignCreative) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CampaignCreative) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetCampaigns

`func (o *CampaignCreative) GetCampaigns() []Campaign`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *CampaignCreative) GetCampaignsOk() (*[]Campaign, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *CampaignCreative) SetCampaigns(v []Campaign)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *CampaignCreative) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetDateCreated

`func (o *CampaignCreative) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CampaignCreative) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CampaignCreative) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CampaignCreative) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateModified

`func (o *CampaignCreative) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *CampaignCreative) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *CampaignCreative) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.

### HasDateModified

`func (o *CampaignCreative) HasDateModified() bool`

HasDateModified returns a boolean if a field has been set.

### GetObject

`func (o *CampaignCreative) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CampaignCreative) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CampaignCreative) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CampaignCreative) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


