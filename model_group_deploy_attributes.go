/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.67.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
)

// checks if the GroupDeployAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupDeployAttributes{}

// GroupDeployAttributes struct for GroupDeployAttributes
type GroupDeployAttributes struct {
	// Member repositories' names
	MemberNames []string `json:"memberNames,omitempty"`
	// Pro-only: This field is for the Group Deployment feature available in NXRM Pro.
	WritableMember *string `json:"writableMember,omitempty"`
}

// NewGroupDeployAttributes instantiates a new GroupDeployAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDeployAttributes() *GroupDeployAttributes {
	this := GroupDeployAttributes{}
	return &this
}

// NewGroupDeployAttributesWithDefaults instantiates a new GroupDeployAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDeployAttributesWithDefaults() *GroupDeployAttributes {
	this := GroupDeployAttributes{}
	return &this
}

// GetMemberNames returns the MemberNames field value if set, zero value otherwise.
func (o *GroupDeployAttributes) GetMemberNames() []string {
	if o == nil || IsNil(o.MemberNames) {
		var ret []string
		return ret
	}
	return o.MemberNames
}

// GetMemberNamesOk returns a tuple with the MemberNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDeployAttributes) GetMemberNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.MemberNames) {
		return nil, false
	}
	return o.MemberNames, true
}

// HasMemberNames returns a boolean if a field has been set.
func (o *GroupDeployAttributes) HasMemberNames() bool {
	if o != nil && !IsNil(o.MemberNames) {
		return true
	}

	return false
}

// SetMemberNames gets a reference to the given []string and assigns it to the MemberNames field.
func (o *GroupDeployAttributes) SetMemberNames(v []string) {
	o.MemberNames = v
}

// GetWritableMember returns the WritableMember field value if set, zero value otherwise.
func (o *GroupDeployAttributes) GetWritableMember() string {
	if o == nil || IsNil(o.WritableMember) {
		var ret string
		return ret
	}
	return *o.WritableMember
}

// GetWritableMemberOk returns a tuple with the WritableMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDeployAttributes) GetWritableMemberOk() (*string, bool) {
	if o == nil || IsNil(o.WritableMember) {
		return nil, false
	}
	return o.WritableMember, true
}

// HasWritableMember returns a boolean if a field has been set.
func (o *GroupDeployAttributes) HasWritableMember() bool {
	if o != nil && !IsNil(o.WritableMember) {
		return true
	}

	return false
}

// SetWritableMember gets a reference to the given string and assigns it to the WritableMember field.
func (o *GroupDeployAttributes) SetWritableMember(v string) {
	o.WritableMember = &v
}

func (o GroupDeployAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupDeployAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberNames) {
		toSerialize["memberNames"] = o.MemberNames
	}
	if !IsNil(o.WritableMember) {
		toSerialize["writableMember"] = o.WritableMember
	}
	return toSerialize, nil
}

type NullableGroupDeployAttributes struct {
	value *GroupDeployAttributes
	isSet bool
}

func (v NullableGroupDeployAttributes) Get() *GroupDeployAttributes {
	return v.value
}

func (v *NullableGroupDeployAttributes) Set(val *GroupDeployAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupDeployAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupDeployAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupDeployAttributes(val *GroupDeployAttributes) *NullableGroupDeployAttributes {
	return &NullableGroupDeployAttributes{value: val, isSet: true}
}

func (v NullableGroupDeployAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupDeployAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


