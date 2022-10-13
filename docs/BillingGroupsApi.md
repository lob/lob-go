# \BillingGroupsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingGroupCreate**](BillingGroupsApi.md#BillingGroupCreate) | **Post** /billing_groups | create
[**BillingGroupRetrieve**](BillingGroupsApi.md#BillingGroupRetrieve) | **Get** /billing_groups/{bg_id} | get
[**BillingGroupUpdate**](BillingGroupsApi.md#BillingGroupUpdate) | **Post** /billing_groups/{bg_id} | update
[**BillingGroupsList**](BillingGroupsApi.md#BillingGroupsList) | **Get** /billing_groups | list



## BillingGroupCreate

> BillingGroup BillingGroupCreate(ctx).BillingGroupEditable(billingGroupEditable).Execute()

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
    billingGroupEditable := *openapiclient.NewBillingGroupEditable("Name_example") // BillingGroupEditable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingGroupsApi.BillingGroupCreate(context.Background()).BillingGroupEditable(billingGroupEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.BillingGroupCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGroupCreate`: BillingGroup
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.BillingGroupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingGroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingGroupEditable** | [**BillingGroupEditable**](BillingGroupEditable.md) |  | 

### Return type

[**BillingGroup**](BillingGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGroupRetrieve

> BillingGroup BillingGroupRetrieve(ctx, bgId).Execute()

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
    bgId := "bgId_example" // string | id of the billing_group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingGroupsApi.BillingGroupRetrieve(context.Background(), bgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.BillingGroupRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGroupRetrieve`: BillingGroup
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.BillingGroupRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgId** | **string** | id of the billing_group | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGroupRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingGroup**](BillingGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGroupUpdate

> BillingGroup BillingGroupUpdate(ctx, bgId).BillingGroupEditable(billingGroupEditable).Execute()

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
    bgId := "bgId_example" // string | id of the billing_group
    billingGroupEditable := *openapiclient.NewBillingGroupEditable("Name_example") // BillingGroupEditable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingGroupsApi.BillingGroupUpdate(context.Background(), bgId).BillingGroupEditable(billingGroupEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.BillingGroupUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGroupUpdate`: BillingGroup
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.BillingGroupUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgId** | **string** | id of the billing_group | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGroupUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingGroupEditable** | [**BillingGroupEditable**](BillingGroupEditable.md) |  | 

### Return type

[**BillingGroup**](BillingGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGroupsList

> BillingGroupList BillingGroupsList(ctx).Limit(limit).Offset(offset).Include(include).DateCreated(dateCreated).DateModified(dateModified).SortByDateModified(sortByDateModified).Execute()

list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | How many results to return. (optional) (default to 10)
    offset := int32(56) // int32 | An integer that designates the offset at which to begin returning results. Defaults to 0. (optional) (default to 0)
    include := []string{"Inner_example"} // []string | Request that the response include the total count by specifying `include[]=total_count`.  (optional)
    dateCreated := map[string]time.Time{"key": time.Now()} // map[string]time.Time | Filter by date created. (optional)
    dateModified := map[string]string{"key": "Inner_example"} // map[string]string | Filter by date modified. (optional)
    sortByDateModified := *openapiclient.NewSortByDateModified() // SortByDateModified | Sorts items by ascending or descending dates. Use either `date_created` or `date_modfied`, not both.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingGroupsApi.BillingGroupsList(context.Background()).Limit(limit).Offset(offset).Include(include).DateCreated(dateCreated).DateModified(dateModified).SortByDateModified(sortByDateModified).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.BillingGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGroupsList`: BillingGroupList
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.BillingGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many results to return. | [default to 10]
 **offset** | **int32** | An integer that designates the offset at which to begin returning results. Defaults to 0. | [default to 0]
 **include** | **[]string** | Request that the response include the total count by specifying &#x60;include[]&#x3D;total_count&#x60;.  | 
 **dateCreated** | [**map[string]time.Time**](time.Time.md) | Filter by date created. | 
 **dateModified** | **map[string]string** | Filter by date modified. | 
 **sortByDateModified** | [**SortByDateModified**](SortByDateModified.md) | Sorts items by ascending or descending dates. Use either &#x60;date_created&#x60; or &#x60;date_modfied&#x60;, not both.  | 

### Return type

[**BillingGroupList**](BillingGroupList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

