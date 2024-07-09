# Chunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **interface{}** |  | 
**Score** | **float32** |  | 
**Document** | [**IngestedDoc**](IngestedDoc.md) |  | 
**Text** | **string** |  | 
**PreviousTexts** | Pointer to **[]string** |  | [optional] 
**NextTexts** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChunk

`func NewChunk(object interface{}, score float32, document IngestedDoc, text string, ) *Chunk`

NewChunk instantiates a new Chunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChunkWithDefaults

`func NewChunkWithDefaults() *Chunk`

NewChunkWithDefaults instantiates a new Chunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Chunk) GetObject() interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Chunk) GetObjectOk() (*interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Chunk) SetObject(v interface{})`

SetObject sets Object field to given value.


### SetObjectNil

`func (o *Chunk) SetObjectNil(b bool)`

 SetObjectNil sets the value for Object to be an explicit nil

### UnsetObject
`func (o *Chunk) UnsetObject()`

UnsetObject ensures that no value is present for Object, not even an explicit nil
### GetScore

`func (o *Chunk) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Chunk) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Chunk) SetScore(v float32)`

SetScore sets Score field to given value.


### GetDocument

`func (o *Chunk) GetDocument() IngestedDoc`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *Chunk) GetDocumentOk() (*IngestedDoc, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *Chunk) SetDocument(v IngestedDoc)`

SetDocument sets Document field to given value.


### GetText

`func (o *Chunk) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Chunk) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Chunk) SetText(v string)`

SetText sets Text field to given value.


### GetPreviousTexts

`func (o *Chunk) GetPreviousTexts() []string`

GetPreviousTexts returns the PreviousTexts field if non-nil, zero value otherwise.

### GetPreviousTextsOk

`func (o *Chunk) GetPreviousTextsOk() (*[]string, bool)`

GetPreviousTextsOk returns a tuple with the PreviousTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTexts

`func (o *Chunk) SetPreviousTexts(v []string)`

SetPreviousTexts sets PreviousTexts field to given value.

### HasPreviousTexts

`func (o *Chunk) HasPreviousTexts() bool`

HasPreviousTexts returns a boolean if a field has been set.

### SetPreviousTextsNil

`func (o *Chunk) SetPreviousTextsNil(b bool)`

 SetPreviousTextsNil sets the value for PreviousTexts to be an explicit nil

### UnsetPreviousTexts
`func (o *Chunk) UnsetPreviousTexts()`

UnsetPreviousTexts ensures that no value is present for PreviousTexts, not even an explicit nil
### GetNextTexts

`func (o *Chunk) GetNextTexts() []string`

GetNextTexts returns the NextTexts field if non-nil, zero value otherwise.

### GetNextTextsOk

`func (o *Chunk) GetNextTextsOk() (*[]string, bool)`

GetNextTextsOk returns a tuple with the NextTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTexts

`func (o *Chunk) SetNextTexts(v []string)`

SetNextTexts sets NextTexts field to given value.

### HasNextTexts

`func (o *Chunk) HasNextTexts() bool`

HasNextTexts returns a boolean if a field has been set.

### SetNextTextsNil

`func (o *Chunk) SetNextTextsNil(b bool)`

 SetNextTextsNil sets the value for NextTexts to be an explicit nil

### UnsetNextTexts
`func (o *Chunk) UnsetNextTexts()`

UnsetNextTexts ensures that no value is present for NextTexts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


