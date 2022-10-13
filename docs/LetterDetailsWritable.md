# LetterDetailsWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressPlacement** | Pointer to **string** | Specifies the location of the address information that will show through the double-window envelope.  | [optional] [default to "top_first_page"]
**Cards** | **[]string** | A single-element array containing an existing card id in a string format. See [cards](#tag/Cards) for more information. | 
**Color** | **bool** | Set this key to &#x60;true&#x60; if you would like to print in color. Set to &#x60;false&#x60; if you would like to print in black and white. | 
**DoubleSided** | Pointer to **bool** | Set this attribute to &#x60;true&#x60; for double sided printing, or &#x60;false&#x60; for for single sided printing. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ExtraService** | Pointer to **string** | Add an extra service to your letter. | [optional] 
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to FIRST_CLASS]
**ReturnEnvelope** | Pointer to **bool** |  | [optional] [default to false]
**CustomEnvelope** | Pointer to **NullableString** | Accepts an envelope ID for any customized envelope with available inventory. | [optional] 

## Methods

### NewLetterDetailsWritable

`func NewLetterDetailsWritable(cards []string, color bool, ) *LetterDetailsWritable`

NewLetterDetailsWritable instantiates a new LetterDetailsWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLetterDetailsWritableWithDefaults

`func NewLetterDetailsWritableWithDefaults() *LetterDetailsWritable`

NewLetterDetailsWritableWithDefaults instantiates a new LetterDetailsWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressPlacement

`func (o *LetterDetailsWritable) GetAddressPlacement() string`

GetAddressPlacement returns the AddressPlacement field if non-nil, zero value otherwise.

### GetAddressPlacementOk

`func (o *LetterDetailsWritable) GetAddressPlacementOk() (*string, bool)`

GetAddressPlacementOk returns a tuple with the AddressPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPlacement

`func (o *LetterDetailsWritable) SetAddressPlacement(v string)`

SetAddressPlacement sets AddressPlacement field to given value.

### HasAddressPlacement

`func (o *LetterDetailsWritable) HasAddressPlacement() bool`

HasAddressPlacement returns a boolean if a field has been set.

### GetCards

`func (o *LetterDetailsWritable) GetCards() []string`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *LetterDetailsWritable) GetCardsOk() (*[]string, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *LetterDetailsWritable) SetCards(v []string)`

SetCards sets Cards field to given value.


### SetCardsNil

`func (o *LetterDetailsWritable) SetCardsNil(b bool)`

 SetCardsNil sets the value for Cards to be an explicit nil

### UnsetCards
`func (o *LetterDetailsWritable) UnsetCards()`

UnsetCards ensures that no value is present for Cards, not even an explicit nil
### GetColor

`func (o *LetterDetailsWritable) GetColor() bool`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LetterDetailsWritable) GetColorOk() (*bool, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LetterDetailsWritable) SetColor(v bool)`

SetColor sets Color field to given value.


### GetDoubleSided

`func (o *LetterDetailsWritable) GetDoubleSided() bool`

GetDoubleSided returns the DoubleSided field if non-nil, zero value otherwise.

### GetDoubleSidedOk

`func (o *LetterDetailsWritable) GetDoubleSidedOk() (*bool, bool)`

GetDoubleSidedOk returns a tuple with the DoubleSided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleSided

`func (o *LetterDetailsWritable) SetDoubleSided(v bool)`

SetDoubleSided sets DoubleSided field to given value.

### HasDoubleSided

`func (o *LetterDetailsWritable) HasDoubleSided() bool`

HasDoubleSided returns a boolean if a field has been set.

### GetExtraService

`func (o *LetterDetailsWritable) GetExtraService() string`

GetExtraService returns the ExtraService field if non-nil, zero value otherwise.

### GetExtraServiceOk

`func (o *LetterDetailsWritable) GetExtraServiceOk() (*string, bool)`

GetExtraServiceOk returns a tuple with the ExtraService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraService

`func (o *LetterDetailsWritable) SetExtraService(v string)`

SetExtraService sets ExtraService field to given value.

### HasExtraService

`func (o *LetterDetailsWritable) HasExtraService() bool`

HasExtraService returns a boolean if a field has been set.

### GetMailType

`func (o *LetterDetailsWritable) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *LetterDetailsWritable) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *LetterDetailsWritable) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *LetterDetailsWritable) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetReturnEnvelope

`func (o *LetterDetailsWritable) GetReturnEnvelope() bool`

GetReturnEnvelope returns the ReturnEnvelope field if non-nil, zero value otherwise.

### GetReturnEnvelopeOk

`func (o *LetterDetailsWritable) GetReturnEnvelopeOk() (*bool, bool)`

GetReturnEnvelopeOk returns a tuple with the ReturnEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnEnvelope

`func (o *LetterDetailsWritable) SetReturnEnvelope(v bool)`

SetReturnEnvelope sets ReturnEnvelope field to given value.

### HasReturnEnvelope

`func (o *LetterDetailsWritable) HasReturnEnvelope() bool`

HasReturnEnvelope returns a boolean if a field has been set.

### GetCustomEnvelope

`func (o *LetterDetailsWritable) GetCustomEnvelope() string`

GetCustomEnvelope returns the CustomEnvelope field if non-nil, zero value otherwise.

### GetCustomEnvelopeOk

`func (o *LetterDetailsWritable) GetCustomEnvelopeOk() (*string, bool)`

GetCustomEnvelopeOk returns a tuple with the CustomEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvelope

`func (o *LetterDetailsWritable) SetCustomEnvelope(v string)`

SetCustomEnvelope sets CustomEnvelope field to given value.

### HasCustomEnvelope

`func (o *LetterDetailsWritable) HasCustomEnvelope() bool`

HasCustomEnvelope returns a boolean if a field has been set.

### SetCustomEnvelopeNil

`func (o *LetterDetailsWritable) SetCustomEnvelopeNil(b bool)`

 SetCustomEnvelopeNil sets the value for CustomEnvelope to be an explicit nil

### UnsetCustomEnvelope
`func (o *LetterDetailsWritable) UnsetCustomEnvelope()`

UnsetCustomEnvelope ensures that no value is present for CustomEnvelope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


