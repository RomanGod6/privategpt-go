# EmbeddingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **interface{}** |  | 
**Model** | **interface{}** |  | 
**Data** | [**[]Embedding**](Embedding.md) |  | 

## Methods

### NewEmbeddingsResponse

`func NewEmbeddingsResponse(object interface{}, model interface{}, data []Embedding, ) *EmbeddingsResponse`

NewEmbeddingsResponse instantiates a new EmbeddingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingsResponseWithDefaults

`func NewEmbeddingsResponseWithDefaults() *EmbeddingsResponse`

NewEmbeddingsResponseWithDefaults instantiates a new EmbeddingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *EmbeddingsResponse) GetObject() interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EmbeddingsResponse) GetObjectOk() (*interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EmbeddingsResponse) SetObject(v interface{})`

SetObject sets Object field to given value.


### SetObjectNil

`func (o *EmbeddingsResponse) SetObjectNil(b bool)`

 SetObjectNil sets the value for Object to be an explicit nil

### UnsetObject
`func (o *EmbeddingsResponse) UnsetObject()`

UnsetObject ensures that no value is present for Object, not even an explicit nil
### GetModel

`func (o *EmbeddingsResponse) GetModel() interface{}`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EmbeddingsResponse) GetModelOk() (*interface{}, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EmbeddingsResponse) SetModel(v interface{})`

SetModel sets Model field to given value.


### SetModelNil

`func (o *EmbeddingsResponse) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *EmbeddingsResponse) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetData

`func (o *EmbeddingsResponse) GetData() []Embedding`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmbeddingsResponse) GetDataOk() (*[]Embedding, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmbeddingsResponse) SetData(v []Embedding)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


