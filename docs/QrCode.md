# QrCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **string** | Sets how a QR code is being positioned in the document. | 
**Top** | Pointer to **string** | Vertical distance(in inches) to place QR code from the top. | [optional] 
**Right** | Pointer to **string** | Horizonal distance(in inches) to place QR code from the right. | [optional] 
**Left** | Pointer to **string** | Horizonal distance(in inches) to place QR code from the left. | [optional] 
**Bottom** | Pointer to **string** | Vertical distance(in inches) to place QR code from the bottom. | [optional] 
**RedirectUrl** | **string** | The url to redirect the user when a QR code is scanned. The url must start with &#x60;https://&#x60; | 
**Width** | **string** | The size(in inches) of the QR code. All QR codes are generated as a square. | 

## Methods

### NewQrCode

`func NewQrCode(position string, redirectUrl string, width string, ) *QrCode`

NewQrCode instantiates a new QrCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQrCodeWithDefaults

`func NewQrCodeWithDefaults() *QrCode`

NewQrCodeWithDefaults instantiates a new QrCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *QrCode) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *QrCode) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *QrCode) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetTop

`func (o *QrCode) GetTop() string`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *QrCode) GetTopOk() (*string, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *QrCode) SetTop(v string)`

SetTop sets Top field to given value.

### HasTop

`func (o *QrCode) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetRight

`func (o *QrCode) GetRight() string`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *QrCode) GetRightOk() (*string, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *QrCode) SetRight(v string)`

SetRight sets Right field to given value.

### HasRight

`func (o *QrCode) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetLeft

`func (o *QrCode) GetLeft() string`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *QrCode) GetLeftOk() (*string, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *QrCode) SetLeft(v string)`

SetLeft sets Left field to given value.

### HasLeft

`func (o *QrCode) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### GetBottom

`func (o *QrCode) GetBottom() string`

GetBottom returns the Bottom field if non-nil, zero value otherwise.

### GetBottomOk

`func (o *QrCode) GetBottomOk() (*string, bool)`

GetBottomOk returns a tuple with the Bottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottom

`func (o *QrCode) SetBottom(v string)`

SetBottom sets Bottom field to given value.

### HasBottom

`func (o *QrCode) HasBottom() bool`

HasBottom returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *QrCode) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *QrCode) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *QrCode) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetWidth

`func (o *QrCode) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *QrCode) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *QrCode) SetWidth(v string)`

SetWidth sets Width field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


