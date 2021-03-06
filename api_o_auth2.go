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
)

// Linger please
var (
	_ _context.Context
)

// OAuth2ApiService OAuth2Api service
type OAuth2ApiService service

type ApiGetOauth2AuthRequest struct {
	ctx _context.Context
	ApiService *OAuth2ApiService
	clientId *string
	responseType *string
	responseMode *string
	redirectUri *string
	state *string
	scope *string
	codeChallengeMethod *string
	codeChallenge *string
	nonce *string
}

func (r ApiGetOauth2AuthRequest) ClientId(clientId string) ApiGetOauth2AuthRequest {
	r.clientId = &clientId
	return r
}
func (r ApiGetOauth2AuthRequest) ResponseType(responseType string) ApiGetOauth2AuthRequest {
	r.responseType = &responseType
	return r
}
func (r ApiGetOauth2AuthRequest) ResponseMode(responseMode string) ApiGetOauth2AuthRequest {
	r.responseMode = &responseMode
	return r
}
func (r ApiGetOauth2AuthRequest) RedirectUri(redirectUri string) ApiGetOauth2AuthRequest {
	r.redirectUri = &redirectUri
	return r
}
func (r ApiGetOauth2AuthRequest) State(state string) ApiGetOauth2AuthRequest {
	r.state = &state
	return r
}
func (r ApiGetOauth2AuthRequest) Scope(scope string) ApiGetOauth2AuthRequest {
	r.scope = &scope
	return r
}
func (r ApiGetOauth2AuthRequest) CodeChallengeMethod(codeChallengeMethod string) ApiGetOauth2AuthRequest {
	r.codeChallengeMethod = &codeChallengeMethod
	return r
}
func (r ApiGetOauth2AuthRequest) CodeChallenge(codeChallenge string) ApiGetOauth2AuthRequest {
	r.codeChallenge = &codeChallenge
	return r
}
func (r ApiGetOauth2AuthRequest) Nonce(nonce string) ApiGetOauth2AuthRequest {
	r.nonce = &nonce
	return r
}

func (r ApiGetOauth2AuthRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetOauth2AuthExecute(r)
}

/*
 * GetOauth2Auth OAuth 2.0 Authorize Endpoint
 * To learn more about this flow please refer to the specification: [RFC-6749](https://tools.ietf.org/html/rfc6749)

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetOauth2AuthRequest
 */
