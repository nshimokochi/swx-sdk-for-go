/*
 * Accounts & Users Service - Public API
 *
 * IN PROGRESS->This is the guide to use the different endpoints to manage the clusters.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

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

// MQTTApiService MQTTApi service
type MQTTApiService service

type ApiCreateMQTTLabelCredentialsRequest struct {
	ctx _context.Context
	ApiService *MQTTApiService
	space string
	mQTTPOSTBody *MQTTPOSTBody
}

func (r ApiCreateMQTTLabelCredentialsRequest) MQTTPOSTBody(mQTTPOSTBody MQTTPOSTBody) ApiCreateMQTTLabelCredentialsRequest {
	r.mQTTPOSTBody = &mQTTPOSTBody
	return r
}

func (r ApiCreateMQTTLabelCredentialsRequest) Execute() (MQTTPOSTResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateMQTTLabelCredentialsExecute(r)
}

/*
 * CreateMQTTLabelCredentials Create MQTT credentials for a label
 * Create MQTT credentials for a specific label
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @return ApiCreateMQTTLabelCredentialsRequest
 */
func (a *MQTTApiService) CreateMQTTLabelCredentials(ctx _context.Context, space string) ApiCreateMQTTLabelCredentialsRequest {
	return ApiCreateMQTTLabelCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
	}
}

/*
 * Execute executes the request
 * @return MQTTPOSTResponse
 */
func (a *MQTTApiService) CreateMQTTLabelCredentialsExecute(r ApiCreateMQTTLabelCredentialsRequest) (MQTTPOSTResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MQTTPOSTResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MQTTApiService.CreateMQTTLabelCredentials")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/mqtt/labels"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.mQTTPOSTBody
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

type ApiDeleteMQTTLabelRequest struct {
	ctx _context.Context
	ApiService *MQTTApiService
	space string
	labelId string
}


func (r ApiDeleteMQTTLabelRequest) Execute() (MQTTLabelDeleteResponse, *_nethttp.Response, error) {
	return r.ApiService.DeleteMQTTLabelExecute(r)
}

/*
 * DeleteMQTTLabel Delete MQTT label
 * Delete MQTT capabilities for a label
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param labelId
 * @return ApiDeleteMQTTLabelRequest
 */
func (a *MQTTApiService) DeleteMQTTLabel(ctx _context.Context, space string, labelId string) ApiDeleteMQTTLabelRequest {
	return ApiDeleteMQTTLabelRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		labelId: labelId,
	}
}

/*
 * Execute executes the request
 * @return MQTTLabelDeleteResponse
 */
func (a *MQTTApiService) DeleteMQTTLabelExecute(r ApiDeleteMQTTLabelRequest) (MQTTLabelDeleteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MQTTLabelDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MQTTApiService.DeleteMQTTLabel")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/mqtt/labels/{label-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"label-id"+"}", _neturl.PathEscape(parameterToString(r.labelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiListMQTTcredentialsRequest struct {
	ctx _context.Context
	ApiService *MQTTApiService
	space string
}


func (r ApiListMQTTcredentialsRequest) Execute() (CredentialsResponseList, *_nethttp.Response, error) {
	return r.ApiService.ListMQTTcredentialsExecute(r)
}

/*
 * ListMQTTcredentials List space MQTT credentials
 * List all MQTT credentials for a specific Account
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @return ApiListMQTTcredentialsRequest
 */
func (a *MQTTApiService) ListMQTTcredentials(ctx _context.Context, space string) ApiListMQTTcredentialsRequest {
	return ApiListMQTTcredentialsRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
	}
}

/*
 * Execute executes the request
 * @return CredentialsResponseList
 */
