# Check

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;chk_&#x60;. | 
**To** | [**Address**](Address.md) |  | 
**From** | Pointer to [**Address**](Address.md) |  | [optional] 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**MergeVariables** | Pointer to **map[string]interface{}** | You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: &#x60;{{variable_name}}&#x60;, pass in &#x60;{\&quot;variable_name\&quot;: \&quot;Harry\&quot;}&#x60; to render &#x60;Harry&#x60;. &#x60;merge_variables&#x60; must be an object. Any type of value is accepted as long as the object is valid JSON; you can use &#x60;strings&#x60;, &#x60;numbers&#x60;, &#x60;booleans&#x60;, &#x60;arrays&#x60;, &#x60;objects&#x60;, or &#x60;null&#x60;. The max length of the object is 25,000 characters. If you call &#x60;JSON.stringify&#x60; on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: &#x60;!&#x60;, &#x60;\&quot;&#x60;, &#x60;#&#x60;, &#x60;%&#x60;, &#x60;&amp;&#x60;, &#x60;&#39;&#x60;, &#x60;(&#x60;, &#x60;)&#x60;, &#x60;*&#x60;, &#x60;+&#x60;, &#x60;,&#x60;, &#x60;/&#x60;, &#x60;;&#x60;, &#x60;&lt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;@&#x60;, &#x60;[&#x60;, &#x60;\\&#x60;, &#x60;]&#x60;, &#x60;^&#x60;, &#x60;&#x60; &#x60; &#x60;&#x60;, &#x60;{&#x60;, &#x60;|&#x60;, &#x60;}&#x60;, &#x60;~&#x60;. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string. | [optional] 
**SendDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the &#x60;send_date&#x60; has passed, the mailpiece can be canceled. If a date in the format &#x60;2017-11-01&#x60; is passed, it will evaluate to midnight UTC of that date (&#x60;2017-11-01T00:00:00.000Z&#x60;). If a datetime is passed, that exact time will be used. A &#x60;send_date&#x60; passed with no time zone will default to UTC, while a &#x60;send_date&#x60; passed with a time zone will be converted to UTC. | [optional] 
**MailType** | Pointer to **string** | Checks must be sent &#x60;usps_first_class&#x60; | [optional] [default to "usps_first_class"]
**Memo** | Pointer to **NullableString** | Text to include on the memo line of the check. | [optional] 
**CheckNumber** | Pointer to **int32** | An integer that designates the check number. If &#x60;check_number&#x60; is not provided, checks created from a new &#x60;bank_account&#x60; will start at &#x60;10000&#x60; and increment with each check created with the &#x60;bank_account&#x60;. A provided &#x60;check_number&#x60; overrides the defaults. Subsequent checks created with the same &#x60;bank_account&#x60; will increment from the provided check number. | [optional] 
**Message** | Pointer to **string** | Max of 400 characters to be included at the bottom of the check page. | [optional] 
**Amount** | **float32** | The payment amount to be sent in US dollars. | 
**BankAccount** | [**BankAccount**](BankAccount.md) |  | 
**CheckBottomTemplateId** | Pointer to **string** | Unique identifier prefixed with &#x60;tmpl_&#x60;. ID of a saved [HTML template](#section/HTML-Templates). | [optional] 
**AttachmentTemplateId** | Pointer to **string** | Unique identifier prefixed with &#x60;tmpl_&#x60;. ID of a saved [HTML template](#section/HTML-Templates). | [optional] 
**CheckBottomTemplateVersionId** | Pointer to **string** | Unique identifier prefixed with &#x60;vrsn_&#x60;. | [optional] 
**AttachmentTemplateVersionId** | Pointer to **string** | Unique identifier prefixed with &#x60;vrsn_&#x60;. | [optional] 
**Url** | **string** | A [signed link](#section/Asset-URLs) served over HTTPS. The link returned will expire in 30 days to prevent mis-sharing. Each time a GET request is initiated, a new signed URL will be generated. | 
**Carrier** | **string** |  | [default to "USPS"]
**Thumbnails** | Pointer to [**[]Thumbnail**](Thumbnail.md) |  | [optional] 
**ExpectedDeliveryDate** | Pointer to **string** | A date in YYYY-MM-DD format of the mailpiece&#39;s expected delivery date based on its &#x60;send_date&#x60;. | [optional] 
**TrackingEvents** | Pointer to [**[]TrackingEventNormal**](TrackingEventNormal.md) | An array of tracking_event objects ordered by ascending &#x60;time&#x60;. Will not be populated for checks created in test mode. | [optional] 
**Object** | **string** |  | [default to "check"]
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 

## Methods

### NewCheck

`func NewCheck(id string, to Address, amount float32, bankAccount BankAccount, url string, carrier string, object string, dateCreated time.Time, dateModified time.Time, ) *Check`

NewCheck instantiates a new Check object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithDefaults

`func NewCheckWithDefaults() *Check`

NewCheckWithDefaults instantiates a new Check object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Check) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Check) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Check) SetId(v string)`

SetId sets Id field to given value.


### GetTo

`func (o *Check) GetTo() Address`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Check) GetToOk() (*Address, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Check) SetTo(v Address)`