func (a *OAuth2ApiService) GetOauth2Auth(ctx _context.Context) ApiGetOauth2AuthRequest {
	return ApiGetOauth2AuthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *OAuth2ApiService) GetOauth2AuthExecute(r ApiGetOauth2AuthRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth2ApiService.GetOauth2Auth")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/auth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.clientId == nil {
		return nil, reportError("clientId is required and must be specified")
	}
	if r.responseType == nil {
		return nil, reportError("responseType is required and must be specified")
	}
	if r.responseMode == nil {
		return nil, reportError("responseMode is required and must be specified")
	}
	if r.redirectUri == nil {
		return nil, reportError("redirectUri is required and must be specified")
	}
	if r.state == nil {
		return nil, reportError("state is required and must be specified")
	}
	if r.scope == nil {
		return nil, reportError("scope is required and must be specified")
	}
	if r.codeChallengeMethod == nil {
		return nil, reportError("codeChallengeMethod is required and must be specified")
	}
	if r.codeChallenge == nil {
		return nil, reportError("codeChallenge is required and must be specified")
	}

	localVarQueryParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarQueryParams.Add("response_type", parameterToString(*r.responseType, ""))
	localVarQueryParams.Add("response_mode", parameterToString(*r.responseMode, ""))
	localVarQueryParams.Add("redirect_uri", parameterToString(*r.redirectUri, ""))
	localVarQueryParams.Add("state", parameterToString(*r.state, ""))
	localVarQueryParams.Add("scope", parameterToString(*r.scope, ""))
	if r.nonce != nil {
		localVarQueryParams.Add("nonce", parameterToString(*r.nonce, ""))
	}
	localVarQueryParams.Add("code_challenge_method", parameterToString(*r.codeChallengeMethod, ""))
	localVarQueryParams.Add("code_challenge", parameterToString(*r.codeChallenge, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetOauth2TokenRequest struct {
	ctx _context.Context
	ApiService *OAuth2ApiService
	clientId *string
	clientSecret *string
	code *string
	codeVerifier *string
	grantType *string
	redirectUri *string
	refreshToken *string
	scope *string
}

func (r ApiGetOauth2TokenRequest) ClientId(clientId string) ApiGetOauth2TokenRequest {
	r.clientId = &clientId
	return r
}
func (r ApiGetOauth2TokenRequest) ClientSecret(clientSecret string) ApiGetOauth2TokenRequest {
	r.clientSecret = &clientSecret
	return r
}
func (r ApiGetOauth2TokenRequest) Code(code string) ApiGetOauth2TokenRequest {
	r.code = &code
	return r
}
func (r ApiGetOauth2TokenRequest) CodeVerifier(codeVerifier string) ApiGetOauth2TokenRequest {
	r.codeVerifier = &codeVerifier
	return r
}
func (r ApiGetOauth2TokenRequest) GrantType(grantType string) ApiGetOauth2TokenRequest {
	r.grantType = &grantType
	return r
}
func (r ApiGetOauth2TokenRequest) RedirectUri(redirectUri string) ApiGetOauth2TokenRequest {
	r.redirectUri = &redirectUri
	return r
}
func (r ApiGetOauth2TokenRequest) RefreshToken(refreshToken string) ApiGetOauth2TokenRequest {
	r.refreshToken = &refreshToken
	return r
}
func (r ApiGetOauth2TokenRequest) Scope(scope string) ApiGetOauth2TokenRequest {
	r.scope = &scope
	return r
}

func (r ApiGetOauth2TokenRequest) Execute() (TokenResponse, *_nethttp.Response, error) {
	return r.ApiService.GetOauth2TokenExecute(r)
}

/*
 * GetOauth2Token OAuth 2.0 Token Endpoint
 * The client makes a request to the token endpoint by sending the following parameters using the
`application/x-www-form-urlencoded` HTTP request entity-body.<br/>
To perform this operation, you must be authenticated by means of one of the following methods: `basic`, `oauth2`.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetOauth2TokenRequest
 */
func (a *OAuth2ApiService) GetOauth2Token(ctx _context.Context) ApiGetOauth2TokenRequest {
	return ApiGetOauth2TokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return TokenResponse
 */
func (a *OAuth2ApiService) GetOauth2TokenExecute(r ApiGetOauth2TokenRequest) (TokenResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth2ApiService.GetOauth2Token")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.clientSecret == nil {
		return localVarReturnValue, nil, reportError("clientSecret is required and must be specified")
	}
	if r.code == nil {
		return localVarReturnValue, nil, reportError("code is required and must be specified")
	}
	if r.codeVerifier == nil {
		return localVarReturnValue, nil, reportError("codeVerifier is required and must be specified")
	}
	if r.grantType == nil {
		return localVarReturnValue, nil, reportError("grantType is required and must be specified")
	}
	if r.redirectUri == nil {
		return localVarReturnValue, nil, reportError("redirectUri is required and must be specified")
	}
	if r.refreshToken == nil {
		return localVarReturnValue, nil, reportError("refreshToken is required and must be specified")
	}
	if r.scope == nil {
		return localVarReturnValue, nil, reportError("scope is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarFormParams.Add("client_secret", parameterToString(*r.clientSecret, ""))
	localVarFormParams.Add("code", parameterToString(*r.code, ""))
	localVarFormParams.Add("code_verifier", parameterToString(*r.codeVerifier, ""))
	localVarFormParams.Add("grant_type", parameterToString(*r.grantType, ""))
	localVarFormParams.Add("redirect_uri", parameterToString(*r.redirectUri, ""))
	localVarFormParams.Add("refresh_token", parameterToString(*r.refreshToken, ""))
	localVarFormParams.Add("scope", parameterToString(*r.scope, ""))
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

type ApiPostOauth2RevokeRequest struct {
	ctx _context.Context
	ApiService *OAuth2ApiService
	clientId *string
	clientSecret *string
	token *string
}

func (r ApiPostOauth2RevokeRequest) ClientId(clientId string) ApiPostOauth2RevokeRequest {
	r.clientId = &clientId
	return r
}
func (r ApiPostOauth2RevokeRequest) ClientSecret(clientSecret string) ApiPostOauth2RevokeRequest {
	r.clientSecret = &clientSecret
	return r
}
func (r ApiPostOauth2RevokeRequest) Token(token string) ApiPostOauth2RevokeRequest {
	r.token = &token
	return r
}

func (r ApiPostOauth2RevokeRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PostOauth2RevokeExecute(r)
}

/*
 * PostOauth2Revoke Revoke a token (Access or Refresh)
 * Revoking a token (both `access` and `refresh`) means that the tokens will be ***invalid***.<br/>
A revoked access token can no longer be used to make access requests, and a revoked refresh token can no longer
be used to refresh an access token.<br/>
Revoking a refresh token also invalidates the access token that was created with it.<br/>
**A token may only be revoked by the client the token was generated for!!**

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostOauth2RevokeRequest
 */
func (a *OAuth2ApiService) PostOauth2Revoke(ctx _context.Context) ApiPostOauth2RevokeRequest {
	return ApiPostOauth2RevokeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *OAuth2ApiService) PostOauth2RevokeExecute(r ApiPostOauth2RevokeRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth2ApiService.PostOauth2Revoke")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/revoke"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.clientId == nil {
		return nil, reportError("clientId is required and must be specified")
	}
	if r.clientSecret == nil {
		return nil, reportError("clientSecret is required and must be specified")
	}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarFormParams.Add("client_secret", parameterToString(*r.clientSecret, ""))
	localVarFormParams.Add("token", parameterToString(*r.token, ""))
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
