# \UploadsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get_export**](UploadsApi.md#Get_export) | **Get** /uploads/{upl_id}/exports/{ex_id} | get_export
[**Get**](UploadsApi.md#Get) | **Get** /uploads/{upl_id} | get
[**Create**](UploadsApi.md#Create) | **Post** /uploads | create
[**Delete**](UploadsApi.md#Delete) | **Delete** /uploads/{upl_id} | delete
[**Create_export**](UploadsApi.md#Create_export) | **Post** /uploads/{upl_id}/exports | create_export
[**Upload_file**](UploadsApi.md#Upload_file) | **Post** /uploads/{upl_id}/file | upload_file
[**Update**](UploadsApi.md#Update) | **Patch** /uploads/{upl_id} | update
[**List**](UploadsApi.md#List) | **Get** /uploads | list



## ExportRetrieve

> Export Get_export(ctx, uplId, exId).Execute()

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
    resp, r, err := apiClient.UploadsApi.Get_export(context.Background(), uplId, exId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.Get_export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get_export`: Export
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.Get_export`: %v\n", resp)
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


## Get

> Upload Get(ctx, uplId).Execute()

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
    uplId := "uplId_example" // string | id of the upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.Get(context.Background(), uplId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Upload
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uplId** | **string** | id of the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


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


## UploadCreate

> Upload Create(ctx).UploadWritable(uploadWritable).Execute()

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
    uploadWritable := *openapiclient.NewUploadWritable(string(123)) // UploadWritable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.Create(context.Background()).UploadWritable(uploadWritable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Upload
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.Create`: %v\n", resp)
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

> Delete(ctx, uplId).Execute()

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
    uplId := "uplId_example" // string | id of the upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.Delete(context.Background(), uplId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.Delete``: %v\n", err)
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

> UploadCreateExport Create_export(ctx, uplId).ExportModel(exportModel).Execute()

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
    resp, r, err := apiClient.UploadsApi.Create_export(context.Background(), uplId).ExportModel(exportModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.Create_export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create_export`: UploadCreateExport
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.Create_export`: %v\n", resp)
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

> UploadFile Upload_file(ctx, uplId).File(file).Execute()

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
    resp, r, err := apiClient.UploadsApi.Upload_file(context.Background(), uplId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.Upload_file``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Upload_file`: UploadFile
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.Upload_file`: %v\n", resp)
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


## UploadUpdate

> Upload Update(ctx, uplId).UploadUpdatable(uploadUpdatable).Execute()

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
    uplId := "uplId_example" // string | id of the upload
    uploadUpdatable := *openapiclient.NewUploadUpdatable() // UploadUpdatable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.Update(context.Background(), uplId).UploadUpdatable(uploadUpdatable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Upload
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.Update`: %v\n", resp)
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

> []Upload List(ctx).CampaignId(campaignId).Execute()

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
    campaignId := "campaignId_example" // string | id of the campaign (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.List(context.Background()).CampaignId(campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []Upload
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.List`: %v\n", resp)
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

