# \CampaignsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CampaignCreate**](CampaignsApi.md#CampaignCreate) | **Post** /campaigns | create
[**CampaignDelete**](CampaignsApi.md#CampaignDelete) | **Delete** /campaigns/{cmp_id} | delete
[**CampaignRetrieve**](CampaignsApi.md#CampaignRetrieve) | **Get** /campaigns/{cmp_id} | get
[**CampaignUpdate**](CampaignsApi.md#CampaignUpdate) | **Patch** /campaigns/{cmp_id} | update
[**CampaignsList**](CampaignsApi.md#CampaignsList) | **Get** /campaigns | list



## CampaignCreate

> Campaign CampaignCreate(ctx).CampaignWritable(campaignWritable).XLangOutput(xLangOutput).Execute()

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
    campaignWritable := *openapiclient.NewCampaignWritable("Name_example", openapiclient.cmp_schedule_type("immediate")) // CampaignWritable | 
    xLangOutput := "xLangOutput_example" // string | * `native` - Translate response to the native language of the country in the request * `match` - match the response to the language in the request  Default response is in English.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CampaignCreate(context.Background()).CampaignWritable(campaignWritable).XLangOutput(xLangOutput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CampaignCreate`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CampaignCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCampaignCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignWritable** | [**CampaignWritable**](CampaignWritable.md) |  | 
 **xLangOutput** | **string** | * &#x60;native&#x60; - Translate response to the native language of the country in the request * &#x60;match&#x60; - match the response to the language in the request  Default response is in English.  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CampaignDelete

> CampaignDeletion CampaignDelete(ctx, cmpId).Execute()

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
    cmpId := "cmpId_example" // string | id of the campaign

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CampaignDelete(context.Background(), cmpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CampaignDelete`: CampaignDeletion
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CampaignDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cmpId** | **string** | id of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignDeletion**](CampaignDeletion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CampaignRetrieve

> Campaign CampaignRetrieve(ctx, cmpId).Execute()

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
    cmpId := "cmpId_example" // string | id of the campaign

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CampaignRetrieve(context.Background(), cmpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CampaignRetrieve`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CampaignRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cmpId** | **string** | id of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Campaign**](Campaign.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CampaignUpdate

> Campaign CampaignUpdate(ctx, cmpId).CampaignUpdatable(campaignUpdatable).Execute()

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
    cmpId := "cmpId_example" // string | id of the campaign
    campaignUpdatable := *openapiclient.NewCampaignUpdatable() // CampaignUpdatable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CampaignUpdate(context.Background(), cmpId).CampaignUpdatable(campaignUpdatable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CampaignUpdate`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CampaignUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cmpId** | **string** | id of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignUpdatable** | [**CampaignUpdatable**](CampaignUpdatable.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CampaignsList

> CampaignsList CampaignsList(ctx).Limit(limit).Include(include).Before(before).After(after).Execute()

list



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
    include := []string{"Inner_example"} // []string | Request that the response include the total count by specifying `include[]=total_count`.  (optional)
    before := "before_example" // string | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the `previous_url` field in the return response.  (optional)
    after := "after_example" // string | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the `next_url` field in the return response.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CampaignsList(context.Background()).Limit(limit).Include(include).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CampaignsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CampaignsList`: CampaignsList
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CampaignsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many results to return. | [default to 10]
 **include** | **[]string** | Request that the response include the total count by specifying &#x60;include[]&#x3D;total_count&#x60;.  | 
 **before** | **string** | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the &#x60;previous_url&#x60; field in the return response.  | 
 **after** | **string** | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the &#x60;next_url&#x60; field in the return response.  | 

### Return type

[**CampaignsList**](CampaignsList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