SetTo sets To field to given value.


### GetFrom

`func (o *Check) GetFrom() Address`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Check) GetFromOk() (*Address, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Check) SetFrom(v Address)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Check) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetDescription

`func (o *Check) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Check) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Check) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Check) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Check) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Check) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *Check) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Check) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Check) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Check) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMergeVariables

`func (o *Check) GetMergeVariables() map[string]interface{}`

GetMergeVariables returns the MergeVariables field if non-nil, zero value otherwise.

### GetMergeVariablesOk

`func (o *Check) GetMergeVariablesOk() (*map[string]interface{}, bool)`

GetMergeVariablesOk returns a tuple with the MergeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariables

`func (o *Check) SetMergeVariables(v map[string]interface{})`

SetMergeVariables sets MergeVariables field to given value.

### HasMergeVariables

`func (o *Check) HasMergeVariables() bool`

HasMergeVariables returns a boolean if a field has been set.

### SetMergeVariablesNil

`func (o *Check) SetMergeVariablesNil(b bool)`

 SetMergeVariablesNil sets the value for MergeVariables to be an explicit nil

### UnsetMergeVariables
`func (o *Check) UnsetMergeVariables()`

UnsetMergeVariables ensures that no value is present for MergeVariables, not even an explicit nil
### GetSendDate

`func (o *Check) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *Check) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *Check) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *Check) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### GetMailType

`func (o *Check) GetMailType() string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *Check) GetMailTypeOk() (*string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *Check) SetMailType(v string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *Check) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMemo

`func (o *Check) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *Check) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *Check) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *Check) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *Check) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *Check) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetCheckNumber

`func (o *Check) GetCheckNumber() int32`

GetCheckNumber returns the CheckNumber field if non-nil, zero value otherwise.

### GetCheckNumberOk

`func (o *Check) GetCheckNumberOk() (*int32, bool)`

GetCheckNumberOk returns a tuple with the CheckNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumber

`func (o *Check) SetCheckNumber(v int32)`

SetCheckNumber sets CheckNumber field to given value.

### HasCheckNumber

`func (o *Check) HasCheckNumber() bool`

HasCheckNumber returns a boolean if a field has been set.

### GetMessage

`func (o *Check) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Check) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Check) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Check) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAmount

`func (o *Check) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Check) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Check) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBankAccount

`func (o *Check) GetBankAccount() BankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *Check) GetBankAccountOk() (*BankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *Check) SetBankAccount(v BankAccount)`

SetBankAccount sets BankAccount field to given value.


### GetCheckBottomTemplateId

`func (o *Check) GetCheckBottomTemplateId() string`

GetCheckBottomTemplateId returns the CheckBottomTemplateId field if non-nil, zero value otherwise.

### GetCheckBottomTemplateIdOk

`func (o *Check) GetCheckBottomTemplateIdOk() (*string, bool)`

GetCheckBottomTemplateIdOk returns a tuple with the CheckBottomTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckBottomTemplateId

`func (o *Check) SetCheckBottomTemplateId(v string)`

SetCheckBottomTemplateId sets CheckBottomTemplateId field to given value.

### HasCheckBottomTemplateId

`func (o *Check) HasCheckBottomTemplateId() bool`

HasCheckBottomTemplateId returns a boolean if a field has been set.

### GetAttachmentTemplateId

`func (o *Check) GetAttachmentTemplateId() string`

GetAttachmentTemplateId returns the AttachmentTemplateId field if non-nil, zero value otherwise.

### GetAttachmentTemplateIdOk

`func (o *Check) GetAttachmentTemplateIdOk() (*string, bool)`

GetAttachmentTemplateIdOk returns a tuple with the AttachmentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentTemplateId

