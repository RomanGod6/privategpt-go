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

// checks if the OpenAIChoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenAIChoice{}

// OpenAIChoice Response from AI.  Either the delta or the message will be present, but never both. Sources used will be returned in case context retrieval was enabled.
type OpenAIChoice struct {
	FinishReason NullableString        `json:"finish_reason"`
	Delta        NullableOpenAIDelta   `json:"delta,omitempty"`
	Message      NullableOpenAIMessage `json:"message,omitempty"`
	Sources      []Chunk               `json:"sources,omitempty"`
	Index        *int32                `json:"index,omitempty"`
}

type _OpenAIChoice OpenAIChoice

// NewOpenAIChoice instantiates a new OpenAIChoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenAIChoice(finishReason NullableString) *OpenAIChoice {
	this := OpenAIChoice{}
	this.FinishReason = finishReason
	var index int32 = 0
	this.Index = &index
	return &this
}

// NewOpenAIChoiceWithDefaults instantiates a new OpenAIChoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenAIChoiceWithDefaults() *OpenAIChoice {
	this := OpenAIChoice{}
	var index int32 = 0
	this.Index = &index
	return &this
}

// GetFinishReason returns the FinishReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenAIChoice) GetFinishReason() string {
	if o == nil || o.FinishReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.FinishReason.Get()
}

// GetFinishReasonOk returns a tuple with the FinishReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenAIChoice) GetFinishReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishReason.Get(), o.FinishReason.IsSet()
}

// SetFinishReason sets field value
func (o *OpenAIChoice) SetFinishReason(v string) {
	o.FinishReason.Set(&v)
}

// GetDelta returns the Delta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenAIChoice) GetDelta() OpenAIDelta {
	if o == nil || IsNil(o.Delta.Get()) {
		var ret OpenAIDelta
		return ret
	}
	return *o.Delta.Get()
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenAIChoice) GetDeltaOk() (*OpenAIDelta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Delta.Get(), o.Delta.IsSet()
}

// HasDelta returns a boolean if a field has been set.
func (o *OpenAIChoice) HasDelta() bool {
	if o != nil && o.Delta.IsSet() {
		return true
	}

	return false
}

// SetDelta gets a reference to the given NullableOpenAIDelta and assigns it to the Delta field.
func (o *OpenAIChoice) SetDelta(v OpenAIDelta) {
	o.Delta.Set(&v)
}

// SetDeltaNil sets the value for Delta to be an explicit nil
func (o *OpenAIChoice) SetDeltaNil() {
	o.Delta.Set(nil)
}

// UnsetDelta ensures that no value is present for Delta, not even an explicit nil
func (o *OpenAIChoice) UnsetDelta() {
	o.Delta.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenAIChoice) GetMessage() OpenAIMessage {
	if o == nil || IsNil(o.Message.Get()) {
		var ret OpenAIMessage
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenAIChoice) GetMessageOk() (*OpenAIMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *OpenAIChoice) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableOpenAIMessage and assigns it to the Message field.
func (o *OpenAIChoice) SetMessage(v OpenAIMessage) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *OpenAIChoice) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *OpenAIChoice) UnsetMessage() {
	o.Message.Unset()
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenAIChoice) GetSources() []Chunk {
	if o == nil {
		var ret []Chunk
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenAIChoice) GetSourcesOk() ([]Chunk, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *OpenAIChoice) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []Chunk and assigns it to the Sources field.
func (o *OpenAIChoice) SetSources(v []Chunk) {
	o.Sources = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *OpenAIChoice) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIChoice) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *OpenAIChoice) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *OpenAIChoice) SetIndex(v int32) {
	o.Index = &v
}

func (o OpenAIChoice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenAIChoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["finish_reason"] = o.FinishReason.Get()
	if o.Delta.IsSet() {
		toSerialize["delta"] = o.Delta.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	return toSerialize, nil
}

func (o *OpenAIChoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"finish_reason",
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

	varOpenAIChoice := _OpenAIChoice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenAIChoice)

	if err != nil {
		return err
	}

	*o = OpenAIChoice(varOpenAIChoice)

	return err
}

type NullableOpenAIChoice struct {
	value *OpenAIChoice
	isSet bool
}

func (v NullableOpenAIChoice) Get() *OpenAIChoice {
	return v.value
}

func (v *NullableOpenAIChoice) Set(val *OpenAIChoice) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenAIChoice) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenAIChoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenAIChoice(val *OpenAIChoice) *NullableOpenAIChoice {
	return &NullableOpenAIChoice{value: val, isSet: true}
}

func (v NullableOpenAIChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenAIChoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
