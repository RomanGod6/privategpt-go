# OpenAICompletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "completion"]
**Created** | **int32** |  | 
**Model** | **interface{}** |  | 
**Choices** | [**[]OpenAIChoice**](OpenAIChoice.md) |  | 

## Methods

### NewOpenAICompletion

`func NewOpenAICompletion(id string, created int32, model interface{}, choices []OpenAIChoice, ) *OpenAICompletion`

NewOpenAICompletion instantiates a new OpenAICompletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenAICompletionWithDefaults

`func NewOpenAICompletionWithDefaults() *OpenAICompletion`

NewOpenAICompletionWithDefaults instantiates a new OpenAICompletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenAICompletion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenAICompletion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenAICompletion) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *OpenAICompletion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OpenAICompletion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OpenAICompletion) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OpenAICompletion) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *OpenAICompletion) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OpenAICompletion) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OpenAICompletion) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetModel

`func (o *OpenAICompletion) GetModel() interface{}`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OpenAICompletion) GetModelOk() (*interface{}, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OpenAICompletion) SetModel(v interface{})`

SetModel sets Model field to given value.


### SetModelNil

`func (o *OpenAICompletion) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *OpenAICompletion) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetChoices

`func (o *OpenAICompletion) GetChoices() []OpenAIChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *OpenAICompletion) GetChoicesOk() (*[]OpenAIChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *OpenAICompletion) SetChoices(v []OpenAIChoice)`

SetChoices sets Choices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


