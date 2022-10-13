# MultiLineAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | Pointer to **NullableString** | The intended recipient, typically a person&#39;s or firm&#39;s name. | [optional] 
**Company** | Pointer to **NullableString** | Either &#x60;name&#x60; or &#x60;company&#x60; is required, you may also add both. | [optional] 
**PrimaryLine** | **string** | The primary delivery line (usually the street address) of the address. Combination of the following applicable &#x60;components&#x60;: * &#x60;primary_number&#x60; * &#x60;street_predirection&#x60; * &#x60;street_name&#x60; * &#x60;street_suffix&#x60; * &#x60;street_postdirection&#x60; * &#x60;secondary_designator&#x60; * &#x60;secondary_number&#x60; * &#x60;pmb_designator&#x60; * &#x60;pmb_number&#x60;  | 
**SecondaryLine** | Pointer to **string** | The secondary delivery line of the address. This field is typically empty but may contain information if &#x60;primary_line&#x60; is too long.  | [optional] 
**Urbanization** | Pointer to **string** | Only present for addresses in Puerto Rico. An urbanization refers to an area, sector, or development within a city. See [USPS documentation](https://pe.usps.com/text/pub28/28api_008.htm#:~:text&#x3D;I51.,-4%20Urbanizations&amp;text&#x3D;In%20Puerto%20Rico%2C%20identical%20street,placed%20before%20the%20urbanization%20name.) for clarification.  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_3166-2:US\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166-2&lt;/a&gt; two letter code or subdivision name for the state. &#x60;city&#x60; and &#x60;state&#x60; are required if no &#x60;zip_code&#x60; is passed. | [optional] 
**ZipCode** | Pointer to **string** | Required if &#x60;city&#x60; and &#x60;state&#x60; are not passed in. If included, must be formatted as a US ZIP or ZIP+4 (e.g. &#x60;94107&#x60;, &#x60;941072282&#x60;, &#x60;94107-2282&#x60;). | [optional] 

## Methods

### NewMultiLineAddress

`func NewMultiLineAddress(primaryLine string, ) *MultiLineAddress`

NewMultiLineAddress instantiates a new MultiLineAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiLineAddressWithDefaults

`func NewMultiLineAddressWithDefaults() *MultiLineAddress`

NewMultiLineAddressWithDefaults instantiates a new MultiLineAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *MultiLineAddress) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *MultiLineAddress) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *MultiLineAddress) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *MultiLineAddress) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *MultiLineAddress) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *MultiLineAddress) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetCompany

`func (o *MultiLineAddress) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *MultiLineAddress) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *MultiLineAddress) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *MultiLineAddress) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *MultiLineAddress) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *MultiLineAddress) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetPrimaryLine

`func (o *MultiLineAddress) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *MultiLineAddress) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *MultiLineAddress) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.


### GetSecondaryLine

`func (o *MultiLineAddress) GetSecondaryLine() string`

GetSecondaryLine returns the SecondaryLine field if non-nil, zero value otherwise.

### GetSecondaryLineOk

`func (o *MultiLineAddress) GetSecondaryLineOk() (*string, bool)`

GetSecondaryLineOk returns a tuple with the SecondaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLine

`func (o *MultiLineAddress) SetSecondaryLine(v string)`

SetSecondaryLine sets SecondaryLine field to given value.

### HasSecondaryLine

`func (o *MultiLineAddress) HasSecondaryLine() bool`

HasSecondaryLine returns a boolean if a field has been set.

### GetUrbanization

`func (o *MultiLineAddress) GetUrbanization() string`

GetUrbanization returns the Urbanization field if non-nil, zero value otherwise.

### GetUrbanizationOk

`func (o *MultiLineAddress) GetUrbanizationOk() (*string, bool)`

GetUrbanizationOk returns a tuple with the Urbanization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrbanization

`func (o *MultiLineAddress) SetUrbanization(v string)`

SetUrbanization sets Urbanization field to given value.

### HasUrbanization

`func (o *MultiLineAddress) HasUrbanization() bool`

HasUrbanization returns a boolean if a field has been set.

### GetCity

`func (o *MultiLineAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MultiLineAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MultiLineAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MultiLineAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *MultiLineAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MultiLineAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MultiLineAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MultiLineAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *MultiLineAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *MultiLineAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *MultiLineAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *MultiLineAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