func (a *MQTTApiService) ListMQTTcredentialsExecute(r ApiListMQTTcredentialsRequest) (CredentialsResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CredentialsResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MQTTApiService.ListMQTTcredentials")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/mqtt/credentials"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiListThingMQTTcredentialsRequest struct {
	ctx _context.Context
	ApiService *MQTTApiService
	space string
	thingId string
}


func (r ApiListThingMQTTcredentialsRequest) Execute() (CredentialsResponseThing, *_nethttp.Response, error) {
	return r.ApiService.ListThingMQTTcredentialsExecute(r)
}

/*
 * ListThingMQTTcredentials List Thing MQTT credentials
 * List all MQTT credentials for a specific Thing
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param thingId
 * @return ApiListThingMQTTcredentialsRequest
 */
func (a *MQTTApiService) ListThingMQTTcredentials(ctx _context.Context, space string, thingId string) ApiListThingMQTTcredentialsRequest {
	return ApiListThingMQTTcredentialsRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		thingId: thingId,
	}
}

/*
 * Execute executes the request
 * @return CredentialsResponseThing
 */
func (a *MQTTApiService) ListThingMQTTcredentialsExecute(r ApiListThingMQTTcredentialsRequest) (CredentialsResponseThing, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CredentialsResponseThing
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MQTTApiService.ListThingMQTTcredentials")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/mqtt/things/{thing-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"thing-id"+"}", _neturl.PathEscape(parameterToString(r.thingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiShowMQTTLabelInfoRequest struct {
	ctx _context.Context
	ApiService *MQTTApiService
	space string
	labelId string
}


func (r ApiShowMQTTLabelInfoRequest) Execute() (MQTTLabelShowResponse, *_nethttp.Response, error) {
	return r.ApiService.ShowMQTTLabelInfoExecute(r)
}

/*
 * ShowMQTTLabelInfo Show MQTT label details
 * Show MQTT label details for a specific label
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param labelId
 * @return ApiShowMQTTLabelInfoRequest
 */
func (a *MQTTApiService) ShowMQTTLabelInfo(ctx _context.Context, space string, labelId string) ApiShowMQTTLabelInfoRequest {
	return ApiShowMQTTLabelInfoRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		labelId: labelId,
	}
}

/*
 * Execute executes the request
 * @return MQTTLabelShowResponse
 */
func (a *MQTTApiService) ShowMQTTLabelInfoExecute(r ApiShowMQTTLabelInfoRequest) (MQTTLabelShowResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MQTTLabelShowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MQTTApiService.ShowMQTTLabelInfo")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/mqtt/labels/{label-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"label-id"+"}", _neturl.PathEscape(parameterToString(r.labelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiShowMQTTcredentialsRequest struct {
	ctx _context.Context
	ApiService *MQTTApiService
	space string
	credentialsId string
}


func (r ApiShowMQTTcredentialsRequest) Execute() (CredentialsResponseThing, *_nethttp.Response, error) {
	return r.ApiService.ShowMQTTcredentialsExecute(r)
}

/*
 * ShowMQTTcredentials Show MQTT credentials
 * Show a specific MQTT credentials
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param credentialsId
 * @return ApiShowMQTTcredentialsRequest
 */
func (a *MQTTApiService) ShowMQTTcredentials(ctx _context.Context, space string, credentialsId string) ApiShowMQTTcredentialsRequest {
	return ApiShowMQTTcredentialsRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		credentialsId: credentialsId,
	}
}

/*
 * Execute executes the request
 * @return CredentialsResponseThing
 */
func (a *MQTTApiService) ShowMQTTcredentialsExecute(r ApiShowMQTTcredentialsRequest) (CredentialsResponseThing, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CredentialsResponseThing
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MQTTApiService.ShowMQTTcredentials")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/mqtt/credentials/{credentials-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"credentials-id"+"}", _neturl.PathEscape(parameterToString(r.credentialsId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiUpdateMQTTLabelRequest struct {
	ctx _context.Context
	ApiService *MQTTApiService
	space string
	labelId string
	mQTTLabelCredentials *MQTTLabelCredentials
}

func (r ApiUpdateMQTTLabelRequest) MQTTLabelCredentials(mQTTLabelCredentials MQTTLabelCredentials) ApiUpdateMQTTLabelRequest {
	r.mQTTLabelCredentials = &mQTTLabelCredentials
	return r
}

func (r ApiUpdateMQTTLabelRequest) Execute() (MQTTPOSTResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateMQTTLabelExecute(r)
}

/*
 * UpdateMQTTLabel Update MQTT Label
 * Update the MQTT Label credentials and / or state
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param labelId
 * @return ApiUpdateMQTTLabelRequest
 */
