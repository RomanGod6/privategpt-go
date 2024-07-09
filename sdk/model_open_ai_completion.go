/*
FastAPI- Private GPT

# SDK for GO LANG for Private GPT

API version: 0.1.0
*/
package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OpenAICompletion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenAICompletion{}

// OpenAICompletion Clone of OpenAI Completion model.  For more information see: https://platform.openai.com/docs/api-reference/chat/object
type OpenAICompletion struct {
	Id      string         `json:"id"`
	Object  *string        `json:"object,omitempty"`
	Created int32          `json:"created"`
	Model   interface{}    `json:"model"`
	Choices []OpenAIChoice `json:"choices"`
}

type _OpenAICompletion OpenAICompletion

// NewOpenAICompletion instantiates a new OpenAICompletion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenAICompletion(id string, created int32, model interface{}, choices []OpenAIChoice) *OpenAICompletion {
	this := OpenAICompletion{}
	this.Id = id
	var object string = "completion"
	this.Object = &object
	this.Created = created
	this.Model = model
	this.Choices = choices
	return &this
}

// NewOpenAICompletionWithDefaults instantiates a new OpenAICompletion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenAICompletionWithDefaults() *OpenAICompletion {
	this := OpenAICompletion{}
	var object string = "completion"
	this.Object = &object
	return &this
}

// GetId returns the Id field value
func (o *OpenAICompletion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpenAICompletion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OpenAICompletion) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OpenAICompletion) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICompletion) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OpenAICompletion) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OpenAICompletion) SetObject(v string) {
	o.Object = &v
}

// GetCreated returns the Created field value
func (o *OpenAICompletion) GetCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *OpenAICompletion) GetCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *OpenAICompletion) SetCreated(v int32) {
	o.Created = v
}

// GetModel returns the Model field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OpenAICompletion) GetModel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenAICompletion) GetModelOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *OpenAICompletion) SetModel(v interface{}) {
	o.Model = v
}

// GetChoices returns the Choices field value
func (o *OpenAICompletion) GetChoices() []OpenAIChoice {
	if o == nil {
		var ret []OpenAIChoice
		return ret
	}

	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value
// and a boolean to check if the value has been set.
func (o *OpenAICompletion) GetChoicesOk() ([]OpenAIChoice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Choices, true
}

// SetChoices sets field value
func (o *OpenAICompletion) SetChoices(v []OpenAIChoice) {
	o.Choices = v
}

func (o OpenAICompletion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenAICompletion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["created"] = o.Created
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	toSerialize["choices"] = o.Choices
	return toSerialize, nil
}

func (o *OpenAICompletion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created",
		"model",
		"choices",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOpenAICompletion := _OpenAICompletion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenAICompletion)

	if err != nil {
		return err
	}

	*o = OpenAICompletion(varOpenAICompletion)

	return err
}

type NullableOpenAICompletion struct {
	value *OpenAICompletion
	isSet bool
}

func (v NullableOpenAICompletion) Get() *OpenAICompletion {
	return v.value
}

func (v *NullableOpenAICompletion) Set(val *OpenAICompletion) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenAICompletion) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenAICompletion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenAICompletion(val *OpenAICompletion) *NullableOpenAICompletion {
	return &NullableOpenAICompletion{value: val, isSet: true}
}

func (v NullableOpenAICompletion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenAICompletion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
