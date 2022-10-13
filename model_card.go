/*
Lob

The Lob API is organized around REST. Our API is designed to have predictable, resource-oriented URLs and uses HTTP response codes to indicate any API errors. <p> Looking for our [previous documentation](https://lob.github.io/legacy-docs/)? 

API version: 1.3.0
Contact: lob-openapi@lob.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lob.go

import (
	"encoding/json"
	"time"
)

// Card struct for Card
type Card struct {
	// Unique identifier prefixed with `card_`.
	Id string `json:"id"`
	// The signed link for the card.
	Url string `json:"url"`
	// True if the cards should be auto-reordered.
	AutoReorder bool `json:"auto_reorder"`
	// The number of cards to be reordered. Only present when auto_reorder is True.
	ReorderQuantity NullableInt32 `json:"reorder_quantity,omitempty"`
	// The raw URL of the card.
	RawUrl *string `json:"raw_url,omitempty"`
	// The original URL of the front template.
	FrontOriginalUrl *string `json:"front_original_url,omitempty"`
	// The original URL of the back template.
	BackOriginalUrl *string `json:"back_original_url,omitempty"`
	Thumbnails []Thumbnail `json:"thumbnails"`
	// The available quantity of cards.
	AvailableQuantity int32 `json:"available_quantity"`
	// The pending quantity of cards.
	PendingQuantity int32 `json:"pending_quantity"`
	Status *string `json:"status,omitempty"`
	// The orientation of the card.
	Orientation *string `json:"orientation,omitempty"`
	// The threshold amount of the card
	ThresholdAmount *int32 `json:"threshold_amount,omitempty"`
	// A timestamp in ISO 8601 format of the date the resource was created.
	DateCreated time.Time `json:"date_created"`
	// A timestamp in ISO 8601 format of the date the resource was last modified.
	DateModified time.Time `json:"date_modified"`
	// Only returned if the resource has been successfully deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// object
	Object string `json:"object"`
	// Description of the card.
	Description NullableString `json:"description,omitempty"`
	// The size of the card
	Size string `json:"size"`
}

// NewCard instantiates a new Card object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCard(id string, url string, autoReorder bool, thumbnails []Thumbnail, availableQuantity int32, pendingQuantity int32, dateCreated time.Time, dateModified time.Time, object string, size string) *Card {
	this := Card{}
	this.Id = id
	this.Url = url
	this.AutoReorder = autoReorder
	this.Thumbnails = thumbnails
	this.AvailableQuantity = availableQuantity
	this.PendingQuantity = pendingQuantity
	var orientation string = "horizontal"
	this.Orientation = &orientation
	var thresholdAmount int32 = 0
	this.ThresholdAmount = &thresholdAmount
	this.DateCreated = dateCreated
	this.DateModified = dateModified
	this.Object = object
	this.Size = size
	return &this
}

// NewCardWithDefaults instantiates a new Card object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWithDefaults() *Card {
	this := Card{}
	var autoReorder bool = false
	this.AutoReorder = autoReorder
	var availableQuantity int32 = 0
	this.AvailableQuantity = availableQuantity
	var pendingQuantity int32 = 0
	this.PendingQuantity = pendingQuantity
	var orientation string = "horizontal"
	this.Orientation = &orientation
	var thresholdAmount int32 = 0
	this.ThresholdAmount = &thresholdAmount
	var object string = "card"
	this.Object = object
	var size string = "2.125x3.375"
	this.Size = size
	return &this
}

// GetId returns the Id field value
func (o *Card) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Card) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Card) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Card) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Card) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Card) SetUrl(v string) {
	o.Url = v
}

// GetAutoReorder returns the AutoReorder field value
func (o *Card) GetAutoReorder() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoReorder
}

// GetAutoReorderOk returns a tuple with the AutoReorder field value
// and a boolean to check if the value has been set.
func (o *Card) GetAutoReorderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoReorder, true
}

// SetAutoReorder sets field value
func (o *Card) SetAutoReorder(v bool) {
	o.AutoReorder = v
}

// GetReorderQuantity returns the ReorderQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Card) GetReorderQuantity() int32 {
	if o == nil || o.ReorderQuantity.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ReorderQuantity.Get()
}

// GetReorderQuantityOk returns a tuple with the ReorderQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Card) GetReorderQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReorderQuantity.Get(), o.ReorderQuantity.IsSet()
}

// HasReorderQuantity returns a boolean if a field has been set.
func (o *Card) HasReorderQuantity() bool {
	if o != nil && o.ReorderQuantity.IsSet() {
		return true
	}

	return false
}

// SetReorderQuantity gets a reference to the given NullableInt32 and assigns it to the ReorderQuantity field.
func (o *Card) SetReorderQuantity(v int32) {
	o.ReorderQuantity.Set(&v)
}
// SetReorderQuantityNil sets the value for ReorderQuantity to be an explicit nil
func (o *Card) SetReorderQuantityNil() {
	o.ReorderQuantity.Set(nil)
}

// UnsetReorderQuantity ensures that no value is present for ReorderQuantity, not even an explicit nil
func (o *Card) UnsetReorderQuantity() {
	o.ReorderQuantity.Unset()
}

// GetRawUrl returns the RawUrl field value if set, zero value otherwise.
func (o *Card) GetRawUrl() string {
	if o == nil || o.RawUrl == nil {
		var ret string
		return ret
	}
	return *o.RawUrl
}

// GetRawUrlOk returns a tuple with the RawUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetRawUrlOk() (*string, bool) {
	if o == nil || o.RawUrl == nil {
		return nil, false
	}
	return o.RawUrl, true
}

// HasRawUrl returns a boolean if a field has been set.
func (o *Card) HasRawUrl() bool {
	if o != nil && o.RawUrl != nil {
		return true
	}

	return false
}

// SetRawUrl gets a reference to the given string and assigns it to the RawUrl field.
func (o *Card) SetRawUrl(v string) {
	o.RawUrl = &v
}

// GetFrontOriginalUrl returns the FrontOriginalUrl field value if set, zero value otherwise.
func (o *Card) GetFrontOriginalUrl() string {
	if o == nil || o.FrontOriginalUrl == nil {
		var ret string
		return ret
	}
	return *o.FrontOriginalUrl
}

// GetFrontOriginalUrlOk returns a tuple with the FrontOriginalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetFrontOriginalUrlOk() (*string, bool) {
	if o == nil || o.FrontOriginalUrl == nil {
		return nil, false
	}
	return o.FrontOriginalUrl, true
}

// HasFrontOriginalUrl returns a boolean if a field has been set.
func (o *Card) HasFrontOriginalUrl() bool {
	if o != nil && o.FrontOriginalUrl != nil {
		return true
	}

	return false
}

// SetFrontOriginalUrl gets a reference to the given string and assigns it to the FrontOriginalUrl field.
func (o *Card) SetFrontOriginalUrl(v string) {
	o.FrontOriginalUrl = &v
}

// GetBackOriginalUrl returns the BackOriginalUrl field value if set, zero value otherwise.
func (o *Card) GetBackOriginalUrl() string {
	if o == nil || o.BackOriginalUrl == nil {
		var ret string
		return ret
	}
	return *o.BackOriginalUrl
}

// GetBackOriginalUrlOk returns a tuple with the BackOriginalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBackOriginalUrlOk() (*string, bool) {
	if o == nil || o.BackOriginalUrl == nil {
		return nil, false
	}
	return o.BackOriginalUrl, true
}

// HasBackOriginalUrl returns a boolean if a field has been set.
func (o *Card) HasBackOriginalUrl() bool {
	if o != nil && o.BackOriginalUrl != nil {
		return true
	}

	return false
}

// SetBackOriginalUrl gets a reference to the given string and assigns it to the BackOriginalUrl field.
func (o *Card) SetBackOriginalUrl(v string) {
	o.BackOriginalUrl = &v
}

// GetThumbnails returns the Thumbnails field value
func (o *Card) GetThumbnails() []Thumbnail {
	if o == nil {
		var ret []Thumbnail
		return ret
	}

	return o.Thumbnails
}

// GetThumbnailsOk returns a tuple with the Thumbnails field value
// and a boolean to check if the value has been set.
func (o *Card) GetThumbnailsOk() ([]Thumbnail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumbnails, true
}

// SetThumbnails sets field value
func (o *Card) SetThumbnails(v []Thumbnail) {
	o.Thumbnails = v
}

// GetAvailableQuantity returns the AvailableQuantity field value
func (o *Card) GetAvailableQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvailableQuantity
}

// GetAvailableQuantityOk returns a tuple with the AvailableQuantity field value
// and a boolean to check if the value has been set.
func (o *Card) GetAvailableQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableQuantity, true
}

// SetAvailableQuantity sets field value
func (o *Card) SetAvailableQuantity(v int32) {
	o.AvailableQuantity = v
}

// GetPendingQuantity returns the PendingQuantity field value
func (o *Card) GetPendingQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PendingQuantity
}

// GetPendingQuantityOk returns a tuple with the PendingQuantity field value
// and a boolean to check if the value has been set.
func (o *Card) GetPendingQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingQuantity, true
}

// SetPendingQuantity sets field value
func (o *Card) SetPendingQuantity(v int32) {
	o.PendingQuantity = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Card) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Card) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Card) SetStatus(v string) {
	o.Status = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *Card) GetOrientation() string {
	if o == nil || o.Orientation == nil {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetOrientationOk() (*string, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *Card) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *Card) SetOrientation(v string) {
	o.Orientation = &v
}

// GetThresholdAmount returns the ThresholdAmount field value if set, zero value otherwise.
func (o *Card) GetThresholdAmount() int32 {
	if o == nil || o.ThresholdAmount == nil {
		var ret int32
		return ret
	}
	return *o.ThresholdAmount
}

// GetThresholdAmountOk returns a tuple with the ThresholdAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetThresholdAmountOk() (*int32, bool) {
	if o == nil || o.ThresholdAmount == nil {
		return nil, false
	}
	return o.ThresholdAmount, true
}

// HasThresholdAmount returns a boolean if a field has been set.
func (o *Card) HasThresholdAmount() bool {
	if o != nil && o.ThresholdAmount != nil {
		return true
	}

	return false
}

// SetThresholdAmount gets a reference to the given int32 and assigns it to the ThresholdAmount field.
func (o *Card) SetThresholdAmount(v int32) {
	o.ThresholdAmount = &v
}

// GetDateCreated returns the DateCreated field value
func (o *Card) GetDateCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *Card) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *Card) SetDateCreated(v time.Time) {
	o.DateCreated = v
}

// GetDateModified returns the DateModified field value
func (o *Card) GetDateModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value
// and a boolean to check if the value has been set.
func (o *Card) GetDateModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateModified, true
}

// SetDateModified sets field value
func (o *Card) SetDateModified(v time.Time) {
	o.DateModified = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Card) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Card) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Card) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetObject returns the Object field value
func (o *Card) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Card) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Card) SetObject(v string) {
	o.Object = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Card) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Card) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Card) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Card) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Card) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Card) UnsetDescription() {
	o.Description.Unset()
}

// GetSize returns the Size field value
func (o *Card) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *Card) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *Card) SetSize(v string) {
	o.Size = v
}

func (o Card) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["auto_reorder"] = o.AutoReorder
	}
	if o.ReorderQuantity.IsSet() {
		toSerialize["reorder_quantity"] = o.ReorderQuantity.Get()
	}
	if o.RawUrl != nil {
		toSerialize["raw_url"] = o.RawUrl
	}
	if o.FrontOriginalUrl != nil {
		toSerialize["front_original_url"] = o.FrontOriginalUrl
	}
	if o.BackOriginalUrl != nil {
		toSerialize["back_original_url"] = o.BackOriginalUrl
	}
	if true {
		toSerialize["thumbnails"] = o.Thumbnails
	}
	if true {
		toSerialize["available_quantity"] = o.AvailableQuantity
	}
	if true {
		toSerialize["pending_quantity"] = o.PendingQuantity
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Orientation != nil {
		toSerialize["orientation"] = o.Orientation
	}
	if o.ThresholdAmount != nil {
		toSerialize["threshold_amount"] = o.ThresholdAmount
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
		toSerialize["object"] = o.Object
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableCard struct {
	value *Card
	isSet bool
}

func (v NullableCard) Get() *Card {
	return v.value
}

func (v *NullableCard) Set(val *Card) {
	v.value = val
	v.isSet = true
}

func (v NullableCard) IsSet() bool {
	return v.isSet
}

func (v *NullableCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCard(val *Card) *NullableCard {
	return &NullableCard{value: val, isSet: true}
}

func (v NullableCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


