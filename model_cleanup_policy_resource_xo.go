/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.74.0-05.

API version: 3.74.0-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CleanupPolicyResourceXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CleanupPolicyResourceXO{}

// CleanupPolicyResourceXO struct for CleanupPolicyResourceXO
type CleanupPolicyResourceXO struct {
	// asset name matcher (Remove components that have at least one asset name matching the following regular expression pattern:)
	CriteriaAssetRegex *string `json:"criteriaAssetRegex,omitempty"`
	// component age (Components published over “x” days ago (e.g 1-999))
	CriteriaLastBlobUpdated *int64 `json:"criteriaLastBlobUpdated,omitempty"`
	// component usage (Components downloaded in “x” amount of days (e.g 1-999))
	CriteriaLastDownloaded *int64 `json:"criteriaLastDownloaded,omitempty"`
	// release type (Remove components that are of the following release type:)
	CriteriaReleaseType *string `json:"criteriaReleaseType,omitempty"`
	// repository format
	Format string `json:"format"`
	// policy name
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// description
	Notes *string `json:"notes,omitempty"`
	// keep the latest \"x\" number of versions
	Retain *int32 `json:"retain,omitempty"`
}

type _CleanupPolicyResourceXO CleanupPolicyResourceXO

// NewCleanupPolicyResourceXO instantiates a new CleanupPolicyResourceXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCleanupPolicyResourceXO(format string, name string) *CleanupPolicyResourceXO {
	this := CleanupPolicyResourceXO{}
	this.Format = format
	this.Name = name
	return &this
}

// NewCleanupPolicyResourceXOWithDefaults instantiates a new CleanupPolicyResourceXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCleanupPolicyResourceXOWithDefaults() *CleanupPolicyResourceXO {
	this := CleanupPolicyResourceXO{}
	return &this
}

// GetCriteriaAssetRegex returns the CriteriaAssetRegex field value if set, zero value otherwise.
func (o *CleanupPolicyResourceXO) GetCriteriaAssetRegex() string {
	if o == nil || IsNil(o.CriteriaAssetRegex) {
		var ret string
		return ret
	}
	return *o.CriteriaAssetRegex
}

// GetCriteriaAssetRegexOk returns a tuple with the CriteriaAssetRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanupPolicyResourceXO) GetCriteriaAssetRegexOk() (*string, bool) {
	if o == nil || IsNil(o.CriteriaAssetRegex) {
		return nil, false
	}
	return o.CriteriaAssetRegex, true
}

// HasCriteriaAssetRegex returns a boolean if a field has been set.
func (o *CleanupPolicyResourceXO) HasCriteriaAssetRegex() bool {
	if o != nil && !IsNil(o.CriteriaAssetRegex) {
		return true
	}

	return false
}

// SetCriteriaAssetRegex gets a reference to the given string and assigns it to the CriteriaAssetRegex field.
func (o *CleanupPolicyResourceXO) SetCriteriaAssetRegex(v string) {
	o.CriteriaAssetRegex = &v
}

// GetCriteriaLastBlobUpdated returns the CriteriaLastBlobUpdated field value if set, zero value otherwise.
func (o *CleanupPolicyResourceXO) GetCriteriaLastBlobUpdated() int64 {
	if o == nil || IsNil(o.CriteriaLastBlobUpdated) {
		var ret int64
		return ret
	}
	return *o.CriteriaLastBlobUpdated
}

// GetCriteriaLastBlobUpdatedOk returns a tuple with the CriteriaLastBlobUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanupPolicyResourceXO) GetCriteriaLastBlobUpdatedOk() (*int64, bool) {
	if o == nil || IsNil(o.CriteriaLastBlobUpdated) {
		return nil, false
	}
	return o.CriteriaLastBlobUpdated, true
}

// HasCriteriaLastBlobUpdated returns a boolean if a field has been set.
func (o *CleanupPolicyResourceXO) HasCriteriaLastBlobUpdated() bool {
	if o != nil && !IsNil(o.CriteriaLastBlobUpdated) {
		return true
	}

	return false
}

// SetCriteriaLastBlobUpdated gets a reference to the given int64 and assigns it to the CriteriaLastBlobUpdated field.
func (o *CleanupPolicyResourceXO) SetCriteriaLastBlobUpdated(v int64) {
	o.CriteriaLastBlobUpdated = &v
}

// GetCriteriaLastDownloaded returns the CriteriaLastDownloaded field value if set, zero value otherwise.
func (o *CleanupPolicyResourceXO) GetCriteriaLastDownloaded() int64 {
	if o == nil || IsNil(o.CriteriaLastDownloaded) {
		var ret int64
		return ret
	}
	return *o.CriteriaLastDownloaded
}

// GetCriteriaLastDownloadedOk returns a tuple with the CriteriaLastDownloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanupPolicyResourceXO) GetCriteriaLastDownloadedOk() (*int64, bool) {
	if o == nil || IsNil(o.CriteriaLastDownloaded) {
		return nil, false
	}
	return o.CriteriaLastDownloaded, true
}

