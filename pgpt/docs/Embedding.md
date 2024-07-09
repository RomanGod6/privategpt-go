# Embedding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Object** | **interface{}** |  | 
**Embedding** | **[]float32** |  | 

## Methods

### NewEmbedding

`func NewEmbedding(index int32, object interface{}, embedding []float32, ) *Embedding`

NewEmbedding instantiates a new Embedding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingWithDefaults

`func NewEmbeddingWithDefaults() *Embedding`

NewEmbeddingWithDefaults instantiates a new Embedding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Embedding) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Embedding) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Embedding) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetObject

`func (o *Embedding) GetObject() interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Embedding) GetObjectOk() (*interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Embedding) SetObject(v interface{})`

SetObject sets Object field to given value.


### SetObjectNil

`func (o *Embedding) SetObjectNil(b bool)`

 SetObjectNil sets the value for Object to be an explicit nil

### UnsetObject
`func (o *Embedding) UnsetObject()`

UnsetObject ensures that no value is present for Object, not even an explicit nil
### GetEmbedding

`func (o *Embedding) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *Embedding) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *Embedding) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


