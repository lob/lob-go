# Postcard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;psc_&#x60;. | 
**To** | Pointer to [**Address**](Address.md) |  | [optional] 
**From** | Pointer to [**AddressDomesticExpanded**](AddressDomesticExpanded.md) |  | [optional] 
**Carrier** | Pointer to **string** |  | [optional] [default to "USPS"]
**Thumbnails** | Pointer to [**[]Thumbnail**](Thumbnail.md) |  | [optional] 
**Size** | Pointer to [**PostcardSize**](PostcardSize.md) |  | [optional] [default to _4X6]
**ExpectedDeliveryDate** | Pointer to **string** | A date in YYYY-MM-DD format of the mailpiece&#39;s expected delivery date based on its &#x60;send_date&#x60;. | [optional] 
**DateCreated** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | [optional] 
**DateModified** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | [optional] 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**FrontTemplateId** | Pointer to **NullableString** | The unique ID of the HTML template used for the front of the postcard. | [optional] 
**BackTemplateId** | Pointer to **NullableString** | The unique ID of the HTML template used for the back of the postcard. | [optional] 
**FrontTemplateVersionId** | Pointer to **NullableString** | The unique ID of the specific version of the HTML template used for the front of the postcard. | [optional] 
**BackTemplateVersionId** | Pointer to **NullableString** | The unique ID of the specific version of the HTML template used for the back of the postcard. | [optional] 
**TrackingEvents** | Pointer to [**[]TrackingEventNormal**](TrackingEventNormal.md) | An array of tracking_event objects ordered by ascending &#x60;time&#x60;. Will not be populated for postcards created in test mode. | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "postcard"]
**Url** | **string** | A [signed link](#section/Asset-URLs) served over HTTPS. The link returned will expire in 30 days to prevent mis-sharing. Each time a GET request is initiated, a new signed URL will be generated. | 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to FIRST_CLASS]
**MergeVariables** | Pointer to **map[string]interface{}** | You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: &#x60;{{variable_name}}&#x60;, pass in &#x60;{\&quot;variable_name\&quot;: \&quot;Harry\&quot;}&#x60; to render &#x60;Harry&#x60;. &#x60;merge_variables&#x60; must be an object. Any type of value is accepted as long as the object is valid JSON; you can use &#x60;strings&#x60;, &#x60;numbers&#x60;, &#x60;booleans&#x60;, &#x60;arrays&#x60;, &#x60;objects&#x60;, or &#x60;null&#x60;. The max length of the object is 25,000 characters. If you call &#x60;JSON.stringify&#x60; on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: &#x60;!&#x60;, &#x60;\&quot;&#x60;, &#x60;#&#x60;, &#x60;%&#x60;, &#x60;&amp;&#x60;, &#x60;&#39;&#x60;, &#x60;(&#x60;, &#x60;)&#x60;, &#x60;*&#x60;, &#x60;+&#x60;, &#x60;,&#x60;, &#x60;/&#x60;, &#x60;;&#x60;, &#x60;&lt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;@&#x60;, &#x60;[&#x60;, &#x60;\\&#x60;, &#x60;]&#x60;, &#x60;^&#x60;, &#x60;&#x60; &#x60; &#x60;&#x60;, &#x60;{&#x60;, &#x60;|&#x60;, &#x60;}&#x60;, &#x60;~&#x60;. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string. | [optional] 
**SendDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the &#x60;send_date&#x60; has passed, the mailpiece can be canceled. If a date in the format &#x60;2017-11-01&#x60; is passed, it will evaluate to midnight UTC of that date (&#x60;2017-11-01T00:00:00.000Z&#x60;). If a datetime is passed, that exact time will be used. A &#x60;send_date&#x60; passed with no time zone will default to UTC, while a &#x60;send_date&#x60; passed with a time zone will be converted to UTC. | [optional] 

## Methods

### NewPostcard

`func NewPostcard(id string, url string, ) *Postcard`

NewPostcard instantiates a new Postcard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostcardWithDefaults

`func NewPostcardWithDefaults() *Postcard`

NewPostcardWithDefaults instantiates a new Postcard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Postcard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Postcard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Postcard) SetId(v string)`

SetId sets Id field to given value.


### GetTo

`func (o *Postcard) GetTo() Address`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Postcard) GetToOk() (*Address, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Postcard) SetTo(v Address)`

SetTo sets To field to given value.

### HasTo

