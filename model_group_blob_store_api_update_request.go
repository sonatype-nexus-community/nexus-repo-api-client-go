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

// checks if the GroupBlobStoreApiUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupBlobStoreApiUpdateRequest{}

// GroupBlobStoreApiUpdateRequest struct for GroupBlobStoreApiUpdateRequest
type GroupBlobStoreApiUpdateRequest struct {
	FillPolicy *string `json:"fillPolicy,omitempty"`
	// List of the names of blob stores that are members of this group
	Members []string `json:"members,omitempty"`
	SoftQuota *BlobStoreApiSoftQuota `json:"softQuota,omitempty"`
}

// NewGroupBlobStoreApiUpdateRequest instantiates a new GroupBlobStoreApiUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupBlobStoreApiUpdateRequest() *GroupBlobStoreApiUpdateRequest {
	this := GroupBlobStoreApiUpdateRequest{}
	return &this
}

// NewGroupBlobStoreApiUpdateRequestWithDefaults instantiates a new GroupBlobStoreApiUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupBlobStoreApiUpdateRequestWithDefaults() *GroupBlobStoreApiUpdateRequest {
	this := GroupBlobStoreApiUpdateRequest{}
	return &this
}

// GetFillPolicy returns the FillPolicy field value if set, zero value otherwise.
func (o *GroupBlobStoreApiUpdateRequest) GetFillPolicy() string {
	if o == nil || IsNil(o.FillPolicy) {
		var ret string
		return ret
	}
	return *o.FillPolicy
}

// GetFillPolicyOk returns a tuple with the FillPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBlobStoreApiUpdateRequest) GetFillPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.FillPolicy) {
		return nil, false
	}
	return o.FillPolicy, true
}

// HasFillPolicy returns a boolean if a field has been set.
func (o *GroupBlobStoreApiUpdateRequest) HasFillPolicy() bool {
	if o != nil && !IsNil(o.FillPolicy) {
		return true
	}

	return false
}

// SetFillPolicy gets a reference to the given string and assigns it to the FillPolicy field.
func (o *GroupBlobStoreApiUpdateRequest) SetFillPolicy(v string) {
	o.FillPolicy = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *GroupBlobStoreApiUpdateRequest) GetMembers() []string {
	if o == nil || IsNil(o.Members) {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBlobStoreApiUpdateRequest) GetMembersOk() ([]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *GroupBlobStoreApiUpdateRequest) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *GroupBlobStoreApiUpdateRequest) SetMembers(v []string) {
	o.Members = v
}

// GetSoftQuota returns the SoftQuota field value if set, zero value otherwise.
func (o *GroupBlobStoreApiUpdateRequest) GetSoftQuota() BlobStoreApiSoftQuota {
	if o == nil || IsNil(o.SoftQuota) {
		var ret BlobStoreApiSoftQuota
		return ret
	}
	return *o.SoftQuota
}

// GetSoftQuotaOk returns a tuple with the SoftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBlobStoreApiUpdateRequest) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool) {
	if o == nil || IsNil(o.SoftQuota) {
		return nil, false
	}
	return o.SoftQuota, true
}

// HasSoftQuota returns a boolean if a field has been set.
func (o *GroupBlobStoreApiUpdateRequest) HasSoftQuota() bool {
	if o != nil && !IsNil(o.SoftQuota) {
		return true
	}

	return false
}

// SetSoftQuota gets a reference to the given BlobStoreApiSoftQuota and assigns it to the SoftQuota field.
func (o *GroupBlobStoreApiUpdateRequest) SetSoftQuota(v BlobStoreApiSoftQuota) {
	o.SoftQuota = &v
}

func (o GroupBlobStoreApiUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupBlobStoreApiUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FillPolicy) {
		toSerialize["fillPolicy"] = o.FillPolicy
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.SoftQuota) {
		toSerialize["softQuota"] = o.SoftQuota
	}
	return toSerialize, nil
}

type NullableGroupBlobStoreApiUpdateRequest struct {
	value *GroupBlobStoreApiUpdateRequest
	isSet bool
}

func (v NullableGroupBlobStoreApiUpdateRequest) Get() *GroupBlobStoreApiUpdateRequest {
	return v.value
}

func (v *NullableGroupBlobStoreApiUpdateRequest) Set(val *GroupBlobStoreApiUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupBlobStoreApiUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupBlobStoreApiUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupBlobStoreApiUpdateRequest(val *GroupBlobStoreApiUpdateRequest) *NullableGroupBlobStoreApiUpdateRequest {
	return &NullableGroupBlobStoreApiUpdateRequest{value: val, isSet: true}
}

func (v NullableGroupBlobStoreApiUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupBlobStoreApiUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


