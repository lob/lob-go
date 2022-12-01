# Letter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**Address**](Address.md) |  | 
**From** | [**Address**](Address.md) |  | 
**Carrier** | Pointer to **string** |  | [optional] [default to "USPS"]
**Thumbnails** | Pointer to [**[]Thumbnail**](Thumbnail.md) |  | [optional] 
**ExpectedDeliveryDate** | Pointer to **string** | A date in YYYY-MM-DD format of the mailpiece&#39;s expected delivery date based on its &#x60;send_date&#x60;. | [optional] 
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Id** | **string** | Unique identifier prefixed with &#x60;ltr_&#x60;. | 
**TemplateId** | Pointer to **string** | Unique identifier prefixed with &#x60;tmpl_&#x60;. ID of a saved [HTML template](#section/HTML-Templates). | [optional] 
**TemplateVersionId** | Pointer to **string** | Unique identifier prefixed with &#x60;vrsn_&#x60;. | [optional] 
**Url** | Pointer to **string** | A [signed link](#section/Asset-URLs) served over HTTPS. The link returned will expire in 30 days to prevent mis-sharing. Each time a GET request is initiated, a new signed URL will be generated. | [optional] 
**Object** | **string** |  | [default to "letter"]
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**MergeVariables** | Pointer to **map[string]interface{}** | You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: &#x60;{{variable_name}}&#x60;, pass in &#x60;{\&quot;variable_name\&quot;: \&quot;Harry\&quot;}&#x60; to render &#x60;Harry&#x60;. &#x60;merge_variables&#x60; must be an object. Any type of value is accepted as long as the object is valid JSON; you can use &#x60;strings&#x60;, &#x60;numbers&#x60;, &#x60;booleans&#x60;, &#x60;arrays&#x60;, &#x60;objects&#x60;, or &#x60;null&#x60;. The max length of the object is 25,000 characters. If you call &#x60;JSON.stringify&#x60; on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: &#x60;!&#x60;, &#x60;\&quot;&#x60;, &#x60;#&#x60;, &#x60;%&#x60;, &#x60;&amp;&#x60;, &#x60;&#39;&#x60;, &#x60;(&#x60;, &#x60;)&#x60;, &#x60;*&#x60;, &#x60;+&#x60;, &#x60;,&#x60;, &#x60;/&#x60;, &#x60;;&#x60;, &#x60;&lt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;@&#x60;, &#x60;[&#x60;, &#x60;\\&#x60;, &#x60;]&#x60;, &#x60;^&#x60;, &#x60;&#x60; &#x60; &#x60;&#x60;, &#x60;{&#x60;, &#x60;|&#x60;, &#x60;}&#x60;, &#x60;~&#x60;. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string. | [optional] 
**SendDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the &#x60;send_date&#x60; has passed, the mailpiece can be canceled. If a date in the format &#x60;2017-11-01&#x60; is passed, it will evaluate to midnight UTC of that date (&#x60;2017-11-01T00:00:00.000Z&#x60;). If a datetime is passed, that exact time will be used. A &#x60;send_date&#x60; passed with no time zone will default to UTC, while a &#x60;send_date&#x60; passed with a time zone will be converted to UTC. | [optional] 
**ExtraService** | Pointer to **string** | Add an extra service to your letter. See [pricing](https://www.lob.com/pricing/print-mail#compare) for extra costs incurred. | [optional] 
**TrackingNumber** | Pointer to **NullableString** | The tracking number, if applicable, will appear here when it becomes available. Dummy tracking numbers are not created in test mode. | [optional] 
**TrackingEvents** | Pointer to [**[]TrackingEventNormal**](TrackingEventNormal.md) | Tracking events are not populated for registered or regular (no extra service) letters. | [optional] 
**ReturnAddress** | Pointer to **interface{}** | Specifies the address the return envelope will be sent back to. This is an optional argument that is available if an account is signed up for the return envelope tracking beta, and has &#x60;return_envelope&#x60;, and &#x60;perforated_page&#x60; fields populated in the API request. | [optional] 
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to MAILTYPE_FIRST_CLASS]
**Color** | Pointer to **bool** | Set this key to &#x60;true&#x60; if you would like to print in color. Set to &#x60;false&#x60; if you would like to print in black and white. | [optional] 
**DoubleSided** | Pointer to **bool** | Set this attribute to &#x60;true&#x60; for double sided printing, or &#x60;false&#x60; for for single sided printing. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**AddressPlacement** | Pointer to **string** | Specifies the location of the address information that will show through the double-window envelope.  | [optional] [default to "top_first_page"]
**ReturnEnvelope** | **interface{}** |  | 
**PerforatedPage** | Pointer to **NullableInt32** | Required if &#x60;return_envelope&#x60; is &#x60;true&#x60;. The number of the page that should be perforated for use with the return envelope. Must be greater than or equal to &#x60;1&#x60;. The blank page added by &#x60;address_placement&#x3D;insert_blank_page&#x60; will be ignored when considering the perforated page number. To see how perforation will impact your letter design, view our [perforation guide](https://s3-us-west-2.amazonaws.com/public.lob.com/assets/templates/letter_perf_template.pdf). | [optional] 
**CustomEnvelope** | Pointer to [**NullableLetterCustomEnvelope**](LetterCustomEnvelope.md) |  | [optional] 

## Methods

### NewLetter

`func NewLetter(to Address, from Address, dateCreated time.Time, dateModified time.Time, id string, object string, returnEnvelope interface{}, ) *Letter`

NewLetter instantiates a new Letter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLetterWithDefaults

`func NewLetterWithDefaults() *Letter`

NewLetterWithDefaults instantiates a new Letter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *Letter) GetTo() Address`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Letter) GetToOk() (*Address, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Letter) SetTo(v Address)`

