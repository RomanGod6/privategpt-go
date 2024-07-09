# \ContextChunksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChunksRetrievalV1ChunksPost**](ContextChunksAPI.md#ChunksRetrievalV1ChunksPost) | **Post** /v1/chunks | Chunks Retrieval



## ChunksRetrievalV1ChunksPost

> ChunksResponse ChunksRetrievalV1ChunksPost(ctx).ChunksBody(chunksBody).Execute()

Chunks Retrieval



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
	chunksBody := *openapiclient.NewChunksBody("Text_example") // ChunksBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContextChunksAPI.ChunksRetrievalV1ChunksPost(context.Background()).ChunksBody(chunksBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContextChunksAPI.ChunksRetrievalV1ChunksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChunksRetrievalV1ChunksPost`: ChunksResponse
	fmt.Fprintf(os.Stdout, "Response from `ContextChunksAPI.ChunksRetrievalV1ChunksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChunksRetrievalV1ChunksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chunksBody** | [**ChunksBody**](ChunksBody.md) |  | 

### Return type

[**ChunksResponse**](ChunksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

