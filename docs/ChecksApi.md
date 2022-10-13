# \ChecksApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCancel**](ChecksApi.md#CheckCancel) | **Delete** /checks/{chk_id} | cancel
[**CheckCreate**](ChecksApi.md#CheckCreate) | **Post** /checks | create
[**CheckRetrieve**](ChecksApi.md#CheckRetrieve) | **Get** /checks/{chk_id} | get
[**ChecksList**](ChecksApi.md#ChecksList) | **Get** /checks | list



## CheckCancel

> CheckDeletion CheckCancel(ctx, chkId).Execute()

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
    chkId := "chkId_example" // string | id of the check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.CheckCancel(context.Background(), chkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.CheckCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCancel`: CheckDeletion
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.CheckCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chkId** | **string** | id of the check | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckDeletion**](CheckDeletion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckCreate

> Check CheckCreate(ctx).CheckEditable(checkEditable).IdempotencyKey(idempotencyKey).Execute()

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
    checkEditable := *openapiclient.NewCheckEditable("From_example", "To_example", "BankAccount_example", float32(123)) // CheckEditable | 
    idempotencyKey := "idempotencyKey_example" // string | A string of no longer than 256 characters that uniquely identifies this resource. For more help integrating idempotency keys, refer to our [implementation guide](https://www.lob.com/guides#idempotent_request).  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.CheckCreate(context.Background()).CheckEditable(checkEditable).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.CheckCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCreate`: Check
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.CheckCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkEditable** | [**CheckEditable**](CheckEditable.md) |  | 
 **idempotencyKey** | **string** | A string of no longer than 256 characters that uniquely identifies this resource. For more help integrating idempotency keys, refer to our [implementation guide](https://www.lob.com/guides#idempotent_request).  | 

### Return type

[**Check**](Check.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckRetrieve

> Check CheckRetrieve(ctx, chkId).Execute()

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
    chkId := "chkId_example" // string | id of the check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.CheckRetrieve(context.Background(), chkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.CheckRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckRetrieve`: Check
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.CheckRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chkId** | **string** | id of the check | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Check**](Check.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksList

> CheckList ChecksList(ctx).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Scheduled(scheduled).SendDate(sendDate).MailType(mailType).SortBy(sortBy).Execute()

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
    scheduled := true // bool | * `true` - only return orders (past or future) where `send_date` is greater than `date_created` * `false` - only return orders where `send_date` is equal to `date_created`  (optional)
    sendDate := map[string]string{"key": "Inner_example"} // map[string]string | Filter by date sent. (optional)
    mailType := openapiclient.mail_type("usps_first_class") // MailType | A string designating the mail postage type: * `usps_first_class` - (default) * `usps_standard` - a [cheaper option](https://lob.com/pricing/print-mail#compare) which is less predictable and takes longer to deliver. `usps_standard` cannot be used with `4x6` postcards or for any postcards sent outside of the United States.  (optional) (default to "usps_first_class")
    sortBy := *openapiclient.NewSortBy3() // SortBy3 | Sorts items by ascending or descending dates. Use either `date_created` or `send_date`, not both.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksList(context.Background()).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Scheduled(scheduled).SendDate(sendDate).MailType(mailType).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksList`: CheckList
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChecksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many results to return. | [default to 10]
 **before** | **string** | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the &#x60;previous_url&#x60; field in the return response.  | 
 **after** | **string** | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the &#x60;next_url&#x60; field in the return response.  | 
 **include** | **[]string** | Request that the response include the total count by specifying &#x60;include[]&#x3D;total_count&#x60;.  | 
 **dateCreated** | [**map[string]time.Time**](time.Time.md) | Filter by date created. | 
 **metadata** | **map[string]string** | Filter by metadata key-value pair&#x60;. | 
 **scheduled** | **bool** | * &#x60;true&#x60; - only return orders (past or future) where &#x60;send_date&#x60; is greater than &#x60;date_created&#x60; * &#x60;false&#x60; - only return orders where &#x60;send_date&#x60; is equal to &#x60;date_created&#x60;  | 
 **sendDate** | **map[string]string** | Filter by date sent. | 
 **mailType** | [**MailType**](MailType.md) | A string designating the mail postage type: * &#x60;usps_first_class&#x60; - (default) * &#x60;usps_standard&#x60; - a [cheaper option](https://lob.com/pricing/print-mail#compare) which is less predictable and takes longer to deliver. &#x60;usps_standard&#x60; cannot be used with &#x60;4x6&#x60; postcards or for any postcards sent outside of the United States.  | [default to &quot;usps_first_class&quot;]
 **sortBy** | [**SortBy3**](SortBy3.md) | Sorts items by ascending or descending dates. Use either &#x60;date_created&#x60; or &#x60;send_date&#x60;, not both.  | 

### Return type

[**CheckList**](CheckList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

