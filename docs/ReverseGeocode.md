# ReverseGeocode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;us_reverse_geocode_&#x60;. | [optional] 
**Addresses** | Pointer to [**[]GeocodeAddresses**](GeocodeAddresses.md) | list of addresses  | [optional] 
**Object** | Pointer to **string** | Value is resource type. | [optional] [default to "us_reverse_geocode_lookup"]

## Methods

### NewReverseGeocode

`func NewReverseGeocode() *ReverseGeocode`

NewReverseGeocode instantiates a new ReverseGeocode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseGeocodeWithDefaults

`func NewReverseGeocodeWithDefaults() *ReverseGeocode`

NewReverseGeocodeWithDefaults instantiates a new ReverseGeocode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReverseGeocode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReverseGeocode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReverseGeocode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReverseGeocode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddresses

`func (o *ReverseGeocode) GetAddresses() []GeocodeAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ReverseGeocode) GetAddressesOk() (*[]GeocodeAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ReverseGeocode) SetAddresses(v []GeocodeAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ReverseGeocode) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetObject

`func (o *ReverseGeocode) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ReverseGeocode) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ReverseGeocode) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ReverseGeocode) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


