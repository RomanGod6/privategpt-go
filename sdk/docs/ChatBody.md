# ChatBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]OpenAIMessage**](OpenAIMessage.md) |  | 
**UseContext** | Pointer to **bool** |  | [optional] [default to false]
**ContextFilter** | Pointer to [**NullableContextFilter**](ContextFilter.md) |  | [optional] 
**IncludeSources** | Pointer to **bool** |  | [optional] [default to true]
**Stream** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewChatBody

`func NewChatBody(messages []OpenAIMessage, ) *ChatBody`

NewChatBody instantiates a new ChatBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatBodyWithDefaults

`func NewChatBodyWithDefaults() *ChatBody`

NewChatBodyWithDefaults instantiates a new ChatBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ChatBody) GetMessages() []OpenAIMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatBody) GetMessagesOk() (*[]OpenAIMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatBody) SetMessages(v []OpenAIMessage)`

SetMessages sets Messages field to given value.


### GetUseContext

`func (o *ChatBody) GetUseContext() bool`

GetUseContext returns the UseContext field if non-nil, zero value otherwise.

### GetUseContextOk

`func (o *ChatBody) GetUseContextOk() (*bool, bool)`

GetUseContextOk returns a tuple with the UseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseContext

`func (o *ChatBody) SetUseContext(v bool)`

SetUseContext sets UseContext field to given value.

### HasUseContext

`func (o *ChatBody) HasUseContext() bool`

HasUseContext returns a boolean if a field has been set.

### GetContextFilter

`func (o *ChatBody) GetContextFilter() ContextFilter`

GetContextFilter returns the ContextFilter field if non-nil, zero value otherwise.

### GetContextFilterOk

`func (o *ChatBody) GetContextFilterOk() (*ContextFilter, bool)`

GetContextFilterOk returns a tuple with the ContextFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextFilter

`func (o *ChatBody) SetContextFilter(v ContextFilter)`

SetContextFilter sets ContextFilter field to given value.

### HasContextFilter

`func (o *ChatBody) HasContextFilter() bool`

HasContextFilter returns a boolean if a field has been set.

### SetContextFilterNil

`func (o *ChatBody) SetContextFilterNil(b bool)`

 SetContextFilterNil sets the value for ContextFilter to be an explicit nil

### UnsetContextFilter
`func (o *ChatBody) UnsetContextFilter()`

UnsetContextFilter ensures that no value is present for ContextFilter, not even an explicit nil
### GetIncludeSources

`func (o *ChatBody) GetIncludeSources() bool`

GetIncludeSources returns the IncludeSources field if non-nil, zero value otherwise.

### GetIncludeSourcesOk

`func (o *ChatBody) GetIncludeSourcesOk() (*bool, bool)`

GetIncludeSourcesOk returns a tuple with the IncludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSources

`func (o *ChatBody) SetIncludeSources(v bool)`

SetIncludeSources sets IncludeSources field to given value.

### HasIncludeSources

`func (o *ChatBody) HasIncludeSources() bool`

HasIncludeSources returns a boolean if a field has been set.

### GetStream

`func (o *ChatBody) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatBody) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatBody) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatBody) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


