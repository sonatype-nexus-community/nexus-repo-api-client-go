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

// checks if the ContentSelectorApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSelectorApiResponse{}

// ContentSelectorApiResponse struct for ContentSelectorApiResponse
type ContentSelectorApiResponse struct {
	// A human-readable description
	Description *string `json:"description,omitempty"`
	// The expression used to identify content
	Expression *string `json:"expression,omitempty"`
	// The content selector name cannot be changed after creation
	Name *string `json:"name,omitempty"`
	// The type of content selector the backend is using
	Type *string `json:"type,omitempty"`
}

// NewContentSelectorApiResponse instantiates a new ContentSelectorApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSelectorApiResponse() *ContentSelectorApiResponse {
	this := ContentSelectorApiResponse{}
	return &this
}

// NewContentSelectorApiResponseWithDefaults instantiates a new ContentSelectorApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSelectorApiResponseWithDefaults() *ContentSelectorApiResponse {
	this := ContentSelectorApiResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContentSelectorApiResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSelectorApiResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContentSelectorApiResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContentSelectorApiResponse) SetDescription(v string) {
	o.Description = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *ContentSelectorApiResponse) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSelectorApiResponse) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *ContentSelectorApiResponse) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *ContentSelectorApiResponse) SetExpression(v string) {
	o.Expression = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContentSelectorApiResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSelectorApiResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContentSelectorApiResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContentSelectorApiResponse) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContentSelectorApiResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSelectorApiResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContentSelectorApiResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContentSelectorApiResponse) SetType(v string) {
	o.Type = &v
}

func (o ContentSelectorApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentSelectorApiResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableContentSelectorApiResponse struct {
	value *ContentSelectorApiResponse
	isSet bool
}

func (v NullableContentSelectorApiResponse) Get() *ContentSelectorApiResponse {
	return v.value
}

func (v *NullableContentSelectorApiResponse) Set(val *ContentSelectorApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSelectorApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSelectorApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSelectorApiResponse(val *ContentSelectorApiResponse) *NullableContentSelectorApiResponse {
	return &NullableContentSelectorApiResponse{value: val, isSet: true}
}

func (v NullableContentSelectorApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSelectorApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


