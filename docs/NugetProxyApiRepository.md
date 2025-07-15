# NugetProxyApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Format** | Pointer to **string** |  | [optional] [default to "nuget"]
**HttpClient** | [**HttpClientAttributes**](HttpClientAttributes.md) |  | 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**NegativeCache** | [**NegativeCacheAttributes**](NegativeCacheAttributes.md) |  | 
**NugetProxy** | [**NugetAttributes**](NugetAttributes.md) |  | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Proxy** | [**ProxyAttributes**](ProxyAttributes.md) |  | 
**Replication** | Pointer to [**ReplicationAttributes**](ReplicationAttributes.md) |  | [optional] 
**RoutingRuleName** | Pointer to **string** | The name of the routing rule assigned to this repository | [optional] 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "proxy"]
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewNugetProxyApiRepository

`func NewNugetProxyApiRepository(httpClient HttpClientAttributes, negativeCache NegativeCacheAttributes, nugetProxy NugetAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes, ) *NugetProxyApiRepository`

NewNugetProxyApiRepository instantiates a new NugetProxyApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNugetProxyApiRepositoryWithDefaults

`func NewNugetProxyApiRepositoryWithDefaults() *NugetProxyApiRepository`

NewNugetProxyApiRepositoryWithDefaults instantiates a new NugetProxyApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *NugetProxyApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *NugetProxyApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *NugetProxyApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *NugetProxyApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetFormat

`func (o *NugetProxyApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *NugetProxyApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *NugetProxyApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *NugetProxyApiRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetHttpClient

`func (o *NugetProxyApiRepository) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *NugetProxyApiRepository) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *NugetProxyApiRepository) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetName

`func (o *NugetProxyApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NugetProxyApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NugetProxyApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NugetProxyApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNegativeCache

`func (o *NugetProxyApiRepository) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *NugetProxyApiRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *NugetProxyApiRepository) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetNugetProxy

`func (o *NugetProxyApiRepository) GetNugetProxy() NugetAttributes`

GetNugetProxy returns the NugetProxy field if non-nil, zero value otherwise.

### GetNugetProxyOk

`func (o *NugetProxyApiRepository) GetNugetProxyOk() (*NugetAttributes, bool)`

GetNugetProxyOk returns a tuple with the NugetProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNugetProxy

`func (o *NugetProxyApiRepository) SetNugetProxy(v NugetAttributes)`

SetNugetProxy sets NugetProxy field to given value.


### GetOnline

`func (o *NugetProxyApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *NugetProxyApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *NugetProxyApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetProxy

`func (o *NugetProxyApiRepository) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *NugetProxyApiRepository) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *NugetProxyApiRepository) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetReplication

`func (o *NugetProxyApiRepository) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *NugetProxyApiRepository) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *NugetProxyApiRepository) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *NugetProxyApiRepository) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetRoutingRuleName

`func (o *NugetProxyApiRepository) GetRoutingRuleName() string`

GetRoutingRuleName returns the RoutingRuleName field if non-nil, zero value otherwise.

### GetRoutingRuleNameOk

`func (o *NugetProxyApiRepository) GetRoutingRuleNameOk() (*string, bool)`

GetRoutingRuleNameOk returns a tuple with the RoutingRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRuleName

`func (o *NugetProxyApiRepository) SetRoutingRuleName(v string)`

SetRoutingRuleName sets RoutingRuleName field to given value.

### HasRoutingRuleName

`func (o *NugetProxyApiRepository) HasRoutingRuleName() bool`

HasRoutingRuleName returns a boolean if a field has been set.

### GetStorage

`func (o *NugetProxyApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *NugetProxyApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *NugetProxyApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *NugetProxyApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NugetProxyApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NugetProxyApiRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NugetProxyApiRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *NugetProxyApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NugetProxyApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NugetProxyApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NugetProxyApiRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


