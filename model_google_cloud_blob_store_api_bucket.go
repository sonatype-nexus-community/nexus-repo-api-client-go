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

// checks if the GoogleCloudBlobStoreApiBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleCloudBlobStoreApiBucket{}

// GoogleCloudBlobStoreApiBucket struct for GoogleCloudBlobStoreApiBucket
type GoogleCloudBlobStoreApiBucket struct {
	// The name of the GC Storage bucket
	Name string `json:"name"`
	// The GC Storage blob store (i.e GC Storage object) key prefix
	Prefix *string `json:"prefix,omitempty"`
	// GCP Project ID
	ProjectId *string `json:"projectId,omitempty"`
	// The GCP region to create a new GC Storage bucket in or an existing GC Storage bucket's region
	Region string `json:"region"`
}

type _GoogleCloudBlobStoreApiBucket GoogleCloudBlobStoreApiBucket

// NewGoogleCloudBlobStoreApiBucket instantiates a new GoogleCloudBlobStoreApiBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudBlobStoreApiBucket(name string, region string) *GoogleCloudBlobStoreApiBucket {
	this := GoogleCloudBlobStoreApiBucket{}
	this.Name = name
	this.Region = region
	return &this
}

// NewGoogleCloudBlobStoreApiBucketWithDefaults instantiates a new GoogleCloudBlobStoreApiBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudBlobStoreApiBucketWithDefaults() *GoogleCloudBlobStoreApiBucket {
	this := GoogleCloudBlobStoreApiBucket{}
	return &this
}

// GetName returns the Name field value
func (o *GoogleCloudBlobStoreApiBucket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudBlobStoreApiBucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GoogleCloudBlobStoreApiBucket) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *GoogleCloudBlobStoreApiBucket) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudBlobStoreApiBucket) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *GoogleCloudBlobStoreApiBucket) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *GoogleCloudBlobStoreApiBucket) SetPrefix(v string) {
	o.Prefix = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GoogleCloudBlobStoreApiBucket) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudBlobStoreApiBucket) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GoogleCloudBlobStoreApiBucket) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GoogleCloudBlobStoreApiBucket) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRegion returns the Region field value
func (o *GoogleCloudBlobStoreApiBucket) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudBlobStoreApiBucket) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *GoogleCloudBlobStoreApiBucket) SetRegion(v string) {
	o.Region = v
}

func (o GoogleCloudBlobStoreApiBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleCloudBlobStoreApiBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	toSerialize["region"] = o.Region
	return toSerialize, nil
}

func (o *GoogleCloudBlobStoreApiBucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varGoogleCloudBlobStoreApiBucket := _GoogleCloudBlobStoreApiBucket{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGoogleCloudBlobStoreApiBucket)

	if err != nil {
		return err
	}

	*o = GoogleCloudBlobStoreApiBucket(varGoogleCloudBlobStoreApiBucket)

	return err
}

type NullableGoogleCloudBlobStoreApiBucket struct {
	value *GoogleCloudBlobStoreApiBucket
	isSet bool
}

func (v NullableGoogleCloudBlobStoreApiBucket) Get() *GoogleCloudBlobStoreApiBucket {
	return v.value
}

func (v *NullableGoogleCloudBlobStoreApiBucket) Set(val *GoogleCloudBlobStoreApiBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudBlobStoreApiBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudBlobStoreApiBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudBlobStoreApiBucket(val *GoogleCloudBlobStoreApiBucket) *NullableGoogleCloudBlobStoreApiBucket {
	return &NullableGoogleCloudBlobStoreApiBucket{value: val, isSet: true}
}

func (v NullableGoogleCloudBlobStoreApiBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudBlobStoreApiBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

