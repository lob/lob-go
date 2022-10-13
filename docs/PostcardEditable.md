# PostcardEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | Must either be an address ID or an inline object with correct address parameters. | 
**From** | Pointer to **string** | Required if &#x60;to&#x60; address is international. Must either be an address ID or an inline object with correct address parameters. | [optional] 
**Size** | Pointer to [**PostcardSize**](PostcardSize.md) |  | [optional] [default to _4X6]
**Description** | Pointer to **NullableString** | An internal description that identifies this resource. Must be no longer than 255 characters.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters &#x60;\&quot;&#x60; and &#x60;\\&#x60;. i.e. &#39;{\&quot;customer_id\&quot; : \&quot;NEWYORK2015\&quot;}&#39; Nested objects are not supported.  See [Metadata](#section/Metadata) for more information. | [optional] 
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to FIRST_CLASS]
**MergeVariables** | Pointer to **map[string]interface{}** | You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: &#x60;{{variable_name}}&#x60;, pass in &#x60;{\&quot;variable_name\&quot;: \&quot;Harry\&quot;}&#x60; to render &#x60;Harry&#x60;. &#x60;merge_variables&#x60; must be an object. Any type of value is accepted as long as the object is valid JSON; you can use &#x60;strings&#x60;, &#x60;numbers&#x60;, &#x60;booleans&#x60;, &#x60;arrays&#x60;, &#x60;objects&#x60;, or &#x60;null&#x60;. The max length of the object is 25,000 characters. If you call &#x60;JSON.stringify&#x60; on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: &#x60;!&#x60;, &#x60;\&quot;&#x60;, &#x60;#&#x60;, &#x60;%&#x60;, &#x60;&amp;&#x60;, &#x60;&#39;&#x60;, &#x60;(&#x60;, &#x60;)&#x60;, &#x60;*&#x60;, &#x60;+&#x60;, &#x60;,&#x60;, &#x60;/&#x60;, &#x60;;&#x60;, &#x60;&lt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;@&#x60;, &#x60;[&#x60;, &#x60;\\&#x60;, &#x60;]&#x60;, &#x60;^&#x60;, &#x60;&#x60; &#x60; &#x60;&#x60;, &#x60;{&#x60;, &#x60;|&#x60;, &#x60;}&#x60;, &#x60;~&#x60;. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string. | [optional] 
**SendDate** | Pointer to **time.Time** | A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the &#x60;send_date&#x60; has passed, the mailpiece can be canceled. If a date in the format &#x60;2017-11-01&#x60; is passed, it will evaluate to midnight UTC of that date (&#x60;2017-11-01T00:00:00.000Z&#x60;). If a datetime is passed, that exact time will be used. A &#x60;send_date&#x60; passed with no time zone will default to UTC, while a &#x60;send_date&#x60; passed with a time zone will be converted to UTC. | [optional] 
**Front** | **string** | The artwork to use as the front of your postcard.  | 
**Back** | **string** | The artwork to use as the back of your postcard.  | 
**BillingGroupId** | Pointer to **string** | An optional string with the billing group ID to tag your usage with. Is used for billing purposes. Requires special activation to use. See [Billing Group API](https://lob.github.io/lob-openapi/#tag/Billing-Groups) for more information. | [optional] 

## Methods

### NewPostcardEditable

`func NewPostcardEditable(to string, front string, back string, ) *PostcardEditable`

NewPostcardEditable instantiates a new PostcardEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostcardEditableWithDefaults

`func NewPostcardEditableWithDefaults() *PostcardEditable`

NewPostcardEditableWithDefaults instantiates a new PostcardEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *PostcardEditable) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PostcardEditable) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PostcardEditable) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *PostcardEditable) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PostcardEditable) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PostcardEditable) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PostcardEditable) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSize

`func (o *PostcardEditable) GetSize() PostcardSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostcardEditable) GetSizeOk() (*PostcardSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostcardEditable) SetSize(v PostcardSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *PostcardEditable) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDescription

`func (o *PostcardEditable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostcardEditable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostcardEditable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostcardEditable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PostcardEditable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PostcardEditable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *PostcardEditable) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostcardEditable) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostcardEditable) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PostcardEditable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMailType

`func (o *PostcardEditable) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *PostcardEditable) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *PostcardEditable) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *PostcardEditable) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMergeVariables

`func (o *PostcardEditable) GetMergeVariables() map[string]interface{}`

GetMergeVariables returns the MergeVariables field if non-nil, zero value otherwise.

### GetMergeVariablesOk

`func (o *PostcardEditable) GetMergeVariablesOk() (*map[string]interface{}, bool)`

GetMergeVariablesOk returns a tuple with the MergeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeVariables

`func (o *PostcardEditable) SetMergeVariables(v map[string]interface{})`

SetMergeVariables sets MergeVariables field to given value.

### HasMergeVariables

`func (o *PostcardEditable) HasMergeVariables() bool`

HasMergeVariables returns a boolean if a field has been set.

### SetMergeVariablesNil

`func (o *PostcardEditable) SetMergeVariablesNil(b bool)`

 SetMergeVariablesNil sets the value for MergeVariables to be an explicit nil

### UnsetMergeVariables
`func (o *PostcardEditable) UnsetMergeVariables()`

UnsetMergeVariables ensures that no value is present for MergeVariables, not even an explicit nil
### GetSendDate

`func (o *PostcardEditable) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *PostcardEditable) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *PostcardEditable) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *PostcardEditable) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### GetFront

`func (o *PostcardEditable) GetFront() string`

GetFront returns the Front field if non-nil, zero value otherwise.

### GetFrontOk

`func (o *PostcardEditable) GetFrontOk() (*string, bool)`

GetFrontOk returns a tuple with the Front field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFront

`func (o *PostcardEditable) SetFront(v string)`

SetFront sets Front field to given value.


### GetBack

`func (o *PostcardEditable) GetBack() string`

GetBack returns the Back field if non-nil, zero value otherwise.

### GetBackOk

`func (o *PostcardEditable) GetBackOk() (*string, bool)`

GetBackOk returns a tuple with the Back field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBack

`func (o *PostcardEditable) SetBack(v string)`

SetBack sets Back field to given value.


### GetBillingGroupId

`func (o *PostcardEditable) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *PostcardEditable) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *PostcardEditable) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *PostcardEditable) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


