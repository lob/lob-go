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

// ZipLookupCity struct for ZipLookupCity
type ZipLookupCity struct {
	City string `json:"city"`
	// The [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) two letter code for the state. 
	State string `json:"state"`
	// County name of the address city.
	County string `json:"county"`
	// A 5-digit [FIPS county code](https://en.wikipedia.org/wiki/FIPS_county_code) which uniquely identifies `components[county]`. It consists of a 2-digit state code and a 3-digit county code. 
	CountyFips string `json:"county_fips"`
	// Indicates whether or not the city is the [USPS default city](https://en.wikipedia.org/wiki/ZIP_Code#ZIP_Codes_and_previous_zoning_lines) (preferred city) of a ZIP code. There is only one preferred city per ZIP code, which will always be in position 0 in the array of cities. 
	Preferred bool `json:"preferred"`
}

// NewZipLookupCity instantiates a new ZipLookupCity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZipLookupCity(city string, state string, county string, countyFips string, preferred bool) *ZipLookupCity {
	this := ZipLookupCity{}
	this.City = city
	this.State = state
	this.County = county
	this.CountyFips = countyFips
	this.Preferred = preferred
	return &this
}

// NewZipLookupCityWithDefaults instantiates a new ZipLookupCity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZipLookupCityWithDefaults() *ZipLookupCity {
	this := ZipLookupCity{}
	return &this
}

// GetCity returns the City field value
func (o *ZipLookupCity) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *ZipLookupCity) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *ZipLookupCity) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *ZipLookupCity) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ZipLookupCity) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ZipLookupCity) SetState(v string) {
	o.State = v
}

// GetCounty returns the County field value
func (o *ZipLookupCity) GetCounty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.County
}

// GetCountyOk returns a tuple with the County field value
// and a boolean to check if the value has been set.
func (o *ZipLookupCity) GetCountyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.County, true
}

// SetCounty sets field value
func (o *ZipLookupCity) SetCounty(v string) {
	o.County = v
}

// GetCountyFips returns the CountyFips field value
func (o *ZipLookupCity) GetCountyFips() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountyFips
}

// GetCountyFipsOk returns a tuple with the CountyFips field value
// and a boolean to check if the value has been set.
func (o *ZipLookupCity) GetCountyFipsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountyFips, true
}

// SetCountyFips sets field value
func (o *ZipLookupCity) SetCountyFips(v string) {
	o.CountyFips = v
}

// GetPreferred returns the Preferred field value
func (o *ZipLookupCity) GetPreferred() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value
// and a boolean to check if the value has been set.
func (o *ZipLookupCity) GetPreferredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preferred, true
}

// SetPreferred sets field value
func (o *ZipLookupCity) SetPreferred(v bool) {
	o.Preferred = v
}

func (o ZipLookupCity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["county"] = o.County
	}
	if true {
		toSerialize["county_fips"] = o.CountyFips
	}
	if true {
		toSerialize["preferred"] = o.Preferred
	}
	return json.Marshal(toSerialize)
}

type NullableZipLookupCity struct {
	value *ZipLookupCity
	isSet bool
}

func (v NullableZipLookupCity) Get() *ZipLookupCity {
	return v.value
}

func (v *NullableZipLookupCity) Set(val *ZipLookupCity) {
	v.value = val
	v.isSet = true
}

func (v NullableZipLookupCity) IsSet() bool {
	return v.isSet
}

func (v *NullableZipLookupCity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZipLookupCity(val *ZipLookupCity) *NullableZipLookupCity {
	return &NullableZipLookupCity{value: val, isSet: true}
}

func (v NullableZipLookupCity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZipLookupCity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


