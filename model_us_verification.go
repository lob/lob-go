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
	
)

// UsVerification struct for UsVerification
type UsVerification struct {
	// Unique identifier prefixed with `us_ver_`.
	Id *string `json:"id,omitempty"`
	// The intended recipient, typically a person's or firm's name.
	Recipient NullableString `json:"recipient,omitempty"`
	// The primary delivery line (usually the street address) of the address. Combination of the following applicable `components`: * `primary_number` * `street_predirection` * `street_name` * `street_suffix` * `street_postdirection` * `secondary_designator` * `secondary_number` * `pmb_designator` * `pmb_number` 
	PrimaryLine *string `json:"primary_line,omitempty"`
	// The secondary delivery line of the address. This field is typically empty but may contain information if `primary_line` is too long. 
	SecondaryLine *string `json:"secondary_line,omitempty"`
	// Only present for addresses in Puerto Rico. An urbanization refers to an area, sector, or development within a city. See [USPS documentation](https://pe.usps.com/text/pub28/28api_008.htm#:~:text=I51.,-4%20Urbanizations&text=In%20Puerto%20Rico%2C%20identical%20street,placed%20before%20the%20urbanization%20name.) for clarification. 
	Urbanization *string `json:"urbanization,omitempty"`
	// Combination of the following applicable `components`: * City (`city`) * State (`state`) * ZIP code (`zip_code`) * ZIP+4 (`zip_code_plus_4`) 
	LastLine *string `json:"last_line,omitempty"`
	// Summarizes the deliverability of the `us_verification` object. For full details, see the `deliverability_analysis` field. Possible values are: * `deliverable` – The address is deliverable by the USPS. * `deliverable_unnecessary_unit` – The address is deliverable, but the secondary unit information is unnecessary. * `deliverable_incorrect_unit` – The address is deliverable to the building's default address but the secondary unit provided may not exist. There is a chance the mail will not reach the intended recipient. * `deliverable_missing_unit` – The address is deliverable to the building's default address but is missing secondary unit information. There is a chance the mail will not reach the intended recipient. * `undeliverable` – The address is not deliverable according to the USPS. 
	Deliverability *string `json:"deliverability,omitempty"`
	// This field indicates whether an address was found in a more comprehensive address dataset that includes sources from the USPS, open source mapping data, and our proprietary mail delivery data. This field can be interpreted as a representation of whether an address is a real location or not. Additionally a valid address may contradict the deliverability field since an address can be a real valid location but the USPS may not deliver to that address. 
	ValidAddress *bool `json:"valid_address,omitempty"`
	Components *UsComponents `json:"components,omitempty"`
	DeliverabilityAnalysis *DeliverabilityAnalysis `json:"deliverability_analysis,omitempty"`
	LobConfidenceScore *LobConfidenceScore `json:"lob_confidence_score,omitempty"`
	Object *string `json:"object,omitempty"`
	// ID that is returned in the response body for the verification 
	TransientId *string `json:"transient_id,omitempty"`
}

// NewUsVerification instantiates a new UsVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsVerification() *UsVerification {
	this := UsVerification{}
	var object string = "us_verification"
	this.Object = &object
	return &this
}

// NewUsVerificationWithDefaults instantiates a new UsVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsVerificationWithDefaults() *UsVerification {
	this := UsVerification{}
	var object string = "us_verification"
	this.Object = &object
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UsVerification) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UsVerification) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UsVerification) SetId(v string) {
	o.Id = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsVerification) GetRecipient() string {
	if o == nil || o.Recipient.Get() == nil {
		var ret string
		return ret
	}
	return *o.Recipient.Get()
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsVerification) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipient.Get(), o.Recipient.IsSet()
}

// HasRecipient returns a boolean if a field has been set.
func (o *UsVerification) HasRecipient() bool {
	if o != nil && o.Recipient.IsSet() {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given NullableString and assigns it to the Recipient field.
func (o *UsVerification) SetRecipient(v string) {
	o.Recipient.Set(&v)
}
// SetRecipientNil sets the value for Recipient to be an explicit nil
func (o *UsVerification) SetRecipientNil() {
	o.Recipient.Set(nil)
}

// UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
func (o *UsVerification) UnsetRecipient() {
	o.Recipient.Unset()
}

// GetPrimaryLine returns the PrimaryLine field value if set, zero value otherwise.
func (o *UsVerification) GetPrimaryLine() string {
	if o == nil || o.PrimaryLine == nil {
		var ret string
		return ret
	}
	return *o.PrimaryLine
}

// GetPrimaryLineOk returns a tuple with the PrimaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetPrimaryLineOk() (*string, bool) {
	if o == nil || o.PrimaryLine == nil {
		return nil, false
	}
	return o.PrimaryLine, true
}

// HasPrimaryLine returns a boolean if a field has been set.
func (o *UsVerification) HasPrimaryLine() bool {
	if o != nil && o.PrimaryLine != nil {
		return true
	}

	return false
}

// SetPrimaryLine gets a reference to the given string and assigns it to the PrimaryLine field.
func (o *UsVerification) SetPrimaryLine(v string) {
	o.PrimaryLine = &v
}

// GetSecondaryLine returns the SecondaryLine field value if set, zero value otherwise.
func (o *UsVerification) GetSecondaryLine() string {
	if o == nil || o.SecondaryLine == nil {
		var ret string
		return ret
	}
	return *o.SecondaryLine
}

// GetSecondaryLineOk returns a tuple with the SecondaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetSecondaryLineOk() (*string, bool) {
	if o == nil || o.SecondaryLine == nil {
		return nil, false
	}
	return o.SecondaryLine, true
}

