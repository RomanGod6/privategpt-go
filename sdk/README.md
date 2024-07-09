# Go API client for openapi



## Overview
- API version: 0.1.0
- Package version: 1.0.0
- Generator version: 7.7.0


## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/RomanGod6/privategpt-go/sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ContextChunksAPI* | [**ChunksRetrievalV1ChunksPost**](docs/ContextChunksAPI.md#chunksretrievalv1chunkspost) | **Post** /v1/chunks | Chunks Retrieval
*ContextualCompletionsAPI* | [**ChatCompletionV1ChatCompletionsPost**](docs/ContextualCompletionsAPI.md#chatcompletionv1chatcompletionspost) | **Post** /v1/chat/completions | Chat Completion
*ContextualCompletionsAPI* | [**PromptCompletionV1CompletionsPost**](docs/ContextualCompletionsAPI.md#promptcompletionv1completionspost) | **Post** /v1/completions | Completion
*EmbeddingsAPI* | [**EmbeddingsGenerationV1EmbeddingsPost**](docs/EmbeddingsAPI.md#embeddingsgenerationv1embeddingspost) | **Post** /v1/embeddings | Embeddings Generation
*HealthAPI* | [**HealthHealthGet**](docs/HealthAPI.md#healthhealthget) | **Get** /health | Health
*IngestionAPI* | [**DeleteIngestedV1IngestDocIdDelete**](docs/IngestionAPI.md#deleteingestedv1ingestdociddelete) | **Delete** /v1/ingest/{doc_id} | Delete Ingested
*IngestionAPI* | [**IngestFileV1IngestFilePost**](docs/IngestionAPI.md#ingestfilev1ingestfilepost) | **Post** /v1/ingest/file | Ingest File
*IngestionAPI* | [**IngestTextV1IngestTextPost**](docs/IngestionAPI.md#ingesttextv1ingesttextpost) | **Post** /v1/ingest/text | Ingest Text
*IngestionAPI* | [**IngestV1IngestPost**](docs/IngestionAPI.md#ingestv1ingestpost) | **Post** /v1/ingest | Ingest
*IngestionAPI* | [**ListIngestedV1IngestListGet**](docs/IngestionAPI.md#listingestedv1ingestlistget) | **Get** /v1/ingest/list | List Ingested


## Documentation For Models

 - [ChatBody](docs/ChatBody.md)
 - [Chunk](docs/Chunk.md)
 - [ChunksBody](docs/ChunksBody.md)
 - [ChunksResponse](docs/ChunksResponse.md)
 - [CompletionsBody](docs/CompletionsBody.md)
 - [ContextFilter](docs/ContextFilter.md)
 - [Embedding](docs/Embedding.md)
 - [EmbeddingsBody](docs/EmbeddingsBody.md)
 - [EmbeddingsResponse](docs/EmbeddingsResponse.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [HealthResponse](docs/HealthResponse.md)
 - [IngestResponse](docs/IngestResponse.md)
 - [IngestTextBody](docs/IngestTextBody.md)
 - [IngestedDoc](docs/IngestedDoc.md)
 - [Input](docs/Input.md)
 - [OpenAIChoice](docs/OpenAIChoice.md)
 - [OpenAICompletion](docs/OpenAICompletion.md)
 - [OpenAIDelta](docs/OpenAIDelta.md)
 - [OpenAIMessage](docs/OpenAIMessage.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorLocInner](docs/ValidationErrorLocInner.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



