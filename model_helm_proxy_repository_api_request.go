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

// checks if the HelmProxyRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmProxyRepositoryApiRequest{}

// HelmProxyRepositoryApiRequest struct for HelmProxyRepositoryApiRequest
type HelmProxyRepositoryApiRequest struct {
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient HttpClientAttributes `json:"httpClient"`
	// A unique identifier for this repository
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	NegativeCache NegativeCacheAttributes `json:"negativeCache"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Proxy ProxyAttributes `json:"proxy"`
	Replication *ReplicationAttributes `json:"replication,omitempty"`
	RoutingRule *string `json:"routingRule,omitempty"`
	Storage StorageAttributes `json:"storage"`
}

type _HelmProxyRepositoryApiRequest HelmProxyRepositoryApiRequest

// NewHelmProxyRepositoryApiRequest instantiates a new HelmProxyRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmProxyRepositoryApiRequest(httpClient HttpClientAttributes, name string, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes) *HelmProxyRepositoryApiRequest {
	this := HelmProxyRepositoryApiRequest{}
	this.HttpClient = httpClient
	this.Name = name
	this.NegativeCache = negativeCache
	this.Online = online
	this.Proxy = proxy
	this.Storage = storage
	return &this
}

// NewHelmProxyRepositoryApiRequestWithDefaults instantiates a new HelmProxyRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmProxyRepositoryApiRequestWithDefaults() *HelmProxyRepositoryApiRequest {
	this := HelmProxyRepositoryApiRequest{}
	return &this
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *HelmProxyRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmProxyRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *HelmProxyRepositoryApiRequest) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *HelmProxyRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetHttpClient returns the HttpClient field value
func (o *HelmProxyRepositoryApiRequest) GetHttpClient() HttpClientAttributes {
	if o == nil {
		var ret HttpClientAttributes
		return ret
	}

	return o.HttpClient
}

// GetHttpClientOk returns a tuple with the HttpClient field value
// and a boolean to check if the value has been set.
func (o *HelmProxyRepositoryApiRequest) GetHttpClientOk() (*HttpClientAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpClient, true
}

// SetHttpClient sets field value
func (o *HelmProxyRepositoryApiRequest) SetHttpClient(v HttpClientAttributes) {
	o.HttpClient = v
}

// GetName returns the Name field value
func (o *HelmProxyRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HelmProxyRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HelmProxyRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetNegativeCache returns the NegativeCache field value
func (o *HelmProxyRepositoryApiRequest) GetNegativeCache() NegativeCacheAttributes {
	if o == nil {
		var ret NegativeCacheAttributes
		return ret
	}

	return o.NegativeCache
}

// GetNegativeCacheOk returns a tuple with the NegativeCache field value
// and a boolean to check if the value has been set.
func (o *HelmProxyRepositoryApiRequest) GetNegativeCacheOk() (*NegativeCacheAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeCache, true
}

// SetNegativeCache sets field value
func (o *HelmProxyRepositoryApiRequest) SetNegativeCache(v NegativeCacheAttributes) {
	o.NegativeCache = v
}

// GetOnline returns the Online field value
func (o *HelmProxyRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *HelmProxyRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *HelmProxyRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetProxy returns the Proxy field value
func (o *HelmProxyRepositoryApiRequest) GetProxy() ProxyAttributes {
	if o == nil {
		var ret ProxyAttributes
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *HelmProxyRepositoryApiRequest) GetProxyOk() (*ProxyAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *HelmProxyRepositoryApiRequest) SetProxy(v ProxyAttributes) {
	o.Proxy = v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *HelmProxyRepositoryApiRequest) GetReplication() ReplicationAttributes {
	if o == nil || IsNil(o.Replication) {
		var ret ReplicationAttributes
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmProxyRepositoryApiRequest) GetReplicationOk() (*ReplicationAttributes, bool) {
	if o == nil || IsNil(o.Replication) {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *HelmProxyRepositoryApiRequest) HasReplication() bool {
	if o != nil && !IsNil(o.Replication) {
		return true
	}

	return false
}

// SetReplication gets a reference to the given ReplicationAttributes and assigns it to the Replication field.
func (o *HelmProxyRepositoryApiRequest) SetReplication(v ReplicationAttributes) {
	o.Replication = &v
}

// GetRoutingRule returns the RoutingRule field value if set, zero value otherwise.
func (o *HelmProxyRepositoryApiRequest) GetRoutingRule() string {
	if o == nil || IsNil(o.RoutingRule) {
		var ret string
		return ret
	}
	return *o.RoutingRule
}

// GetRoutingRuleOk returns a tuple with the RoutingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmProxyRepositoryApiRequest) GetRoutingRuleOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingRule) {
		return nil, false
	}
	return o.RoutingRule, true
}

// HasRoutingRule returns a boolean if a field has been set.
func (o *HelmProxyRepositoryApiRequest) HasRoutingRule() bool {
	if o != nil && !IsNil(o.RoutingRule) {
		return true
	}

	return false
}

// SetRoutingRule gets a reference to the given string and assigns it to the RoutingRule field.
func (o *HelmProxyRepositoryApiRequest) SetRoutingRule(v string) {
	o.RoutingRule = &v
}

// GetStorage returns the Storage field value
func (o *HelmProxyRepositoryApiRequest) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *HelmProxyRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *HelmProxyRepositoryApiRequest) SetStorage(v StorageAttributes) {
	o.Storage = v
}

func (o HelmProxyRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmProxyRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	toSerialize["httpClient"] = o.HttpClient
	toSerialize["name"] = o.Name
	toSerialize["negativeCache"] = o.NegativeCache
	toSerialize["online"] = o.Online
	toSerialize["proxy"] = o.Proxy
	if !IsNil(o.Replication) {
		toSerialize["replication"] = o.Replication
	}
	if !IsNil(o.RoutingRule) {
		toSerialize["routingRule"] = o.RoutingRule
	}
	toSerialize["storage"] = o.Storage
	return toSerialize, nil
}

func (o *HelmProxyRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
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

	varHelmProxyRepositoryApiRequest := _HelmProxyRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHelmProxyRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = HelmProxyRepositoryApiRequest(varHelmProxyRepositoryApiRequest)

	return err
}

type NullableHelmProxyRepositoryApiRequest struct {
	value *HelmProxyRepositoryApiRequest
	isSet bool
}

func (v NullableHelmProxyRepositoryApiRequest) Get() *HelmProxyRepositoryApiRequest {
	return v.value
}

func (v *NullableHelmProxyRepositoryApiRequest) Set(val *HelmProxyRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmProxyRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmProxyRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmProxyRepositoryApiRequest(val *HelmProxyRepositoryApiRequest) *NullableHelmProxyRepositoryApiRequest {
	return &NullableHelmProxyRepositoryApiRequest{value: val, isSet: true}
}

func (v NullableHelmProxyRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmProxyRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


