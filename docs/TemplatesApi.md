# \TemplatesApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplate**](TemplatesApi.md#CreateTemplate) | **Post** /templates | create
[**TemplateDelete**](TemplatesApi.md#TemplateDelete) | **Delete** /templates/{tmpl_id} | delete
[**TemplateRetrieve**](TemplatesApi.md#TemplateRetrieve) | **Get** /templates/{tmpl_id} | get
[**TemplateUpdate**](TemplatesApi.md#TemplateUpdate) | **Post** /templates/{tmpl_id} | update
[**TemplatesList**](TemplatesApi.md#TemplatesList) | **Get** /templates | list



## CreateTemplate

> Template CreateTemplate(ctx).TemplateWritable(templateWritable).Execute()

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
    templateWritable := *openapiclient.NewTemplateWritable("Html_example") // TemplateWritable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.CreateTemplate(context.Background()).TemplateWritable(templateWritable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.CreateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTemplate`: Template
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateWritable** | [**TemplateWritable**](TemplateWritable.md) |  | 

### Return type

[**Template**](Template.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateDelete

> TemplateDeletion TemplateDelete(ctx, tmplId).Execute()

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
    tmplId := "tmplId_example" // string | id of the template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateDelete(context.Background(), tmplId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateDelete`: TemplateDeletion
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmplId** | **string** | id of the template | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateDeletion**](TemplateDeletion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateRetrieve

> Template TemplateRetrieve(ctx, tmplId).Execute()

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
    tmplId := "tmplId_example" // string | id of the template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateRetrieve(context.Background(), tmplId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateRetrieve`: Template
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmplId** | **string** | id of the template | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Template**](Template.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateUpdate

> Template TemplateUpdate(ctx, tmplId).TemplateUpdate(templateUpdate).Execute()

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
    tmplId := "tmplId_example" // string | id of the template
    templateUpdate := *openapiclient.NewTemplateUpdate() // TemplateUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateUpdate(context.Background(), tmplId).TemplateUpdate(templateUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateUpdate`: Template
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmplId** | **string** | id of the template | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateUpdate** | [**TemplateUpdate**](TemplateUpdate.md) |  | 

### Return type

[**Template**](Template.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesList

> TemplateList TemplatesList(ctx).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Execute()

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
    before := "before_example" // string | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the `previous_url` field in the return response.  (optional)
    after := "after_example" // string | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the `next_url` field in the return response.  (optional)
    include := []string{"Inner_example"} // []string | Request that the response include the total count by specifying `include[]=total_count`.  (optional)
    dateCreated := map[string]time.Time{"key": time.Now()} // map[string]time.Time | Filter by date created. (optional)
    metadata := map[string]string{"key": "Inner_example"} // map[string]string | Filter by metadata key-value pair`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesList(context.Background()).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesList`: TemplateList
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many results to return. | [default to 10]
 **before** | **string** | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the &#x60;previous_url&#x60; field in the return response.  | 
 **after** | **string** | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the &#x60;next_url&#x60; field in the return response.  | 
 **include** | **[]string** | Request that the response include the total count by specifying &#x60;include[]&#x3D;total_count&#x60;.  | 
 **dateCreated** | [**map[string]time.Time**](time.Time.md) | Filter by date created. | 
 **metadata** | **map[string]string** | Filter by metadata key-value pair&#x60;. | 

### Return type

[**TemplateList**](TemplateList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

