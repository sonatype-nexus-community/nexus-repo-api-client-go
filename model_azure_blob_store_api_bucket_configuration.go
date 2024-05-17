/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.68.1-02.

API version: 3.68.1-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AzureBlobStoreApiBucketConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureBlobStoreApiBucketConfiguration{}

// AzureBlobStoreApiBucketConfiguration struct for AzureBlobStoreApiBucketConfiguration
type AzureBlobStoreApiBucketConfiguration struct {
	// Account name found under Access keys for the storage account.
	AccountName string `json:"accountName"`
	Authentication AzureBlobStoreApiAuthentication `json:"authentication"`
	// The name of an existing container to be used for storage.
	ContainerName string `json:"containerName"`
}

type _AzureBlobStoreApiBucketConfiguration AzureBlobStoreApiBucketConfiguration

// NewAzureBlobStoreApiBucketConfiguration instantiates a new AzureBlobStoreApiBucketConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStoreApiBucketConfiguration(accountName string, authentication AzureBlobStoreApiAuthentication, containerName string) *AzureBlobStoreApiBucketConfiguration {
	this := AzureBlobStoreApiBucketConfiguration{}
	this.AccountName = accountName
	this.Authentication = authentication
	this.ContainerName = containerName
	return &this
}

// NewAzureBlobStoreApiBucketConfigurationWithDefaults instantiates a new AzureBlobStoreApiBucketConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStoreApiBucketConfigurationWithDefaults() *AzureBlobStoreApiBucketConfiguration {
	this := AzureBlobStoreApiBucketConfiguration{}
	return &this
}

// GetAccountName returns the AccountName field value
func (o *AzureBlobStoreApiBucketConfiguration) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStoreApiBucketConfiguration) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *AzureBlobStoreApiBucketConfiguration) SetAccountName(v string) {
	o.AccountName = v
}

// GetAuthentication returns the Authentication field value
func (o *AzureBlobStoreApiBucketConfiguration) GetAuthentication() AzureBlobStoreApiAuthentication {
	if o == nil {
		var ret AzureBlobStoreApiAuthentication
		return ret
	}

	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStoreApiBucketConfiguration) GetAuthenticationOk() (*AzureBlobStoreApiAuthentication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value
func (o *AzureBlobStoreApiBucketConfiguration) SetAuthentication(v AzureBlobStoreApiAuthentication) {
	o.Authentication = v
}

// GetContainerName returns the ContainerName field value
func (o *AzureBlobStoreApiBucketConfiguration) GetContainerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStoreApiBucketConfiguration) GetContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerName, true
}

// SetContainerName sets field value
func (o *AzureBlobStoreApiBucketConfiguration) SetContainerName(v string) {
	o.ContainerName = v
}

func (o AzureBlobStoreApiBucketConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureBlobStoreApiBucketConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountName"] = o.AccountName
	toSerialize["authentication"] = o.Authentication
	toSerialize["containerName"] = o.ContainerName
	return toSerialize, nil
}

func (o *AzureBlobStoreApiBucketConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountName",
		"authentication",
		"containerName",
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

	varAzureBlobStoreApiBucketConfiguration := _AzureBlobStoreApiBucketConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAzureBlobStoreApiBucketConfiguration)

	if err != nil {
		return err
	}

	*o = AzureBlobStoreApiBucketConfiguration(varAzureBlobStoreApiBucketConfiguration)

	return err
}

type NullableAzureBlobStoreApiBucketConfiguration struct {
	value *AzureBlobStoreApiBucketConfiguration
	isSet bool
}

func (v NullableAzureBlobStoreApiBucketConfiguration) Get() *AzureBlobStoreApiBucketConfiguration {
	return v.value
}

func (v *NullableAzureBlobStoreApiBucketConfiguration) Set(val *AzureBlobStoreApiBucketConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStoreApiBucketConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStoreApiBucketConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStoreApiBucketConfiguration(val *AzureBlobStoreApiBucketConfiguration) *NullableAzureBlobStoreApiBucketConfiguration {
	return &NullableAzureBlobStoreApiBucketConfiguration{value: val, isSet: true}
}

func (v NullableAzureBlobStoreApiBucketConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStoreApiBucketConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


