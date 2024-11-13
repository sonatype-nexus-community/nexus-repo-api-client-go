/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.74.0-05.

API version: 3.74.0-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PypiGroupRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PypiGroupRepositoryApiRequest{}

// PypiGroupRepositoryApiRequest struct for PypiGroupRepositoryApiRequest
type PypiGroupRepositoryApiRequest struct {
	Group GroupAttributes `json:"group"`
	// A unique identifier for this repository
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage StorageAttributes `json:"storage"`
}

type _PypiGroupRepositoryApiRequest PypiGroupRepositoryApiRequest

// NewPypiGroupRepositoryApiRequest instantiates a new PypiGroupRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPypiGroupRepositoryApiRequest(group GroupAttributes, name string, online bool, storage StorageAttributes) *PypiGroupRepositoryApiRequest {
	this := PypiGroupRepositoryApiRequest{}
	this.Group = group
	this.Name = name
	this.Online = online
	this.Storage = storage
	return &this
}

// NewPypiGroupRepositoryApiRequestWithDefaults instantiates a new PypiGroupRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPypiGroupRepositoryApiRequestWithDefaults() *PypiGroupRepositoryApiRequest {
	this := PypiGroupRepositoryApiRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *PypiGroupRepositoryApiRequest) GetGroup() GroupAttributes {
	if o == nil {
		var ret GroupAttributes
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *PypiGroupRepositoryApiRequest) GetGroupOk() (*GroupAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *PypiGroupRepositoryApiRequest) SetGroup(v GroupAttributes) {
	o.Group = v
}

// GetName returns the Name field value
func (o *PypiGroupRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PypiGroupRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PypiGroupRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetOnline returns the Online field value
func (o *PypiGroupRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *PypiGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *PypiGroupRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *PypiGroupRepositoryApiRequest) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *PypiGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *PypiGroupRepositoryApiRequest) SetStorage(v StorageAttributes) {
	o.Storage = v
}

func (o PypiGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PypiGroupRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["name"] = o.Name
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	return toSerialize, nil
}

func (o *PypiGroupRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group",
		"name",
		"online",
		"storage",
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

	varPypiGroupRepositoryApiRequest := _PypiGroupRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPypiGroupRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = PypiGroupRepositoryApiRequest(varPypiGroupRepositoryApiRequest)

	return err
}

type NullablePypiGroupRepositoryApiRequest struct {
	value *PypiGroupRepositoryApiRequest
	isSet bool
}

func (v NullablePypiGroupRepositoryApiRequest) Get() *PypiGroupRepositoryApiRequest {
	return v.value
}

func (v *NullablePypiGroupRepositoryApiRequest) Set(val *PypiGroupRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePypiGroupRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePypiGroupRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePypiGroupRepositoryApiRequest(val *PypiGroupRepositoryApiRequest) *NullablePypiGroupRepositoryApiRequest {
	return &NullablePypiGroupRepositoryApiRequest{value: val, isSet: true}
}

func (v NullablePypiGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePypiGroupRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


