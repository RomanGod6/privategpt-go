# CompletionsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | **string** |  | 
**SystemPrompt** | Pointer to **NullableString** |  | [optional] 
**UseContext** | Pointer to **bool** |  | [optional] [default to false]
**ContextFilter** | Pointer to [**NullableContextFilter**](ContextFilter.md) |  | [optional] 
**IncludeSources** | Pointer to **bool** |  | [optional] [default to true]
**Stream** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCompletionsBody

`func NewCompletionsBody(prompt string, ) *CompletionsBody`

NewCompletionsBody instantiates a new CompletionsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionsBodyWithDefaults

`func NewCompletionsBodyWithDefaults() *CompletionsBody`

NewCompletionsBodyWithDefaults instantiates a new CompletionsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *CompletionsBody) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CompletionsBody) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CompletionsBody) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetSystemPrompt

`func (o *CompletionsBody) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *CompletionsBody) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *CompletionsBody) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *CompletionsBody) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *CompletionsBody) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *CompletionsBody) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
### GetUseContext

`func (o *CompletionsBody) GetUseContext() bool`

GetUseContext returns the UseContext field if non-nil, zero value otherwise.

### GetUseContextOk

`func (o *CompletionsBody) GetUseContextOk() (*bool, bool)`

GetUseContextOk returns a tuple with the UseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseContext

`func (o *CompletionsBody) SetUseContext(v bool)`

SetUseContext sets UseContext field to given value.

### HasUseContext

`func (o *CompletionsBody) HasUseContext() bool`

HasUseContext returns a boolean if a field has been set.

### GetContextFilter

`func (o *CompletionsBody) GetContextFilter() ContextFilter`

GetContextFilter returns the ContextFilter field if non-nil, zero value otherwise.

### GetContextFilterOk

`func (o *CompletionsBody) GetContextFilterOk() (*ContextFilter, bool)`

GetContextFilterOk returns a tuple with the ContextFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextFilter

`func (o *CompletionsBody) SetContextFilter(v ContextFilter)`

SetContextFilter sets ContextFilter field to given value.

### HasContextFilter

`func (o *CompletionsBody) HasContextFilter() bool`

HasContextFilter returns a boolean if a field has been set.

### SetContextFilterNil

`func (o *CompletionsBody) SetContextFilterNil(b bool)`

 SetContextFilterNil sets the value for ContextFilter to be an explicit nil

### UnsetContextFilter
`func (o *CompletionsBody) UnsetContextFilter()`

UnsetContextFilter ensures that no value is present for ContextFilter, not even an explicit nil
### GetIncludeSources

`func (o *CompletionsBody) GetIncludeSources() bool`

GetIncludeSources returns the IncludeSources field if non-nil, zero value otherwise.

### GetIncludeSourcesOk

`func (o *CompletionsBody) GetIncludeSourcesOk() (*bool, bool)`

GetIncludeSourcesOk returns a tuple with the IncludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSources

`func (o *CompletionsBody) SetIncludeSources(v bool)`

SetIncludeSources sets IncludeSources field to given value.

### HasIncludeSources

`func (o *CompletionsBody) HasIncludeSources() bool`

HasIncludeSources returns a boolean if a field has been set.

### GetStream

`func (o *CompletionsBody) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CompletionsBody) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CompletionsBody) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CompletionsBody) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


