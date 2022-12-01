# \BuckslipOrdersApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](BuckslipOrdersApi.md#create) | **Post** /buckslips/{buckslip_id}/orders | create
[**get**](BuckslipOrdersApi.md#get) | **Get** /buckslips/{buckslip_id}/orders | get



## BuckslipOrderCreate

> BuckslipOrder create(ctx, buckslipId).BuckslipOrderEditable(buckslipOrderEditable).Execute()

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
    buckslipId := "buckslipId_example" // string | The ID of the buckslip to which the buckslip orders belong.
    buckslipOrderEditable := *openapiclient.NewBuckslipOrderEditable(int32(123)) // BuckslipOrderEditable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuckslipOrdersApi.create(context.Background(), buckslipId).BuckslipOrderEditable(buckslipOrderEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipOrdersApi.create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `create`: BuckslipOrder
    fmt.Fprintf(os.Stdout, "Response from `BuckslipOrdersApi.create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buckslipId** | **string** | The ID of the buckslip to which the buckslip orders belong. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuckslipOrderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buckslipOrderEditable** | [**BuckslipOrderEditable**](BuckslipOrderEditable.md) |  | 

### Return type

[**BuckslipOrder**](BuckslipOrder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuckslipOrdersRetrieve

> BuckslipOrdersList get(ctx, buckslipId).Limit(limit).Offset(offset).Execute()

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
    buckslipId := "buckslipId_example" // string | The ID of the buckslip to which the buckslip orders belong.
    limit := int32(56) // int32 | How many results to return. (optional) (default to 10)
    offset := int32(56) // int32 | An integer that designates the offset at which to begin returning results. Defaults to 0. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuckslipOrdersApi.get(context.Background(), buckslipId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipOrdersApi.get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `get`: BuckslipOrdersList
    fmt.Fprintf(os.Stdout, "Response from `BuckslipOrdersApi.get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buckslipId** | **string** | The ID of the buckslip to which the buckslip orders belong. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuckslipOrdersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | How many results to return. | [default to 10]
 **offset** | **int32** | An integer that designates the offset at which to begin returning results. Defaults to 0. | [default to 0]

### Return type

[**BuckslipOrdersList**](BuckslipOrdersList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

