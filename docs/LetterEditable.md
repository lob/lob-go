# LetterEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to FIRST_CLASS]
**MergeVariables** | Pointer to **map[string]interface{}** | You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: &#x60;{{variable_name}}&#x60;, pass in &#x60;{\&quot;variable_name\&quot;: \&quot;Harry\&quot;}&#x60; to render &#x60;Harry&#x60;. &#x60;merge_variables&#x60; must be an object. Any type of value is accepted as long as the object is valid JSON; you can use &#x60;strings&#x60;, &#x60;numbers&#x60;, &#x60;booleans&#x60;, &#x60;arrays&#x60;, &#x60;objects&#x60;, or &#x60;null&#x60;. The max length of the object is 25,000 characters. If you call &#x60;JSON.stringify&#x60; on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: &#x60;!&#x60;, &#x60;\&quot;&#x60;, &#x60;#&#x60;, &#x60;%&#x60;, &#x60;&amp;&#x60;, &#x60;&#39;&#x60;, &#x60;(&#x60;, &#x60;)&#x60;, &#x60;*&#x60;, &#x60;+&#x60;, &#x60;,&#x60;, &#x60;/&#x60;, &#x60;;&#x60;, &#x60;&lt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;@&#x60;, &#x60;[&#x60;, &#x60;\\&#x60;, &#x60;]&#x60;, &#x60;^&#x60;, &#x60;&#x60; &#x60; &#x60;&#x60;, &#x60;{&#x60;, &#x60;|&#x60;, &#x60;}&#x60;, &#x60;~&#x60;. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string. | [optional] 
**SendDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the &#x60;send_date&#x60; has passed, the mailpiece can be canceled. If a date in the format &#x60;2017-11-01&#x60; is passed, it will evaluate to midnight UTC of that date (&#x60;2017-11-01T00:00:00.000Z&#x60;). If a datetime is passed, that exact time will be used. A &#x60;send_date&#x60; passed with no time zone will default to UTC, while a &#x60;send_date&#x60; passed with a time zone will be converted to UTC. | [optional] 
**Color** | **bool** | Set this key to &#x60;true&#x60; if you would like to print in color. Set to &#x60;false&#x60; if you would like to print in black and white. | 
**DoubleSided** | Pointer to **bool** | Set this attribute to &#x60;true&#x60; for double sided printing, or &#x60;false&#x60; for for single sided printing. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**AddressPlacement** | Pointer to **string** | Specifies the location of the address information that will show through the double-window envelope. To see how this will impact your letter design, view our letter template.   * &#x60;top_first_page&#x60; - (default) print address information at the top of your provided first page   * &#x60;insert_blank_page&#x60; - insert a blank address page at the beginning of your file (you will be charged for the extra page)   * &#x60;bottom_first_page_center&#x60; - **(deprecation planned within a few months)** print address information at the bottom center of your provided first page   * &#x60;bottom_first_page&#x60; - print address information at the bottom of your provided first page  | [optional] [default to "top_first_page"]
**ReturnEnvelope** | Pointer to **interface{}** | indicates if a return envelope is requested for the letter. The value corresponding to this field is by default a boolean. But if the account is signed up for custom return envelopes, the value is of type string and is &#x60;no_9_single_window&#x60; for a standard return envelope and a custom &#x60;return_envelope_id&#x60; for non-standard return envelopes.  To include a return envelope with your letter, set to &#x60;true&#x60; and specify the &#x60;perforated_page&#x60;. See [pricing](https://www.lob.com/pricing/print-mail#compare) for extra costs incurred. | [optional] 
**PerforatedPage** | Pointer to **NullableInt32** | Required if &#x60;return_envelope&#x60; is &#x60;true&#x60;. The number of the page that should be perforated for use with the return envelope. Must be greater than or equal to &#x60;1&#x60;. The blank page added by &#x60;address_placement&#x3D;insert_blank_page&#x60; will be ignored when considering the perforated page number. To see how perforation will impact your letter design, view our [perforation guide](https://s3-us-west-2.amazonaws.com/public.lob.com/assets/templates/letter_perf_template.pdf). | [optional] 
**CustomEnvelope** | Pointer to **NullableString** |  | [optional] 
**To** | **string** | Must either be an address ID or an inline object with correct address parameters. | 
**From** | **string** | Must either be an address ID or an inline object with correct address parameters. | 
**File** | **string** | PDF file containing the letter&#39;s formatting. | 
**ExtraService** | Pointer to **NullableString** | Add an extra service to your letter:   * &#x60;certified&#x60; - track and confirm delivery for domestic destinations. An extra sheet (1 PDF page single-sided or 2 PDF pages double-sided) is added to the beginning of your letter for address and barcode information. See here for templates: [#10 envelope](https://s3-us-west-2.amazonaws.com/public.lob.com/assets/templates/letter_certified_template.pdf) and [flat envelope](https://s3-us-west-2.amazonaws.com/public.lob.com/assets/templates/letter_certified_flat_template.pdf) (used for letters over 6 pages single-sided or 12 pages double-sided). You will not be charged for this extra sheet.   * &#x60;certified_return_receipt&#x60; - request an electronic copy of the recipient&#39;s signature to prove delivery of your certified letter   * &#x60;registered&#x60; - provides tracking and confirmation for international addresses  | [optional] 
**Cards** | Pointer to **[]string** | A single-element array containing an existing card id in a string format. See [cards](#tag/Cards) for more information. | [optional] 
**BillingGroupId** | Pointer to **string** | An optional string with the billing group ID to tag your usage with. Is used for billing purposes. Requires special activation to use. See [Billing Group API](https://lob.github.io/lob-openapi/#tag/Billing-Groups) for more information. | [optional] 

## Methods

### NewLetterEditable

`func NewLetterEditable(color bool, to string, from string, file string, ) *LetterEditable`

NewLetterEditable instantiates a new LetterEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLetterEditableWithDefaults

`func NewLetterEditableWithDefaults() *LetterEditable`

NewLetterEditableWithDefaults instantiates a new LetterEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LetterEditable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LetterEditable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LetterEditable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LetterEditable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LetterEditable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LetterEditable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *LetterEditable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LetterEditable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LetterEditable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LetterEditable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMailType

`func (o *LetterEditable) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *LetterEditable) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *LetterEditable) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *LetterEditable) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMergeVariables

`func (o *LetterEditable) GetMergeVariables() map[string]interface{}`

GetMergeVariables returns the MergeVariables field if non-nil, zero value otherwise.

### GetMergeVariablesOk

`func (o *LetterEditable) GetMergeVariablesOk() (*map[string]interface{}, bool)`

GetMergeVariablesOk returns a tuple with the MergeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariables

`func (o *LetterEditable) SetMergeVariables(v map[string]interface{})`

SetMergeVariables sets MergeVariables field to given value.

### HasMergeVariables

`func (o *LetterEditable) HasMergeVariables() bool`

HasMergeVariables returns a boolean if a field has been set.

### SetMergeVariablesNil

`func (o *LetterEditable) SetMergeVariablesNil(b bool)`

 SetMergeVariablesNil sets the value for MergeVariables to be an explicit nil

### UnsetMergeVariables
`func (o *LetterEditable) UnsetMergeVariables()`

UnsetMergeVariables ensures that no value is present for MergeVariables, not even an explicit nil
### GetSendDate

`func (o *LetterEditable) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *LetterEditable) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *LetterEditable) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *LetterEditable) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### GetColor

