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

// checks if the PageTagXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageTagXO{}

// PageTagXO struct for PageTagXO
type PageTagXO struct {
	ContinuationToken *string `json:"continuationToken,omitempty"`
	Items []TagXO `json:"items,omitempty"`
}

// NewPageTagXO instantiates a new PageTagXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageTagXO() *PageTagXO {
	this := PageTagXO{}
	return &this
}

// NewPageTagXOWithDefaults instantiates a new PageTagXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageTagXOWithDefaults() *PageTagXO {
	this := PageTagXO{}
	return &this
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *PageTagXO) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken) {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTagXO) GetContinuationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContinuationToken) {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *PageTagXO) HasContinuationToken() bool {
	if o != nil && !IsNil(o.ContinuationToken) {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *PageTagXO) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PageTagXO) GetItems() []TagXO {
	if o == nil || IsNil(o.Items) {
		var ret []TagXO
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTagXO) GetItemsOk() ([]TagXO, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PageTagXO) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TagXO and assigns it to the Items field.
func (o *PageTagXO) SetItems(v []TagXO) {
	o.Items = v
}

func (o PageTagXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageTagXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContinuationToken) {
		toSerialize["continuationToken"] = o.ContinuationToken
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullablePageTagXO struct {
	value *PageTagXO
	isSet bool
}

func (v NullablePageTagXO) Get() *PageTagXO {
	return v.value
}

func (v *NullablePageTagXO) Set(val *PageTagXO) {
	v.value = val
	v.isSet = true
}

func (v NullablePageTagXO) IsSet() bool {
	return v.isSet
}

func (v *NullablePageTagXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageTagXO(val *PageTagXO) *NullablePageTagXO {
	return &NullablePageTagXO{value: val, isSet: true}
}

func (v NullablePageTagXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageTagXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


