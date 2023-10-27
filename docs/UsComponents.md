# UsComponents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryNumber** | **string** | The numeric or alphanumeric part of an address preceding the street name. Often the house, building, or PO Box number. | 
**StreetPredirection** | **string** | Geographic direction preceding a street name (&#x60;N&#x60;, &#x60;S&#x60;, &#x60;E&#x60;, &#x60;W&#x60;, &#x60;NE&#x60;, &#x60;SW&#x60;, &#x60;SE&#x60;, &#x60;NW&#x60;).  | 
**StreetName** | **string** | The name of the street. | 
**StreetSuffix** | **string** | The standard USPS abbreviation for the street suffix (&#x60;ST&#x60;, &#x60;AVE&#x60;, &#x60;BLVD&#x60;, etc).  | 
**StreetPostdirection** | **string** | Geographic direction following a street name (&#x60;N&#x60;, &#x60;S&#x60;, &#x60;E&#x60;, &#x60;W&#x60;, &#x60;NE&#x60;, &#x60;SW&#x60;, &#x60;SE&#x60;, &#x60;NW&#x60;).  | 
**SecondaryDesignator** | **string** | The standard USPS abbreviation describing the &#x60;components[secondary_number]&#x60; (&#x60;STE&#x60;, &#x60;APT&#x60;, &#x60;BLDG&#x60;, etc).  | 
**SecondaryNumber** | **string** | Number of the apartment/unit/etc.  | 
**PmbDesignator** | **string** | Designator of a [CMRA-authorized](https://en.wikipedia.org/wiki/Commercial_mail_receiving_agency) private mailbox.  | 
**PmbNumber** | **string** | Number of a [CMRA-authorized](https://en.wikipedia.org/wiki/Commercial_mail_receiving_agency) private mailbox.  | 
**ExtraSecondaryDesignator** | **string** | An extra (often unnecessary) secondary designator provided with the input address.  | 
**ExtraSecondaryNumber** | **string** | An extra (often unnecessary) secondary number provided with the input address.  | 
**City** | **string** |  | 
**State** | **string** | The [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) two letter code for the state.  | 
**ZipCode** | **string** | The 5-digit ZIP code | 
**ZipCodePlus4** | **string** |  | 
**ZipCodeType** | [**ZipCodeType**](ZipCodeType.md) |  | 
**DeliveryPointBarcode** | **string** | A 12-digit identifier that uniquely identifies a delivery point (location where mail can be sent and received). It consists of the 5-digit ZIP code (&#x60;zip_code&#x60;), 4-digit ZIP+4 add-on (&#x60;zip_code_plus_4&#x60;), 2-digit delivery point, and 1-digit delivery point check digit.  | 
**AddressType** | **string** | Uses USPS&#39;s [Residential Delivery Indicator (RDI)](https://www.usps.com/nationalpremieraccounts/rdi.htm) to identify whether an address is classified as residential or business. Possible values are: * &#x60;residential&#x60; –– The address is residential or a PO Box. * &#x60;commercial&#x60; –– The address is commercial. * &#x60;&#39;&#39;&#x60; –– Not enough information provided to be determined.  | 
**RecordType** | **string** | A description of the type of address. Populated if a DPV match is made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;Y&#x60;, &#x60;S&#x60;, or &#x60;D&#x60;). For more detailed information about each record type, see [US Verification Details](#tag/US-Verification-Types).  | 
**DefaultBuildingAddress** | **bool** | Designates whether or not the address is the default address for a building containing multiple delivery points.  | 
**County** | **string** | County name of the address city. | 
**CountyFips** | **string** | A 5-digit [FIPS county code](https://en.wikipedia.org/wiki/FIPS_county_code) which uniquely identifies &#x60;components[county]&#x60;. It consists of a 2-digit state code and a 3-digit county code.  | 
**CarrierRoute** | **string** | A 4-character code assigned to a mail delivery route within a ZIP code.  | 
**CarrierRouteType** | **string** | The type of &#x60;components[carrier_route]&#x60;. For more detailed information about each carrier route type, see [US Verification Details](#tag/US-Verification-Types).  | 
**PoBoxOnlyFlag** | **string** | Indicates the mailing facility for an address only supports PO Box deliveries and other forms of mail delivery are not available.  | 
**Latitude** | Pointer to **NullableFloat32** | A positive or negative decimal indicating the geographic latitude of the address, specifying the north-to-south position of a location. This should be used with &#x60;longitude&#x60; to pinpoint locations on a map. Will not be returned for undeliverable addresses or military addresses (state is &#x60;AA&#x60;, &#x60;AE&#x60;, or &#x60;AP&#x60;).  | [optional] 
**Longitude** | Pointer to **NullableFloat32** | A positive or negative decimal indicating the geographic longitude of the address, specifying the north-to-south position of a location. This should be used with &#x60;latitude&#x60; to pinpoint locations on a map. Will not be returned for undeliverable addresses or military addresses (state is &#x60;AA&#x60;, &#x60;AE&#x60;, or &#x60;AP&#x60;).  | [optional] 

## Methods

### NewUsComponents

`func NewUsComponents(primaryNumber string, streetPredirection string, streetName string, streetSuffix string, streetPostdirection string, secondaryDesignator string, secondaryNumber string, pmbDesignator string, pmbNumber string, extraSecondaryDesignator string, extraSecondaryNumber string, city string, state string, zipCode string, zipCodePlus4 string, zipCodeType ZipCodeType, deliveryPointBarcode string, addressType string, recordType string, defaultBuildingAddress bool, county string, countyFips string, carrierRoute string, carrierRouteType string, poBoxOnlyFlag string, ) *UsComponents`

NewUsComponents instantiates a new UsComponents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsComponentsWithDefaults

`func NewUsComponentsWithDefaults() *UsComponents`

NewUsComponentsWithDefaults instantiates a new UsComponents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryNumber

`func (o *UsComponents) GetPrimaryNumber() string`

GetPrimaryNumber returns the PrimaryNumber field if non-nil, zero value otherwise.

### GetPrimaryNumberOk

`func (o *UsComponents) GetPrimaryNumberOk() (*string, bool)`

GetPrimaryNumberOk returns a tuple with the PrimaryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNumber

`func (o *UsComponents) SetPrimaryNumber(v string)`

SetPrimaryNumber sets PrimaryNumber field to given value.


### GetStreetPredirection

`func (o *UsComponents) GetStreetPredirection() string`

GetStreetPredirection returns the StreetPredirection field if non-nil, zero value otherwise.

### GetStreetPredirectionOk

`func (o *UsComponents) GetStreetPredirectionOk() (*string, bool)`

GetStreetPredirectionOk returns a tuple with the StreetPredirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetPredirection

`func (o *UsComponents) SetStreetPredirection(v string)`

SetStreetPredirection sets StreetPredirection field to given value.


### GetStreetName

`func (o *UsComponents) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *UsComponents) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *UsComponents) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.


### GetStreetSuffix

`func (o *UsComponents) GetStreetSuffix() string`

GetStreetSuffix returns the StreetSuffix field if non-nil, zero value otherwise.

### GetStreetSuffixOk

`func (o *UsComponents) GetStreetSuffixOk() (*string, bool)`

GetStreetSuffixOk returns a tuple with the StreetSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetSuffix

`func (o *UsComponents) SetStreetSuffix(v string)`

SetStreetSuffix sets StreetSuffix field to given value.


### GetStreetPostdirection

`func (o *UsComponents) GetStreetPostdirection() string`

GetStreetPostdirection returns the StreetPostdirection field if non-nil, zero value otherwise.

### GetStreetPostdirectionOk

`func (o *UsComponents) GetStreetPostdirectionOk() (*string, bool)`

GetStreetPostdirectionOk returns a tuple with the StreetPostdirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetPostdirection

`func (o *UsComponents) SetStreetPostdirection(v string)`

SetStreetPostdirection sets StreetPostdirection field to given value.


### GetSecondaryDesignator

`func (o *UsComponents) GetSecondaryDesignator() string`

GetSecondaryDesignator returns the SecondaryDesignator field if non-nil, zero value otherwise.

### GetSecondaryDesignatorOk

`func (o *UsComponents) GetSecondaryDesignatorOk() (*string, bool)`

GetSecondaryDesignatorOk returns a tuple with the SecondaryDesignator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDesignator

`func (o *UsComponents) SetSecondaryDesignator(v string)`

SetSecondaryDesignator sets SecondaryDesignator field to given value.


### GetSecondaryNumber

`func (o *UsComponents) GetSecondaryNumber() string`

GetSecondaryNumber returns the SecondaryNumber field if non-nil, zero value otherwise.

### GetSecondaryNumberOk

`func (o *UsComponents) GetSecondaryNumberOk() (*string, bool)`

GetSecondaryNumberOk returns a tuple with the SecondaryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryNumber

`func (o *UsComponents) SetSecondaryNumber(v string)`

SetSecondaryNumber sets SecondaryNumber field to given value.


### GetPmbDesignator

`func (o *UsComponents) GetPmbDesignator() string`

GetPmbDesignator returns the PmbDesignator field if non-nil, zero value otherwise.

### GetPmbDesignatorOk

`func (o *UsComponents) GetPmbDesignatorOk() (*string, bool)`

GetPmbDesignatorOk returns a tuple with the PmbDesignator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmbDesignator

`func (o *UsComponents) SetPmbDesignator(v string)`

SetPmbDesignator sets PmbDesignator field to given value.


### GetPmbNumber

`func (o *UsComponents) GetPmbNumber() string`

GetPmbNumber returns the PmbNumber field if non-nil, zero value otherwise.

### GetPmbNumberOk

`func (o *UsComponents) GetPmbNumberOk() (*string, bool)`

GetPmbNumberOk returns a tuple with the PmbNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmbNumber

`func (o *UsComponents) SetPmbNumber(v string)`

SetPmbNumber sets PmbNumber field to given value.


### GetExtraSecondaryDesignator

`func (o *UsComponents) GetExtraSecondaryDesignator() string`

GetExtraSecondaryDesignator returns the ExtraSecondaryDesignator field if non-nil, zero value otherwise.

### GetExtraSecondaryDesignatorOk

`func (o *UsComponents) GetExtraSecondaryDesignatorOk() (*string, bool)`

GetExtraSecondaryDesignatorOk returns a tuple with the ExtraSecondaryDesignator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraSecondaryDesignator

`func (o *UsComponents) SetExtraSecondaryDesignator(v string)`

SetExtraSecondaryDesignator sets ExtraSecondaryDesignator field to given value.


### GetExtraSecondaryNumber

`func (o *UsComponents) GetExtraSecondaryNumber() string`

GetExtraSecondaryNumber returns the ExtraSecondaryNumber field if non-nil, zero value otherwise.

### GetExtraSecondaryNumberOk

`func (o *UsComponents) GetExtraSecondaryNumberOk() (*string, bool)`

GetExtraSecondaryNumberOk returns a tuple with the ExtraSecondaryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraSecondaryNumber

`func (o *UsComponents) SetExtraSecondaryNumber(v string)`

SetExtraSecondaryNumber sets ExtraSecondaryNumber field to given value.


### GetCity

`func (o *UsComponents) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UsComponents) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UsComponents) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *UsComponents) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UsComponents) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UsComponents) SetState(v string)`

SetState sets State field to given value.


### GetZipCode

`func (o *UsComponents) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UsComponents) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UsComponents) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.


