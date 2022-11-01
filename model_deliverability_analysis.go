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

// DeliverabilityAnalysis A nested object containing a breakdown of the deliverability of an address.
type DeliverabilityAnalysis struct {
	// Result of Delivery Point Validation (DPV), which determines whether or not the address is deliverable by the USPS. Possible values are: * `Y` –– The address is deliverable by the USPS. * `S` –– The address is deliverable by removing the provided secondary unit designator. This information may be incorrect or unnecessary. * `D` –– The address is deliverable to the building's default address but is missing a secondary unit designator and/or number.   There is a chance the mail will not reach the intended recipient. * `N` –– The address is not deliverable according to the USPS, but parts of the address are valid (such as the street and ZIP code). * `''` –– This address is not deliverable. No matching street could be found within the city or ZIP code. 
	DpvConfirmation string `json:"dpv_confirmation"`
	// indicates whether or not the address is [CMRA-authorized](https://en.wikipedia.org/wiki/Commercial_mail_receiving_agency). Possible values are: * `Y` –– Address is CMRA-authorized. * `N` –– Address is not CMRA-authorized. * `''` –– A DPV match is not made (`deliverability_analysis[dpv_confirmation]` is `N` or an empty string). 
	DpvCmra string `json:"dpv_cmra"`
	// indicates that an address was once deliverable, but has become vacant and is no longer receiving deliveries. Possible values are: * `Y` –– Address is vacant. * `N` –– Address is not vacant. * `''` –– A DPV match is not made (`deliverability_analysis[dpv_confirmation]` is `N` or an empty string). 
	DpvVacant string `json:"dpv_vacant"`
	// Corresponds to the USPS field `dpv_no_stat`. Indicates that an address has been vacated in the recent past, and is no longer receiving deliveries. If it's been unoccupied for 90+ days, or temporarily vacant, this will be flagged. Possible values are: * `Y` –– Address is active. * `N` –– Address is not active. * `''` –– A DPV match is not made (`deliverability_analysis[dpv_confirmation]` is `N` or an empty string). 
	DpvActive string `json:"dpv_active"`
	// An array of 2-character strings that gives more insight into how `deliverability_analysis[dpv_confirmation]` was determined. Will always include at least 1 string, and can include up to 3. For details, see [US Verification Details](#tag/US-Verification-Types). 
	DpvFootnotes []DpvFootnote `json:"dpv_footnotes"`
	// indicates whether or not an address has been flagged in the [Early Warning System](https://docs.informatica.com/data-engineering/data-engineering-quality/10-4-0/address-validator-port-reference/postal-carrier-certification-data-ports/early-warning-system-return-code.html), meaning the address is under development and not yet ready to receive mail. However, it should become available in a few months. 
	EwsMatch bool `json:"ews_match"`
	// indicates whether this address has been converted by [LACS<sup>Link</sup>](https://postalpro.usps.com/address-quality/lacslink). LACS<sup>Link</sup> corrects outdated addresses into their modern counterparts. Possible values are: * `Y` –– New address produced with a matching record in LACS<sup>Link</sup>. * `N` –– New address could not be produced with a matching record in LACS<sup>Link</sup>. * `''` –– A DPV match is not made (`deliverability_analysis[dpv_confirmation]` is `N` or an empty string). 
	LacsIndicator string `json:"lacs_indicator"`
	// A code indicating how `deliverability_analysis[lacs_indicator]` was determined. Possible values are: * `A` — A new address was produced because a match was found in LACS<sup>Link</sup>. * `92` — A LACS<sup>Link</sup> record was matched after dropping secondary information. * `14` — A match was found in LACS<sup>Link</sup>, but could not be converted to a deliverable address. * `00` — A match was not found in LACS<sup>Link</sup>, and no new address was produced. * `''` — LACS<sup>Link</sup> was not attempted. 
	LacsReturnCode string `json:"lacs_return_code"`
	// A return code that indicates whether the address was matched and corrected by [Suite<sup>Link</sup>](https://postalpro.usps.com/address-quality-solutions/suitelink). Suite<sup>Link</sup> attempts to provide secondary information to business addresses. Possible values are: * `A` –– A Suite<sup>Link</sup> match was found and secondary information was added. * `00` –– A Suite<sup>Link</sup> match could not be found and no secondary information was added. * `''` –– Suite<sup>Link</sup> lookup was not attempted. 
	SuiteReturnCode string `json:"suite_return_code"`
}

// NewDeliverabilityAnalysis instantiates a new DeliverabilityAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverabilityAnalysis(dpvConfirmation string, dpvCmra string, dpvVacant string, dpvActive string, dpvFootnotes []DpvFootnote, ewsMatch bool, lacsIndicator string, lacsReturnCode string, suiteReturnCode string) *DeliverabilityAnalysis {
	this := DeliverabilityAnalysis{}
	this.DpvConfirmation = dpvConfirmation
	this.DpvCmra = dpvCmra
	this.DpvVacant = dpvVacant
	this.DpvActive = dpvActive
	this.DpvFootnotes = dpvFootnotes
	this.EwsMatch = ewsMatch
	this.LacsIndicator = lacsIndicator
	this.LacsReturnCode = lacsReturnCode
	this.SuiteReturnCode = suiteReturnCode
	return &this
}

// NewDeliverabilityAnalysisWithDefaults instantiates a new DeliverabilityAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverabilityAnalysisWithDefaults() *DeliverabilityAnalysis {
	this := DeliverabilityAnalysis{}
	return &this
}

// GetDpvConfirmation returns the DpvConfirmation field value
func (o *DeliverabilityAnalysis) GetDpvConfirmation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DpvConfirmation
}