`func (o *Check) SetAttachmentTemplateId(v string)`

SetAttachmentTemplateId sets AttachmentTemplateId field to given value.

### HasAttachmentTemplateId

`func (o *Check) HasAttachmentTemplateId() bool`

HasAttachmentTemplateId returns a boolean if a field has been set.

### GetCheckBottomTemplateVersionId

`func (o *Check) GetCheckBottomTemplateVersionId() string`

GetCheckBottomTemplateVersionId returns the CheckBottomTemplateVersionId field if non-nil, zero value otherwise.

### GetCheckBottomTemplateVersionIdOk

`func (o *Check) GetCheckBottomTemplateVersionIdOk() (*string, bool)`

GetCheckBottomTemplateVersionIdOk returns a tuple with the CheckBottomTemplateVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckBottomTemplateVersionId

`func (o *Check) SetCheckBottomTemplateVersionId(v string)`

SetCheckBottomTemplateVersionId sets CheckBottomTemplateVersionId field to given value.

### HasCheckBottomTemplateVersionId

`func (o *Check) HasCheckBottomTemplateVersionId() bool`

HasCheckBottomTemplateVersionId returns a boolean if a field has been set.

### GetAttachmentTemplateVersionId

`func (o *Check) GetAttachmentTemplateVersionId() string`

GetAttachmentTemplateVersionId returns the AttachmentTemplateVersionId field if non-nil, zero value otherwise.

### GetAttachmentTemplateVersionIdOk

`func (o *Check) GetAttachmentTemplateVersionIdOk() (*string, bool)`

GetAttachmentTemplateVersionIdOk returns a tuple with the AttachmentTemplateVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentTemplateVersionId

`func (o *Check) SetAttachmentTemplateVersionId(v string)`

SetAttachmentTemplateVersionId sets AttachmentTemplateVersionId field to given value.

### HasAttachmentTemplateVersionId

`func (o *Check) HasAttachmentTemplateVersionId() bool`

HasAttachmentTemplateVersionId returns a boolean if a field has been set.

### GetUrl

`func (o *Check) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Check) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Check) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCarrier

`func (o *Check) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *Check) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *Check) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.


### GetThumbnails

`func (o *Check) GetThumbnails() []Thumbnail`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *Check) GetThumbnailsOk() (*[]Thumbnail, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *Check) SetThumbnails(v []Thumbnail)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *Check) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### GetExpectedDeliveryDate

`func (o *Check) GetExpectedDeliveryDate() string`

GetExpectedDeliveryDate returns the ExpectedDeliveryDate field if non-nil, zero value otherwise.

### GetExpectedDeliveryDateOk

`func (o *Check) GetExpectedDeliveryDateOk() (*string, bool)`

GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDeliveryDate

`func (o *Check) SetExpectedDeliveryDate(v string)`

SetExpectedDeliveryDate sets ExpectedDeliveryDate field to given value.

### HasExpectedDeliveryDate

`func (o *Check) HasExpectedDeliveryDate() bool`

HasExpectedDeliveryDate returns a boolean if a field has been set.

### GetTrackingEvents

`func (o *Check) GetTrackingEvents() []TrackingEventNormal`

GetTrackingEvents returns the TrackingEvents field if non-nil, zero value otherwise.

### GetTrackingEventsOk

`func (o *Check) GetTrackingEventsOk() (*[]TrackingEventNormal, bool)`

GetTrackingEventsOk returns a tuple with the TrackingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingEvents

`func (o *Check) SetTrackingEvents(v []TrackingEventNormal)`

SetTrackingEvents sets TrackingEvents field to given value.

### HasTrackingEvents

`func (o *Check) HasTrackingEvents() bool`

HasTrackingEvents returns a boolean if a field has been set.

### SetTrackingEventsNil

`func (o *Check) SetTrackingEventsNil(b bool)`

 SetTrackingEventsNil sets the value for TrackingEvents to be an explicit nil

### UnsetTrackingEvents
`func (o *Check) UnsetTrackingEvents()`

UnsetTrackingEvents ensures that no value is present for TrackingEvents, not even an explicit nil
### GetObject

`func (o *Check) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Check) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Check) SetObject(v string)`

SetObject sets Object field to given value.


### GetDateCreated

`func (o *Check) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Check) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Check) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *Check) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *Check) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *Check) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetDeleted

`func (o *Check) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Check) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Check) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Check) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


