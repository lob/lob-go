# \AddressesApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressCreate**](AddressesApi.md#AddressCreate) | **Post** /addresses | create
[**AddressDelete**](AddressesApi.md#AddressDelete) | **Delete** /addresses/{adr_id} | delete
[**AddressRetrieve**](AddressesApi.md#AddressRetrieve) | **Get** /addresses/{adr_id} | get
[**AddressesList**](AddressesApi.md#AddressesList) | **Get** /addresses | list



## AddressCreate

> Address AddressCreate(ctx).AddressEditable(addressEditable).Execute()

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
    addressEditable := *openapiclient.NewAddressEditable() // AddressEditable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.AddressCreate(context.Background()).AddressEditable(addressEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.AddressCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressCreate`: Address
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.AddressCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressEditable** | [**AddressEditable**](AddressEditable.md) |  | 

### Return type

[**Address**](Address.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressDelete

> AddressDeletion AddressDelete(ctx, adrId).Execute()

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
    adrId := "adrId_example" // string | id of the address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.AddressDelete(context.Background(), adrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.AddressDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressDelete`: AddressDeletion
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.AddressDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adrId** | **string** | id of the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressDeletion**](AddressDeletion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressRetrieve

> Address AddressRetrieve(ctx, adrId).Execute()

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
    adrId := "adrId_example" // string | id of the address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.AddressRetrieve(context.Background(), adrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.AddressRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressRetrieve`: Address
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.AddressRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adrId** | **string** | id of the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Address**](Address.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressesList

> AddressList AddressesList(ctx).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Execute()

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
    resp, r, err := apiClient.AddressesApi.AddressesList(context.Background()).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.AddressesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressesList`: AddressList
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.AddressesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many results to return. | [default to 10]
 **before** | **string** | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the &#x60;previous_url&#x60; field in the return response.  | 
 **after** | **string** | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the &#x60;next_url&#x60; field in the return response.  | 
 **include** | **[]string** | Request that the response include the total count by specifying &#x60;include[]&#x3D;total_count&#x60;.  | 
 **dateCreated** | [**map[string]time.Time**](time.Time.md) | Filter by date created. | 
 **metadata** | **map[string]string** | Filter by metadata key-value pair&#x60;. | 

### Return type

[**AddressList**](AddressList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

