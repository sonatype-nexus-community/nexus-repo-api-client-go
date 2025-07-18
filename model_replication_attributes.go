/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.82.0-08.

API version: 3.82.0-08
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ReplicationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicationAttributes{}

// ReplicationAttributes struct for ReplicationAttributes
type ReplicationAttributes struct {
	// Regular Expression of Asset Paths to pull pre-emptively pull
	AssetPathRegex *string `json:"assetPathRegex,omitempty"`
	// Whether pre-emptive pull is enabled
	PreemptivePullEnabled bool `json:"preemptivePullEnabled"`
}

type _ReplicationAttributes ReplicationAttributes

// NewReplicationAttributes instantiates a new ReplicationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationAttributes(preemptivePullEnabled bool) *ReplicationAttributes {
	this := ReplicationAttributes{}
	this.PreemptivePullEnabled = preemptivePullEnabled
	return &this
}

// NewReplicationAttributesWithDefaults instantiates a new ReplicationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationAttributesWithDefaults() *ReplicationAttributes {
	this := ReplicationAttributes{}
	return &this
}

// GetAssetPathRegex returns the AssetPathRegex field value if set, zero value otherwise.
func (o *ReplicationAttributes) GetAssetPathRegex() string {
	if o == nil || IsNil(o.AssetPathRegex) {
		var ret string
		return ret
	}
	return *o.AssetPathRegex
}

// GetAssetPathRegexOk returns a tuple with the AssetPathRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAttributes) GetAssetPathRegexOk() (*string, bool) {
	if o == nil || IsNil(o.AssetPathRegex) {
		return nil, false
	}
	return o.AssetPathRegex, true
}

// HasAssetPathRegex returns a boolean if a field has been set.
func (o *ReplicationAttributes) HasAssetPathRegex() bool {
	if o != nil && !IsNil(o.AssetPathRegex) {
		return true
	}

	return false
}

// SetAssetPathRegex gets a reference to the given string and assigns it to the AssetPathRegex field.
func (o *ReplicationAttributes) SetAssetPathRegex(v string) {
	o.AssetPathRegex = &v
}

// GetPreemptivePullEnabled returns the PreemptivePullEnabled field value
func (o *ReplicationAttributes) GetPreemptivePullEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PreemptivePullEnabled
}

// GetPreemptivePullEnabledOk returns a tuple with the PreemptivePullEnabled field value
// and a boolean to check if the value has been set.
func (o *ReplicationAttributes) GetPreemptivePullEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreemptivePullEnabled, true
}

// SetPreemptivePullEnabled sets field value
func (o *ReplicationAttributes) SetPreemptivePullEnabled(v bool) {
	o.PreemptivePullEnabled = v
}

func (o ReplicationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetPathRegex) {
		toSerialize["assetPathRegex"] = o.AssetPathRegex
	}
	toSerialize["preemptivePullEnabled"] = o.PreemptivePullEnabled
	return toSerialize, nil
}

func (o *ReplicationAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"preemptivePullEnabled",
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

	varReplicationAttributes := _ReplicationAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplicationAttributes)

	if err != nil {
		return err
	}

	*o = ReplicationAttributes(varReplicationAttributes)

	return err
}

type NullableReplicationAttributes struct {
	value *ReplicationAttributes
	isSet bool
}

func (v NullableReplicationAttributes) Get() *ReplicationAttributes {
	return v.value
}

func (v *NullableReplicationAttributes) Set(val *ReplicationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationAttributes(val *ReplicationAttributes) *NullableReplicationAttributes {
	return &NullableReplicationAttributes{value: val, isSet: true}
}

func (v NullableReplicationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


