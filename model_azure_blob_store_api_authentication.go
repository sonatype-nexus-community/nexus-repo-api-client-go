/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.74.0-05.

API version: 3.74.0-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AzureBlobStoreApiAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureBlobStoreApiAuthentication{}

// AzureBlobStoreApiAuthentication struct for AzureBlobStoreApiAuthentication
type AzureBlobStoreApiAuthentication struct {
	// The account key.
	AccountKey *string `json:"accountKey,omitempty"`
	// The type of Azure authentication to use.
	AuthenticationMethod string `json:"authenticationMethod"`
}

type _AzureBlobStoreApiAuthentication AzureBlobStoreApiAuthentication

// NewAzureBlobStoreApiAuthentication instantiates a new AzureBlobStoreApiAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStoreApiAuthentication(authenticationMethod string) *AzureBlobStoreApiAuthentication {
	this := AzureBlobStoreApiAuthentication{}
	this.AuthenticationMethod = authenticationMethod
	return &this
}

// NewAzureBlobStoreApiAuthenticationWithDefaults instantiates a new AzureBlobStoreApiAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStoreApiAuthenticationWithDefaults() *AzureBlobStoreApiAuthentication {
	this := AzureBlobStoreApiAuthentication{}
	return &this
}

// GetAccountKey returns the AccountKey field value if set, zero value otherwise.
func (o *AzureBlobStoreApiAuthentication) GetAccountKey() string {
	if o == nil || IsNil(o.AccountKey) {
		var ret string
		return ret
	}
	return *o.AccountKey
}

// GetAccountKeyOk returns a tuple with the AccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStoreApiAuthentication) GetAccountKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccountKey) {
		return nil, false
	}
	return o.AccountKey, true
}

// HasAccountKey returns a boolean if a field has been set.
func (o *AzureBlobStoreApiAuthentication) HasAccountKey() bool {
	if o != nil && !IsNil(o.AccountKey) {
		return true
	}

	return false
}

// SetAccountKey gets a reference to the given string and assigns it to the AccountKey field.
func (o *AzureBlobStoreApiAuthentication) SetAccountKey(v string) {
	o.AccountKey = &v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *AzureBlobStoreApiAuthentication) GetAuthenticationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStoreApiAuthentication) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *AzureBlobStoreApiAuthentication) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = v
}

func (o AzureBlobStoreApiAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureBlobStoreApiAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountKey) {
		toSerialize["accountKey"] = o.AccountKey
	}
	toSerialize["authenticationMethod"] = o.AuthenticationMethod
	return toSerialize, nil
}

func (o *AzureBlobStoreApiAuthentication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authenticationMethod",
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

	varAzureBlobStoreApiAuthentication := _AzureBlobStoreApiAuthentication{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAzureBlobStoreApiAuthentication)

	if err != nil {
		return err
	}

	*o = AzureBlobStoreApiAuthentication(varAzureBlobStoreApiAuthentication)

	return err
}

type NullableAzureBlobStoreApiAuthentication struct {
	value *AzureBlobStoreApiAuthentication
	isSet bool
}

func (v NullableAzureBlobStoreApiAuthentication) Get() *AzureBlobStoreApiAuthentication {
	return v.value
}

func (v *NullableAzureBlobStoreApiAuthentication) Set(val *AzureBlobStoreApiAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStoreApiAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStoreApiAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStoreApiAuthentication(val *AzureBlobStoreApiAuthentication) *NullableAzureBlobStoreApiAuthentication {
	return &NullableAzureBlobStoreApiAuthentication{value: val, isSet: true}
}

func (v NullableAzureBlobStoreApiAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStoreApiAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


