# \EmbeddingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmbeddingsGenerationV1EmbeddingsPost**](EmbeddingsAPI.md#EmbeddingsGenerationV1EmbeddingsPost) | **Post** /v1/embeddings | Embeddings Generation



## EmbeddingsGenerationV1EmbeddingsPost

> EmbeddingsResponse EmbeddingsGenerationV1EmbeddingsPost(ctx).EmbeddingsBody(embeddingsBody).Execute()

Embeddings Generation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RomanGod6/privategpt-go/sdk"
)

func main() {
	embeddingsBody := *openapiclient.NewEmbeddingsBody(*openapiclient.NewInput()) // EmbeddingsBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.EmbeddingsGenerationV1EmbeddingsPost(context.Background()).EmbeddingsBody(embeddingsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.EmbeddingsGenerationV1EmbeddingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmbeddingsGenerationV1EmbeddingsPost`: EmbeddingsResponse
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.EmbeddingsGenerationV1EmbeddingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmbeddingsGenerationV1EmbeddingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **embeddingsBody** | [**EmbeddingsBody**](EmbeddingsBody.md) |  | 

### Return type

[**EmbeddingsResponse**](EmbeddingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

