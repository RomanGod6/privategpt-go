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

// checks if the IngestedDoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestedDoc{}

// IngestedDoc struct for IngestedDoc
type IngestedDoc struct {
	Object      interface{}            `json:"object"`
	DocId       string                 `json:"doc_id"`
	DocMetadata map[string]interface{} `json:"doc_metadata"`
}

type _IngestedDoc IngestedDoc

// NewIngestedDoc instantiates a new IngestedDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestedDoc(object interface{}, docId string, docMetadata map[string]interface{}) *IngestedDoc {
	this := IngestedDoc{}
	this.Object = object
	this.DocId = docId
	this.DocMetadata = docMetadata
	return &this
}

// NewIngestedDocWithDefaults instantiates a new IngestedDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestedDocWithDefaults() *IngestedDoc {
	this := IngestedDoc{}
	return &this
}

// GetObject returns the Object field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *IngestedDoc) GetObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestedDoc) GetObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *IngestedDoc) SetObject(v interface{}) {
	o.Object = v
}

// GetDocId returns the DocId field value
func (o *IngestedDoc) GetDocId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value
// and a boolean to check if the value has been set.
func (o *IngestedDoc) GetDocIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocId, true
}

// SetDocId sets field value
func (o *IngestedDoc) SetDocId(v string) {
	o.DocId = v
}

// GetDocMetadata returns the DocMetadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *IngestedDoc) GetDocMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.DocMetadata
}

// GetDocMetadataOk returns a tuple with the DocMetadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestedDoc) GetDocMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DocMetadata) {
		return map[string]interface{}{}, false
	}
	return o.DocMetadata, true
}

// SetDocMetadata sets field value
func (o *IngestedDoc) SetDocMetadata(v map[string]interface{}) {
	o.DocMetadata = v
}

func (o IngestedDoc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestedDoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	toSerialize["doc_id"] = o.DocId
	if o.DocMetadata != nil {
		toSerialize["doc_metadata"] = o.DocMetadata
	}
	return toSerialize, nil
}

func (o *IngestedDoc) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"doc_id",
		"doc_metadata",
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

	varIngestedDoc := _IngestedDoc{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestedDoc)

	if err != nil {
		return err
	}

	*o = IngestedDoc(varIngestedDoc)

	return err
}

type NullableIngestedDoc struct {
	value *IngestedDoc
	isSet bool
}

func (v NullableIngestedDoc) Get() *IngestedDoc {
	return v.value
}

func (v *NullableIngestedDoc) Set(val *IngestedDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestedDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestedDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestedDoc(val *IngestedDoc) *NullableIngestedDoc {
	return &NullableIngestedDoc{value: val, isSet: true}
}

func (v NullableIngestedDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestedDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
