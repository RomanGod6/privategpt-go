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

// checks if the ContextFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextFilter{}

// ContextFilter struct for ContextFilter
type ContextFilter struct {
	DocsIds []string `json:"docs_ids"`
}

type _ContextFilter ContextFilter

// NewContextFilter instantiates a new ContextFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextFilter(docsIds []string) *ContextFilter {
	this := ContextFilter{}
	this.DocsIds = docsIds
	return &this
}

// NewContextFilterWithDefaults instantiates a new ContextFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextFilterWithDefaults() *ContextFilter {
	this := ContextFilter{}
	return &this
}

// GetDocsIds returns the DocsIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ContextFilter) GetDocsIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DocsIds
}

// GetDocsIdsOk returns a tuple with the DocsIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContextFilter) GetDocsIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DocsIds) {
		return nil, false
	}
	return o.DocsIds, true
}

// SetDocsIds sets field value
func (o *ContextFilter) SetDocsIds(v []string) {
	o.DocsIds = v
}

func (o ContextFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DocsIds != nil {
		toSerialize["docs_ids"] = o.DocsIds
	}
	return toSerialize, nil
}

func (o *ContextFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"docs_ids",
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

	varContextFilter := _ContextFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContextFilter)

	if err != nil {
		return err
	}

	*o = ContextFilter(varContextFilter)

	return err
}

type NullableContextFilter struct {
	value *ContextFilter
	isSet bool
}

func (v NullableContextFilter) Get() *ContextFilter {
	return v.value
}

func (v *NullableContextFilter) Set(val *ContextFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableContextFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableContextFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextFilter(val *ContextFilter) *NullableContextFilter {
	return &NullableContextFilter{value: val, isSet: true}
}

func (v NullableContextFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
