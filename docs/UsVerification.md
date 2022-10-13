# UsVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;us_ver_&#x60;. | [optional] 
**Recipient** | Pointer to **NullableString** | The intended recipient, typically a person&#39;s or firm&#39;s name. | [optional] 
**PrimaryLine** | Pointer to **string** | The primary delivery line (usually the street address) of the address. Combination of the following applicable &#x60;components&#x60;: * &#x60;primary_number&#x60; * &#x60;street_predirection&#x60; * &#x60;street_name&#x60; * &#x60;street_suffix&#x60; * &#x60;street_postdirection&#x60; * &#x60;secondary_designator&#x60; * &#x60;secondary_number&#x60; * &#x60;pmb_designator&#x60; * &#x60;pmb_number&#x60;  | [optional] 
**SecondaryLine** | Pointer to **string** | The secondary delivery line of the address. This field is typically empty but may contain information if &#x60;primary_line&#x60; is too long.  | [optional] 
**Urbanization** | Pointer to **string** | Only present for addresses in Puerto Rico. An urbanization refers to an area, sector, or development within a city. See [USPS documentation](https://pe.usps.com/text/pub28/28api_008.htm#:~:text&#x3D;I51.,-4%20Urbanizations&amp;text&#x3D;In%20Puerto%20Rico%2C%20identical%20street,placed%20before%20the%20urbanization%20name.) for clarification.  | [optional] 
**LastLine** | Pointer to **string** | Combination of the following applicable &#x60;components&#x60;: * City (&#x60;city&#x60;) * State (&#x60;state&#x60;) * ZIP code (&#x60;zip_code&#x60;) * ZIP+4 (&#x60;zip_code_plus_4&#x60;)  | [optional] 
**Deliverability** | Pointer to **string** | Summarizes the deliverability of the &#x60;us_verification&#x60; object. For full details, see the &#x60;deliverability_analysis&#x60; field. Possible values are: * &#x60;deliverable&#x60; – The address is deliverable by the USPS. * &#x60;deliverable_unnecessary_unit&#x60; – The address is deliverable, but the secondary unit information is unnecessary. * &#x60;deliverable_incorrect_unit&#x60; – The address is deliverable to the building&#39;s default address but the secondary unit provided may not exist. There is a chance the mail will not reach the intended recipient. * &#x60;deliverable_missing_unit&#x60; – The address is deliverable to the building&#39;s default address but is missing secondary unit information. There is a chance the mail will not reach the intended recipient. * &#x60;undeliverable&#x60; – The address is not deliverable according to the USPS.  | [optional] 
**Components** | Pointer to [**UsComponents**](UsComponents.md) |  | [optional] 
**DeliverabilityAnalysis** | Pointer to [**DeliverabilityAnalysis**](DeliverabilityAnalysis.md) |  | [optional] 
**LobConfidenceScore** | Pointer to [**LobConfidenceScore**](LobConfidenceScore.md) |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "us_verification"]

## Methods

### NewUsVerification

`func NewUsVerification() *UsVerification`

NewUsVerification instantiates a new UsVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsVerificationWithDefaults

`func NewUsVerificationWithDefaults() *UsVerification`

NewUsVerificationWithDefaults instantiates a new UsVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsVerification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsVerification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsVerification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsVerification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipient

`func (o *UsVerification) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *UsVerification) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *UsVerification) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *UsVerification) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *UsVerification) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *UsVerification) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetPrimaryLine

`func (o *UsVerification) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *UsVerification) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *UsVerification) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.

### HasPrimaryLine

`func (o *UsVerification) HasPrimaryLine() bool`

HasPrimaryLine returns a boolean if a field has been set.

### GetSecondaryLine

`func (o *UsVerification) GetSecondaryLine() string`

GetSecondaryLine returns the SecondaryLine field if non-nil, zero value otherwise.

### GetSecondaryLineOk

`func (o *UsVerification) GetSecondaryLineOk() (*string, bool)`

GetSecondaryLineOk returns a tuple with the SecondaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLine

