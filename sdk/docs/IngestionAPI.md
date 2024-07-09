# \IngestionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIngestedV1IngestDocIdDelete**](IngestionAPI.md#DeleteIngestedV1IngestDocIdDelete) | **Delete** /v1/ingest/{doc_id} | Delete Ingested
[**IngestFileV1IngestFilePost**](IngestionAPI.md#IngestFileV1IngestFilePost) | **Post** /v1/ingest/file | Ingest File
[**IngestTextV1IngestTextPost**](IngestionAPI.md#IngestTextV1IngestTextPost) | **Post** /v1/ingest/text | Ingest Text
[**IngestV1IngestPost**](IngestionAPI.md#IngestV1IngestPost) | **Post** /v1/ingest | Ingest
[**ListIngestedV1IngestListGet**](IngestionAPI.md#ListIngestedV1IngestListGet) | **Get** /v1/ingest/list | List Ingested



## DeleteIngestedV1IngestDocIdDelete

> interface{} DeleteIngestedV1IngestDocIdDelete(ctx, docId).Execute()

Delete Ingested



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
	docId := "docId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionAPI.DeleteIngestedV1IngestDocIdDelete(context.Background(), docId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionAPI.DeleteIngestedV1IngestDocIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIngestedV1IngestDocIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IngestionAPI.DeleteIngestedV1IngestDocIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**docId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestedV1IngestDocIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestFileV1IngestFilePost

> IngestResponse IngestFileV1IngestFilePost(ctx).File(file).Execute()

Ingest File



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
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionAPI.IngestFileV1IngestFilePost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionAPI.IngestFileV1IngestFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IngestFileV1IngestFilePost`: IngestResponse
	fmt.Fprintf(os.Stdout, "Response from `IngestionAPI.IngestFileV1IngestFilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestFileV1IngestFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**IngestResponse**](IngestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestTextV1IngestTextPost

> IngestResponse IngestTextV1IngestTextPost(ctx).IngestTextBody(ingestTextBody).Execute()

Ingest Text



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
	ingestTextBody := *openapiclient.NewIngestTextBody("FileName_example", "Text_example") // IngestTextBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionAPI.IngestTextV1IngestTextPost(context.Background()).IngestTextBody(ingestTextBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionAPI.IngestTextV1IngestTextPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IngestTextV1IngestTextPost`: IngestResponse
	fmt.Fprintf(os.Stdout, "Response from `IngestionAPI.IngestTextV1IngestTextPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestTextV1IngestTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestTextBody** | [**IngestTextBody**](IngestTextBody.md) |  | 

### Return type

[**IngestResponse**](IngestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestV1IngestPost

> IngestResponse IngestV1IngestPost(ctx).File(file).Execute()

Ingest



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
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionAPI.IngestV1IngestPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionAPI.IngestV1IngestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IngestV1IngestPost`: IngestResponse
	fmt.Fprintf(os.Stdout, "Response from `IngestionAPI.IngestV1IngestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestV1IngestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**IngestResponse**](IngestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIngestedV1IngestListGet

> IngestResponse ListIngestedV1IngestListGet(ctx).Execute()

List Ingested



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionAPI.ListIngestedV1IngestListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionAPI.ListIngestedV1IngestListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIngestedV1IngestListGet`: IngestResponse
	fmt.Fprintf(os.Stdout, "Response from `IngestionAPI.ListIngestedV1IngestListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIngestedV1IngestListGetRequest struct via the builder pattern


### Return type

[**IngestResponse**](IngestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

