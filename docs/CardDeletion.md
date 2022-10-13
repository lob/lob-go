# CardDeletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;card_&#x60;. | [optional] 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Object** | Pointer to **string** | Value is type of resource. | [optional] [default to "card_deleted"]

## Methods

### NewCardDeletion

`func NewCardDeletion() *CardDeletion`

NewCardDeletion instantiates a new CardDeletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDeletionWithDefaults

`func NewCardDeletionWithDefaults() *CardDeletion`

NewCardDeletionWithDefaults instantiates a new CardDeletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardDeletion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardDeletion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardDeletion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardDeletion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *CardDeletion) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CardDeletion) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CardDeletion) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CardDeletion) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetObject

`func (o *CardDeletion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CardDeletion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CardDeletion) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CardDeletion) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


