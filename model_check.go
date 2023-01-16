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

// Check struct for Check
type Check struct {
	// Unique identifier prefixed with `chk_`.
	Id string `json:"id"`
	To Address `json:"to"`
	From *Address `json:"from,omitempty"`
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
	// An integer that designates the check number. If `check_number` is not provided, checks created from a new `bank_account` will start at `10000` and increment with each check created with the `bank_account`. A provided `check_number` overrides the defaults. Subsequent checks created with the same `bank_account` will increment from the provided check number.
	CheckNumber *int32 `json:"check_number,omitempty"`
	// Max of 400 characters to be included at the bottom of the check page.
	Message *string `json:"message,omitempty"`
	// The payment amount to be sent in US dollars.
	Amount float32 `json:"amount"`
	BankAccount BankAccount `json:"bank_account"`
	// Unique identifier prefixed with `tmpl_`. ID of a saved [HTML template](#section/HTML-Templates).
	CheckBottomTemplateId *string `json:"check_bottom_template_id,omitempty"`
	// Unique identifier prefixed with `tmpl_`. ID of a saved [HTML template](#section/HTML-Templates).
	AttachmentTemplateId *string `json:"attachment_template_id,omitempty"`
	// Unique identifier prefixed with `vrsn_`.
	CheckBottomTemplateVersionId *string `json:"check_bottom_template_version_id,omitempty"`
	// Unique identifier prefixed with `vrsn_`.
	AttachmentTemplateVersionId *string `json:"attachment_template_version_id,omitempty"`
	// A [signed link](#section/Asset-URLs) served over HTTPS. The link returned will expire in 30 days to prevent mis-sharing. Each time a GET request is initiated, a new signed URL will be generated.
	Url string `json:"url"`
	Carrier string `json:"carrier"`
	Thumbnails []Thumbnail `json:"thumbnails,omitempty"`
	// A date in YYYY-MM-DD format of the mailpiece's expected delivery date based on its `send_date`.
	ExpectedDeliveryDate *string `json:"expected_delivery_date,omitempty"`
	// An array of tracking_event objects ordered by ascending `time`. Will not be populated for checks created in test mode.
	TrackingEvents []TrackingEventNormal `json:"tracking_events,omitempty"`
	Object string `json:"object"`
	// A timestamp in ISO 8601 format of the date the resource was created.
	DateCreated time.Time `json:"date_created"`
	// A timestamp in ISO 8601 format of the date the resource was last modified.
	DateModified time.Time `json:"date_modified"`
	// Only returned if the resource has been successfully deleted.
	Deleted *bool `json:"deleted,omitempty"`
	UseType NullableChkUseType `json:"use_type"`
}

// NewCheck instantiates a new Check object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheck(id string, to Address, amount float32, bankAccount BankAccount, url string, carrier string, object string, dateCreated time.Time, dateModified time.Time, useType NullableChkUseType) *Check {
	this := Check{}
	this.Id = id
	this.To = to
	var mailType string = "usps_first_class"
	this.MailType = &mailType
	this.Amount = amount
	this.BankAccount = bankAccount
	this.Url = url
	this.Carrier = carrier
	this.Object = object
	this.DateCreated = dateCreated
	this.DateModified = dateModified
	this.UseType = useType
	return &this
}

// NewCheckWithDefaults instantiates a new Check object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckWithDefaults() *Check {
	this := Check{}
	var mailType string = "usps_first_class"
	this.MailType = &mailType
	var carrier string = "USPS"
	this.Carrier = carrier
	var object string = "check"
	this.Object = object
	return &this
}

// GetId returns the Id field value
func (o *Check) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Check) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Check) SetId(v string) {
	o.Id = v
}

// GetTo returns the To field value
func (o *Check) GetTo() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Check) GetToOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *Check) SetTo(v Address) {
	o.To = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *Check) GetFrom() Address {
	if o == nil || o.From == nil {
		var ret Address
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetFromOk() (*Address, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *Check) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given Address and assigns it to the From field.
func (o *Check) SetFrom(v Address) {
	o.From = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Check) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Check) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Check) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Check) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Check) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Check) UnsetDescription() {
	o.Description.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Check) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Check) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Check) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMergeVariables returns the MergeVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Check) GetMergeVariables() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MergeVariables
}

// GetMergeVariablesOk returns a tuple with the MergeVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Check) GetMergeVariablesOk() (map[string]interface{}, bool) {
	if o == nil || o.MergeVariables == nil {
		return nil, false
	}
	return o.MergeVariables, true
}

// HasMergeVariables returns a boolean if a field has been set.
func (o *Check) HasMergeVariables() bool {
	if o != nil && o.MergeVariables != nil {
		return true
	}

	return false
}

