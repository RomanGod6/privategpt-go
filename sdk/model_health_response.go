/*
FastAPI- Private GPT

SDK for GO LANG for Private GPT

API version: 0.1.0
*/

package openapi

import (
	"encoding/json"
)

// checks if the HealthResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthResponse{}

// HealthResponse struct for HealthResponse
type HealthResponse struct {
	Status interface{} `json:"status,omitempty"`
}

// NewHealthResponse instantiates a new HealthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthResponse() *HealthResponse {
	this := HealthResponse{}
	return &this
}

// NewHealthResponseWithDefaults instantiates a new HealthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthResponseWithDefaults() *HealthResponse {
	this := HealthResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthResponse) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthResponse) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *HealthResponse) SetStatus(v interface{}) {
	o.Status = v
}

func (o HealthResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableHealthResponse struct {
	value *HealthResponse
	isSet bool
}

func (v NullableHealthResponse) Get() *HealthResponse {
	return v.value
}

func (v *NullableHealthResponse) Set(val *HealthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthResponse(val *HealthResponse) *NullableHealthResponse {
	return &NullableHealthResponse{value: val, isSet: true}
}

func (v NullableHealthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
