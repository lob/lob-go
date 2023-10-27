# IdentityValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;id_validation_&#x60;. | [optional] 
**Recipient** | Pointer to **NullableString** | The intended recipient, typically a person&#39;s or firm&#39;s name. | [optional] 
**PrimaryLine** | Pointer to **string** | The primary delivery line (usually the street address) of the address. Combination of the following applicable &#x60;components&#x60;: * &#x60;primary_number&#x60; * &#x60;street_predirection&#x60; * &#x60;street_name&#x60; * &#x60;street_suffix&#x60; * &#x60;street_postdirection&#x60; * &#x60;secondary_designator&#x60; * &#x60;secondary_number&#x60; * &#x60;pmb_designator&#x60; * &#x60;pmb_number&#x60;  | [optional] 
**SecondaryLine** | Pointer to **string** | The secondary delivery line of the address. This field is typically empty but may contain information if &#x60;primary_line&#x60; is too long.  | [optional] 
**Urbanization** | Pointer to **string** | Only present for addresses in Puerto Rico. An urbanization refers to an area, sector, or development within a city. See [USPS documentation](https://pe.usps.com/text/pub28/28api_008.htm#:~:text&#x3D;I51.,-4%20Urbanizations&amp;text&#x3D;In%20Puerto%20Rico%2C%20identical%20street,placed%20before%20the%20urbanization%20name.) for clarification.  | [optional] 
**LastLine** | Pointer to **string** | Combination of the following applicable &#x60;components&#x60;: * City (&#x60;city&#x60;) * State (&#x60;state&#x60;) * ZIP code (&#x60;zip_code&#x60;) * ZIP+4 (&#x60;zip_code_plus_4&#x60;)  | [optional] 
**Score** | Pointer to **NullableFloat32** | A numerical score between 0 and 100 that represents the likelihood the provided name is associated with a physical address.  | [optional] 
**Confidence** | Pointer to **string** | Indicates the likelihood the recipient name and address match based on our custom internal calculation. Possible values are: - &#x60;high&#x60; — Has a Lob confidence score greater than 70. - &#x60;medium&#x60; — Has a Lob confidence score between 40 and 70. - &#x60;low&#x60; — Has a Lob confidence score less than 40. - &#x60;\&quot;\&quot;&#x60; — No tracking data exists for this address.  | [optional] 
**Object** | Pointer to **string** | Value is resource type. | [optional] [default to "id_validation"]

## Methods

### NewIdentityValidation

`func NewIdentityValidation() *IdentityValidation`

NewIdentityValidation instantiates a new IdentityValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityValidationWithDefaults

`func NewIdentityValidationWithDefaults() *IdentityValidation`

NewIdentityValidationWithDefaults instantiates a new IdentityValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityValidation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityValidation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityValidation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityValidation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipient

`func (o *IdentityValidation) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *IdentityValidation) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *IdentityValidation) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *IdentityValidation) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *IdentityValidation) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *IdentityValidation) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetPrimaryLine

`func (o *IdentityValidation) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *IdentityValidation) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *IdentityValidation) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.

### HasPrimaryLine

`func (o *IdentityValidation) HasPrimaryLine() bool`

HasPrimaryLine returns a boolean if a field has been set.

### GetSecondaryLine

`func (o *IdentityValidation) GetSecondaryLine() string`

GetSecondaryLine returns the SecondaryLine field if non-nil, zero value otherwise.

### GetSecondaryLineOk

`func (o *IdentityValidation) GetSecondaryLineOk() (*string, bool)`

GetSecondaryLineOk returns a tuple with the SecondaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLine

`func (o *IdentityValidation) SetSecondaryLine(v string)`

SetSecondaryLine sets SecondaryLine field to given value.

### HasSecondaryLine

`func (o *IdentityValidation) HasSecondaryLine() bool`

HasSecondaryLine returns a boolean if a field has been set.

### GetUrbanization

`func (o *IdentityValidation) GetUrbanization() string`

GetUrbanization returns the Urbanization field if non-nil, zero value otherwise.

### GetUrbanizationOk

`func (o *IdentityValidation) GetUrbanizationOk() (*string, bool)`

GetUrbanizationOk returns a tuple with the Urbanization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrbanization

`func (o *IdentityValidation) SetUrbanization(v string)`

SetUrbanization sets Urbanization field to given value.

### HasUrbanization

`func (o *IdentityValidation) HasUrbanization() bool`

HasUrbanization returns a boolean if a field has been set.

### GetLastLine

`func (o *IdentityValidation) GetLastLine() string`

GetLastLine returns the LastLine field if non-nil, zero value otherwise.

### GetLastLineOk

`func (o *IdentityValidation) GetLastLineOk() (*string, bool)`

GetLastLineOk returns a tuple with the LastLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLine

`func (o *IdentityValidation) SetLastLine(v string)`

SetLastLine sets LastLine field to given value.

### HasLastLine

`func (o *IdentityValidation) HasLastLine() bool`

HasLastLine returns a boolean if a field has been set.

### GetScore

`func (o *IdentityValidation) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *IdentityValidation) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *IdentityValidation) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *IdentityValidation) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *IdentityValidation) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *IdentityValidation) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetConfidence

`func (o *IdentityValidation) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *IdentityValidation) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *IdentityValidation) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *IdentityValidation) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetObject

`func (o *IdentityValidation) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IdentityValidation) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IdentityValidation) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *IdentityValidation) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


