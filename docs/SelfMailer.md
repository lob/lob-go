# SelfMailer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;sfm_&#x60;. | 
**To** | **interface{}** |  | 
**From** | Pointer to **interface{}** |  | [optional] 
**Size** | Pointer to [**SelfMailerSize**](SelfMailerSize.md) |  | [optional] [default to _6X18_BIFOLD]
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to FIRST_CLASS]
**MergeVariables** | Pointer to **map[string]interface{}** | You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: &#x60;{{variable_name}}&#x60;, pass in &#x60;{\&quot;variable_name\&quot;: \&quot;Harry\&quot;}&#x60; to render &#x60;Harry&#x60;. &#x60;merge_variables&#x60; must be an object. Any type of value is accepted as long as the object is valid JSON; you can use &#x60;strings&#x60;, &#x60;numbers&#x60;, &#x60;booleans&#x60;, &#x60;arrays&#x60;, &#x60;objects&#x60;, or &#x60;null&#x60;. The max length of the object is 25,000 characters. If you call &#x60;JSON.stringify&#x60; on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: &#x60;!&#x60;, &#x60;\&quot;&#x60;, &#x60;#&#x60;, &#x60;%&#x60;, &#x60;&amp;&#x60;, &#x60;&#39;&#x60;, &#x60;(&#x60;, &#x60;)&#x60;, &#x60;*&#x60;, &#x60;+&#x60;, &#x60;,&#x60;, &#x60;/&#x60;, &#x60;;&#x60;, &#x60;&lt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;@&#x60;, &#x60;[&#x60;, &#x60;\\&#x60;, &#x60;]&#x60;, &#x60;^&#x60;, &#x60;&#x60; &#x60; &#x60;&#x60;, &#x60;{&#x60;, &#x60;|&#x60;, &#x60;}&#x60;, &#x60;~&#x60;. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string. | [optional] 
**SendDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the &#x60;send_date&#x60; has passed, the mailpiece can be canceled. If a date in the format &#x60;2017-11-01&#x60; is passed, it will evaluate to midnight UTC of that date (&#x60;2017-11-01T00:00:00.000Z&#x60;). If a datetime is passed, that exact time will be used. A &#x60;send_date&#x60; passed with no time zone will default to UTC, while a &#x60;send_date&#x60; passed with a time zone will be converted to UTC. | [optional] 
**OutsideTemplateId** | Pointer to **NullableString** | The unique ID of the HTML template used for the outside of the self mailer. | [optional] 
**InsideTemplateId** | Pointer to **NullableString** | The unique ID of the HTML template used for the inside of the self mailer. | [optional] 
**OutsideTemplateVersionId** | Pointer to **NullableString** | The unique ID of the specific version of the HTML template used for the outside of the self mailer. | [optional] 
**InsideTemplateVersionId** | Pointer to **NullableString** | The unique ID of the specific version of the HTML template used for the inside of the self mailer. | [optional] 
**Object** | Pointer to **string** | Value is resource type. | [optional] [default to "self_mailer"]
**TrackingEvents** | Pointer to [**[]TrackingEventCertified**](TrackingEventCertified.md) | An array of certified tracking events ordered by ascending &#x60;time&#x60;. Not populated in test mode. | [optional] 
**Url** | **string** | A [signed link](#section/Asset-URLs) served over HTTPS. The link returned will expire in 30 days to prevent mis-sharing. Each time a GET request is initiated, a new signed URL will be generated. | 

## Methods

### NewSelfMailer

`func NewSelfMailer(id string, to interface{}, url string, ) *SelfMailer`

NewSelfMailer instantiates a new SelfMailer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfMailerWithDefaults

`func NewSelfMailerWithDefaults() *SelfMailer`

NewSelfMailerWithDefaults instantiates a new SelfMailer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SelfMailer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelfMailer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelfMailer) SetId(v string)`

SetId sets Id field to given value.


### GetTo

`func (o *SelfMailer) GetTo() interface{}`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SelfMailer) GetToOk() (*interface{}, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SelfMailer) SetTo(v interface{})`

