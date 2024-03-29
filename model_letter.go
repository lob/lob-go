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

// Letter struct for Letter
type Letter struct {
	To Address `json:"to"`
	From Address `json:"from"`
	Carrier *string `json:"carrier,omitempty"`
	Thumbnails []Thumbnail `json:"thumbnails,omitempty"`
	// A date in YYYY-MM-DD format of the mailpiece's expected delivery date based on its `send_date`.
	ExpectedDeliveryDate *string `json:"expected_delivery_date,omitempty"`
	// A timestamp in ISO 8601 format of the date the resource was created.
	DateCreated time.Time `json:"date_created"`
	// A timestamp in ISO 8601 format of the date the resource was last modified.
	DateModified time.Time `json:"date_modified"`
	// Only returned if the resource has been successfully deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// Unique identifier prefixed with `ltr_`.
	Id string `json:"id"`
	// Unique identifier prefixed with `tmpl_`. ID of a saved [HTML template](#section/HTML-Templates).
	TemplateId *string `json:"template_id,omitempty"`
	// Unique identifier prefixed with `vrsn_`.
	TemplateVersionId *string `json:"template_version_id,omitempty"`
	// A [signed link](#section/Asset-URLs) served over HTTPS. The link returned will expire in 30 days to prevent mis-sharing. Each time a GET request is initiated, a new signed URL will be generated.
	Url *string `json:"url,omitempty"`
	Object string `json:"object"`
	// An internal description that identifies this resource. Must be no longer than 255 characters. 
	Description NullableString `json:"description,omitempty"`
	// Use metadata to store custom information for tagging and labeling back to your internal systems. Must be an object with up to 20 key-value pairs. Keys must be at most 40 characters and values must be at most 500 characters. Neither can contain the characters `\"` and `\\`. i.e. '{\"customer_id\" : \"NEWYORK2015\"}' Nested objects are not supported.  See [Metadata](#section/Metadata) for more information.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// You can input a merge variable payload object to your template to render dynamic content. For example, if you have a template like: `{{variable_name}}`, pass in `{\"variable_name\": \"Harry\"}` to render `Harry`. `merge_variables` must be an object. Any type of value is accepted as long as the object is valid JSON; you can use `strings`, `numbers`, `booleans`, `arrays`, `objects`, or `null`. The max length of the object is 25,000 characters. If you call `JSON.stringify` on your object, it can be no longer than 25,000 characters. Your variable names cannot contain any whitespace or any of the following special characters: `!`, `\"`, `#`, `%`, `&`, `'`, `(`, `)`, `*`, `+`, `,`, `/`, `;`, `<`, `=`, `>`, `@`, `[`, `\\`, `]`, `^`, `` ` ``, `{`, `|`, `}`, `~`. More instructions can be found in [our guide to using html and merge variables](https://lob.com/resources/guides/general/using-html-and-merge-variables). Depending on your [Merge Variable strictness](https://dashboard.lob.com/#/settings/account) setting, if you define variables in your HTML but do not pass them here, you will either receive an error or the variable will render as an empty string.
	MergeVariables map[string]interface{} `json:"merge_variables,omitempty"`
	// A timestamp in ISO 8601 format which specifies a date after the current time and up to 180 days in the future to send the letter off for production. Setting a send date overrides the default [cancellation window](#section/Cancellation-Windows) applied to the mailpiece. Until the `send_date` has passed, the mailpiece can be canceled. If a date in the format `2017-11-01` is passed, it will evaluate to midnight UTC of that date (`2017-11-01T00:00:00.000Z`). If a datetime is passed, that exact time will be used. A `send_date` passed with no time zone will default to UTC, while a `send_date` passed with a time zone will be converted to UTC.
	SendDate *time.Time `json:"send_date,omitempty"`
	// Add an extra service to your letter. See [pricing](https://www.lob.com/pricing/print-mail#compare) for extra costs incurred.
	ExtraService *string `json:"extra_service,omitempty"`
	// The tracking number, if applicable, will appear here when it becomes available. Dummy tracking numbers are not created in test mode.
	TrackingNumber NullableString `json:"tracking_number,omitempty"`
	// Tracking events are not populated for registered or regular (no extra service) letters.
	TrackingEvents []TrackingEventNormal `json:"tracking_events,omitempty"`
	// Specifies the address the return envelope will be sent back to. This is an optional argument that is available if an account is signed up for the return envelope tracking beta, and has `return_envelope`, and `perforated_page` fields populated in the API request.
	ReturnAddress interface{} `json:"return_address,omitempty"`
	MailType *MailType `json:"mail_type,omitempty"`
	// Set this key to `true` if you would like to print in color. Set to `false` if you would like to print in black and white.
	Color *bool `json:"color,omitempty"`
	// Set this attribute to `true` for double sided printing, or `false` for for single sided printing. Defaults to `true`.
	DoubleSided *bool `json:"double_sided,omitempty"`
	// Specifies the location of the address information that will show through the double-window envelope. 
	AddressPlacement *string `json:"address_placement,omitempty"`
	ReturnEnvelope interface{} `json:"return_envelope"`
	// Required if `return_envelope` is `true`. The number of the page that should be perforated for use with the return envelope. Must be greater than or equal to `1`. The blank page added by `address_placement=insert_blank_page` will be ignored when considering the perforated page number. To see how perforation will impact your letter design, view our [perforation guide](https://s3-us-west-2.amazonaws.com/public.lob.com/assets/templates/letter_perf_template.pdf).
	PerforatedPage NullableInt32 `json:"perforated_page,omitempty"`
	CustomEnvelope NullableLetterCustomEnvelope `json:"custom_envelope,omitempty"`
	// The unique ID of the associated campaign if the resource was generated from a campaign.
	CampaignId NullableString `json:"campaign_id,omitempty"`
	UseType NullableLtrUseType `json:"use_type"`
}

