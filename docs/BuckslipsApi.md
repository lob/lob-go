# \BuckslipsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](BuckslipsApi.md#Create) | **Post** /buckslips | create
[**Delete**](BuckslipsApi.md#Delete) | **Delete** /buckslips/{buckslip_id} | delete
[**Get**](BuckslipsApi.md#Get) | **Get** /buckslips/{buckslip_id} | get
[**Update**](BuckslipsApi.md#Update) | **Patch** /buckslips/{buckslip_id} | update
[**List**](BuckslipsApi.md#List) | **Get** /buckslips | List



## BuckslipCreate

> Buckslip Create(ctx).BuckslipEditable(buckslipEditable).Execute()

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
    resp, r, err := apiClient.BuckslipsApi.Create(context.Background()).BuckslipEditable(buckslipEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Buckslip
    fmt.Fprintf(os.Stdout, "Response from `BuckslipsApi.Create`: %v\n", resp)
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

> BuckslipDeletion Delete(ctx, buckslipId).Execute()

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
    resp, r, err := apiClient.BuckslipsApi.Delete(context.Background(), buckslipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: BuckslipDeletion
    fmt.Fprintf(os.Stdout, "Response from `BuckslipsApi.Delete`: %v\n", resp)
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

> Buckslip Get(ctx, buckslipId).Execute()

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
    resp, r, err := apiClient.BuckslipsApi.Get(context.Background(), buckslipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Buckslip
    fmt.Fprintf(os.Stdout, "Response from `BuckslipsApi.Get`: %v\n", resp)
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

> Buckslip Update(ctx, buckslipId).BuckslipUpdatable(buckslipUpdatable).Execute()

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
    resp, r, err := apiClient.BuckslipsApi.Update(context.Background(), buckslipId).BuckslipUpdatable(buckslipUpdatable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuckslipsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Buckslip
    fmt.Fprintf(os.Stdout, "Response from `BuckslipsApi.Update`: %v\n", resp)
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

