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

// checks if the IngestTextBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestTextBody{}

// IngestTextBody struct for IngestTextBody
type IngestTextBody struct {
	FileName string `json:"file_name"`
	Text     string `json:"text"`
}

type _IngestTextBody IngestTextBody

// NewIngestTextBody instantiates a new IngestTextBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestTextBody(fileName string, text string) *IngestTextBody {
	this := IngestTextBody{}
	this.FileName = fileName
	this.Text = text
	return &this
}

// NewIngestTextBodyWithDefaults instantiates a new IngestTextBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestTextBodyWithDefaults() *IngestTextBody {
	this := IngestTextBody{}
	return &this
}

// GetFileName returns the FileName field value
func (o *IngestTextBody) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *IngestTextBody) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *IngestTextBody) SetFileName(v string) {
	o.FileName = v
}

// GetText returns the Text field value
func (o *IngestTextBody) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *IngestTextBody) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *IngestTextBody) SetText(v string) {
	o.Text = v
}

func (o IngestTextBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestTextBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_name"] = o.FileName
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

func (o *IngestTextBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file_name",
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

	varIngestTextBody := _IngestTextBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestTextBody)

	if err != nil {
		return err
	}

	*o = IngestTextBody(varIngestTextBody)

	return err
}

type NullableIngestTextBody struct {
	value *IngestTextBody
	isSet bool
}

func (v NullableIngestTextBody) Get() *IngestTextBody {
	return v.value
}

func (v *NullableIngestTextBody) Set(val *IngestTextBody) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestTextBody) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestTextBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestTextBody(val *IngestTextBody) *NullableIngestTextBody {
	return &NullableIngestTextBody{value: val, isSet: true}
}

func (v NullableIngestTextBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestTextBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
