/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.39
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// GlobalCustomerApiService GlobalCustomerApi service
type GlobalCustomerApiService service

type ApiGlobalCustomerGetEndpointV1Request struct {
	ctx _context.Context
	ApiService *GlobalCustomerApiService
	pksCustomerCode string
	sInfrastructureproductCode *string
}

func (r ApiGlobalCustomerGetEndpointV1Request) SInfrastructureproductCode(sInfrastructureproductCode string) ApiGlobalCustomerGetEndpointV1Request {
	r.sInfrastructureproductCode = &sInfrastructureproductCode
	return r
}

func (r ApiGlobalCustomerGetEndpointV1Request) Execute() (GlobalCustomerGetEndpointV1Response, *_nethttp.Response, error) {
	return r.ApiService.GlobalCustomerGetEndpointV1Execute(r)
}

/*
 * GlobalCustomerGetEndpointV1 Get customer endpoint
 * Retrieve the customer's specific server endpoint where to send requests. This will help locate the proper region (ie: sInfrastructureregionCode) and the proper environment (ie: sInfrastructureenvironmenttypeDescription) where the customer's data is stored.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pksCustomerCode The customer code assigned to your account
 * @return ApiGlobalCustomerGetEndpointV1Request
 */
func (a *GlobalCustomerApiService) GlobalCustomerGetEndpointV1(ctx _context.Context, pksCustomerCode string) ApiGlobalCustomerGetEndpointV1Request {
	return ApiGlobalCustomerGetEndpointV1Request{
		ApiService: a,
		ctx: ctx,
		pksCustomerCode: pksCustomerCode,
	}
}

/*
 * Execute executes the request
 * @return GlobalCustomerGetEndpointV1Response
 */
func (a *GlobalCustomerApiService) GlobalCustomerGetEndpointV1Execute(r ApiGlobalCustomerGetEndpointV1Request) (GlobalCustomerGetEndpointV1Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GlobalCustomerGetEndpointV1Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalCustomerApiService.GlobalCustomerGetEndpointV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/customer/{pksCustomerCode}/endpoint"
	localVarPath = strings.Replace(localVarPath, "{"+"pksCustomerCode"+"}", _neturl.PathEscape(parameterToString(r.pksCustomerCode, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.sInfrastructureproductCode != nil {
		localVarQueryParams.Add("sInfrastructureproductCode", parameterToString(*r.sInfrastructureproductCode, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v CommonResponseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
