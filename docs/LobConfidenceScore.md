# LobConfidenceScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **NullableFloat32** | A numerical score between 0 and 100 that represents the percentage of mailpieces Lob has sent to this addresses that have been delivered successfully over the past 2 years. Will be &#x60;null&#x60; if no tracking data exists for this address.  | [optional] 
**Level** | Pointer to **string** | indicates the likelihood that the address is a valid, mail-receiving address. Possible values are:   - &#x60;high&#x60; — Over 70% of mailpieces Lob has sent to this address were delivered successfully and recent mailings were also successful.   - &#x60;medium&#x60; — Between 40% and 70% of mailpieces Lob has sent to this address were delivered successfully.   - &#x60;low&#x60; — Less than 40% of mailpieces Lob has sent to this address were delivered successfully and recent mailings weren&#39;t successful.   - &#x60;\&quot;\&quot;&#x60; — No tracking data exists for this address or lob deliverability was unable to find a corresponding level of mail success.  | [optional] 

## Methods

### NewLobConfidenceScore

`func NewLobConfidenceScore() *LobConfidenceScore`

NewLobConfidenceScore instantiates a new LobConfidenceScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobConfidenceScoreWithDefaults

`func NewLobConfidenceScoreWithDefaults() *LobConfidenceScore`

NewLobConfidenceScoreWithDefaults instantiates a new LobConfidenceScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *LobConfidenceScore) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *LobConfidenceScore) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *LobConfidenceScore) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *LobConfidenceScore) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *LobConfidenceScore) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *LobConfidenceScore) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetLevel

`func (o *LobConfidenceScore) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LobConfidenceScore) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LobConfidenceScore) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *LobConfidenceScore) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


