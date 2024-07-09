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

// checks if the OpenAIMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenAIMessage{}

// OpenAIMessage Inference result, with the source of the message.  Role could be the assistant or system (providing a default response, not AI generated).
type OpenAIMessage struct {
	Role    *string        `json:"role,omitempty"`
	Content NullableString `json:"content"`
}

type _OpenAIMessage OpenAIMessage

// NewOpenAIMessage instantiates a new OpenAIMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenAIMessage(content NullableString) *OpenAIMessage {
	this := OpenAIMessage{}
	var role string = "user"
	this.Role = &role
	this.Content = content
	return &this
}

// NewOpenAIMessageWithDefaults instantiates a new OpenAIMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenAIMessageWithDefaults() *OpenAIMessage {
	this := OpenAIMessage{}
	var role string = "user"
	this.Role = &role
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OpenAIMessage) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIMessage) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OpenAIMessage) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OpenAIMessage) SetRole(v string) {
	o.Role = &v
}

// GetContent returns the Content field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenAIMessage) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}

	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenAIMessage) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// SetContent sets field value
func (o *OpenAIMessage) SetContent(v string) {
	o.Content.Set(&v)
}

func (o OpenAIMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenAIMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	toSerialize["content"] = o.Content.Get()
	return toSerialize, nil
}

func (o *OpenAIMessage) UnmarshalJSON(data []byte) (err error) {
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

	varOpenAIMessage := _OpenAIMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenAIMessage)

	if err != nil {
		return err
	}

	*o = OpenAIMessage(varOpenAIMessage)

	return err
}

type NullableOpenAIMessage struct {
	value *OpenAIMessage
	isSet bool
}

func (v NullableOpenAIMessage) Get() *OpenAIMessage {
	return v.value
}

func (v *NullableOpenAIMessage) Set(val *OpenAIMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenAIMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenAIMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenAIMessage(val *OpenAIMessage) *NullableOpenAIMessage {
	return &NullableOpenAIMessage{value: val, isSet: true}
}

func (v NullableOpenAIMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenAIMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
