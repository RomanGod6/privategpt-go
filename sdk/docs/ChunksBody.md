# ChunksBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**ContextFilter** | Pointer to [**NullableContextFilter**](ContextFilter.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] [default to 10]
**PrevNextChunks** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewChunksBody

`func NewChunksBody(text string, ) *ChunksBody`

NewChunksBody instantiates a new ChunksBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChunksBodyWithDefaults

`func NewChunksBodyWithDefaults() *ChunksBody`

NewChunksBodyWithDefaults instantiates a new ChunksBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ChunksBody) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChunksBody) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChunksBody) SetText(v string)`

SetText sets Text field to given value.


### GetContextFilter

`func (o *ChunksBody) GetContextFilter() ContextFilter`

GetContextFilter returns the ContextFilter field if non-nil, zero value otherwise.

### GetContextFilterOk

`func (o *ChunksBody) GetContextFilterOk() (*ContextFilter, bool)`

GetContextFilterOk returns a tuple with the ContextFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextFilter

`func (o *ChunksBody) SetContextFilter(v ContextFilter)`

SetContextFilter sets ContextFilter field to given value.

### HasContextFilter

`func (o *ChunksBody) HasContextFilter() bool`

HasContextFilter returns a boolean if a field has been set.

### SetContextFilterNil

`func (o *ChunksBody) SetContextFilterNil(b bool)`

 SetContextFilterNil sets the value for ContextFilter to be an explicit nil

### UnsetContextFilter
`func (o *ChunksBody) UnsetContextFilter()`

UnsetContextFilter ensures that no value is present for ContextFilter, not even an explicit nil
### GetLimit

`func (o *ChunksBody) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ChunksBody) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ChunksBody) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ChunksBody) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPrevNextChunks

`func (o *ChunksBody) GetPrevNextChunks() int32`

GetPrevNextChunks returns the PrevNextChunks field if non-nil, zero value otherwise.

### GetPrevNextChunksOk

`func (o *ChunksBody) GetPrevNextChunksOk() (*int32, bool)`

GetPrevNextChunksOk returns a tuple with the PrevNextChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevNextChunks

`func (o *ChunksBody) SetPrevNextChunks(v int32)`

SetPrevNextChunks sets PrevNextChunks field to given value.

### HasPrevNextChunks

`func (o *ChunksBody) HasPrevNextChunks() bool`

HasPrevNextChunks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


