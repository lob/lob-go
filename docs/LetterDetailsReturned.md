# LetterDetailsReturned

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **bool** | Set this key to &#x60;true&#x60; if you would like to print in color, false for black and white. | 
**Cards** | **[]string** | A single-element array containing an existing card id in a string format. See [cards](#tag/Cards) for more information. | 
**AddressPlacement** | Pointer to **string** | Specifies the location of the address information that will show through the double-window envelope.  | [optional] [default to "top_first_page"]
**CustomEnvelope** | Pointer to [**NullableCustomEnvelopeReturned**](CustomEnvelopeReturned.md) |  | [optional] 
**DoubleSided** | Pointer to **bool** | Set this attribute to &#x60;true&#x60; for double sided printing,  &#x60;false&#x60; for for single sided printing. | [optional] [default to true]
**ExtraService** | Pointer to **string** | Add an extra service to your letter. | [optional] 
**MailType** | Pointer to [**MailType**](MailType.md) |  | [optional] [default to MAILTYPE_FIRST_CLASS]
**ReturnEnvelope** | Pointer to **interface{}** |  | [optional] 
**Bleed** | Pointer to **bool** | Allows for letter bleed. Enabled only with specific feature flags. | [optional] [default to false]
**FileOriginalUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLetterDetailsReturned

`func NewLetterDetailsReturned(color bool, cards []string, ) *LetterDetailsReturned`

NewLetterDetailsReturned instantiates a new LetterDetailsReturned object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLetterDetailsReturnedWithDefaults

`func NewLetterDetailsReturnedWithDefaults() *LetterDetailsReturned`

NewLetterDetailsReturnedWithDefaults instantiates a new LetterDetailsReturned object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *LetterDetailsReturned) GetColor() bool`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LetterDetailsReturned) GetColorOk() (*bool, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LetterDetailsReturned) SetColor(v bool)`

SetColor sets Color field to given value.


### GetCards

