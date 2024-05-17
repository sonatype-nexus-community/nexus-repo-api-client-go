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

// checks if the S3BlobStoreApiBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3BlobStoreApiBucket{}

// S3BlobStoreApiBucket struct for S3BlobStoreApiBucket
type S3BlobStoreApiBucket struct {
	// How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable)
	Expiration int32 `json:"expiration"`
	// The name of the S3 bucket
	Name string `json:"name"`
	// The S3 blob store (i.e S3 object) key prefix
	Prefix *string `json:"prefix,omitempty"`
	// The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	Region string `json:"region"`
}

type _S3BlobStoreApiBucket S3BlobStoreApiBucket

// NewS3BlobStoreApiBucket instantiates a new S3BlobStoreApiBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3BlobStoreApiBucket(expiration int32, name string, region string) *S3BlobStoreApiBucket {
	this := S3BlobStoreApiBucket{}
	this.Expiration = expiration
	this.Name = name
	this.Region = region
	return &this
}

// NewS3BlobStoreApiBucketWithDefaults instantiates a new S3BlobStoreApiBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3BlobStoreApiBucketWithDefaults() *S3BlobStoreApiBucket {
	this := S3BlobStoreApiBucket{}
	return &this
}

// GetExpiration returns the Expiration field value
func (o *S3BlobStoreApiBucket) GetExpiration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiBucket) GetExpirationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *S3BlobStoreApiBucket) SetExpiration(v int32) {
	o.Expiration = v
}

// GetName returns the Name field value
func (o *S3BlobStoreApiBucket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiBucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *S3BlobStoreApiBucket) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *S3BlobStoreApiBucket) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiBucket) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *S3BlobStoreApiBucket) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *S3BlobStoreApiBucket) SetPrefix(v string) {
	o.Prefix = &v
}

// GetRegion returns the Region field value
func (o *S3BlobStoreApiBucket) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *S3BlobStoreApiBucket) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *S3BlobStoreApiBucket) SetRegion(v string) {
	o.Region = v
}

func (o S3BlobStoreApiBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3BlobStoreApiBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expiration"] = o.Expiration
	toSerialize["name"] = o.Name
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	toSerialize["region"] = o.Region
	return toSerialize, nil
}

func (o *S3BlobStoreApiBucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expiration",
		"name",
		"region",
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

	varS3BlobStoreApiBucket := _S3BlobStoreApiBucket{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varS3BlobStoreApiBucket)

	if err != nil {
		return err
	}

	*o = S3BlobStoreApiBucket(varS3BlobStoreApiBucket)

	return err
}

type NullableS3BlobStoreApiBucket struct {
	value *S3BlobStoreApiBucket
	isSet bool
}

func (v NullableS3BlobStoreApiBucket) Get() *S3BlobStoreApiBucket {
	return v.value
}

func (v *NullableS3BlobStoreApiBucket) Set(val *S3BlobStoreApiBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableS3BlobStoreApiBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableS3BlobStoreApiBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3BlobStoreApiBucket(val *S3BlobStoreApiBucket) *NullableS3BlobStoreApiBucket {
	return &NullableS3BlobStoreApiBucket{value: val, isSet: true}
}

func (v NullableS3BlobStoreApiBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3BlobStoreApiBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


