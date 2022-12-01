# \IntlVerificationsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**verifyBulk**](IntlVerificationsApi.md#verifyBulk) | **Post** /bulk/intl_verifications | verifyBulk
[**verifySingle**](IntlVerificationsApi.md#verifySingle) | **Post** /intl_verifications | verifySingle



## BulkIntlVerifications

> IntlVerifications verifyBulk(ctx).IntlVerificationsPayload(intlVerificationsPayload).Execute()

verifyBulk



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    intlVerificationsPayload := *openapiclient.NewIntlVerificationsPayload([]openapiclient.MultipleComponentsIntl{*openapiclient.NewMultipleComponentsIntl("PrimaryLine_example", openapiclient.country_extended("AD"))}) // IntlVerificationsPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntlVerificationsApi.verifyBulk(context.Background()).IntlVerificationsPayload(intlVerificationsPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntlVerificationsApi.verifyBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `verifyBulk`: IntlVerifications
    fmt.Fprintf(os.Stdout, "Response from `IntlVerificationsApi.verifyBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkIntlVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intlVerificationsPayload** | [**IntlVerificationsPayload**](IntlVerificationsPayload.md) |  | 

### Return type

[**IntlVerifications**](IntlVerifications.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntlVerification

> IntlVerification verifySingle(ctx).IntlVerificationWritable(intlVerificationWritable).XLangOutput(xLangOutput).Execute()

verifySingle



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    intlVerificationWritable := *openapiclient.NewIntlVerificationWritable() // IntlVerificationWritable | 
    xLangOutput := "xLangOutput_example" // string | * `native` - Translate response to the native language of the country in the request * `match` - match the response to the language in the request  Default response is in English.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntlVerificationsApi.verifySingle(context.Background()).IntlVerificationWritable(intlVerificationWritable).XLangOutput(xLangOutput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntlVerificationsApi.verifySingle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `verifySingle`: IntlVerification
    fmt.Fprintf(os.Stdout, "Response from `IntlVerificationsApi.verifySingle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntlVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intlVerificationWritable** | [**IntlVerificationWritable**](IntlVerificationWritable.md) |  | 
 **xLangOutput** | **string** | * &#x60;native&#x60; - Translate response to the native language of the country in the request * &#x60;match&#x60; - match the response to the language in the request  Default response is in English.  | 

### Return type

[**IntlVerification**](IntlVerification.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

