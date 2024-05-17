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

// checks if the RoleXORequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleXORequest{}

// RoleXORequest struct for RoleXORequest
type RoleXORequest struct {
	// The description of this role.
	Description *string `json:"description,omitempty"`
	// The id of the role.
	Id *string `json:"id,omitempty"`
	// The name of the role.
	Name *string `json:"name,omitempty"`
	// The list of privileges assigned to this role.
	Privileges []string `json:"privileges,omitempty"`
	// The list of roles assigned to this role.
	Roles []string `json:"roles,omitempty"`
}

// NewRoleXORequest instantiates a new RoleXORequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleXORequest() *RoleXORequest {
	this := RoleXORequest{}
	return &this
}

// NewRoleXORequestWithDefaults instantiates a new RoleXORequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleXORequestWithDefaults() *RoleXORequest {
	this := RoleXORequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleXORequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXORequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleXORequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleXORequest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleXORequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXORequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleXORequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleXORequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleXORequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXORequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleXORequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleXORequest) SetName(v string) {
	o.Name = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *RoleXORequest) GetPrivileges() []string {
	if o == nil || IsNil(o.Privileges) {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXORequest) GetPrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *RoleXORequest) HasPrivileges() bool {
	if o != nil && !IsNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *RoleXORequest) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *RoleXORequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXORequest) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *RoleXORequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *RoleXORequest) SetRoles(v []string) {
	o.Roles = v
}

func (o RoleXORequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleXORequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableRoleXORequest struct {
	value *RoleXORequest
	isSet bool
}

func (v NullableRoleXORequest) Get() *RoleXORequest {
	return v.value
}

func (v *NullableRoleXORequest) Set(val *RoleXORequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleXORequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleXORequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleXORequest(val *RoleXORequest) *NullableRoleXORequest {
	return &NullableRoleXORequest{value: val, isSet: true}
}

func (v NullableRoleXORequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleXORequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


