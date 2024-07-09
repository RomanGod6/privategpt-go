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

// checks if the Chunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Chunk{}

// Chunk struct for Chunk
type Chunk struct {
	Object        interface{} `json:"object"`
	Score         float32     `json:"score"`
	Document      IngestedDoc `json:"document"`
	Text          string      `json:"text"`
	PreviousTexts []string    `json:"previous_texts,omitempty"`
	NextTexts     []string    `json:"next_texts,omitempty"`
}

type _Chunk Chunk

// NewChunk instantiates a new Chunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChunk(object interface{}, score float32, document IngestedDoc, text string) *Chunk {
	this := Chunk{}
	this.Object = object
	this.Score = score
	this.Document = document
	this.Text = text
	return &this
}

// NewChunkWithDefaults instantiates a new Chunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChunkWithDefaults() *Chunk {
	this := Chunk{}
	return &this
}

// GetObject returns the Object field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Chunk) GetObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chunk) GetObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Chunk) SetObject(v interface{}) {
	o.Object = v
}

// GetScore returns the Score field value
func (o *Chunk) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *Chunk) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *Chunk) SetScore(v float32) {
	o.Score = v
}

// GetDocument returns the Document field value
func (o *Chunk) GetDocument() IngestedDoc {
	if o == nil {
		var ret IngestedDoc
		return ret
	}

	return o.Document
}

// GetDocumentOk returns a tuple with the Document field value
// and a boolean to check if the value has been set.
func (o *Chunk) GetDocumentOk() (*IngestedDoc, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Document, true
}

// SetDocument sets field value
func (o *Chunk) SetDocument(v IngestedDoc) {
	o.Document = v
}

// GetText returns the Text field value
func (o *Chunk) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *Chunk) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *Chunk) SetText(v string) {
	o.Text = v
}

// GetPreviousTexts returns the PreviousTexts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chunk) GetPreviousTexts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PreviousTexts
}

// GetPreviousTextsOk returns a tuple with the PreviousTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chunk) GetPreviousTextsOk() ([]string, bool) {
	if o == nil || IsNil(o.PreviousTexts) {
		return nil, false
	}
	return o.PreviousTexts, true
}

// HasPreviousTexts returns a boolean if a field has been set.
func (o *Chunk) HasPreviousTexts() bool {
	if o != nil && !IsNil(o.PreviousTexts) {
		return true
	}

	return false
}

// SetPreviousTexts gets a reference to the given []string and assigns it to the PreviousTexts field.
func (o *Chunk) SetPreviousTexts(v []string) {
	o.PreviousTexts = v
}

// GetNextTexts returns the NextTexts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chunk) GetNextTexts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NextTexts
}

// GetNextTextsOk returns a tuple with the NextTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chunk) GetNextTextsOk() ([]string, bool) {
	if o == nil || IsNil(o.NextTexts) {
		return nil, false
	}
	return o.NextTexts, true
}

// HasNextTexts returns a boolean if a field has been set.
func (o *Chunk) HasNextTexts() bool {
	if o != nil && !IsNil(o.NextTexts) {
		return true
	}

	return false
}

// SetNextTexts gets a reference to the given []string and assigns it to the NextTexts field.
func (o *Chunk) SetNextTexts(v []string) {
	o.NextTexts = v
}

func (o Chunk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Chunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	toSerialize["score"] = o.Score
	toSerialize["document"] = o.Document
	toSerialize["text"] = o.Text
	if o.PreviousTexts != nil {
		toSerialize["previous_texts"] = o.PreviousTexts
	}
	if o.NextTexts != nil {
		toSerialize["next_texts"] = o.NextTexts
	}
	return toSerialize, nil
}

func (o *Chunk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"score",
		"document",
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

	varChunk := _Chunk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChunk)

	if err != nil {
		return err
	}

	*o = Chunk(varChunk)

	return err
}

type NullableChunk struct {
	value *Chunk
	isSet bool
}

func (v NullableChunk) Get() *Chunk {
	return v.value
}

func (v *NullableChunk) Set(val *Chunk) {
	v.value = val
	v.isSet = true
}

func (v NullableChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChunk(val *Chunk) *NullableChunk {
	return &NullableChunk{value: val, isSet: true}
}

func (v NullableChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