// NewLetter instantiates a new Letter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLetter(to Address, from Address, dateCreated time.Time, dateModified time.Time, id string, object string, returnEnvelope interface{}, useType NullableLtrUseType) *Letter {
	this := Letter{}
	this.To = to
	this.From = from
	var carrier string = "USPS"
	this.Carrier = &carrier
	this.DateCreated = dateCreated
	this.DateModified = dateModified
	this.Id = id
	this.Object = object
	var mailType MailType = MAILTYPE_FIRST_CLASS
	this.MailType = &mailType
	var doubleSided bool = true
	this.DoubleSided = &doubleSided
	var addressPlacement string = "top_first_page"
	this.AddressPlacement = &addressPlacement
	this.ReturnEnvelope = returnEnvelope
	this.UseType = useType
	return &this
}

// NewLetterWithDefaults instantiates a new Letter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLetterWithDefaults() *Letter {
	this := Letter{}
	var carrier string = "USPS"
	this.Carrier = &carrier
	var object string = "letter"
	this.Object = object
	var mailType MailType = MAILTYPE_FIRST_CLASS
	this.MailType = &mailType
	var doubleSided bool = true
	this.DoubleSided = &doubleSided
	var addressPlacement string = "top_first_page"
	this.AddressPlacement = &addressPlacement
	return &this
}

// GetTo returns the To field value
func (o *Letter) GetTo() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Letter) GetToOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *Letter) SetTo(v Address) {
	o.To = v
}

