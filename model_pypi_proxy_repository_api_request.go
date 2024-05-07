/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.67.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PypiProxyRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PypiProxyRepositoryApiRequest{}

// PypiProxyRepositoryApiRequest struct for PypiProxyRepositoryApiRequest
type PypiProxyRepositoryApiRequest struct {
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient HttpClientAttributes `json:"httpClient"`
	// A unique identifier for this repository
	Name string `json:"name"`
	NegativeCache NegativeCacheAttributes `json:"negativeCache"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Proxy ProxyAttributes `json:"proxy"`
	Pypi *PyPiProxyAttributes `json:"pypi,omitempty"`
	Replication *ReplicationAttributes `json:"replication,omitempty"`
	RoutingRule *string `json:"routingRule,omitempty"`
	Storage StorageAttributes `json:"storage"`
}

type _PypiProxyRepositoryApiRequest PypiProxyRepositoryApiRequest

// NewPypiProxyRepositoryApiRequest instantiates a new PypiProxyRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPypiProxyRepositoryApiRequest(httpClient HttpClientAttributes, name string, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes) *PypiProxyRepositoryApiRequest {
	this := PypiProxyRepositoryApiRequest{}
	this.HttpClient = httpClient
	this.Name = name
	this.NegativeCache = negativeCache
	this.Online = online
	this.Proxy = proxy
	this.Storage = storage
	return &this
}

// NewPypiProxyRepositoryApiRequestWithDefaults instantiates a new PypiProxyRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPypiProxyRepositoryApiRequestWithDefaults() *PypiProxyRepositoryApiRequest {
	this := PypiProxyRepositoryApiRequest{}
	return &this
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *PypiProxyRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *PypiProxyRepositoryApiRequest) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *PypiProxyRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetHttpClient returns the HttpClient field value
func (o *PypiProxyRepositoryApiRequest) GetHttpClient() HttpClientAttributes {
	if o == nil {
		var ret HttpClientAttributes
		return ret
	}

	return o.HttpClient
}

// GetHttpClientOk returns a tuple with the HttpClient field value
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetHttpClientOk() (*HttpClientAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpClient, true
}

// SetHttpClient sets field value
func (o *PypiProxyRepositoryApiRequest) SetHttpClient(v HttpClientAttributes) {
	o.HttpClient = v
}

// GetName returns the Name field value
func (o *PypiProxyRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PypiProxyRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetNegativeCache returns the NegativeCache field value
func (o *PypiProxyRepositoryApiRequest) GetNegativeCache() NegativeCacheAttributes {
	if o == nil {
		var ret NegativeCacheAttributes
		return ret
	}

	return o.NegativeCache
}

// GetNegativeCacheOk returns a tuple with the NegativeCache field value
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetNegativeCacheOk() (*NegativeCacheAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeCache, true
}

// SetNegativeCache sets field value
func (o *PypiProxyRepositoryApiRequest) SetNegativeCache(v NegativeCacheAttributes) {
	o.NegativeCache = v
}

// GetOnline returns the Online field value
func (o *PypiProxyRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *PypiProxyRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetProxy returns the Proxy field value
func (o *PypiProxyRepositoryApiRequest) GetProxy() ProxyAttributes {
	if o == nil {
		var ret ProxyAttributes
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetProxyOk() (*ProxyAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *PypiProxyRepositoryApiRequest) SetProxy(v ProxyAttributes) {
	o.Proxy = v
}

// GetPypi returns the Pypi field value if set, zero value otherwise.
func (o *PypiProxyRepositoryApiRequest) GetPypi() PyPiProxyAttributes {
	if o == nil || IsNil(o.Pypi) {
		var ret PyPiProxyAttributes
		return ret
	}
	return *o.Pypi
}

// GetPypiOk returns a tuple with the Pypi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetPypiOk() (*PyPiProxyAttributes, bool) {
	if o == nil || IsNil(o.Pypi) {
		return nil, false
	}
	return o.Pypi, true
}

// HasPypi returns a boolean if a field has been set.
func (o *PypiProxyRepositoryApiRequest) HasPypi() bool {
	if o != nil && !IsNil(o.Pypi) {
		return true
	}

	return false
}

// SetPypi gets a reference to the given PyPiProxyAttributes and assigns it to the Pypi field.
func (o *PypiProxyRepositoryApiRequest) SetPypi(v PyPiProxyAttributes) {
	o.Pypi = &v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *PypiProxyRepositoryApiRequest) GetReplication() ReplicationAttributes {
	if o == nil || IsNil(o.Replication) {
		var ret ReplicationAttributes
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetReplicationOk() (*ReplicationAttributes, bool) {
	if o == nil || IsNil(o.Replication) {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *PypiProxyRepositoryApiRequest) HasReplication() bool {
	if o != nil && !IsNil(o.Replication) {
		return true
	}

	return false
}

// SetReplication gets a reference to the given ReplicationAttributes and assigns it to the Replication field.
func (o *PypiProxyRepositoryApiRequest) SetReplication(v ReplicationAttributes) {
	o.Replication = &v
}

// GetRoutingRule returns the RoutingRule field value if set, zero value otherwise.
func (o *PypiProxyRepositoryApiRequest) GetRoutingRule() string {
	if o == nil || IsNil(o.RoutingRule) {
		var ret string
		return ret
	}
	return *o.RoutingRule
}

// GetRoutingRuleOk returns a tuple with the RoutingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetRoutingRuleOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingRule) {
		return nil, false
	}
	return o.RoutingRule, true
}

// HasRoutingRule returns a boolean if a field has been set.
func (o *PypiProxyRepositoryApiRequest) HasRoutingRule() bool {
	if o != nil && !IsNil(o.RoutingRule) {
		return true
	}

	return false
}

// SetRoutingRule gets a reference to the given string and assigns it to the RoutingRule field.
func (o *PypiProxyRepositoryApiRequest) SetRoutingRule(v string) {
	o.RoutingRule = &v
}

// GetStorage returns the Storage field value
func (o *PypiProxyRepositoryApiRequest) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *PypiProxyRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *PypiProxyRepositoryApiRequest) SetStorage(v StorageAttributes) {
	o.Storage = v
}

func (o PypiProxyRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PypiProxyRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	toSerialize["httpClient"] = o.HttpClient
	toSerialize["name"] = o.Name
	toSerialize["negativeCache"] = o.NegativeCache
	toSerialize["online"] = o.Online
	toSerialize["proxy"] = o.Proxy
	if !IsNil(o.Pypi) {
		toSerialize["pypi"] = o.Pypi
	}
	if !IsNil(o.Replication) {
		toSerialize["replication"] = o.Replication
	}
	if !IsNil(o.RoutingRule) {
		toSerialize["routingRule"] = o.RoutingRule
	}
	toSerialize["storage"] = o.Storage
	return toSerialize, nil
}

func (o *PypiProxyRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"httpClient",
		"name",
		"negativeCache",
		"online",
		"proxy",
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

	varPypiProxyRepositoryApiRequest := _PypiProxyRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPypiProxyRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = PypiProxyRepositoryApiRequest(varPypiProxyRepositoryApiRequest)

	return err
}

type NullablePypiProxyRepositoryApiRequest struct {
	value *PypiProxyRepositoryApiRequest
	isSet bool
}

func (v NullablePypiProxyRepositoryApiRequest) Get() *PypiProxyRepositoryApiRequest {
	return v.value
}

func (v *NullablePypiProxyRepositoryApiRequest) Set(val *PypiProxyRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePypiProxyRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePypiProxyRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePypiProxyRepositoryApiRequest(val *PypiProxyRepositoryApiRequest) *NullablePypiProxyRepositoryApiRequest {
	return &NullablePypiProxyRepositoryApiRequest{value: val, isSet: true}
}

func (v NullablePypiProxyRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePypiProxyRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


