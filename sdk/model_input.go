/*
FastAPI- Private GPT

SDK for GO LANG for Private GPT

API version: 0.1.0
*/

package openapi

import (
	"encoding/json"
	"fmt"
)

// Input struct for Input
type Input struct {
	ArrayOfString *[]string
	String        *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Input) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ArrayOfString
	err = json.Unmarshal(data, &dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			return nil // data stored in dst.ArrayOfString, return on the first match
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Input)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Input) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableInput struct {
	value *Input
	isSet bool
}

func (v NullableInput) Get() *Input {
	return v.value
}

func (v *NullableInput) Set(val *Input) {
	v.value = val
	v.isSet = true
}

func (v NullableInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInput(val *Input) *NullableInput {
	return &NullableInput{value: val, isSet: true}
}

func (v NullableInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
