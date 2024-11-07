/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.68.1-02.

API version: 3.68.1-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
)

// checks if the ScriptXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScriptXO{}

// ScriptXO struct for ScriptXO
type ScriptXO struct {
	Content *string `json:"content,omitempty"`
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	Type *string `json:"type,omitempty"`
}

// NewScriptXO instantiates a new ScriptXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptXO() *ScriptXO {
	this := ScriptXO{}
	return &this
}

// NewScriptXOWithDefaults instantiates a new ScriptXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptXOWithDefaults() *ScriptXO {
	this := ScriptXO{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ScriptXO) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptXO) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ScriptXO) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ScriptXO) SetContent(v string) {
	o.Content = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScriptXO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptXO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScriptXO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScriptXO) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScriptXO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptXO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScriptXO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScriptXO) SetType(v string) {
	o.Type = &v
}

func (o ScriptXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScriptXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableScriptXO struct {
	value *ScriptXO
	isSet bool
}

func (v NullableScriptXO) Get() *ScriptXO {
	return v.value
}

func (v *NullableScriptXO) Set(val *ScriptXO) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptXO) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptXO(val *ScriptXO) *NullableScriptXO {
	return &NullableScriptXO{value: val, isSet: true}
}

func (v NullableScriptXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


