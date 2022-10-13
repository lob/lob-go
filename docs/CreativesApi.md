# \CreativesApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreativeCreate**](CreativesApi.md#CreativeCreate) | **Post** /creatives | create
[**CreativeRetrieve**](CreativesApi.md#CreativeRetrieve) | **Get** /creatives/{crv_id} | get
[**CreativeUpdate**](CreativesApi.md#CreativeUpdate) | **Patch** /creatives/{crv_id} | update



## CreativeCreate

> CreativeResponse CreativeCreate(ctx).CreativeWritable(creativeWritable).XLangOutput(xLangOutput).Execute()

create



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
    creativeWritable := *openapiclient.NewCreativeWritable(interface{}(123), "ResourceType_example", "CampaignId_example") // CreativeWritable | 
    xLangOutput := "xLangOutput_example" // string | * `native` - Translate response to the native language of the country in the request * `match` - match the response to the language in the request  Default response is in English.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativesApi.CreativeCreate(context.Background()).CreativeWritable(creativeWritable).XLangOutput(xLangOutput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesApi.CreativeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreativeCreate`: CreativeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreativesApi.CreativeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreativeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creativeWritable** | [**CreativeWritable**](CreativeWritable.md) |  | 
 **xLangOutput** | **string** | * &#x60;native&#x60; - Translate response to the native language of the country in the request * &#x60;match&#x60; - match the response to the language in the request  Default response is in English.  | 

### Return type

[**CreativeResponse**](CreativeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreativeRetrieve

> CreativeResponse CreativeRetrieve(ctx, crvId).Execute()

get



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
    crvId := "crvId_example" // string | id of the creative

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativesApi.CreativeRetrieve(context.Background(), crvId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesApi.CreativeRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreativeRetrieve`: CreativeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreativesApi.CreativeRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crvId** | **string** | id of the creative | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreativeRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreativeResponse**](CreativeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreativeUpdate

> CreativeResponse CreativeUpdate(ctx, crvId).CreativePatch(creativePatch).Execute()

update



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
    crvId := "crvId_example" // string | id of the creative
    creativePatch := *openapiclient.NewCreativePatch() // CreativePatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativesApi.CreativeUpdate(context.Background(), crvId).CreativePatch(creativePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesApi.CreativeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreativeUpdate`: CreativeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreativesApi.CreativeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crvId** | **string** | id of the creative | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreativeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creativePatch** | [**CreativePatch**](CreativePatch.md) |  | 

### Return type

[**CreativeResponse**](CreativeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

