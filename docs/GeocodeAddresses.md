# GeocodeAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**GeocodeComponents**](GeocodeComponents.md) |  | [optional] 
**LocationAnalysis** | Pointer to [**LocationAnalysis**](LocationAnalysis.md) |  | [optional] 

## Methods

### NewGeocodeAddresses

`func NewGeocodeAddresses() *GeocodeAddresses`

NewGeocodeAddresses instantiates a new GeocodeAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeocodeAddressesWithDefaults

`func NewGeocodeAddressesWithDefaults() *GeocodeAddresses`

NewGeocodeAddressesWithDefaults instantiates a new GeocodeAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *GeocodeAddresses) GetComponents() GeocodeComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GeocodeAddresses) GetComponentsOk() (*GeocodeComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GeocodeAddresses) SetComponents(v GeocodeComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *GeocodeAddresses) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetLocationAnalysis

`func (o *GeocodeAddresses) GetLocationAnalysis() LocationAnalysis`

GetLocationAnalysis returns the LocationAnalysis field if non-nil, zero value otherwise.

### GetLocationAnalysisOk

`func (o *GeocodeAddresses) GetLocationAnalysisOk() (*LocationAnalysis, bool)`

GetLocationAnalysisOk returns a tuple with the LocationAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAnalysis

`func (o *GeocodeAddresses) SetLocationAnalysis(v LocationAnalysis)`

SetLocationAnalysis sets LocationAnalysis field to given value.

### HasLocationAnalysis

`func (o *GeocodeAddresses) HasLocationAnalysis() bool`

HasLocationAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


