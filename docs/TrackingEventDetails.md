# TrackingEventDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | Find the full table [here](#tag/Tracking-Events). A detailed substatus about the event: * &#x60;package_accepted&#x60; - Package has been accepted into the carrier network for delivery. * &#x60;package_arrived&#x60; - Package has arrived at an intermediate location in the carrier network. * &#x60;package_departed&#x60; - Package has departed from an intermediate location in the carrier network. * &#x60;package_processing&#x60; - Package is processing at an intermediate location in the carrier network. * &#x60;package_processed&#x60; - Package has been processed at an intermediate location. * &#x60;package_in_local_area&#x60; - Package is at a location near the end destination. * &#x60;delivery_scheduled&#x60; - Package is scheduled for delivery. * &#x60;out_for_delivery&#x60; - Package is out for delivery. * &#x60;pickup_available&#x60; - Package is available for pickup at carrier location. * &#x60;delivered&#x60; - Package has been delivered. * &#x60;package_forwarded&#x60; - Package has been forwarded. * &#x60;returned_to_sender&#x60; - Package is to be returned to sender. * &#x60;address_issue&#x60; - Address information is incorrect. Contact carrier to ensure delivery. * &#x60;contact_carrier&#x60; - Contact the carrier for more information. * &#x60;delayed&#x60; - Delivery of package is delayed. * &#x60;delivery_attempted&#x60; - Delivery of package has been attempted. Contact carrier to ensure delivery. * &#x60;delivery_rescheduled&#x60; - Delivery of package has been rescheduled. * &#x60;location_inaccessible&#x60; - Delivery location inaccessible to carrier. Contact carrier to ensure delivery. * &#x60;notice_left&#x60; - Carrier left notice during attempted delivery. Follow carrier instructions on notice. * &#x60;package_damaged&#x60; - Package has been damaged. Contact carrier for more details. * &#x60;package_disposed&#x60; - Package has been disposed. * &#x60;package_held&#x60; - Package held at carrier location. Contact carrier for more details. * &#x60;package_lost&#x60; - Package has been lost. Contact carrier for more details. * &#x60;package_unclaimed&#x60; - Package is unclaimed. * &#x60;package_undeliverable&#x60; - Package is not able to be delivered. * &#x60;reschedule_delivery&#x60; - Contact carrier to reschedule delivery. * &#x60;other&#x60; - Unrecognized carrier status.  | 
**Description** | **string** | The description as listed in the description for event. | 
**Notes** | Pointer to **string** | Event-specific notes from USPS about the tracking event. | [optional] 
**ActionRequired** | **bool** | &#x60;true&#x60; if action is required by the end recipient, &#x60;false&#x60; otherwise.  | 

## Methods

### NewTrackingEventDetails

`func NewTrackingEventDetails(event string, description string, actionRequired bool, ) *TrackingEventDetails`

NewTrackingEventDetails instantiates a new TrackingEventDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingEventDetailsWithDefaults

`func NewTrackingEventDetailsWithDefaults() *TrackingEventDetails`

NewTrackingEventDetailsWithDefaults instantiates a new TrackingEventDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *TrackingEventDetails) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TrackingEventDetails) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TrackingEventDetails) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetDescription

`func (o *TrackingEventDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrackingEventDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrackingEventDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotes

`func (o *TrackingEventDetails) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TrackingEventDetails) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TrackingEventDetails) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TrackingEventDetails) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetActionRequired

`func (o *TrackingEventDetails) GetActionRequired() bool`

GetActionRequired returns the ActionRequired field if non-nil, zero value otherwise.

### GetActionRequiredOk

`func (o *TrackingEventDetails) GetActionRequiredOk() (*bool, bool)`

GetActionRequiredOk returns a tuple with the ActionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRequired

`func (o *TrackingEventDetails) SetActionRequired(v bool)`

SetActionRequired sets ActionRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


