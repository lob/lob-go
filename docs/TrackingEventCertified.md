# TrackingEventCertified

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | a Certified letter tracking event | 
**Name** | **string** | Name of tracking event for Certified letters. Letters sent with USPS Certified Mail are fully tracked by USPS, therefore their tracking events have an additional details object with more detailed information about the tracking event. Some certified tracking event names have multiple meanings, noted in the list here. See the description of the details object for the full set of combined certified tracking event name meanings.    * &#x60;Mailed&#x60; - Package has been accepted into the carrier network for delivery.    * &#x60;In Transit&#x60; - Maps to four distinct stages of transit.    * &#x60;In Local Area&#x60; - Package is at a location near the end destination.    * &#x60;Processed for Delivery&#x60; - Maps to two distinct stages of delivery.    * &#x60;Pickup Available&#x60; - Package is available for pickup at carrier location.    * &#x60;Delivered&#x60; - Package has been delivered.    * &#x60;Re-Routed&#x60; - Package has been forwarded.    * &#x60;Returned to Sender&#x60; - Package is to be returned to sender.    * &#x60;Issue&#x60; - Maps to (at least) 15 possible issues, some of which are actionable.  | 
**Details** | Pointer to [**TrackingEventDetails**](TrackingEventDetails.md) |  | [optional] 
**Location** | Pointer to **NullableString** | The zip code in which the event occurred if it exists, otherwise will be the name of a Regional Distribution Center if it exists, otherwise will be null.  | [optional] 
**Id** | **string** | Unique identifier prefixed with &#x60;evnt_&#x60;. | 
**Time** | Pointer to **time.Time** | A timestamp in ISO 8601 format of the date USPS registered the event. | [optional] 
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | 
**Object** | **string** |  | [default to "tracking_event"]

## Methods

### NewTrackingEventCertified

`func NewTrackingEventCertified(type_ string, name string, id string, dateCreated time.Time, dateModified time.Time, object string, ) *TrackingEventCertified`

NewTrackingEventCertified instantiates a new TrackingEventCertified object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingEventCertifiedWithDefaults

`func NewTrackingEventCertifiedWithDefaults() *TrackingEventCertified`

NewTrackingEventCertifiedWithDefaults instantiates a new TrackingEventCertified object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TrackingEventCertified) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrackingEventCertified) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrackingEventCertified) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TrackingEventCertified) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackingEventCertified) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackingEventCertified) SetName(v string)`

SetName sets Name field to given value.


### GetDetails

`func (o *TrackingEventCertified) GetDetails() TrackingEventDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TrackingEventCertified) GetDetailsOk() (*TrackingEventDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TrackingEventCertified) SetDetails(v TrackingEventDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *TrackingEventCertified) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetLocation

`func (o *TrackingEventCertified) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TrackingEventCertified) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TrackingEventCertified) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TrackingEventCertified) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *TrackingEventCertified) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *TrackingEventCertified) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetId

`func (o *TrackingEventCertified) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingEventCertified) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingEventCertified) SetId(v string)`

SetId sets Id field to given value.


### GetTime

`func (o *TrackingEventCertified) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TrackingEventCertified) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TrackingEventCertified) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *TrackingEventCertified) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetDateCreated

`func (o *TrackingEventCertified) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TrackingEventCertified) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TrackingEventCertified) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *TrackingEventCertified) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *TrackingEventCertified) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *TrackingEventCertified) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetObject

`func (o *TrackingEventCertified) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TrackingEventCertified) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TrackingEventCertified) SetObject(v string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


