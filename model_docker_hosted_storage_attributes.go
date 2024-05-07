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

// checks if the DockerHostedStorageAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DockerHostedStorageAttributes{}

// DockerHostedStorageAttributes struct for DockerHostedStorageAttributes
type DockerHostedStorageAttributes struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName"`
	// Whether to allow redeploying the 'latest' tag but defer to the Deployment Policy for all other tags
	LatestPolicy *bool `json:"latestPolicy,omitempty"`
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`
	// Controls if deployments of and updates to assets are allowed
	WritePolicy string `json:"writePolicy"`
}

type _DockerHostedStorageAttributes DockerHostedStorageAttributes

// NewDockerHostedStorageAttributes instantiates a new DockerHostedStorageAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerHostedStorageAttributes(blobStoreName string, strictContentTypeValidation bool, writePolicy string) *DockerHostedStorageAttributes {
	this := DockerHostedStorageAttributes{}
	this.BlobStoreName = blobStoreName
	this.StrictContentTypeValidation = strictContentTypeValidation
	this.WritePolicy = writePolicy
	return &this
}

// NewDockerHostedStorageAttributesWithDefaults instantiates a new DockerHostedStorageAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerHostedStorageAttributesWithDefaults() *DockerHostedStorageAttributes {
	this := DockerHostedStorageAttributes{}
	return &this
}

// GetBlobStoreName returns the BlobStoreName field value
func (o *DockerHostedStorageAttributes) GetBlobStoreName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlobStoreName
}

// GetBlobStoreNameOk returns a tuple with the BlobStoreName field value
// and a boolean to check if the value has been set.
func (o *DockerHostedStorageAttributes) GetBlobStoreNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlobStoreName, true
}

// SetBlobStoreName sets field value
func (o *DockerHostedStorageAttributes) SetBlobStoreName(v string) {
	o.BlobStoreName = v
}

// GetLatestPolicy returns the LatestPolicy field value if set, zero value otherwise.
func (o *DockerHostedStorageAttributes) GetLatestPolicy() bool {
	if o == nil || IsNil(o.LatestPolicy) {
		var ret bool
		return ret
	}
	return *o.LatestPolicy
}

// GetLatestPolicyOk returns a tuple with the LatestPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerHostedStorageAttributes) GetLatestPolicyOk() (*bool, bool) {
	if o == nil || IsNil(o.LatestPolicy) {
		return nil, false
	}
	return o.LatestPolicy, true
}

// HasLatestPolicy returns a boolean if a field has been set.
func (o *DockerHostedStorageAttributes) HasLatestPolicy() bool {
	if o != nil && !IsNil(o.LatestPolicy) {
		return true
	}

	return false
}

// SetLatestPolicy gets a reference to the given bool and assigns it to the LatestPolicy field.
func (o *DockerHostedStorageAttributes) SetLatestPolicy(v bool) {
	o.LatestPolicy = &v
}

// GetStrictContentTypeValidation returns the StrictContentTypeValidation field value
func (o *DockerHostedStorageAttributes) GetStrictContentTypeValidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StrictContentTypeValidation
}

// GetStrictContentTypeValidationOk returns a tuple with the StrictContentTypeValidation field value
// and a boolean to check if the value has been set.
func (o *DockerHostedStorageAttributes) GetStrictContentTypeValidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StrictContentTypeValidation, true
}

// SetStrictContentTypeValidation sets field value
func (o *DockerHostedStorageAttributes) SetStrictContentTypeValidation(v bool) {
	o.StrictContentTypeValidation = v
}

// GetWritePolicy returns the WritePolicy field value
func (o *DockerHostedStorageAttributes) GetWritePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WritePolicy
}

// GetWritePolicyOk returns a tuple with the WritePolicy field value
// and a boolean to check if the value has been set.
func (o *DockerHostedStorageAttributes) GetWritePolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WritePolicy, true
}

// SetWritePolicy sets field value
func (o *DockerHostedStorageAttributes) SetWritePolicy(v string) {
	o.WritePolicy = v
}

func (o DockerHostedStorageAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DockerHostedStorageAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blobStoreName"] = o.BlobStoreName
	if !IsNil(o.LatestPolicy) {
		toSerialize["latestPolicy"] = o.LatestPolicy
	}
	toSerialize["strictContentTypeValidation"] = o.StrictContentTypeValidation
	toSerialize["writePolicy"] = o.WritePolicy
	return toSerialize, nil
}

func (o *DockerHostedStorageAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"blobStoreName",
		"strictContentTypeValidation",
		"writePolicy",
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

	varDockerHostedStorageAttributes := _DockerHostedStorageAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDockerHostedStorageAttributes)

	if err != nil {
		return err
	}

	*o = DockerHostedStorageAttributes(varDockerHostedStorageAttributes)

	return err
}

type NullableDockerHostedStorageAttributes struct {
	value *DockerHostedStorageAttributes
	isSet bool
}

func (v NullableDockerHostedStorageAttributes) Get() *DockerHostedStorageAttributes {
	return v.value
}

func (v *NullableDockerHostedStorageAttributes) Set(val *DockerHostedStorageAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerHostedStorageAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerHostedStorageAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerHostedStorageAttributes(val *DockerHostedStorageAttributes) *NullableDockerHostedStorageAttributes {
	return &NullableDockerHostedStorageAttributes{value: val, isSet: true}
}

func (v NullableDockerHostedStorageAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerHostedStorageAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


