/*
FastAPI- Private GPT

SDK for GO LANG for Private GPT

API version: 0.1.0
*/

package pgpt

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ContextualCompletionsAPIService ContextualCompletionsAPI service
type ContextualCompletionsAPIService service

type ApiChatCompletionV1ChatCompletionsPostRequest struct {
	ctx        context.Context
	ApiService *ContextualCompletionsAPIService
	chatBody   *ChatBody
}

func (r ApiChatCompletionV1ChatCompletionsPostRequest) ChatBody(chatBody ChatBody) ApiChatCompletionV1ChatCompletionsPostRequest {
	r.chatBody = &chatBody
	return r
}

func (r ApiChatCompletionV1ChatCompletionsPostRequest) Execute() (*OpenAICompletion, *http.Response, error) {
	return r.ApiService.ChatCompletionV1ChatCompletionsPostExecute(r)
}

/*
ChatCompletionV1ChatCompletionsPost Chat Completion

Given a list of messages comprising a conversation, return a response.

Optionally include an initial `role: system` message to influence the way
the LLM answers.

If `use_context` is set to `true`, the model will use context coming
from the ingested documents to create the response. The documents being used can
be filtered using the `context_filter` and passing the document IDs to be used.
Ingested documents IDs can be found using `/ingest/list` endpoint. If you want
all ingested documents to be used, remove `context_filter` altogether.

When using `'include_sources': true`, the API will return the source Chunks used
to create the response, which come from the context provided.

When using `'stream': true`, the API will return data chunks following [OpenAI's
streaming model](https://platform.openai.com/docs/api-reference/chat/streaming):
```
{"id":"12345","object":"completion.chunk","created":1694268190,
"model":"private-gpt","choices":[{"index":0,"delta":{"content":"Hello"},
"finish_reason":null}]}
```

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiChatCompletionV1ChatCompletionsPostRequest
*/
func (a *ContextualCompletionsAPIService) ChatCompletionV1ChatCompletionsPost(ctx context.Context) ApiChatCompletionV1ChatCompletionsPostRequest {
	return ApiChatCompletionV1ChatCompletionsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OpenAICompletion
func (a *ContextualCompletionsAPIService) ChatCompletionV1ChatCompletionsPostExecute(r ApiChatCompletionV1ChatCompletionsPostRequest) (*OpenAICompletion, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OpenAICompletion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextualCompletionsAPIService.ChatCompletionV1ChatCompletionsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat/completions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.chatBody == nil {
		return localVarReturnValue, nil, reportError("chatBody is required and must be specified")
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
	localVarPostBody = r.chatBody
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

type ApiPromptCompletionV1CompletionsPostRequest struct {
	ctx             context.Context
	ApiService      *ContextualCompletionsAPIService
	completionsBody *CompletionsBody
}

func (r ApiPromptCompletionV1CompletionsPostRequest) CompletionsBody(completionsBody CompletionsBody) ApiPromptCompletionV1CompletionsPostRequest {
	r.completionsBody = &completionsBody
	return r
}

func (r ApiPromptCompletionV1CompletionsPostRequest) Execute() (*OpenAICompletion, *http.Response, error) {
	return r.ApiService.PromptCompletionV1CompletionsPostExecute(r)
}

/*
PromptCompletionV1CompletionsPost Completion

We recommend most users use our Chat completions API.

Given a prompt, the model will return one predicted completion.

Optionally include a `system_prompt` to influence the way the LLM answers.

If `use_context`
is set to `true`, the model will use context coming from the ingested documents
to create the response. The documents being used can be filtered using the
`context_filter` and passing the document IDs to be used. Ingested documents IDs
can be found using `/ingest/list` endpoint. If you want all ingested documents to
be used, remove `context_filter` altogether.

When using `'include_sources': true`, the API will return the source Chunks used
to create the response, which come from the context provided.

When using `'stream': true`, the API will return data chunks following [OpenAI's
streaming model](https://platform.openai.com/docs/api-reference/chat/streaming):
```
{"id":"12345","object":"completion.chunk","created":1694268190,
"model":"private-gpt","choices":[{"index":0,"delta":{"content":"Hello"},
"finish_reason":null}]}
```

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPromptCompletionV1CompletionsPostRequest
*/
func (a *ContextualCompletionsAPIService) PromptCompletionV1CompletionsPost(ctx context.Context) ApiPromptCompletionV1CompletionsPostRequest {
	return ApiPromptCompletionV1CompletionsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OpenAICompletion
func (a *ContextualCompletionsAPIService) PromptCompletionV1CompletionsPostExecute(r ApiPromptCompletionV1CompletionsPostRequest) (*OpenAICompletion, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OpenAICompletion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextualCompletionsAPIService.PromptCompletionV1CompletionsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/completions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.completionsBody == nil {
		return localVarReturnValue, nil, reportError("completionsBody is required and must be specified")
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
	localVarPostBody = r.completionsBody
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
