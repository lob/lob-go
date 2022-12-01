# CheckEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **interface{}** | Must either be an address ID or an inline object with correct address parameters. | 
**To** | **interface{}** | Must either be an address ID or an inline object with correct address parameters. | 
**BankAccount** | **NullableString** |  | 
**Amount** | **float32** | The payment amount to be sent in US dollars. | 
**Logo** | Pointer to **string** | Accepts a remote URL or local file upload to an image to print (in grayscale) in the upper-left corner of your check. | [optional] 
**CheckBottom** | Pointer to **string** | The artwork to use on the bottom of the check page.  Notes: - HTML merge variables should not include delimiting whitespace. - PDF, PNG, and JPGs must be sized at 8.5\&quot;x11\&quot; at 300 DPI, while supplied HTML will be rendered and trimmed to fit on a 8.5\&quot;x11\&quot; page. - The check bottom will always be printed in black &amp; white. - Must conform to [this template](https://s3-us-west-2.amazonaws.com/public.lob.com/assets/templates/check_bottom_template.pdf).  Need more help? Consult our [HTML examples](#section/HTML-Examples). | [optional] 
**Attachment** | Pointer to **string** | A document to include with the check. | [optional] 
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**MergeVariables** | Pointer to **map[string]interface{}** | You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: &#x60;{{variable_name}}&#x60;, pass in &#x60;{\&quot;variable_name\&quot;: \&quot;Harry\&quot;}&#x60; to render &#x60;Harry&#x60;. &#x60;merge_variables&#x60; must be an object. Any type of value is accepted as long as the object is valid JSON; you can use &#x60;strings&#x60;, &#x60;numbers&#x60;, &#x60;booleans&#x60;, &#x60;arrays&#x60;, &#x60;objects&#x60;, or &#x60;null&#x60;. The max length of the object is 25,000 characters. If you call &#x60;JSON.stringify&#x60; on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: &#x60;!&#x60;, &#x60;\&quot;&#x60;, &#x60;#&#x60;, &#x60;%&#x60;, &#x60;&amp;&#x60;, &#x60;&#39;&#x60;, &#x60;(&#x60;, &#x60;)&#x60;, &#x60;*&#x60;, &#x60;+&#x60;, &#x60;,&#x60;, &#x60;/&#x60;, &#x60;;&#x60;, &#x60;&lt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;@&#x60;, &#x60;[&#x60;, &#x60;\\&#x60;, &#x60;]&#x60;, &#x60;^&#x60;, &#x60;&#x60; &#x60; &#x60;&#x60;, &#x60;{&#x60;, &#x60;|&#x60;, &#x60;}&#x60;, &#x60;~&#x60;. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string. | [optional] 
**SendDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the &#x60;send_date&#x60; has passed, the mailpiece can be canceled. If a date in the format &#x60;2017-11-01&#x60; is passed, it will evaluate to midnight UTC of that date (&#x60;2017-11-01T00:00:00.000Z&#x60;). If a datetime is passed, that exact time will be used. A &#x60;send_date&#x60; passed with no time zone will default to UTC, while a &#x60;send_date&#x60; passed with a time zone will be converted to UTC. | [optional] 
**MailType** | Pointer to **string** | Checks must be sent &#x60;usps_first_class&#x60; | [optional] [default to "usps_first_class"]
**Memo** | Pointer to **NullableString** | Text to include on the memo line of the check. | [optional] 
**CheckNumber** | Pointer to **int32** | An integer that designates the check number. | [optional] 
**Message** | Pointer to **string** | Max of 400 characters to be included at the bottom of the check page. | [optional] 
**BillingGroupId** | Pointer to **string** | An optional string with the billing group ID to tag your usage with. Is used for billing purposes. Requires special activation to use. See [Billing Group API](https://lob.github.io/lob-openapi/#tag/Billing-Groups) for more information. | [optional] 

## Methods

### NewCheckEditable

`func NewCheckEditable(from interface{}, to interface{}, bankAccount NullableString, amount float32, ) *CheckEditable`

NewCheckEditable instantiates a new CheckEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckEditableWithDefaults

`func NewCheckEditableWithDefaults() *CheckEditable`

NewCheckEditableWithDefaults instantiates a new CheckEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CheckEditable) GetFrom() interface{}`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CheckEditable) GetFromOk() (*interface{}, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CheckEditable) SetFrom(v interface{})`

SetFrom sets From field to given value.


### SetFromNil

`func (o *CheckEditable) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *CheckEditable) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetTo

`func (o *CheckEditable) GetTo() interface{}`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CheckEditable) GetToOk() (*interface{}, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CheckEditable) SetTo(v interface{})`

SetTo sets To field to given value.


### SetToNil

`func (o *CheckEditable) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *CheckEditable) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetBankAccount

`func (o *CheckEditable) GetBankAccount() string`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *CheckEditable) GetBankAccountOk() (*string, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *CheckEditable) SetBankAccount(v string)`

SetBankAccount sets BankAccount field to given value.


### SetBankAccountNil

`func (o *CheckEditable) SetBankAccountNil(b bool)`

 SetBankAccountNil sets the value for BankAccount to be an explicit nil

### UnsetBankAccount
`func (o *CheckEditable) UnsetBankAccount()`

UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
### GetAmount

`func (o *CheckEditable) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckEditable) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckEditable) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetLogo

`func (o *CheckEditable) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CheckEditable) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CheckEditable) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CheckEditable) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetCheckBottom

