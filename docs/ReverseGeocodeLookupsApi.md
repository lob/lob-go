# \ReverseGeocodeLookupsApi

All URIs are relative to *https://api.lob.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Lookup**](ReverseGeocodeLookupsApi.md#Lookup) | **Post** /us_reverse_geocode_lookups | lookup



## ReverseGeocodeLookup

> ReverseGeocode Lookup(ctx).Location(location).Size(size).Execute()

lookup



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
    location := *openapiclient.NewLocation(NullableFloat32(123), NullableFloat32(123)) // Location | 
    size := int32(56) // int32 | Determines the number of locations returned. Possible values are between 1 and 50 and any number higher will be rounded down to 50. Default size is a list of 5 reverse geocoded locations. (optional) (default to 5)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReverseGeocodeLookupsApi.Lookup(context.Background()).Location(location).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseGeocodeLookupsApi.Lookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Lookup`: ReverseGeocode
    fmt.Fprintf(os.Stdout, "Response from `ReverseGeocodeLookupsApi.Lookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReverseGeocodeLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | [**Location**](Location.md) |  | 
 **size** | **int32** | Determines the number of locations returned. Possible values are between 1 and 50 and any number higher will be rounded down to 50. Default size is a list of 5 reverse geocoded locations. | [default to 5]

### Return type

[**ReverseGeocode**](ReverseGeocode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

