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

// checks if the YumHostedApiRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &YumHostedApiRepository{}

// YumHostedApiRepository struct for YumHostedApiRepository
type YumHostedApiRepository struct {
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes `json:"component,omitempty"`
	Format *string `json:"format,omitempty"`
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
	Yum YumAttributes `json:"yum"`
}

type _YumHostedApiRepository YumHostedApiRepository

// NewYumHostedApiRepository instantiates a new YumHostedApiRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYumHostedApiRepository(online bool, storage HostedStorageAttributes, yum YumAttributes) *YumHostedApiRepository {
	this := YumHostedApiRepository{}
	var format string = "yum"
	this.Format = &format
	this.Online = online
	this.Storage = storage
	var type_ string = "hosted"
	this.Type = &type_
	this.Yum = yum
	return &this
}

// NewYumHostedApiRepositoryWithDefaults instantiates a new YumHostedApiRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYumHostedApiRepositoryWithDefaults() *YumHostedApiRepository {
	this := YumHostedApiRepository{}
	var format string = "yum"
	this.Format = &format
	var type_ string = "hosted"
	this.Type = &type_
	return &this
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *YumHostedApiRepository) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YumHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *YumHostedApiRepository) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *YumHostedApiRepository) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *YumHostedApiRepository) GetComponent() ComponentAttributes {
	if o == nil || IsNil(o.Component) {
		var ret ComponentAttributes
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YumHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *YumHostedApiRepository) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ComponentAttributes and assigns it to the Component field.
func (o *YumHostedApiRepository) SetComponent(v ComponentAttributes) {
	o.Component = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *YumHostedApiRepository) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YumHostedApiRepository) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *YumHostedApiRepository) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *YumHostedApiRepository) SetFormat(v string) {
	o.Format = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *YumHostedApiRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YumHostedApiRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *YumHostedApiRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *YumHostedApiRepository) SetName(v string) {
	o.Name = &v
}

// GetOnline returns the Online field value
func (o *YumHostedApiRepository) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *YumHostedApiRepository) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *YumHostedApiRepository) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *YumHostedApiRepository) GetStorage() HostedStorageAttributes {
	if o == nil {
		var ret HostedStorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *YumHostedApiRepository) GetStorageOk() (*HostedStorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *YumHostedApiRepository) SetStorage(v HostedStorageAttributes) {
	o.Storage = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *YumHostedApiRepository) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YumHostedApiRepository) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *YumHostedApiRepository) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *YumHostedApiRepository) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *YumHostedApiRepository) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YumHostedApiRepository) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *YumHostedApiRepository) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *YumHostedApiRepository) SetUrl(v string) {
	o.Url = &v
}

// GetYum returns the Yum field value
func (o *YumHostedApiRepository) GetYum() YumAttributes {
	if o == nil {
		var ret YumAttributes
		return ret
	}

	return o.Yum
}

// GetYumOk returns a tuple with the Yum field value
// and a boolean to check if the value has been set.
func (o *YumHostedApiRepository) GetYumOk() (*YumAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yum, true
}

// SetYum sets field value
func (o *YumHostedApiRepository) SetYum(v YumAttributes) {
	o.Yum = v
}

func (o YumHostedApiRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o YumHostedApiRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
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
	toSerialize["yum"] = o.Yum
	return toSerialize, nil
}

func (o *YumHostedApiRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"online",
		"storage",
		"yum",
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

	varYumHostedApiRepository := _YumHostedApiRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varYumHostedApiRepository)

	if err != nil {
		return err
	}

	*o = YumHostedApiRepository(varYumHostedApiRepository)

	return err
}

type NullableYumHostedApiRepository struct {
	value *YumHostedApiRepository
	isSet bool
}

func (v NullableYumHostedApiRepository) Get() *YumHostedApiRepository {
	return v.value
}

func (v *NullableYumHostedApiRepository) Set(val *YumHostedApiRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableYumHostedApiRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableYumHostedApiRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYumHostedApiRepository(val *YumHostedApiRepository) *NullableYumHostedApiRepository {
	return &NullableYumHostedApiRepository{value: val, isSet: true}
}

func (v NullableYumHostedApiRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYumHostedApiRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


