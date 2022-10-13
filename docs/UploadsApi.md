# \UploadsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportRetrieve**](UploadsApi.md#ExportRetrieve) | **Get** /uploads/{upl_id}/exports/{ex_id} | get_export
[**UploadCreate**](UploadsApi.md#UploadCreate) | **Post** /uploads | create_upload
[**UploadDelete**](UploadsApi.md#UploadDelete) | **Delete** /uploads/{upl_id} | delete_upload
[**UploadExportCreate**](UploadsApi.md#UploadExportCreate) | **Post** /uploads/{upl_id}/exports | create_export
[**UploadFileCreate**](UploadsApi.md#UploadFileCreate) | **Post** /uploads/{upl_id}/file | upload_file
[**UploadRetrieve**](UploadsApi.md#UploadRetrieve) | **Get** /uploads/{upl_id} | get_upload
[**UploadUpdate**](UploadsApi.md#UploadUpdate) | **Patch** /uploads/{upl_id} | update_upload
[**UploadsList**](UploadsApi.md#UploadsList) | **Get** /uploads | list_upload



## ExportRetrieve

> Export ExportRetrieve(ctx, uplId, exId).Execute()

get_export



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
    uplId := "uplId_example" // string | ID of the upload
    exId := "exId_example" // string | ID of the export

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.ExportRetrieve(context.Background(), uplId, exId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.ExportRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRetrieve`: Export
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.ExportRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uplId** | **string** | ID of the upload | 
**exId** | **string** | ID of the export | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Export**](Export.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCreate

> Upload UploadCreate(ctx).UploadWritable(uploadWritable).Execute()

create_upload



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
    uploadWritable := *openapiclient.NewUploadWritable("CampaignId_example", map[string]interface{}(123)) // UploadWritable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadCreate(context.Background()).UploadWritable(uploadWritable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadCreate`: Upload
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadWritable** | [**UploadWritable**](UploadWritable.md) |  | 

### Return type

[**Upload**](Upload.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadDelete

> UploadDelete(ctx, uplId).Execute()

delete_upload



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
    uplId := "uplId_example" // string | id of the upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadDelete(context.Background(), uplId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uplId** | **string** | id of the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadExportCreate

> UploadCreateExport UploadExportCreate(ctx, uplId).ExportModel(exportModel).Execute()

create_export



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
    uplId := "uplId_example" // string | ID of the upload
    exportModel := *openapiclient.NewExportModel() // ExportModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadExportCreate(context.Background(), uplId).ExportModel(exportModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadExportCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadExportCreate`: UploadCreateExport
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadExportCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uplId** | **string** | ID of the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadExportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportModel** | [**ExportModel**](ExportModel.md) |  | 

### Return type

[**UploadCreateExport**](UploadCreateExport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileCreate

> UploadFile UploadFileCreate(ctx, uplId).File(file).Execute()

upload_file



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
    uplId := "uplId_example" // string | ID of the upload
    file := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadFileCreate(context.Background(), uplId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadFileCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFileCreate`: UploadFile
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadFileCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uplId** | **string** | ID of the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**interface{}**](interface{}.md) |  | 

### Return type

[**UploadFile**](UploadFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadRetrieve

> Upload UploadRetrieve(ctx, uplId).Execute()

get_upload



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
    uplId := "uplId_example" // string | id of the upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadRetrieve(context.Background(), uplId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadRetrieve`: Upload
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uplId** | **string** | id of the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Upload**](Upload.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadUpdate

> Upload UploadUpdate(ctx, uplId).UploadUpdatable(uploadUpdatable).Execute()

update_upload



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
    uplId := "uplId_example" // string | id of the upload
    uploadUpdatable := *openapiclient.NewUploadUpdatable() // UploadUpdatable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadUpdate(context.Background(), uplId).UploadUpdatable(uploadUpdatable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadUpdate`: Upload
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uplId** | **string** | id of the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadUpdatable** | [**UploadUpdatable**](UploadUpdatable.md) |  | 

### Return type

[**Upload**](Upload.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsList

> []Upload UploadsList(ctx).CampaignId(campaignId).Execute()

list_upload



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
    campaignId := "campaignId_example" // string | id of the campaign (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsList(context.Background()).CampaignId(campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsList`: []Upload
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string** | id of the campaign | 

### Return type

[**[]Upload**](Upload.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

