# OpenAIChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishReason** | **NullableString** |  | 
**Delta** | Pointer to [**NullableOpenAIDelta**](OpenAIDelta.md) |  | [optional] 
**Message** | Pointer to [**NullableOpenAIMessage**](OpenAIMessage.md) |  | [optional] 
**Sources** | Pointer to [**[]Chunk**](Chunk.md) |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewOpenAIChoice

`func NewOpenAIChoice(finishReason NullableString, ) *OpenAIChoice`

NewOpenAIChoice instantiates a new OpenAIChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenAIChoiceWithDefaults

`func NewOpenAIChoiceWithDefaults() *OpenAIChoice`

NewOpenAIChoiceWithDefaults instantiates a new OpenAIChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishReason

`func (o *OpenAIChoice) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *OpenAIChoice) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *OpenAIChoice) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.


### SetFinishReasonNil

`func (o *OpenAIChoice) SetFinishReasonNil(b bool)`

 SetFinishReasonNil sets the value for FinishReason to be an explicit nil

### UnsetFinishReason
`func (o *OpenAIChoice) UnsetFinishReason()`

UnsetFinishReason ensures that no value is present for FinishReason, not even an explicit nil
### GetDelta

`func (o *OpenAIChoice) GetDelta() OpenAIDelta`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *OpenAIChoice) GetDeltaOk() (*OpenAIDelta, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *OpenAIChoice) SetDelta(v OpenAIDelta)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *OpenAIChoice) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### SetDeltaNil

`func (o *OpenAIChoice) SetDeltaNil(b bool)`

 SetDeltaNil sets the value for Delta to be an explicit nil

### UnsetDelta
`func (o *OpenAIChoice) UnsetDelta()`

UnsetDelta ensures that no value is present for Delta, not even an explicit nil
### GetMessage

`func (o *OpenAIChoice) GetMessage() OpenAIMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OpenAIChoice) GetMessageOk() (*OpenAIMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OpenAIChoice) SetMessage(v OpenAIMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OpenAIChoice) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *OpenAIChoice) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *OpenAIChoice) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSources

`func (o *OpenAIChoice) GetSources() []Chunk`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *OpenAIChoice) GetSourcesOk() (*[]Chunk, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *OpenAIChoice) SetSources(v []Chunk)`

SetSources sets Sources field to given value.

### HasSources

`func (o *OpenAIChoice) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *OpenAIChoice) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *OpenAIChoice) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetIndex

`func (o *OpenAIChoice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OpenAIChoice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OpenAIChoice) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *OpenAIChoice) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


