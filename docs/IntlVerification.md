# IntlVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;intl_ver_&#x60;. | [optional] 
**Recipient** | Pointer to **NullableString** | The intended recipient, typically a person&#39;s or firm&#39;s name. | [optional] 
**PrimaryLine** | Pointer to **string** | The primary delivery line (usually the street address) of the address.  | [optional] 
**SecondaryLine** | Pointer to **string** | The secondary delivery line of the address. This field is typically empty but may contain information if &#x60;primary_line&#x60; is too long.  | [optional] 
**LastLine** | Pointer to **string** | Combination of the following applicable &#x60;components&#x60;: * &#x60;city&#x60; * &#x60;state&#x60; * &#x60;zip_code&#x60; * &#x60;zip_code_plus_4&#x60;  | [optional] 
**Country** | Pointer to **string** | The country of the address. Will be returned as a 2 letter country short-name code (ISO 3166). | [optional] 
**Coverage** | Pointer to **string** | The coverage level for the country. This represents the maximum level of accuracy an input address can be verified to.  * &#x60;SUBBUILDING&#x60; - Coverage down to unit numbers. For example, in an apartment or a large building * &#x60;HOUSENUMBER/BUILDING&#x60; - Coverage down to house number. For example, the address where a house or building may be located * &#x60;STREET&#x60; - Coverage down to street. This means that we can verify that an street exists in a city, state, country * &#x60;LOCALITY&#x60; - Coverage down to city, state, or village or province. This means that we can verify that a city, village, province, or state exists in a country. Countries differ in how they define what is a province, state, city, village, etc. This attempts to group eveyrthing together. * &#x60;SPARSE&#x60; - Some addresses for this country exist in our databases  | [optional] 
**Deliverability** | Pointer to **string** | Summarizes the deliverability of the &#x60;intl_verification&#x60; object. Possible values are: * &#x60;deliverable&#x60; — The address is deliverable. * &#x60;deliverable_missing_info&#x60; — The address is missing some information, but is most likely deliverable. * &#x60;undeliverable&#x60; — The address is most likely not deliverable. Some components of the address (such as city or postal code) may have been found. * &#x60;no_match&#x60; — This address is not deliverable. No matching street could be found within the city or postal code.  | [optional] 
**Status** | Pointer to **string** | The status level for the country. This represents the maximum level of accuracy an input address can be verified to.  * &#x60;LV4&#x60; - Verified. The input data is correct. All input data was able to match in databases. * &#x60;LV3&#x60; - Verified. The input data is correct. All input data was able to match in databases &lt;em&gt;after&lt;/em&gt; some or all elements were standarized. The input data may also be using outdated city, state, or country names. * &#x60;LV2&#x60; - Verified. The input data is correct although some input data is unverifiable due to incomplete data. * &#x60;LV1&#x60; - Verified. The input data is acceptable but in an attempt to standarize user input, errors were introduced. * &#x60;LF4&#x60; - Fixed. The input data is matched and fixed. (Example: Brighon, UK -&gt; Brighton, UK) * &#x60;LF3&#x60; - Fixed. The input data is matched and fixed but some elements such as Subbuilding number and house number cannot be checked. * &#x60;LF2&#x60; - Fixed. The input data is matched but some elements such as Street cannot be checked. * &#x60;LF1&#x60; - Fixed. The input data is acceptable but in an attempt to standarize user input, errors were introduced. * &#x60;LM4&#x60; - Missing Info. The input data cannot be corrected completely. * &#x60;LM3&#x60; - Missing Info. The input data cannot be corrected completely and there were multiple matches found in databases. * &#x60;LM2&#x60; - Missing Info. The input data cannot be corrected completely and only some elements were found. * &#x60;LU1&#x60; - Unverified. The input data cannot be corrected or matched.  | [optional] 
**Components** | Pointer to [**IntlComponents**](IntlComponents.md) |  | [optional] 
**Object** | Pointer to **string** | Value is resource type. | [optional] [default to "intl_verification"]

## Methods

### NewIntlVerification

`func NewIntlVerification() *IntlVerification`

NewIntlVerification instantiates a new IntlVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntlVerificationWithDefaults

`func NewIntlVerificationWithDefaults() *IntlVerification`

NewIntlVerificationWithDefaults instantiates a new IntlVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntlVerification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntlVerification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntlVerification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntlVerification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipient

`func (o *IntlVerification) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *IntlVerification) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *IntlVerification) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *IntlVerification) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *IntlVerification) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *IntlVerification) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetPrimaryLine

`func (o *IntlVerification) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *IntlVerification) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *IntlVerification) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.

### HasPrimaryLine

`func (o *IntlVerification) HasPrimaryLine() bool`

HasPrimaryLine returns a boolean if a field has been set.

### GetSecondaryLine

`func (o *IntlVerification) GetSecondaryLine() string`

GetSecondaryLine returns the SecondaryLine field if non-nil, zero value otherwise.

### GetSecondaryLineOk

`func (o *IntlVerification) GetSecondaryLineOk() (*string, bool)`

GetSecondaryLineOk returns a tuple with the SecondaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLine

`func (o *IntlVerification) SetSecondaryLine(v string)`

SetSecondaryLine sets SecondaryLine field to given value.

### HasSecondaryLine

`func (o *IntlVerification) HasSecondaryLine() bool`

HasSecondaryLine returns a boolean if a field has been set.

### GetLastLine

`func (o *IntlVerification) GetLastLine() string`

GetLastLine returns the LastLine field if non-nil, zero value otherwise.

### GetLastLineOk

`func (o *IntlVerification) GetLastLineOk() (*string, bool)`

GetLastLineOk returns a tuple with the LastLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLine

`func (o *IntlVerification) SetLastLine(v string)`

SetLastLine sets LastLine field to given value.

### HasLastLine

`func (o *IntlVerification) HasLastLine() bool`

HasLastLine returns a boolean if a field has been set.

### GetCountry

`func (o *IntlVerification) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IntlVerification) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IntlVerification) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *IntlVerification) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCoverage

`func (o *IntlVerification) GetCoverage() string`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *IntlVerification) GetCoverageOk() (*string, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *IntlVerification) SetCoverage(v string)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *IntlVerification) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetDeliverability

`func (o *IntlVerification) GetDeliverability() string`

GetDeliverability returns the Deliverability field if non-nil, zero value otherwise.

### GetDeliverabilityOk

`func (o *IntlVerification) GetDeliverabilityOk() (*string, bool)`

GetDeliverabilityOk returns a tuple with the Deliverability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverability

`func (o *IntlVerification) SetDeliverability(v string)`

SetDeliverability sets Deliverability field to given value.

### HasDeliverability

`func (o *IntlVerification) HasDeliverability() bool`

HasDeliverability returns a boolean if a field has been set.

### GetStatus

`func (o *IntlVerification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntlVerification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntlVerification) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntlVerification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComponents

`func (o *IntlVerification) GetComponents() IntlComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *IntlVerification) GetComponentsOk() (*IntlComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *IntlVerification) SetComponents(v IntlComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *IntlVerification) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetObject

`func (o *IntlVerification) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IntlVerification) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IntlVerification) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *IntlVerification) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


