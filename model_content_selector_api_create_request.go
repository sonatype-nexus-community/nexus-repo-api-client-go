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

// checks if the ContentSelectorApiCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSelectorApiCreateRequest{}

// ContentSelectorApiCreateRequest struct for ContentSelectorApiCreateRequest
type ContentSelectorApiCreateRequest struct {
	// A human-readable description
	Description *string `json:"description,omitempty"`
	// The expression used to identify content
	Expression *string `json:"expression,omitempty"`
	// The content selector name cannot be changed after creation
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
}

// NewContentSelectorApiCreateRequest instantiates a new ContentSelectorApiCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSelectorApiCreateRequest() *ContentSelectorApiCreateRequest {
	this := ContentSelectorApiCreateRequest{}
	return &this
}

// NewContentSelectorApiCreateRequestWithDefaults instantiates a new ContentSelectorApiCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSelectorApiCreateRequestWithDefaults() *ContentSelectorApiCreateRequest {
	this := ContentSelectorApiCreateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContentSelectorApiCreateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSelectorApiCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContentSelectorApiCreateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContentSelectorApiCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *ContentSelectorApiCreateRequest) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSelectorApiCreateRequest) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *ContentSelectorApiCreateRequest) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *ContentSelectorApiCreateRequest) SetExpression(v string) {
	o.Expression = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContentSelectorApiCreateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSelectorApiCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContentSelectorApiCreateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContentSelectorApiCreateRequest) SetName(v string) {
	o.Name = &v
}

func (o ContentSelectorApiCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentSelectorApiCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableContentSelectorApiCreateRequest struct {
	value *ContentSelectorApiCreateRequest
	isSet bool
}

func (v NullableContentSelectorApiCreateRequest) Get() *ContentSelectorApiCreateRequest {
	return v.value
}

func (v *NullableContentSelectorApiCreateRequest) Set(val *ContentSelectorApiCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSelectorApiCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSelectorApiCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSelectorApiCreateRequest(val *ContentSelectorApiCreateRequest) *NullableContentSelectorApiCreateRequest {
	return &NullableContentSelectorApiCreateRequest{value: val, isSet: true}
}

func (v NullableContentSelectorApiCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSelectorApiCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


