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

// checks if the SystemCheckResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemCheckResultDTO{}

// SystemCheckResultDTO struct for SystemCheckResultDTO
type SystemCheckResultDTO struct {
	// Whether the system check succeeded of failed
	Healthy *bool `json:"healthy,omitempty"`
	// A description of the success or failure
	Message *string `json:"message,omitempty"`
}

// NewSystemCheckResultDTO instantiates a new SystemCheckResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemCheckResultDTO() *SystemCheckResultDTO {
	this := SystemCheckResultDTO{}
	return &this
}

// NewSystemCheckResultDTOWithDefaults instantiates a new SystemCheckResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemCheckResultDTOWithDefaults() *SystemCheckResultDTO {
	this := SystemCheckResultDTO{}
	return &this
}

// GetHealthy returns the Healthy field value if set, zero value otherwise.
func (o *SystemCheckResultDTO) GetHealthy() bool {
	if o == nil || IsNil(o.Healthy) {
		var ret bool
		return ret
	}
	return *o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemCheckResultDTO) GetHealthyOk() (*bool, bool) {
	if o == nil || IsNil(o.Healthy) {
		return nil, false
	}
	return o.Healthy, true
}

// HasHealthy returns a boolean if a field has been set.
func (o *SystemCheckResultDTO) HasHealthy() bool {
	if o != nil && !IsNil(o.Healthy) {
		return true
	}

	return false
}

// SetHealthy gets a reference to the given bool and assigns it to the Healthy field.
func (o *SystemCheckResultDTO) SetHealthy(v bool) {
	o.Healthy = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SystemCheckResultDTO) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemCheckResultDTO) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SystemCheckResultDTO) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SystemCheckResultDTO) SetMessage(v string) {
	o.Message = &v
}

func (o SystemCheckResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemCheckResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Healthy) {
		toSerialize["healthy"] = o.Healthy
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableSystemCheckResultDTO struct {
	value *SystemCheckResultDTO
	isSet bool
}

func (v NullableSystemCheckResultDTO) Get() *SystemCheckResultDTO {
	return v.value
}

func (v *NullableSystemCheckResultDTO) Set(val *SystemCheckResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemCheckResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemCheckResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemCheckResultDTO(val *SystemCheckResultDTO) *NullableSystemCheckResultDTO {
	return &NullableSystemCheckResultDTO{value: val, isSet: true}
}

func (v NullableSystemCheckResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemCheckResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


