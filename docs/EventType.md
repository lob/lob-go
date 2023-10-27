# EventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EnabledForTest** | **bool** | Value is &#x60;true&#x60; if the event type is enabled in both the test and live environments. | 
**Resource** | **string** |  | 
**Object** | **string** | Value is resource type. | [default to "event_type"]

## Methods

### NewEventType

`func NewEventType(id string, enabledForTest bool, resource string, object string, ) *EventType`

NewEventType instantiates a new EventType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTypeWithDefaults

`func NewEventTypeWithDefaults() *EventType`

NewEventTypeWithDefaults instantiates a new EventType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventType) SetId(v string)`

SetId sets Id field to given value.


### GetEnabledForTest

`func (o *EventType) GetEnabledForTest() bool`

GetEnabledForTest returns the EnabledForTest field if non-nil, zero value otherwise.

### GetEnabledForTestOk

`func (o *EventType) GetEnabledForTestOk() (*bool, bool)`

GetEnabledForTestOk returns a tuple with the EnabledForTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledForTest

`func (o *EventType) SetEnabledForTest(v bool)`

SetEnabledForTest sets EnabledForTest field to given value.


### GetResource

`func (o *EventType) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EventType) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EventType) SetResource(v string)`

SetResource sets Resource field to given value.


### GetObject

`func (o *EventType) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EventType) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EventType) SetObject(v string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