SetTo sets To field to given value.


### GetFrom

`func (o *Letter) GetFrom() Address`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Letter) GetFromOk() (*Address, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Letter) SetFrom(v Address)`

SetFrom sets From field to given value.


### GetCarrier

`func (o *Letter) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *Letter) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *Letter) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *Letter) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetThumbnails

`func (o *Letter) GetThumbnails() []Thumbnail`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *Letter) GetThumbnailsOk() (*[]Thumbnail, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *Letter) SetThumbnails(v []Thumbnail)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *Letter) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### GetExpectedDeliveryDate

`func (o *Letter) GetExpectedDeliveryDate() string`

GetExpectedDeliveryDate returns the ExpectedDeliveryDate field if non-nil, zero value otherwise.

### GetExpectedDeliveryDateOk

`func (o *Letter) GetExpectedDeliveryDateOk() (*string, bool)`

GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDeliveryDate

`func (o *Letter) SetExpectedDeliveryDate(v string)`

SetExpectedDeliveryDate sets ExpectedDeliveryDate field to given value.

### HasExpectedDeliveryDate

`func (o *Letter) HasExpectedDeliveryDate() bool`

HasExpectedDeliveryDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *Letter) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Letter) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Letter) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *Letter) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *Letter) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *Letter) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetDeleted

`func (o *Letter) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Letter) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Letter) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Letter) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *Letter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Letter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Letter) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *Letter) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Letter) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Letter) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Letter) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVersionId

`func (o *Letter) GetTemplateVersionId() string`

GetTemplateVersionId returns the TemplateVersionId field if non-nil, zero value otherwise.

### GetTemplateVersionIdOk

`func (o *Letter) GetTemplateVersionIdOk() (*string, bool)`

GetTemplateVersionIdOk returns a tuple with the TemplateVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersionId

`func (o *Letter) SetTemplateVersionId(v string)`

SetTemplateVersionId sets TemplateVersionId field to given value.

### HasTemplateVersionId

`func (o *Letter) HasTemplateVersionId() bool`

HasTemplateVersionId returns a boolean if a field has been set.

### GetUrl

`func (o *Letter) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Letter) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Letter) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Letter) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetObject

`func (o *Letter) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Letter) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Letter) SetObject(v string)`

