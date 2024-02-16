# UsVerificationOrError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;us_ver_&#x60;. | [optional] 
**Recipient** | Pointer to **NullableString** | The intended recipient, typically a person&#39;s or firm&#39;s name. | [optional] 
**PrimaryLine** | Pointer to **string** | The primary delivery line (usually the street address) of the address. Combination of the following applicable &#x60;components&#x60;: * &#x60;primary_number&#x60; * &#x60;street_predirection&#x60; * &#x60;street_name&#x60; * &#x60;street_suffix&#x60; * &#x60;street_postdirection&#x60; * &#x60;secondary_designator&#x60; * &#x60;secondary_number&#x60; * &#x60;pmb_designator&#x60; * &#x60;pmb_number&#x60;  | [optional] 
**SecondaryLine** | Pointer to **string** | The secondary delivery line of the address. This field is typically empty but may contain information if &#x60;primary_line&#x60; is too long.  | [optional] 
**Urbanization** | Pointer to **string** | Only present for addresses in Puerto Rico. An urbanization refers to an area, sector, or development within a city. See [USPS documentation](https://pe.usps.com/text/pub28/28api_008.htm#:~:text&#x3D;I51.,-4%20Urbanizations&amp;text&#x3D;In%20Puerto%20Rico%2C%20identical%20street,placed%20before%20the%20urbanization%20name.) for clarification.  | [optional] 
**LastLine** | Pointer to **string** |  | [optional] 
**Deliverability** | Pointer to **string** |  | [optional] 
**Components** | Pointer to [**UsComponents**](UsComponents.md) |  | [optional] 
**DeliverabilityAnalysis** | Pointer to [**DeliverabilityAnalysis**](DeliverabilityAnalysis.md) |  | [optional] 
**LobConfidenceScore** | Pointer to [**LobConfidenceScore**](LobConfidenceScore.md) |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "us_verification"]
**TransientId** | Pointer to **string** | ID that is returned in the response body for the verification  | [optional] 
**Error** | Pointer to [**BulkError**](BulkError.md) |  | [optional] 

## Methods

### NewUsVerificationOrError

`func NewUsVerificationOrError() *UsVerificationOrError`

NewUsVerificationOrError instantiates a new UsVerificationOrError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsVerificationOrErrorWithDefaults

`func NewUsVerificationOrErrorWithDefaults() *UsVerificationOrError`

NewUsVerificationOrErrorWithDefaults instantiates a new UsVerificationOrError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsVerificationOrError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsVerificationOrError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsVerificationOrError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsVerificationOrError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipient

`func (o *UsVerificationOrError) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *UsVerificationOrError) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *UsVerificationOrError) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *UsVerificationOrError) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *UsVerificationOrError) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *UsVerificationOrError) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetPrimaryLine

`func (o *UsVerificationOrError) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *UsVerificationOrError) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *UsVerificationOrError) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.

### HasPrimaryLine

`func (o *UsVerificationOrError) HasPrimaryLine() bool`

HasPrimaryLine returns a boolean if a field has been set.

### GetSecondaryLine

`func (o *UsVerificationOrError) GetSecondaryLine() string`

GetSecondaryLine returns the SecondaryLine field if non-nil, zero value otherwise.

### GetSecondaryLineOk

`func (o *UsVerificationOrError) GetSecondaryLineOk() (*string, bool)`

GetSecondaryLineOk returns a tuple with the SecondaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLine

`func (o *UsVerificationOrError) SetSecondaryLine(v string)`

SetSecondaryLine sets SecondaryLine field to given value.

### HasSecondaryLine

`func (o *UsVerificationOrError) HasSecondaryLine() bool`

HasSecondaryLine returns a boolean if a field has been set.

### GetUrbanization

`func (o *UsVerificationOrError) GetUrbanization() string`

GetUrbanization returns the Urbanization field if non-nil, zero value otherwise.

### GetUrbanizationOk

`func (o *UsVerificationOrError) GetUrbanizationOk() (*string, bool)`

GetUrbanizationOk returns a tuple with the Urbanization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrbanization

`func (o *UsVerificationOrError) SetUrbanization(v string)`

SetUrbanization sets Urbanization field to given value.

### HasUrbanization

`func (o *UsVerificationOrError) HasUrbanization() bool`