`func (o *CheckEditable) GetCheckBottom() string`

GetCheckBottom returns the CheckBottom field if non-nil, zero value otherwise.

### GetCheckBottomOk

`func (o *CheckEditable) GetCheckBottomOk() (*string, bool)`

GetCheckBottomOk returns a tuple with the CheckBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckBottom

`func (o *CheckEditable) SetCheckBottom(v string)`

SetCheckBottom sets CheckBottom field to given value.

### HasCheckBottom

`func (o *CheckEditable) HasCheckBottom() bool`

HasCheckBottom returns a boolean if a field has been set.

### GetAttachment

`func (o *CheckEditable) GetAttachment() string`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *CheckEditable) GetAttachmentOk() (*string, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *CheckEditable) SetAttachment(v string)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *CheckEditable) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetDescription

`func (o *CheckEditable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckEditable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckEditable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CheckEditable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CheckEditable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CheckEditable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *CheckEditable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckEditable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckEditable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckEditable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMergeVariables

`func (o *CheckEditable) GetMergeVariables() map[string]interface{}`

GetMergeVariables returns the MergeVariables field if non-nil, zero value otherwise.

### GetMergeVariablesOk

`func (o *CheckEditable) GetMergeVariablesOk() (*map[string]interface{}, bool)`

GetMergeVariablesOk returns a tuple with the MergeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariables

`func (o *CheckEditable) SetMergeVariables(v map[string]interface{})`

SetMergeVariables sets MergeVariables field to given value.

### HasMergeVariables

`func (o *CheckEditable) HasMergeVariables() bool`

HasMergeVariables returns a boolean if a field has been set.

### SetMergeVariablesNil

`func (o *CheckEditable) SetMergeVariablesNil(b bool)`

 SetMergeVariablesNil sets the value for MergeVariables to be an explicit nil

### UnsetMergeVariables
`func (o *CheckEditable) UnsetMergeVariables()`

UnsetMergeVariables ensures that no value is present for MergeVariables, not even an explicit nil
### GetSendDate

`func (o *CheckEditable) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *CheckEditable) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *CheckEditable) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *CheckEditable) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### GetMailType

`func (o *CheckEditable) GetMailType() string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *CheckEditable) GetMailTypeOk() (*string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *CheckEditable) SetMailType(v string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *CheckEditable) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMemo

`func (o *CheckEditable) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CheckEditable) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CheckEditable) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CheckEditable) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *CheckEditable) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *CheckEditable) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetCheckNumber

`func (o *CheckEditable) GetCheckNumber() int32`

GetCheckNumber returns the CheckNumber field if non-nil, zero value otherwise.

### GetCheckNumberOk

`func (o *CheckEditable) GetCheckNumberOk() (*int32, bool)`

GetCheckNumberOk returns a tuple with the CheckNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumber

`func (o *CheckEditable) SetCheckNumber(v int32)`

SetCheckNumber sets CheckNumber field to given value.

### HasCheckNumber

`func (o *CheckEditable) HasCheckNumber() bool`

HasCheckNumber returns a boolean if a field has been set.

### GetMessage

`func (o *CheckEditable) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CheckEditable) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CheckEditable) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CheckEditable) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *CheckEditable) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *CheckEditable) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *CheckEditable) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *CheckEditable) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


