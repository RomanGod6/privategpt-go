/*
FastAPI- Private GPT

# SDK for GO LANG for Private GPT

API version: 0.1.0
*/
package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ChatBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatBody{}

// ChatBody struct for ChatBody
type ChatBody struct {
	Messages       []OpenAIMessage       `json:"messages"`
	UseContext     *bool                 `json:"use_context,omitempty"`
	ContextFilter  NullableContextFilter `json:"context_filter,omitempty"`
	IncludeSources *bool                 `json:"include_sources,omitempty"`
	Stream         *bool                 `json:"stream,omitempty"`
}

type _ChatBody ChatBody

// NewChatBody instantiates a new ChatBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatBody(messages []OpenAIMessage) *ChatBody {
	this := ChatBody{}
	this.Messages = messages
	var useContext bool = false
	this.UseContext = &useContext
	var includeSources bool = true
	this.IncludeSources = &includeSources
	var stream bool = false
	this.Stream = &stream
	return &this
}

// NewChatBodyWithDefaults instantiates a new ChatBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatBodyWithDefaults() *ChatBody {
	this := ChatBody{}
	var useContext bool = false
	this.UseContext = &useContext
	var includeSources bool = true
	this.IncludeSources = &includeSources
	var stream bool = false
	this.Stream = &stream
	return &this
}

// GetMessages returns the Messages field value
func (o *ChatBody) GetMessages() []OpenAIMessage {
	if o == nil {
		var ret []OpenAIMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *ChatBody) GetMessagesOk() ([]OpenAIMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *ChatBody) SetMessages(v []OpenAIMessage) {
	o.Messages = v
}

// GetUseContext returns the UseContext field value if set, zero value otherwise.
func (o *ChatBody) GetUseContext() bool {
	if o == nil || IsNil(o.UseContext) {
		var ret bool
		return ret
	}
	return *o.UseContext
}

// GetUseContextOk returns a tuple with the UseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatBody) GetUseContextOk() (*bool, bool) {
	if o == nil || IsNil(o.UseContext) {
		return nil, false
	}
	return o.UseContext, true
}

// HasUseContext returns a boolean if a field has been set.
func (o *ChatBody) HasUseContext() bool {
	if o != nil && !IsNil(o.UseContext) {
		return true
	}

	return false
}

// SetUseContext gets a reference to the given bool and assigns it to the UseContext field.
func (o *ChatBody) SetUseContext(v bool) {
	o.UseContext = &v
}

// GetContextFilter returns the ContextFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatBody) GetContextFilter() ContextFilter {
	if o == nil || IsNil(o.ContextFilter.Get()) {
		var ret ContextFilter
		return ret
	}
	return *o.ContextFilter.Get()
}

// GetContextFilterOk returns a tuple with the ContextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatBody) GetContextFilterOk() (*ContextFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContextFilter.Get(), o.ContextFilter.IsSet()
}

// HasContextFilter returns a boolean if a field has been set.
func (o *ChatBody) HasContextFilter() bool {
	if o != nil && o.ContextFilter.IsSet() {
		return true
	}

	return false
}

// SetContextFilter gets a reference to the given NullableContextFilter and assigns it to the ContextFilter field.
func (o *ChatBody) SetContextFilter(v ContextFilter) {
	o.ContextFilter.Set(&v)
}

// SetContextFilterNil sets the value for ContextFilter to be an explicit nil
func (o *ChatBody) SetContextFilterNil() {
	o.ContextFilter.Set(nil)
}

// UnsetContextFilter ensures that no value is present for ContextFilter, not even an explicit nil
func (o *ChatBody) UnsetContextFilter() {
	o.ContextFilter.Unset()
}

// GetIncludeSources returns the IncludeSources field value if set, zero value otherwise.
func (o *ChatBody) GetIncludeSources() bool {
	if o == nil || IsNil(o.IncludeSources) {
		var ret bool
		return ret
	}
	return *o.IncludeSources
}

// GetIncludeSourcesOk returns a tuple with the IncludeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatBody) GetIncludeSourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSources) {
		return nil, false
	}
	return o.IncludeSources, true
}

// HasIncludeSources returns a boolean if a field has been set.
func (o *ChatBody) HasIncludeSources() bool {
	if o != nil && !IsNil(o.IncludeSources) {
		return true
	}

	return false
}

// SetIncludeSources gets a reference to the given bool and assigns it to the IncludeSources field.
func (o *ChatBody) SetIncludeSources(v bool) {
	o.IncludeSources = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *ChatBody) GetStream() bool {
	if o == nil || IsNil(o.Stream) {
		var ret bool
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatBody) GetStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *ChatBody) HasStream() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given bool and assigns it to the Stream field.
func (o *ChatBody) SetStream(v bool) {
	o.Stream = &v
}

func (o ChatBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messages"] = o.Messages
	if !IsNil(o.UseContext) {
		toSerialize["use_context"] = o.UseContext
	}
	if o.ContextFilter.IsSet() {
		toSerialize["context_filter"] = o.ContextFilter.Get()
	}
	if !IsNil(o.IncludeSources) {
		toSerialize["include_sources"] = o.IncludeSources
	}
	if !IsNil(o.Stream) {
		toSerialize["stream"] = o.Stream
	}
	return toSerialize, nil
}

func (o *ChatBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messages",
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

	varChatBody := _ChatBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatBody)

	if err != nil {
		return err
	}

	*o = ChatBody(varChatBody)

	return err
}

type NullableChatBody struct {
	value *ChatBody
	isSet bool
}

func (v NullableChatBody) Get() *ChatBody {
	return v.value
}

func (v *NullableChatBody) Set(val *ChatBody) {
	v.value = val
	v.isSet = true
}

func (v NullableChatBody) IsSet() bool {
	return v.isSet
}

func (v *NullableChatBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatBody(val *ChatBody) *NullableChatBody {
	return &NullableChatBody{value: val, isSet: true}
}

func (v NullableChatBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
