/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.68.1-02.

API version: 3.68.1-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GitLfsHostedRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitLfsHostedRepositoryApiRequest{}

// GitLfsHostedRepositoryApiRequest struct for GitLfsHostedRepositoryApiRequest
type GitLfsHostedRepositoryApiRequest struct {
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes `json:"component,omitempty"`
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
}

type _GitLfsHostedRepositoryApiRequest GitLfsHostedRepositoryApiRequest

// NewGitLfsHostedRepositoryApiRequest instantiates a new GitLfsHostedRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitLfsHostedRepositoryApiRequest(name string, online bool, storage HostedStorageAttributes) *GitLfsHostedRepositoryApiRequest {
	this := GitLfsHostedRepositoryApiRequest{}
	this.Name = name
	this.Online = online
	this.Storage = storage
	return &this
}

// NewGitLfsHostedRepositoryApiRequestWithDefaults instantiates a new GitLfsHostedRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitLfsHostedRepositoryApiRequestWithDefaults() *GitLfsHostedRepositoryApiRequest {
	this := GitLfsHostedRepositoryApiRequest{}
	return &this
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *GitLfsHostedRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitLfsHostedRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *GitLfsHostedRepositoryApiRequest) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *GitLfsHostedRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *GitLfsHostedRepositoryApiRequest) GetComponent() ComponentAttributes {
	if o == nil || IsNil(o.Component) {
		var ret ComponentAttributes
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitLfsHostedRepositoryApiRequest) GetComponentOk() (*ComponentAttributes, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *GitLfsHostedRepositoryApiRequest) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ComponentAttributes and assigns it to the Component field.
func (o *GitLfsHostedRepositoryApiRequest) SetComponent(v ComponentAttributes) {
	o.Component = &v
}

// GetName returns the Name field value
func (o *GitLfsHostedRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitLfsHostedRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitLfsHostedRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetOnline returns the Online field value
func (o *GitLfsHostedRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *GitLfsHostedRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *GitLfsHostedRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *GitLfsHostedRepositoryApiRequest) GetStorage() HostedStorageAttributes {
	if o == nil {
		var ret HostedStorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *GitLfsHostedRepositoryApiRequest) GetStorageOk() (*HostedStorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *GitLfsHostedRepositoryApiRequest) SetStorage(v HostedStorageAttributes) {
	o.Storage = v
}

func (o GitLfsHostedRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitLfsHostedRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
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

func (o *GitLfsHostedRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGitLfsHostedRepositoryApiRequest := _GitLfsHostedRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGitLfsHostedRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = GitLfsHostedRepositoryApiRequest(varGitLfsHostedRepositoryApiRequest)

	return err
}

type NullableGitLfsHostedRepositoryApiRequest struct {
	value *GitLfsHostedRepositoryApiRequest
	isSet bool
}

func (v NullableGitLfsHostedRepositoryApiRequest) Get() *GitLfsHostedRepositoryApiRequest {
	return v.value
}

func (v *NullableGitLfsHostedRepositoryApiRequest) Set(val *GitLfsHostedRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGitLfsHostedRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGitLfsHostedRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitLfsHostedRepositoryApiRequest(val *GitLfsHostedRepositoryApiRequest) *NullableGitLfsHostedRepositoryApiRequest {
	return &NullableGitLfsHostedRepositoryApiRequest{value: val, isSet: true}
}

func (v NullableGitLfsHostedRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitLfsHostedRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


