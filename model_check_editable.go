/*
Lob

The Lob API is organized around REST. Our API is designed to have predictable, resource-oriented URLs and uses HTTP response codes to indicate any API errors. <p> Looking for our [previous documentation](https://lob.github.io/legacy-docs/)? 

API version: 1.3.0
Contact: lob-openapi@lob.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lob

import (
	"encoding/json"
	
	"time"
)

// CheckEditable struct for CheckEditable
type CheckEditable struct {
	// Must either be an address ID or an inline object with correct address parameters.
	From interface{} `json:"from"`
	// Must either be an address ID or an inline object with correct address parameters.
	To interface{} `json:"to"`
	BankAccount NullableString `json:"bank_account"`
	// The payment amount to be sent in US dollars.
	Amount float32 `json:"amount"`
	// Accepts a remote URL or local file upload to an image to print (in grayscale) in the upper-left corner of your check.
	Logo *string `json:"logo,omitempty"`
	// The artwork to use on the bottom of the check page.  Notes: - HTML merge variables should not include delimiting whitespace. - PDF, PNG, and JPGs must be sized at 8.5\"x11\" at 300 DPI, while supplied HTML will be rendered and trimmed to fit on a 8.5\"x11\" page. - The check bottom will always be printed in black & white. - Must conform to [this template](https://s3-us-west-2.amazonaws.com/public.lob.com/assets/templates/check_bottom_template.pdf).  Need more help? Consult our [HTML examples](#section/HTML-Examples).
	CheckBottom *string `json:"check_bottom,omitempty"`
	// A document to include with the check.
	Attachment *string `json:"attachment,omitempty"`
	// An internal description that identifies this resource. Must be no longer than 255 characters. 
	Description NullableString `json:"description,omitempty"`
	// Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters `\"` and `\\`. i.e. '{\"customer_id\" : \"NEWYORK2015\"}' Nested objects are not supported.  See [Metadata](#section/Metadata) for more information.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: `{{variable_name}}`, pass in `{\"variable_name\": \"Harry\"}` to render `Harry`. `merge_variables` must be an object. Any type of value is accepted as long as the object is valid JSON; you can use `strings`, `numbers`, `booleans`, `arrays`, `objects`, or `null`. The max length of the object is 25,000 characters. If you call `JSON.stringify` on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: `!`, `\"`, `#`, `%`, `&`, `'`, `(`, `)`, `*`, `+`, `,`, `/`, `;`, `<`, `=`, `>`, `@`, `[`, `\\`, `]`, `^`, `` ` ``, `{`, `|`, `}`, `~`. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string.
	MergeVariables map[string]interface{} `json:"merge_variables,omitempty"`
	// A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the `send_date` has passed, the mailpiece can be canceled. If a date in the format `2017-11-01` is passed, it will evaluate to midnight UTC of that date (`2017-11-01T00:00:00.000Z`). If a datetime is passed, that exact time will be used. A `send_date` passed with no time zone will default to UTC, while a `send_date` passed with a time zone will be converted to UTC.
	SendDate *time.Time `json:"send_date,omitempty"`
	// Checks must be sent `usps_first_class`
	MailType *string `json:"mail_type,omitempty"`
	// Text to include on the memo line of the check.
	Memo NullableString `json:"memo,omitempty"`
	// An integer that designates the check number.
	CheckNumber *int32 `json:"check_number,omitempty"`
	// Max of 400 characters to be included at the bottom of the check page.
	Message *string `json:"message,omitempty"`
	// An optional string with the billing group ID to tag your usage with. Is used for billing purposes. Requires special activation to use. See [Billing Group API](https://lob.github.io/lob-openapi/#tag/Billing-Groups) for more information.
	BillingGroupId *string `json:"billing_group_id,omitempty"`
}

// NewCheckEditable instantiates a new CheckEditable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckEditable(from interface{}, to interface{}, bankAccount NullableString, amount float32) *CheckEditable {
	this := CheckEditable{}
	this.From = from
	this.To = to
	this.BankAccount = bankAccount
	this.Amount = amount
	var mailType string = "usps_first_class"
	this.MailType = &mailType
	return &this
}

// NewCheckEditableWithDefaults instantiates a new CheckEditable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckEditableWithDefaults() *CheckEditable {
	this := CheckEditable{}
	var mailType string = "usps_first_class"
	this.MailType = &mailType
	return &this
}

// GetFrom returns the From field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CheckEditable) GetFrom() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckEditable) GetFromOk() (*interface{}, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CheckEditable) SetFrom(v interface{}) {
	o.From = v
}

// GetTo returns the To field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CheckEditable) GetTo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckEditable) GetToOk() (*interface{}, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *CheckEditable) SetTo(v interface{}) {
	o.To = v
}

// GetBankAccount returns the BankAccount field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckEditable) GetBankAccount() string {
	if o == nil || o.BankAccount.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankAccount.Get()
}

// GetBankAccountOk returns a tuple with the BankAccount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckEditable) GetBankAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankAccount.Get(), o.BankAccount.IsSet()
}

// SetBankAccount sets field value
func (o *CheckEditable) SetBankAccount(v string) {
	o.BankAccount.Set(&v)
}

// GetAmount returns the Amount field value
func (o *CheckEditable) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CheckEditable) SetAmount(v float32) {
	o.Amount = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CheckEditable) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CheckEditable) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *CheckEditable) SetLogo(v string) {
	o.Logo = &v
}

// GetCheckBottom returns the CheckBottom field value if set, zero value otherwise.
func (o *CheckEditable) GetCheckBottom() string {
	if o == nil || o.CheckBottom == nil {
		var ret string
		return ret
	}
	return *o.CheckBottom
}

// GetCheckBottomOk returns a tuple with the CheckBottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetCheckBottomOk() (*string, bool) {
	if o == nil || o.CheckBottom == nil {
		return nil, false
	}
	return o.CheckBottom, true
}

// HasCheckBottom returns a boolean if a field has been set.
func (o *CheckEditable) HasCheckBottom() bool {
	if o != nil && o.CheckBottom != nil {
		return true
	}

	return false
}

// SetCheckBottom gets a reference to the given string and assigns it to the CheckBottom field.
func (o *CheckEditable) SetCheckBottom(v string) {
	o.CheckBottom = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *CheckEditable) GetAttachment() string {
	if o == nil || o.Attachment == nil {
		var ret string
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetAttachmentOk() (*string, bool) {
	if o == nil || o.Attachment == nil {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *CheckEditable) HasAttachment() bool {
	if o != nil && o.Attachment != nil {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given string and assigns it to the Attachment field.
func (o *CheckEditable) SetAttachment(v string) {
	o.Attachment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckEditable) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckEditable) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CheckEditable) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CheckEditable) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CheckEditable) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CheckEditable) UnsetDescription() {
	o.Description.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CheckEditable) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CheckEditable) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CheckEditable) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMergeVariables returns the MergeVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckEditable) GetMergeVariables() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MergeVariables
}

// GetMergeVariablesOk returns a tuple with the MergeVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckEditable) GetMergeVariablesOk() (map[string]interface{}, bool) {
	if o == nil || o.MergeVariables == nil {
		return nil, false
	}
	return o.MergeVariables, true
}

// HasMergeVariables returns a boolean if a field has been set.
func (o *CheckEditable) HasMergeVariables() bool {
	if o != nil && o.MergeVariables != nil {
		return true
	}

	return false
}

// SetMergeVariables gets a reference to the given map[string]interface{} and assigns it to the MergeVariables field.
func (o *CheckEditable) SetMergeVariables(v map[string]interface{}) {
	o.MergeVariables = v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *CheckEditable) GetSendDate() time.Time {
	if o == nil || o.SendDate == nil {
		var ret time.Time
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetSendDateOk() (*time.Time, bool) {
	if o == nil || o.SendDate == nil {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *CheckEditable) HasSendDate() bool {
	if o != nil && o.SendDate != nil {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given time.Time and assigns it to the SendDate field.
func (o *CheckEditable) SetSendDate(v time.Time) {
	o.SendDate = &v
}

// GetMailType returns the MailType field value if set, zero value otherwise.
func (o *CheckEditable) GetMailType() string {
	if o == nil || o.MailType == nil {
		var ret string
		return ret
	}
	return *o.MailType
}

// GetMailTypeOk returns a tuple with the MailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetMailTypeOk() (*string, bool) {
	if o == nil || o.MailType == nil {
		return nil, false
	}
	return o.MailType, true
}

// HasMailType returns a boolean if a field has been set.
func (o *CheckEditable) HasMailType() bool {
	if o != nil && o.MailType != nil {
		return true
	}

	return false
}

// SetMailType gets a reference to the given string and assigns it to the MailType field.
func (o *CheckEditable) SetMailType(v string) {
	o.MailType = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckEditable) GetMemo() string {
	if o == nil || o.Memo.Get() == nil {
		var ret string
		return ret
	}
	return *o.Memo.Get()
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckEditable) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memo.Get(), o.Memo.IsSet()
}

// HasMemo returns a boolean if a field has been set.
func (o *CheckEditable) HasMemo() bool {
	if o != nil && o.Memo.IsSet() {
		return true
	}

	return false
}

// SetMemo gets a reference to the given NullableString and assigns it to the Memo field.
func (o *CheckEditable) SetMemo(v string) {
	o.Memo.Set(&v)
}
// SetMemoNil sets the value for Memo to be an explicit nil
func (o *CheckEditable) SetMemoNil() {
	o.Memo.Set(nil)
}

// UnsetMemo ensures that no value is present for Memo, not even an explicit nil
func (o *CheckEditable) UnsetMemo() {
	o.Memo.Unset()
}

// GetCheckNumber returns the CheckNumber field value if set, zero value otherwise.
func (o *CheckEditable) GetCheckNumber() int32 {
	if o == nil || o.CheckNumber == nil {
		var ret int32
		return ret
	}
	return *o.CheckNumber
}

// GetCheckNumberOk returns a tuple with the CheckNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetCheckNumberOk() (*int32, bool) {
	if o == nil || o.CheckNumber == nil {
		return nil, false
	}
	return o.CheckNumber, true
}

// HasCheckNumber returns a boolean if a field has been set.
func (o *CheckEditable) HasCheckNumber() bool {
	if o != nil && o.CheckNumber != nil {
		return true
	}

	return false
}

// SetCheckNumber gets a reference to the given int32 and assigns it to the CheckNumber field.
func (o *CheckEditable) SetCheckNumber(v int32) {
	o.CheckNumber = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CheckEditable) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CheckEditable) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CheckEditable) SetMessage(v string) {
	o.Message = &v
}

// GetBillingGroupId returns the BillingGroupId field value if set, zero value otherwise.
func (o *CheckEditable) GetBillingGroupId() string {
	if o == nil || o.BillingGroupId == nil {
		var ret string
		return ret
	}
	return *o.BillingGroupId
}

// GetBillingGroupIdOk returns a tuple with the BillingGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckEditable) GetBillingGroupIdOk() (*string, bool) {
	if o == nil || o.BillingGroupId == nil {
		return nil, false
	}
	return o.BillingGroupId, true
}

// HasBillingGroupId returns a boolean if a field has been set.
func (o *CheckEditable) HasBillingGroupId() bool {
	if o != nil && o.BillingGroupId != nil {
		return true
	}

	return false
}

// SetBillingGroupId gets a reference to the given string and assigns it to the BillingGroupId field.
func (o *CheckEditable) SetBillingGroupId(v string) {
	o.BillingGroupId = &v
}

func (o CheckEditable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["bank_account"] = o.BankAccount.Get()
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.CheckBottom != nil {
		toSerialize["check_bottom"] = o.CheckBottom
	}
	if o.Attachment != nil {
		toSerialize["attachment"] = o.Attachment
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MergeVariables != nil {
		toSerialize["merge_variables"] = o.MergeVariables
	}
	if o.SendDate != nil {
		toSerialize["send_date"] = o.SendDate
	}
	if o.MailType != nil {
		toSerialize["mail_type"] = o.MailType
	}
	if o.Memo.IsSet() {
		toSerialize["memo"] = o.Memo.Get()
	}
	if o.CheckNumber != nil {
		toSerialize["check_number"] = o.CheckNumber
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.BillingGroupId != nil {
		toSerialize["billing_group_id"] = o.BillingGroupId
	}
	return json.Marshal(toSerialize)
}

type NullableCheckEditable struct {
	value *CheckEditable
	isSet bool
}

func (v NullableCheckEditable) Get() *CheckEditable {
	return v.value
}

func (v *NullableCheckEditable) Set(val *CheckEditable) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckEditable) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckEditable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckEditable(val *CheckEditable) *NullableCheckEditable {
	return &NullableCheckEditable{value: val, isSet: true}
}

func (v NullableCheckEditable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckEditable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


