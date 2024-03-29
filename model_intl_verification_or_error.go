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

// IntlVerificationOrError A model used to represent an entry in a result list where the entry can either be a intl_verification or an Error. The SDK will perform necessary casting into the correct corresponding type. 
type IntlVerificationOrError struct {
	// Unique identifier prefixed with `intl_ver_`.
	Id *string `json:"id,omitempty"`
	// The intended recipient, typically a person's or firm's name.
	Recipient NullableString `json:"recipient,omitempty"`
	PrimaryLine *string `json:"primary_line,omitempty"`
	// The secondary delivery line of the address. This field is typically empty but may contain information if `primary_line` is too long. 
	SecondaryLine *string `json:"secondary_line,omitempty"`
	LastLine *string `json:"last_line,omitempty"`
	Country *string `json:"country,omitempty"`
	Coverage *string `json:"coverage,omitempty"`
	Deliverability *string `json:"deliverability,omitempty"`
	Status *string `json:"status,omitempty"`
	Components *IntlComponents `json:"components,omitempty"`
	Object *string `json:"object,omitempty"`
	Error *BulkError `json:"error,omitempty"`
}

// NewIntlVerificationOrError instantiates a new IntlVerificationOrError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntlVerificationOrError() *IntlVerificationOrError {
	this := IntlVerificationOrError{}
	var object string = "intl_verification"
	this.Object = &object
	return &this
}

// NewIntlVerificationOrErrorWithDefaults instantiates a new IntlVerificationOrError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntlVerificationOrErrorWithDefaults() *IntlVerificationOrError {
	this := IntlVerificationOrError{}
	var object string = "intl_verification"
	this.Object = &object
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntlVerificationOrError) SetId(v string) {
	o.Id = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntlVerificationOrError) GetRecipient() string {
	if o == nil || o.Recipient.Get() == nil {
		var ret string
		return ret
	}
	return *o.Recipient.Get()
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntlVerificationOrError) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipient.Get(), o.Recipient.IsSet()
}

// HasRecipient returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasRecipient() bool {
	if o != nil && o.Recipient.IsSet() {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given NullableString and assigns it to the Recipient field.
func (o *IntlVerificationOrError) SetRecipient(v string) {
	o.Recipient.Set(&v)
}
// SetRecipientNil sets the value for Recipient to be an explicit nil
func (o *IntlVerificationOrError) SetRecipientNil() {
	o.Recipient.Set(nil)
}

// UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
func (o *IntlVerificationOrError) UnsetRecipient() {
	o.Recipient.Unset()
}

// GetPrimaryLine returns the PrimaryLine field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetPrimaryLine() string {
	if o == nil || o.PrimaryLine == nil {
		var ret string
		return ret
	}
	return *o.PrimaryLine
}

// GetPrimaryLineOk returns a tuple with the PrimaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetPrimaryLineOk() (*string, bool) {
	if o == nil || o.PrimaryLine == nil {
		return nil, false
	}
	return o.PrimaryLine, true
}

// HasPrimaryLine returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasPrimaryLine() bool {
	if o != nil && o.PrimaryLine != nil {
		return true
	}

	return false
}

// SetPrimaryLine gets a reference to the given string and assigns it to the PrimaryLine field.
func (o *IntlVerificationOrError) SetPrimaryLine(v string) {
	o.PrimaryLine = &v
}

// GetSecondaryLine returns the SecondaryLine field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetSecondaryLine() string {
	if o == nil || o.SecondaryLine == nil {
		var ret string
		return ret
	}
	return *o.SecondaryLine
}

// GetSecondaryLineOk returns a tuple with the SecondaryLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetSecondaryLineOk() (*string, bool) {
	if o == nil || o.SecondaryLine == nil {
		return nil, false
	}
	return o.SecondaryLine, true
}

// HasSecondaryLine returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasSecondaryLine() bool {
	if o != nil && o.SecondaryLine != nil {
		return true
	}

	return false
}

