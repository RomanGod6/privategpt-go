/*
FastAPI- Private GPT

SDK for GO LANG for Private GPT

API version: 0.1.0
*/

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ContextChunksAPIService ContextChunksAPI service
type ContextChunksAPIService service

type ApiChunksRetrievalV1ChunksPostRequest struct {
	ctx        context.Context
	ApiService *ContextChunksAPIService
	chunksBody *ChunksBody
}

func (r ApiChunksRetrievalV1ChunksPostRequest) ChunksBody(chunksBody ChunksBody) ApiChunksRetrievalV1ChunksPostRequest {
	r.chunksBody = &chunksBody
	return r
}

func (r ApiChunksRetrievalV1ChunksPostRequest) Execute() (*ChunksResponse, *http.Response, error) {
	return r.ApiService.ChunksRetrievalV1ChunksPostExecute(r)
}

/*
ChunksRetrievalV1ChunksPost Chunks Retrieval

Given a `text`, returns the most relevant chunks from the ingested documents.

The returned information can be used to generate prompts that can be
passed to `/completions` or `/chat/completions` APIs. Note: it is usually a very
fast API, because only the Embeddings model is involved, not the LLM. The
returned information contains the relevant chunk `text` together with the source
`document` it is coming from. It also contains a score that can be used to
compare different results.

The max number of chunks to be returned is set using the `limit` param.

Previous and next chunks (pieces of text that appear right before or after in the
document) can be fetched by using the `prev_next_chunks` field.

The documents being used can be filtered using the `context_filter` and passing
the document IDs to be used. Ingested documents IDs can be found using
`/ingest/list` endpoint. If you want all ingested documents to be used,
remove `context_filter` altogether.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiChunksRetrievalV1ChunksPostRequest
*/
func (a *ContextChunksAPIService) ChunksRetrievalV1ChunksPost(ctx context.Context) ApiChunksRetrievalV1ChunksPostRequest {
	return ApiChunksRetrievalV1ChunksPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChunksResponse
func (a *ContextChunksAPIService) ChunksRetrievalV1ChunksPostExecute(r ApiChunksRetrievalV1ChunksPostRequest) (*ChunksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ChunksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextChunksAPIService.ChunksRetrievalV1ChunksPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chunks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.chunksBody == nil {
		return localVarReturnValue, nil, reportError("chunksBody is required and must be specified")
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
	localVarPostBody = r.chunksBody
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
