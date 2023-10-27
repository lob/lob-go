# Zip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZipCode** | Pointer to **string** | A 5-digit ZIP code. | [optional] 
**Id** | **string** | Unique identifier prefixed with &#x60;us_zip_&#x60;. | 
**Cities** | [**[]ZipLookupCity**](ZipLookupCity.md) | An array of city objects containing valid cities for the &#x60;zip_code&#x60;. Multiple cities will be returned if more than one city is associated with the input ZIP code.  | 
**ZipCodeType** | [**ZipCodeType**](ZipCodeType.md) |  | 
**Object** | **string** |  | [default to "us_zip_lookup"]

## Methods

### NewZip

`func NewZip(id string, cities []ZipLookupCity, zipCodeType ZipCodeType, object string, ) *Zip`

NewZip instantiates a new Zip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZipWithDefaults

`func NewZipWithDefaults() *Zip`

NewZipWithDefaults instantiates a new Zip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZipCode

`func (o *Zip) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Zip) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Zip) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *Zip) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetId

`func (o *Zip) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Zip) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Zip) SetId(v string)`

SetId sets Id field to given value.


### GetCities

`func (o *Zip) GetCities() []ZipLookupCity`

GetCities returns the Cities field if non-nil, zero value otherwise.

### GetCitiesOk

`func (o *Zip) GetCitiesOk() (*[]ZipLookupCity, bool)`

GetCitiesOk returns a tuple with the Cities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCities

`func (o *Zip) SetCities(v []ZipLookupCity)`

SetCities sets Cities field to given value.


### GetZipCodeType

`func (o *Zip) GetZipCodeType() ZipCodeType`

GetZipCodeType returns the ZipCodeType field if non-nil, zero value otherwise.

### GetZipCodeTypeOk

`func (o *Zip) GetZipCodeTypeOk() (*ZipCodeType, bool)`

GetZipCodeTypeOk returns a tuple with the ZipCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeType

`func (o *Zip) SetZipCodeType(v ZipCodeType)`

SetZipCodeType sets ZipCodeType field to given value.


### GetObject

`func (o *Zip) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Zip) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Zip) SetObject(v string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


