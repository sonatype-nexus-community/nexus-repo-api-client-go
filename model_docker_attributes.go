/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.67.1-01.

API version: 3.67.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DockerAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DockerAttributes{}

// DockerAttributes struct for DockerAttributes
type DockerAttributes struct {
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	ForceBasicAuth bool `json:"forceBasicAuth"`
	// Create an HTTP connector at specified port
	HttpPort *int32 `json:"httpPort,omitempty"`
	// Create an HTTPS connector at specified port
	HttpsPort *int32 `json:"httpsPort,omitempty"`
	// Allows to use subdomain
	Subdomain *string `json:"subdomain,omitempty"`
	// Whether to allow clients to use the V1 API to interact with this repository
	V1Enabled bool `json:"v1Enabled"`
}

type _DockerAttributes DockerAttributes

// NewDockerAttributes instantiates a new DockerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerAttributes(forceBasicAuth bool, v1Enabled bool) *DockerAttributes {
	this := DockerAttributes{}
	this.ForceBasicAuth = forceBasicAuth
	this.V1Enabled = v1Enabled
	return &this
}

// NewDockerAttributesWithDefaults instantiates a new DockerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerAttributesWithDefaults() *DockerAttributes {
	this := DockerAttributes{}
	return &this
}

// GetForceBasicAuth returns the ForceBasicAuth field value
func (o *DockerAttributes) GetForceBasicAuth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForceBasicAuth
}

// GetForceBasicAuthOk returns a tuple with the ForceBasicAuth field value
// and a boolean to check if the value has been set.
func (o *DockerAttributes) GetForceBasicAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForceBasicAuth, true
}

// SetForceBasicAuth sets field value
func (o *DockerAttributes) SetForceBasicAuth(v bool) {
	o.ForceBasicAuth = v
}

// GetHttpPort returns the HttpPort field value if set, zero value otherwise.
func (o *DockerAttributes) GetHttpPort() int32 {
	if o == nil || IsNil(o.HttpPort) {
		var ret int32
		return ret
	}
	return *o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerAttributes) GetHttpPortOk() (*int32, bool) {
	if o == nil || IsNil(o.HttpPort) {
		return nil, false
	}
	return o.HttpPort, true
}

// HasHttpPort returns a boolean if a field has been set.
func (o *DockerAttributes) HasHttpPort() bool {
	if o != nil && !IsNil(o.HttpPort) {
		return true
	}

	return false
}

// SetHttpPort gets a reference to the given int32 and assigns it to the HttpPort field.
func (o *DockerAttributes) SetHttpPort(v int32) {
	o.HttpPort = &v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *DockerAttributes) GetHttpsPort() int32 {
	if o == nil || IsNil(o.HttpsPort) {
		var ret int32
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerAttributes) GetHttpsPortOk() (*int32, bool) {
	if o == nil || IsNil(o.HttpsPort) {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *DockerAttributes) HasHttpsPort() bool {
	if o != nil && !IsNil(o.HttpsPort) {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int32 and assigns it to the HttpsPort field.
func (o *DockerAttributes) SetHttpsPort(v int32) {
	o.HttpsPort = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *DockerAttributes) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerAttributes) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *DockerAttributes) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *DockerAttributes) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetV1Enabled returns the V1Enabled field value
func (o *DockerAttributes) GetV1Enabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.V1Enabled
}

// GetV1EnabledOk returns a tuple with the V1Enabled field value
// and a boolean to check if the value has been set.
func (o *DockerAttributes) GetV1EnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.V1Enabled, true
}

// SetV1Enabled sets field value
func (o *DockerAttributes) SetV1Enabled(v bool) {
	o.V1Enabled = v
}

func (o DockerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DockerAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["forceBasicAuth"] = o.ForceBasicAuth
	if !IsNil(o.HttpPort) {
		toSerialize["httpPort"] = o.HttpPort
	}
	if !IsNil(o.HttpsPort) {
		toSerialize["httpsPort"] = o.HttpsPort
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	toSerialize["v1Enabled"] = o.V1Enabled
	return toSerialize, nil
}

func (o *DockerAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"forceBasicAuth",
		"v1Enabled",
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

	varDockerAttributes := _DockerAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDockerAttributes)

	if err != nil {
		return err
	}

	*o = DockerAttributes(varDockerAttributes)

	return err
}

type NullableDockerAttributes struct {
	value *DockerAttributes
	isSet bool
}

func (v NullableDockerAttributes) Get() *DockerAttributes {
	return v.value
}

func (v *NullableDockerAttributes) Set(val *DockerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerAttributes(val *DockerAttributes) *NullableDockerAttributes {
	return &NullableDockerAttributes{value: val, isSet: true}
}

func (v NullableDockerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


