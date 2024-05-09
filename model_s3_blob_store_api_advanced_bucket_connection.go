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

// checks if the S3BlobStoreApiAdvancedBucketConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3BlobStoreApiAdvancedBucketConnection{}

// S3BlobStoreApiAdvancedBucketConnection struct for S3BlobStoreApiAdvancedBucketConnection
type S3BlobStoreApiAdvancedBucketConnection struct {
	// A custom endpoint URL for third party object stores using the S3 API.
	Endpoint *string `json:"endpoint,omitempty"`
	// Setting this flag will result in path-style access being used for all requests.
	ForcePathStyle *bool `json:"forcePathStyle,omitempty"`
	// Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	MaxConnectionPoolSize *int32 `json:"maxConnectionPoolSize,omitempty"`
	// An API signature version which may be required for third party object stores using the S3 API.
	SignerType *string `json:"signerType,omitempty"`
}

// NewS3BlobStoreApiAdvancedBucketConnection instantiates a new S3BlobStoreApiAdvancedBucketConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3BlobStoreApiAdvancedBucketConnection() *S3BlobStoreApiAdvancedBucketConnection {
	this := S3BlobStoreApiAdvancedBucketConnection{}
	return &this
}

// NewS3BlobStoreApiAdvancedBucketConnectionWithDefaults instantiates a new S3BlobStoreApiAdvancedBucketConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3BlobStoreApiAdvancedBucketConnectionWithDefaults() *S3BlobStoreApiAdvancedBucketConnection {
	this := S3BlobStoreApiAdvancedBucketConnection{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *S3BlobStoreApiAdvancedBucketConnection) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiAdvancedBucketConnection) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *S3BlobStoreApiAdvancedBucketConnection) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *S3BlobStoreApiAdvancedBucketConnection) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetForcePathStyle returns the ForcePathStyle field value if set, zero value otherwise.
func (o *S3BlobStoreApiAdvancedBucketConnection) GetForcePathStyle() bool {
	if o == nil || IsNil(o.ForcePathStyle) {
		var ret bool
		return ret
	}
	return *o.ForcePathStyle
}

// GetForcePathStyleOk returns a tuple with the ForcePathStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiAdvancedBucketConnection) GetForcePathStyleOk() (*bool, bool) {
	if o == nil || IsNil(o.ForcePathStyle) {
		return nil, false
	}
	return o.ForcePathStyle, true
}

// HasForcePathStyle returns a boolean if a field has been set.
func (o *S3BlobStoreApiAdvancedBucketConnection) HasForcePathStyle() bool {
	if o != nil && !IsNil(o.ForcePathStyle) {
		return true
	}

	return false
}

// SetForcePathStyle gets a reference to the given bool and assigns it to the ForcePathStyle field.
func (o *S3BlobStoreApiAdvancedBucketConnection) SetForcePathStyle(v bool) {
	o.ForcePathStyle = &v
}

// GetMaxConnectionPoolSize returns the MaxConnectionPoolSize field value if set, zero value otherwise.
func (o *S3BlobStoreApiAdvancedBucketConnection) GetMaxConnectionPoolSize() int32 {
	if o == nil || IsNil(o.MaxConnectionPoolSize) {
		var ret int32
		return ret
	}
	return *o.MaxConnectionPoolSize
}

// GetMaxConnectionPoolSizeOk returns a tuple with the MaxConnectionPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiAdvancedBucketConnection) GetMaxConnectionPoolSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConnectionPoolSize) {
		return nil, false
	}
	return o.MaxConnectionPoolSize, true
}

// HasMaxConnectionPoolSize returns a boolean if a field has been set.
func (o *S3BlobStoreApiAdvancedBucketConnection) HasMaxConnectionPoolSize() bool {
	if o != nil && !IsNil(o.MaxConnectionPoolSize) {
		return true
	}

	return false
}

// SetMaxConnectionPoolSize gets a reference to the given int32 and assigns it to the MaxConnectionPoolSize field.
func (o *S3BlobStoreApiAdvancedBucketConnection) SetMaxConnectionPoolSize(v int32) {
	o.MaxConnectionPoolSize = &v
}

// GetSignerType returns the SignerType field value if set, zero value otherwise.
func (o *S3BlobStoreApiAdvancedBucketConnection) GetSignerType() string {
	if o == nil || IsNil(o.SignerType) {
		var ret string
		return ret
	}
	return *o.SignerType
}

// GetSignerTypeOk returns a tuple with the SignerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiAdvancedBucketConnection) GetSignerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SignerType) {
		return nil, false
	}
	return o.SignerType, true
}

// HasSignerType returns a boolean if a field has been set.
func (o *S3BlobStoreApiAdvancedBucketConnection) HasSignerType() bool {
	if o != nil && !IsNil(o.SignerType) {
		return true
	}

	return false
}

// SetSignerType gets a reference to the given string and assigns it to the SignerType field.
func (o *S3BlobStoreApiAdvancedBucketConnection) SetSignerType(v string) {
	o.SignerType = &v
}

func (o S3BlobStoreApiAdvancedBucketConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3BlobStoreApiAdvancedBucketConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.ForcePathStyle) {
		toSerialize["forcePathStyle"] = o.ForcePathStyle
	}
	if !IsNil(o.MaxConnectionPoolSize) {
		toSerialize["maxConnectionPoolSize"] = o.MaxConnectionPoolSize
	}
	if !IsNil(o.SignerType) {
		toSerialize["signerType"] = o.SignerType
	}
	return toSerialize, nil
}

type NullableS3BlobStoreApiAdvancedBucketConnection struct {
	value *S3BlobStoreApiAdvancedBucketConnection
	isSet bool
}

func (v NullableS3BlobStoreApiAdvancedBucketConnection) Get() *S3BlobStoreApiAdvancedBucketConnection {
	return v.value
}

func (v *NullableS3BlobStoreApiAdvancedBucketConnection) Set(val *S3BlobStoreApiAdvancedBucketConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableS3BlobStoreApiAdvancedBucketConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableS3BlobStoreApiAdvancedBucketConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3BlobStoreApiAdvancedBucketConnection(val *S3BlobStoreApiAdvancedBucketConnection) *NullableS3BlobStoreApiAdvancedBucketConnection {
	return &NullableS3BlobStoreApiAdvancedBucketConnection{value: val, isSet: true}
}

func (v NullableS3BlobStoreApiAdvancedBucketConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3BlobStoreApiAdvancedBucketConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