### GetZipCodePlus4

`func (o *UsComponents) GetZipCodePlus4() string`

GetZipCodePlus4 returns the ZipCodePlus4 field if non-nil, zero value otherwise.

### GetZipCodePlus4Ok

`func (o *UsComponents) GetZipCodePlus4Ok() (*string, bool)`

GetZipCodePlus4Ok returns a tuple with the ZipCodePlus4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodePlus4

`func (o *UsComponents) SetZipCodePlus4(v string)`

SetZipCodePlus4 sets ZipCodePlus4 field to given value.


### GetZipCodeType

`func (o *UsComponents) GetZipCodeType() ZipCodeType`

GetZipCodeType returns the ZipCodeType field if non-nil, zero value otherwise.

### GetZipCodeTypeOk

`func (o *UsComponents) GetZipCodeTypeOk() (*ZipCodeType, bool)`

GetZipCodeTypeOk returns a tuple with the ZipCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeType

`func (o *UsComponents) SetZipCodeType(v ZipCodeType)`

SetZipCodeType sets ZipCodeType field to given value.


### GetDeliveryPointBarcode

`func (o *UsComponents) GetDeliveryPointBarcode() string`

GetDeliveryPointBarcode returns the DeliveryPointBarcode field if non-nil, zero value otherwise.

