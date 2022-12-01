# \BillingGroupsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](BillingGroupsApi.md#create) | **Post** /billing_groups | create
[**get**](BillingGroupsApi.md#get) | **Get** /billing_groups/{bg_id} | get
[**update**](BillingGroupsApi.md#update) | **Post** /billing_groups/{bg_id} | update
[**list**](BillingGroupsApi.md#list) | **Get** /billing_groups | list



## BillingGroupCreate

> BillingGroup create(ctx).BillingGroupEditable(billingGroupEditable).Execute()

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
    resp, r, err := apiClient.BillingGroupsApi.create(context.Background()).BillingGroupEditable(billingGroupEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `create`: BillingGroup
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.create`: %v\n", resp)
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

> BillingGroup get(ctx, bgId).Execute()

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
    resp, r, err := apiClient.BillingGroupsApi.get(context.Background(), bgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `get`: BillingGroup
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.get`: %v\n", resp)
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

> BillingGroup update(ctx, bgId).BillingGroupEditable(billingGroupEditable).Execute()

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
    resp, r, err := apiClient.BillingGroupsApi.update(context.Background(), bgId).BillingGroupEditable(billingGroupEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `update`: BillingGroup
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.update`: %v\n", resp)
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

> BillingGroupList list(ctx).Limit(limit).Offset(offset).Include(include).DateCreated(dateCreated).DateModified(dateModified).SortByDateModified(sortByDateModified).Execute()

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
    resp, r, err := apiClient.BillingGroupsApi.list(context.Background()).Limit(limit).Offset(offset).Include(include).DateCreated(dateCreated).DateModified(dateModified).SortByDateModified(sortByDateModified).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.list``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `list`: BillingGroupList
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.list`: %v\n", resp)
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

