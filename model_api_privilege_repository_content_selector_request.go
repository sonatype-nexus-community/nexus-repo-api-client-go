/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.82.0-08.

API version: 3.82.0-08
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the ApiPrivilegeRepositoryContentSelectorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPrivilegeRepositoryContentSelectorRequest{}

// ApiPrivilegeRepositoryContentSelectorRequest struct for ApiPrivilegeRepositoryContentSelectorRequest
type ApiPrivilegeRepositoryContentSelectorRequest struct {
	// A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges.
	Actions []string `json:"actions,omitempty"`
	// The name of a content selector that will be used to grant access to content via this privilege.
	ContentSelector *string `json:"contentSelector,omitempty"`
	Description *string `json:"description,omitempty"`
	// The repository format (i.e 'nuget', 'npm') this privilege will grant access to (or * for all).
	Format *string `json:"format,omitempty"`
	// The name of the privilege.  This value cannot be changed.
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// The name of the repository this privilege will grant access to (or * for all).
	Repository *string `json:"repository,omitempty"`
}

// NewApiPrivilegeRepositoryContentSelectorRequest instantiates a new ApiPrivilegeRepositoryContentSelectorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPrivilegeRepositoryContentSelectorRequest() *ApiPrivilegeRepositoryContentSelectorRequest {
	this := ApiPrivilegeRepositoryContentSelectorRequest{}
	return &this
}

// NewApiPrivilegeRepositoryContentSelectorRequestWithDefaults instantiates a new ApiPrivilegeRepositoryContentSelectorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPrivilegeRepositoryContentSelectorRequestWithDefaults() *ApiPrivilegeRepositoryContentSelectorRequest {
	this := ApiPrivilegeRepositoryContentSelectorRequest{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetActions() []string {
	if o == nil || IsNil(o.Actions) {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetActions(v []string) {
	o.Actions = v
}

// GetContentSelector returns the ContentSelector field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetContentSelector() string {
	if o == nil || IsNil(o.ContentSelector) {
		var ret string
		return ret
	}
	return *o.ContentSelector
}

// GetContentSelectorOk returns a tuple with the ContentSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetContentSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.ContentSelector) {
		return nil, false
	}
	return o.ContentSelector, true
}

// HasContentSelector returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasContentSelector() bool {
	if o != nil && !IsNil(o.ContentSelector) {
		return true
	}

	return false
}

// SetContentSelector gets a reference to the given string and assigns it to the ContentSelector field.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetContentSelector(v string) {
	o.ContentSelector = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetFormat(v string) {
	o.Format = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetName(v string) {
	o.Name = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetRepository(v string) {
	o.Repository = &v
}

func (o ApiPrivilegeRepositoryContentSelectorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPrivilegeRepositoryContentSelectorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.ContentSelector) {
		toSerialize["contentSelector"] = o.ContentSelector
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

type NullableApiPrivilegeRepositoryContentSelectorRequest struct {
	value *ApiPrivilegeRepositoryContentSelectorRequest
	isSet bool
}

func (v NullableApiPrivilegeRepositoryContentSelectorRequest) Get() *ApiPrivilegeRepositoryContentSelectorRequest {
	return v.value
}

func (v *NullableApiPrivilegeRepositoryContentSelectorRequest) Set(val *ApiPrivilegeRepositoryContentSelectorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPrivilegeRepositoryContentSelectorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPrivilegeRepositoryContentSelectorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPrivilegeRepositoryContentSelectorRequest(val *ApiPrivilegeRepositoryContentSelectorRequest) *NullableApiPrivilegeRepositoryContentSelectorRequest {
	return &NullableApiPrivilegeRepositoryContentSelectorRequest{value: val, isSet: true}
}

func (v NullableApiPrivilegeRepositoryContentSelectorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPrivilegeRepositoryContentSelectorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


