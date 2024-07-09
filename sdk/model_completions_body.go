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

// checks if the CompletionsBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompletionsBody{}

// CompletionsBody struct for CompletionsBody
type CompletionsBody struct {
	Prompt         string                `json:"prompt"`
	SystemPrompt   NullableString        `json:"system_prompt,omitempty"`
	UseContext     *bool                 `json:"use_context,omitempty"`
	ContextFilter  NullableContextFilter `json:"context_filter,omitempty"`
	IncludeSources *bool                 `json:"include_sources,omitempty"`
	Stream         *bool                 `json:"stream,omitempty"`
}

type _CompletionsBody CompletionsBody

// NewCompletionsBody instantiates a new CompletionsBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletionsBody(prompt string) *CompletionsBody {
	this := CompletionsBody{}
	this.Prompt = prompt
	var useContext bool = false
	this.UseContext = &useContext
	var includeSources bool = true
	this.IncludeSources = &includeSources
	var stream bool = false
	this.Stream = &stream
	return &this
}

// NewCompletionsBodyWithDefaults instantiates a new CompletionsBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletionsBodyWithDefaults() *CompletionsBody {
	this := CompletionsBody{}
	var useContext bool = false
	this.UseContext = &useContext
	var includeSources bool = true
	this.IncludeSources = &includeSources
	var stream bool = false
	this.Stream = &stream
	return &this
}

// GetPrompt returns the Prompt field value
func (o *CompletionsBody) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *CompletionsBody) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *CompletionsBody) SetPrompt(v string) {
	o.Prompt = v
}

// GetSystemPrompt returns the SystemPrompt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompletionsBody) GetSystemPrompt() string {
	if o == nil || IsNil(o.SystemPrompt.Get()) {
		var ret string
		return ret
	}
	return *o.SystemPrompt.Get()
}

// GetSystemPromptOk returns a tuple with the SystemPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompletionsBody) GetSystemPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SystemPrompt.Get(), o.SystemPrompt.IsSet()
}

// HasSystemPrompt returns a boolean if a field has been set.
func (o *CompletionsBody) HasSystemPrompt() bool {
	if o != nil && o.SystemPrompt.IsSet() {
		return true
	}

	return false
}

// SetSystemPrompt gets a reference to the given NullableString and assigns it to the SystemPrompt field.
func (o *CompletionsBody) SetSystemPrompt(v string) {
	o.SystemPrompt.Set(&v)
}

// SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil
func (o *CompletionsBody) SetSystemPromptNil() {
	o.SystemPrompt.Set(nil)
}

// UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
func (o *CompletionsBody) UnsetSystemPrompt() {
	o.SystemPrompt.Unset()
}

// GetUseContext returns the UseContext field value if set, zero value otherwise.
func (o *CompletionsBody) GetUseContext() bool {
	if o == nil || IsNil(o.UseContext) {
		var ret bool
		return ret
	}
	return *o.UseContext
}

// GetUseContextOk returns a tuple with the UseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletionsBody) GetUseContextOk() (*bool, bool) {
	if o == nil || IsNil(o.UseContext) {
		return nil, false
	}
	return o.UseContext, true
}

// HasUseContext returns a boolean if a field has been set.
func (o *CompletionsBody) HasUseContext() bool {
	if o != nil && !IsNil(o.UseContext) {
		return true
	}

	return false
}

// SetUseContext gets a reference to the given bool and assigns it to the UseContext field.
func (o *CompletionsBody) SetUseContext(v bool) {
	o.UseContext = &v
}

// GetContextFilter returns the ContextFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompletionsBody) GetContextFilter() ContextFilter {
	if o == nil || IsNil(o.ContextFilter.Get()) {
		var ret ContextFilter
		return ret
	}
	return *o.ContextFilter.Get()
}

// GetContextFilterOk returns a tuple with the ContextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompletionsBody) GetContextFilterOk() (*ContextFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContextFilter.Get(), o.ContextFilter.IsSet()
}

// HasContextFilter returns a boolean if a field has been set.
func (o *CompletionsBody) HasContextFilter() bool {
	if o != nil && o.ContextFilter.IsSet() {
		return true
	}

	return false
}

// SetContextFilter gets a reference to the given NullableContextFilter and assigns it to the ContextFilter field.
func (o *CompletionsBody) SetContextFilter(v ContextFilter) {
	o.ContextFilter.Set(&v)
}

// SetContextFilterNil sets the value for ContextFilter to be an explicit nil
func (o *CompletionsBody) SetContextFilterNil() {
	o.ContextFilter.Set(nil)
}

// UnsetContextFilter ensures that no value is present for ContextFilter, not even an explicit nil
func (o *CompletionsBody) UnsetContextFilter() {
	o.ContextFilter.Unset()
}

// GetIncludeSources returns the IncludeSources field value if set, zero value otherwise.
func (o *CompletionsBody) GetIncludeSources() bool {
	if o == nil || IsNil(o.IncludeSources) {
		var ret bool
		return ret
	}
	return *o.IncludeSources
}

// GetIncludeSourcesOk returns a tuple with the IncludeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletionsBody) GetIncludeSourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSources) {
		return nil, false
	}
	return o.IncludeSources, true
}

// HasIncludeSources returns a boolean if a field has been set.
func (o *CompletionsBody) HasIncludeSources() bool {
	if o != nil && !IsNil(o.IncludeSources) {
		return true
	}

	return false
}

// SetIncludeSources gets a reference to the given bool and assigns it to the IncludeSources field.
func (o *CompletionsBody) SetIncludeSources(v bool) {
	o.IncludeSources = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *CompletionsBody) GetStream() bool {
	if o == nil || IsNil(o.Stream) {
		var ret bool
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletionsBody) GetStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *CompletionsBody) HasStream() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given bool and assigns it to the Stream field.
func (o *CompletionsBody) SetStream(v bool) {
	o.Stream = &v
}

func (o CompletionsBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompletionsBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prompt"] = o.Prompt
	if o.SystemPrompt.IsSet() {
		toSerialize["system_prompt"] = o.SystemPrompt.Get()
	}
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

func (o *CompletionsBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prompt",
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

	varCompletionsBody := _CompletionsBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompletionsBody)

	if err != nil {
		return err
	}

	*o = CompletionsBody(varCompletionsBody)

	return err
}

type NullableCompletionsBody struct {
	value *CompletionsBody
	isSet bool
}

func (v NullableCompletionsBody) Get() *CompletionsBody {
	return v.value
}

func (v *NullableCompletionsBody) Set(val *CompletionsBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletionsBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletionsBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletionsBody(val *CompletionsBody) *NullableCompletionsBody {
	return &NullableCompletionsBody{value: val, isSet: true}
}

func (v NullableCompletionsBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletionsBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
