# IntlAutocompletions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;intl_auto_&#x60;. | [optional] 
**Suggestions** | Pointer to [**[]IntlSuggestions**](IntlSuggestions.md) | An array of objects representing suggested addresses.  | [optional] 

## Methods

### NewIntlAutocompletions

`func NewIntlAutocompletions() *IntlAutocompletions`

NewIntlAutocompletions instantiates a new IntlAutocompletions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntlAutocompletionsWithDefaults

`func NewIntlAutocompletionsWithDefaults() *IntlAutocompletions`

NewIntlAutocompletionsWithDefaults instantiates a new IntlAutocompletions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntlAutocompletions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntlAutocompletions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntlAutocompletions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntlAutocompletions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSuggestions

`func (o *IntlAutocompletions) GetSuggestions() []IntlSuggestions`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *IntlAutocompletions) GetSuggestionsOk() (*[]IntlSuggestions, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *IntlAutocompletions) SetSuggestions(v []IntlSuggestions)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *IntlAutocompletions) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