// GetFrom returns the From field value
func (o *Letter) GetFrom() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *Letter) GetFromOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *Letter) SetFrom(v Address) {
	o.From = v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *Letter) GetCarrier() string {
	if o == nil || o.Carrier == nil {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetCarrierOk() (*string, bool) {
	if o == nil || o.Carrier == nil {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *Letter) HasCarrier() bool {
	if o != nil && o.Carrier != nil {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *Letter) SetCarrier(v string) {
	o.Carrier = &v
}

// GetThumbnails returns the Thumbnails field value if set, zero value otherwise.
func (o *Letter) GetThumbnails() []Thumbnail {
	if o == nil || o.Thumbnails == nil {
		var ret []Thumbnail
		return ret
	}
	return o.Thumbnails
}

// GetThumbnailsOk returns a tuple with the Thumbnails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetThumbnailsOk() ([]Thumbnail, bool) {
	if o == nil || o.Thumbnails == nil {
		return nil, false
	}
	return o.Thumbnails, true
}

// HasThumbnails returns a boolean if a field has been set.
func (o *Letter) HasThumbnails() bool {
	if o != nil && o.Thumbnails != nil {
		return true
	}

	return false
}

// SetThumbnails gets a reference to the given []Thumbnail and assigns it to the Thumbnails field.
func (o *Letter) SetThumbnails(v []Thumbnail) {
	o.Thumbnails = v
}

// GetExpectedDeliveryDate returns the ExpectedDeliveryDate field value if set, zero value otherwise.
func (o *Letter) GetExpectedDeliveryDate() string {
	if o == nil || o.ExpectedDeliveryDate == nil {
		var ret string
		return ret
	}
	return *o.ExpectedDeliveryDate
}

// GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetExpectedDeliveryDateOk() (*string, bool) {
	if o == nil || o.ExpectedDeliveryDate == nil {
		return nil, false
	}
	return o.ExpectedDeliveryDate, true
}

// HasExpectedDeliveryDate returns a boolean if a field has been set.
func (o *Letter) HasExpectedDeliveryDate() bool {
	if o != nil && o.ExpectedDeliveryDate != nil {
		return true
	}

	return false
}

// SetExpectedDeliveryDate gets a reference to the given string and assigns it to the ExpectedDeliveryDate field.
func (o *Letter) SetExpectedDeliveryDate(v string) {
	o.ExpectedDeliveryDate = &v
}

// GetDateCreated returns the DateCreated field value
func (o *Letter) GetDateCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *Letter) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *Letter) SetDateCreated(v time.Time) {
	o.DateCreated = v
}

// GetDateModified returns the DateModified field value
func (o *Letter) GetDateModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value
// and a boolean to check if the value has been set.
func (o *Letter) GetDateModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateModified, true
}

// SetDateModified sets field value
func (o *Letter) SetDateModified(v time.Time) {
	o.DateModified = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Letter) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Letter) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Letter) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetId returns the Id field value
func (o *Letter) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Letter) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Letter) SetId(v string) {
	o.Id = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *Letter) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *Letter) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *Letter) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateVersionId returns the TemplateVersionId field value if set, zero value otherwise.
func (o *Letter) GetTemplateVersionId() string {
	if o == nil || o.TemplateVersionId == nil {
		var ret string
		return ret
	}
	return *o.TemplateVersionId
}

// GetTemplateVersionIdOk returns a tuple with the TemplateVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetTemplateVersionIdOk() (*string, bool) {
	if o == nil || o.TemplateVersionId == nil {
		return nil, false
	}
	return o.TemplateVersionId, true
}

// HasTemplateVersionId returns a boolean if a field has been set.
func (o *Letter) HasTemplateVersionId() bool {
	if o != nil && o.TemplateVersionId != nil {
		return true
	}

	return false
}

// SetTemplateVersionId gets a reference to the given string and assigns it to the TemplateVersionId field.
func (o *Letter) SetTemplateVersionId(v string) {
	o.TemplateVersionId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Letter) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Letter) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Letter) SetUrl(v string) {
	o.Url = &v
}

// GetObject returns the Object field value
func (o *Letter) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Letter) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Letter) SetObject(v string) {
	o.Object = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Letter) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Letter) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Letter) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Letter) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Letter) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Letter) UnsetDescription() {
	o.Description.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Letter) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Letter) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Letter) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMergeVariables returns the MergeVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Letter) GetMergeVariables() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MergeVariables
}

// GetMergeVariablesOk returns a tuple with the MergeVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Letter) GetMergeVariablesOk() (map[string]interface{}, bool) {
	if o == nil || o.MergeVariables == nil {
		return nil, false
	}
	return o.MergeVariables, true
}

// HasMergeVariables returns a boolean if a field has been set.
func (o *Letter) HasMergeVariables() bool {
	if o != nil && o.MergeVariables != nil {
		return true
	}

	return false
}

