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

// checks if the HttpClientConnectionAuthenticationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpClientConnectionAuthenticationAttributes{}

// HttpClientConnectionAuthenticationAttributes struct for HttpClientConnectionAuthenticationAttributes
type HttpClientConnectionAuthenticationAttributes struct {
	BearerToken *string `json:"bearerToken,omitempty"`
	NtlmDomain *string `json:"ntlmDomain,omitempty"`
	NtlmHost *string `json:"ntlmHost,omitempty"`
	Password *string `json:"password,omitempty"`
	// Whether to use pre-emptive authentication. Use with caution. Defaults to false.
	Preemptive *bool `json:"preemptive,omitempty"`
	// Authentication type
	Type *string `json:"type,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewHttpClientConnectionAuthenticationAttributes instantiates a new HttpClientConnectionAuthenticationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpClientConnectionAuthenticationAttributes() *HttpClientConnectionAuthenticationAttributes {
	this := HttpClientConnectionAuthenticationAttributes{}
	return &this
}

// NewHttpClientConnectionAuthenticationAttributesWithDefaults instantiates a new HttpClientConnectionAuthenticationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpClientConnectionAuthenticationAttributesWithDefaults() *HttpClientConnectionAuthenticationAttributes {
	this := HttpClientConnectionAuthenticationAttributes{}
	return &this
}

// GetBearerToken returns the BearerToken field value if set, zero value otherwise.
func (o *HttpClientConnectionAuthenticationAttributes) GetBearerToken() string {
	if o == nil || IsNil(o.BearerToken) {
		var ret string
		return ret
	}
	return *o.BearerToken
}

// GetBearerTokenOk returns a tuple with the BearerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAuthenticationAttributes) GetBearerTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BearerToken) {
		return nil, false
	}
	return o.BearerToken, true
}

// HasBearerToken returns a boolean if a field has been set.
func (o *HttpClientConnectionAuthenticationAttributes) HasBearerToken() bool {
	if o != nil && !IsNil(o.BearerToken) {
		return true
	}

	return false
}

// SetBearerToken gets a reference to the given string and assigns it to the BearerToken field.
func (o *HttpClientConnectionAuthenticationAttributes) SetBearerToken(v string) {
	o.BearerToken = &v
}

// GetNtlmDomain returns the NtlmDomain field value if set, zero value otherwise.
func (o *HttpClientConnectionAuthenticationAttributes) GetNtlmDomain() string {
	if o == nil || IsNil(o.NtlmDomain) {
		var ret string
		return ret
	}
	return *o.NtlmDomain
}

// GetNtlmDomainOk returns a tuple with the NtlmDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAuthenticationAttributes) GetNtlmDomainOk() (*string, bool) {
	if o == nil || IsNil(o.NtlmDomain) {
		return nil, false
	}
	return o.NtlmDomain, true
}

// HasNtlmDomain returns a boolean if a field has been set.
func (o *HttpClientConnectionAuthenticationAttributes) HasNtlmDomain() bool {
	if o != nil && !IsNil(o.NtlmDomain) {
		return true
	}

	return false
}

// SetNtlmDomain gets a reference to the given string and assigns it to the NtlmDomain field.
func (o *HttpClientConnectionAuthenticationAttributes) SetNtlmDomain(v string) {
	o.NtlmDomain = &v
}

// GetNtlmHost returns the NtlmHost field value if set, zero value otherwise.
func (o *HttpClientConnectionAuthenticationAttributes) GetNtlmHost() string {
	if o == nil || IsNil(o.NtlmHost) {
		var ret string
		return ret
	}
	return *o.NtlmHost
}

// GetNtlmHostOk returns a tuple with the NtlmHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAuthenticationAttributes) GetNtlmHostOk() (*string, bool) {
	if o == nil || IsNil(o.NtlmHost) {
		return nil, false
	}
	return o.NtlmHost, true
}

// HasNtlmHost returns a boolean if a field has been set.
func (o *HttpClientConnectionAuthenticationAttributes) HasNtlmHost() bool {
	if o != nil && !IsNil(o.NtlmHost) {
		return true
	}

	return false
}

// SetNtlmHost gets a reference to the given string and assigns it to the NtlmHost field.
func (o *HttpClientConnectionAuthenticationAttributes) SetNtlmHost(v string) {
	o.NtlmHost = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *HttpClientConnectionAuthenticationAttributes) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAuthenticationAttributes) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *HttpClientConnectionAuthenticationAttributes) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *HttpClientConnectionAuthenticationAttributes) SetPassword(v string) {
	o.Password = &v
}

// GetPreemptive returns the Preemptive field value if set, zero value otherwise.
func (o *HttpClientConnectionAuthenticationAttributes) GetPreemptive() bool {
	if o == nil || IsNil(o.Preemptive) {
		var ret bool
		return ret
	}
	return *o.Preemptive
}

// GetPreemptiveOk returns a tuple with the Preemptive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAuthenticationAttributes) GetPreemptiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Preemptive) {
		return nil, false
	}
	return o.Preemptive, true
}

// HasPreemptive returns a boolean if a field has been set.
func (o *HttpClientConnectionAuthenticationAttributes) HasPreemptive() bool {
	if o != nil && !IsNil(o.Preemptive) {
		return true
	}

	return false
}

// SetPreemptive gets a reference to the given bool and assigns it to the Preemptive field.
func (o *HttpClientConnectionAuthenticationAttributes) SetPreemptive(v bool) {
	o.Preemptive = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HttpClientConnectionAuthenticationAttributes) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAuthenticationAttributes) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HttpClientConnectionAuthenticationAttributes) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HttpClientConnectionAuthenticationAttributes) SetType(v string) {
	o.Type = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *HttpClientConnectionAuthenticationAttributes) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAuthenticationAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *HttpClientConnectionAuthenticationAttributes) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *HttpClientConnectionAuthenticationAttributes) SetUsername(v string) {
	o.Username = &v
}

func (o HttpClientConnectionAuthenticationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpClientConnectionAuthenticationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BearerToken) {
		toSerialize["bearerToken"] = o.BearerToken
	}
	if !IsNil(o.NtlmDomain) {
		toSerialize["ntlmDomain"] = o.NtlmDomain
	}
	if !IsNil(o.NtlmHost) {
		toSerialize["ntlmHost"] = o.NtlmHost
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Preemptive) {
		toSerialize["preemptive"] = o.Preemptive
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableHttpClientConnectionAuthenticationAttributes struct {
	value *HttpClientConnectionAuthenticationAttributes
	isSet bool
}

func (v NullableHttpClientConnectionAuthenticationAttributes) Get() *HttpClientConnectionAuthenticationAttributes {
	return v.value
}

func (v *NullableHttpClientConnectionAuthenticationAttributes) Set(val *HttpClientConnectionAuthenticationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpClientConnectionAuthenticationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpClientConnectionAuthenticationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpClientConnectionAuthenticationAttributes(val *HttpClientConnectionAuthenticationAttributes) *NullableHttpClientConnectionAuthenticationAttributes {
	return &NullableHttpClientConnectionAuthenticationAttributes{value: val, isSet: true}
}

func (v NullableHttpClientConnectionAuthenticationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpClientConnectionAuthenticationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


