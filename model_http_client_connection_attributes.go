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

// checks if the HttpClientConnectionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpClientConnectionAttributes{}

// HttpClientConnectionAttributes struct for HttpClientConnectionAttributes
type HttpClientConnectionAttributes struct {
	// Whether to enable redirects to the same location (may be required by some servers)
	EnableCircularRedirects *bool `json:"enableCircularRedirects,omitempty"`
	// Whether to allow cookies to be stored and used
	EnableCookies *bool `json:"enableCookies,omitempty"`
	// Total retries if the initial connection attempt suffers a timeout
	Retries *int32 `json:"retries,omitempty"`
	// Seconds to wait for activity before stopping and retrying the connection
	Timeout *int32 `json:"timeout,omitempty"`
	// Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	UseTrustStore *bool `json:"useTrustStore,omitempty"`
	// Custom fragment to append to User-Agent header in HTTP requests
	UserAgentSuffix *string `json:"userAgentSuffix,omitempty"`
}

// NewHttpClientConnectionAttributes instantiates a new HttpClientConnectionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpClientConnectionAttributes() *HttpClientConnectionAttributes {
	this := HttpClientConnectionAttributes{}
	return &this
}

// NewHttpClientConnectionAttributesWithDefaults instantiates a new HttpClientConnectionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpClientConnectionAttributesWithDefaults() *HttpClientConnectionAttributes {
	this := HttpClientConnectionAttributes{}
	return &this
}

// GetEnableCircularRedirects returns the EnableCircularRedirects field value if set, zero value otherwise.
func (o *HttpClientConnectionAttributes) GetEnableCircularRedirects() bool {
	if o == nil || IsNil(o.EnableCircularRedirects) {
		var ret bool
		return ret
	}
	return *o.EnableCircularRedirects
}

// GetEnableCircularRedirectsOk returns a tuple with the EnableCircularRedirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAttributes) GetEnableCircularRedirectsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCircularRedirects) {
		return nil, false
	}
	return o.EnableCircularRedirects, true
}

// HasEnableCircularRedirects returns a boolean if a field has been set.
func (o *HttpClientConnectionAttributes) HasEnableCircularRedirects() bool {
	if o != nil && !IsNil(o.EnableCircularRedirects) {
		return true
	}

	return false
}

// SetEnableCircularRedirects gets a reference to the given bool and assigns it to the EnableCircularRedirects field.
func (o *HttpClientConnectionAttributes) SetEnableCircularRedirects(v bool) {
	o.EnableCircularRedirects = &v
}

// GetEnableCookies returns the EnableCookies field value if set, zero value otherwise.
func (o *HttpClientConnectionAttributes) GetEnableCookies() bool {
	if o == nil || IsNil(o.EnableCookies) {
		var ret bool
		return ret
	}
	return *o.EnableCookies
}

// GetEnableCookiesOk returns a tuple with the EnableCookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAttributes) GetEnableCookiesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCookies) {
		return nil, false
	}
	return o.EnableCookies, true
}

// HasEnableCookies returns a boolean if a field has been set.
func (o *HttpClientConnectionAttributes) HasEnableCookies() bool {
	if o != nil && !IsNil(o.EnableCookies) {
		return true
	}

	return false
}

// SetEnableCookies gets a reference to the given bool and assigns it to the EnableCookies field.
func (o *HttpClientConnectionAttributes) SetEnableCookies(v bool) {
	o.EnableCookies = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *HttpClientConnectionAttributes) GetRetries() int32 {
	if o == nil || IsNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAttributes) GetRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *HttpClientConnectionAttributes) HasRetries() bool {
	if o != nil && !IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *HttpClientConnectionAttributes) SetRetries(v int32) {
	o.Retries = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *HttpClientConnectionAttributes) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAttributes) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *HttpClientConnectionAttributes) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *HttpClientConnectionAttributes) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetUseTrustStore returns the UseTrustStore field value if set, zero value otherwise.
func (o *HttpClientConnectionAttributes) GetUseTrustStore() bool {
	if o == nil || IsNil(o.UseTrustStore) {
		var ret bool
		return ret
	}
	return *o.UseTrustStore
}

// GetUseTrustStoreOk returns a tuple with the UseTrustStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAttributes) GetUseTrustStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTrustStore) {
		return nil, false
	}
	return o.UseTrustStore, true
}

// HasUseTrustStore returns a boolean if a field has been set.
func (o *HttpClientConnectionAttributes) HasUseTrustStore() bool {
	if o != nil && !IsNil(o.UseTrustStore) {
		return true
	}

	return false
}

// SetUseTrustStore gets a reference to the given bool and assigns it to the UseTrustStore field.
func (o *HttpClientConnectionAttributes) SetUseTrustStore(v bool) {
	o.UseTrustStore = &v
}

// GetUserAgentSuffix returns the UserAgentSuffix field value if set, zero value otherwise.
func (o *HttpClientConnectionAttributes) GetUserAgentSuffix() string {
	if o == nil || IsNil(o.UserAgentSuffix) {
		var ret string
		return ret
	}
	return *o.UserAgentSuffix
}

// GetUserAgentSuffixOk returns a tuple with the UserAgentSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientConnectionAttributes) GetUserAgentSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgentSuffix) {
		return nil, false
	}
	return o.UserAgentSuffix, true
}

// HasUserAgentSuffix returns a boolean if a field has been set.
func (o *HttpClientConnectionAttributes) HasUserAgentSuffix() bool {
	if o != nil && !IsNil(o.UserAgentSuffix) {
		return true
	}

	return false
}

// SetUserAgentSuffix gets a reference to the given string and assigns it to the UserAgentSuffix field.
func (o *HttpClientConnectionAttributes) SetUserAgentSuffix(v string) {
	o.UserAgentSuffix = &v
}

func (o HttpClientConnectionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpClientConnectionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableCircularRedirects) {
		toSerialize["enableCircularRedirects"] = o.EnableCircularRedirects
	}
	if !IsNil(o.EnableCookies) {
		toSerialize["enableCookies"] = o.EnableCookies
	}
	if !IsNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.UseTrustStore) {
		toSerialize["useTrustStore"] = o.UseTrustStore
	}
	if !IsNil(o.UserAgentSuffix) {
		toSerialize["userAgentSuffix"] = o.UserAgentSuffix
	}
	return toSerialize, nil
}

type NullableHttpClientConnectionAttributes struct {
	value *HttpClientConnectionAttributes
	isSet bool
}

func (v NullableHttpClientConnectionAttributes) Get() *HttpClientConnectionAttributes {
	return v.value
}

func (v *NullableHttpClientConnectionAttributes) Set(val *HttpClientConnectionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpClientConnectionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpClientConnectionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpClientConnectionAttributes(val *HttpClientConnectionAttributes) *NullableHttpClientConnectionAttributes {
	return &NullableHttpClientConnectionAttributes{value: val, isSet: true}
}

func (v NullableHttpClientConnectionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpClientConnectionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