// SetMergeVariables gets a reference to the given map[string]interface{} and assigns it to the MergeVariables field.
func (o *Letter) SetMergeVariables(v map[string]interface{}) {
	o.MergeVariables = v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *Letter) GetSendDate() time.Time {
	if o == nil || o.SendDate == nil {
		var ret time.Time
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetSendDateOk() (*time.Time, bool) {
	if o == nil || o.SendDate == nil {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *Letter) HasSendDate() bool {
	if o != nil && o.SendDate != nil {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given time.Time and assigns it to the SendDate field.
func (o *Letter) SetSendDate(v time.Time) {
	o.SendDate = &v
}

// GetExtraService returns the ExtraService field value if set, zero value otherwise.
func (o *Letter) GetExtraService() string {
	if o == nil || o.ExtraService == nil {
		var ret string
		return ret
	}
	return *o.ExtraService
}

// GetExtraServiceOk returns a tuple with the ExtraService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetExtraServiceOk() (*string, bool) {
	if o == nil || o.ExtraService == nil {
		return nil, false
	}
	return o.ExtraService, true
}

// HasExtraService returns a boolean if a field has been set.
func (o *Letter) HasExtraService() bool {
	if o != nil && o.ExtraService != nil {
		return true
	}

	return false
}

// SetExtraService gets a reference to the given string and assigns it to the ExtraService field.
func (o *Letter) SetExtraService(v string) {
	o.ExtraService = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Letter) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.TrackingNumber.Get()
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Letter) GetTrackingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingNumber.Get(), o.TrackingNumber.IsSet()
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *Letter) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber.IsSet() {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given NullableString and assigns it to the TrackingNumber field.
func (o *Letter) SetTrackingNumber(v string) {
	o.TrackingNumber.Set(&v)
}
// SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil
func (o *Letter) SetTrackingNumberNil() {
	o.TrackingNumber.Set(nil)
}

// UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
func (o *Letter) UnsetTrackingNumber() {
	o.TrackingNumber.Unset()
}

// GetTrackingEvents returns the TrackingEvents field value if set, zero value otherwise.
func (o *Letter) GetTrackingEvents() []TrackingEventNormal {
	if o == nil || o.TrackingEvents == nil {
		var ret []TrackingEventNormal
		return ret
	}
	return o.TrackingEvents
}

// GetTrackingEventsOk returns a tuple with the TrackingEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetTrackingEventsOk() ([]TrackingEventNormal, bool) {
	if o == nil || o.TrackingEvents == nil {
		return nil, false
	}
	return o.TrackingEvents, true
}

// HasTrackingEvents returns a boolean if a field has been set.
func (o *Letter) HasTrackingEvents() bool {
	if o != nil && o.TrackingEvents != nil {
		return true
	}

	return false
}

// SetTrackingEvents gets a reference to the given []TrackingEventNormal and assigns it to the TrackingEvents field.
func (o *Letter) SetTrackingEvents(v []TrackingEventNormal) {
	o.TrackingEvents = v
}

// GetReturnAddress returns the ReturnAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Letter) GetReturnAddress() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReturnAddress
}

// GetReturnAddressOk returns a tuple with the ReturnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Letter) GetReturnAddressOk() (*interface{}, bool) {
	if o == nil || o.ReturnAddress == nil {
		return nil, false
	}
	return &o.ReturnAddress, true
}

// HasReturnAddress returns a boolean if a field has been set.
func (o *Letter) HasReturnAddress() bool {
	if o != nil && o.ReturnAddress != nil {
		return true
	}

	return false
}

// SetReturnAddress gets a reference to the given interface{} and assigns it to the ReturnAddress field.
func (o *Letter) SetReturnAddress(v interface{}) {
	o.ReturnAddress = v
}

// GetMailType returns the MailType field value if set, zero value otherwise.
func (o *Letter) GetMailType() MailType {
	if o == nil || o.MailType == nil {
		var ret MailType
		return ret
	}
	return *o.MailType
}

// GetMailTypeOk returns a tuple with the MailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetMailTypeOk() (*MailType, bool) {
	if o == nil || o.MailType == nil {
		return nil, false
	}
	return o.MailType, true
}

// HasMailType returns a boolean if a field has been set.
func (o *Letter) HasMailType() bool {
	if o != nil && o.MailType != nil {
		return true
	}

	return false
}

