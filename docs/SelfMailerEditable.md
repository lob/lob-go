# SelfMailerEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **interface{}** | Must either be an address ID or an inline object with correct address parameters. | 
**From** | Pointer to **interface{}** | Must either be an address ID or an inline object with correct address parameters. | [optional] 
**Size** | Pointer to [**SelfMailerSize**](SelfMailerSize.md) |  | [optional] [default to SELFMAILERSIZE__6X18_BIFOLD]
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to MAILTYPE_FIRST_CLASS]
**MergeVariables** | Pointer to **map[string]interface{}** | You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: &#x60;{{variable_name}}&#x60;, pass in &#x60;{\&quot;variable_name\&quot;: \&quot;Harry\&quot;}&#x60; to render &#x60;Harry&#x60;. &#x60;merge_variables&#x60; must be an object. Any type of value is accepted as long as the object is valid JSON; you can use &#x60;strings&#x60;, &#x60;numbers&#x60;, &#x60;booleans&#x60;, &#x60;arrays&#x60;, &#x60;objects&#x60;, or &#x60;null&#x60;. The max length of the object is 25,000 characters. If you call &#x60;JSON.stringify&#x60; on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: &#x60;!&#x60;, &#x60;\&quot;&#x60;, &#x60;#&#x60;, &#x60;%&#x60;, &#x60;&amp;&#x60;, &#x60;&#39;&#x60;, &#x60;(&#x60;, &#x60;)&#x60;, &#x60;*&#x60;, &#x60;+&#x60;, &#x60;,&#x60;, &#x60;/&#x60;, &#x60;;&#x60;, &#x60;&lt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;@&#x60;, &#x60;[&#x60;, &#x60;\\&#x60;, &#x60;]&#x60;, &#x60;^&#x60;, &#x60;&#x60; &#x60; &#x60;&#x60;, &#x60;{&#x60;, &#x60;|&#x60;, &#x60;}&#x60;, &#x60;~&#x60;. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string. | [optional] 
**SendDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the &#x60;send_date&#x60; has passed, the mailpiece can be canceled. If a date in the format &#x60;2017-11-01&#x60; is passed, it will evaluate to midnight UTC of that date (&#x60;2017-11-01T00:00:00.000Z&#x60;). If a datetime is passed, that exact time will be used. A &#x60;send_date&#x60; passed with no time zone will default to UTC, while a &#x60;send_date&#x60; passed with a time zone will be converted to UTC. | [optional] 
**Inside** | **string** | The artwork to use as the inside of your self mailer.  | 
**Outside** | **string** | The artwork to use as the outside of your self mailer.  | 
**BillingGroupId** | Pointer to **string** | An optional string with the billing group ID to tag your usage with. Is used for billing purposes. Requires special activation to use. See [Billing Group API](https://lob.github.io/lob-openapi/#tag/Billing-Groups) for more information. | [optional] 
**UseType** | [**NullableSfmUseType**](SfmUseType.md) |  | 

## Methods

### NewSelfMailerEditable

`func NewSelfMailerEditable(to interface{}, inside string, outside string, useType NullableSfmUseType, ) *SelfMailerEditable`

NewSelfMailerEditable instantiates a new SelfMailerEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfMailerEditableWithDefaults

`func NewSelfMailerEditableWithDefaults() *SelfMailerEditable`

NewSelfMailerEditableWithDefaults instantiates a new SelfMailerEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *SelfMailerEditable) GetTo() interface{}`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SelfMailerEditable) GetToOk() (*interface{}, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SelfMailerEditable) SetTo(v interface{})`

SetTo sets To field to given value.


### SetToNil

`func (o *SelfMailerEditable) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *SelfMailerEditable) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetFrom

`func (o *SelfMailerEditable) GetFrom() interface{}`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SelfMailerEditable) GetFromOk() (*interface{}, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SelfMailerEditable) SetFrom(v interface{})`

SetFrom sets From field to given value.

### HasFrom

`func (o *SelfMailerEditable) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *SelfMailerEditable) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *SelfMailerEditable) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetSize

`func (o *SelfMailerEditable) GetSize() SelfMailerSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SelfMailerEditable) GetSizeOk() (*SelfMailerSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SelfMailerEditable) SetSize(v SelfMailerSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *SelfMailerEditable) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDescription

`func (o *SelfMailerEditable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SelfMailerEditable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SelfMailerEditable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SelfMailerEditable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SelfMailerEditable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SelfMailerEditable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *SelfMailerEditable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SelfMailerEditable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SelfMailerEditable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SelfMailerEditable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMailType

`func (o *SelfMailerEditable) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *SelfMailerEditable) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *SelfMailerEditable) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *SelfMailerEditable) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMergeVariables

`func (o *SelfMailerEditable) GetMergeVariables() map[string]interface{}`

GetMergeVariables returns the MergeVariables field if non-nil, zero value otherwise.

### GetMergeVariablesOk

`func (o *SelfMailerEditable) GetMergeVariablesOk() (*map[string]interface{}, bool)`

GetMergeVariablesOk returns a tuple with the MergeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariables

`func (o *SelfMailerEditable) SetMergeVariables(v map[string]interface{})`

SetMergeVariables sets MergeVariables field to given value.

### HasMergeVariables

`func (o *SelfMailerEditable) HasMergeVariables() bool`

HasMergeVariables returns a boolean if a field has been set.

### SetMergeVariablesNil

`func (o *SelfMailerEditable) SetMergeVariablesNil(b bool)`

 SetMergeVariablesNil sets the value for MergeVariables to be an explicit nil

### UnsetMergeVariables
`func (o *SelfMailerEditable) UnsetMergeVariables()`

UnsetMergeVariables ensures that no value is present for MergeVariables, not even an explicit nil
### GetSendDate

`func (o *SelfMailerEditable) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *SelfMailerEditable) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *SelfMailerEditable) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *SelfMailerEditable) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### GetInside

`func (o *SelfMailerEditable) GetInside() string`

GetInside returns the Inside field if non-nil, zero value otherwise.

### GetInsideOk

`func (o *SelfMailerEditable) GetInsideOk() (*string, bool)`

GetInsideOk returns a tuple with the Inside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInside

`func (o *SelfMailerEditable) SetInside(v string)`

SetInside sets Inside field to given value.


### GetOutside

`func (o *SelfMailerEditable) GetOutside() string`

GetOutside returns the Outside field if non-nil, zero value otherwise.

### GetOutsideOk

`func (o *SelfMailerEditable) GetOutsideOk() (*string, bool)`

GetOutsideOk returns a tuple with the Outside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutside

`func (o *SelfMailerEditable) SetOutside(v string)`

SetOutside sets Outside field to given value.


### GetBillingGroupId

`func (o *SelfMailerEditable) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *SelfMailerEditable) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *SelfMailerEditable) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *SelfMailerEditable) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetUseType

`func (o *SelfMailerEditable) GetUseType() SfmUseType`

GetUseType returns the UseType field if non-nil, zero value otherwise.

### GetUseTypeOk

`func (o *SelfMailerEditable) GetUseTypeOk() (*SfmUseType, bool)`

GetUseTypeOk returns a tuple with the UseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseType

`func (o *SelfMailerEditable) SetUseType(v SfmUseType)`

SetUseType sets UseType field to given value.


### SetUseTypeNil

`func (o *SelfMailerEditable) SetUseTypeNil(b bool)`

 SetUseTypeNil sets the value for UseType to be an explicit nil

### UnsetUseType
`func (o *SelfMailerEditable) UnsetUseType()`

UnsetUseType ensures that no value is present for UseType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


