# \CardOrdersApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CardOrdersApi.md#Create) | **Post** /cards/{card_id}/orders | create
[**Get**](CardOrdersApi.md#Get) | **Get** /cards/{card_id}/orders | get



## CardOrderCreate

> CardOrder Create(ctx, cardId).CardOrderEditable(cardOrderEditable).Execute()

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
    cardId := "cardId_example" // string | The ID of the card to which the card orders belong.
    cardOrderEditable := *openapiclient.NewCardOrderEditable(int32(123)) // CardOrderEditable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardOrdersApi.Create(context.Background(), cardId).CardOrderEditable(cardOrderEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardOrdersApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: CardOrder
    fmt.Fprintf(os.Stdout, "Response from `CardOrdersApi.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** | The ID of the card to which the card orders belong. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCardOrderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cardOrderEditable** | [**CardOrderEditable**](CardOrderEditable.md) |  | 

### Return type

[**CardOrder**](CardOrder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CardOrdersRetrieve

> CardOrderList Get(ctx, cardId).Limit(limit).Offset(offset).Execute()

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
    cardId := "cardId_example" // string | The ID of the card to which the card orders belong.
    limit := int32(56) // int32 | How many results to return. (optional) (default to 10)
    offset := int32(56) // int32 | An integer that designates the offset at which to begin returning results. Defaults to 0. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardOrdersApi.Get(context.Background(), cardId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardOrdersApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: CardOrderList
    fmt.Fprintf(os.Stdout, "Response from `CardOrdersApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** | The ID of the card to which the card orders belong. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCardOrdersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | How many results to return. | [default to 10]
 **offset** | **int32** | An integer that designates the offset at which to begin returning results. Defaults to 0. | [default to 0]

### Return type

[**CardOrderList**](CardOrderList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

