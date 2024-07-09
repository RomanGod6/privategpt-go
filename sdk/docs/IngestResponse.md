# IngestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **interface{}** |  | 
**Model** | **interface{}** |  | 
**Data** | [**[]IngestedDoc**](IngestedDoc.md) |  | 

## Methods

### NewIngestResponse

`func NewIngestResponse(object interface{}, model interface{}, data []IngestedDoc, ) *IngestResponse`

NewIngestResponse instantiates a new IngestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestResponseWithDefaults

`func NewIngestResponseWithDefaults() *IngestResponse`

NewIngestResponseWithDefaults instantiates a new IngestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *IngestResponse) GetObject() interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IngestResponse) GetObjectOk() (*interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IngestResponse) SetObject(v interface{})`

SetObject sets Object field to given value.


### SetObjectNil

`func (o *IngestResponse) SetObjectNil(b bool)`

 SetObjectNil sets the value for Object to be an explicit nil

### UnsetObject
`func (o *IngestResponse) UnsetObject()`

UnsetObject ensures that no value is present for Object, not even an explicit nil
### GetModel

`func (o *IngestResponse) GetModel() interface{}`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *IngestResponse) GetModelOk() (*interface{}, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *IngestResponse) SetModel(v interface{})`

SetModel sets Model field to given value.


### SetModelNil

`func (o *IngestResponse) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *IngestResponse) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetData

`func (o *IngestResponse) GetData() []IngestedDoc`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IngestResponse) GetDataOk() (*[]IngestedDoc, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IngestResponse) SetData(v []IngestedDoc)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


