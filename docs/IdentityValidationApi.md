# \IdentityValidationApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityValidation**](IdentityValidationApi.md#IdentityValidation) | **Post** /identity_validation | validate



## IdentityValidation

> IdentityValidation IdentityValidation(ctx).MultiLineAddress(multiLineAddress).Execute()

validate



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
    multiLineAddress := *openapiclient.NewMultiLineAddress("PrimaryLine_example") // MultiLineAddress | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityValidationApi.IdentityValidation(context.Background()).MultiLineAddress(multiLineAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityValidationApi.IdentityValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityValidation`: IdentityValidation
    fmt.Fprintf(os.Stdout, "Response from `IdentityValidationApi.IdentityValidation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiLineAddress** | [**MultiLineAddress**](MultiLineAddress.md) |  | 

### Return type

[**IdentityValidation**](IdentityValidation.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