`func (o *UsVerification) SetSecondaryLine(v string)`

SetSecondaryLine sets SecondaryLine field to given value.

### HasSecondaryLine

`func (o *UsVerification) HasSecondaryLine() bool`

HasSecondaryLine returns a boolean if a field has been set.

### GetUrbanization

`func (o *UsVerification) GetUrbanization() string`

GetUrbanization returns the Urbanization field if non-nil, zero value otherwise.

### GetUrbanizationOk

`func (o *UsVerification) GetUrbanizationOk() (*string, bool)`

GetUrbanizationOk returns a tuple with the Urbanization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrbanization

`func (o *UsVerification) SetUrbanization(v string)`

SetUrbanization sets Urbanization field to given value.

### HasUrbanization

`func (o *UsVerification) HasUrbanization() bool`

HasUrbanization returns a boolean if a field has been set.

### GetLastLine

`func (o *UsVerification) GetLastLine() string`

GetLastLine returns the LastLine field if non-nil, zero value otherwise.

### GetLastLineOk

`func (o *UsVerification) GetLastLineOk() (*string, bool)`

GetLastLineOk returns a tuple with the LastLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLine

`func (o *UsVerification) SetLastLine(v string)`

SetLastLine sets LastLine field to given value.

### HasLastLine

`func (o *UsVerification) HasLastLine() bool`

HasLastLine returns a boolean if a field has been set.

### GetDeliverability

`func (o *UsVerification) GetDeliverability() string`

GetDeliverability returns the Deliverability field if non-nil, zero value otherwise.

### GetDeliverabilityOk

`func (o *UsVerification) GetDeliverabilityOk() (*string, bool)`

GetDeliverabilityOk returns a tuple with the Deliverability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverability

`func (o *UsVerification) SetDeliverability(v string)`

SetDeliverability sets Deliverability field to given value.

### HasDeliverability

`func (o *UsVerification) HasDeliverability() bool`

HasDeliverability returns a boolean if a field has been set.

### GetComponents

`func (o *UsVerification) GetComponents() UsComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *UsVerification) GetComponentsOk() (*UsComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *UsVerification) SetComponents(v UsComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *UsVerification) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetDeliverabilityAnalysis

`func (o *UsVerification) GetDeliverabilityAnalysis() DeliverabilityAnalysis`

GetDeliverabilityAnalysis returns the DeliverabilityAnalysis field if non-nil, zero value otherwise.

### GetDeliverabilityAnalysisOk

`func (o *UsVerification) GetDeliverabilityAnalysisOk() (*DeliverabilityAnalysis, bool)`

GetDeliverabilityAnalysisOk returns a tuple with the DeliverabilityAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverabilityAnalysis

`func (o *UsVerification) SetDeliverabilityAnalysis(v DeliverabilityAnalysis)`

SetDeliverabilityAnalysis sets DeliverabilityAnalysis field to given value.

### HasDeliverabilityAnalysis

`func (o *UsVerification) HasDeliverabilityAnalysis() bool`

HasDeliverabilityAnalysis returns a boolean if a field has been set.

### GetLobConfidenceScore

`func (o *UsVerification) GetLobConfidenceScore() LobConfidenceScore`

GetLobConfidenceScore returns the LobConfidenceScore field if non-nil, zero value otherwise.

### GetLobConfidenceScoreOk

`func (o *UsVerification) GetLobConfidenceScoreOk() (*LobConfidenceScore, bool)`

GetLobConfidenceScoreOk returns a tuple with the LobConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobConfidenceScore

`func (o *UsVerification) SetLobConfidenceScore(v LobConfidenceScore)`

SetLobConfidenceScore sets LobConfidenceScore field to given value.

### HasLobConfidenceScore

`func (o *UsVerification) HasLobConfidenceScore() bool`

HasLobConfidenceScore returns a boolean if a field has been set.

### GetObject

`func (o *UsVerification) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UsVerification) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UsVerification) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *UsVerification) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


