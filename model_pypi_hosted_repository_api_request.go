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

// checks if the PypiHostedRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PypiHostedRepositoryApiRequest{}

// PypiHostedRepositoryApiRequest struct for PypiHostedRepositoryApiRequest
type PypiHostedRepositoryApiRequest struct {
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes `json:"component,omitempty"`
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
}

type _PypiHostedRepositoryApiRequest PypiHostedRepositoryApiRequest

// NewPypiHostedRepositoryApiRequest instantiates a new PypiHostedRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPypiHostedRepositoryApiRequest(name string, online bool, storage HostedStorageAttributes) *PypiHostedRepositoryApiRequest {
	this := PypiHostedRepositoryApiRequest{}
	this.Name = name
	this.Online = online
	this.Storage = storage
	return &this
}

// NewPypiHostedRepositoryApiRequestWithDefaults instantiates a new PypiHostedRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPypiHostedRepositoryApiRequestWithDefaults() *PypiHostedRepositoryApiRequest {
	this := PypiHostedRepositoryApiRequest{}
	return &this
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *PypiHostedRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PypiHostedRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *PypiHostedRepositoryApiRequest) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *PypiHostedRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *PypiHostedRepositoryApiRequest) GetComponent() ComponentAttributes {
	if o == nil || IsNil(o.Component) {
		var ret ComponentAttributes
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PypiHostedRepositoryApiRequest) GetComponentOk() (*ComponentAttributes, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *PypiHostedRepositoryApiRequest) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ComponentAttributes and assigns it to the Component field.
func (o *PypiHostedRepositoryApiRequest) SetComponent(v ComponentAttributes) {
	o.Component = &v
}

// GetName returns the Name field value
func (o *PypiHostedRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PypiHostedRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PypiHostedRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetOnline returns the Online field value
func (o *PypiHostedRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *PypiHostedRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *PypiHostedRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *PypiHostedRepositoryApiRequest) GetStorage() HostedStorageAttributes {
	if o == nil {
		var ret HostedStorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *PypiHostedRepositoryApiRequest) GetStorageOk() (*HostedStorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *PypiHostedRepositoryApiRequest) SetStorage(v HostedStorageAttributes) {
	o.Storage = v
}

func (o PypiHostedRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PypiHostedRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	toSerialize["name"] = o.Name
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	return toSerialize, nil
}

func (o *PypiHostedRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varPypiHostedRepositoryApiRequest := _PypiHostedRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPypiHostedRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = PypiHostedRepositoryApiRequest(varPypiHostedRepositoryApiRequest)

	return err
}

type NullablePypiHostedRepositoryApiRequest struct {
	value *PypiHostedRepositoryApiRequest
	isSet bool
}

func (v NullablePypiHostedRepositoryApiRequest) Get() *PypiHostedRepositoryApiRequest {
	return v.value
}

func (v *NullablePypiHostedRepositoryApiRequest) Set(val *PypiHostedRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePypiHostedRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePypiHostedRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePypiHostedRepositoryApiRequest(val *PypiHostedRepositoryApiRequest) *NullablePypiHostedRepositoryApiRequest {
	return &NullablePypiHostedRepositoryApiRequest{value: val, isSet: true}
}

func (v NullablePypiHostedRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePypiHostedRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


