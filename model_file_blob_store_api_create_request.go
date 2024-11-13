/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.74.0-05.

API version: 3.74.0-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
)

// checks if the FileBlobStoreApiCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileBlobStoreApiCreateRequest{}

// FileBlobStoreApiCreateRequest struct for FileBlobStoreApiCreateRequest
type FileBlobStoreApiCreateRequest struct {
	Name *string `json:"name,omitempty"`
	// The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to or it can be a path relative to the sonatype-work directory.
	Path *string `json:"path,omitempty"`
	SoftQuota *BlobStoreApiSoftQuota `json:"softQuota,omitempty"`
}

// NewFileBlobStoreApiCreateRequest instantiates a new FileBlobStoreApiCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileBlobStoreApiCreateRequest() *FileBlobStoreApiCreateRequest {
	this := FileBlobStoreApiCreateRequest{}
	return &this
}

// NewFileBlobStoreApiCreateRequestWithDefaults instantiates a new FileBlobStoreApiCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileBlobStoreApiCreateRequestWithDefaults() *FileBlobStoreApiCreateRequest {
	this := FileBlobStoreApiCreateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FileBlobStoreApiCreateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBlobStoreApiCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FileBlobStoreApiCreateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FileBlobStoreApiCreateRequest) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FileBlobStoreApiCreateRequest) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBlobStoreApiCreateRequest) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FileBlobStoreApiCreateRequest) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FileBlobStoreApiCreateRequest) SetPath(v string) {
	o.Path = &v
}

// GetSoftQuota returns the SoftQuota field value if set, zero value otherwise.
func (o *FileBlobStoreApiCreateRequest) GetSoftQuota() BlobStoreApiSoftQuota {
	if o == nil || IsNil(o.SoftQuota) {
		var ret BlobStoreApiSoftQuota
		return ret
	}
	return *o.SoftQuota
}

// GetSoftQuotaOk returns a tuple with the SoftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBlobStoreApiCreateRequest) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool) {
	if o == nil || IsNil(o.SoftQuota) {
		return nil, false
	}
	return o.SoftQuota, true
}

// HasSoftQuota returns a boolean if a field has been set.
func (o *FileBlobStoreApiCreateRequest) HasSoftQuota() bool {
	if o != nil && !IsNil(o.SoftQuota) {
		return true
	}

	return false
}

// SetSoftQuota gets a reference to the given BlobStoreApiSoftQuota and assigns it to the SoftQuota field.
func (o *FileBlobStoreApiCreateRequest) SetSoftQuota(v BlobStoreApiSoftQuota) {
	o.SoftQuota = &v
}

func (o FileBlobStoreApiCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileBlobStoreApiCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.SoftQuota) {
		toSerialize["softQuota"] = o.SoftQuota
	}
	return toSerialize, nil
}

type NullableFileBlobStoreApiCreateRequest struct {
	value *FileBlobStoreApiCreateRequest
	isSet bool
}

func (v NullableFileBlobStoreApiCreateRequest) Get() *FileBlobStoreApiCreateRequest {
	return v.value
}

func (v *NullableFileBlobStoreApiCreateRequest) Set(val *FileBlobStoreApiCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileBlobStoreApiCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileBlobStoreApiCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileBlobStoreApiCreateRequest(val *FileBlobStoreApiCreateRequest) *NullableFileBlobStoreApiCreateRequest {
	return &NullableFileBlobStoreApiCreateRequest{value: val, isSet: true}
}

func (v NullableFileBlobStoreApiCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileBlobStoreApiCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


