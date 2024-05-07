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

// checks if the DockerProxyApiRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DockerProxyApiRepository{}

// DockerProxyApiRepository struct for DockerProxyApiRepository
type DockerProxyApiRepository struct {
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Docker DockerAttributes `json:"docker"`
	DockerProxy DockerProxyAttributes `json:"dockerProxy"`
	HttpClient HttpClientAttributes `json:"httpClient"`
	// A unique identifier for this repository
	Name *string `json:"name,omitempty"`
	NegativeCache NegativeCacheAttributes `json:"negativeCache"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Proxy ProxyAttributes `json:"proxy"`
	Replication *ReplicationAttributes `json:"replication,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`
	Storage StorageAttributes `json:"storage"`
}

type _DockerProxyApiRepository DockerProxyApiRepository

// NewDockerProxyApiRepository instantiates a new DockerProxyApiRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerProxyApiRepository(docker DockerAttributes, dockerProxy DockerProxyAttributes, httpClient HttpClientAttributes, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes) *DockerProxyApiRepository {
	this := DockerProxyApiRepository{}
	this.Docker = docker
	this.DockerProxy = dockerProxy
	this.HttpClient = httpClient
	this.NegativeCache = negativeCache
	this.Online = online
	this.Proxy = proxy
	this.Storage = storage
	return &this
}

// NewDockerProxyApiRepositoryWithDefaults instantiates a new DockerProxyApiRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerProxyApiRepositoryWithDefaults() *DockerProxyApiRepository {
	this := DockerProxyApiRepository{}
	return &this
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *DockerProxyApiRepository) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *DockerProxyApiRepository) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *DockerProxyApiRepository) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetDocker returns the Docker field value
func (o *DockerProxyApiRepository) GetDocker() DockerAttributes {
	if o == nil {
		var ret DockerAttributes
		return ret
	}

	return o.Docker
}

// GetDockerOk returns a tuple with the Docker field value
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetDockerOk() (*DockerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Docker, true
}

// SetDocker sets field value
func (o *DockerProxyApiRepository) SetDocker(v DockerAttributes) {
	o.Docker = v
}

// GetDockerProxy returns the DockerProxy field value
func (o *DockerProxyApiRepository) GetDockerProxy() DockerProxyAttributes {
	if o == nil {
		var ret DockerProxyAttributes
		return ret
	}

	return o.DockerProxy
}

// GetDockerProxyOk returns a tuple with the DockerProxy field value
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetDockerProxyOk() (*DockerProxyAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerProxy, true
}

// SetDockerProxy sets field value
func (o *DockerProxyApiRepository) SetDockerProxy(v DockerProxyAttributes) {
	o.DockerProxy = v
}

// GetHttpClient returns the HttpClient field value
func (o *DockerProxyApiRepository) GetHttpClient() HttpClientAttributes {
	if o == nil {
		var ret HttpClientAttributes
		return ret
	}

	return o.HttpClient
}

// GetHttpClientOk returns a tuple with the HttpClient field value
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetHttpClientOk() (*HttpClientAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpClient, true
}

// SetHttpClient sets field value
func (o *DockerProxyApiRepository) SetHttpClient(v HttpClientAttributes) {
	o.HttpClient = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DockerProxyApiRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DockerProxyApiRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DockerProxyApiRepository) SetName(v string) {
	o.Name = &v
}

// GetNegativeCache returns the NegativeCache field value
func (o *DockerProxyApiRepository) GetNegativeCache() NegativeCacheAttributes {
	if o == nil {
		var ret NegativeCacheAttributes
		return ret
	}

	return o.NegativeCache
}

// GetNegativeCacheOk returns a tuple with the NegativeCache field value
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeCache, true
}

// SetNegativeCache sets field value
func (o *DockerProxyApiRepository) SetNegativeCache(v NegativeCacheAttributes) {
	o.NegativeCache = v
}

// GetOnline returns the Online field value
func (o *DockerProxyApiRepository) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *DockerProxyApiRepository) SetOnline(v bool) {
	o.Online = v
}

// GetProxy returns the Proxy field value
func (o *DockerProxyApiRepository) GetProxy() ProxyAttributes {
	if o == nil {
		var ret ProxyAttributes
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetProxyOk() (*ProxyAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *DockerProxyApiRepository) SetProxy(v ProxyAttributes) {
	o.Proxy = v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *DockerProxyApiRepository) GetReplication() ReplicationAttributes {
	if o == nil || IsNil(o.Replication) {
		var ret ReplicationAttributes
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetReplicationOk() (*ReplicationAttributes, bool) {
	if o == nil || IsNil(o.Replication) {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *DockerProxyApiRepository) HasReplication() bool {
	if o != nil && !IsNil(o.Replication) {
		return true
	}

	return false
}

// SetReplication gets a reference to the given ReplicationAttributes and assigns it to the Replication field.
func (o *DockerProxyApiRepository) SetReplication(v ReplicationAttributes) {
	o.Replication = &v
}

// GetRoutingRuleName returns the RoutingRuleName field value if set, zero value otherwise.
func (o *DockerProxyApiRepository) GetRoutingRuleName() string {
	if o == nil || IsNil(o.RoutingRuleName) {
		var ret string
		return ret
	}
	return *o.RoutingRuleName
}

// GetRoutingRuleNameOk returns a tuple with the RoutingRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetRoutingRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingRuleName) {
		return nil, false
	}
	return o.RoutingRuleName, true
}

// HasRoutingRuleName returns a boolean if a field has been set.
func (o *DockerProxyApiRepository) HasRoutingRuleName() bool {
	if o != nil && !IsNil(o.RoutingRuleName) {
		return true
	}

	return false
}

// SetRoutingRuleName gets a reference to the given string and assigns it to the RoutingRuleName field.
func (o *DockerProxyApiRepository) SetRoutingRuleName(v string) {
	o.RoutingRuleName = &v
}

// GetStorage returns the Storage field value
func (o *DockerProxyApiRepository) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *DockerProxyApiRepository) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *DockerProxyApiRepository) SetStorage(v StorageAttributes) {
	o.Storage = v
}

func (o DockerProxyApiRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DockerProxyApiRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	toSerialize["docker"] = o.Docker
	toSerialize["dockerProxy"] = o.DockerProxy
	toSerialize["httpClient"] = o.HttpClient
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["negativeCache"] = o.NegativeCache
	toSerialize["online"] = o.Online
	toSerialize["proxy"] = o.Proxy
	if !IsNil(o.Replication) {
		toSerialize["replication"] = o.Replication
	}
	if !IsNil(o.RoutingRuleName) {
		toSerialize["routingRuleName"] = o.RoutingRuleName
	}
	toSerialize["storage"] = o.Storage
	return toSerialize, nil
}

func (o *DockerProxyApiRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"docker",
		"dockerProxy",
		"httpClient",
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

	varDockerProxyApiRepository := _DockerProxyApiRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDockerProxyApiRepository)

	if err != nil {
		return err
	}

	*o = DockerProxyApiRepository(varDockerProxyApiRepository)

	return err
}

type NullableDockerProxyApiRepository struct {
	value *DockerProxyApiRepository
	isSet bool
}

func (v NullableDockerProxyApiRepository) Get() *DockerProxyApiRepository {
	return v.value
}

func (v *NullableDockerProxyApiRepository) Set(val *DockerProxyApiRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerProxyApiRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerProxyApiRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerProxyApiRepository(val *DockerProxyApiRepository) *NullableDockerProxyApiRepository {
	return &NullableDockerProxyApiRepository{value: val, isSet: true}
}

func (v NullableDockerProxyApiRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerProxyApiRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


