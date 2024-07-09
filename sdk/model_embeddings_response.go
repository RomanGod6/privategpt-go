/*
FastAPI- Private GPT

SDK for GO LANG for Private GPT

API version: 0.1.0
*/

package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EmbeddingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddingsResponse{}

// EmbeddingsResponse struct for EmbeddingsResponse
type EmbeddingsResponse struct {
	Object interface{} `json:"object"`
	Model  interface{} `json:"model"`
	Data   []Embedding `json:"data"`
}

type _EmbeddingsResponse EmbeddingsResponse

// NewEmbeddingsResponse instantiates a new EmbeddingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddingsResponse(object interface{}, model interface{}, data []Embedding) *EmbeddingsResponse {
	this := EmbeddingsResponse{}
	this.Object = object
	this.Model = model
	this.Data = data
	return &this
}

// NewEmbeddingsResponseWithDefaults instantiates a new EmbeddingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddingsResponseWithDefaults() *EmbeddingsResponse {
	this := EmbeddingsResponse{}
	return &this
}

// GetObject returns the Object field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *EmbeddingsResponse) GetObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbeddingsResponse) GetObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *EmbeddingsResponse) SetObject(v interface{}) {
	o.Object = v
}

// GetModel returns the Model field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *EmbeddingsResponse) GetModel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbeddingsResponse) GetModelOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *EmbeddingsResponse) SetModel(v interface{}) {
	o.Model = v
}

// GetData returns the Data field value
func (o *EmbeddingsResponse) GetData() []Embedding {
	if o == nil {
		var ret []Embedding
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EmbeddingsResponse) GetDataOk() ([]Embedding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *EmbeddingsResponse) SetData(v []Embedding) {
	o.Data = v
}

func (o EmbeddingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EmbeddingsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"model",
		"data",
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

	varEmbeddingsResponse := _EmbeddingsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmbeddingsResponse)

	if err != nil {
		return err
	}

	*o = EmbeddingsResponse(varEmbeddingsResponse)

	return err
}

type NullableEmbeddingsResponse struct {
	value *EmbeddingsResponse
	isSet bool
}

func (v NullableEmbeddingsResponse) Get() *EmbeddingsResponse {
	return v.value
}

func (v *NullableEmbeddingsResponse) Set(val *EmbeddingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddingsResponse(val *EmbeddingsResponse) *NullableEmbeddingsResponse {
	return &NullableEmbeddingsResponse{value: val, isSet: true}
}

func (v NullableEmbeddingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
