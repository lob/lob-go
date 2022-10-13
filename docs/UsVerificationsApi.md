# \UsVerificationsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUsVerifications**](UsVerificationsApi.md#BulkUsVerifications) | **Post** /bulk/us_verifications | verifyBulk
[**UsVerification**](UsVerificationsApi.md#UsVerification) | **Post** /us_verifications | verifySingle



## BulkUsVerifications

> UsVerifications BulkUsVerifications(ctx).MultipleComponentsList(multipleComponentsList).Case_(case_).Execute()

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
    resp, r, err := apiClient.UsVerificationsApi.BulkUsVerifications(context.Background()).MultipleComponentsList(multipleComponentsList).Case_(case_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsVerificationsApi.BulkUsVerifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUsVerifications`: UsVerifications
    fmt.Fprintf(os.Stdout, "Response from `UsVerificationsApi.BulkUsVerifications`: %v\n", resp)
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

> UsVerification UsVerification(ctx).UsVerificationsWritable(usVerificationsWritable).Case_(case_).Execute()

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
    resp, r, err := apiClient.UsVerificationsApi.UsVerification(context.Background()).UsVerificationsWritable(usVerificationsWritable).Case_(case_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsVerificationsApi.UsVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsVerification`: UsVerification
    fmt.Fprintf(os.Stdout, "Response from `UsVerificationsApi.UsVerification`: %v\n", resp)
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