### GetDeliveryPointBarcodeOk

`func (o *UsComponents) GetDeliveryPointBarcodeOk() (*string, bool)`

GetDeliveryPointBarcodeOk returns a tuple with the DeliveryPointBarcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPointBarcode

`func (o *UsComponents) SetDeliveryPointBarcode(v string)`

SetDeliveryPointBarcode sets DeliveryPointBarcode field to given value.


### GetAddressType

`func (o *UsComponents) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *UsComponents) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *UsComponents) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetRecordType

`func (o *UsComponents) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *UsComponents) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *UsComponents) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetDefaultBuildingAddress

`func (o *UsComponents) GetDefaultBuildingAddress() bool`

GetDefaultBuildingAddress returns the DefaultBuildingAddress field if non-nil, zero value otherwise.

### GetDefaultBuildingAddressOk

`func (o *UsComponents) GetDefaultBuildingAddressOk() (*bool, bool)`

GetDefaultBuildingAddressOk returns a tuple with the DefaultBuildingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBuildingAddress

`func (o *UsComponents) SetDefaultBuildingAddress(v bool)`

SetDefaultBuildingAddress sets DefaultBuildingAddress field to given value.


### GetCounty

`func (o *UsComponents) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *UsComponents) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *UsComponents) SetCounty(v string)`

