# \CampaignsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CampaignsApi.md#Create) | **Post** /campaigns | create
[**Delete**](CampaignsApi.md#Delete) | **Delete** /campaigns/{cmp_id} | delete
[**Get**](CampaignsApi.md#Get) | **Get** /campaigns/{cmp_id} | get
[**Update**](CampaignsApi.md#Update) | **Patch** /campaigns/{cmp_id} | update
[**List**](CampaignsApi.md#List) | **Get** /campaigns | list



## CampaignCreate

> Campaign Create(ctx).CampaignWritable(campaignWritable).XLangOutput(xLangOutput).Execute()

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
    campaignWritable := *openapiclient.NewCampaignWritable("Name_example", openapiclient.cmp_schedule_type("immediate"), "TODO") // CampaignWritable | 
    xLangOutput := "xLangOutput_example" // string | * `native` - Translate response to the native language of the country in the request * `match` - match the response to the language in the request  Default response is in English.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.Create(context.Background()).CampaignWritable(campaignWritable).XLangOutput(xLangOutput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.Create`: %v\n", resp)
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

> CampaignDeletion Delete(ctx, cmpId).Execute()

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
    resp, r, err := apiClient.CampaignsApi.Delete(context.Background(), cmpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: CampaignDeletion
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.Delete`: %v\n", resp)
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

> Campaign Get(ctx, cmpId).Execute()

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
    resp, r, err := apiClient.CampaignsApi.Get(context.Background(), cmpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.Get`: %v\n", resp)
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

> Campaign Update(ctx, cmpId).CampaignUpdatable(campaignUpdatable).Execute()

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
    resp, r, err := apiClient.CampaignsApi.Update(context.Background(), cmpId).CampaignUpdatable(campaignUpdatable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.Update`: %v\n", resp)
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

> CampaignsList List(ctx).Limit(limit).Include(include).Before(before).After(after).Execute()

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
    resp, r, err := apiClient.CampaignsApi.List(context.Background()).Limit(limit).Include(include).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: CampaignsList
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.List`: %v\n", resp)
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