// HasSecondaryLine returns a boolean if a field has been set.
func (o *UsVerification) HasSecondaryLine() bool {
	if o != nil && o.SecondaryLine != nil {
		return true
	}

	return false
}

// SetSecondaryLine gets a reference to the given string and assigns it to the SecondaryLine field.
func (o *UsVerification) SetSecondaryLine(v string) {
	o.SecondaryLine = &v
}

// GetUrbanization returns the Urbanization field value if set, zero value otherwise.
func (o *UsVerification) GetUrbanization() string {
	if o == nil || o.Urbanization == nil {
		var ret string
		return ret
	}
	return *o.Urbanization
}

// GetUrbanizationOk returns a tuple with the Urbanization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetUrbanizationOk() (*string, bool) {
	if o == nil || o.Urbanization == nil {
		return nil, false
	}
	return o.Urbanization, true
}

// HasUrbanization returns a boolean if a field has been set.
func (o *UsVerification) HasUrbanization() bool {
	if o != nil && o.Urbanization != nil {
		return true
	}

	return false
}

// SetUrbanization gets a reference to the given string and assigns it to the Urbanization field.
func (o *UsVerification) SetUrbanization(v string) {
	o.Urbanization = &v
}

// GetLastLine returns the LastLine field value if set, zero value otherwise.
func (o *UsVerification) GetLastLine() string {
	if o == nil || o.LastLine == nil {
		var ret string
		return ret
	}
	return *o.LastLine
}

// GetLastLineOk returns a tuple with the LastLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetLastLineOk() (*string, bool) {
	if o == nil || o.LastLine == nil {
		return nil, false
	}
	return o.LastLine, true
}

// HasLastLine returns a boolean if a field has been set.
func (o *UsVerification) HasLastLine() bool {
	if o != nil && o.LastLine != nil {
		return true
	}

	return false
}

// SetLastLine gets a reference to the given string and assigns it to the LastLine field.
func (o *UsVerification) SetLastLine(v string) {
	o.LastLine = &v
}

// GetDeliverability returns the Deliverability field value if set, zero value otherwise.
func (o *UsVerification) GetDeliverability() string {
	if o == nil || o.Deliverability == nil {
		var ret string
		return ret
	}
	return *o.Deliverability
}

// GetDeliverabilityOk returns a tuple with the Deliverability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetDeliverabilityOk() (*string, bool) {
	if o == nil || o.Deliverability == nil {
		return nil, false
	}
	return o.Deliverability, true
}

// HasDeliverability returns a boolean if a field has been set.
func (o *UsVerification) HasDeliverability() bool {
	if o != nil && o.Deliverability != nil {
		return true
	}

	return false
}

// SetDeliverability gets a reference to the given string and assigns it to the Deliverability field.
func (o *UsVerification) SetDeliverability(v string) {
	o.Deliverability = &v
}

// GetValidAddress returns the ValidAddress field value if set, zero value otherwise.
func (o *UsVerification) GetValidAddress() bool {
	if o == nil || o.ValidAddress == nil {
		var ret bool
		return ret
	}
	return *o.ValidAddress
}

// GetValidAddressOk returns a tuple with the ValidAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetValidAddressOk() (*bool, bool) {
	if o == nil || o.ValidAddress == nil {
		return nil, false
	}
	return o.ValidAddress, true
}

// HasValidAddress returns a boolean if a field has been set.
func (o *UsVerification) HasValidAddress() bool {
	if o != nil && o.ValidAddress != nil {
		return true
	}

	return false
}