`func (o *Postcard) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *Postcard) GetFrom() AddressDomesticExpanded`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Postcard) GetFromOk() (*AddressDomesticExpanded, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Postcard) SetFrom(v AddressDomesticExpanded)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Postcard) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetCarrier

`func (o *Postcard) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *Postcard) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *Postcard) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *Postcard) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetThumbnails

`func (o *Postcard) GetThumbnails() []Thumbnail`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *Postcard) GetThumbnailsOk() (*[]Thumbnail, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *Postcard) SetThumbnails(v []Thumbnail)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *Postcard) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### GetSize

`func (o *Postcard) GetSize() PostcardSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Postcard) GetSizeOk() (*PostcardSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Postcard) SetSize(v PostcardSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *Postcard) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetExpectedDeliveryDate

`func (o *Postcard) GetExpectedDeliveryDate() string`

GetExpectedDeliveryDate returns the ExpectedDeliveryDate field if non-nil, zero value otherwise.

### GetExpectedDeliveryDateOk

`func (o *Postcard) GetExpectedDeliveryDateOk() (*string, bool)`

GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDeliveryDate

`func (o *Postcard) SetExpectedDeliveryDate(v string)`

SetExpectedDeliveryDate sets ExpectedDeliveryDate field to given value.

### HasExpectedDeliveryDate

`func (o *Postcard) HasExpectedDeliveryDate() bool`

HasExpectedDeliveryDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *Postcard) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Postcard) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Postcard) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Postcard) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateModified

`func (o *Postcard) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *Postcard) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *Postcard) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.

### HasDateModified

`func (o *Postcard) HasDateModified() bool`

HasDateModified returns a boolean if a field has been set.

### GetDeleted

`func (o *Postcard) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Postcard) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Postcard) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Postcard) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFrontTemplateId

`func (o *Postcard) GetFrontTemplateId() string`

GetFrontTemplateId returns the FrontTemplateId field if non-nil, zero value otherwise.

### GetFrontTemplateIdOk

`func (o *Postcard) GetFrontTemplateIdOk() (*string, bool)`

GetFrontTemplateIdOk returns a tuple with the FrontTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontTemplateId

`func (o *Postcard) SetFrontTemplateId(v string)`

SetFrontTemplateId sets FrontTemplateId field to given value.

### HasFrontTemplateId

`func (o *Postcard) HasFrontTemplateId() bool`

HasFrontTemplateId returns a boolean if a field has been set.

### SetFrontTemplateIdNil

`func (o *Postcard) SetFrontTemplateIdNil(b bool)`

 SetFrontTemplateIdNil sets the value for FrontTemplateId to be an explicit nil

### UnsetFrontTemplateId
`func (o *Postcard) UnsetFrontTemplateId()`

UnsetFrontTemplateId ensures that no value is present for FrontTemplateId, not even an explicit nil
### GetBackTemplateId

`func (o *Postcard) GetBackTemplateId() string`

GetBackTemplateId returns the BackTemplateId field if non-nil, zero value otherwise.

### GetBackTemplateIdOk

`func (o *Postcard) GetBackTemplateIdOk() (*string, bool)`

GetBackTemplateIdOk returns a tuple with the BackTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackTemplateId

`func (o *Postcard) SetBackTemplateId(v string)`

SetBackTemplateId sets BackTemplateId field to given value.

### HasBackTemplateId

`func (o *Postcard) HasBackTemplateId() bool`

HasBackTemplateId returns a boolean if a field has been set.

### SetBackTemplateIdNil

`func (o *Postcard) SetBackTemplateIdNil(b bool)`

 SetBackTemplateIdNil sets the value for BackTemplateId to be an explicit nil

### UnsetBackTemplateId
`func (o *Postcard) UnsetBackTemplateId()`

UnsetBackTemplateId ensures that no value is present for BackTemplateId, not even an explicit nil
### GetFrontTemplateVersionId

`func (o *Postcard) GetFrontTemplateVersionId() string`

GetFrontTemplateVersionId returns the FrontTemplateVersionId field if non-nil, zero value otherwise.

### GetFrontTemplateVersionIdOk

`func (o *Postcard) GetFrontTemplateVersionIdOk() (*string, bool)`

GetFrontTemplateVersionIdOk returns a tuple with the FrontTemplateVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontTemplateVersionId

`func (o *Postcard) SetFrontTemplateVersionId(v string)`

SetFrontTemplateVersionId sets FrontTemplateVersionId field to given value.

