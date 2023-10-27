# UsAutocompletions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;us_auto_&#x60;. | [optional] 
**Suggestions** | Pointer to [**[]Suggestions**](Suggestions.md) | An array of objects representing suggested addresses.  | [optional] 
**Object** | Pointer to **string** | Value is resource type. | [optional] [default to "us_autocompletion"]

## Methods

### NewUsAutocompletions

`func NewUsAutocompletions() *UsAutocompletions`

NewUsAutocompletions instantiates a new UsAutocompletions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsAutocompletionsWithDefaults

`func NewUsAutocompletionsWithDefaults() *UsAutocompletions`

NewUsAutocompletionsWithDefaults instantiates a new UsAutocompletions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsAutocompletions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsAutocompletions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsAutocompletions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsAutocompletions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSuggestions

`func (o *UsAutocompletions) GetSuggestions() []Suggestions`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *UsAutocompletions) GetSuggestionsOk() (*[]Suggestions, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *UsAutocompletions) SetSuggestions(v []Suggestions)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *UsAutocompletions) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### GetObject

`func (o *UsAutocompletions) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UsAutocompletions) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UsAutocompletions) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *UsAutocompletions) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


