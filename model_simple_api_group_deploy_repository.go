/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.82.0-08.

API version: 3.82.0-08
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SimpleApiGroupDeployRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleApiGroupDeployRepository{}

// SimpleApiGroupDeployRepository struct for SimpleApiGroupDeployRepository
type SimpleApiGroupDeployRepository struct {
	Format *string `json:"format,omitempty"`
	Group GroupDeployAttributes `json:"group"`
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
}

type _SimpleApiGroupDeployRepository SimpleApiGroupDeployRepository

// NewSimpleApiGroupDeployRepository instantiates a new SimpleApiGroupDeployRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleApiGroupDeployRepository(group GroupDeployAttributes, online bool, storage StorageAttributes) *SimpleApiGroupDeployRepository {
	this := SimpleApiGroupDeployRepository{}
	this.Group = group
	this.Online = online
	this.Storage = storage
	var type_ string = "group"
	this.Type = &type_
	return &this
}

// NewSimpleApiGroupDeployRepositoryWithDefaults instantiates a new SimpleApiGroupDeployRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleApiGroupDeployRepositoryWithDefaults() *SimpleApiGroupDeployRepository {
	this := SimpleApiGroupDeployRepository{}
	var type_ string = "group"
	this.Type = &type_
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *SimpleApiGroupDeployRepository) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleApiGroupDeployRepository) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *SimpleApiGroupDeployRepository) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *SimpleApiGroupDeployRepository) SetFormat(v string) {
	o.Format = &v
}

// GetGroup returns the Group field value
func (o *SimpleApiGroupDeployRepository) GetGroup() GroupDeployAttributes {
	if o == nil {
		var ret GroupDeployAttributes
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SimpleApiGroupDeployRepository) GetGroupOk() (*GroupDeployAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SimpleApiGroupDeployRepository) SetGroup(v GroupDeployAttributes) {
	o.Group = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimpleApiGroupDeployRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleApiGroupDeployRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimpleApiGroupDeployRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimpleApiGroupDeployRepository) SetName(v string) {
	o.Name = &v
}

// GetOnline returns the Online field value
func (o *SimpleApiGroupDeployRepository) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *SimpleApiGroupDeployRepository) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *SimpleApiGroupDeployRepository) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *SimpleApiGroupDeployRepository) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *SimpleApiGroupDeployRepository) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *SimpleApiGroupDeployRepository) SetStorage(v StorageAttributes) {
	o.Storage = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimpleApiGroupDeployRepository) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleApiGroupDeployRepository) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimpleApiGroupDeployRepository) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SimpleApiGroupDeployRepository) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SimpleApiGroupDeployRepository) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleApiGroupDeployRepository) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SimpleApiGroupDeployRepository) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SimpleApiGroupDeployRepository) SetUrl(v string) {
	o.Url = &v
}

func (o SimpleApiGroupDeployRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleApiGroupDeployRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	toSerialize["group"] = o.Group
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

func (o *SimpleApiGroupDeployRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group",
		"online",
		"storage",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSimpleApiGroupDeployRepository := _SimpleApiGroupDeployRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimpleApiGroupDeployRepository)

	if err != nil {
		return err
	}

	*o = SimpleApiGroupDeployRepository(varSimpleApiGroupDeployRepository)

	return err
}

type NullableSimpleApiGroupDeployRepository struct {
	value *SimpleApiGroupDeployRepository
	isSet bool
}

func (v NullableSimpleApiGroupDeployRepository) Get() *SimpleApiGroupDeployRepository {
	return v.value
}

func (v *NullableSimpleApiGroupDeployRepository) Set(val *SimpleApiGroupDeployRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleApiGroupDeployRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleApiGroupDeployRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleApiGroupDeployRepository(val *SimpleApiGroupDeployRepository) *NullableSimpleApiGroupDeployRepository {
	return &NullableSimpleApiGroupDeployRepository{value: val, isSet: true}
}

func (v NullableSimpleApiGroupDeployRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleApiGroupDeployRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


