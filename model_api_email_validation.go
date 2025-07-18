/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.82.0-08.

API version: 3.82.0-08
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the ApiEmailValidation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiEmailValidation{}

// ApiEmailValidation struct for ApiEmailValidation
type ApiEmailValidation struct {
	Reason *string `json:"reason,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewApiEmailValidation instantiates a new ApiEmailValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiEmailValidation() *ApiEmailValidation {
	this := ApiEmailValidation{}
	return &this
}

// NewApiEmailValidationWithDefaults instantiates a new ApiEmailValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiEmailValidationWithDefaults() *ApiEmailValidation {
	this := ApiEmailValidation{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ApiEmailValidation) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailValidation) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ApiEmailValidation) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ApiEmailValidation) SetReason(v string) {
	o.Reason = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiEmailValidation) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailValidation) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiEmailValidation) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiEmailValidation) SetSuccess(v bool) {
	o.Success = &v
}

func (o ApiEmailValidation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiEmailValidation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableApiEmailValidation struct {
	value *ApiEmailValidation
	isSet bool
}

func (v NullableApiEmailValidation) Get() *ApiEmailValidation {
	return v.value
}

func (v *NullableApiEmailValidation) Set(val *ApiEmailValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableApiEmailValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableApiEmailValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiEmailValidation(val *ApiEmailValidation) *NullableApiEmailValidation {
	return &NullableApiEmailValidation{value: val, isSet: true}
}

func (v NullableApiEmailValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiEmailValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


