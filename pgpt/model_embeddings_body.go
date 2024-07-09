/*
FastAPI- Private GPT

SDK for GO LANG for Private GPT

API version: 0.1.0
*/

package pgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EmbeddingsBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddingsBody{}

// EmbeddingsBody struct for EmbeddingsBody
type EmbeddingsBody struct {
	Input Input `json:"input"`
}

type _EmbeddingsBody EmbeddingsBody

// NewEmbeddingsBody instantiates a new EmbeddingsBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddingsBody(input Input) *EmbeddingsBody {
	this := EmbeddingsBody{}
	this.Input = input
	return &this
}

// NewEmbeddingsBodyWithDefaults instantiates a new EmbeddingsBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddingsBodyWithDefaults() *EmbeddingsBody {
	this := EmbeddingsBody{}
	return &this
}

// GetInput returns the Input field value
func (o *EmbeddingsBody) GetInput() Input {
	if o == nil {
		var ret Input
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *EmbeddingsBody) GetInputOk() (*Input, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *EmbeddingsBody) SetInput(v Input) {
	o.Input = v
}

func (o EmbeddingsBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddingsBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input"] = o.Input
	return toSerialize, nil
}

func (o *EmbeddingsBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"input",
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

	varEmbeddingsBody := _EmbeddingsBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmbeddingsBody)

	if err != nil {
		return err
	}

	*o = EmbeddingsBody(varEmbeddingsBody)

	return err
}

type NullableEmbeddingsBody struct {
	value *EmbeddingsBody
	isSet bool
}

func (v NullableEmbeddingsBody) Get() *EmbeddingsBody {
	return v.value
}

func (v *NullableEmbeddingsBody) Set(val *EmbeddingsBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddingsBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddingsBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddingsBody(val *EmbeddingsBody) *NullableEmbeddingsBody {
	return &NullableEmbeddingsBody{value: val, isSet: true}
}

func (v NullableEmbeddingsBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddingsBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
