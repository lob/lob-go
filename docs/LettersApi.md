# \LettersApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LetterCancel**](LettersApi.md#LetterCancel) | **Delete** /letters/{ltr_id} | cancel
[**LetterCreate**](LettersApi.md#LetterCreate) | **Post** /letters | create
[**LetterRetrieve**](LettersApi.md#LetterRetrieve) | **Get** /letters/{ltr_id} | get
[**LettersList**](LettersApi.md#LettersList) | **Get** /letters | list



## LetterCancel

> LetterDeletion LetterCancel(ctx, ltrId).Execute()

cancel



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
    ltrId := "ltrId_example" // string | id of the letter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LettersApi.LetterCancel(context.Background(), ltrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LettersApi.LetterCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LetterCancel`: LetterDeletion
    fmt.Fprintf(os.Stdout, "Response from `LettersApi.LetterCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ltrId** | **string** | id of the letter | 

### Other Parameters

Other parameters are passed through a pointer to a apiLetterCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LetterDeletion**](LetterDeletion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LetterCreate

> Letter LetterCreate(ctx).LetterEditable(letterEditable).IdempotencyKey(idempotencyKey).Execute()

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
    letterEditable := *openapiclient.NewLetterEditable(false, "To_example", "From_example", "File_example") // LetterEditable | 
    idempotencyKey := "idempotencyKey_example" // string | A string of no longer than 256 characters that uniquely identifies this resource. For more help integrating idempotency keys, refer to our [implementation guide](https://www.lob.com/guides#idempotent_request).  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LettersApi.LetterCreate(context.Background()).LetterEditable(letterEditable).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LettersApi.LetterCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LetterCreate`: Letter
    fmt.Fprintf(os.Stdout, "Response from `LettersApi.LetterCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLetterCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **letterEditable** | [**LetterEditable**](LetterEditable.md) |  | 
 **idempotencyKey** | **string** | A string of no longer than 256 characters that uniquely identifies this resource. For more help integrating idempotency keys, refer to our [implementation guide](https://www.lob.com/guides#idempotent_request).  | 

### Return type

[**Letter**](Letter.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LetterRetrieve

> Letter LetterRetrieve(ctx, ltrId).Execute()

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
    ltrId := "ltrId_example" // string | id of the letter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LettersApi.LetterRetrieve(context.Background(), ltrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LettersApi.LetterRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LetterRetrieve`: Letter
    fmt.Fprintf(os.Stdout, "Response from `LettersApi.LetterRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ltrId** | **string** | id of the letter | 

### Other Parameters

Other parameters are passed through a pointer to a apiLetterRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Letter**](Letter.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LettersList

> LetterList LettersList(ctx).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Color(color).Scheduled(scheduled).SendDate(sendDate).MailType(mailType).SortBy(sortBy).Execute()

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
    color := true // bool | Set to `true` to return only color letters. Set to `false` to return only black & white letters. (optional)
    scheduled := true // bool | * `true` - only return orders (past or future) where `send_date` is greater than `date_created` * `false` - only return orders where `send_date` is equal to `date_created`  (optional)
    sendDate := map[string]string{"key": "Inner_example"} // map[string]string | Filter by date sent. (optional)
    mailType := openapiclient.mail_type("usps_first_class") // MailType | A string designating the mail postage type: * `usps_first_class` - (default) * `usps_standard` - a [cheaper option](https://lob.com/pricing/print-mail#compare) which is less predictable and takes longer to deliver. `usps_standard` cannot be used with `4x6` postcards or for any postcards sent outside of the United States.  (optional) (default to "usps_first_class")
    sortBy := *openapiclient.NewSortBy3() // SortBy3 | Sorts items by ascending or descending dates. Use either `date_created` or `send_date`, not both.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LettersApi.LettersList(context.Background()).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Color(color).Scheduled(scheduled).SendDate(sendDate).MailType(mailType).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LettersApi.LettersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LettersList`: LetterList
    fmt.Fprintf(os.Stdout, "Response from `LettersApi.LettersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLettersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many results to return. | [default to 10]
 **before** | **string** | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the &#x60;previous_url&#x60; field in the return response.  | 
 **after** | **string** | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the &#x60;next_url&#x60; field in the return response.  | 
 **include** | **[]string** | Request that the response include the total count by specifying &#x60;include[]&#x3D;total_count&#x60;.  | 
 **dateCreated** | [**map[string]time.Time**](time.Time.md) | Filter by date created. | 
 **metadata** | **map[string]string** | Filter by metadata key-value pair&#x60;. | 
 **color** | **bool** | Set to &#x60;true&#x60; to return only color letters. Set to &#x60;false&#x60; to return only black &amp; white letters. | 
 **scheduled** | **bool** | * &#x60;true&#x60; - only return orders (past or future) where &#x60;send_date&#x60; is greater than &#x60;date_created&#x60; * &#x60;false&#x60; - only return orders where &#x60;send_date&#x60; is equal to &#x60;date_created&#x60;  | 
 **sendDate** | **map[string]string** | Filter by date sent. | 
 **mailType** | [**MailType**](MailType.md) | A string designating the mail postage type: * &#x60;usps_first_class&#x60; - (default) * &#x60;usps_standard&#x60; - a [cheaper option](https://lob.com/pricing/print-mail#compare) which is less predictable and takes longer to deliver. &#x60;usps_standard&#x60; cannot be used with &#x60;4x6&#x60; postcards or for any postcards sent outside of the United States.  | [default to &quot;usps_first_class&quot;]
 **sortBy** | [**SortBy3**](SortBy3.md) | Sorts items by ascending or descending dates. Use either &#x60;date_created&#x60; or &#x60;send_date&#x60;, not both.  | 

### Return type

[**LetterList**](LetterList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

