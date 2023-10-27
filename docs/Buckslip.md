# Buckslip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;bck_&#x60;. | 
**AutoReorder** | **bool** | True if the buckslips should be auto-reordered. | [default to false]
**ReorderQuantity** | **NullableInt32** | The number of buckslips to be reordered. | 
**ThresholdAmount** | **int32** | The threshold amount of the buckslip | [default to 0]
**Url** | **string** | The signed link for the buckslip. | 
**RawUrl** | **string** | The raw URL of the buckslip. | 
**FrontOriginalUrl** | **string** | The original URL of the front template. | 
**BackOriginalUrl** | **string** | The original URL of the back template. | 
**Thumbnails** | [**[]Thumbnail**](Thumbnail.md) |  | 
**AvailableQuantity** | **float32** | The available quantity of buckslips. | [default to 0]
**AllocatedQuantity** | **float32** | The allocated quantity of buckslips. | [default to 0]
**OnhandQuantity** | **float32** | The onhand quantity of buckslips. | [default to 0]
**PendingQuantity** | **float32** | The pending quantity of buckslips. | [default to 0]
**ProjectedQuantity** | **float32** | The sum of pending and onhand quantities of buckslips. | [default to 0]
**BuckslipOrders** | [**[]BuckslipOrder**](BuckslipOrder.md) | An array of buckslip orders that are associated with the buckslip. | 
**Stock** | **string** |  | 
**Weight** | **string** |  | 
**Finish** | **string** |  | 
**Status** | **string** |  | 
**Object** | **string** | object | [default to "buckslip"]
**Description** | **NullableString** | Description of the buckslip. | 
**Size** | Pointer to **string** | The size of the buckslip | [optional] [default to "8.75x3.75"]

## Methods

### NewBuckslip

`func NewBuckslip(id string, autoReorder bool, reorderQuantity NullableInt32, thresholdAmount int32, url string, rawUrl string, frontOriginalUrl string, backOriginalUrl string, thumbnails []Thumbnail, availableQuantity float32, allocatedQuantity float32, onhandQuantity float32, pendingQuantity float32, projectedQuantity float32, buckslipOrders []BuckslipOrder, stock string, weight string, finish string, status string, object string, description NullableString, ) *Buckslip`

NewBuckslip instantiates a new Buckslip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuckslipWithDefaults

`func NewBuckslipWithDefaults() *Buckslip`

NewBuckslipWithDefaults instantiates a new Buckslip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Buckslip) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Buckslip) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Buckslip) SetId(v string)`

SetId sets Id field to given value.


### GetAutoReorder

`func (o *Buckslip) GetAutoReorder() bool`

GetAutoReorder returns the AutoReorder field if non-nil, zero value otherwise.

### GetAutoReorderOk

`func (o *Buckslip) GetAutoReorderOk() (*bool, bool)`

GetAutoReorderOk returns a tuple with the AutoReorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoReorder

`func (o *Buckslip) SetAutoReorder(v bool)`

SetAutoReorder sets AutoReorder field to given value.


### GetReorderQuantity

`func (o *Buckslip) GetReorderQuantity() int32`

GetReorderQuantity returns the ReorderQuantity field if non-nil, zero value otherwise.

### GetReorderQuantityOk

`func (o *Buckslip) GetReorderQuantityOk() (*int32, bool)`

GetReorderQuantityOk returns a tuple with the ReorderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderQuantity

`func (o *Buckslip) SetReorderQuantity(v int32)`

SetReorderQuantity sets ReorderQuantity field to given value.


### SetReorderQuantityNil

`func (o *Buckslip) SetReorderQuantityNil(b bool)`

 SetReorderQuantityNil sets the value for ReorderQuantity to be an explicit nil

### UnsetReorderQuantity
`func (o *Buckslip) UnsetReorderQuantity()`

UnsetReorderQuantity ensures that no value is present for ReorderQuantity, not even an explicit nil
### GetThresholdAmount

`func (o *Buckslip) GetThresholdAmount() int32`

GetThresholdAmount returns the ThresholdAmount field if non-nil, zero value otherwise.

### GetThresholdAmountOk

`func (o *Buckslip) GetThresholdAmountOk() (*int32, bool)`

GetThresholdAmountOk returns a tuple with the ThresholdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdAmount

`func (o *Buckslip) SetThresholdAmount(v int32)`

SetThresholdAmount sets ThresholdAmount field to given value.


### GetUrl

`func (o *Buckslip) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Buckslip) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Buckslip) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRawUrl

`func (o *Buckslip) GetRawUrl() string`

GetRawUrl returns the RawUrl field if non-nil, zero value otherwise.

### GetRawUrlOk

`func (o *Buckslip) GetRawUrlOk() (*string, bool)`

GetRawUrlOk returns a tuple with the RawUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawUrl

`func (o *Buckslip) SetRawUrl(v string)`

SetRawUrl sets RawUrl field to given value.


### GetFrontOriginalUrl

`func (o *Buckslip) GetFrontOriginalUrl() string`

GetFrontOriginalUrl returns the FrontOriginalUrl field if non-nil, zero value otherwise.

### GetFrontOriginalUrlOk

`func (o *Buckslip) GetFrontOriginalUrlOk() (*string, bool)`

GetFrontOriginalUrlOk returns a tuple with the FrontOriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontOriginalUrl

