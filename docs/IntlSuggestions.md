# IntlSuggestions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryNumberRange** | **string** | The primary number range of the address that identifies a building at street level.  | 
**PrimaryLine** | **string** | The primary delivery line (usually the street address) of the address. Combination of the following applicable &#x60;components&#x60; (primary number &amp; secondary information may be missing or inaccurate): * &#x60;primary_number&#x60; * &#x60;street_predirection&#x60; * &#x60;street_name&#x60; * &#x60;street_suffix&#x60; * &#x60;street_postdirection&#x60; * &#x60;secondary_designator&#x60; * &#x60;secondary_number&#x60; * &#x60;pmb_designator&#x60; * &#x60;pmb_number&#x60;  | 
**City** | **string** |  | 
**State** | **string** | The [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) two letter code for the state.  | 
**Country** | [**CountryExtended**](CountryExtended.md) |  | 
**ZipCode** | **string** | A 5-digit zip code. Left empty if a test key is used. | 
**Object** | Pointer to **string** | Value is resource type. | [optional] [default to "intl_autocompletion"]

## Methods

### NewIntlSuggestions

`func NewIntlSuggestions(primaryNumberRange string, primaryLine string, city string, state string, country CountryExtended, zipCode string, ) *IntlSuggestions`

NewIntlSuggestions instantiates a new IntlSuggestions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntlSuggestionsWithDefaults

`func NewIntlSuggestionsWithDefaults() *IntlSuggestions`

NewIntlSuggestionsWithDefaults instantiates a new IntlSuggestions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryNumberRange

`func (o *IntlSuggestions) GetPrimaryNumberRange() string`

GetPrimaryNumberRange returns the PrimaryNumberRange field if non-nil, zero value otherwise.

### GetPrimaryNumberRangeOk

`func (o *IntlSuggestions) GetPrimaryNumberRangeOk() (*string, bool)`

GetPrimaryNumberRangeOk returns a tuple with the PrimaryNumberRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNumberRange

`func (o *IntlSuggestions) SetPrimaryNumberRange(v string)`

SetPrimaryNumberRange sets PrimaryNumberRange field to given value.


### GetPrimaryLine

`func (o *IntlSuggestions) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *IntlSuggestions) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *IntlSuggestions) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.


### GetCity

`func (o *IntlSuggestions) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *IntlSuggestions) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *IntlSuggestions) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *IntlSuggestions) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IntlSuggestions) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IntlSuggestions) SetState(v string)`

SetState sets State field to given value.


### GetCountry

`func (o *IntlSuggestions) GetCountry() CountryExtended`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IntlSuggestions) GetCountryOk() (*CountryExtended, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IntlSuggestions) SetCountry(v CountryExtended)`

SetCountry sets Country field to given value.


### GetZipCode

`func (o *IntlSuggestions) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *IntlSuggestions) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *IntlSuggestions) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.


### GetObject

`func (o *IntlSuggestions) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IntlSuggestions) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IntlSuggestions) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *IntlSuggestions) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


