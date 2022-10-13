# \UsAutocompletionsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsAutocompletion**](UsAutocompletionsApi.md#UsAutocompletion) | **Post** /us_autocompletions | autocomplete



## UsAutocompletion

> UsAutocompletions UsAutocompletion(ctx).UsAutocompletionsWritable(usAutocompletionsWritable).Execute()

autocomplete



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
    usAutocompletionsWritable := *openapiclient.NewUsAutocompletionsWritable("AddressPrefix_example") // UsAutocompletionsWritable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsAutocompletionsApi.UsAutocompletion(context.Background()).UsAutocompletionsWritable(usAutocompletionsWritable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsAutocompletionsApi.UsAutocompletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsAutocompletion`: UsAutocompletions
    fmt.Fprintf(os.Stdout, "Response from `UsAutocompletionsApi.UsAutocompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsAutocompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usAutocompletionsWritable** | [**UsAutocompletionsWritable**](UsAutocompletionsWritable.md) |  | 

### Return type

[**UsAutocompletions**](UsAutocompletions.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

