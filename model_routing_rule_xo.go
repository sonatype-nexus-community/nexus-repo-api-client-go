/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.67.1-01.

API version: 3.67.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
)

// checks if the RoutingRuleXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingRuleXO{}

// RoutingRuleXO struct for RoutingRuleXO
type RoutingRuleXO struct {
	Description *string `json:"description,omitempty"`
	// Regular expressions used to identify request paths that are allowed or blocked (depending on mode)
	Matchers []string `json:"matchers,omitempty"`
	// Determines what should be done with requests when their path matches any of the matchers
	Mode *string `json:"mode,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewRoutingRuleXO instantiates a new RoutingRuleXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingRuleXO() *RoutingRuleXO {
	this := RoutingRuleXO{}
	return &this
}

// NewRoutingRuleXOWithDefaults instantiates a new RoutingRuleXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingRuleXOWithDefaults() *RoutingRuleXO {
	this := RoutingRuleXO{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoutingRuleXO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleXO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoutingRuleXO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoutingRuleXO) SetDescription(v string) {
	o.Description = &v
}

// GetMatchers returns the Matchers field value if set, zero value otherwise.
func (o *RoutingRuleXO) GetMatchers() []string {
	if o == nil || IsNil(o.Matchers) {
		var ret []string
		return ret
	}
	return o.Matchers
}

// GetMatchersOk returns a tuple with the Matchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleXO) GetMatchersOk() ([]string, bool) {
	if o == nil || IsNil(o.Matchers) {
		return nil, false
	}
	return o.Matchers, true
}

// HasMatchers returns a boolean if a field has been set.
func (o *RoutingRuleXO) HasMatchers() bool {
	if o != nil && !IsNil(o.Matchers) {
		return true
	}

	return false
}

// SetMatchers gets a reference to the given []string and assigns it to the Matchers field.
func (o *RoutingRuleXO) SetMatchers(v []string) {
	o.Matchers = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RoutingRuleXO) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleXO) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RoutingRuleXO) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RoutingRuleXO) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoutingRuleXO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleXO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoutingRuleXO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoutingRuleXO) SetName(v string) {
	o.Name = &v
}

func (o RoutingRuleXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingRuleXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Matchers) {
		toSerialize["matchers"] = o.Matchers
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableRoutingRuleXO struct {
	value *RoutingRuleXO
	isSet bool
}

func (v NullableRoutingRuleXO) Get() *RoutingRuleXO {
	return v.value
}

func (v *NullableRoutingRuleXO) Set(val *RoutingRuleXO) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingRuleXO) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingRuleXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingRuleXO(val *RoutingRuleXO) *NullableRoutingRuleXO {
	return &NullableRoutingRuleXO{value: val, isSet: true}
}

func (v NullableRoutingRuleXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingRuleXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


