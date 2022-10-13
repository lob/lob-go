# TrackingEventNormal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | non-Certified postcards, self mailers, letters, and checks | 
**Name** | **string** | Name of tracking event (for normal postcards, self mailers, letters, and checks):    * &#x60;In Transit&#x60; - The mailpiece is being processed at the entry/origin facility.    * &#x60;In Local Area&#x60; - The mailpiece is being processed at the destination facility.    * &#x60;Processed for Delivery&#x60; - The mailpiece has been greenlit for     delivery at the recipient&#39;s nearest postal facility. The mailpiece     should reach the mailbox within 1 business day of this tracking     event.    * &#x60;Re-Routed&#x60; - The mailpiece is re-routed due to recipient change of     address, address errors, or USPS relabeling of barcode/ID tag     area.    * &#x60;Returned to Sender&#x60; - The mailpiece is being returned to sender due     to barcode, ID tag area, or address errors.    * &#x60;Mailed&#x60; - The mailpiece has been handed off to and accepted by USPS     and is en route. [More about     Mailed.](https://support.lob.com/hc/en-us/articles/360001724400-What-does-a-Mailed-tracking-event-mean-)     Note this data is only available in Enterprise editions of     Lob. [Contact Sales](https://lob.com/support/contact#contact) if     you want access to this feature.  [More about tracking](https://support.lob.com/hc/en-us/articles/115000097404-Can-I-track-my-mail-)  | 
**Details** | Pointer to **map[string]interface{}** | Will be &#x60;null&#x60; for &#x60;type&#x3D;normal&#x60; events | [optional] 
**Location** | Pointer to **NullableString** | The zip code in which the scan event occurred. Null for &#x60;Mailed&#x60; events.  | [optional] 
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;evnt_&#x60;. | [optional] 
**Time** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date USPS registered the event. | [optional] 
**DateCreated** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | [optional] 
**DateModified** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "tracking_event"]

## Methods

### NewTrackingEventNormal

`func NewTrackingEventNormal(type_ string, name string, ) *TrackingEventNormal`

NewTrackingEventNormal instantiates a new TrackingEventNormal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingEventNormalWithDefaults

`func NewTrackingEventNormalWithDefaults() *TrackingEventNormal`

NewTrackingEventNormalWithDefaults instantiates a new TrackingEventNormal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TrackingEventNormal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrackingEventNormal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrackingEventNormal) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TrackingEventNormal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackingEventNormal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackingEventNormal) SetName(v string)`

SetName sets Name field to given value.


### GetDetails

`func (o *TrackingEventNormal) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TrackingEventNormal) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TrackingEventNormal) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *TrackingEventNormal) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *TrackingEventNormal) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *TrackingEventNormal) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetLocation

`func (o *TrackingEventNormal) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TrackingEventNormal) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TrackingEventNormal) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TrackingEventNormal) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *TrackingEventNormal) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *TrackingEventNormal) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetId

`func (o *TrackingEventNormal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingEventNormal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingEventNormal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackingEventNormal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTime

`func (o *TrackingEventNormal) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TrackingEventNormal) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TrackingEventNormal) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *TrackingEventNormal) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetDateCreated

`func (o *TrackingEventNormal) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TrackingEventNormal) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TrackingEventNormal) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TrackingEventNormal) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateModified

`func (o *TrackingEventNormal) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *TrackingEventNormal) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *TrackingEventNormal) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.

### HasDateModified

`func (o *TrackingEventNormal) HasDateModified() bool`

HasDateModified returns a boolean if a field has been set.

### GetObject

`func (o *TrackingEventNormal) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TrackingEventNormal) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TrackingEventNormal) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *TrackingEventNormal) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


