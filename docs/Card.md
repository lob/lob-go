# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;card_&#x60;. | 
**Url** | **string** | The signed link for the card. | 
**AutoReorder** | **bool** | True if the cards should be auto-reordered. | [default to false]
**ReorderQuantity** | Pointer to **NullableInt32** | The number of cards to be reordered. Only present when auto_reorder is True. | [optional] 
**RawUrl** | Pointer to **string** | The raw URL of the card. | [optional] 
**FrontOriginalUrl** | Pointer to **string** | The original URL of the front template. | [optional] 
**BackOriginalUrl** | Pointer to **string** | The original URL of the back template. | [optional] 
**Thumbnails** | [**[]Thumbnail**](Thumbnail.md) |  | 
**AvailableQuantity** | **int32** | The available quantity of cards. | [default to 0]
**PendingQuantity** | **int32** | The pending quantity of cards. | [default to 0]
**Status** | Pointer to **string** |  | [optional] 
**Orientation** | Pointer to **string** | The orientation of the card. | [optional] [default to "horizontal"]
**ThresholdAmount** | Pointer to **int32** | The threshold amount of the card | [optional] [default to 0]
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Object** | **string** | object | [default to "card"]
**Description** | Pointer to **NullableString** | Description of the card. | [optional] 
**Size** | **string** | The size of the card | [default to "2.125x3.375"]

## Methods

### NewCard

`func NewCard(id string, url string, autoReorder bool, thumbnails []Thumbnail, availableQuantity int32, pendingQuantity int32, dateCreated time.Time, dateModified time.Time, object string, size string, ) *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Card) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Card) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Card) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *Card) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Card) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Card) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAutoReorder

`func (o *Card) GetAutoReorder() bool`

GetAutoReorder returns the AutoReorder field if non-nil, zero value otherwise.

### GetAutoReorderOk

`func (o *Card) GetAutoReorderOk() (*bool, bool)`

GetAutoReorderOk returns a tuple with the AutoReorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoReorder

`func (o *Card) SetAutoReorder(v bool)`

SetAutoReorder sets AutoReorder field to given value.


### GetReorderQuantity

`func (o *Card) GetReorderQuantity() int32`

GetReorderQuantity returns the ReorderQuantity field if non-nil, zero value otherwise.

### GetReorderQuantityOk

`func (o *Card) GetReorderQuantityOk() (*int32, bool)`

GetReorderQuantityOk returns a tuple with the ReorderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderQuantity

`func (o *Card) SetReorderQuantity(v int32)`

SetReorderQuantity sets ReorderQuantity field to given value.

### HasReorderQuantity

`func (o *Card) HasReorderQuantity() bool`

HasReorderQuantity returns a boolean if a field has been set.

### SetReorderQuantityNil

`func (o *Card) SetReorderQuantityNil(b bool)`

 SetReorderQuantityNil sets the value for ReorderQuantity to be an explicit nil

### UnsetReorderQuantity
`func (o *Card) UnsetReorderQuantity()`

UnsetReorderQuantity ensures that no value is present for ReorderQuantity, not even an explicit nil
### GetRawUrl

`func (o *Card) GetRawUrl() string`

GetRawUrl returns the RawUrl field if non-nil, zero value otherwise.

### GetRawUrlOk

`func (o *Card) GetRawUrlOk() (*string, bool)`

GetRawUrlOk returns a tuple with the RawUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawUrl

`func (o *Card) SetRawUrl(v string)`

SetRawUrl sets RawUrl field to given value.

### HasRawUrl

`func (o *Card) HasRawUrl() bool`

HasRawUrl returns a boolean if a field has been set.

### GetFrontOriginalUrl

`func (o *Card) GetFrontOriginalUrl() string`

GetFrontOriginalUrl returns the FrontOriginalUrl field if non-nil, zero value otherwise.

### GetFrontOriginalUrlOk

`func (o *Card) GetFrontOriginalUrlOk() (*string, bool)`

GetFrontOriginalUrlOk returns a tuple with the FrontOriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontOriginalUrl

`func (o *Card) SetFrontOriginalUrl(v string)`

SetFrontOriginalUrl sets FrontOriginalUrl field to given value.

### HasFrontOriginalUrl

`func (o *Card) HasFrontOriginalUrl() bool`

HasFrontOriginalUrl returns a boolean if a field has been set.

### GetBackOriginalUrl

`func (o *Card) GetBackOriginalUrl() string`

GetBackOriginalUrl returns the BackOriginalUrl field if non-nil, zero value otherwise.

### GetBackOriginalUrlOk

`func (o *Card) GetBackOriginalUrlOk() (*string, bool)`

GetBackOriginalUrlOk returns a tuple with the BackOriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackOriginalUrl

`func (o *Card) SetBackOriginalUrl(v string)`

SetBackOriginalUrl sets BackOriginalUrl field to given value.

### HasBackOriginalUrl

`func (o *Card) HasBackOriginalUrl() bool`

HasBackOriginalUrl returns a boolean if a field has been set.

### GetThumbnails

`func (o *Card) GetThumbnails() []Thumbnail`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *Card) GetThumbnailsOk() (*[]Thumbnail, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *Card) SetThumbnails(v []Thumbnail)`

SetThumbnails sets Thumbnails field to given value.


### GetAvailableQuantity

`func (o *Card) GetAvailableQuantity() int32`

GetAvailableQuantity returns the AvailableQuantity field if non-nil, zero value otherwise.

### GetAvailableQuantityOk

`func (o *Card) GetAvailableQuantityOk() (*int32, bool)`

GetAvailableQuantityOk returns a tuple with the AvailableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableQuantity

`func (o *Card) SetAvailableQuantity(v int32)`

SetAvailableQuantity sets AvailableQuantity field to given value.


### GetPendingQuantity

`func (o *Card) GetPendingQuantity() int32`

GetPendingQuantity returns the PendingQuantity field if non-nil, zero value otherwise.

### GetPendingQuantityOk

`func (o *Card) GetPendingQuantityOk() (*int32, bool)`

GetPendingQuantityOk returns a tuple with the PendingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingQuantity

`func (o *Card) SetPendingQuantity(v int32)`

SetPendingQuantity sets PendingQuantity field to given value.


### GetStatus

`func (o *Card) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Card) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Card) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Card) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrientation

`func (o *Card) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *Card) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *Card) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *Card) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetThresholdAmount

`func (o *Card) GetThresholdAmount() int32`

GetThresholdAmount returns the ThresholdAmount field if non-nil, zero value otherwise.

### GetThresholdAmountOk

`func (o *Card) GetThresholdAmountOk() (*int32, bool)`

GetThresholdAmountOk returns a tuple with the ThresholdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdAmount

`func (o *Card) SetThresholdAmount(v int32)`

SetThresholdAmount sets ThresholdAmount field to given value.

### HasThresholdAmount

`func (o *Card) HasThresholdAmount() bool`

HasThresholdAmount returns a boolean if a field has been set.

### GetDateCreated

`func (o *Card) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Card) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Card) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *Card) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *Card) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *Card) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetDeleted

`func (o *Card) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Card) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Card) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Card) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetObject

`func (o *Card) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Card) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Card) SetObject(v string)`

SetObject sets Object field to given value.


### GetDescription

`func (o *Card) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Card) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Card) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Card) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Card) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Card) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSize

`func (o *Card) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Card) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Card) SetSize(v string)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