SetTo sets To field to given value.


### SetToNil

`func (o *SelfMailer) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *SelfMailer) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetFrom

`func (o *SelfMailer) GetFrom() interface{}`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SelfMailer) GetFromOk() (*interface{}, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SelfMailer) SetFrom(v interface{})`

SetFrom sets From field to given value.

### HasFrom

`func (o *SelfMailer) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *SelfMailer) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *SelfMailer) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetSize

`func (o *SelfMailer) GetSize() SelfMailerSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SelfMailer) GetSizeOk() (*SelfMailerSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SelfMailer) SetSize(v SelfMailerSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *SelfMailer) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDescription

`func (o *SelfMailer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SelfMailer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SelfMailer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SelfMailer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SelfMailer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SelfMailer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *SelfMailer) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SelfMailer) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SelfMailer) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SelfMailer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMailType

`func (o *SelfMailer) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *SelfMailer) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *SelfMailer) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *SelfMailer) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMergeVariables

`func (o *SelfMailer) GetMergeVariables() map[string]interface{}`

GetMergeVariables returns the MergeVariables field if non-nil, zero value otherwise.

### GetMergeVariablesOk

`func (o *SelfMailer) GetMergeVariablesOk() (*map[string]interface{}, bool)`

GetMergeVariablesOk returns a tuple with the MergeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariables

`func (o *SelfMailer) SetMergeVariables(v map[string]interface{})`

SetMergeVariables sets MergeVariables field to given value.

### HasMergeVariables

`func (o *SelfMailer) HasMergeVariables() bool`

HasMergeVariables returns a boolean if a field has been set.

### SetMergeVariablesNil

`func (o *SelfMailer) SetMergeVariablesNil(b bool)`

 SetMergeVariablesNil sets the value for MergeVariables to be an explicit nil

### UnsetMergeVariables
`func (o *SelfMailer) UnsetMergeVariables()`

UnsetMergeVariables ensures that no value is present for MergeVariables, not even an explicit nil
### GetSendDate

