/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.67.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PyPiProxyAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PyPiProxyAttributes{}

// PyPiProxyAttributes struct for PyPiProxyAttributes
type PyPiProxyAttributes struct {
	// Remove Quarantined Versions
	RemoveQuarantined bool `json:"removeQuarantined"`
}

type _PyPiProxyAttributes PyPiProxyAttributes

// NewPyPiProxyAttributes instantiates a new PyPiProxyAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPyPiProxyAttributes(removeQuarantined bool) *PyPiProxyAttributes {
	this := PyPiProxyAttributes{}
	this.RemoveQuarantined = removeQuarantined
	return &this
}

// NewPyPiProxyAttributesWithDefaults instantiates a new PyPiProxyAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPyPiProxyAttributesWithDefaults() *PyPiProxyAttributes {
	this := PyPiProxyAttributes{}
	return &this
}

// GetRemoveQuarantined returns the RemoveQuarantined field value
func (o *PyPiProxyAttributes) GetRemoveQuarantined() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RemoveQuarantined
}

// GetRemoveQuarantinedOk returns a tuple with the RemoveQuarantined field value
// and a boolean to check if the value has been set.
func (o *PyPiProxyAttributes) GetRemoveQuarantinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoveQuarantined, true
}

// SetRemoveQuarantined sets field value
func (o *PyPiProxyAttributes) SetRemoveQuarantined(v bool) {
	o.RemoveQuarantined = v
}

func (o PyPiProxyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PyPiProxyAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["removeQuarantined"] = o.RemoveQuarantined
	return toSerialize, nil
}

func (o *PyPiProxyAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"removeQuarantined",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPyPiProxyAttributes := _PyPiProxyAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPyPiProxyAttributes)

	if err != nil {
		return err
	}

	*o = PyPiProxyAttributes(varPyPiProxyAttributes)

	return err
}

type NullablePyPiProxyAttributes struct {
	value *PyPiProxyAttributes
	isSet bool
}

func (v NullablePyPiProxyAttributes) Get() *PyPiProxyAttributes {
	return v.value
}

func (v *NullablePyPiProxyAttributes) Set(val *PyPiProxyAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePyPiProxyAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePyPiProxyAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePyPiProxyAttributes(val *PyPiProxyAttributes) *NullablePyPiProxyAttributes {
	return &NullablePyPiProxyAttributes{value: val, isSet: true}
}

func (v NullablePyPiProxyAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePyPiProxyAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