// SetMergeVariables gets a reference to the given map[string]interface{} and assigns it to the MergeVariables field.
func (o *Check) SetMergeVariables(v map[string]interface{}) {
	o.MergeVariables = v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *Check) GetSendDate() time.Time {
	if o == nil || o.SendDate == nil {
		var ret time.Time
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetSendDateOk() (*time.Time, bool) {
	if o == nil || o.SendDate == nil {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *Check) HasSendDate() bool {
	if o != nil && o.SendDate != nil {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given time.Time and assigns it to the SendDate field.
func (o *Check) SetSendDate(v time.Time) {
	o.SendDate = &v
}

// GetMailType returns the MailType field value if set, zero value otherwise.
func (o *Check) GetMailType() string {
	if o == nil || o.MailType == nil {
		var ret string
		return ret
	}
	return *o.MailType
}

// GetMailTypeOk returns a tuple with the MailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetMailTypeOk() (*string, bool) {
	if o == nil || o.MailType == nil {
		return nil, false
	}
	return o.MailType, true
}

// HasMailType returns a boolean if a field has been set.
func (o *Check) HasMailType() bool {
	if o != nil && o.MailType != nil {
		return true
	}

	return false
}

// SetMailType gets a reference to the given string and assigns it to the MailType field.
func (o *Check) SetMailType(v string) {
	o.MailType = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Check) GetMemo() string {
	if o == nil || o.Memo.Get() == nil {
		var ret string
		return ret
	}
	return *o.Memo.Get()
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Check) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memo.Get(), o.Memo.IsSet()
}

// HasMemo returns a boolean if a field has been set.
func (o *Check) HasMemo() bool {
	if o != nil && o.Memo.IsSet() {
		return true
	}

	return false
}

// SetMemo gets a reference to the given NullableString and assigns it to the Memo field.
func (o *Check) SetMemo(v string) {
	o.Memo.Set(&v)
}
// SetMemoNil sets the value for Memo to be an explicit nil
func (o *Check) SetMemoNil() {
	o.Memo.Set(nil)
}

// UnsetMemo ensures that no value is present for Memo, not even an explicit nil
func (o *Check) UnsetMemo() {
	o.Memo.Unset()
}

// GetCheckNumber returns the CheckNumber field value if set, zero value otherwise.
func (o *Check) GetCheckNumber() int32 {
	if o == nil || o.CheckNumber == nil {
		var ret int32
		return ret
	}
	return *o.CheckNumber
}

// GetCheckNumberOk returns a tuple with the CheckNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetCheckNumberOk() (*int32, bool) {
	if o == nil || o.CheckNumber == nil {
		return nil, false
	}
	return o.CheckNumber, true
}

// HasCheckNumber returns a boolean if a field has been set.
func (o *Check) HasCheckNumber() bool {
	if o != nil && o.CheckNumber != nil {
		return true
	}

	return false
}

// SetCheckNumber gets a reference to the given int32 and assigns it to the CheckNumber field.
func (o *Check) SetCheckNumber(v int32) {
	o.CheckNumber = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Check) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Check) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Check) SetMessage(v string) {
	o.Message = &v
}

// GetAmount returns the Amount field value
func (o *Check) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Check) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Check) SetAmount(v float32) {
	o.Amount = v
}

// GetBankAccount returns the BankAccount field value
func (o *Check) GetBankAccount() BankAccount {
	if o == nil {
		var ret BankAccount
		return ret
	}

	return o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value
// and a boolean to check if the value has been set.
func (o *Check) GetBankAccountOk() (*BankAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccount, true
}

// SetBankAccount sets field value
func (o *Check) SetBankAccount(v BankAccount) {
	o.BankAccount = v
}

// GetCheckBottomTemplateId returns the CheckBottomTemplateId field value if set, zero value otherwise.
func (o *Check) GetCheckBottomTemplateId() string {
	if o == nil || o.CheckBottomTemplateId == nil {
		var ret string
		return ret
	}
	return *o.CheckBottomTemplateId
}

// GetCheckBottomTemplateIdOk returns a tuple with the CheckBottomTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetCheckBottomTemplateIdOk() (*string, bool) {
	if o == nil || o.CheckBottomTemplateId == nil {
		return nil, false
	}
	return o.CheckBottomTemplateId, true
}

// HasCheckBottomTemplateId returns a boolean if a field has been set.
func (o *Check) HasCheckBottomTemplateId() bool {
	if o != nil && o.CheckBottomTemplateId != nil {
		return true
	}

	return false
}

// SetCheckBottomTemplateId gets a reference to the given string and assigns it to the CheckBottomTemplateId field.
func (o *Check) SetCheckBottomTemplateId(v string) {
	o.CheckBottomTemplateId = &v
}

