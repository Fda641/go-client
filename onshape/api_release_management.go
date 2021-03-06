/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
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

// ReleaseManagementApiService ReleaseManagementApi service
type ReleaseManagementApiService service

type apiCreateObsoletionPackageRequest struct {
	ctx _context.Context
	apiService *ReleaseManagementApiService
	wfid string
	revisionId *string
}


func (r apiCreateObsoletionPackageRequest) RevisionId(revisionId string) apiCreateObsoletionPackageRequest {
	r.revisionId = &revisionId
	return r
}

/*
CreateObsoletionPackage Method for CreateObsoletionPackage
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param wfid
@return apiCreateObsoletionPackageRequest
*/
func (a *ReleaseManagementApiService) CreateObsoletionPackage(ctx _context.Context, wfid string) apiCreateObsoletionPackageRequest {
	return apiCreateObsoletionPackageRequest{
		apiService: a,
		ctx: ctx,
		wfid: wfid,
	}
}

/*
Execute executes the request

*/
func (r apiCreateObsoletionPackageRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.CreateObsoletionPackage")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/releasepackages/obsoletion/{wfid}"
	localVarPath = strings.Replace(localVarPath, "{"+"wfid"+"}", _neturl.QueryEscape(parameterToString(r.wfid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	if r.revisionId != nil {
		localVarQueryParams.Add("revisionId", parameterToString(*r.revisionId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
type apiCreateReleasePackageRequest struct {
	ctx _context.Context
	apiService *ReleaseManagementApiService
	wfid string
	bTReleasePackageParams *BTReleasePackageParams
}


func (r apiCreateReleasePackageRequest) BTReleasePackageParams(bTReleasePackageParams BTReleasePackageParams) apiCreateReleasePackageRequest {
	r.bTReleasePackageParams = &bTReleasePackageParams
	return r
}

/*
CreateReleasePackage Method for CreateReleasePackage
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param wfid
@return apiCreateReleasePackageRequest
*/
func (a *ReleaseManagementApiService) CreateReleasePackage(ctx _context.Context, wfid string) apiCreateReleasePackageRequest {
	return apiCreateReleasePackageRequest{
		apiService: a,
		ctx: ctx,
		wfid: wfid,
	}
}

/*
Execute executes the request

*/
func (r apiCreateReleasePackageRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.CreateReleasePackage")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/releasepackages/release/{wfid}"
	localVarPath = strings.Replace(localVarPath, "{"+"wfid"+"}", _neturl.QueryEscape(parameterToString(r.wfid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	if r.bTReleasePackageParams == nil {
		return nil, reportError("bTReleasePackageParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTReleasePackageParams
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
type apiGetCompanyReleaseWorkflowRequest struct {
	ctx _context.Context
	apiService *ReleaseManagementApiService
	documentId *string
}


func (r apiGetCompanyReleaseWorkflowRequest) DocumentId(documentId string) apiGetCompanyReleaseWorkflowRequest {
	r.documentId = &documentId
	return r
}

/*
GetCompanyReleaseWorkflow Method for GetCompanyReleaseWorkflow
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiGetCompanyReleaseWorkflowRequest
*/
func (a *ReleaseManagementApiService) GetCompanyReleaseWorkflow(ctx _context.Context) apiGetCompanyReleaseWorkflowRequest {
	return apiGetCompanyReleaseWorkflowRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return BTActiveWorkflowInfo
*/
func (r apiGetCompanyReleaseWorkflowRequest) Execute() (BTActiveWorkflowInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTActiveWorkflowInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.GetCompanyReleaseWorkflow")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/releasepackages/companyreleaseworkflow"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.documentId != nil {
		localVarQueryParams.Add("documentId", parameterToString(*r.documentId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTActiveWorkflowInfo
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiGetReleasePackageRequest struct {
	ctx _context.Context
	apiService *ReleaseManagementApiService
	rpid string
	detailed *bool
}


func (r apiGetReleasePackageRequest) Detailed(detailed bool) apiGetReleasePackageRequest {
	r.detailed = &detailed
	return r
}

/*
GetReleasePackage Method for GetReleasePackage
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param rpid
@return apiGetReleasePackageRequest
*/
func (a *ReleaseManagementApiService) GetReleasePackage(ctx _context.Context, rpid string) apiGetReleasePackageRequest {
	return apiGetReleasePackageRequest{
		apiService: a,
		ctx: ctx,
		rpid: rpid,
	}
}

/*
Execute executes the request
 @return BTReleasePackageInfo
*/
func (r apiGetReleasePackageRequest) Execute() (BTReleasePackageInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTReleasePackageInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.GetReleasePackage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/releasepackages/{rpid}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpid"+"}", _neturl.QueryEscape(parameterToString(r.rpid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	if r.detailed != nil {
		localVarQueryParams.Add("detailed", parameterToString(*r.detailed, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTReleasePackageInfo
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiUpdateReleasePackageRequest struct {
	ctx _context.Context
	apiService *ReleaseManagementApiService
	rpid string
	bTUpdateReleasePackageParams *BTUpdateReleasePackageParams
	action *string
	wfaction *string
}


func (r apiUpdateReleasePackageRequest) BTUpdateReleasePackageParams(bTUpdateReleasePackageParams BTUpdateReleasePackageParams) apiUpdateReleasePackageRequest {
	r.bTUpdateReleasePackageParams = &bTUpdateReleasePackageParams
	return r
}

func (r apiUpdateReleasePackageRequest) Action(action string) apiUpdateReleasePackageRequest {
	r.action = &action
	return r
}

func (r apiUpdateReleasePackageRequest) Wfaction(wfaction string) apiUpdateReleasePackageRequest {
	r.wfaction = &wfaction
	return r
}

/*
UpdateReleasePackage Method for UpdateReleasePackage
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param rpid
@return apiUpdateReleasePackageRequest
*/
func (a *ReleaseManagementApiService) UpdateReleasePackage(ctx _context.Context, rpid string) apiUpdateReleasePackageRequest {
	return apiUpdateReleasePackageRequest{
		apiService: a,
		ctx: ctx,
		rpid: rpid,
	}
}

/*
Execute executes the request
 @return BTReleasePackageInfo
*/
func (r apiUpdateReleasePackageRequest) Execute() (BTReleasePackageInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTReleasePackageInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.UpdateReleasePackage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/releasepackages/{rpid}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpid"+"}", _neturl.QueryEscape(parameterToString(r.rpid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	if r.bTUpdateReleasePackageParams == nil {
		return localVarReturnValue, nil, reportError("bTUpdateReleasePackageParams is required and must be specified")
	}
		
	if r.action != nil {
		localVarQueryParams.Add("action", parameterToString(*r.action, ""))
	}
	if r.wfaction != nil {
		localVarQueryParams.Add("wfaction", parameterToString(*r.wfaction, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTUpdateReleasePackageParams
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTReleasePackageInfo
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
