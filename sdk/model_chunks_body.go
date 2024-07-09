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

// checks if the ChunksBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChunksBody{}

// ChunksBody struct for ChunksBody
type ChunksBody struct {
	Text           string                `json:"text"`
	ContextFilter  NullableContextFilter `json:"context_filter,omitempty"`
	Limit          *int32                `json:"limit,omitempty"`
	PrevNextChunks *int32                `json:"prev_next_chunks,omitempty"`
}

type _ChunksBody ChunksBody

// NewChunksBody instantiates a new ChunksBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChunksBody(text string) *ChunksBody {
	this := ChunksBody{}
	this.Text = text
	var limit int32 = 10
	this.Limit = &limit
	var prevNextChunks int32 = 0
	this.PrevNextChunks = &prevNextChunks
	return &this
}

// NewChunksBodyWithDefaults instantiates a new ChunksBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChunksBodyWithDefaults() *ChunksBody {
	this := ChunksBody{}
	var limit int32 = 10
	this.Limit = &limit
	var prevNextChunks int32 = 0
	this.PrevNextChunks = &prevNextChunks
	return &this
}

// GetText returns the Text field value
func (o *ChunksBody) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *ChunksBody) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *ChunksBody) SetText(v string) {
	o.Text = v
}

// GetContextFilter returns the ContextFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChunksBody) GetContextFilter() ContextFilter {
	if o == nil || IsNil(o.ContextFilter.Get()) {
		var ret ContextFilter
		return ret
	}
	return *o.ContextFilter.Get()
}

// GetContextFilterOk returns a tuple with the ContextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChunksBody) GetContextFilterOk() (*ContextFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContextFilter.Get(), o.ContextFilter.IsSet()
}

// HasContextFilter returns a boolean if a field has been set.
func (o *ChunksBody) HasContextFilter() bool {
	if o != nil && o.ContextFilter.IsSet() {
		return true
	}

	return false
}

// SetContextFilter gets a reference to the given NullableContextFilter and assigns it to the ContextFilter field.
func (o *ChunksBody) SetContextFilter(v ContextFilter) {
	o.ContextFilter.Set(&v)
}

// SetContextFilterNil sets the value for ContextFilter to be an explicit nil
func (o *ChunksBody) SetContextFilterNil() {
	o.ContextFilter.Set(nil)
}

// UnsetContextFilter ensures that no value is present for ContextFilter, not even an explicit nil
func (o *ChunksBody) UnsetContextFilter() {
	o.ContextFilter.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ChunksBody) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunksBody) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ChunksBody) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ChunksBody) SetLimit(v int32) {
	o.Limit = &v
}

// GetPrevNextChunks returns the PrevNextChunks field value if set, zero value otherwise.
func (o *ChunksBody) GetPrevNextChunks() int32 {
	if o == nil || IsNil(o.PrevNextChunks) {
		var ret int32
		return ret
	}
	return *o.PrevNextChunks
}

// GetPrevNextChunksOk returns a tuple with the PrevNextChunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunksBody) GetPrevNextChunksOk() (*int32, bool) {
	if o == nil || IsNil(o.PrevNextChunks) {
		return nil, false
	}
	return o.PrevNextChunks, true
}

// HasPrevNextChunks returns a boolean if a field has been set.
func (o *ChunksBody) HasPrevNextChunks() bool {
	if o != nil && !IsNil(o.PrevNextChunks) {
		return true
	}

	return false
}

// SetPrevNextChunks gets a reference to the given int32 and assigns it to the PrevNextChunks field.
func (o *ChunksBody) SetPrevNextChunks(v int32) {
	o.PrevNextChunks = &v
}

func (o ChunksBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChunksBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if o.ContextFilter.IsSet() {
		toSerialize["context_filter"] = o.ContextFilter.Get()
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.PrevNextChunks) {
		toSerialize["prev_next_chunks"] = o.PrevNextChunks
	}
	return toSerialize, nil
}

func (o *ChunksBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
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

	varChunksBody := _ChunksBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChunksBody)

	if err != nil {
		return err
	}

	*o = ChunksBody(varChunksBody)

	return err
}

type NullableChunksBody struct {
	value *ChunksBody
	isSet bool
}

func (v NullableChunksBody) Get() *ChunksBody {
	return v.value
}

func (v *NullableChunksBody) Set(val *ChunksBody) {
	v.value = val
	v.isSet = true
}

func (v NullableChunksBody) IsSet() bool {
	return v.isSet
}

func (v *NullableChunksBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChunksBody(val *ChunksBody) *NullableChunksBody {
	return &NullableChunksBody{value: val, isSet: true}
}

func (v NullableChunksBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChunksBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
