# \DefaultApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Placeholder_no_call**](DefaultApi.md#Placeholder_no_call) | **Get** /shared_dont_call | placeholder_no_call



## Placeholder

> PlaceholderModel Placeholder_no_call(ctx).Execute()

placeholder_no_call



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Placeholder_no_call(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Placeholder_no_call``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Placeholder_no_call`: PlaceholderModel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Placeholder_no_call`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlaceholderRequest struct via the builder pattern


### Return type

[**PlaceholderModel**](PlaceholderModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

