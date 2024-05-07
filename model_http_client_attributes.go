/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.67.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HttpClientAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpClientAttributes{}

// HttpClientAttributes struct for HttpClientAttributes
type HttpClientAttributes struct {
	Authentication *HttpClientConnectionAuthenticationAttributes `json:"authentication,omitempty"`
	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	AutoBlock bool `json:"autoBlock"`
	// Whether to block outbound connections on the repository
	Blocked bool `json:"blocked"`
	Connection *HttpClientConnectionAttributes `json:"connection,omitempty"`
}

type _HttpClientAttributes HttpClientAttributes

// NewHttpClientAttributes instantiates a new HttpClientAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpClientAttributes(autoBlock bool, blocked bool) *HttpClientAttributes {
	this := HttpClientAttributes{}
	this.AutoBlock = autoBlock
	this.Blocked = blocked
	return &this
}

// NewHttpClientAttributesWithDefaults instantiates a new HttpClientAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpClientAttributesWithDefaults() *HttpClientAttributes {
	this := HttpClientAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *HttpClientAttributes) GetAuthentication() HttpClientConnectionAuthenticationAttributes {
	if o == nil || IsNil(o.Authentication) {
		var ret HttpClientConnectionAuthenticationAttributes
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientAttributes) GetAuthenticationOk() (*HttpClientConnectionAuthenticationAttributes, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *HttpClientAttributes) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given HttpClientConnectionAuthenticationAttributes and assigns it to the Authentication field.
func (o *HttpClientAttributes) SetAuthentication(v HttpClientConnectionAuthenticationAttributes) {
	o.Authentication = &v
}

// GetAutoBlock returns the AutoBlock field value
func (o *HttpClientAttributes) GetAutoBlock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoBlock
}

// GetAutoBlockOk returns a tuple with the AutoBlock field value
// and a boolean to check if the value has been set.
func (o *HttpClientAttributes) GetAutoBlockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoBlock, true
}

// SetAutoBlock sets field value
func (o *HttpClientAttributes) SetAutoBlock(v bool) {
	o.AutoBlock = v
}

// GetBlocked returns the Blocked field value
func (o *HttpClientAttributes) GetBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value
// and a boolean to check if the value has been set.
func (o *HttpClientAttributes) GetBlockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blocked, true
}

// SetBlocked sets field value
func (o *HttpClientAttributes) SetBlocked(v bool) {
	o.Blocked = v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *HttpClientAttributes) GetConnection() HttpClientConnectionAttributes {
	if o == nil || IsNil(o.Connection) {
		var ret HttpClientConnectionAttributes
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientAttributes) GetConnectionOk() (*HttpClientConnectionAttributes, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *HttpClientAttributes) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given HttpClientConnectionAttributes and assigns it to the Connection field.
func (o *HttpClientAttributes) SetConnection(v HttpClientConnectionAttributes) {
	o.Connection = &v
}

func (o HttpClientAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpClientAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	toSerialize["autoBlock"] = o.AutoBlock
	toSerialize["blocked"] = o.Blocked
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	return toSerialize, nil
}

func (o *HttpClientAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"autoBlock",
		"blocked",
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

	varHttpClientAttributes := _HttpClientAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHttpClientAttributes)

	if err != nil {
		return err
	}

	*o = HttpClientAttributes(varHttpClientAttributes)

	return err
}

type NullableHttpClientAttributes struct {
	value *HttpClientAttributes
	isSet bool
}

func (v NullableHttpClientAttributes) Get() *HttpClientAttributes {
	return v.value
}

func (v *NullableHttpClientAttributes) Set(val *HttpClientAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpClientAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpClientAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpClientAttributes(val *HttpClientAttributes) *NullableHttpClientAttributes {
	return &NullableHttpClientAttributes{value: val, isSet: true}
}

func (v NullableHttpClientAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpClientAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


