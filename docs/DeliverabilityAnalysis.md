# DeliverabilityAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpvConfirmation** | **string** | Result of Delivery Point Validation (DPV), which determines whether or not the address is deliverable by the USPS. Possible values are: * &#x60;Y&#x60; –– The address is deliverable by the USPS. * &#x60;S&#x60; –– The address is deliverable by removing the provided secondary unit designator. This information may be incorrect or unnecessary. * &#x60;D&#x60; –– The address is deliverable to the building&#39;s default address but is missing a secondary unit designator and/or number.   There is a chance the mail will not reach the intended recipient. * &#x60;N&#x60; –– The address is not deliverable according to the USPS, but parts of the address are valid (such as the street and ZIP code). * &#x60;&#39;&#39;&#x60; –– This address is not deliverable. No matching street could be found within the city or ZIP code.  | 
**DpvCmra** | **string** | indicates whether or not the address is [CMRA-authorized](https://en.wikipedia.org/wiki/Commercial_mail_receiving_agency). Possible values are: * &#x60;Y&#x60; –– Address is CMRA-authorized. * &#x60;N&#x60; –– Address is not CMRA-authorized. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string).  | 
**DpvVacant** | **string** | indicates that an address was once deliverable, but has become vacant and is no longer receiving deliveries. Possible values are: * &#x60;Y&#x60; –– Address is vacant. * &#x60;N&#x60; –– Address is not vacant. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string).  | 
**DpvActive** | **string** | Corresponds to the USPS field &#x60;dpv_no_stat&#x60;. Indicates that an address has been vacated in the recent past, and is no longer receiving deliveries. If it&#39;s been unoccupied for 90+ days, or temporarily vacant, this will be flagged. Possible values are: * &#x60;Y&#x60; –– Address is active. * &#x60;N&#x60; –– Address is not active. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string).  | 
**DpvInactiveReason** | **string** | Indicates the reason why an address is vacant or no longer receiving deliveries. Possible values are: * &#x60;01&#x60; –– Address does not receive mail from the USPS directly, but is serviced by a drop address. * &#x60;02&#x60; –– Address not yet deliverable. * &#x60;03&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string). * &#x60;04&#x60; –– Address is a College, Military Zone, or other type. * &#x60;05&#x60; –– Address no longer receives deliveries. * &#x60;06&#x60; –– Address is missing required secondary information. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made or the address is active.  | 
**DpvThrowback** | **string** | Indicates a street address for which mail is delivered to a PO Box. Possible values are: * &#x60;Y&#x60; –– Address is a PO Box throwback delivery point. * &#x60;N&#x60; –– Address is not a PO Box throwback delivery point. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string).  | 
**DpvNonDeliveryDayFlag** | **string** | Indicates whether deliveries are not performed on one or more days of the week at an address. Possible values are: * &#x60;Y&#x60; –– Mail delivery does not occur on some days of the week. * &#x60;N&#x60; –– Mail delivery occurs every day of the week. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string).  | 
**DpvNonDeliveryDayValues** | **string** | Indicates days of the week (starting on Sunday) deliveries are not performed at an address. For example: * &#x60;YNNNNNN&#x60; –– Mail delivery does not occur on Sunday&#39;s. * &#x60;NYNNNYN&#x60; –– Mail delivery does not occur on Monday&#39;s or Friday&#39;s. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string) or address receives mail every day of the week (&#x60;deliverability_analysis[dpv_non_delivery_day_flag]&#x60; is &#x60;N&#x60; or an empty string).  | 
**DpvNoSecureLocation** | **string** | Indicates packages to this address will not be left due to security concerns. Possible values are: * &#x60;Y&#x60; –– Address does not have a secure mailbox. * &#x60;N&#x60; –– Address has a secure mailbox. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string).  | 
**DpvDoorNotAccessible** | **string** | Indicates the door of the address is not accessible for mail delivery. Possible values are: * &#x60;Y&#x60; –– Door is not accessible. * &#x60;N&#x60; –– Door is accessible. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string).  | 
**DpvFootnotes** | [**[]DpvFootnote**](DpvFootnote.md) | An array of 2-character strings that gives more insight into how &#x60;deliverability_analysis[dpv_confirmation]&#x60; was determined. Will always include at least 1 string, and can include up to 3. For details, see [US Verification Details](#tag/US-Verification-Types).  | 
**EwsMatch** | **bool** | indicates whether or not an address has been flagged in the [Early Warning System](https://docs.informatica.com/data-engineering/data-engineering-quality/10-4-0/address-validator-port-reference/postal-carrier-certification-data-ports/early-warning-system-return-code.html), meaning the address is under development and not yet ready to receive mail. However, it should become available in a few months.  | 
**LacsIndicator** | **string** | indicates whether this address has been converted by [LACS&lt;sup&gt;Link&lt;/sup&gt;](https://postalpro.usps.com/address-quality/lacslink). LACS&lt;sup&gt;Link&lt;/sup&gt; corrects outdated addresses into their modern counterparts. Possible values are: * &#x60;Y&#x60; –– New address produced with a matching record in LACS&lt;sup&gt;Link&lt;/sup&gt;. * &#x60;N&#x60; –– New address could not be produced with a matching record in LACS&lt;sup&gt;Link&lt;/sup&gt;. * &#x60;&#39;&#39;&#x60; –– A DPV match is not made (&#x60;deliverability_analysis[dpv_confirmation]&#x60; is &#x60;N&#x60; or an empty string).  | 
**LacsReturnCode** | **string** | A code indicating how &#x60;deliverability_analysis[lacs_indicator]&#x60; was determined. Possible values are: * &#x60;A&#x60; — A new address was produced because a match was found in LACS&lt;sup&gt;Link&lt;/sup&gt;. * &#x60;92&#x60; — A LACS&lt;sup&gt;Link&lt;/sup&gt; record was matched after dropping secondary information. * &#x60;14&#x60; — A match was found in LACS&lt;sup&gt;Link&lt;/sup&gt;, but could not be converted to a deliverable address. * &#x60;00&#x60; — A match was not found in LACS&lt;sup&gt;Link&lt;/sup&gt;, and no new address was produced. * &#x60;&#39;&#39;&#x60; — LACS&lt;sup&gt;Link&lt;/sup&gt; was not attempted.  | 
**SuiteReturnCode** | **string** | A return code that indicates whether the address was matched and corrected by [Suite&lt;sup&gt;Link&lt;/sup&gt;](https://postalpro.usps.com/address-quality-solutions/suitelink). Suite&lt;sup&gt;Link&lt;/sup&gt; attempts to provide secondary information to business addresses. Possible values are: * &#x60;A&#x60; –– A Suite&lt;sup&gt;Link&lt;/sup&gt; match was found and secondary information was added. * &#x60;00&#x60; –– A Suite&lt;sup&gt;Link&lt;/sup&gt; match could not be found and no secondary information was added. * &#x60;&#39;&#39;&#x60; –– Suite&lt;sup&gt;Link&lt;/sup&gt; lookup was not attempted.  | 

## Methods

### NewDeliverabilityAnalysis

`func NewDeliverabilityAnalysis(dpvConfirmation string, dpvCmra string, dpvVacant string, dpvActive string, dpvInactiveReason string, dpvThrowback string, dpvNonDeliveryDayFlag string, dpvNonDeliveryDayValues string, dpvNoSecureLocation string, dpvDoorNotAccessible string, dpvFootnotes []DpvFootnote, ewsMatch bool, lacsIndicator string, lacsReturnCode string, suiteReturnCode string, ) *DeliverabilityAnalysis`

NewDeliverabilityAnalysis instantiates a new DeliverabilityAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverabilityAnalysisWithDefaults

`func NewDeliverabilityAnalysisWithDefaults() *DeliverabilityAnalysis`

NewDeliverabilityAnalysisWithDefaults instantiates a new DeliverabilityAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpvConfirmation

`func (o *DeliverabilityAnalysis) GetDpvConfirmation() string`

GetDpvConfirmation returns the DpvConfirmation field if non-nil, zero value otherwise.

### GetDpvConfirmationOk

`func (o *DeliverabilityAnalysis) GetDpvConfirmationOk() (*string, bool)`

GetDpvConfirmationOk returns a tuple with the DpvConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvConfirmation

`func (o *DeliverabilityAnalysis) SetDpvConfirmation(v string)`

SetDpvConfirmation sets DpvConfirmation field to given value.


### GetDpvCmra

`func (o *DeliverabilityAnalysis) GetDpvCmra() string`

GetDpvCmra returns the DpvCmra field if non-nil, zero value otherwise.

### GetDpvCmraOk

`func (o *DeliverabilityAnalysis) GetDpvCmraOk() (*string, bool)`

GetDpvCmraOk returns a tuple with the DpvCmra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvCmra

`func (o *DeliverabilityAnalysis) SetDpvCmra(v string)`

SetDpvCmra sets DpvCmra field to given value.


### GetDpvVacant

`func (o *DeliverabilityAnalysis) GetDpvVacant() string`

GetDpvVacant returns the DpvVacant field if non-nil, zero value otherwise.

### GetDpvVacantOk

`func (o *DeliverabilityAnalysis) GetDpvVacantOk() (*string, bool)`

GetDpvVacantOk returns a tuple with the DpvVacant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvVacant

`func (o *DeliverabilityAnalysis) SetDpvVacant(v string)`

SetDpvVacant sets DpvVacant field to given value.


### GetDpvActive

`func (o *DeliverabilityAnalysis) GetDpvActive() string`

GetDpvActive returns the DpvActive field if non-nil, zero value otherwise.

### GetDpvActiveOk

`func (o *DeliverabilityAnalysis) GetDpvActiveOk() (*string, bool)`

GetDpvActiveOk returns a tuple with the DpvActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvActive

`func (o *DeliverabilityAnalysis) SetDpvActive(v string)`

SetDpvActive sets DpvActive field to given value.


### GetDpvInactiveReason

`func (o *DeliverabilityAnalysis) GetDpvInactiveReason() string`

GetDpvInactiveReason returns the DpvInactiveReason field if non-nil, zero value otherwise.

### GetDpvInactiveReasonOk

`func (o *DeliverabilityAnalysis) GetDpvInactiveReasonOk() (*string, bool)`

GetDpvInactiveReasonOk returns a tuple with the DpvInactiveReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvInactiveReason

`func (o *DeliverabilityAnalysis) SetDpvInactiveReason(v string)`

SetDpvInactiveReason sets DpvInactiveReason field to given value.


### GetDpvThrowback

`func (o *DeliverabilityAnalysis) GetDpvThrowback() string`

GetDpvThrowback returns the DpvThrowback field if non-nil, zero value otherwise.

### GetDpvThrowbackOk

`func (o *DeliverabilityAnalysis) GetDpvThrowbackOk() (*string, bool)`

GetDpvThrowbackOk returns a tuple with the DpvThrowback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvThrowback

`func (o *DeliverabilityAnalysis) SetDpvThrowback(v string)`

SetDpvThrowback sets DpvThrowback field to given value.


### GetDpvNonDeliveryDayFlag

`func (o *DeliverabilityAnalysis) GetDpvNonDeliveryDayFlag() string`

GetDpvNonDeliveryDayFlag returns the DpvNonDeliveryDayFlag field if non-nil, zero value otherwise.

### GetDpvNonDeliveryDayFlagOk

`func (o *DeliverabilityAnalysis) GetDpvNonDeliveryDayFlagOk() (*string, bool)`

GetDpvNonDeliveryDayFlagOk returns a tuple with the DpvNonDeliveryDayFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvNonDeliveryDayFlag

`func (o *DeliverabilityAnalysis) SetDpvNonDeliveryDayFlag(v string)`

SetDpvNonDeliveryDayFlag sets DpvNonDeliveryDayFlag field to given value.


### GetDpvNonDeliveryDayValues

`func (o *DeliverabilityAnalysis) GetDpvNonDeliveryDayValues() string`

GetDpvNonDeliveryDayValues returns the DpvNonDeliveryDayValues field if non-nil, zero value otherwise.

### GetDpvNonDeliveryDayValuesOk

`func (o *DeliverabilityAnalysis) GetDpvNonDeliveryDayValuesOk() (*string, bool)`

GetDpvNonDeliveryDayValuesOk returns a tuple with the DpvNonDeliveryDayValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvNonDeliveryDayValues

`func (o *DeliverabilityAnalysis) SetDpvNonDeliveryDayValues(v string)`

SetDpvNonDeliveryDayValues sets DpvNonDeliveryDayValues field to given value.


### GetDpvNoSecureLocation

`func (o *DeliverabilityAnalysis) GetDpvNoSecureLocation() string`

GetDpvNoSecureLocation returns the DpvNoSecureLocation field if non-nil, zero value otherwise.

### GetDpvNoSecureLocationOk

`func (o *DeliverabilityAnalysis) GetDpvNoSecureLocationOk() (*string, bool)`

GetDpvNoSecureLocationOk returns a tuple with the DpvNoSecureLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvNoSecureLocation

`func (o *DeliverabilityAnalysis) SetDpvNoSecureLocation(v string)`

SetDpvNoSecureLocation sets DpvNoSecureLocation field to given value.


### GetDpvDoorNotAccessible

`func (o *DeliverabilityAnalysis) GetDpvDoorNotAccessible() string`

GetDpvDoorNotAccessible returns the DpvDoorNotAccessible field if non-nil, zero value otherwise.

### GetDpvDoorNotAccessibleOk

`func (o *DeliverabilityAnalysis) GetDpvDoorNotAccessibleOk() (*string, bool)`

GetDpvDoorNotAccessibleOk returns a tuple with the DpvDoorNotAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvDoorNotAccessible

`func (o *DeliverabilityAnalysis) SetDpvDoorNotAccessible(v string)`

SetDpvDoorNotAccessible sets DpvDoorNotAccessible field to given value.


### GetDpvFootnotes

`func (o *DeliverabilityAnalysis) GetDpvFootnotes() []DpvFootnote`

GetDpvFootnotes returns the DpvFootnotes field if non-nil, zero value otherwise.

### GetDpvFootnotesOk

`func (o *DeliverabilityAnalysis) GetDpvFootnotesOk() (*[]DpvFootnote, bool)`

GetDpvFootnotesOk returns a tuple with the DpvFootnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpvFootnotes

`func (o *DeliverabilityAnalysis) SetDpvFootnotes(v []DpvFootnote)`

SetDpvFootnotes sets DpvFootnotes field to given value.


### GetEwsMatch

`func (o *DeliverabilityAnalysis) GetEwsMatch() bool`

GetEwsMatch returns the EwsMatch field if non-nil, zero value otherwise.

### GetEwsMatchOk

`func (o *DeliverabilityAnalysis) GetEwsMatchOk() (*bool, bool)`

GetEwsMatchOk returns a tuple with the EwsMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwsMatch

`func (o *DeliverabilityAnalysis) SetEwsMatch(v bool)`

SetEwsMatch sets EwsMatch field to given value.


### GetLacsIndicator

`func (o *DeliverabilityAnalysis) GetLacsIndicator() string`

GetLacsIndicator returns the LacsIndicator field if non-nil, zero value otherwise.

### GetLacsIndicatorOk

`func (o *DeliverabilityAnalysis) GetLacsIndicatorOk() (*string, bool)`

GetLacsIndicatorOk returns a tuple with the LacsIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacsIndicator

`func (o *DeliverabilityAnalysis) SetLacsIndicator(v string)`

SetLacsIndicator sets LacsIndicator field to given value.


### GetLacsReturnCode

`func (o *DeliverabilityAnalysis) GetLacsReturnCode() string`

GetLacsReturnCode returns the LacsReturnCode field if non-nil, zero value otherwise.

### GetLacsReturnCodeOk

`func (o *DeliverabilityAnalysis) GetLacsReturnCodeOk() (*string, bool)`

GetLacsReturnCodeOk returns a tuple with the LacsReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacsReturnCode

`func (o *DeliverabilityAnalysis) SetLacsReturnCode(v string)`

SetLacsReturnCode sets LacsReturnCode field to given value.


### GetSuiteReturnCode

`func (o *DeliverabilityAnalysis) GetSuiteReturnCode() string`

GetSuiteReturnCode returns the SuiteReturnCode field if non-nil, zero value otherwise.

### GetSuiteReturnCodeOk

`func (o *DeliverabilityAnalysis) GetSuiteReturnCodeOk() (*string, bool)`

GetSuiteReturnCodeOk returns a tuple with the SuiteReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteReturnCode

`func (o *DeliverabilityAnalysis) SetSuiteReturnCode(v string)`

SetSuiteReturnCode sets SuiteReturnCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