// SetValidAddress gets a reference to the given bool and assigns it to the ValidAddress field.
func (o *UsVerification) SetValidAddress(v bool) {
	o.ValidAddress = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *UsVerification) GetComponents() UsComponents {
	if o == nil || o.Components == nil {
		var ret UsComponents
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetComponentsOk() (*UsComponents, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *UsVerification) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given UsComponents and assigns it to the Components field.
func (o *UsVerification) SetComponents(v UsComponents) {
	o.Components = &v
}

// GetDeliverabilityAnalysis returns the DeliverabilityAnalysis field value if set, zero value otherwise.
func (o *UsVerification) GetDeliverabilityAnalysis() DeliverabilityAnalysis {
	if o == nil || o.DeliverabilityAnalysis == nil {
		var ret DeliverabilityAnalysis
		return ret
	}
	return *o.DeliverabilityAnalysis
}

// GetDeliverabilityAnalysisOk returns a tuple with the DeliverabilityAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetDeliverabilityAnalysisOk() (*DeliverabilityAnalysis, bool) {
	if o == nil || o.DeliverabilityAnalysis == nil {
		return nil, false
	}
	return o.DeliverabilityAnalysis, true
}

// HasDeliverabilityAnalysis returns a boolean if a field has been set.
func (o *UsVerification) HasDeliverabilityAnalysis() bool {
	if o != nil && o.DeliverabilityAnalysis != nil {
		return true
	}

	return false
}

// SetDeliverabilityAnalysis gets a reference to the given DeliverabilityAnalysis and assigns it to the DeliverabilityAnalysis field.
func (o *UsVerification) SetDeliverabilityAnalysis(v DeliverabilityAnalysis) {
	o.DeliverabilityAnalysis = &v
}

// GetLobConfidenceScore returns the LobConfidenceScore field value if set, zero value otherwise.
func (o *UsVerification) GetLobConfidenceScore() LobConfidenceScore {
	if o == nil || o.LobConfidenceScore == nil {
		var ret LobConfidenceScore
		return ret
	}
	return *o.LobConfidenceScore
}

// GetLobConfidenceScoreOk returns a tuple with the LobConfidenceScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetLobConfidenceScoreOk() (*LobConfidenceScore, bool) {
	if o == nil || o.LobConfidenceScore == nil {
		return nil, false
	}
	return o.LobConfidenceScore, true
}

// HasLobConfidenceScore returns a boolean if a field has been set.
func (o *UsVerification) HasLobConfidenceScore() bool {
	if o != nil && o.LobConfidenceScore != nil {
		return true
	}

	return false
}

// SetLobConfidenceScore gets a reference to the given LobConfidenceScore and assigns it to the LobConfidenceScore field.
func (o *UsVerification) SetLobConfidenceScore(v LobConfidenceScore) {
	o.LobConfidenceScore = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *UsVerification) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *UsVerification) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *UsVerification) SetObject(v string) {
	o.Object = &v
}

// GetTransientId returns the TransientId field value if set, zero value otherwise.
func (o *UsVerification) GetTransientId() string {
	if o == nil || o.TransientId == nil {
		var ret string
		return ret
	}
	return *o.TransientId
}

// GetTransientIdOk returns a tuple with the TransientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsVerification) GetTransientIdOk() (*string, bool) {
	if o == nil || o.TransientId == nil {
		return nil, false
	}
	return o.TransientId, true
}

// HasTransientId returns a boolean if a field has been set.
func (o *UsVerification) HasTransientId() bool {
	if o != nil && o.TransientId != nil {
		return true
	}

	return false
}

// SetTransientId gets a reference to the given string and assigns it to the TransientId field.
func (o *UsVerification) SetTransientId(v string) {
	o.TransientId = &v
}

func (o UsVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Recipient.IsSet() {
		toSerialize["recipient"] = o.Recipient.Get()
	}
	if o.PrimaryLine != nil {
		toSerialize["primary_line"] = o.PrimaryLine
	}
	if o.SecondaryLine != nil {
		toSerialize["secondary_line"] = o.SecondaryLine
	}
	if o.Urbanization != nil {
		toSerialize["urbanization"] = o.Urbanization
	}
	if o.LastLine != nil {
		toSerialize["last_line"] = o.LastLine
	}
	if o.Deliverability != nil {
		toSerialize["deliverability"] = o.Deliverability
	}
	if o.ValidAddress != nil {
		toSerialize["valid_address"] = o.ValidAddress
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.DeliverabilityAnalysis != nil {
		toSerialize["deliverability_analysis"] = o.DeliverabilityAnalysis
	}
	if o.LobConfidenceScore != nil {
		toSerialize["lob_confidence_score"] = o.LobConfidenceScore
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.TransientId != nil {
		toSerialize["transient_id"] = o.TransientId
	}
	return json.Marshal(toSerialize)
}

type NullableUsVerification struct {
	value *UsVerification
	isSet bool
}

func (v NullableUsVerification) Get() *UsVerification {
	return v.value
}

func (v *NullableUsVerification) Set(val *UsVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableUsVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableUsVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsVerification(val *UsVerification) *NullableUsVerification {
	return &NullableUsVerification{value: val, isSet: true}
}

func (v NullableUsVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