func (a *MQTTApiService) UpdateMQTTLabel(ctx _context.Context, space string, labelId string) ApiUpdateMQTTLabelRequest {
	return ApiUpdateMQTTLabelRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		labelId: labelId,
	}
}

/*
 * Execute executes the request
 * @return MQTTPOSTResponse
 */
func (a *MQTTApiService) UpdateMQTTLabelExecute(r ApiUpdateMQTTLabelRequest) (MQTTPOSTResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MQTTPOSTResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MQTTApiService.UpdateMQTTLabel")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/mqtt/labels/{label-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"label-id"+"}", _neturl.PathEscape(parameterToString(r.labelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.mQTTLabelCredentials == nil {
		return localVarReturnValue, nil, reportError("mQTTLabelCredentials is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.mQTTLabelCredentials
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

type ApiUpdateMQTTcredentialsRequest struct {
	ctx _context.Context
	ApiService *MQTTApiService
	space string
	credentialsId string
	mQTTPUTBody *MQTTPUTBody
}

func (r ApiUpdateMQTTcredentialsRequest) MQTTPUTBody(mQTTPUTBody MQTTPUTBody) ApiUpdateMQTTcredentialsRequest {
	r.mQTTPUTBody = &mQTTPUTBody
	return r
}

func (r ApiUpdateMQTTcredentialsRequest) Execute() (CredentialsResponsePUT, *_nethttp.Response, error) {
	return r.ApiService.UpdateMQTTcredentialsExecute(r)
}

/*
 * UpdateMQTTcredentials Update MQTT credentials
 * Update a specific MQTT credentials. If you send an empty body, a random password will be generated
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param credentialsId
 * @return ApiUpdateMQTTcredentialsRequest
 */
func (a *MQTTApiService) UpdateMQTTcredentials(ctx _context.Context, space string, credentialsId string) ApiUpdateMQTTcredentialsRequest {
	return ApiUpdateMQTTcredentialsRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		credentialsId: credentialsId,
	}
}

/*
 * Execute executes the request
 * @return CredentialsResponsePUT
 */
func (a *MQTTApiService) UpdateMQTTcredentialsExecute(r ApiUpdateMQTTcredentialsRequest) (CredentialsResponsePUT, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CredentialsResponsePUT
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MQTTApiService.UpdateMQTTcredentials")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/mqtt/credentials/{credentials-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"credentials-id"+"}", _neturl.PathEscape(parameterToString(r.credentialsId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.mQTTPUTBody
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

type ApiUpdateMQTTpasswordRequest struct {
	ctx _context.Context
	ApiService *MQTTApiService
	space string
	mqttUsername string
	mQTTPassword *MQTTPassword
}

func (r ApiUpdateMQTTpasswordRequest) MQTTPassword(mQTTPassword MQTTPassword) ApiUpdateMQTTpasswordRequest {
	r.mQTTPassword = &mQTTPassword
	return r
}

func (r ApiUpdateMQTTpasswordRequest) Execute() (CredentialsResponsePUT, *_nethttp.Response, error) {
	return r.ApiService.UpdateMQTTpasswordExecute(r)
}

/*
 * UpdateMQTTpassword Update MQTT password
 * Update the MQTT password for a specific mqtt-username.  If you send an empty body, a random password will be generated
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param mqttUsername
 * @return ApiUpdateMQTTpasswordRequest
 */
func (a *MQTTApiService) UpdateMQTTpassword(ctx _context.Context, space string, mqttUsername string) ApiUpdateMQTTpasswordRequest {
	return ApiUpdateMQTTpasswordRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		mqttUsername: mqttUsername,
	}
}

/*
 * Execute executes the request
 * @return CredentialsResponsePUT
 */
func (a *MQTTApiService) UpdateMQTTpasswordExecute(r ApiUpdateMQTTpasswordRequest) (CredentialsResponsePUT, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CredentialsResponsePUT
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MQTTApiService.UpdateMQTTpassword")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/mqtt/credentials-username/{mqtt-username}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqtt-username"+"}", _neturl.PathEscape(parameterToString(r.mqttUsername, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.mQTTPassword
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