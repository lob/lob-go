# \IntlAutocompletionsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Autocomplete**](IntlAutocompletionsApi.md#Autocomplete) | **Post** /intl_autocompletions | autocomplete



## IntlAutocompletion

> IntlAutocompletions Autocomplete(ctx).IntlAutocompletionsWritable(intlAutocompletionsWritable).XLangOutput(xLangOutput).Execute()

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
    intlAutocompletionsWritable := *openapiclient.NewIntlAutocompletionsWritable("AddressPrefix_example", openapiclient.country_extended("AD")) // IntlAutocompletionsWritable | 
    xLangOutput := "xLangOutput_example" // string | * `native` - Translate response to the native language of the country in the request * `match` - match the response to the language in the request  Default response is in English.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntlAutocompletionsApi.Autocomplete(context.Background()).IntlAutocompletionsWritable(intlAutocompletionsWritable).XLangOutput(xLangOutput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntlAutocompletionsApi.Autocomplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Autocomplete`: IntlAutocompletions
    fmt.Fprintf(os.Stdout, "Response from `IntlAutocompletionsApi.Autocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntlAutocompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intlAutocompletionsWritable** | [**IntlAutocompletionsWritable**](IntlAutocompletionsWritable.md) |  | 
 **xLangOutput** | **string** | * &#x60;native&#x60; - Translate response to the native language of the country in the request * &#x60;match&#x60; - match the response to the language in the request  Default response is in English.  | 

### Return type

[**IntlAutocompletions**](IntlAutocompletions.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

