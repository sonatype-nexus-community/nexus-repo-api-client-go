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

// checks if the ApiPrivilegeApplicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPrivilegeApplicationRequest{}

// ApiPrivilegeApplicationRequest struct for ApiPrivilegeApplicationRequest
type ApiPrivilegeApplicationRequest struct {
	// A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges.
	Actions []string `json:"actions,omitempty"`
	Description *string `json:"description,omitempty"`
	// The domain (i.e. 'blobstores', 'capabilities' or even '*' for all) that this privilege is granting access to.  Note that creating new privileges with a domain is only necessary when using plugins that define their own domain(s).
	Domain *string `json:"domain,omitempty"`
	// The name of the privilege.  This value cannot be changed.
	Name *string `json:"name,omitempty"`
}

// NewApiPrivilegeApplicationRequest instantiates a new ApiPrivilegeApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPrivilegeApplicationRequest() *ApiPrivilegeApplicationRequest {
	this := ApiPrivilegeApplicationRequest{}
	return &this
}

// NewApiPrivilegeApplicationRequestWithDefaults instantiates a new ApiPrivilegeApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPrivilegeApplicationRequestWithDefaults() *ApiPrivilegeApplicationRequest {
	this := ApiPrivilegeApplicationRequest{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ApiPrivilegeApplicationRequest) GetActions() []string {
	if o == nil || IsNil(o.Actions) {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeApplicationRequest) GetActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ApiPrivilegeApplicationRequest) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *ApiPrivilegeApplicationRequest) SetActions(v []string) {
	o.Actions = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiPrivilegeApplicationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeApplicationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiPrivilegeApplicationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiPrivilegeApplicationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiPrivilegeApplicationRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeApplicationRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiPrivilegeApplicationRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiPrivilegeApplicationRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiPrivilegeApplicationRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiPrivilegeApplicationRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiPrivilegeApplicationRequest) SetName(v string) {
	o.Name = &v
}

func (o ApiPrivilegeApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPrivilegeApplicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableApiPrivilegeApplicationRequest struct {
	value *ApiPrivilegeApplicationRequest
	isSet bool
}

func (v NullableApiPrivilegeApplicationRequest) Get() *ApiPrivilegeApplicationRequest {
	return v.value
}

func (v *NullableApiPrivilegeApplicationRequest) Set(val *ApiPrivilegeApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPrivilegeApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPrivilegeApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPrivilegeApplicationRequest(val *ApiPrivilegeApplicationRequest) *NullableApiPrivilegeApplicationRequest {
	return &NullableApiPrivilegeApplicationRequest{value: val, isSet: true}
}

func (v NullableApiPrivilegeApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPrivilegeApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