// HasCriteriaLastDownloaded returns a boolean if a field has been set.
func (o *CleanupPolicyResourceXO) HasCriteriaLastDownloaded() bool {
	if o != nil && !IsNil(o.CriteriaLastDownloaded) {
		return true
	}

	return false
}

// SetCriteriaLastDownloaded gets a reference to the given int64 and assigns it to the CriteriaLastDownloaded field.
func (o *CleanupPolicyResourceXO) SetCriteriaLastDownloaded(v int64) {
	o.CriteriaLastDownloaded = &v
}

// GetCriteriaReleaseType returns the CriteriaReleaseType field value if set, zero value otherwise.
func (o *CleanupPolicyResourceXO) GetCriteriaReleaseType() string {
	if o == nil || IsNil(o.CriteriaReleaseType) {
		var ret string
		return ret
	}
	return *o.CriteriaReleaseType
}

// GetCriteriaReleaseTypeOk returns a tuple with the CriteriaReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanupPolicyResourceXO) GetCriteriaReleaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CriteriaReleaseType) {
		return nil, false
	}
	return o.CriteriaReleaseType, true
}

// HasCriteriaReleaseType returns a boolean if a field has been set.
func (o *CleanupPolicyResourceXO) HasCriteriaReleaseType() bool {
	if o != nil && !IsNil(o.CriteriaReleaseType) {
		return true
	}

	return false
}

// SetCriteriaReleaseType gets a reference to the given string and assigns it to the CriteriaReleaseType field.
func (o *CleanupPolicyResourceXO) SetCriteriaReleaseType(v string) {
	o.CriteriaReleaseType = &v
}

// GetFormat returns the Format field value
func (o *CleanupPolicyResourceXO) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *CleanupPolicyResourceXO) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *CleanupPolicyResourceXO) SetFormat(v string) {
	o.Format = v
}

// GetName returns the Name field value
func (o *CleanupPolicyResourceXO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CleanupPolicyResourceXO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CleanupPolicyResourceXO) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CleanupPolicyResourceXO) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanupPolicyResourceXO) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CleanupPolicyResourceXO) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CleanupPolicyResourceXO) SetNotes(v string) {
	o.Notes = &v
}

// GetRetain returns the Retain field value if set, zero value otherwise.
func (o *CleanupPolicyResourceXO) GetRetain() int32 {
	if o == nil || IsNil(o.Retain) {
		var ret int32
		return ret
	}
	return *o.Retain
}

// GetRetainOk returns a tuple with the Retain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanupPolicyResourceXO) GetRetainOk() (*int32, bool) {
	if o == nil || IsNil(o.Retain) {
		return nil, false
	}
	return o.Retain, true
}

// HasRetain returns a boolean if a field has been set.
func (o *CleanupPolicyResourceXO) HasRetain() bool {
	if o != nil && !IsNil(o.Retain) {
		return true
	}

	return false
}

// SetRetain gets a reference to the given int32 and assigns it to the Retain field.
func (o *CleanupPolicyResourceXO) SetRetain(v int32) {
	o.Retain = &v
}

func (o CleanupPolicyResourceXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CleanupPolicyResourceXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CriteriaAssetRegex) {
		toSerialize["criteriaAssetRegex"] = o.CriteriaAssetRegex
	}
	if !IsNil(o.CriteriaLastBlobUpdated) {
		toSerialize["criteriaLastBlobUpdated"] = o.CriteriaLastBlobUpdated
	}
	if !IsNil(o.CriteriaLastDownloaded) {
		toSerialize["criteriaLastDownloaded"] = o.CriteriaLastDownloaded
	}
	if !IsNil(o.CriteriaReleaseType) {
		toSerialize["criteriaReleaseType"] = o.CriteriaReleaseType
	}
	toSerialize["format"] = o.Format
	toSerialize["name"] = o.Name
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Retain) {
		toSerialize["retain"] = o.Retain
	}
	return toSerialize, nil
}

func (o *CleanupPolicyResourceXO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"format",
		"name",
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

	varCleanupPolicyResourceXO := _CleanupPolicyResourceXO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCleanupPolicyResourceXO)

	if err != nil {
		return err
	}

	*o = CleanupPolicyResourceXO(varCleanupPolicyResourceXO)

	return err
}

type NullableCleanupPolicyResourceXO struct {
	value *CleanupPolicyResourceXO
	isSet bool
}

func (v NullableCleanupPolicyResourceXO) Get() *CleanupPolicyResourceXO {
	return v.value
}

func (v *NullableCleanupPolicyResourceXO) Set(val *CleanupPolicyResourceXO) {
	v.value = val
	v.isSet = true
}

func (v NullableCleanupPolicyResourceXO) IsSet() bool {
	return v.isSet
}

func (v *NullableCleanupPolicyResourceXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCleanupPolicyResourceXO(val *CleanupPolicyResourceXO) *NullableCleanupPolicyResourceXO {
	return &NullableCleanupPolicyResourceXO{value: val, isSet: true}
}

func (v NullableCleanupPolicyResourceXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCleanupPolicyResourceXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

