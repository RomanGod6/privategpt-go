/*
FastAPI- Private GPT

SDK for GO LANG for Private GPT

API version: 0.1.0
*/

package sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// EmbeddingsAPIService EmbeddingsAPI service
type EmbeddingsAPIService service

type ApiEmbeddingsGenerationV1EmbeddingsPostRequest struct {
	ctx            context.Context
	ApiService     *EmbeddingsAPIService
	embeddingsBody *EmbeddingsBody
}

func (r ApiEmbeddingsGenerationV1EmbeddingsPostRequest) EmbeddingsBody(embeddingsBody EmbeddingsBody) ApiEmbeddingsGenerationV1EmbeddingsPostRequest {
	r.embeddingsBody = &embeddingsBody
	return r
}

func (r ApiEmbeddingsGenerationV1EmbeddingsPostRequest) Execute() (*EmbeddingsResponse, *http.Response, error) {
	return r.ApiService.EmbeddingsGenerationV1EmbeddingsPostExecute(r)
}

/*
EmbeddingsGenerationV1EmbeddingsPost Embeddings Generation

Get a vector representation of a given input.

That vector representation can be easily consumed
by machine learning models and algorithms.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEmbeddingsGenerationV1EmbeddingsPostRequest
*/
func (a *EmbeddingsAPIService) EmbeddingsGenerationV1EmbeddingsPost(ctx context.Context) ApiEmbeddingsGenerationV1EmbeddingsPostRequest {
	return ApiEmbeddingsGenerationV1EmbeddingsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EmbeddingsResponse
func (a *EmbeddingsAPIService) EmbeddingsGenerationV1EmbeddingsPostExecute(r ApiEmbeddingsGenerationV1EmbeddingsPostRequest) (*EmbeddingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EmbeddingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbeddingsAPIService.EmbeddingsGenerationV1EmbeddingsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/embeddings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.embeddingsBody == nil {
		return localVarReturnValue, nil, reportError("embeddingsBody is required and must be specified")
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
	localVarPostBody = r.embeddingsBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
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
