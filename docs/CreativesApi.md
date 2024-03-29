# \CreativesApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CreativesApi.md#Create) | **Post** /creatives | create
[**Get**](CreativesApi.md#Get) | **Get** /creatives/{crv_id} | get
[**Update**](CreativesApi.md#Update) | **Patch** /creatives/{crv_id} | update



## CreativeCreate

> CreativeResponse Create(ctx).CreativeWritable(creativeWritable).XLangOutput(xLangOutput).Execute()

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
    resp, r, err := apiClient.CreativesApi.Create(context.Background()).CreativeWritable(creativeWritable).XLangOutput(xLangOutput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: CreativeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreativesApi.Create`: %v\n", resp)
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

> CreativeResponse Get(ctx, crvId).Execute()

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
    resp, r, err := apiClient.CreativesApi.Get(context.Background(), crvId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: CreativeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreativesApi.Get`: %v\n", resp)
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

> CreativeResponse Update(ctx, crvId).CreativePatch(creativePatch).Execute()

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
    resp, r, err := apiClient.CreativesApi.Update(context.Background(), crvId).CreativePatch(creativePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: CreativeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreativesApi.Update`: %v\n", resp)
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

