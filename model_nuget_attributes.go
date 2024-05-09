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

// checks if the NugetAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NugetAttributes{}

// NugetAttributes struct for NugetAttributes
type NugetAttributes struct {
	// Nuget protocol version
	NugetVersion *string `json:"nugetVersion,omitempty"`
	// How long to cache query results from the proxied repository (in seconds)
	QueryCacheItemMaxAge *int32 `json:"queryCacheItemMaxAge,omitempty"`
}

// NewNugetAttributes instantiates a new NugetAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNugetAttributes() *NugetAttributes {
	this := NugetAttributes{}
	return &this
}

// NewNugetAttributesWithDefaults instantiates a new NugetAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNugetAttributesWithDefaults() *NugetAttributes {
	this := NugetAttributes{}
	return &this
}

// GetNugetVersion returns the NugetVersion field value if set, zero value otherwise.
func (o *NugetAttributes) GetNugetVersion() string {
	if o == nil || IsNil(o.NugetVersion) {
		var ret string
		return ret
	}
	return *o.NugetVersion
}

// GetNugetVersionOk returns a tuple with the NugetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NugetAttributes) GetNugetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.NugetVersion) {
		return nil, false
	}
	return o.NugetVersion, true
}

// HasNugetVersion returns a boolean if a field has been set.
func (o *NugetAttributes) HasNugetVersion() bool {
	if o != nil && !IsNil(o.NugetVersion) {
		return true
	}

	return false
}

// SetNugetVersion gets a reference to the given string and assigns it to the NugetVersion field.
func (o *NugetAttributes) SetNugetVersion(v string) {
	o.NugetVersion = &v
}

// GetQueryCacheItemMaxAge returns the QueryCacheItemMaxAge field value if set, zero value otherwise.
func (o *NugetAttributes) GetQueryCacheItemMaxAge() int32 {
	if o == nil || IsNil(o.QueryCacheItemMaxAge) {
		var ret int32
		return ret
	}
	return *o.QueryCacheItemMaxAge
}

// GetQueryCacheItemMaxAgeOk returns a tuple with the QueryCacheItemMaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NugetAttributes) GetQueryCacheItemMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.QueryCacheItemMaxAge) {
		return nil, false
	}
	return o.QueryCacheItemMaxAge, true
}

// HasQueryCacheItemMaxAge returns a boolean if a field has been set.
func (o *NugetAttributes) HasQueryCacheItemMaxAge() bool {
	if o != nil && !IsNil(o.QueryCacheItemMaxAge) {
		return true
	}

	return false
}

// SetQueryCacheItemMaxAge gets a reference to the given int32 and assigns it to the QueryCacheItemMaxAge field.
func (o *NugetAttributes) SetQueryCacheItemMaxAge(v int32) {
	o.QueryCacheItemMaxAge = &v
}

func (o NugetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NugetAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NugetVersion) {
		toSerialize["nugetVersion"] = o.NugetVersion
	}
	if !IsNil(o.QueryCacheItemMaxAge) {
		toSerialize["queryCacheItemMaxAge"] = o.QueryCacheItemMaxAge
	}
	return toSerialize, nil
}

type NullableNugetAttributes struct {
	value *NugetAttributes
	isSet bool
}

func (v NullableNugetAttributes) Get() *NugetAttributes {
	return v.value
}

func (v *NullableNugetAttributes) Set(val *NugetAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNugetAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNugetAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNugetAttributes(val *NugetAttributes) *NullableNugetAttributes {
	return &NullableNugetAttributes{value: val, isSet: true}
}

func (v NullableNugetAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNugetAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


