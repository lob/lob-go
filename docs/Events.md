# Events

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier prefixed with &#x60;evt_&#x60;. | 
**Body** | **map[string]interface{}** | The body of the associated resource as they were at the time of the event, i.e. the [postcard object](https://docs.lob.com/#tag/Postcards/operation/postcard_retrieve), [the letter object](https://docs.lob.com/#tag/Letters/operation/letter_retrieve), [the check object](https://docs.lob.com/#tag/Checks/operation/check_retrieve), [the address object](https://docs.lob.com/#tag/Addresses/operation/address_retrieve), or [the bank account object](https://docs.lob.com/#tag/Bank-Accounts/operation/bank_account_retrieve). For &#x60;.deleted&#x60; events, the body matches the response for the corresponding delete endpoint for that resource (e.g. the [postcard cancel response](https://docs.lob.com/#tag/Postcards/operation/postcard_delete)). | 
**ReferenceId** | **string** | Unique identifier of the related resource for the event. | 
**EventType** | [**EventType**](EventType.md) |  | 
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**Object** | **string** | Value is resource type. | [default to "event"]

## Methods

### NewEvents

`func NewEvents(id string, body map[string]interface{}, referenceId string, eventType EventType, dateCreated time.Time, object string, ) *Events`

NewEvents instantiates a new Events object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsWithDefaults

`func NewEventsWithDefaults() *Events`

NewEventsWithDefaults instantiates a new Events object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Events) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Events) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Events) SetId(v string)`

SetId sets Id field to given value.


### GetBody

`func (o *Events) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Events) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Events) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.


### GetReferenceId

`func (o *Events) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Events) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Events) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetEventType

`func (o *Events) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Events) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Events) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetDateCreated

`func (o *Events) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Events) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Events) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetObject

`func (o *Events) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Events) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Events) SetObject(v string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


