# BillingGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the billing group. | [optional] 
**Name** | **string** | Name of the billing group. | 
**Id** | **string** | Unique identifier prefixed with &#x60;bg_&#x60;. | 
**DateCreated** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was created. | 
**DateModified** | **time.Time** | A timestamp in ISO 8601 format of the date the resource was last modified. | 
**Object** | **string** | Value is resource type. | [default to "billing_group"]

## Methods

### NewBillingGroup

`func NewBillingGroup(name string, id string, dateCreated time.Time, dateModified time.Time, object string, ) *BillingGroup`

NewBillingGroup instantiates a new BillingGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingGroupWithDefaults

`func NewBillingGroupWithDefaults() *BillingGroup`

NewBillingGroupWithDefaults instantiates a new BillingGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BillingGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *BillingGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingGroup) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *BillingGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingGroup) SetId(v string)`

SetId sets Id field to given value.


### GetDateCreated

`func (o *BillingGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BillingGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BillingGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateModified

`func (o *BillingGroup) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *BillingGroup) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *BillingGroup) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.


### GetObject

`func (o *BillingGroup) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BillingGroup) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BillingGroup) SetObject(v string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


