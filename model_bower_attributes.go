/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.67.1-01.

API version: 3.67.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BowerAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BowerAttributes{}

// BowerAttributes struct for BowerAttributes
type BowerAttributes struct {
	// Whether to force Bower to retrieve packages through this proxy repository
	RewritePackageUrls bool `json:"rewritePackageUrls"`
}

type _BowerAttributes BowerAttributes

// NewBowerAttributes instantiates a new BowerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBowerAttributes(rewritePackageUrls bool) *BowerAttributes {
	this := BowerAttributes{}
	this.RewritePackageUrls = rewritePackageUrls
	return &this
}

// NewBowerAttributesWithDefaults instantiates a new BowerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBowerAttributesWithDefaults() *BowerAttributes {
	this := BowerAttributes{}
	return &this
}

// GetRewritePackageUrls returns the RewritePackageUrls field value
func (o *BowerAttributes) GetRewritePackageUrls() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RewritePackageUrls
}

// GetRewritePackageUrlsOk returns a tuple with the RewritePackageUrls field value
// and a boolean to check if the value has been set.
func (o *BowerAttributes) GetRewritePackageUrlsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewritePackageUrls, true
}

// SetRewritePackageUrls sets field value
func (o *BowerAttributes) SetRewritePackageUrls(v bool) {
	o.RewritePackageUrls = v
}

func (o BowerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BowerAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rewritePackageUrls"] = o.RewritePackageUrls
	return toSerialize, nil
}

func (o *BowerAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rewritePackageUrls",
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

	varBowerAttributes := _BowerAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBowerAttributes)

	if err != nil {
		return err
	}

	*o = BowerAttributes(varBowerAttributes)

	return err
}

type NullableBowerAttributes struct {
	value *BowerAttributes
	isSet bool
}

func (v NullableBowerAttributes) Get() *BowerAttributes {
	return v.value
}

func (v *NullableBowerAttributes) Set(val *BowerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBowerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBowerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBowerAttributes(val *BowerAttributes) *NullableBowerAttributes {
	return &NullableBowerAttributes{value: val, isSet: true}
}

func (v NullableBowerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBowerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