// SetSecondaryLine gets a reference to the given string and assigns it to the SecondaryLine field.
func (o *IntlVerificationOrError) SetSecondaryLine(v string) {
	o.SecondaryLine = &v
}

// GetLastLine returns the LastLine field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetLastLine() string {
	if o == nil || o.LastLine == nil {
		var ret string
		return ret
	}
	return *o.LastLine
}

// GetLastLineOk returns a tuple with the LastLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetLastLineOk() (*string, bool) {
	if o == nil || o.LastLine == nil {
		return nil, false
	}
	return o.LastLine, true
}

// HasLastLine returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasLastLine() bool {
	if o != nil && o.LastLine != nil {
		return true
	}

	return false
}

// SetLastLine gets a reference to the given string and assigns it to the LastLine field.
func (o *IntlVerificationOrError) SetLastLine(v string) {
	o.LastLine = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *IntlVerificationOrError) SetCountry(v string) {
	o.Country = &v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetCoverage() string {
	if o == nil || o.Coverage == nil {
		var ret string
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetCoverageOk() (*string, bool) {
	if o == nil || o.Coverage == nil {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasCoverage() bool {
	if o != nil && o.Coverage != nil {
		return true
	}

	return false
}

// SetCoverage gets a reference to the given string and assigns it to the Coverage field.
func (o *IntlVerificationOrError) SetCoverage(v string) {
	o.Coverage = &v
}

// GetDeliverability returns the Deliverability field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetDeliverability() string {
	if o == nil || o.Deliverability == nil {
		var ret string
		return ret
	}
	return *o.Deliverability
}

// GetDeliverabilityOk returns a tuple with the Deliverability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetDeliverabilityOk() (*string, bool) {
	if o == nil || o.Deliverability == nil {
		return nil, false
	}
	return o.Deliverability, true
}

// HasDeliverability returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasDeliverability() bool {
	if o != nil && o.Deliverability != nil {
		return true
	}

	return false
}

// SetDeliverability gets a reference to the given string and assigns it to the Deliverability field.
func (o *IntlVerificationOrError) SetDeliverability(v string) {
	o.Deliverability = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IntlVerificationOrError) SetStatus(v string) {
	o.Status = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetComponents() IntlComponents {
	if o == nil || o.Components == nil {
		var ret IntlComponents
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetComponentsOk() (*IntlComponents, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given IntlComponents and assigns it to the Components field.
func (o *IntlVerificationOrError) SetComponents(v IntlComponents) {
	o.Components = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *IntlVerificationOrError) SetObject(v string) {
	o.Object = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *IntlVerificationOrError) GetError() BulkError {
	if o == nil || o.Error == nil {
		var ret BulkError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntlVerificationOrError) GetErrorOk() (*BulkError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *IntlVerificationOrError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given BulkError and assigns it to the Error field.
func (o *IntlVerificationOrError) SetError(v BulkError) {
	o.Error = &v
}

func (o IntlVerificationOrError) MarshalJSON() ([]byte, error) {
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
	if o.LastLine != nil {
		toSerialize["last_line"] = o.LastLine
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
	}
	if o.Deliverability != nil {
		toSerialize["deliverability"] = o.Deliverability
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableIntlVerificationOrError struct {
	value *IntlVerificationOrError
	isSet bool
}

func (v NullableIntlVerificationOrError) Get() *IntlVerificationOrError {
	return v.value
}

func (v *NullableIntlVerificationOrError) Set(val *IntlVerificationOrError) {
	v.value = val
	v.isSet = true
}

func (v NullableIntlVerificationOrError) IsSet() bool {
	return v.isSet
}

func (v *NullableIntlVerificationOrError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntlVerificationOrError(val *IntlVerificationOrError) *NullableIntlVerificationOrError {
	return &NullableIntlVerificationOrError{value: val, isSet: true}
}

func (v NullableIntlVerificationOrError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntlVerificationOrError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


