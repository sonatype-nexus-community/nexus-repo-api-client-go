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

// checks if the UploadDefinitionXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadDefinitionXO{}

// UploadDefinitionXO struct for UploadDefinitionXO
type UploadDefinitionXO struct {
	AssetFields []UploadFieldDefinitionXO `json:"assetFields,omitempty"`
	ComponentFields []UploadFieldDefinitionXO `json:"componentFields,omitempty"`
	Format *string `json:"format,omitempty"`
	MultipleUpload *bool `json:"multipleUpload,omitempty"`
}

// NewUploadDefinitionXO instantiates a new UploadDefinitionXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadDefinitionXO() *UploadDefinitionXO {
	this := UploadDefinitionXO{}
	return &this
}

// NewUploadDefinitionXOWithDefaults instantiates a new UploadDefinitionXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadDefinitionXOWithDefaults() *UploadDefinitionXO {
	this := UploadDefinitionXO{}
	return &this
}

// GetAssetFields returns the AssetFields field value if set, zero value otherwise.
func (o *UploadDefinitionXO) GetAssetFields() []UploadFieldDefinitionXO {
	if o == nil || IsNil(o.AssetFields) {
		var ret []UploadFieldDefinitionXO
		return ret
	}
	return o.AssetFields
}

// GetAssetFieldsOk returns a tuple with the AssetFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadDefinitionXO) GetAssetFieldsOk() ([]UploadFieldDefinitionXO, bool) {
	if o == nil || IsNil(o.AssetFields) {
		return nil, false
	}
	return o.AssetFields, true
}

// HasAssetFields returns a boolean if a field has been set.
func (o *UploadDefinitionXO) HasAssetFields() bool {
	if o != nil && !IsNil(o.AssetFields) {
		return true
	}

	return false
}

// SetAssetFields gets a reference to the given []UploadFieldDefinitionXO and assigns it to the AssetFields field.
func (o *UploadDefinitionXO) SetAssetFields(v []UploadFieldDefinitionXO) {
	o.AssetFields = v
}

// GetComponentFields returns the ComponentFields field value if set, zero value otherwise.
func (o *UploadDefinitionXO) GetComponentFields() []UploadFieldDefinitionXO {
	if o == nil || IsNil(o.ComponentFields) {
		var ret []UploadFieldDefinitionXO
		return ret
	}
	return o.ComponentFields
}

// GetComponentFieldsOk returns a tuple with the ComponentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadDefinitionXO) GetComponentFieldsOk() ([]UploadFieldDefinitionXO, bool) {
	if o == nil || IsNil(o.ComponentFields) {
		return nil, false
	}
	return o.ComponentFields, true
}

// HasComponentFields returns a boolean if a field has been set.
func (o *UploadDefinitionXO) HasComponentFields() bool {
	if o != nil && !IsNil(o.ComponentFields) {
		return true
	}

	return false
}

// SetComponentFields gets a reference to the given []UploadFieldDefinitionXO and assigns it to the ComponentFields field.
func (o *UploadDefinitionXO) SetComponentFields(v []UploadFieldDefinitionXO) {
	o.ComponentFields = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *UploadDefinitionXO) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadDefinitionXO) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *UploadDefinitionXO) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *UploadDefinitionXO) SetFormat(v string) {
	o.Format = &v
}

// GetMultipleUpload returns the MultipleUpload field value if set, zero value otherwise.
func (o *UploadDefinitionXO) GetMultipleUpload() bool {
	if o == nil || IsNil(o.MultipleUpload) {
		var ret bool
		return ret
	}
	return *o.MultipleUpload
}

// GetMultipleUploadOk returns a tuple with the MultipleUpload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadDefinitionXO) GetMultipleUploadOk() (*bool, bool) {
	if o == nil || IsNil(o.MultipleUpload) {
		return nil, false
	}
	return o.MultipleUpload, true
}

// HasMultipleUpload returns a boolean if a field has been set.
func (o *UploadDefinitionXO) HasMultipleUpload() bool {
	if o != nil && !IsNil(o.MultipleUpload) {
		return true
	}

	return false
}

// SetMultipleUpload gets a reference to the given bool and assigns it to the MultipleUpload field.
func (o *UploadDefinitionXO) SetMultipleUpload(v bool) {
	o.MultipleUpload = &v
}

func (o UploadDefinitionXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadDefinitionXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetFields) {
		toSerialize["assetFields"] = o.AssetFields
	}
	if !IsNil(o.ComponentFields) {
		toSerialize["componentFields"] = o.ComponentFields
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.MultipleUpload) {
		toSerialize["multipleUpload"] = o.MultipleUpload
	}
	return toSerialize, nil
}

type NullableUploadDefinitionXO struct {
	value *UploadDefinitionXO
	isSet bool
}

func (v NullableUploadDefinitionXO) Get() *UploadDefinitionXO {
	return v.value
}

func (v *NullableUploadDefinitionXO) Set(val *UploadDefinitionXO) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadDefinitionXO) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadDefinitionXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadDefinitionXO(val *UploadDefinitionXO) *NullableUploadDefinitionXO {
	return &NullableUploadDefinitionXO{value: val, isSet: true}
}

func (v NullableUploadDefinitionXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadDefinitionXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