`func (o *SelfMailer) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *SelfMailer) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *SelfMailer) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *SelfMailer) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### GetOutsideTemplateId

`func (o *SelfMailer) GetOutsideTemplateId() string`

GetOutsideTemplateId returns the OutsideTemplateId field if non-nil, zero value otherwise.

### GetOutsideTemplateIdOk

`func (o *SelfMailer) GetOutsideTemplateIdOk() (*string, bool)`

GetOutsideTemplateIdOk returns a tuple with the OutsideTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideTemplateId

`func (o *SelfMailer) SetOutsideTemplateId(v string)`

SetOutsideTemplateId sets OutsideTemplateId field to given value.

### HasOutsideTemplateId

`func (o *SelfMailer) HasOutsideTemplateId() bool`

HasOutsideTemplateId returns a boolean if a field has been set.

### SetOutsideTemplateIdNil

`func (o *SelfMailer) SetOutsideTemplateIdNil(b bool)`

 SetOutsideTemplateIdNil sets the value for OutsideTemplateId to be an explicit nil

### UnsetOutsideTemplateId
`func (o *SelfMailer) UnsetOutsideTemplateId()`

UnsetOutsideTemplateId ensures that no value is present for OutsideTemplateId, not even an explicit nil
### GetInsideTemplateId

`func (o *SelfMailer) GetInsideTemplateId() string`

GetInsideTemplateId returns the InsideTemplateId field if non-nil, zero value otherwise.

### GetInsideTemplateIdOk

`func (o *SelfMailer) GetInsideTemplateIdOk() (*string, bool)`

GetInsideTemplateIdOk returns a tuple with the InsideTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsideTemplateId

`func (o *SelfMailer) SetInsideTemplateId(v string)`

SetInsideTemplateId sets InsideTemplateId field to given value.

### HasInsideTemplateId

`func (o *SelfMailer) HasInsideTemplateId() bool`

HasInsideTemplateId returns a boolean if a field has been set.

### SetInsideTemplateIdNil

`func (o *SelfMailer) SetInsideTemplateIdNil(b bool)`

 SetInsideTemplateIdNil sets the value for InsideTemplateId to be an explicit nil

### UnsetInsideTemplateId
`func (o *SelfMailer) UnsetInsideTemplateId()`

UnsetInsideTemplateId ensures that no value is present for InsideTemplateId, not even an explicit nil
### GetOutsideTemplateVersionId

`func (o *SelfMailer) GetOutsideTemplateVersionId() string`

GetOutsideTemplateVersionId returns the OutsideTemplateVersionId field if non-nil, zero value otherwise.

### GetOutsideTemplateVersionIdOk

`func (o *SelfMailer) GetOutsideTemplateVersionIdOk() (*string, bool)`

GetOutsideTemplateVersionIdOk returns a tuple with the OutsideTemplateVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideTemplateVersionId

`func (o *SelfMailer) SetOutsideTemplateVersionId(v string)`

SetOutsideTemplateVersionId sets OutsideTemplateVersionId field to given value.

### HasOutsideTemplateVersionId

`func (o *SelfMailer) HasOutsideTemplateVersionId() bool`

HasOutsideTemplateVersionId returns a boolean if a field has been set.

### SetOutsideTemplateVersionIdNil

`func (o *SelfMailer) SetOutsideTemplateVersionIdNil(b bool)`

 SetOutsideTemplateVersionIdNil sets the value for OutsideTemplateVersionId to be an explicit nil

### UnsetOutsideTemplateVersionId
`func (o *SelfMailer) UnsetOutsideTemplateVersionId()`

UnsetOutsideTemplateVersionId ensures that no value is present for OutsideTemplateVersionId, not even an explicit nil
### GetInsideTemplateVersionId

`func (o *SelfMailer) GetInsideTemplateVersionId() string`

GetInsideTemplateVersionId returns the InsideTemplateVersionId field if non-nil, zero value otherwise.

### GetInsideTemplateVersionIdOk

`func (o *SelfMailer) GetInsideTemplateVersionIdOk() (*string, bool)`

GetInsideTemplateVersionIdOk returns a tuple with the InsideTemplateVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsideTemplateVersionId

`func (o *SelfMailer) SetInsideTemplateVersionId(v string)`

SetInsideTemplateVersionId sets InsideTemplateVersionId field to given value.

### HasInsideTemplateVersionId

`func (o *SelfMailer) HasInsideTemplateVersionId() bool`

HasInsideTemplateVersionId returns a boolean if a field has been set.

### SetInsideTemplateVersionIdNil

`func (o *SelfMailer) SetInsideTemplateVersionIdNil(b bool)`

 SetInsideTemplateVersionIdNil sets the value for InsideTemplateVersionId to be an explicit nil

### UnsetInsideTemplateVersionId
`func (o *SelfMailer) UnsetInsideTemplateVersionId()`

UnsetInsideTemplateVersionId ensures that no value is present for InsideTemplateVersionId, not even an explicit nil
### GetObject

`func (o *SelfMailer) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SelfMailer) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SelfMailer) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SelfMailer) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetTrackingEvents

`func (o *SelfMailer) GetTrackingEvents() []TrackingEventCertified`

GetTrackingEvents returns the TrackingEvents field if non-nil, zero value otherwise.

### GetTrackingEventsOk

`func (o *SelfMailer) GetTrackingEventsOk() (*[]TrackingEventCertified, bool)`

GetTrackingEventsOk returns a tuple with the TrackingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingEvents

`func (o *SelfMailer) SetTrackingEvents(v []TrackingEventCertified)`

SetTrackingEvents sets TrackingEvents field to given value.

### HasTrackingEvents

`func (o *SelfMailer) HasTrackingEvents() bool`

HasTrackingEvents returns a boolean if a field has been set.

### GetUrl

`func (o *SelfMailer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SelfMailer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SelfMailer) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


