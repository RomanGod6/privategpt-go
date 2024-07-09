/*
FastAPI- Private GPT

SDK for GO LANG for Private GPT

API version: 0.1.0
*/

package sdk

import (
	"encoding/json"
	"fmt"
)

// ValidationErrorLocInner struct for ValidationErrorLocInner
type ValidationErrorLocInner struct {
	Int32  *int32
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ValidationErrorLocInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Int32
	err = json.Unmarshal(data, &dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ValidationErrorLocInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ValidationErrorLocInner) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableValidationErrorLocInner struct {
	value *ValidationErrorLocInner
	isSet bool
}

func (v NullableValidationErrorLocInner) Get() *ValidationErrorLocInner {
	return v.value
}

func (v *NullableValidationErrorLocInner) Set(val *ValidationErrorLocInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationErrorLocInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationErrorLocInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationErrorLocInner(val *ValidationErrorLocInner) *NullableValidationErrorLocInner {
	return &NullableValidationErrorLocInner{value: val, isSet: true}
}

func (v NullableValidationErrorLocInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationErrorLocInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
