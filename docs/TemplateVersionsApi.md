# \TemplateVersionsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplateVersion**](TemplateVersionsApi.md#CreateTemplateVersion) | **Post** /templates/{tmpl_id}/versions | create
[**TemplateVersionDelete**](TemplateVersionsApi.md#TemplateVersionDelete) | **Delete** /templates/{tmpl_id}/versions/{vrsn_id} | delete
[**TemplateVersionRetrieve**](TemplateVersionsApi.md#TemplateVersionRetrieve) | **Get** /templates/{tmpl_id}/versions/{vrsn_id} | get
[**TemplateVersionUpdate**](TemplateVersionsApi.md#TemplateVersionUpdate) | **Post** /templates/{tmpl_id}/versions/{vrsn_id} | update
[**TemplateVersionsList**](TemplateVersionsApi.md#TemplateVersionsList) | **Get** /templates/{tmpl_id}/versions | list



## CreateTemplateVersion

> TemplateVersion CreateTemplateVersion(ctx, tmplId).TemplateVersionWritable(templateVersionWritable).Execute()

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
    tmplId := "tmplId_example" // string | The ID of the template the new version will be attached to
    templateVersionWritable := *openapiclient.NewTemplateVersionWritable("Html_example") // TemplateVersionWritable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateVersionsApi.CreateTemplateVersion(context.Background(), tmplId).TemplateVersionWritable(templateVersionWritable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateVersionsApi.CreateTemplateVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTemplateVersion`: TemplateVersion
    fmt.Fprintf(os.Stdout, "Response from `TemplateVersionsApi.CreateTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmplId** | **string** | The ID of the template the new version will be attached to | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateVersionWritable** | [**TemplateVersionWritable**](TemplateVersionWritable.md) |  | 

### Return type

[**TemplateVersion**](TemplateVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateVersionDelete

> TemplateVersionDeletion TemplateVersionDelete(ctx, tmplId, vrsnId).Execute()

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
    tmplId := "tmplId_example" // string | The ID of the template to which the version belongs.
    vrsnId := "vrsnId_example" // string | id of the template_version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateVersionsApi.TemplateVersionDelete(context.Background(), tmplId, vrsnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateVersionsApi.TemplateVersionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateVersionDelete`: TemplateVersionDeletion
    fmt.Fprintf(os.Stdout, "Response from `TemplateVersionsApi.TemplateVersionDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmplId** | **string** | The ID of the template to which the version belongs. | 
**vrsnId** | **string** | id of the template_version | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateVersionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TemplateVersionDeletion**](TemplateVersionDeletion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateVersionRetrieve

> TemplateVersion TemplateVersionRetrieve(ctx, tmplId, vrsnId).Execute()

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
    tmplId := "tmplId_example" // string | The ID of the template to which the version belongs.
    vrsnId := "vrsnId_example" // string | id of the template_version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateVersionsApi.TemplateVersionRetrieve(context.Background(), tmplId, vrsnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateVersionsApi.TemplateVersionRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateVersionRetrieve`: TemplateVersion
    fmt.Fprintf(os.Stdout, "Response from `TemplateVersionsApi.TemplateVersionRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmplId** | **string** | The ID of the template to which the version belongs. | 
**vrsnId** | **string** | id of the template_version | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateVersionRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TemplateVersion**](TemplateVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateVersionUpdate

> TemplateVersion TemplateVersionUpdate(ctx, tmplId, vrsnId).TemplateVersionUpdatable(templateVersionUpdatable).Execute()

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
    tmplId := "tmplId_example" // string | The ID of the template to which the version belongs.
    vrsnId := "vrsnId_example" // string | id of the template_version
    templateVersionUpdatable := *openapiclient.NewTemplateVersionUpdatable() // TemplateVersionUpdatable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateVersionsApi.TemplateVersionUpdate(context.Background(), tmplId, vrsnId).TemplateVersionUpdatable(templateVersionUpdatable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateVersionsApi.TemplateVersionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateVersionUpdate`: TemplateVersion
    fmt.Fprintf(os.Stdout, "Response from `TemplateVersionsApi.TemplateVersionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmplId** | **string** | The ID of the template to which the version belongs. | 
**vrsnId** | **string** | id of the template_version | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateVersionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **templateVersionUpdatable** | [**TemplateVersionUpdatable**](TemplateVersionUpdatable.md) |  | 

### Return type

[**TemplateVersion**](TemplateVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateVersionsList

> TemplateVersionList TemplateVersionsList(ctx, tmplId).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Execute()

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
    tmplId := "tmplId_example" // string | The ID of the template associated with the retrieved versions
    limit := int32(56) // int32 | How many results to return. (optional) (default to 10)
    before := "before_example" // string | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the `previous_url` field in the return response.  (optional)
    after := "after_example" // string | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the `next_url` field in the return response.  (optional)
    include := []string{"Inner_example"} // []string | Request that the response include the total count by specifying `include[]=total_count`.  (optional)
    dateCreated := map[string]time.Time{"key": time.Now()} // map[string]time.Time | Filter by date created. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateVersionsApi.TemplateVersionsList(context.Background(), tmplId).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateVersionsApi.TemplateVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateVersionsList`: TemplateVersionList
    fmt.Fprintf(os.Stdout, "Response from `TemplateVersionsApi.TemplateVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmplId** | **string** | The ID of the template associated with the retrieved versions | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | How many results to return. | [default to 10]
 **before** | **string** | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the &#x60;previous_url&#x60; field in the return response.  | 
 **after** | **string** | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the &#x60;next_url&#x60; field in the return response.  | 
 **include** | **[]string** | Request that the response include the total count by specifying &#x60;include[]&#x3D;total_count&#x60;.  | 
 **dateCreated** | [**map[string]time.Time**](time.Time.md) | Filter by date created. | 

### Return type

[**TemplateVersionList**](TemplateVersionList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