`func (o *LetterDetailsReturned) GetCards() []string`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *LetterDetailsReturned) GetCardsOk() (*[]string, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *LetterDetailsReturned) SetCards(v []string)`

SetCards sets Cards field to given value.


### SetCardsNil

`func (o *LetterDetailsReturned) SetCardsNil(b bool)`

 SetCardsNil sets the value for Cards to be an explicit nil

### UnsetCards
`func (o *LetterDetailsReturned) UnsetCards()`

UnsetCards ensures that no value is present for Cards, not even an explicit nil
### GetAddressPlacement

`func (o *LetterDetailsReturned) GetAddressPlacement() string`

GetAddressPlacement returns the AddressPlacement field if non-nil, zero value otherwise.

### GetAddressPlacementOk

`func (o *LetterDetailsReturned) GetAddressPlacementOk() (*string, bool)`

GetAddressPlacementOk returns a tuple with the AddressPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPlacement

`func (o *LetterDetailsReturned) SetAddressPlacement(v string)`

SetAddressPlacement sets AddressPlacement field to given value.

### HasAddressPlacement

`func (o *LetterDetailsReturned) HasAddressPlacement() bool`

HasAddressPlacement returns a boolean if a field has been set.

### GetCustomEnvelope

`func (o *LetterDetailsReturned) GetCustomEnvelope() CustomEnvelopeReturned`

GetCustomEnvelope returns the CustomEnvelope field if non-nil, zero value otherwise.

### GetCustomEnvelopeOk

`func (o *LetterDetailsReturned) GetCustomEnvelopeOk() (*CustomEnvelopeReturned, bool)`

GetCustomEnvelopeOk returns a tuple with the CustomEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvelope

`func (o *LetterDetailsReturned) SetCustomEnvelope(v CustomEnvelopeReturned)`

SetCustomEnvelope sets CustomEnvelope field to given value.

### HasCustomEnvelope

`func (o *LetterDetailsReturned) HasCustomEnvelope() bool`

HasCustomEnvelope returns a boolean if a field has been set.

### SetCustomEnvelopeNil

`func (o *LetterDetailsReturned) SetCustomEnvelopeNil(b bool)`

 SetCustomEnvelopeNil sets the value for CustomEnvelope to be an explicit nil

### UnsetCustomEnvelope
`func (o *LetterDetailsReturned) UnsetCustomEnvelope()`

UnsetCustomEnvelope ensures that no value is present for CustomEnvelope, not even an explicit nil
### GetDoubleSided

`func (o *LetterDetailsReturned) GetDoubleSided() bool`

GetDoubleSided returns the DoubleSided field if non-nil, zero value otherwise.

### GetDoubleSidedOk

`func (o *LetterDetailsReturned) GetDoubleSidedOk() (*bool, bool)`

GetDoubleSidedOk returns a tuple with the DoubleSided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleSided

`func (o *LetterDetailsReturned) SetDoubleSided(v bool)`

SetDoubleSided sets DoubleSided field to given value.

### HasDoubleSided

`func (o *LetterDetailsReturned) HasDoubleSided() bool`

HasDoubleSided returns a boolean if a field has been set.

### GetExtraService

`func (o *LetterDetailsReturned) GetExtraService() string`

GetExtraService returns the ExtraService field if non-nil, zero value otherwise.

### GetExtraServiceOk

`func (o *LetterDetailsReturned) GetExtraServiceOk() (*string, bool)`

GetExtraServiceOk returns a tuple with the ExtraService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraService

`func (o *LetterDetailsReturned) SetExtraService(v string)`

SetExtraService sets ExtraService field to given value.

### HasExtraService

`func (o *LetterDetailsReturned) HasExtraService() bool`

HasExtraService returns a boolean if a field has been set.

### GetMailType

`func (o *LetterDetailsReturned) GetMailType() MailType`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *LetterDetailsReturned) GetMailTypeOk() (*MailType, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *LetterDetailsReturned) SetMailType(v MailType)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *LetterDetailsReturned) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetReturnEnvelope

`func (o *LetterDetailsReturned) GetReturnEnvelope() interface{}`

GetReturnEnvelope returns the ReturnEnvelope field if non-nil, zero value otherwise.

### GetReturnEnvelopeOk

`func (o *LetterDetailsReturned) GetReturnEnvelopeOk() (*interface{}, bool)`

GetReturnEnvelopeOk returns a tuple with the ReturnEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnEnvelope

`func (o *LetterDetailsReturned) SetReturnEnvelope(v interface{})`

SetReturnEnvelope sets ReturnEnvelope field to given value.

### HasReturnEnvelope

`func (o *LetterDetailsReturned) HasReturnEnvelope() bool`

HasReturnEnvelope returns a boolean if a field has been set.

### SetReturnEnvelopeNil

`func (o *LetterDetailsReturned) SetReturnEnvelopeNil(b bool)`

 SetReturnEnvelopeNil sets the value for ReturnEnvelope to be an explicit nil

### UnsetReturnEnvelope
`func (o *LetterDetailsReturned) UnsetReturnEnvelope()`

UnsetReturnEnvelope ensures that no value is present for ReturnEnvelope, not even an explicit nil
### GetBleed

`func (o *LetterDetailsReturned) GetBleed() bool`

GetBleed returns the Bleed field if non-nil, zero value otherwise.

### GetBleedOk

`func (o *LetterDetailsReturned) GetBleedOk() (*bool, bool)`

GetBleedOk returns a tuple with the Bleed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleed

`func (o *LetterDetailsReturned) SetBleed(v bool)`

SetBleed sets Bleed field to given value.

### HasBleed

`func (o *LetterDetailsReturned) HasBleed() bool`

HasBleed returns a boolean if a field has been set.

### GetFileOriginalUrl

`func (o *LetterDetailsReturned) GetFileOriginalUrl() string`

GetFileOriginalUrl returns the FileOriginalUrl field if non-nil, zero value otherwise.

### GetFileOriginalUrlOk

`func (o *LetterDetailsReturned) GetFileOriginalUrlOk() (*string, bool)`

GetFileOriginalUrlOk returns a tuple with the FileOriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOriginalUrl

`func (o *LetterDetailsReturned) SetFileOriginalUrl(v string)`

SetFileOriginalUrl sets FileOriginalUrl field to given value.

### HasFileOriginalUrl

`func (o *LetterDetailsReturned) HasFileOriginalUrl() bool`

HasFileOriginalUrl returns a boolean if a field has been set.

### SetFileOriginalUrlNil

`func (o *LetterDetailsReturned) SetFileOriginalUrlNil(b bool)`

 SetFileOriginalUrlNil sets the value for FileOriginalUrl to be an explicit nil

### UnsetFileOriginalUrl
`func (o *LetterDetailsReturned) UnsetFileOriginalUrl()`

UnsetFileOriginalUrl ensures that no value is present for FileOriginalUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