// GetAttachmentTemplateId returns the AttachmentTemplateId field value if set, zero value otherwise.
func (o *Check) GetAttachmentTemplateId() string {
	if o == nil || o.AttachmentTemplateId == nil {
		var ret string
		return ret
	}
	return *o.AttachmentTemplateId
}

// GetAttachmentTemplateIdOk returns a tuple with the AttachmentTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetAttachmentTemplateIdOk() (*string, bool) {
	if o == nil || o.AttachmentTemplateId == nil {
		return nil, false
	}
	return o.AttachmentTemplateId, true
}

// HasAttachmentTemplateId returns a boolean if a field has been set.
func (o *Check) HasAttachmentTemplateId() bool {
	if o != nil && o.AttachmentTemplateId != nil {
		return true
	}

	return false
}

// SetAttachmentTemplateId gets a reference to the given string and assigns it to the AttachmentTemplateId field.
func (o *Check) SetAttachmentTemplateId(v string) {
	o.AttachmentTemplateId = &v
}

// GetCheckBottomTemplateVersionId returns the CheckBottomTemplateVersionId field value if set, zero value otherwise.
func (o *Check) GetCheckBottomTemplateVersionId() string {
	if o == nil || o.CheckBottomTemplateVersionId == nil {
		var ret string
		return ret
	}
	return *o.CheckBottomTemplateVersionId
}

// GetCheckBottomTemplateVersionIdOk returns a tuple with the CheckBottomTemplateVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetCheckBottomTemplateVersionIdOk() (*string, bool) {
	if o == nil || o.CheckBottomTemplateVersionId == nil {
		return nil, false
	}
	return o.CheckBottomTemplateVersionId, true
}

// HasCheckBottomTemplateVersionId returns a boolean if a field has been set.
func (o *Check) HasCheckBottomTemplateVersionId() bool {
	if o != nil && o.CheckBottomTemplateVersionId != nil {
		return true
	}

	return false
}

// SetCheckBottomTemplateVersionId gets a reference to the given string and assigns it to the CheckBottomTemplateVersionId field.
func (o *Check) SetCheckBottomTemplateVersionId(v string) {
	o.CheckBottomTemplateVersionId = &v
}

// GetAttachmentTemplateVersionId returns the AttachmentTemplateVersionId field value if set, zero value otherwise.
func (o *Check) GetAttachmentTemplateVersionId() string {
	if o == nil || o.AttachmentTemplateVersionId == nil {
		var ret string
		return ret
	}
	return *o.AttachmentTemplateVersionId
}

// GetAttachmentTemplateVersionIdOk returns a tuple with the AttachmentTemplateVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetAttachmentTemplateVersionIdOk() (*string, bool) {
	if o == nil || o.AttachmentTemplateVersionId == nil {
		return nil, false
	}
	return o.AttachmentTemplateVersionId, true
}

// HasAttachmentTemplateVersionId returns a boolean if a field has been set.
func (o *Check) HasAttachmentTemplateVersionId() bool {
	if o != nil && o.AttachmentTemplateVersionId != nil {
		return true
	}

	return false
}

// SetAttachmentTemplateVersionId gets a reference to the given string and assigns it to the AttachmentTemplateVersionId field.
func (o *Check) SetAttachmentTemplateVersionId(v string) {
	o.AttachmentTemplateVersionId = &v
}

// GetUrl returns the Url field value
func (o *Check) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Check) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Check) SetUrl(v string) {
	o.Url = v
}

// GetCarrier returns the Carrier field value
func (o *Check) GetCarrier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value
// and a boolean to check if the value has been set.
func (o *Check) GetCarrierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Carrier, true
}

// SetCarrier sets field value
func (o *Check) SetCarrier(v string) {
	o.Carrier = v
}

// GetThumbnails returns the Thumbnails field value if set, zero value otherwise.
func (o *Check) GetThumbnails() []Thumbnail {
	if o == nil || o.Thumbnails == nil {
		var ret []Thumbnail
		return ret
	}
	return o.Thumbnails
}

// GetThumbnailsOk returns a tuple with the Thumbnails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetThumbnailsOk() ([]Thumbnail, bool) {
	if o == nil || o.Thumbnails == nil {
		return nil, false
	}
	return o.Thumbnails, true
}

// HasThumbnails returns a boolean if a field has been set.
func (o *Check) HasThumbnails() bool {
	if o != nil && o.Thumbnails != nil {
		return true
	}

	return false
}

// SetThumbnails gets a reference to the given []Thumbnail and assigns it to the Thumbnails field.
func (o *Check) SetThumbnails(v []Thumbnail) {
	o.Thumbnails = v
}