### HasFrontTemplateVersionId

`func (o *Postcard) HasFrontTemplateVersionId() bool`

HasFrontTemplateVersionId returns a boolean if a field has been set.

### SetFrontTemplateVersionIdNil

`func (o *Postcard) SetFrontTemplateVersionIdNil(b bool)`

 SetFrontTemplateVersionIdNil sets the value for FrontTemplateVersionId to be an explicit nil

### UnsetFrontTemplateVersionId
`func (o *Postcard) UnsetFrontTemplateVersionId()`

UnsetFrontTemplateVersionId ensures that no value is present for FrontTemplateVersionId, not even an explicit nil
### GetBackTemplateVersionId

`func (o *Postcard) GetBackTemplateVersionId() string`

GetBackTemplateVersionId returns the BackTemplateVersionId field if non-nil, zero value otherwise.

### GetBackTemplateVersionIdOk

`func (o *Postcard) GetBackTemplateVersionIdOk() (*string, bool)`

GetBackTemplateVersionIdOk returns a tuple with the BackTemplateVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackTemplateVersionId

`func (o *Postcard) SetBackTemplateVersionId(v string)`

SetBackTemplateVersionId sets BackTemplateVersionId field to given value.

### HasBackTemplateVersionId

`func (o *Postcard) HasBackTemplateVersionId() bool`

HasBackTemplateVersionId returns a boolean if a field has been set.

### SetBackTemplateVersionIdNil

`func (o *Postcard) SetBackTemplateVersionIdNil(b bool)`

 SetBackTemplateVersionIdNil sets the value for BackTemplateVersionId to be an explicit nil

### UnsetBackTemplateVersionId
`func (o *Postcard) UnsetBackTemplateVersionId()`

UnsetBackTemplateVersionId ensures that no value is present for BackTemplateVersionId, not even an explicit nil
### GetTrackingEvents

`func (o *Postcard) GetTrackingEvents() []TrackingEventNormal`

GetTrackingEvents returns the TrackingEvents field if non-nil, zero value otherwise.

### GetTrackingEventsOk

`func (o *Postcard) GetTrackingEventsOk() (*[]TrackingEventNormal, bool)`

GetTrackingEventsOk returns a tuple with the TrackingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingEvents

`func (o *Postcard) SetTrackingEvents(v []TrackingEventNormal)`

SetTrackingEvents sets TrackingEvents field to given value.

### HasTrackingEvents

`func (o *Postcard) HasTrackingEvents() bool`

HasTrackingEvents returns a boolean if a field has been set.

### SetTrackingEventsNil

`func (o *Postcard) SetTrackingEventsNil(b bool)`

 SetTrackingEventsNil sets the value for TrackingEvents to be an explicit nil

### UnsetTrackingEvents
`func (o *Postcard) UnsetTrackingEvents()`

UnsetTrackingEvents ensures that no value is present for TrackingEvents, not even an explicit nil
### GetObject

`func (o *Postcard) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Postcard) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Postcard) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Postcard) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUrl

`func (o *Postcard) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Postcard) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Postcard) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *Postcard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Postcard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Postcard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Postcard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Postcard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Postcard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *Postcard) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Postcard) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Postcard) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Postcard) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMailType

`func (o *Postcard) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *Postcard) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *Postcard) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *Postcard) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMergeVariables

`func (o *Postcard) GetMergeVariables() map[string]interface{}`

GetMergeVariables returns the MergeVariables field if non-nil, zero value otherwise.

### GetMergeVariablesOk

`func (o *Postcard) GetMergeVariablesOk() (*map[string]interface{}, bool)`

GetMergeVariablesOk returns a tuple with the MergeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariables

`func (o *Postcard) SetMergeVariables(v map[string]interface{})`

SetMergeVariables sets MergeVariables field to given value.

### HasMergeVariables

`func (o *Postcard) HasMergeVariables() bool`

HasMergeVariables returns a boolean if a field has been set.

### SetMergeVariablesNil

`func (o *Postcard) SetMergeVariablesNil(b bool)`

 SetMergeVariablesNil sets the value for MergeVariables to be an explicit nil

### UnsetMergeVariables
`func (o *Postcard) UnsetMergeVariables()`

UnsetMergeVariables ensures that no value is present for MergeVariables, not even an explicit nil
### GetSendDate

`func (o *Postcard) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *Postcard) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *Postcard) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *Postcard) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


