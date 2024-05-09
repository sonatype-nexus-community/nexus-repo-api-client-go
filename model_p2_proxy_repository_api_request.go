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

// checks if the P2ProxyRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &P2ProxyRepositoryApiRequest{}

// P2ProxyRepositoryApiRequest struct for P2ProxyRepositoryApiRequest
type P2ProxyRepositoryApiRequest struct {
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient HttpClientAttributes `json:"httpClient"`
	// A unique identifier for this repository
	Name string `json:"name"`
	NegativeCache NegativeCacheAttributes `json:"negativeCache"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Proxy ProxyAttributes `json:"proxy"`
	Replication *ReplicationAttributes `json:"replication,omitempty"`
	RoutingRule *string `json:"routingRule,omitempty"`
	Storage StorageAttributes `json:"storage"`
}

type _P2ProxyRepositoryApiRequest P2ProxyRepositoryApiRequest

// NewP2ProxyRepositoryApiRequest instantiates a new P2ProxyRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewP2ProxyRepositoryApiRequest(httpClient HttpClientAttributes, name string, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes) *P2ProxyRepositoryApiRequest {
	this := P2ProxyRepositoryApiRequest{}
	this.HttpClient = httpClient
	this.Name = name
	this.NegativeCache = negativeCache
	this.Online = online
	this.Proxy = proxy
	this.Storage = storage
	return &this
}

// NewP2ProxyRepositoryApiRequestWithDefaults instantiates a new P2ProxyRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewP2ProxyRepositoryApiRequestWithDefaults() *P2ProxyRepositoryApiRequest {
	this := P2ProxyRepositoryApiRequest{}
	return &this
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *P2ProxyRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P2ProxyRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *P2ProxyRepositoryApiRequest) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *P2ProxyRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetHttpClient returns the HttpClient field value
func (o *P2ProxyRepositoryApiRequest) GetHttpClient() HttpClientAttributes {
	if o == nil {
		var ret HttpClientAttributes
		return ret
	}

	return o.HttpClient
}

// GetHttpClientOk returns a tuple with the HttpClient field value
// and a boolean to check if the value has been set.
func (o *P2ProxyRepositoryApiRequest) GetHttpClientOk() (*HttpClientAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpClient, true
}

// SetHttpClient sets field value
func (o *P2ProxyRepositoryApiRequest) SetHttpClient(v HttpClientAttributes) {
	o.HttpClient = v
}

// GetName returns the Name field value
func (o *P2ProxyRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *P2ProxyRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *P2ProxyRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetNegativeCache returns the NegativeCache field value
func (o *P2ProxyRepositoryApiRequest) GetNegativeCache() NegativeCacheAttributes {
	if o == nil {
		var ret NegativeCacheAttributes
		return ret
	}

	return o.NegativeCache
}

// GetNegativeCacheOk returns a tuple with the NegativeCache field value
// and a boolean to check if the value has been set.
func (o *P2ProxyRepositoryApiRequest) GetNegativeCacheOk() (*NegativeCacheAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeCache, true
}

// SetNegativeCache sets field value
func (o *P2ProxyRepositoryApiRequest) SetNegativeCache(v NegativeCacheAttributes) {
	o.NegativeCache = v
}

// GetOnline returns the Online field value
func (o *P2ProxyRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *P2ProxyRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *P2ProxyRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetProxy returns the Proxy field value
func (o *P2ProxyRepositoryApiRequest) GetProxy() ProxyAttributes {
	if o == nil {
		var ret ProxyAttributes
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *P2ProxyRepositoryApiRequest) GetProxyOk() (*ProxyAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *P2ProxyRepositoryApiRequest) SetProxy(v ProxyAttributes) {
	o.Proxy = v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *P2ProxyRepositoryApiRequest) GetReplication() ReplicationAttributes {
	if o == nil || IsNil(o.Replication) {
		var ret ReplicationAttributes
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P2ProxyRepositoryApiRequest) GetReplicationOk() (*ReplicationAttributes, bool) {
	if o == nil || IsNil(o.Replication) {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *P2ProxyRepositoryApiRequest) HasReplication() bool {
	if o != nil && !IsNil(o.Replication) {
		return true
	}

	return false
}

// SetReplication gets a reference to the given ReplicationAttributes and assigns it to the Replication field.
func (o *P2ProxyRepositoryApiRequest) SetReplication(v ReplicationAttributes) {
	o.Replication = &v
}

// GetRoutingRule returns the RoutingRule field value if set, zero value otherwise.
func (o *P2ProxyRepositoryApiRequest) GetRoutingRule() string {
	if o == nil || IsNil(o.RoutingRule) {
		var ret string
		return ret
	}
	return *o.RoutingRule
}

// GetRoutingRuleOk returns a tuple with the RoutingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P2ProxyRepositoryApiRequest) GetRoutingRuleOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingRule) {
		return nil, false
	}
	return o.RoutingRule, true
}

// HasRoutingRule returns a boolean if a field has been set.
func (o *P2ProxyRepositoryApiRequest) HasRoutingRule() bool {
	if o != nil && !IsNil(o.RoutingRule) {
		return true
	}

	return false
}

// SetRoutingRule gets a reference to the given string and assigns it to the RoutingRule field.
func (o *P2ProxyRepositoryApiRequest) SetRoutingRule(v string) {
	o.RoutingRule = &v
}

// GetStorage returns the Storage field value
func (o *P2ProxyRepositoryApiRequest) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *P2ProxyRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *P2ProxyRepositoryApiRequest) SetStorage(v StorageAttributes) {
	o.Storage = v
}

func (o P2ProxyRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o P2ProxyRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
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

func (o *P2ProxyRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
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

	varP2ProxyRepositoryApiRequest := _P2ProxyRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varP2ProxyRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = P2ProxyRepositoryApiRequest(varP2ProxyRepositoryApiRequest)

	return err
}

type NullableP2ProxyRepositoryApiRequest struct {
	value *P2ProxyRepositoryApiRequest
	isSet bool
}

func (v NullableP2ProxyRepositoryApiRequest) Get() *P2ProxyRepositoryApiRequest {
	return v.value
}

func (v *NullableP2ProxyRepositoryApiRequest) Set(val *P2ProxyRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableP2ProxyRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableP2ProxyRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableP2ProxyRepositoryApiRequest(val *P2ProxyRepositoryApiRequest) *NullableP2ProxyRepositoryApiRequest {
	return &NullableP2ProxyRepositoryApiRequest{value: val, isSet: true}
}

func (v NullableP2ProxyRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableP2ProxyRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


