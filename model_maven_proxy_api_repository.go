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

// checks if the MavenProxyApiRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MavenProxyApiRepository{}

// MavenProxyApiRepository struct for MavenProxyApiRepository
type MavenProxyApiRepository struct {
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Format *string `json:"format,omitempty"`
	HttpClient HttpClientAttributes `json:"httpClient"`
	Maven MavenAttributes `json:"maven"`
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
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
}

type _MavenProxyApiRepository MavenProxyApiRepository

// NewMavenProxyApiRepository instantiates a new MavenProxyApiRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMavenProxyApiRepository(httpClient HttpClientAttributes, maven MavenAttributes, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes) *MavenProxyApiRepository {
	this := MavenProxyApiRepository{}
	var format string = "maven2"
	this.Format = &format
	this.HttpClient = httpClient
	this.Maven = maven
	this.NegativeCache = negativeCache
	this.Online = online
	this.Proxy = proxy
	this.Storage = storage
	var type_ string = "hosted"
	this.Type = &type_
	return &this
}

// NewMavenProxyApiRepositoryWithDefaults instantiates a new MavenProxyApiRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMavenProxyApiRepositoryWithDefaults() *MavenProxyApiRepository {
	this := MavenProxyApiRepository{}
	var format string = "maven2"
	this.Format = &format
	var type_ string = "hosted"
	this.Type = &type_
	return &this
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *MavenProxyApiRepository) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *MavenProxyApiRepository) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *MavenProxyApiRepository) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *MavenProxyApiRepository) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MavenProxyApiRepository) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *MavenProxyApiRepository) SetFormat(v string) {
	o.Format = &v
}

// GetHttpClient returns the HttpClient field value
func (o *MavenProxyApiRepository) GetHttpClient() HttpClientAttributes {
	if o == nil {
		var ret HttpClientAttributes
		return ret
	}

	return o.HttpClient
}

// GetHttpClientOk returns a tuple with the HttpClient field value
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetHttpClientOk() (*HttpClientAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpClient, true
}

// SetHttpClient sets field value
func (o *MavenProxyApiRepository) SetHttpClient(v HttpClientAttributes) {
	o.HttpClient = v
}

// GetMaven returns the Maven field value
func (o *MavenProxyApiRepository) GetMaven() MavenAttributes {
	if o == nil {
		var ret MavenAttributes
		return ret
	}

	return o.Maven
}

// GetMavenOk returns a tuple with the Maven field value
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetMavenOk() (*MavenAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maven, true
}

// SetMaven sets field value
func (o *MavenProxyApiRepository) SetMaven(v MavenAttributes) {
	o.Maven = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MavenProxyApiRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MavenProxyApiRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MavenProxyApiRepository) SetName(v string) {
	o.Name = &v
}

// GetNegativeCache returns the NegativeCache field value
func (o *MavenProxyApiRepository) GetNegativeCache() NegativeCacheAttributes {
	if o == nil {
		var ret NegativeCacheAttributes
		return ret
	}

	return o.NegativeCache
}

// GetNegativeCacheOk returns a tuple with the NegativeCache field value
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeCache, true
}

// SetNegativeCache sets field value
func (o *MavenProxyApiRepository) SetNegativeCache(v NegativeCacheAttributes) {
	o.NegativeCache = v
}

// GetOnline returns the Online field value
func (o *MavenProxyApiRepository) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *MavenProxyApiRepository) SetOnline(v bool) {
	o.Online = v
}

// GetProxy returns the Proxy field value
func (o *MavenProxyApiRepository) GetProxy() ProxyAttributes {
	if o == nil {
		var ret ProxyAttributes
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetProxyOk() (*ProxyAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *MavenProxyApiRepository) SetProxy(v ProxyAttributes) {
	o.Proxy = v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *MavenProxyApiRepository) GetReplication() ReplicationAttributes {
	if o == nil || IsNil(o.Replication) {
		var ret ReplicationAttributes
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetReplicationOk() (*ReplicationAttributes, bool) {
	if o == nil || IsNil(o.Replication) {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *MavenProxyApiRepository) HasReplication() bool {
	if o != nil && !IsNil(o.Replication) {
		return true
	}

	return false
}

// SetReplication gets a reference to the given ReplicationAttributes and assigns it to the Replication field.
func (o *MavenProxyApiRepository) SetReplication(v ReplicationAttributes) {
	o.Replication = &v
}

// GetRoutingRuleName returns the RoutingRuleName field value if set, zero value otherwise.
func (o *MavenProxyApiRepository) GetRoutingRuleName() string {
	if o == nil || IsNil(o.RoutingRuleName) {
		var ret string
		return ret
	}
	return *o.RoutingRuleName
}

// GetRoutingRuleNameOk returns a tuple with the RoutingRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetRoutingRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingRuleName) {
		return nil, false
	}
	return o.RoutingRuleName, true
}

// HasRoutingRuleName returns a boolean if a field has been set.
func (o *MavenProxyApiRepository) HasRoutingRuleName() bool {
	if o != nil && !IsNil(o.RoutingRuleName) {
		return true
	}

	return false
}

// SetRoutingRuleName gets a reference to the given string and assigns it to the RoutingRuleName field.
func (o *MavenProxyApiRepository) SetRoutingRuleName(v string) {
	o.RoutingRuleName = &v
}

// GetStorage returns the Storage field value
func (o *MavenProxyApiRepository) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *MavenProxyApiRepository) SetStorage(v StorageAttributes) {
	o.Storage = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MavenProxyApiRepository) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MavenProxyApiRepository) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MavenProxyApiRepository) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MavenProxyApiRepository) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenProxyApiRepository) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MavenProxyApiRepository) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MavenProxyApiRepository) SetUrl(v string) {
	o.Url = &v
}

func (o MavenProxyApiRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MavenProxyApiRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	toSerialize["httpClient"] = o.HttpClient
	toSerialize["maven"] = o.Maven
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

func (o *MavenProxyApiRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"httpClient",
		"maven",
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

	varMavenProxyApiRepository := _MavenProxyApiRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMavenProxyApiRepository)

	if err != nil {
		return err
	}

	*o = MavenProxyApiRepository(varMavenProxyApiRepository)

	return err
}

type NullableMavenProxyApiRepository struct {
	value *MavenProxyApiRepository
	isSet bool
}

func (v NullableMavenProxyApiRepository) Get() *MavenProxyApiRepository {
	return v.value
}

func (v *NullableMavenProxyApiRepository) Set(val *MavenProxyApiRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableMavenProxyApiRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableMavenProxyApiRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMavenProxyApiRepository(val *MavenProxyApiRepository) *NullableMavenProxyApiRepository {
	return &NullableMavenProxyApiRepository{value: val, isSet: true}
}

func (v NullableMavenProxyApiRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMavenProxyApiRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


