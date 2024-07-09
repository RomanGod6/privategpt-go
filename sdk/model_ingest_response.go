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

// checks if the IngestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestResponse{}

// IngestResponse struct for IngestResponse
type IngestResponse struct {
	Object interface{}   `json:"object"`
	Model  interface{}   `json:"model"`
	Data   []IngestedDoc `json:"data"`
}

type _IngestResponse IngestResponse

// NewIngestResponse instantiates a new IngestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestResponse(object interface{}, model interface{}, data []IngestedDoc) *IngestResponse {
	this := IngestResponse{}
	this.Object = object
	this.Model = model
	this.Data = data
	return &this
}

// NewIngestResponseWithDefaults instantiates a new IngestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestResponseWithDefaults() *IngestResponse {
	this := IngestResponse{}
	return &this
}

// GetObject returns the Object field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *IngestResponse) GetObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestResponse) GetObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *IngestResponse) SetObject(v interface{}) {
	o.Object = v
}

// GetModel returns the Model field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *IngestResponse) GetModel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestResponse) GetModelOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *IngestResponse) SetModel(v interface{}) {
	o.Model = v
}

// GetData returns the Data field value
func (o *IngestResponse) GetData() []IngestedDoc {
	if o == nil {
		var ret []IngestedDoc
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IngestResponse) GetDataOk() ([]IngestedDoc, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *IngestResponse) SetData(v []IngestedDoc) {
	o.Data = v
}

func (o IngestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *IngestResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"model",
		"data",
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

	varIngestResponse := _IngestResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestResponse)

	if err != nil {
		return err
	}

	*o = IngestResponse(varIngestResponse)

	return err
}

type NullableIngestResponse struct {
	value *IngestResponse
	isSet bool
}

func (v NullableIngestResponse) Get() *IngestResponse {
	return v.value
}

func (v *NullableIngestResponse) Set(val *IngestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestResponse(val *IngestResponse) *NullableIngestResponse {
	return &NullableIngestResponse{value: val, isSet: true}
}

func (v NullableIngestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
