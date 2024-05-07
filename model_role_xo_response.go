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

// checks if the RoleXOResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleXOResponse{}

// RoleXOResponse struct for RoleXOResponse
type RoleXOResponse struct {
	// The description of this role.
	Description *string `json:"description,omitempty"`
	// The id of the role.
	Id *string `json:"id,omitempty"`
	// The name of the role.
	Name *string `json:"name,omitempty"`
	// The list of privileges assigned to this role.
	Privileges []string `json:"privileges,omitempty"`
	// Indicates whether the role can be changed. The system will ignore any supplied external values.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// The list of roles assigned to this role.
	Roles []string `json:"roles,omitempty"`
	// The user source which is the origin of this role.
	Source *string `json:"source,omitempty"`
}

// NewRoleXOResponse instantiates a new RoleXOResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleXOResponse() *RoleXOResponse {
	this := RoleXOResponse{}
	return &this
}

// NewRoleXOResponseWithDefaults instantiates a new RoleXOResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleXOResponseWithDefaults() *RoleXOResponse {
	this := RoleXOResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleXOResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXOResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleXOResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleXOResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleXOResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXOResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleXOResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleXOResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleXOResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXOResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleXOResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleXOResponse) SetName(v string) {
	o.Name = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *RoleXOResponse) GetPrivileges() []string {
	if o == nil || IsNil(o.Privileges) {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXOResponse) GetPrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *RoleXOResponse) HasPrivileges() bool {
	if o != nil && !IsNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *RoleXOResponse) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *RoleXOResponse) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXOResponse) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *RoleXOResponse) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *RoleXOResponse) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *RoleXOResponse) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXOResponse) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *RoleXOResponse) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *RoleXOResponse) SetRoles(v []string) {
	o.Roles = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RoleXOResponse) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleXOResponse) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RoleXOResponse) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RoleXOResponse) SetSource(v string) {
	o.Source = &v
}

func (o RoleXOResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleXOResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableRoleXOResponse struct {
	value *RoleXOResponse
	isSet bool
}

func (v NullableRoleXOResponse) Get() *RoleXOResponse {
	return v.value
}

func (v *NullableRoleXOResponse) Set(val *RoleXOResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleXOResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleXOResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleXOResponse(val *RoleXOResponse) *NullableRoleXOResponse {
	return &NullableRoleXOResponse{value: val, isSet: true}
}

func (v NullableRoleXOResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleXOResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