`func (o *LetterEditable) GetColor() bool`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LetterEditable) GetColorOk() (*bool, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LetterEditable) SetColor(v bool)`

SetColor sets Color field to given value.


### GetDoubleSided

`func (o *LetterEditable) GetDoubleSided() bool`

GetDoubleSided returns the DoubleSided field if non-nil, zero value otherwise.

### GetDoubleSidedOk

`func (o *LetterEditable) GetDoubleSidedOk() (*bool, bool)`

GetDoubleSidedOk returns a tuple with the DoubleSided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleSided

`func (o *LetterEditable) SetDoubleSided(v bool)`

SetDoubleSided sets DoubleSided field to given value.

### HasDoubleSided

`func (o *LetterEditable) HasDoubleSided() bool`

HasDoubleSided returns a boolean if a field has been set.

### GetAddressPlacement

`func (o *LetterEditable) GetAddressPlacement() string`

GetAddressPlacement returns the AddressPlacement field if non-nil, zero value otherwise.

### GetAddressPlacementOk

`func (o *LetterEditable) GetAddressPlacementOk() (*string, bool)`

GetAddressPlacementOk returns a tuple with the AddressPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPlacement

`func (o *LetterEditable) SetAddressPlacement(v string)`

SetAddressPlacement sets AddressPlacement field to given value.

### HasAddressPlacement

`func (o *LetterEditable) HasAddressPlacement() bool`

HasAddressPlacement returns a boolean if a field has been set.

### GetReturnEnvelope

`func (o *LetterEditable) GetReturnEnvelope() interface{}`

GetReturnEnvelope returns the ReturnEnvelope field if non-nil, zero value otherwise.

### GetReturnEnvelopeOk

`func (o *LetterEditable) GetReturnEnvelopeOk() (*interface{}, bool)`

GetReturnEnvelopeOk returns a tuple with the ReturnEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnEnvelope

`func (o *LetterEditable) SetReturnEnvelope(v interface{})`

SetReturnEnvelope sets ReturnEnvelope field to given value.

### HasReturnEnvelope

`func (o *LetterEditable) HasReturnEnvelope() bool`

HasReturnEnvelope returns a boolean if a field has been set.

### SetReturnEnvelopeNil

`func (o *LetterEditable) SetReturnEnvelopeNil(b bool)`

 SetReturnEnvelopeNil sets the value for ReturnEnvelope to be an explicit nil

### UnsetReturnEnvelope
`func (o *LetterEditable) UnsetReturnEnvelope()`

UnsetReturnEnvelope ensures that no value is present for ReturnEnvelope, not even an explicit nil
### GetPerforatedPage

`func (o *LetterEditable) GetPerforatedPage() int32`

GetPerforatedPage returns the PerforatedPage field if non-nil, zero value otherwise.

### GetPerforatedPageOk

`func (o *LetterEditable) GetPerforatedPageOk() (*int32, bool)`

GetPerforatedPageOk returns a tuple with the PerforatedPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerforatedPage

`func (o *LetterEditable) SetPerforatedPage(v int32)`

SetPerforatedPage sets PerforatedPage field to given value.

### HasPerforatedPage

`func (o *LetterEditable) HasPerforatedPage() bool`

HasPerforatedPage returns a boolean if a field has been set.

### SetPerforatedPageNil

`func (o *LetterEditable) SetPerforatedPageNil(b bool)`

 SetPerforatedPageNil sets the value for PerforatedPage to be an explicit nil

### UnsetPerforatedPage
`func (o *LetterEditable) UnsetPerforatedPage()`

UnsetPerforatedPage ensures that no value is present for PerforatedPage, not even an explicit nil
### GetCustomEnvelope

`func (o *LetterEditable) GetCustomEnvelope() string`

GetCustomEnvelope returns the CustomEnvelope field if non-nil, zero value otherwise.

### GetCustomEnvelopeOk

`func (o *LetterEditable) GetCustomEnvelopeOk() (*string, bool)`

GetCustomEnvelopeOk returns a tuple with the CustomEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvelope

`func (o *LetterEditable) SetCustomEnvelope(v string)`

SetCustomEnvelope sets CustomEnvelope field to given value.

### HasCustomEnvelope

`func (o *LetterEditable) HasCustomEnvelope() bool`

HasCustomEnvelope returns a boolean if a field has been set.

### SetCustomEnvelopeNil

`func (o *LetterEditable) SetCustomEnvelopeNil(b bool)`

 SetCustomEnvelopeNil sets the value for CustomEnvelope to be an explicit nil

### UnsetCustomEnvelope
`func (o *LetterEditable) UnsetCustomEnvelope()`

UnsetCustomEnvelope ensures that no value is present for CustomEnvelope, not even an explicit nil
### GetTo

`func (o *LetterEditable) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LetterEditable) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LetterEditable) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *LetterEditable) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LetterEditable) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LetterEditable) SetFrom(v string)`

SetFrom sets From field to given value.


### GetFile

`func (o *LetterEditable) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *LetterEditable) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *LetterEditable) SetFile(v string)`

SetFile sets File field to given value.


### GetExtraService

`func (o *LetterEditable) GetExtraService() string`

GetExtraService returns the ExtraService field if non-nil, zero value otherwise.

### GetExtraServiceOk

`func (o *LetterEditable) GetExtraServiceOk() (*string, bool)`

GetExtraServiceOk returns a tuple with the ExtraService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraService

`func (o *LetterEditable) SetExtraService(v string)`

SetExtraService sets ExtraService field to given value.

### HasExtraService

`func (o *LetterEditable) HasExtraService() bool`

HasExtraService returns a boolean if a field has been set.

### SetExtraServiceNil

`func (o *LetterEditable) SetExtraServiceNil(b bool)`

 SetExtraServiceNil sets the value for ExtraService to be an explicit nil

### UnsetExtraService
`func (o *LetterEditable) UnsetExtraService()`

UnsetExtraService ensures that no value is present for ExtraService, not even an explicit nil
### GetCards

`func (o *LetterEditable) GetCards() []string`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *LetterEditable) GetCardsOk() (*[]string, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *LetterEditable) SetCards(v []string)`

SetCards sets Cards field to given value.

### HasCards

`func (o *LetterEditable) HasCards() bool`

HasCards returns a boolean if a field has been set.

### SetCardsNil

`func (o *LetterEditable) SetCardsNil(b bool)`

 SetCardsNil sets the value for Cards to be an explicit nil

### UnsetCards
`func (o *LetterEditable) UnsetCards()`

UnsetCards ensures that no value is present for Cards, not even an explicit nil
### GetBillingGroupId

`func (o *LetterEditable) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *LetterEditable) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *LetterEditable) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *LetterEditable) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


