# \BuckslipsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](BuckslipsApi.md#create) | **Post** /buckslips | create
[**delete**](BuckslipsApi.md#delete) | **Delete** /buckslips/{buckslip_id} | delete
[**get**](BuckslipsApi.md#get) | **Get** /buckslips/{buckslip_id} | get
[**update**](BuckslipsApi.md#update) | **Patch** /buckslips/{buckslip_id} | update
[**List**](BuckslipsApi.md#List) | **Get** /buckslips | List



## BuckslipCreate

> Buckslip create(ctx).BuckslipEditable(buckslipEditable).Execute()

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
    buckslipEditable := *openapiclient.NewBuckslipEditable("Front_example") // BuckslipEditable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuckslipsApi.create(context.Background()).BuckslipEditable(buckslipEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipsApi.create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `create`: Buckslip
    fmt.Fprintf(os.Stdout, "Response from `BuckslipsApi.create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuckslipCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buckslipEditable** | [**BuckslipEditable**](BuckslipEditable.md) |  | 

### Return type

[**Buckslip**](Buckslip.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuckslipDelete

> BuckslipDeletion delete(ctx, buckslipId).Execute()

delete



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
    buckslipId := "buckslipId_example" // string | id of the buckslip

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuckslipsApi.delete(context.Background(), buckslipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipsApi.delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `delete`: BuckslipDeletion
    fmt.Fprintf(os.Stdout, "Response from `BuckslipsApi.delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buckslipId** | **string** | id of the buckslip | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuckslipDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuckslipDeletion**](BuckslipDeletion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuckslipRetrieve

> Buckslip get(ctx, buckslipId).Execute()

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
    buckslipId := "buckslipId_example" // string | id of the buckslip

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuckslipsApi.get(context.Background(), buckslipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipsApi.get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `get`: Buckslip
    fmt.Fprintf(os.Stdout, "Response from `BuckslipsApi.get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buckslipId** | **string** | id of the buckslip | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuckslipRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Buckslip**](Buckslip.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuckslipUpdate

> Buckslip update(ctx, buckslipId).BuckslipUpdatable(buckslipUpdatable).Execute()

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
    buckslipId := "buckslipId_example" // string | id of the buckslip
    buckslipUpdatable := *openapiclient.NewBuckslipUpdatable() // BuckslipUpdatable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuckslipsApi.update(context.Background(), buckslipId).BuckslipUpdatable(buckslipUpdatable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipsApi.update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `update`: Buckslip
    fmt.Fprintf(os.Stdout, "Response from `BuckslipsApi.update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buckslipId** | **string** | id of the buckslip | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuckslipUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buckslipUpdatable** | [**BuckslipUpdatable**](BuckslipUpdatable.md) |  | 

### Return type

[**Buckslip**](Buckslip.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuckslipsList

> BuckslipsList List(ctx).Limit(limit).Before(before).After(after).Include(include).Execute()

List



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
    limit := int32(56) // int32 | How many results to return. (optional) (default to 10)
    before := "before_example" // string | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the `previous_url` field in the return response.  (optional)
    after := "after_example" // string | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the `next_url` field in the return response.  (optional)
    include := []string{"Inner_example"} // []string | Request that the response include the total count by specifying `include[]=total_count`.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuckslipsApi.List(context.Background()).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: BuckslipsList
    fmt.Fprintf(os.Stdout, "Response from `BuckslipsApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuckslipsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many results to return. | [default to 10]
 **before** | **string** | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the &#x60;previous_url&#x60; field in the return response.  | 
 **after** | **string** | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the &#x60;next_url&#x60; field in the return response.  | 
 **include** | **[]string** | Request that the response include the total count by specifying &#x60;include[]&#x3D;total_count&#x60;.  | 

### Return type

[**BuckslipsList**](BuckslipsList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

