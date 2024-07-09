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

// checks if the Embedding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Embedding{}

// Embedding struct for Embedding
type Embedding struct {
	Index     int32       `json:"index"`
	Object    interface{} `json:"object"`
	Embedding []float32   `json:"embedding"`
}

type _Embedding Embedding

// NewEmbedding instantiates a new Embedding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbedding(index int32, object interface{}, embedding []float32) *Embedding {
	this := Embedding{}
	this.Index = index
	this.Object = object
	this.Embedding = embedding
	return &this
}

// NewEmbeddingWithDefaults instantiates a new Embedding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddingWithDefaults() *Embedding {
	this := Embedding{}
	return &this
}

// GetIndex returns the Index field value
func (o *Embedding) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *Embedding) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *Embedding) SetIndex(v int32) {
	o.Index = v
}

// GetObject returns the Object field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Embedding) GetObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Embedding) GetObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Embedding) SetObject(v interface{}) {
	o.Object = v
}

// GetEmbedding returns the Embedding field value
func (o *Embedding) GetEmbedding() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Embedding
}

// GetEmbeddingOk returns a tuple with the Embedding field value
// and a boolean to check if the value has been set.
func (o *Embedding) GetEmbeddingOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embedding, true
}

// SetEmbedding sets field value
func (o *Embedding) SetEmbedding(v []float32) {
	o.Embedding = v
}

func (o Embedding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Embedding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	toSerialize["embedding"] = o.Embedding
	return toSerialize, nil
}

func (o *Embedding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"object",
		"embedding",
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

	varEmbedding := _Embedding{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmbedding)

	if err != nil {
		return err
	}

	*o = Embedding(varEmbedding)

	return err
}

type NullableEmbedding struct {
	value *Embedding
	isSet bool
}

func (v NullableEmbedding) Get() *Embedding {
	return v.value
}

func (v *NullableEmbedding) Set(val *Embedding) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbedding) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbedding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbedding(val *Embedding) *NullableEmbedding {
	return &NullableEmbedding{value: val, isSet: true}
}

func (v NullableEmbedding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbedding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
