# \BillingGroupsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](BillingGroupsApi.md#Create) | **Post** /billing_groups | create
[**Get**](BillingGroupsApi.md#Get) | **Get** /billing_groups/{bg_id} | get
[**Update**](BillingGroupsApi.md#Update) | **Post** /billing_groups/{bg_id} | update
[**List**](BillingGroupsApi.md#List) | **Get** /billing_groups | list



## BillingGroupCreate

> BillingGroup Create(ctx).BillingGroupEditable(billingGroupEditable).Execute()

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
    resp, r, err := apiClient.BillingGroupsApi.Create(context.Background()).BillingGroupEditable(billingGroupEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: BillingGroup
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.Create`: %v\n", resp)
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

> BillingGroup Get(ctx, bgId).Execute()

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
    resp, r, err := apiClient.BillingGroupsApi.Get(context.Background(), bgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: BillingGroup
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.Get`: %v\n", resp)
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

> BillingGroup Update(ctx, bgId).BillingGroupEditable(billingGroupEditable).Execute()

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
    resp, r, err := apiClient.BillingGroupsApi.Update(context.Background(), bgId).BillingGroupEditable(billingGroupEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: BillingGroup
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.Update`: %v\n", resp)
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

> BillingGroupList List(ctx).Limit(limit).Offset(offset).Include(include).DateCreated(dateCreated).DateModified(dateModified).SortByDateModified(sortByDateModified).Execute()

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
    resp, r, err := apiClient.BillingGroupsApi.List(context.Background()).Limit(limit).Offset(offset).Include(include).DateCreated(dateCreated).DateModified(dateModified).SortByDateModified(sortByDateModified).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: BillingGroupList
    fmt.Fprintf(os.Stdout, "Response from `BillingGroupsApi.List`: %v\n", resp)
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

