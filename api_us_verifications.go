/*
Lob

The Lob API is organized around REST. Our API is designed to have predictable, resource-oriented URLs and uses HTTP response codes to indicate any API errors. <p> Looking for our [previous documentation](https://lob.github.io/legacy-docs/)? 

API version: 1.3.0
Contact: lob-openapi@lob.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lob

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// UsVerificationsApiService UsVerificationsApi service
type UsVerificationsApiService service

type ApiBulkUsVerificationsRequest struct {
	ctx context.Context
	ApiService *UsVerificationsApiService
	multipleComponentsList *MultipleComponentsList
	case_ *string
}

func (r ApiBulkUsVerificationsRequest) MultipleComponentsList(multipleComponentsList MultipleComponentsList) ApiBulkUsVerificationsRequest {
	r.multipleComponentsList = &multipleComponentsList
	return r
}

// Casing of the verified address. Possible values are &#x60;upper&#x60; and &#x60;proper&#x60; for uppercased (e.g. \&quot;PO BOX\&quot;) and proper-cased (e.g. \&quot;PO Box\&quot;), respectively.
func (r ApiBulkUsVerificationsRequest) Case_(case_ string) ApiBulkUsVerificationsRequest {
	r.case_ = &case_
	return r
}

func (r ApiBulkUsVerificationsRequest) Execute() (*UsVerifications, *http.Response, error) {
	return r.ApiService.BulkUsVerificationsExecute(r)
}

/*
BulkUsVerifications verifyBulk

Verify a list of US or US territory addresses with a live API key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkUsVerificationsRequest
*/
func (a *UsVerificationsApiService) BulkUsVerifications(ctx context.Context) ApiBulkUsVerificationsRequest {
	return ApiBulkUsVerificationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsVerifications
func (a *UsVerificationsApiService) BulkUsVerificationsExecute(r ApiBulkUsVerificationsRequest) (*UsVerifications, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsVerifications
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsVerificationsApiService.BulkUsVerifications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bulk/us_verifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.multipleComponentsList == nil {
		return localVarReturnValue, nil, reportError("multipleComponentsList is required and must be specified")
	}

	if r.case_ != nil {
		localVarQueryParams.Add("case", parameterToString(*r.case_, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.multipleComponentsList
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v LobError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsVerificationRequest struct {
	ctx context.Context
	ApiService *UsVerificationsApiService
	usVerificationsWritable *UsVerificationsWritable
	case_ *string
}

func (r ApiUsVerificationRequest) UsVerificationsWritable(usVerificationsWritable UsVerificationsWritable) ApiUsVerificationRequest {
	r.usVerificationsWritable = &usVerificationsWritable
	return r
}

// Casing of the verified address. Possible values are &#x60;upper&#x60; and &#x60;proper&#x60; for uppercased (e.g. \&quot;PO BOX\&quot;) and proper-cased (e.g. \&quot;PO Box\&quot;), respectively.
func (r ApiUsVerificationRequest) Case_(case_ string) ApiUsVerificationRequest {
	r.case_ = &case_
	return r
}

func (r ApiUsVerificationRequest) Execute() (*UsVerification, *http.Response, error) {
	return r.ApiService.UsVerificationExecute(r)
}

/*
UsVerification verifySingle

Verify a US or US territory address with a live API key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUsVerificationRequest
*/
func (a *UsVerificationsApiService) UsVerification(ctx context.Context) ApiUsVerificationRequest {
	return ApiUsVerificationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsVerification
func (a *UsVerificationsApiService) UsVerificationExecute(r ApiUsVerificationRequest) (*UsVerification, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsVerification
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsVerificationsApiService.UsVerification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/us_verifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.usVerificationsWritable == nil {
		return localVarReturnValue, nil, reportError("usVerificationsWritable is required and must be specified")
	}

	if r.case_ != nil {
		localVarQueryParams.Add("case", parameterToString(*r.case_, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.usVerificationsWritable
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v LobError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}