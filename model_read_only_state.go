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

// checks if the ReadOnlyState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadOnlyState{}

// ReadOnlyState struct for ReadOnlyState
type ReadOnlyState struct {
	Frozen *bool `json:"frozen,omitempty"`
	SummaryReason *string `json:"summaryReason,omitempty"`
	SystemInitiated *bool `json:"systemInitiated,omitempty"`
}

// NewReadOnlyState instantiates a new ReadOnlyState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlyState() *ReadOnlyState {
	this := ReadOnlyState{}
	return &this
}

// NewReadOnlyStateWithDefaults instantiates a new ReadOnlyState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlyStateWithDefaults() *ReadOnlyState {
	this := ReadOnlyState{}
	return &this
}

// GetFrozen returns the Frozen field value if set, zero value otherwise.
func (o *ReadOnlyState) GetFrozen() bool {
	if o == nil || IsNil(o.Frozen) {
		var ret bool
		return ret
	}
	return *o.Frozen
}

// GetFrozenOk returns a tuple with the Frozen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyState) GetFrozenOk() (*bool, bool) {
	if o == nil || IsNil(o.Frozen) {
		return nil, false
	}
	return o.Frozen, true
}

// HasFrozen returns a boolean if a field has been set.
func (o *ReadOnlyState) HasFrozen() bool {
	if o != nil && !IsNil(o.Frozen) {
		return true
	}

	return false
}

// SetFrozen gets a reference to the given bool and assigns it to the Frozen field.
func (o *ReadOnlyState) SetFrozen(v bool) {
	o.Frozen = &v
}

// GetSummaryReason returns the SummaryReason field value if set, zero value otherwise.
func (o *ReadOnlyState) GetSummaryReason() string {
	if o == nil || IsNil(o.SummaryReason) {
		var ret string
		return ret
	}
	return *o.SummaryReason
}

// GetSummaryReasonOk returns a tuple with the SummaryReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyState) GetSummaryReasonOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryReason) {
		return nil, false
	}
	return o.SummaryReason, true
}

// HasSummaryReason returns a boolean if a field has been set.
func (o *ReadOnlyState) HasSummaryReason() bool {
	if o != nil && !IsNil(o.SummaryReason) {
		return true
	}

	return false
}

// SetSummaryReason gets a reference to the given string and assigns it to the SummaryReason field.
func (o *ReadOnlyState) SetSummaryReason(v string) {
	o.SummaryReason = &v
}

// GetSystemInitiated returns the SystemInitiated field value if set, zero value otherwise.
func (o *ReadOnlyState) GetSystemInitiated() bool {
	if o == nil || IsNil(o.SystemInitiated) {
		var ret bool
		return ret
	}
	return *o.SystemInitiated
}

// GetSystemInitiatedOk returns a tuple with the SystemInitiated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyState) GetSystemInitiatedOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemInitiated) {
		return nil, false
	}
	return o.SystemInitiated, true
}

// HasSystemInitiated returns a boolean if a field has been set.
func (o *ReadOnlyState) HasSystemInitiated() bool {
	if o != nil && !IsNil(o.SystemInitiated) {
		return true
	}

	return false
}

// SetSystemInitiated gets a reference to the given bool and assigns it to the SystemInitiated field.
func (o *ReadOnlyState) SetSystemInitiated(v bool) {
	o.SystemInitiated = &v
}

func (o ReadOnlyState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadOnlyState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Frozen) {
		toSerialize["frozen"] = o.Frozen
	}
	if !IsNil(o.SummaryReason) {
		toSerialize["summaryReason"] = o.SummaryReason
	}
	if !IsNil(o.SystemInitiated) {
		toSerialize["systemInitiated"] = o.SystemInitiated
	}
	return toSerialize, nil
}

type NullableReadOnlyState struct {
	value *ReadOnlyState
	isSet bool
}

func (v NullableReadOnlyState) Get() *ReadOnlyState {
	return v.value
}

func (v *NullableReadOnlyState) Set(val *ReadOnlyState) {
	v.value = val
	v.isSet = true
}

func (v NullableReadOnlyState) IsSet() bool {
	return v.isSet
}

func (v *NullableReadOnlyState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadOnlyState(val *ReadOnlyState) *NullableReadOnlyState {
	return &NullableReadOnlyState{value: val, isSet: true}
}

func (v NullableReadOnlyState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadOnlyState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