// GetExpectedDeliveryDate returns the ExpectedDeliveryDate field value if set, zero value otherwise.
func (o *Check) GetExpectedDeliveryDate() string {
	if o == nil || o.ExpectedDeliveryDate == nil {
		var ret string
		return ret
	}
	return *o.ExpectedDeliveryDate
}

// GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetExpectedDeliveryDateOk() (*string, bool) {
	if o == nil || o.ExpectedDeliveryDate == nil {
		return nil, false
	}
	return o.ExpectedDeliveryDate, true
}

// HasExpectedDeliveryDate returns a boolean if a field has been set.
func (o *Check) HasExpectedDeliveryDate() bool {
	if o != nil && o.ExpectedDeliveryDate != nil {
		return true
	}

	return false
}

// SetExpectedDeliveryDate gets a reference to the given string and assigns it to the ExpectedDeliveryDate field.
func (o *Check) SetExpectedDeliveryDate(v string) {
	o.ExpectedDeliveryDate = &v
}

// GetTrackingEvents returns the TrackingEvents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Check) GetTrackingEvents() []TrackingEventNormal {
	if o == nil {
		var ret []TrackingEventNormal
		return ret
	}
	return o.TrackingEvents
}

// GetTrackingEventsOk returns a tuple with the TrackingEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Check) GetTrackingEventsOk() ([]TrackingEventNormal, bool) {
	if o == nil || o.TrackingEvents == nil {
		return nil, false
	}
	return o.TrackingEvents, true
}

// HasTrackingEvents returns a boolean if a field has been set.
func (o *Check) HasTrackingEvents() bool {
	if o != nil && o.TrackingEvents != nil {
		return true
	}

	return false
}

// SetTrackingEvents gets a reference to the given []TrackingEventNormal and assigns it to the TrackingEvents field.
func (o *Check) SetTrackingEvents(v []TrackingEventNormal) {
	o.TrackingEvents = v
}

// GetObject returns the Object field value
func (o *Check) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Check) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Check) SetObject(v string) {
	o.Object = v
}

// GetDateCreated returns the DateCreated field value
func (o *Check) GetDateCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *Check) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *Check) SetDateCreated(v time.Time) {
	o.DateCreated = v
}

// GetDateModified returns the DateModified field value
func (o *Check) GetDateModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value
// and a boolean to check if the value has been set.
func (o *Check) GetDateModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateModified, true
}

// SetDateModified sets field value
func (o *Check) SetDateModified(v time.Time) {
	o.DateModified = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Check) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Check) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Check) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetUseType returns the UseType field value
// If the value is explicit nil, the zero value for ChkUseType will be returned
func (o *Check) GetUseType() ChkUseType {
	if o == nil || o.UseType.Get() == nil {
		var ret ChkUseType
		return ret
	}

	return *o.UseType.Get()
}

// GetUseTypeOk returns a tuple with the UseType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Check) GetUseTypeOk() (*ChkUseType, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseType.Get(), o.UseType.IsSet()
}

// SetUseType sets field value
func (o *Check) SetUseType(v ChkUseType) {
	o.UseType.Set(&v)
}

func (o Check) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["to"] = o.To
	}
	if o.From != nil {
		toSerialize["from"] = o.From
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
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["bank_account"] = o.BankAccount
	}
	if o.CheckBottomTemplateId != nil {
		toSerialize["check_bottom_template_id"] = o.CheckBottomTemplateId
	}
	if o.AttachmentTemplateId != nil {
		toSerialize["attachment_template_id"] = o.AttachmentTemplateId
	}
	if o.CheckBottomTemplateVersionId != nil {
		toSerialize["check_bottom_template_version_id"] = o.CheckBottomTemplateVersionId
	}
	if o.AttachmentTemplateVersionId != nil {
		toSerialize["attachment_template_version_id"] = o.AttachmentTemplateVersionId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["carrier"] = o.Carrier
	}
	if o.Thumbnails != nil {
		toSerialize["thumbnails"] = o.Thumbnails
	}
	if o.ExpectedDeliveryDate != nil {
		toSerialize["expected_delivery_date"] = o.ExpectedDeliveryDate
	}
	if o.TrackingEvents != nil {
		toSerialize["tracking_events"] = o.TrackingEvents
	}
	if true {
		toSerialize["object"] = o.Object
	}
	if true {
		toSerialize["date_created"] = o.DateCreated
	}
	if true {
		toSerialize["date_modified"] = o.DateModified
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if true {
		toSerialize["use_type"] = o.UseType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCheck struct {
	value *Check
	isSet bool
}

func (v NullableCheck) Get() *Check {
	return v.value
}

func (v *NullableCheck) Set(val *Check) {
	v.value = val
	v.isSet = true
}

func (v NullableCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheck(val *Check) *NullableCheck {
	return &NullableCheck{value: val, isSet: true}
}

func (v NullableCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


