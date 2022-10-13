# ZipLookupCity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** |  | 
**State** | **string** | The [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) two letter code for the state.  | 
**County** | **string** | County name of the address city. | 
**CountyFips** | **string** | A 5-digit [FIPS county code](https://en.wikipedia.org/wiki/FIPS_county_code) which uniquely identifies &#x60;components[county]&#x60;. It consists of a 2-digit state code and a 3-digit county code.  | 
**Preferred** | **bool** | Indicates whether or not the city is the [USPS default city](https://en.wikipedia.org/wiki/ZIP_Code#ZIP_Codes_and_previous_zoning_lines) (preferred city) of a ZIP code. There is only one preferred city per ZIP code, which will always be in position 0 in the array of cities.  | 

## Methods

### NewZipLookupCity

`func NewZipLookupCity(city string, state string, county string, countyFips string, preferred bool, ) *ZipLookupCity`

NewZipLookupCity instantiates a new ZipLookupCity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZipLookupCityWithDefaults

`func NewZipLookupCityWithDefaults() *ZipLookupCity`

NewZipLookupCityWithDefaults instantiates a new ZipLookupCity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *ZipLookupCity) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ZipLookupCity) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ZipLookupCity) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *ZipLookupCity) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ZipLookupCity) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ZipLookupCity) SetState(v string)`

SetState sets State field to given value.


### GetCounty

`func (o *ZipLookupCity) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *ZipLookupCity) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *ZipLookupCity) SetCounty(v string)`

SetCounty sets County field to given value.


### GetCountyFips

`func (o *ZipLookupCity) GetCountyFips() string`

GetCountyFips returns the CountyFips field if non-nil, zero value otherwise.

### GetCountyFipsOk

`func (o *ZipLookupCity) GetCountyFipsOk() (*string, bool)`

GetCountyFipsOk returns a tuple with the CountyFips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyFips

`func (o *ZipLookupCity) SetCountyFips(v string)`

SetCountyFips sets CountyFips field to given value.


### GetPreferred

`func (o *ZipLookupCity) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *ZipLookupCity) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *ZipLookupCity) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


