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

// checks if the AuthSettingsXo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthSettingsXo{}

// AuthSettingsXo struct for AuthSettingsXo
type AuthSettingsXo struct {
	// auth enabled
	Enabled bool `json:"enabled"`
	// Windows NTLM Domain
	NtlmDomain string `json:"ntlmDomain"`
	// Windows NTLM Hostname
	NtlmHost string `json:"ntlmHost"`
	// user password
	Password string `json:"password"`
	// user name
	Username string `json:"username"`
}

type _AuthSettingsXo AuthSettingsXo

// NewAuthSettingsXo instantiates a new AuthSettingsXo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthSettingsXo(enabled bool, ntlmDomain string, ntlmHost string, password string, username string) *AuthSettingsXo {
	this := AuthSettingsXo{}
	this.Enabled = enabled
	this.NtlmDomain = ntlmDomain
	this.NtlmHost = ntlmHost
	this.Password = password
	this.Username = username
	return &this
}

// NewAuthSettingsXoWithDefaults instantiates a new AuthSettingsXo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthSettingsXoWithDefaults() *AuthSettingsXo {
	this := AuthSettingsXo{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *AuthSettingsXo) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AuthSettingsXo) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AuthSettingsXo) SetEnabled(v bool) {
	o.Enabled = v
}

// GetNtlmDomain returns the NtlmDomain field value
func (o *AuthSettingsXo) GetNtlmDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NtlmDomain
}

// GetNtlmDomainOk returns a tuple with the NtlmDomain field value
// and a boolean to check if the value has been set.
func (o *AuthSettingsXo) GetNtlmDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NtlmDomain, true
}

// SetNtlmDomain sets field value
func (o *AuthSettingsXo) SetNtlmDomain(v string) {
	o.NtlmDomain = v
}

// GetNtlmHost returns the NtlmHost field value
func (o *AuthSettingsXo) GetNtlmHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NtlmHost
}

// GetNtlmHostOk returns a tuple with the NtlmHost field value
// and a boolean to check if the value has been set.
func (o *AuthSettingsXo) GetNtlmHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NtlmHost, true
}

// SetNtlmHost sets field value
func (o *AuthSettingsXo) SetNtlmHost(v string) {
	o.NtlmHost = v
}

// GetPassword returns the Password field value
func (o *AuthSettingsXo) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthSettingsXo) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AuthSettingsXo) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *AuthSettingsXo) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthSettingsXo) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthSettingsXo) SetUsername(v string) {
	o.Username = v
}

func (o AuthSettingsXo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthSettingsXo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["ntlmDomain"] = o.NtlmDomain
	toSerialize["ntlmHost"] = o.NtlmHost
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *AuthSettingsXo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"ntlmDomain",
		"ntlmHost",
		"password",
		"username",
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

	varAuthSettingsXo := _AuthSettingsXo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthSettingsXo)

	if err != nil {
		return err
	}

	*o = AuthSettingsXo(varAuthSettingsXo)

	return err
}

type NullableAuthSettingsXo struct {
	value *AuthSettingsXo
	isSet bool
}

func (v NullableAuthSettingsXo) Get() *AuthSettingsXo {
	return v.value
}

func (v *NullableAuthSettingsXo) Set(val *AuthSettingsXo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthSettingsXo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthSettingsXo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthSettingsXo(val *AuthSettingsXo) *NullableAuthSettingsXo {
	return &NullableAuthSettingsXo{value: val, isSet: true}
}

func (v NullableAuthSettingsXo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthSettingsXo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