// SetMailType gets a reference to the given MailType and assigns it to the MailType field.
func (o *Letter) SetMailType(v MailType) {
	o.MailType = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Letter) GetColor() bool {
	if o == nil || o.Color == nil {
		var ret bool
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetColorOk() (*bool, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Letter) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given bool and assigns it to the Color field.
func (o *Letter) SetColor(v bool) {
	o.Color = &v
}

// GetDoubleSided returns the DoubleSided field value if set, zero value otherwise.
func (o *Letter) GetDoubleSided() bool {
	if o == nil || o.DoubleSided == nil {
		var ret bool
		return ret
	}
	return *o.DoubleSided
}

// GetDoubleSidedOk returns a tuple with the DoubleSided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetDoubleSidedOk() (*bool, bool) {
	if o == nil || o.DoubleSided == nil {
		return nil, false
	}
	return o.DoubleSided, true
}

// HasDoubleSided returns a boolean if a field has been set.
func (o *Letter) HasDoubleSided() bool {
	if o != nil && o.DoubleSided != nil {
		return true
	}

	return false
}

// SetDoubleSided gets a reference to the given bool and assigns it to the DoubleSided field.
func (o *Letter) SetDoubleSided(v bool) {
	o.DoubleSided = &v
}

// GetAddressPlacement returns the AddressPlacement field value if set, zero value otherwise.
func (o *Letter) GetAddressPlacement() string {
	if o == nil || o.AddressPlacement == nil {
		var ret string
		return ret
	}
	return *o.AddressPlacement
}

// GetAddressPlacementOk returns a tuple with the AddressPlacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Letter) GetAddressPlacementOk() (*string, bool) {
	if o == nil || o.AddressPlacement == nil {
		return nil, false
	}
	return o.AddressPlacement, true
}

// HasAddressPlacement returns a boolean if a field has been set.
func (o *Letter) HasAddressPlacement() bool {
	if o != nil && o.AddressPlacement != nil {
		return true
	}

	return false
}

// SetAddressPlacement gets a reference to the given string and assigns it to the AddressPlacement field.
func (o *Letter) SetAddressPlacement(v string) {
	o.AddressPlacement = &v
}

// GetReturnEnvelope returns the ReturnEnvelope field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Letter) GetReturnEnvelope() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ReturnEnvelope
}

// GetReturnEnvelopeOk returns a tuple with the ReturnEnvelope field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Letter) GetReturnEnvelopeOk() (*interface{}, bool) {
	if o == nil || o.ReturnEnvelope == nil {
		return nil, false
	}
	return &o.ReturnEnvelope, true
}

// SetReturnEnvelope sets field value
func (o *Letter) SetReturnEnvelope(v interface{}) {
	o.ReturnEnvelope = v
}

// GetPerforatedPage returns the PerforatedPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Letter) GetPerforatedPage() int32 {
	if o == nil || o.PerforatedPage.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PerforatedPage.Get()
}

// GetPerforatedPageOk returns a tuple with the PerforatedPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Letter) GetPerforatedPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerforatedPage.Get(), o.PerforatedPage.IsSet()
}

// HasPerforatedPage returns a boolean if a field has been set.
func (o *Letter) HasPerforatedPage() bool {
	if o != nil && o.PerforatedPage.IsSet() {
		return true
	}

	return false
}

// SetPerforatedPage gets a reference to the given NullableInt32 and assigns it to the PerforatedPage field.
func (o *Letter) SetPerforatedPage(v int32) {
	o.PerforatedPage.Set(&v)
}
// SetPerforatedPageNil sets the value for PerforatedPage to be an explicit nil
func (o *Letter) SetPerforatedPageNil() {
	o.PerforatedPage.Set(nil)
}

// UnsetPerforatedPage ensures that no value is present for PerforatedPage, not even an explicit nil
func (o *Letter) UnsetPerforatedPage() {
	o.PerforatedPage.Unset()
}

// GetCustomEnvelope returns the CustomEnvelope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Letter) GetCustomEnvelope() LetterCustomEnvelope {
	if o == nil || o.CustomEnvelope.Get() == nil {
		var ret LetterCustomEnvelope
		return ret
	}
	return *o.CustomEnvelope.Get()
}

// GetCustomEnvelopeOk returns a tuple with the CustomEnvelope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Letter) GetCustomEnvelopeOk() (*LetterCustomEnvelope, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomEnvelope.Get(), o.CustomEnvelope.IsSet()
}

// HasCustomEnvelope returns a boolean if a field has been set.
func (o *Letter) HasCustomEnvelope() bool {
	if o != nil && o.CustomEnvelope.IsSet() {
		return true
	}

	return false
}

