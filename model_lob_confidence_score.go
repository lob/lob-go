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

// LobConfidenceScore Lob Confidence Score is a nested object that provides a numerical value between 0-100 of the likelihood that an address is deliverable based on Lob’s mail delivery data to over half of US households.
type LobConfidenceScore struct {
	// A numerical score between 0 and 100 that represents the percentage of mailpieces Lob has sent to this addresses that have been delivered successfully over the past 2 years. Will be `null` if no tracking data exists for this address. 
	Score NullableFloat32 `json:"score,omitempty"`
	// indicates the likelihood that the address is a valid, mail-receiving address. Possible values are:   - `high` — Over 70% of mailpieces Lob has sent to this address were delivered successfully and recent mailings were also successful.   - `medium` — Between 40% and 70% of mailpieces Lob has sent to this address were delivered successfully.   - `low` — Less than 40% of mailpieces Lob has sent to this address were delivered successfully and recent mailings weren't successful.   - `\"\"` — No tracking data exists for this address or lob deliverability was unable to find a corresponding level of mail success. 
	Level *string `json:"level,omitempty"`
}

// NewLobConfidenceScore instantiates a new LobConfidenceScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLobConfidenceScore() *LobConfidenceScore {
	this := LobConfidenceScore{}
	return &this
}

// NewLobConfidenceScoreWithDefaults instantiates a new LobConfidenceScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLobConfidenceScoreWithDefaults() *LobConfidenceScore {
	this := LobConfidenceScore{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LobConfidenceScore) GetScore() float32 {
	if o == nil || o.Score.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LobConfidenceScore) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *LobConfidenceScore) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableFloat32 and assigns it to the Score field.
func (o *LobConfidenceScore) SetScore(v float32) {
	o.Score.Set(&v)
}
// SetScoreNil sets the value for Score to be an explicit nil
func (o *LobConfidenceScore) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *LobConfidenceScore) UnsetScore() {
	o.Score.Unset()
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *LobConfidenceScore) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LobConfidenceScore) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *LobConfidenceScore) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *LobConfidenceScore) SetLevel(v string) {
	o.Level = &v
}

func (o LobConfidenceScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableLobConfidenceScore struct {
	value *LobConfidenceScore
	isSet bool
}

func (v NullableLobConfidenceScore) Get() *LobConfidenceScore {
	return v.value
}

func (v *NullableLobConfidenceScore) Set(val *LobConfidenceScore) {
	v.value = val
	v.isSet = true
}

func (v NullableLobConfidenceScore) IsSet() bool {
	return v.isSet
}

func (v *NullableLobConfidenceScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLobConfidenceScore(val *LobConfidenceScore) *NullableLobConfidenceScore {
	return &NullableLobConfidenceScore{value: val, isSet: true}
}

func (v NullableLobConfidenceScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLobConfidenceScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