`func (o *Buckslip) SetFrontOriginalUrl(v string)`

SetFrontOriginalUrl sets FrontOriginalUrl field to given value.


### GetBackOriginalUrl

`func (o *Buckslip) GetBackOriginalUrl() string`

GetBackOriginalUrl returns the BackOriginalUrl field if non-nil, zero value otherwise.

### GetBackOriginalUrlOk

`func (o *Buckslip) GetBackOriginalUrlOk() (*string, bool)`

GetBackOriginalUrlOk returns a tuple with the BackOriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackOriginalUrl

`func (o *Buckslip) SetBackOriginalUrl(v string)`

SetBackOriginalUrl sets BackOriginalUrl field to given value.


### GetThumbnails

`func (o *Buckslip) GetThumbnails() []Thumbnail`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *Buckslip) GetThumbnailsOk() (*[]Thumbnail, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *Buckslip) SetThumbnails(v []Thumbnail)`

SetThumbnails sets Thumbnails field to given value.


### GetAvailableQuantity

`func (o *Buckslip) GetAvailableQuantity() float32`

GetAvailableQuantity returns the AvailableQuantity field if non-nil, zero value otherwise.

### GetAvailableQuantityOk

`func (o *Buckslip) GetAvailableQuantityOk() (*float32, bool)`

GetAvailableQuantityOk returns a tuple with the AvailableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableQuantity

`func (o *Buckslip) SetAvailableQuantity(v float32)`

SetAvailableQuantity sets AvailableQuantity field to given value.


### GetAllocatedQuantity

`func (o *Buckslip) GetAllocatedQuantity() float32`

GetAllocatedQuantity returns the AllocatedQuantity field if non-nil, zero value otherwise.

### GetAllocatedQuantityOk

`func (o *Buckslip) GetAllocatedQuantityOk() (*float32, bool)`

GetAllocatedQuantityOk returns a tuple with the AllocatedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedQuantity

`func (o *Buckslip) SetAllocatedQuantity(v float32)`

SetAllocatedQuantity sets AllocatedQuantity field to given value.


### GetOnhandQuantity

`func (o *Buckslip) GetOnhandQuantity() float32`

GetOnhandQuantity returns the OnhandQuantity field if non-nil, zero value otherwise.

### GetOnhandQuantityOk

`func (o *Buckslip) GetOnhandQuantityOk() (*float32, bool)`

GetOnhandQuantityOk returns a tuple with the OnhandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnhandQuantity

`func (o *Buckslip) SetOnhandQuantity(v float32)`

SetOnhandQuantity sets OnhandQuantity field to given value.


### GetPendingQuantity

`func (o *Buckslip) GetPendingQuantity() float32`

GetPendingQuantity returns the PendingQuantity field if non-nil, zero value otherwise.

### GetPendingQuantityOk

`func (o *Buckslip) GetPendingQuantityOk() (*float32, bool)`

GetPendingQuantityOk returns a tuple with the PendingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingQuantity

`func (o *Buckslip) SetPendingQuantity(v float32)`

SetPendingQuantity sets PendingQuantity field to given value.


### GetProjectedQuantity

`func (o *Buckslip) GetProjectedQuantity() float32`

GetProjectedQuantity returns the ProjectedQuantity field if non-nil, zero value otherwise.

### GetProjectedQuantityOk

`func (o *Buckslip) GetProjectedQuantityOk() (*float32, bool)`

GetProjectedQuantityOk returns a tuple with the ProjectedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedQuantity

`func (o *Buckslip) SetProjectedQuantity(v float32)`

SetProjectedQuantity sets ProjectedQuantity field to given value.


### GetBuckslipOrders

`func (o *Buckslip) GetBuckslipOrders() []BuckslipOrder`

GetBuckslipOrders returns the BuckslipOrders field if non-nil, zero value otherwise.

### GetBuckslipOrdersOk

`func (o *Buckslip) GetBuckslipOrdersOk() (*[]BuckslipOrder, bool)`

GetBuckslipOrdersOk returns a tuple with the BuckslipOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckslipOrders

`func (o *Buckslip) SetBuckslipOrders(v []BuckslipOrder)`

SetBuckslipOrders sets BuckslipOrders field to given value.


### GetStock

`func (o *Buckslip) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *Buckslip) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *Buckslip) SetStock(v string)`

SetStock sets Stock field to given value.


### GetWeight

`func (o *Buckslip) GetWeight() string`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Buckslip) GetWeightOk() (*string, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Buckslip) SetWeight(v string)`

SetWeight sets Weight field to given value.


### GetFinish

`func (o *Buckslip) GetFinish() string`

GetFinish returns the Finish field if non-nil, zero value otherwise.

### GetFinishOk

`func (o *Buckslip) GetFinishOk() (*string, bool)`

GetFinishOk returns a tuple with the Finish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinish

`func (o *Buckslip) SetFinish(v string)`

SetFinish sets Finish field to given value.


### GetStatus

`func (o *Buckslip) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Buckslip) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Buckslip) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetObject

`func (o *Buckslip) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Buckslip) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Buckslip) SetObject(v string)`

SetObject sets Object field to given value.


### GetDescription

`func (o *Buckslip) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Buckslip) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Buckslip) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Buckslip) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Buckslip) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSize

`func (o *Buckslip) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Buckslip) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Buckslip) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Buckslip) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


