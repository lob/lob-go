# \ZipLookupsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**lookup**](ZipLookupsApi.md#lookup) | **Post** /us_zip_lookups | lookup



## ZipLookup

> Zip lookup(ctx).ZipEditable(zipEditable).Execute()

lookup



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
    zipEditable := *openapiclient.NewZipEditable() // ZipEditable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZipLookupsApi.lookup(context.Background()).ZipEditable(zipEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZipLookupsApi.lookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `lookup`: Zip
    fmt.Fprintf(os.Stdout, "Response from `ZipLookupsApi.lookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZipLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zipEditable** | [**ZipEditable**](ZipEditable.md) |  | 

### Return type

[**Zip**](Zip.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

