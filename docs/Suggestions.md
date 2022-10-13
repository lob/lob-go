# Suggestions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryLine** | **string** | The primary delivery line (usually the street address) of the address. Combination of the following applicable &#x60;components&#x60; (primary number &amp; secondary information may be missing or inaccurate): * &#x60;primary_number&#x60; * &#x60;street_predirection&#x60; * &#x60;street_name&#x60; * &#x60;street_suffix&#x60; * &#x60;street_postdirection&#x60; * &#x60;secondary_designator&#x60; * &#x60;secondary_number&#x60; * &#x60;pmb_designator&#x60; * &#x60;pmb_number&#x60;  | 
**City** | **string** |  | 
**State** | **string** | The [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) two letter code for the state.  | 
**ZipCode** | **string** | A 5-digit zip code. Left empty if a test key is used. | 
**Object** | Pointer to **string** | Value is resource type. | [optional] [default to "us_autocompletion"]

## Methods

### NewSuggestions

`func NewSuggestions(primaryLine string, city string, state string, zipCode string, ) *Suggestions`

NewSuggestions instantiates a new Suggestions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestionsWithDefaults

`func NewSuggestionsWithDefaults() *Suggestions`

NewSuggestionsWithDefaults instantiates a new Suggestions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryLine

`func (o *Suggestions) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *Suggestions) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *Suggestions) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.


### GetCity

`func (o *Suggestions) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Suggestions) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Suggestions) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *Suggestions) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Suggestions) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Suggestions) SetState(v string)`

SetState sets State field to given value.


### GetZipCode

`func (o *Suggestions) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Suggestions) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Suggestions) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.


### GetObject

`func (o *Suggestions) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Suggestions) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Suggestions) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Suggestions) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


