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

// checks if the ApiPrivilegeRepositoryAdminRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPrivilegeRepositoryAdminRequest{}

// ApiPrivilegeRepositoryAdminRequest struct for ApiPrivilegeRepositoryAdminRequest
type ApiPrivilegeRepositoryAdminRequest struct {
	// A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges.
	Actions []string `json:"actions,omitempty"`
	Description *string `json:"description,omitempty"`
	// The repository format (i.e 'nuget', 'npm') this privilege will grant access to (or * for all).
	Format *string `json:"format,omitempty"`
	// The name of the privilege.  This value cannot be changed.
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// The name of the repository this privilege will grant access to (or * for all).
	Repository *string `json:"repository,omitempty"`
}

// NewApiPrivilegeRepositoryAdminRequest instantiates a new ApiPrivilegeRepositoryAdminRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPrivilegeRepositoryAdminRequest() *ApiPrivilegeRepositoryAdminRequest {
	this := ApiPrivilegeRepositoryAdminRequest{}
	return &this
}

// NewApiPrivilegeRepositoryAdminRequestWithDefaults instantiates a new ApiPrivilegeRepositoryAdminRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPrivilegeRepositoryAdminRequestWithDefaults() *ApiPrivilegeRepositoryAdminRequest {
	this := ApiPrivilegeRepositoryAdminRequest{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryAdminRequest) GetActions() []string {
	if o == nil || IsNil(o.Actions) {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) GetActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *ApiPrivilegeRepositoryAdminRequest) SetActions(v []string) {
	o.Actions = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryAdminRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiPrivilegeRepositoryAdminRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryAdminRequest) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ApiPrivilegeRepositoryAdminRequest) SetFormat(v string) {
	o.Format = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryAdminRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiPrivilegeRepositoryAdminRequest) SetName(v string) {
	o.Name = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryAdminRequest) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryAdminRequest) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *ApiPrivilegeRepositoryAdminRequest) SetRepository(v string) {
	o.Repository = &v
}

func (o ApiPrivilegeRepositoryAdminRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPrivilegeRepositoryAdminRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

type NullableApiPrivilegeRepositoryAdminRequest struct {
	value *ApiPrivilegeRepositoryAdminRequest
	isSet bool
}

func (v NullableApiPrivilegeRepositoryAdminRequest) Get() *ApiPrivilegeRepositoryAdminRequest {
	return v.value
}

func (v *NullableApiPrivilegeRepositoryAdminRequest) Set(val *ApiPrivilegeRepositoryAdminRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPrivilegeRepositoryAdminRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPrivilegeRepositoryAdminRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPrivilegeRepositoryAdminRequest(val *ApiPrivilegeRepositoryAdminRequest) *NullableApiPrivilegeRepositoryAdminRequest {
	return &NullableApiPrivilegeRepositoryAdminRequest{value: val, isSet: true}
}

func (v NullableApiPrivilegeRepositoryAdminRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPrivilegeRepositoryAdminRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