SetObject sets Object field to given value.


### GetDescription

`func (o *Letter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Letter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Letter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Letter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Letter) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Letter) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *Letter) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Letter) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Letter) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Letter) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMergeVariables

`func (o *Letter) GetMergeVariables() map[string]interface{}`

GetMergeVariables returns the MergeVariables field if non-nil, zero value otherwise.

### GetMergeVariablesOk

`func (o *Letter) GetMergeVariablesOk() (*map[string]interface{}, bool)`

GetMergeVariablesOk returns a tuple with the MergeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariables

`func (o *Letter) SetMergeVariables(v map[string]interface{})`

SetMergeVariables sets MergeVariables field to given value.

### HasMergeVariables

`func (o *Letter) HasMergeVariables() bool`

HasMergeVariables returns a boolean if a field has been set.

### SetMergeVariablesNil

`func (o *Letter) SetMergeVariablesNil(b bool)`

 SetMergeVariablesNil sets the value for MergeVariables to be an explicit nil

### UnsetMergeVariables
`func (o *Letter) UnsetMergeVariables()`

UnsetMergeVariables ensures that no value is present for MergeVariables, not even an explicit nil
### GetSendDate

`func (o *Letter) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *Letter) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *Letter) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *Letter) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### GetExtraService

`func (o *Letter) GetExtraService() string`

GetExtraService returns the ExtraService field if non-nil, zero value otherwise.

### GetExtraServiceOk

`func (o *Letter) GetExtraServiceOk() (*string, bool)`

GetExtraServiceOk returns a tuple with the ExtraService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraService

`func (o *Letter) SetExtraService(v string)`

SetExtraService sets ExtraService field to given value.

### HasExtraService

`func (o *Letter) HasExtraService() bool`

HasExtraService returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *Letter) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *Letter) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *Letter) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *Letter) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *Letter) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *Letter) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetTrackingEvents

`func (o *Letter) GetTrackingEvents() []TrackingEventNormal`

GetTrackingEvents returns the TrackingEvents field if non-nil, zero value otherwise.

### GetTrackingEventsOk

`func (o *Letter) GetTrackingEventsOk() (*[]TrackingEventNormal, bool)`

GetTrackingEventsOk returns a tuple with the TrackingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingEvents

`func (o *Letter) SetTrackingEvents(v []TrackingEventNormal)`

SetTrackingEvents sets TrackingEvents field to given value.

### HasTrackingEvents

`func (o *Letter) HasTrackingEvents() bool`

HasTrackingEvents returns a boolean if a field has been set.

### GetReturnAddress

`func (o *Letter) GetReturnAddress() interface{}`

GetReturnAddress returns the ReturnAddress field if non-nil, zero value otherwise.

### GetReturnAddressOk

`func (o *Letter) GetReturnAddressOk() (*interface{}, bool)`

GetReturnAddressOk returns a tuple with the ReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAddress

`func (o *Letter) SetReturnAddress(v interface{})`

SetReturnAddress sets ReturnAddress field to given value.

### HasReturnAddress

`func (o *Letter) HasReturnAddress() bool`

HasReturnAddress returns a boolean if a field has been set.

### SetReturnAddressNil

`func (o *Letter) SetReturnAddressNil(b bool)`

 SetReturnAddressNil sets the value for ReturnAddress to be an explicit nil

### UnsetReturnAddress
`func (o *Letter) UnsetReturnAddress()`

UnsetReturnAddress ensures that no value is present for ReturnAddress, not even an explicit nil
### GetMailType

`func (o *Letter) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *Letter) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *Letter) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *Letter) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetColor

`func (o *Letter) GetColor() bool`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Letter) GetColorOk() (*bool, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Letter) SetColor(v bool)`

SetColor sets Color field to given value.

### HasColor

