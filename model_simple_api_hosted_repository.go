/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.67.1-01.

API version: 3.67.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SimpleApiHostedRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleApiHostedRepository{}

// SimpleApiHostedRepository struct for SimpleApiHostedRepository
type SimpleApiHostedRepository struct {
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes `json:"component,omitempty"`
	// A unique identifier for this repository
	Name *string `json:"name,omitempty"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
}

type _SimpleApiHostedRepository SimpleApiHostedRepository

// NewSimpleApiHostedRepository instantiates a new SimpleApiHostedRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleApiHostedRepository(online bool, storage HostedStorageAttributes) *SimpleApiHostedRepository {
	this := SimpleApiHostedRepository{}
	this.Online = online
	this.Storage = storage
	return &this
}

// NewSimpleApiHostedRepositoryWithDefaults instantiates a new SimpleApiHostedRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleApiHostedRepositoryWithDefaults() *SimpleApiHostedRepository {
	this := SimpleApiHostedRepository{}
	return &this
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *SimpleApiHostedRepository) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleApiHostedRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *SimpleApiHostedRepository) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *SimpleApiHostedRepository) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *SimpleApiHostedRepository) GetComponent() ComponentAttributes {
	if o == nil || IsNil(o.Component) {
		var ret ComponentAttributes
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleApiHostedRepository) GetComponentOk() (*ComponentAttributes, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *SimpleApiHostedRepository) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ComponentAttributes and assigns it to the Component field.
func (o *SimpleApiHostedRepository) SetComponent(v ComponentAttributes) {
	o.Component = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimpleApiHostedRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleApiHostedRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimpleApiHostedRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimpleApiHostedRepository) SetName(v string) {
	o.Name = &v
}

// GetOnline returns the Online field value
func (o *SimpleApiHostedRepository) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *SimpleApiHostedRepository) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *SimpleApiHostedRepository) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *SimpleApiHostedRepository) GetStorage() HostedStorageAttributes {
	if o == nil {
		var ret HostedStorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *SimpleApiHostedRepository) GetStorageOk() (*HostedStorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *SimpleApiHostedRepository) SetStorage(v HostedStorageAttributes) {
	o.Storage = v
}

func (o SimpleApiHostedRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleApiHostedRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	return toSerialize, nil
}

func (o *SimpleApiHostedRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varSimpleApiHostedRepository := _SimpleApiHostedRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimpleApiHostedRepository)

	if err != nil {
		return err
	}

	*o = SimpleApiHostedRepository(varSimpleApiHostedRepository)

	return err
}

type NullableSimpleApiHostedRepository struct {
	value *SimpleApiHostedRepository
	isSet bool
}

func (v NullableSimpleApiHostedRepository) Get() *SimpleApiHostedRepository {
	return v.value
}

func (v *NullableSimpleApiHostedRepository) Set(val *SimpleApiHostedRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleApiHostedRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleApiHostedRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleApiHostedRepository(val *SimpleApiHostedRepository) *NullableSimpleApiHostedRepository {
	return &NullableSimpleApiHostedRepository{value: val, isSet: true}
}

func (v NullableSimpleApiHostedRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleApiHostedRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


