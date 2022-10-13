# \BankAccountsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BankAccountCreate**](BankAccountsApi.md#BankAccountCreate) | **Post** /bank_accounts | create
[**BankAccountDelete**](BankAccountsApi.md#BankAccountDelete) | **Delete** /bank_accounts/{bank_id} | delete
[**BankAccountRetrieve**](BankAccountsApi.md#BankAccountRetrieve) | **Get** /bank_accounts/{bank_id} | get
[**BankAccountVerify**](BankAccountsApi.md#BankAccountVerify) | **Post** /bank_accounts/{bank_id}/verify | verify
[**BankAccountsList**](BankAccountsApi.md#BankAccountsList) | **Get** /bank_accounts | list



## BankAccountCreate

> BankAccount BankAccountCreate(ctx).BankAccountWritable(bankAccountWritable).Execute()

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
    bankAccountWritable := *openapiclient.NewBankAccountWritable("RoutingNumber_example", "AccountNumber_example", openapiclient.bank_type_enum("company"), "Signatory_example") // BankAccountWritable | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BankAccountsApi.BankAccountCreate(context.Background()).BankAccountWritable(bankAccountWritable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.BankAccountCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankAccountCreate`: BankAccount
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.BankAccountCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankAccountCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankAccountWritable** | [**BankAccountWritable**](BankAccountWritable.md) |  | 

### Return type

[**BankAccount**](BankAccount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankAccountDelete

> BankAccountDeletion BankAccountDelete(ctx, bankId).Execute()

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
    bankId := "bankId_example" // string | id of the bank account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BankAccountsApi.BankAccountDelete(context.Background(), bankId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.BankAccountDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankAccountDelete`: BankAccountDeletion
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.BankAccountDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** | id of the bank account | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankAccountDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankAccountDeletion**](BankAccountDeletion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankAccountRetrieve

> BankAccount BankAccountRetrieve(ctx, bankId).Execute()

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
    bankId := "bankId_example" // string | id of the bank account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BankAccountsApi.BankAccountRetrieve(context.Background(), bankId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.BankAccountRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankAccountRetrieve`: BankAccount
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.BankAccountRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** | id of the bank account | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankAccountRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankAccount**](BankAccount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankAccountVerify

> BankAccount BankAccountVerify(ctx, bankId).BankAccountVerify(bankAccountVerify).Execute()

verify



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
    bankId := "bankId_example" // string | id of the bank account to be verified
    bankAccountVerify := *openapiclient.NewBankAccountVerify([]int32{int32(123)}) // BankAccountVerify | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BankAccountsApi.BankAccountVerify(context.Background(), bankId).BankAccountVerify(bankAccountVerify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.BankAccountVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankAccountVerify`: BankAccount
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.BankAccountVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** | id of the bank account to be verified | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankAccountVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bankAccountVerify** | [**BankAccountVerify**](BankAccountVerify.md) |  | 

### Return type

[**BankAccount**](BankAccount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankAccountsList

> BankAccountList BankAccountsList(ctx).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Execute()

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
    resp, r, err := apiClient.BankAccountsApi.BankAccountsList(context.Background()).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.BankAccountsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankAccountsList`: BankAccountList
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.BankAccountsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankAccountsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many results to return. | [default to 10]
 **before** | **string** | A reference to a list entry used for paginating to the previous set of entries. This field is pre-populated in the &#x60;previous_url&#x60; field in the return response.  | 
 **after** | **string** | A reference to a list entry used for paginating to the next set of entries. This field is pre-populated in the &#x60;next_url&#x60; field in the return response.  | 
 **include** | **[]string** | Request that the response include the total count by specifying &#x60;include[]&#x3D;total_count&#x60;.  | 
 **dateCreated** | [**map[string]time.Time**](time.Time.md) | Filter by date created. | 
 **metadata** | **map[string]string** | Filter by metadata key-value pair&#x60;. | 

### Return type

[**BankAccountList**](BankAccountList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

