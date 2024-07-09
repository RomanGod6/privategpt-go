# OpenAIMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] [default to "user"]
**Content** | **NullableString** |  | 

## Methods

### NewOpenAIMessage

`func NewOpenAIMessage(content NullableString, ) *OpenAIMessage`

NewOpenAIMessage instantiates a new OpenAIMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenAIMessageWithDefaults

`func NewOpenAIMessageWithDefaults() *OpenAIMessage`

NewOpenAIMessageWithDefaults instantiates a new OpenAIMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OpenAIMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OpenAIMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OpenAIMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OpenAIMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *OpenAIMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OpenAIMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OpenAIMessage) SetContent(v string)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *OpenAIMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *OpenAIMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


