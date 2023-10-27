# LocationAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **NullableFloat32** | A positive or negative decimal indicating the geographic latitude of the address. | 
**Longitude** | **NullableFloat32** | A positive or negative decimal indicating the geographic longitude of the address. | 
**Distance** | **float32** | The distance from the input location to this exact zip code in miles. | 

## Methods

### NewLocationAnalysis

`func NewLocationAnalysis(latitude NullableFloat32, longitude NullableFloat32, distance float32, ) *LocationAnalysis`

NewLocationAnalysis instantiates a new LocationAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAnalysisWithDefaults

`func NewLocationAnalysisWithDefaults() *LocationAnalysis`

NewLocationAnalysisWithDefaults instantiates a new LocationAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *LocationAnalysis) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationAnalysis) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationAnalysis) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### SetLatitudeNil

`func (o *LocationAnalysis) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *LocationAnalysis) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *LocationAnalysis) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationAnalysis) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationAnalysis) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### SetLongitudeNil

`func (o *LocationAnalysis) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *LocationAnalysis) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetDistance

`func (o *LocationAnalysis) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *LocationAnalysis) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *LocationAnalysis) SetDistance(v float32)`

SetDistance sets Distance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