// GetDpvConfirmationOk returns a tuple with the DpvConfirmation field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityAnalysis) GetDpvConfirmationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpvConfirmation, true
}

// SetDpvConfirmation sets field value
func (o *DeliverabilityAnalysis) SetDpvConfirmation(v string) {
	o.DpvConfirmation = v
}

// GetDpvCmra returns the DpvCmra field value
func (o *DeliverabilityAnalysis) GetDpvCmra() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DpvCmra
}

// GetDpvCmraOk returns a tuple with the DpvCmra field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityAnalysis) GetDpvCmraOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpvCmra, true
}

// SetDpvCmra sets field value
func (o *DeliverabilityAnalysis) SetDpvCmra(v string) {
	o.DpvCmra = v
}

// GetDpvVacant returns the DpvVacant field value
func (o *DeliverabilityAnalysis) GetDpvVacant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DpvVacant
}

// GetDpvVacantOk returns a tuple with the DpvVacant field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityAnalysis) GetDpvVacantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpvVacant, true
}

// SetDpvVacant sets field value
func (o *DeliverabilityAnalysis) SetDpvVacant(v string) {
	o.DpvVacant = v
}

// GetDpvActive returns the DpvActive field value
func (o *DeliverabilityAnalysis) GetDpvActive() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DpvActive
}

// GetDpvActiveOk returns a tuple with the DpvActive field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityAnalysis) GetDpvActiveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpvActive, true
}

// SetDpvActive sets field value
func (o *DeliverabilityAnalysis) SetDpvActive(v string) {
	o.DpvActive = v
}

// GetDpvFootnotes returns the DpvFootnotes field value
func (o *DeliverabilityAnalysis) GetDpvFootnotes() []DpvFootnote {
	if o == nil {
		var ret []DpvFootnote
		return ret
	}

	return o.DpvFootnotes
}

// GetDpvFootnotesOk returns a tuple with the DpvFootnotes field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityAnalysis) GetDpvFootnotesOk() ([]DpvFootnote, bool) {
	if o == nil {
		return nil, false
	}
	return o.DpvFootnotes, true
}

// SetDpvFootnotes sets field value
func (o *DeliverabilityAnalysis) SetDpvFootnotes(v []DpvFootnote) {
	o.DpvFootnotes = v
}

// GetEwsMatch returns the EwsMatch field value
func (o *DeliverabilityAnalysis) GetEwsMatch() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EwsMatch
}

// GetEwsMatchOk returns a tuple with the EwsMatch field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityAnalysis) GetEwsMatchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EwsMatch, true
}

// SetEwsMatch sets field value
func (o *DeliverabilityAnalysis) SetEwsMatch(v bool) {
	o.EwsMatch = v
}

// GetLacsIndicator returns the LacsIndicator field value
func (o *DeliverabilityAnalysis) GetLacsIndicator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LacsIndicator
}

// GetLacsIndicatorOk returns a tuple with the LacsIndicator field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityAnalysis) GetLacsIndicatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LacsIndicator, true
}

// SetLacsIndicator sets field value
func (o *DeliverabilityAnalysis) SetLacsIndicator(v string) {
	o.LacsIndicator = v
}

// GetLacsReturnCode returns the LacsReturnCode field value
func (o *DeliverabilityAnalysis) GetLacsReturnCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LacsReturnCode
}

// GetLacsReturnCodeOk returns a tuple with the LacsReturnCode field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityAnalysis) GetLacsReturnCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LacsReturnCode, true
}

// SetLacsReturnCode sets field value
func (o *DeliverabilityAnalysis) SetLacsReturnCode(v string) {
	o.LacsReturnCode = v
}

// GetSuiteReturnCode returns the SuiteReturnCode field value
func (o *DeliverabilityAnalysis) GetSuiteReturnCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuiteReturnCode
}

// GetSuiteReturnCodeOk returns a tuple with the SuiteReturnCode field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityAnalysis) GetSuiteReturnCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuiteReturnCode, true
}

// SetSuiteReturnCode sets field value
func (o *DeliverabilityAnalysis) SetSuiteReturnCode(v string) {
	o.SuiteReturnCode = v
}

func (o DeliverabilityAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dpv_confirmation"] = o.DpvConfirmation
	}
	if true {
		toSerialize["dpv_cmra"] = o.DpvCmra
	}
	if true {
		toSerialize["dpv_vacant"] = o.DpvVacant
	}
	if true {
		toSerialize["dpv_active"] = o.DpvActive
	}
	if true {
		toSerialize["dpv_footnotes"] = o.DpvFootnotes
	}
	if true {
		toSerialize["ews_match"] = o.EwsMatch
	}
	if true {
		toSerialize["lacs_indicator"] = o.LacsIndicator
	}
	if true {
		toSerialize["lacs_return_code"] = o.LacsReturnCode
	}
	if true {
		toSerialize["suite_return_code"] = o.SuiteReturnCode
	}
	return json.Marshal(toSerialize)
}

type NullableDeliverabilityAnalysis struct {
	value *DeliverabilityAnalysis
	isSet bool
}

func (v NullableDeliverabilityAnalysis) Get() *DeliverabilityAnalysis {
	return v.value
}

func (v *NullableDeliverabilityAnalysis) Set(val *DeliverabilityAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverabilityAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverabilityAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverabilityAnalysis(val *DeliverabilityAnalysis) *NullableDeliverabilityAnalysis {
	return &NullableDeliverabilityAnalysis{value: val, isSet: true}
}

func (v NullableDeliverabilityAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverabilityAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