// SetCustomEnvelope gets a reference to the given NullableLetterCustomEnvelope and assigns it to the CustomEnvelope field.
func (o *Letter) SetCustomEnvelope(v LetterCustomEnvelope) {
	o.CustomEnvelope.Set(&v)
}
// SetCustomEnvelopeNil sets the value for CustomEnvelope to be an explicit nil
func (o *Letter) SetCustomEnvelopeNil() {
	o.CustomEnvelope.Set(nil)
}

// UnsetCustomEnvelope ensures that no value is present for CustomEnvelope, not even an explicit nil
func (o *Letter) UnsetCustomEnvelope() {
	o.CustomEnvelope.Unset()
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Letter) GetCampaignId() string {
	if o == nil || o.CampaignId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CampaignId.Get()
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Letter) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignId.Get(), o.CampaignId.IsSet()
}

// HasCampaignId returns a boolean if a field has been set.
func (o *Letter) HasCampaignId() bool {
	if o != nil && o.CampaignId.IsSet() {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given NullableString and assigns it to the CampaignId field.
func (o *Letter) SetCampaignId(v string) {
	o.CampaignId.Set(&v)
}
// SetCampaignIdNil sets the value for CampaignId to be an explicit nil
func (o *Letter) SetCampaignIdNil() {
	o.CampaignId.Set(nil)
}

// UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
func (o *Letter) UnsetCampaignId() {
	o.CampaignId.Unset()
}

// GetUseType returns the UseType field value
// If the value is explicit nil, the zero value for LtrUseType will be returned
func (o *Letter) GetUseType() LtrUseType {
	if o == nil || o.UseType.Get() == nil {
		var ret LtrUseType
		return ret
	}

	return *o.UseType.Get()
}

// GetUseTypeOk returns a tuple with the UseType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Letter) GetUseTypeOk() (*LtrUseType, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseType.Get(), o.UseType.IsSet()
}

// SetUseType sets field value
func (o *Letter) SetUseType(v LtrUseType) {
	o.UseType.Set(&v)
}

func (o Letter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["from"] = o.From
	}
	if o.Carrier != nil {
		toSerialize["carrier"] = o.Carrier
	}
	if o.Thumbnails != nil {
		toSerialize["thumbnails"] = o.Thumbnails
	}
	if o.ExpectedDeliveryDate != nil {
		toSerialize["expected_delivery_date"] = o.ExpectedDeliveryDate
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
		toSerialize["id"] = o.Id
	}
	if o.TemplateId != nil {
		toSerialize["template_id"] = o.TemplateId
	}
	if o.TemplateVersionId != nil {
		toSerialize["template_version_id"] = o.TemplateVersionId
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["object"] = o.Object
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
	if o.ExtraService != nil {
		toSerialize["extra_service"] = o.ExtraService
	}
	if o.TrackingNumber.IsSet() {
		toSerialize["tracking_number"] = o.TrackingNumber.Get()
	}
	if o.TrackingEvents != nil {
		toSerialize["tracking_events"] = o.TrackingEvents
	}
	if o.ReturnAddress != nil {
		toSerialize["return_address"] = o.ReturnAddress
	}
	if o.MailType != nil {
		toSerialize["mail_type"] = o.MailType
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.DoubleSided != nil {
		toSerialize["double_sided"] = o.DoubleSided
	}
	if o.AddressPlacement != nil {
		toSerialize["address_placement"] = o.AddressPlacement
	}
	if o.ReturnEnvelope != nil {
		toSerialize["return_envelope"] = o.ReturnEnvelope
	}
	if o.PerforatedPage.IsSet() {
		toSerialize["perforated_page"] = o.PerforatedPage.Get()
	}
	if o.CustomEnvelope.IsSet() {
		toSerialize["custom_envelope"] = o.CustomEnvelope.Get()
	}
	if o.CampaignId.IsSet() {
		toSerialize["campaign_id"] = o.CampaignId.Get()
	}
	if true {
		toSerialize["use_type"] = o.UseType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLetter struct {
	value *Letter
	isSet bool
}

func (v NullableLetter) Get() *Letter {
	return v.value
}

func (v *NullableLetter) Set(val *Letter) {
	v.value = val
	v.isSet = true
}

func (v NullableLetter) IsSet() bool {
	return v.isSet
}

func (v *NullableLetter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLetter(val *Letter) *NullableLetter {
	return &NullableLetter{value: val, isSet: true}
}

func (v NullableLetter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLetter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


