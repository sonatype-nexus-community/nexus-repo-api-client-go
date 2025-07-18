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

// checks if the RepositoryXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryXO{}

// RepositoryXO struct for RepositoryXO
type RepositoryXO struct {
	Attributes map[string]map[string]interface{} `json:"attributes,omitempty"`
	Format *string `json:"format,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewRepositoryXO instantiates a new RepositoryXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryXO() *RepositoryXO {
	this := RepositoryXO{}
	return &this
}

// NewRepositoryXOWithDefaults instantiates a new RepositoryXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryXOWithDefaults() *RepositoryXO {
	this := RepositoryXO{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RepositoryXO) GetAttributes() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryXO) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RepositoryXO) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *RepositoryXO) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *RepositoryXO) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryXO) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *RepositoryXO) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *RepositoryXO) SetFormat(v string) {
	o.Format = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RepositoryXO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryXO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RepositoryXO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RepositoryXO) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RepositoryXO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryXO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RepositoryXO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RepositoryXO) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RepositoryXO) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryXO) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RepositoryXO) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RepositoryXO) SetUrl(v string) {
	o.Url = &v
}

func (o RepositoryXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableRepositoryXO struct {
	value *RepositoryXO
	isSet bool
}

func (v NullableRepositoryXO) Get() *RepositoryXO {
	return v.value
}

func (v *NullableRepositoryXO) Set(val *RepositoryXO) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryXO) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryXO(val *RepositoryXO) *NullableRepositoryXO {
	return &NullableRepositoryXO{value: val, isSet: true}
}

func (v NullableRepositoryXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


