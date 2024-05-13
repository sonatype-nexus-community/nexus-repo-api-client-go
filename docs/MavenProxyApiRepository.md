# MavenProxyApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**HttpClient** | [**HttpClientAttributes**](HttpClientAttributes.md) |  | 
**Maven** | [**MavenAttributes**](MavenAttributes.md) |  | 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**NegativeCache** | [**NegativeCacheAttributes**](NegativeCacheAttributes.md) |  | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Proxy** | [**ProxyAttributes**](ProxyAttributes.md) |  | 
**Replication** | Pointer to [**ReplicationAttributes**](ReplicationAttributes.md) |  | [optional] 
**RoutingRuleName** | Pointer to **string** | The name of the routing rule assigned to this repository | [optional] 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewMavenProxyApiRepository

`func NewMavenProxyApiRepository(httpClient HttpClientAttributes, maven MavenAttributes, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes, ) *MavenProxyApiRepository`

NewMavenProxyApiRepository instantiates a new MavenProxyApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenProxyApiRepositoryWithDefaults

`func NewMavenProxyApiRepositoryWithDefaults() *MavenProxyApiRepository`

NewMavenProxyApiRepositoryWithDefaults instantiates a new MavenProxyApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *MavenProxyApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *MavenProxyApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *MavenProxyApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *MavenProxyApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetHttpClient

`func (o *MavenProxyApiRepository) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *MavenProxyApiRepository) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *MavenProxyApiRepository) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetMaven

`func (o *MavenProxyApiRepository) GetMaven() MavenAttributes`

GetMaven returns the Maven field if non-nil, zero value otherwise.

### GetMavenOk

`func (o *MavenProxyApiRepository) GetMavenOk() (*MavenAttributes, bool)`

GetMavenOk returns a tuple with the Maven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven

`func (o *MavenProxyApiRepository) SetMaven(v MavenAttributes)`

SetMaven sets Maven field to given value.


### GetName

`func (o *MavenProxyApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MavenProxyApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MavenProxyApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MavenProxyApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNegativeCache

`func (o *MavenProxyApiRepository) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *MavenProxyApiRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *MavenProxyApiRepository) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetOnline

`func (o *MavenProxyApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *MavenProxyApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *MavenProxyApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetProxy

`func (o *MavenProxyApiRepository) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *MavenProxyApiRepository) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *MavenProxyApiRepository) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetReplication

`func (o *MavenProxyApiRepository) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *MavenProxyApiRepository) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *MavenProxyApiRepository) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *MavenProxyApiRepository) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetRoutingRuleName

`func (o *MavenProxyApiRepository) GetRoutingRuleName() string`

GetRoutingRuleName returns the RoutingRuleName field if non-nil, zero value otherwise.

### GetRoutingRuleNameOk

`func (o *MavenProxyApiRepository) GetRoutingRuleNameOk() (*string, bool)`

GetRoutingRuleNameOk returns a tuple with the RoutingRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRuleName

`func (o *MavenProxyApiRepository) SetRoutingRuleName(v string)`

SetRoutingRuleName sets RoutingRuleName field to given value.

### HasRoutingRuleName

`func (o *MavenProxyApiRepository) HasRoutingRuleName() bool`

HasRoutingRuleName returns a boolean if a field has been set.

### GetStorage

`func (o *MavenProxyApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *MavenProxyApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *MavenProxyApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetUrl

`func (o *MavenProxyApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MavenProxyApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MavenProxyApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MavenProxyApiRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


