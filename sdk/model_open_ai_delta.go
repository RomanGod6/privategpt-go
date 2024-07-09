/*
FastAPI- Private GPT

SDK for GO LANG for Private GPT

API version: 0.1.0
*/

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OpenAIDelta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenAIDelta{}

// OpenAIDelta A piece of completion that needs to be concatenated to get the full message.
type OpenAIDelta struct {
	Content NullableString `json:"content"`
}

type _OpenAIDelta OpenAIDelta

// NewOpenAIDelta instantiates a new OpenAIDelta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenAIDelta(content NullableString) *OpenAIDelta {
	this := OpenAIDelta{}
	this.Content = content
	return &this
}

// NewOpenAIDeltaWithDefaults instantiates a new OpenAIDelta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenAIDeltaWithDefaults() *OpenAIDelta {
	this := OpenAIDelta{}
	return &this
}

// GetContent returns the Content field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenAIDelta) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}

	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenAIDelta) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// SetContent sets field value
func (o *OpenAIDelta) SetContent(v string) {
	o.Content.Set(&v)
}

func (o OpenAIDelta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenAIDelta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content.Get()
	return toSerialize, nil
}

func (o *OpenAIDelta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
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

	varOpenAIDelta := _OpenAIDelta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenAIDelta)

	if err != nil {
		return err
	}

	*o = OpenAIDelta(varOpenAIDelta)

	return err
}

type NullableOpenAIDelta struct {
	value *OpenAIDelta
	isSet bool
}

func (v NullableOpenAIDelta) Get() *OpenAIDelta {
	return v.value
}

func (v *NullableOpenAIDelta) Set(val *OpenAIDelta) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenAIDelta) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenAIDelta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenAIDelta(val *OpenAIDelta) *NullableOpenAIDelta {
	return &NullableOpenAIDelta{value: val, isSet: true}
}

func (v NullableOpenAIDelta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenAIDelta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
