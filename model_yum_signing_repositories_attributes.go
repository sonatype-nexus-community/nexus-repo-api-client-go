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

// checks if the YumSigningRepositoriesAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &YumSigningRepositoriesAttributes{}

// YumSigningRepositoriesAttributes struct for YumSigningRepositoriesAttributes
type YumSigningRepositoriesAttributes struct {
	// PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)
	Keypair *string `json:"keypair,omitempty"`
	// Passphrase to access PGP signing key
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewYumSigningRepositoriesAttributes instantiates a new YumSigningRepositoriesAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYumSigningRepositoriesAttributes() *YumSigningRepositoriesAttributes {
	this := YumSigningRepositoriesAttributes{}
	return &this
}

// NewYumSigningRepositoriesAttributesWithDefaults instantiates a new YumSigningRepositoriesAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYumSigningRepositoriesAttributesWithDefaults() *YumSigningRepositoriesAttributes {
	this := YumSigningRepositoriesAttributes{}
	return &this
}

// GetKeypair returns the Keypair field value if set, zero value otherwise.
func (o *YumSigningRepositoriesAttributes) GetKeypair() string {
	if o == nil || IsNil(o.Keypair) {
		var ret string
		return ret
	}
	return *o.Keypair
}

// GetKeypairOk returns a tuple with the Keypair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YumSigningRepositoriesAttributes) GetKeypairOk() (*string, bool) {
	if o == nil || IsNil(o.Keypair) {
		return nil, false
	}
	return o.Keypair, true
}

// HasKeypair returns a boolean if a field has been set.
func (o *YumSigningRepositoriesAttributes) HasKeypair() bool {
	if o != nil && !IsNil(o.Keypair) {
		return true
	}

	return false
}

// SetKeypair gets a reference to the given string and assigns it to the Keypair field.
func (o *YumSigningRepositoriesAttributes) SetKeypair(v string) {
	o.Keypair = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *YumSigningRepositoriesAttributes) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YumSigningRepositoriesAttributes) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *YumSigningRepositoriesAttributes) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *YumSigningRepositoriesAttributes) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o YumSigningRepositoriesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o YumSigningRepositoriesAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keypair) {
		toSerialize["keypair"] = o.Keypair
	}
	if !IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	return toSerialize, nil
}

type NullableYumSigningRepositoriesAttributes struct {
	value *YumSigningRepositoriesAttributes
	isSet bool
}

func (v NullableYumSigningRepositoriesAttributes) Get() *YumSigningRepositoriesAttributes {
	return v.value
}

func (v *NullableYumSigningRepositoriesAttributes) Set(val *YumSigningRepositoriesAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableYumSigningRepositoriesAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableYumSigningRepositoriesAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYumSigningRepositoriesAttributes(val *YumSigningRepositoriesAttributes) *NullableYumSigningRepositoriesAttributes {
	return &NullableYumSigningRepositoriesAttributes{value: val, isSet: true}
}

func (v NullableYumSigningRepositoriesAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYumSigningRepositoriesAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


