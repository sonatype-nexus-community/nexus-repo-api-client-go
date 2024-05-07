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

// checks if the AptSigningRepositoriesAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AptSigningRepositoriesAttributes{}

// AptSigningRepositoriesAttributes struct for AptSigningRepositoriesAttributes
type AptSigningRepositoriesAttributes struct {
	// PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)
	Keypair *string `json:"keypair,omitempty"`
	// Passphrase to access PGP signing key
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewAptSigningRepositoriesAttributes instantiates a new AptSigningRepositoriesAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAptSigningRepositoriesAttributes() *AptSigningRepositoriesAttributes {
	this := AptSigningRepositoriesAttributes{}
	return &this
}

// NewAptSigningRepositoriesAttributesWithDefaults instantiates a new AptSigningRepositoriesAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAptSigningRepositoriesAttributesWithDefaults() *AptSigningRepositoriesAttributes {
	this := AptSigningRepositoriesAttributes{}
	return &this
}

// GetKeypair returns the Keypair field value if set, zero value otherwise.
func (o *AptSigningRepositoriesAttributes) GetKeypair() string {
	if o == nil || IsNil(o.Keypair) {
		var ret string
		return ret
	}
	return *o.Keypair
}

// GetKeypairOk returns a tuple with the Keypair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptSigningRepositoriesAttributes) GetKeypairOk() (*string, bool) {
	if o == nil || IsNil(o.Keypair) {
		return nil, false
	}
	return o.Keypair, true
}

// HasKeypair returns a boolean if a field has been set.
func (o *AptSigningRepositoriesAttributes) HasKeypair() bool {
	if o != nil && !IsNil(o.Keypair) {
		return true
	}

	return false
}

// SetKeypair gets a reference to the given string and assigns it to the Keypair field.
func (o *AptSigningRepositoriesAttributes) SetKeypair(v string) {
	o.Keypair = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *AptSigningRepositoriesAttributes) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptSigningRepositoriesAttributes) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *AptSigningRepositoriesAttributes) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *AptSigningRepositoriesAttributes) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o AptSigningRepositoriesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AptSigningRepositoriesAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keypair) {
		toSerialize["keypair"] = o.Keypair
	}
	if !IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	return toSerialize, nil
}

type NullableAptSigningRepositoriesAttributes struct {
	value *AptSigningRepositoriesAttributes
	isSet bool
}

func (v NullableAptSigningRepositoriesAttributes) Get() *AptSigningRepositoriesAttributes {
	return v.value
}

func (v *NullableAptSigningRepositoriesAttributes) Set(val *AptSigningRepositoriesAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAptSigningRepositoriesAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAptSigningRepositoriesAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAptSigningRepositoriesAttributes(val *AptSigningRepositoriesAttributes) *NullableAptSigningRepositoriesAttributes {
	return &NullableAptSigningRepositoriesAttributes{value: val, isSet: true}
}

func (v NullableAptSigningRepositoriesAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAptSigningRepositoriesAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