HasUrbanization returns a boolean if a field has been set.

### GetLastLine

`func (o *UsVerificationOrError) GetLastLine() string`

GetLastLine returns the LastLine field if non-nil, zero value otherwise.

### GetLastLineOk

`func (o *UsVerificationOrError) GetLastLineOk() (*string, bool)`

GetLastLineOk returns a tuple with the LastLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLine

`func (o *UsVerificationOrError) SetLastLine(v string)`

SetLastLine sets LastLine field to given value.

### HasLastLine

`func (o *UsVerificationOrError) HasLastLine() bool`

HasLastLine returns a boolean if a field has been set.

### GetDeliverability

`func (o *UsVerificationOrError) GetDeliverability() string`

GetDeliverability returns the Deliverability field if non-nil, zero value otherwise.

### GetDeliverabilityOk

`func (o *UsVerificationOrError) GetDeliverabilityOk() (*string, bool)`

GetDeliverabilityOk returns a tuple with the Deliverability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverability

`func (o *UsVerificationOrError) SetDeliverability(v string)`

SetDeliverability sets Deliverability field to given value.

### HasDeliverability

`func (o *UsVerificationOrError) HasDeliverability() bool`

HasDeliverability returns a boolean if a field has been set.

### GetComponents

`func (o *UsVerificationOrError) GetComponents() UsComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *UsVerificationOrError) GetComponentsOk() (*UsComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *UsVerificationOrError) SetComponents(v UsComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *UsVerificationOrError) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetDeliverabilityAnalysis

`func (o *UsVerificationOrError) GetDeliverabilityAnalysis() DeliverabilityAnalysis`

GetDeliverabilityAnalysis returns the DeliverabilityAnalysis field if non-nil, zero value otherwise.

### GetDeliverabilityAnalysisOk

`func (o *UsVerificationOrError) GetDeliverabilityAnalysisOk() (*DeliverabilityAnalysis, bool)`

GetDeliverabilityAnalysisOk returns a tuple with the DeliverabilityAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverabilityAnalysis

`func (o *UsVerificationOrError) SetDeliverabilityAnalysis(v DeliverabilityAnalysis)`

SetDeliverabilityAnalysis sets DeliverabilityAnalysis field to given value.

### HasDeliverabilityAnalysis

`func (o *UsVerificationOrError) HasDeliverabilityAnalysis() bool`

HasDeliverabilityAnalysis returns a boolean if a field has been set.

### GetLobConfidenceScore

`func (o *UsVerificationOrError) GetLobConfidenceScore() LobConfidenceScore`

GetLobConfidenceScore returns the LobConfidenceScore field if non-nil, zero value otherwise.

### GetLobConfidenceScoreOk

`func (o *UsVerificationOrError) GetLobConfidenceScoreOk() (*LobConfidenceScore, bool)`

GetLobConfidenceScoreOk returns a tuple with the LobConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobConfidenceScore

`func (o *UsVerificationOrError) SetLobConfidenceScore(v LobConfidenceScore)`

SetLobConfidenceScore sets LobConfidenceScore field to given value.

### HasLobConfidenceScore

`func (o *UsVerificationOrError) HasLobConfidenceScore() bool`

HasLobConfidenceScore returns a boolean if a field has been set.

### GetObject

`func (o *UsVerificationOrError) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UsVerificationOrError) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UsVerificationOrError) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *UsVerificationOrError) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetTransientId

`func (o *UsVerificationOrError) GetTransientId() string`

GetTransientId returns the TransientId field if non-nil, zero value otherwise.

### GetTransientIdOk

`func (o *UsVerificationOrError) GetTransientIdOk() (*string, bool)`

GetTransientIdOk returns a tuple with the TransientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransientId

`func (o *UsVerificationOrError) SetTransientId(v string)`

SetTransientId sets TransientId field to given value.

### HasTransientId

`func (o *UsVerificationOrError) HasTransientId() bool`

HasTransientId returns a boolean if a field has been set.

### GetError

`func (o *UsVerificationOrError) GetError() BulkError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UsVerificationOrError) GetErrorOk() (*BulkError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UsVerificationOrError) SetError(v BulkError)`

SetError sets Error field to given value.

### HasError

`func (o *UsVerificationOrError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