SetCounty sets County field to given value.


### GetCountyFips

`func (o *UsComponents) GetCountyFips() string`

GetCountyFips returns the CountyFips field if non-nil, zero value otherwise.

### GetCountyFipsOk

`func (o *UsComponents) GetCountyFipsOk() (*string, bool)`

GetCountyFipsOk returns a tuple with the CountyFips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyFips

`func (o *UsComponents) SetCountyFips(v string)`

SetCountyFips sets CountyFips field to given value.


### GetCarrierRoute

`func (o *UsComponents) GetCarrierRoute() string`

GetCarrierRoute returns the CarrierRoute field if non-nil, zero value otherwise.

### GetCarrierRouteOk

`func (o *UsComponents) GetCarrierRouteOk() (*string, bool)`

GetCarrierRouteOk returns a tuple with the CarrierRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierRoute

`func (o *UsComponents) SetCarrierRoute(v string)`

SetCarrierRoute sets CarrierRoute field to given value.


### GetCarrierRouteType

`func (o *UsComponents) GetCarrierRouteType() string`

GetCarrierRouteType returns the CarrierRouteType field if non-nil, zero value otherwise.

### GetCarrierRouteTypeOk

`func (o *UsComponents) GetCarrierRouteTypeOk() (*string, bool)`

GetCarrierRouteTypeOk returns a tuple with the CarrierRouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierRouteType

`func (o *UsComponents) SetCarrierRouteType(v string)`

SetCarrierRouteType sets CarrierRouteType field to given value.


### GetPoBoxOnlyFlag

`func (o *UsComponents) GetPoBoxOnlyFlag() string`

GetPoBoxOnlyFlag returns the PoBoxOnlyFlag field if non-nil, zero value otherwise.

### GetPoBoxOnlyFlagOk

`func (o *UsComponents) GetPoBoxOnlyFlagOk() (*string, bool)`

GetPoBoxOnlyFlagOk returns a tuple with the PoBoxOnlyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoBoxOnlyFlag

`func (o *UsComponents) SetPoBoxOnlyFlag(v string)`

SetPoBoxOnlyFlag sets PoBoxOnlyFlag field to given value.


### GetLatitude

`func (o *UsComponents) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *UsComponents) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *UsComponents) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *UsComponents) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *UsComponents) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *UsComponents) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *UsComponents) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *UsComponents) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *UsComponents) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *UsComponents) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *UsComponents) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *UsComponents) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


