# \UsVerificationsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyBulk**](UsVerificationsApi.md#VerifyBulk) | **Post** /bulk/us_verifications | verifyBulk
[**VerifySingle**](UsVerificationsApi.md#VerifySingle) | **Post** /us_verifications | verifySingle



## BulkUsVerifications

> UsVerifications VerifyBulk(ctx).MultipleComponentsList(multipleComponentsList).Case_(case_).Execute()

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
    multipleComponentsList := *openapiclient.NewMultipleComponentsList([]openapiclient.MultipleComponents{*openapiclient.NewMultipleComponents("PrimaryLine_example")}) // MultipleComponentsList | 
    case_ := "case__example" // string | Casing of the verified address. Possible values are `upper` and `proper` for uppercased (e.g. \"PO BOX\") and proper-cased (e.g. \"PO Box\"), respectively. (optional) (default to "upper")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsVerificationsApi.VerifyBulk(context.Background()).MultipleComponentsList(multipleComponentsList).Case_(case_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsVerificationsApi.VerifyBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyBulk`: UsVerifications
    fmt.Fprintf(os.Stdout, "Response from `UsVerificationsApi.VerifyBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUsVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipleComponentsList** | [**MultipleComponentsList**](MultipleComponentsList.md) |  | 
 **case_** | **string** | Casing of the verified address. Possible values are &#x60;upper&#x60; and &#x60;proper&#x60; for uppercased (e.g. \&quot;PO BOX\&quot;) and proper-cased (e.g. \&quot;PO Box\&quot;), respectively. | [default to &quot;upper&quot;]

### Return type

[**UsVerifications**](UsVerifications.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsVerification

> UsVerification VerifySingle(ctx).UsVerificationsWritable(usVerificationsWritable).Case_(case_).Execute()

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
    usVerificationsWritable := *openapiclient.NewUsVerificationsWritable() // UsVerificationsWritable | 
    case_ := "case__example" // string | Casing of the verified address. Possible values are `upper` and `proper` for uppercased (e.g. \"PO BOX\") and proper-cased (e.g. \"PO Box\"), respectively. (optional) (default to "upper")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsVerificationsApi.VerifySingle(context.Background()).UsVerificationsWritable(usVerificationsWritable).Case_(case_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsVerificationsApi.VerifySingle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifySingle`: UsVerification
    fmt.Fprintf(os.Stdout, "Response from `UsVerificationsApi.VerifySingle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usVerificationsWritable** | [**UsVerificationsWritable**](UsVerificationsWritable.md) |  | 
 **case_** | **string** | Casing of the verified address. Possible values are &#x60;upper&#x60; and &#x60;proper&#x60; for uppercased (e.g. \&quot;PO BOX\&quot;) and proper-cased (e.g. \&quot;PO Box\&quot;), respectively. | [default to &quot;upper&quot;]

### Return type

[**UsVerification**](UsVerification.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

