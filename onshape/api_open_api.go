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
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// OpenAPIApiService OpenAPIApi service
type OpenAPIApiService service

type apiGetOpenApiRequest struct {
	ctx _context.Context
	apiService *OpenAPIApiService
	excludedTags *string
	includedTags *string
	includeDeprecated *bool
	documentationStatus *[]string
}


func (r apiGetOpenApiRequest) ExcludedTags(excludedTags string) apiGetOpenApiRequest {
	r.excludedTags = &excludedTags
	return r
}

func (r apiGetOpenApiRequest) IncludedTags(includedTags string) apiGetOpenApiRequest {
	r.includedTags = &includedTags
	return r
}

func (r apiGetOpenApiRequest) IncludeDeprecated(includeDeprecated bool) apiGetOpenApiRequest {
	r.includeDeprecated = &includeDeprecated
	return r
}

func (r apiGetOpenApiRequest) DocumentationStatus(documentationStatus []string) apiGetOpenApiRequest {
	r.documentationStatus = &documentationStatus
	return r
}

/*
GetOpenApi OpenAPI spec documentation for the Onshape REST API.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiGetOpenApiRequest
*/
func (a *OpenAPIApiService) GetOpenApi(ctx _context.Context) apiGetOpenApiRequest {
	return apiGetOpenApiRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request

*/
func (r apiGetOpenApiRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "OpenAPIApiService.GetOpenApi")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/openapi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
				
	if r.excludedTags != nil {
		localVarQueryParams.Add("excludedTags", parameterToString(*r.excludedTags, ""))
	}
	if r.includedTags != nil {
		localVarQueryParams.Add("includedTags", parameterToString(*r.includedTags, ""))
	}
	if r.includeDeprecated != nil {
		localVarQueryParams.Add("includeDeprecated", parameterToString(*r.includeDeprecated, ""))
	}
	if r.documentationStatus != nil {
		t := *r.documentationStatus
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("documentationStatus", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("documentationStatus", parameterToString(t, "multi"))
		}
	}
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
