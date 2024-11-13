/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.74.0-05.

API version: 3.74.0-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
)

// checks if the RawAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RawAttributes{}

// RawAttributes struct for RawAttributes
type RawAttributes struct {
	// Content Disposition
	ContentDisposition *string `json:"contentDisposition,omitempty"`
}

// NewRawAttributes instantiates a new RawAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawAttributes() *RawAttributes {
	this := RawAttributes{}
	return &this
}

// NewRawAttributesWithDefaults instantiates a new RawAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawAttributesWithDefaults() *RawAttributes {
	this := RawAttributes{}
	return &this
}

// GetContentDisposition returns the ContentDisposition field value if set, zero value otherwise.
func (o *RawAttributes) GetContentDisposition() string {
	if o == nil || IsNil(o.ContentDisposition) {
		var ret string
		return ret
	}
	return *o.ContentDisposition
}

// GetContentDispositionOk returns a tuple with the ContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawAttributes) GetContentDispositionOk() (*string, bool) {
	if o == nil || IsNil(o.ContentDisposition) {
		return nil, false
	}
	return o.ContentDisposition, true
}

// HasContentDisposition returns a boolean if a field has been set.
func (o *RawAttributes) HasContentDisposition() bool {
	if o != nil && !IsNil(o.ContentDisposition) {
		return true
	}

	return false
}

// SetContentDisposition gets a reference to the given string and assigns it to the ContentDisposition field.
func (o *RawAttributes) SetContentDisposition(v string) {
	o.ContentDisposition = &v
}

func (o RawAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RawAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentDisposition) {
		toSerialize["contentDisposition"] = o.ContentDisposition
	}
	return toSerialize, nil
}

type NullableRawAttributes struct {
	value *RawAttributes
	isSet bool
}

func (v NullableRawAttributes) Get() *RawAttributes {
	return v.value
}

func (v *NullableRawAttributes) Set(val *RawAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableRawAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableRawAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawAttributes(val *RawAttributes) *NullableRawAttributes {
	return &NullableRawAttributes{value: val, isSet: true}
}

func (v NullableRawAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


