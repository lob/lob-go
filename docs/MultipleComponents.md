# MultipleComponents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | Pointer to **NullableString** | The intended recipient, typically a person&#39;s or firm&#39;s name. | [optional] 
**PrimaryLine** | **string** | The primary delivery line (usually the street address) of the address. Combination of the following applicable &#x60;components&#x60;: * &#x60;primary_number&#x60; * &#x60;street_predirection&#x60; * &#x60;street_name&#x60; * &#x60;street_suffix&#x60; * &#x60;street_postdirection&#x60; * &#x60;secondary_designator&#x60; * &#x60;secondary_number&#x60; * &#x60;pmb_designator&#x60; * &#x60;pmb_number&#x60;  | 
**SecondaryLine** | Pointer to **string** | The secondary delivery line of the address. This field is typically empty but may contain information if &#x60;primary_line&#x60; is too long.  | [optional] 
**Urbanization** | Pointer to **string** | Only present for addresses in Puerto Rico. An urbanization refers to an area, sector, or development within a city. See [USPS documentation](https://pe.usps.com/text/pub28/28api_008.htm#:~:text&#x3D;I51.,-4%20Urbanizations&amp;text&#x3D;In%20Puerto%20Rico%2C%20identical%20street,placed%20before%20the%20urbanization%20name.) for clarification.  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2:US) two letter code or subdivision name for the state. &#x60;city&#x60; and &#x60;state&#x60; are required if no &#x60;zip_code&#x60; is passed. | [optional] 
**ZipCode** | Pointer to **string** | Required if &#x60;city&#x60; and &#x60;state&#x60; are not passed in. If included, must be formatted as a US ZIP or ZIP+4 (e.g. &#x60;94107&#x60;, &#x60;941072282&#x60;, &#x60;94107-2282&#x60;). | [optional] 

## Methods

### NewMultipleComponents

`func NewMultipleComponents(primaryLine string, ) *MultipleComponents`

NewMultipleComponents instantiates a new MultipleComponents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleComponentsWithDefaults

`func NewMultipleComponentsWithDefaults() *MultipleComponents`

NewMultipleComponentsWithDefaults instantiates a new MultipleComponents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *MultipleComponents) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *MultipleComponents) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *MultipleComponents) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *MultipleComponents) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *MultipleComponents) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *MultipleComponents) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetPrimaryLine

`func (o *MultipleComponents) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *MultipleComponents) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *MultipleComponents) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.


### GetSecondaryLine

`func (o *MultipleComponents) GetSecondaryLine() string`

GetSecondaryLine returns the SecondaryLine field if non-nil, zero value otherwise.

### GetSecondaryLineOk

`func (o *MultipleComponents) GetSecondaryLineOk() (*string, bool)`

GetSecondaryLineOk returns a tuple with the SecondaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLine

`func (o *MultipleComponents) SetSecondaryLine(v string)`

SetSecondaryLine sets SecondaryLine field to given value.

### HasSecondaryLine

`func (o *MultipleComponents) HasSecondaryLine() bool`

HasSecondaryLine returns a boolean if a field has been set.

### GetUrbanization

`func (o *MultipleComponents) GetUrbanization() string`

GetUrbanization returns the Urbanization field if non-nil, zero value otherwise.

### GetUrbanizationOk

`func (o *MultipleComponents) GetUrbanizationOk() (*string, bool)`

GetUrbanizationOk returns a tuple with the Urbanization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrbanization

`func (o *MultipleComponents) SetUrbanization(v string)`

SetUrbanization sets Urbanization field to given value.

### HasUrbanization

`func (o *MultipleComponents) HasUrbanization() bool`

HasUrbanization returns a boolean if a field has been set.

### GetCity

`func (o *MultipleComponents) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MultipleComponents) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MultipleComponents) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MultipleComponents) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *MultipleComponents) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MultipleComponents) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MultipleComponents) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MultipleComponents) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *MultipleComponents) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *MultipleComponents) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *MultipleComponents) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *MultipleComponents) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


