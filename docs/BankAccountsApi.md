# \BankAccountsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](BankAccountsApi.md#create) | **Post** /bank_accounts | create
[**delete**](BankAccountsApi.md#delete) | **Delete** /bank_accounts/{bank_id} | delete
[**get**](BankAccountsApi.md#get) | **Get** /bank_accounts/{bank_id} | get
[**verify**](BankAccountsApi.md#verify) | **Post** /bank_accounts/{bank_id}/verify | verify
[**list**](BankAccountsApi.md#list) | **Get** /bank_accounts | list



## BankAccountCreate

> BankAccount create(ctx).BankAccountWritable(bankAccountWritable).Execute()

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
    resp, r, err := apiClient.BankAccountsApi.create(context.Background()).BankAccountWritable(bankAccountWritable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `create`: BankAccount
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.create`: %v\n", resp)
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

> BankAccountDeletion delete(ctx, bankId).Execute()

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
    resp, r, err := apiClient.BankAccountsApi.delete(context.Background(), bankId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `delete`: BankAccountDeletion
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.delete`: %v\n", resp)
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

> BankAccount get(ctx, bankId).Execute()

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
    resp, r, err := apiClient.BankAccountsApi.get(context.Background(), bankId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `get`: BankAccount
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.get`: %v\n", resp)
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

> BankAccount verify(ctx, bankId).BankAccountVerify(bankAccountVerify).Execute()

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
    resp, r, err := apiClient.BankAccountsApi.verify(context.Background(), bankId).BankAccountVerify(bankAccountVerify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.verify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `verify`: BankAccount
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.verify`: %v\n", resp)
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

> BankAccountList list(ctx).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Execute()

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
    resp, r, err := apiClient.BankAccountsApi.list(context.Background()).Limit(limit).Before(before).After(after).Include(include).DateCreated(dateCreated).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankAccountsApi.list``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `list`: BankAccountList
    fmt.Fprintf(os.Stdout, "Response from `BankAccountsApi.list`: %v\n", resp)
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