`func (o *Letter) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDoubleSided

`func (o *Letter) GetDoubleSided() bool`

GetDoubleSided returns the DoubleSided field if non-nil, zero value otherwise.

### GetDoubleSidedOk

`func (o *Letter) GetDoubleSidedOk() (*bool, bool)`

GetDoubleSidedOk returns a tuple with the DoubleSided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleSided

`func (o *Letter) SetDoubleSided(v bool)`

SetDoubleSided sets DoubleSided field to given value.

### HasDoubleSided

`func (o *Letter) HasDoubleSided() bool`

HasDoubleSided returns a boolean if a field has been set.

### GetAddressPlacement

`func (o *Letter) GetAddressPlacement() string`

GetAddressPlacement returns the AddressPlacement field if non-nil, zero value otherwise.

### GetAddressPlacementOk

`func (o *Letter) GetAddressPlacementOk() (*string, bool)`

GetAddressPlacementOk returns a tuple with the AddressPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPlacement

`func (o *Letter) SetAddressPlacement(v string)`

SetAddressPlacement sets AddressPlacement field to given value.

### HasAddressPlacement

`func (o *Letter) HasAddressPlacement() bool`

HasAddressPlacement returns a boolean if a field has been set.

### GetReturnEnvelope

`func (o *Letter) GetReturnEnvelope() interface{}`

GetReturnEnvelope returns the ReturnEnvelope field if non-nil, zero value otherwise.

### GetReturnEnvelopeOk

`func (o *Letter) GetReturnEnvelopeOk() (*interface{}, bool)`

GetReturnEnvelopeOk returns a tuple with the ReturnEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnEnvelope

`func (o *Letter) SetReturnEnvelope(v interface{})`

SetReturnEnvelope sets ReturnEnvelope field to given value.


### SetReturnEnvelopeNil

`func (o *Letter) SetReturnEnvelopeNil(b bool)`

 SetReturnEnvelopeNil sets the value for ReturnEnvelope to be an explicit nil

### UnsetReturnEnvelope
`func (o *Letter) UnsetReturnEnvelope()`

UnsetReturnEnvelope ensures that no value is present for ReturnEnvelope, not even an explicit nil
### GetPerforatedPage

`func (o *Letter) GetPerforatedPage() int32`

GetPerforatedPage returns the PerforatedPage field if non-nil, zero value otherwise.

### GetPerforatedPageOk

`func (o *Letter) GetPerforatedPageOk() (*int32, bool)`

GetPerforatedPageOk returns a tuple with the PerforatedPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerforatedPage

`func (o *Letter) SetPerforatedPage(v int32)`

SetPerforatedPage sets PerforatedPage field to given value.

### HasPerforatedPage

`func (o *Letter) HasPerforatedPage() bool`

HasPerforatedPage returns a boolean if a field has been set.

### SetPerforatedPageNil

`func (o *Letter) SetPerforatedPageNil(b bool)`

 SetPerforatedPageNil sets the value for PerforatedPage to be an explicit nil

### UnsetPerforatedPage
`func (o *Letter) UnsetPerforatedPage()`

UnsetPerforatedPage ensures that no value is present for PerforatedPage, not even an explicit nil
### GetCustomEnvelope

`func (o *Letter) GetCustomEnvelope() LetterCustomEnvelope`

GetCustomEnvelope returns the CustomEnvelope field if non-nil, zero value otherwise.

### GetCustomEnvelopeOk

`func (o *Letter) GetCustomEnvelopeOk() (*LetterCustomEnvelope, bool)`

GetCustomEnvelopeOk returns a tuple with the CustomEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvelope

`func (o *Letter) SetCustomEnvelope(v LetterCustomEnvelope)`

SetCustomEnvelope sets CustomEnvelope field to given value.

### HasCustomEnvelope

`func (o *Letter) HasCustomEnvelope() bool`

HasCustomEnvelope returns a boolean if a field has been set.

### SetCustomEnvelopeNil

`func (o *Letter) SetCustomEnvelopeNil(b bool)`

 SetCustomEnvelopeNil sets the value for CustomEnvelope to be an explicit nil

### UnsetCustomEnvelope
`func (o *Letter) UnsetCustomEnvelope()`

UnsetCustomEnvelope ensures that no value is present for CustomEnvelope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


