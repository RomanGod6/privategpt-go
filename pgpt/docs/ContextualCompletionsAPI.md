# \ContextualCompletionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatCompletionV1ChatCompletionsPost**](ContextualCompletionsAPI.md#ChatCompletionV1ChatCompletionsPost) | **Post** /v1/chat/completions | Chat Completion
[**PromptCompletionV1CompletionsPost**](ContextualCompletionsAPI.md#PromptCompletionV1CompletionsPost) | **Post** /v1/completions | Completion



## ChatCompletionV1ChatCompletionsPost

> OpenAICompletion ChatCompletionV1ChatCompletionsPost(ctx).ChatBody(chatBody).Execute()

Chat Completion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sdk "github.com/RomanGod6/privategpt-go/pgpt"
)

func main() {
	chatBody := *pgpt.NewChatBody([]pgpt.OpenAIMessage{*pgpt.NewOpenAIMessage("Content_example")}) // ChatBody | 

	configuration := pgpt.NewConfiguration()
	apiClient := pgpt.NewAPIClient(configuration)
	resp, r, err := apiClient.ContextualCompletionsAPI.ChatCompletionV1ChatCompletionsPost(context.Background()).ChatBody(chatBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContextualCompletionsAPI.ChatCompletionV1ChatCompletionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatCompletionV1ChatCompletionsPost`: OpenAICompletion
	fmt.Fprintf(os.Stdout, "Response from `ContextualCompletionsAPI.ChatCompletionV1ChatCompletionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatCompletionV1ChatCompletionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatBody** | [**ChatBody**](ChatBody.md) |  | 

### Return type

[**OpenAICompletion**](OpenAICompletion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptCompletionV1CompletionsPost

> OpenAICompletion PromptCompletionV1CompletionsPost(ctx).CompletionsBody(completionsBody).Execute()

Completion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sdk "github.com/RomanGod6/privategpt-go/pgpt"
)

func main() {
	completionsBody := *pgpt.NewCompletionsBody("Prompt_example") // CompletionsBody | 

	configuration := pgpt.NewConfiguration()
	apiClient := pgpt.NewAPIClient(configuration)
	resp, r, err := apiClient.ContextualCompletionsAPI.PromptCompletionV1CompletionsPost(context.Background()).CompletionsBody(completionsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContextualCompletionsAPI.PromptCompletionV1CompletionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptCompletionV1CompletionsPost`: OpenAICompletion
	fmt.Fprintf(os.Stdout, "Response from `ContextualCompletionsAPI.PromptCompletionV1CompletionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPromptCompletionV1CompletionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionsBody** | [**CompletionsBody**](CompletionsBody.md) |  | 

### Return type

[**OpenAICompletion**](OpenAICompletion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

