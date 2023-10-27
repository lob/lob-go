# IntlVerificationOrError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;intl_ver_&#x60;. | [optional] 
**Recipient** | Pointer to **NullableString** | The intended recipient, typically a person&#39;s or firm&#39;s name. | [optional] 
**PrimaryLine** | Pointer to **string** |  | [optional] 
**SecondaryLine** | Pointer to **string** | The secondary delivery line of the address. This field is typically empty but may contain information if &#x60;primary_line&#x60; is too long.  | [optional] 
**LastLine** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Coverage** | Pointer to **string** |  | [optional] 
**Deliverability** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Components** | Pointer to [**IntlComponents**](IntlComponents.md) |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "intl_verification"]
**Error** | Pointer to [**BulkError**](BulkError.md) |  | [optional] 

## Methods

### NewIntlVerificationOrError

`func NewIntlVerificationOrError() *IntlVerificationOrError`

NewIntlVerificationOrError instantiates a new IntlVerificationOrError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntlVerificationOrErrorWithDefaults

`func NewIntlVerificationOrErrorWithDefaults() *IntlVerificationOrError`

NewIntlVerificationOrErrorWithDefaults instantiates a new IntlVerificationOrError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntlVerificationOrError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntlVerificationOrError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntlVerificationOrError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntlVerificationOrError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipient

`func (o *IntlVerificationOrError) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *IntlVerificationOrError) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *IntlVerificationOrError) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *IntlVerificationOrError) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *IntlVerificationOrError) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *IntlVerificationOrError) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetPrimaryLine

`func (o *IntlVerificationOrError) GetPrimaryLine() string`

GetPrimaryLine returns the PrimaryLine field if non-nil, zero value otherwise.

### GetPrimaryLineOk

`func (o *IntlVerificationOrError) GetPrimaryLineOk() (*string, bool)`

GetPrimaryLineOk returns a tuple with the PrimaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLine

`func (o *IntlVerificationOrError) SetPrimaryLine(v string)`

SetPrimaryLine sets PrimaryLine field to given value.

### HasPrimaryLine

`func (o *IntlVerificationOrError) HasPrimaryLine() bool`

HasPrimaryLine returns a boolean if a field has been set.

### GetSecondaryLine

`func (o *IntlVerificationOrError) GetSecondaryLine() string`

GetSecondaryLine returns the SecondaryLine field if non-nil, zero value otherwise.

### GetSecondaryLineOk

`func (o *IntlVerificationOrError) GetSecondaryLineOk() (*string, bool)`

GetSecondaryLineOk returns a tuple with the SecondaryLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLine

`func (o *IntlVerificationOrError) SetSecondaryLine(v string)`

SetSecondaryLine sets SecondaryLine field to given value.

### HasSecondaryLine

`func (o *IntlVerificationOrError) HasSecondaryLine() bool`

HasSecondaryLine returns a boolean if a field has been set.

### GetLastLine

`func (o *IntlVerificationOrError) GetLastLine() string`

GetLastLine returns the LastLine field if non-nil, zero value otherwise.

### GetLastLineOk

`func (o *IntlVerificationOrError) GetLastLineOk() (*string, bool)`

GetLastLineOk returns a tuple with the LastLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLine

`func (o *IntlVerificationOrError) SetLastLine(v string)`

SetLastLine sets LastLine field to given value.

### HasLastLine

`func (o *IntlVerificationOrError) HasLastLine() bool`

HasLastLine returns a boolean if a field has been set.

### GetCountry

`func (o *IntlVerificationOrError) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IntlVerificationOrError) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IntlVerificationOrError) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *IntlVerificationOrError) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCoverage

`func (o *IntlVerificationOrError) GetCoverage() string`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *IntlVerificationOrError) GetCoverageOk() (*string, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *IntlVerificationOrError) SetCoverage(v string)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *IntlVerificationOrError) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetDeliverability

`func (o *IntlVerificationOrError) GetDeliverability() string`

GetDeliverability returns the Deliverability field if non-nil, zero value otherwise.

### GetDeliverabilityOk

`func (o *IntlVerificationOrError) GetDeliverabilityOk() (*string, bool)`

GetDeliverabilityOk returns a tuple with the Deliverability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverability

`func (o *IntlVerificationOrError) SetDeliverability(v string)`

SetDeliverability sets Deliverability field to given value.

### HasDeliverability

`func (o *IntlVerificationOrError) HasDeliverability() bool`

HasDeliverability returns a boolean if a field has been set.

### GetStatus

`func (o *IntlVerificationOrError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntlVerificationOrError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntlVerificationOrError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntlVerificationOrError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComponents

`func (o *IntlVerificationOrError) GetComponents() IntlComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *IntlVerificationOrError) GetComponentsOk() (*IntlComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *IntlVerificationOrError) SetComponents(v IntlComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *IntlVerificationOrError) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetObject

`func (o *IntlVerificationOrError) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IntlVerificationOrError) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IntlVerificationOrError) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *IntlVerificationOrError) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetError

`func (o *IntlVerificationOrError) GetError() BulkError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IntlVerificationOrError) GetErrorOk() (*BulkError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IntlVerificationOrError) SetError(v BulkError)`

SetError sets Error field to given value.

### HasError

`func (o *IntlVerificationOrError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


