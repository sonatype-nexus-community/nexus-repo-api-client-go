/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.68.1-02.

API version: 3.68.1-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
)

// checks if the S3BlobStoreApiEncryption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3BlobStoreApiEncryption{}

// S3BlobStoreApiEncryption struct for S3BlobStoreApiEncryption
type S3BlobStoreApiEncryption struct {
	// The encryption key.
	EncryptionKey *string `json:"encryptionKey,omitempty"`
	// The type of S3 server side encryption to use.
	EncryptionType *string `json:"encryptionType,omitempty"`
}

// NewS3BlobStoreApiEncryption instantiates a new S3BlobStoreApiEncryption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3BlobStoreApiEncryption() *S3BlobStoreApiEncryption {
	this := S3BlobStoreApiEncryption{}
	return &this
}

// NewS3BlobStoreApiEncryptionWithDefaults instantiates a new S3BlobStoreApiEncryption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3BlobStoreApiEncryptionWithDefaults() *S3BlobStoreApiEncryption {
	this := S3BlobStoreApiEncryption{}
	return &this
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *S3BlobStoreApiEncryption) GetEncryptionKey() string {
	if o == nil || IsNil(o.EncryptionKey) {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiEncryption) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionKey) {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *S3BlobStoreApiEncryption) HasEncryptionKey() bool {
	if o != nil && !IsNil(o.EncryptionKey) {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *S3BlobStoreApiEncryption) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetEncryptionType returns the EncryptionType field value if set, zero value otherwise.
func (o *S3BlobStoreApiEncryption) GetEncryptionType() string {
	if o == nil || IsNil(o.EncryptionType) {
		var ret string
		return ret
	}
	return *o.EncryptionType
}

// GetEncryptionTypeOk returns a tuple with the EncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiEncryption) GetEncryptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionType) {
		return nil, false
	}
	return o.EncryptionType, true
}

// HasEncryptionType returns a boolean if a field has been set.
func (o *S3BlobStoreApiEncryption) HasEncryptionType() bool {
	if o != nil && !IsNil(o.EncryptionType) {
		return true
	}

	return false
}

// SetEncryptionType gets a reference to the given string and assigns it to the EncryptionType field.
func (o *S3BlobStoreApiEncryption) SetEncryptionType(v string) {
	o.EncryptionType = &v
}

func (o S3BlobStoreApiEncryption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3BlobStoreApiEncryption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EncryptionKey) {
		toSerialize["encryptionKey"] = o.EncryptionKey
	}
	if !IsNil(o.EncryptionType) {
		toSerialize["encryptionType"] = o.EncryptionType
	}
	return toSerialize, nil
}

type NullableS3BlobStoreApiEncryption struct {
	value *S3BlobStoreApiEncryption
	isSet bool
}

func (v NullableS3BlobStoreApiEncryption) Get() *S3BlobStoreApiEncryption {
	return v.value
}

func (v *NullableS3BlobStoreApiEncryption) Set(val *S3BlobStoreApiEncryption) {
	v.value = val
	v.isSet = true
}

func (v NullableS3BlobStoreApiEncryption) IsSet() bool {
	return v.isSet
}

func (v *NullableS3BlobStoreApiEncryption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3BlobStoreApiEncryption(val *S3BlobStoreApiEncryption) *NullableS3BlobStoreApiEncryption {
	return &NullableS3BlobStoreApiEncryption{value: val, isSet: true}
}

func (v NullableS3BlobStoreApiEncryption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3BlobStoreApiEncryption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


