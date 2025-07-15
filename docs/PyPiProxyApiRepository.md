# PyPiProxyApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Format** | **string** |  | [default to "pypi"]
**HttpClient** | [**HttpClientAttributes**](HttpClientAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**NegativeCache** | [**NegativeCacheAttributes**](NegativeCacheAttributes.md) |  | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Proxy** | [**ProxyAttributes**](ProxyAttributes.md) |  | 
**Pypi** | [**PyPiProxyAttributes**](PyPiProxyAttributes.md) |  | 
**Replication** | Pointer to [**ReplicationAttributes**](ReplicationAttributes.md) |  | [optional] 
**RoutingRuleName** | Pointer to **string** |  | [optional] 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Type** | **string** |  | [default to "pypi"]
**Url** | **string** |  | 

## Methods

### NewPyPiProxyApiRepository

`func NewPyPiProxyApiRepository(format string, httpClient HttpClientAttributes, name string, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, pypi PyPiProxyAttributes, storage StorageAttributes, type_ string, url string, ) *PyPiProxyApiRepository`

NewPyPiProxyApiRepository instantiates a new PyPiProxyApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPyPiProxyApiRepositoryWithDefaults

`func NewPyPiProxyApiRepositoryWithDefaults() *PyPiProxyApiRepository`

NewPyPiProxyApiRepositoryWithDefaults instantiates a new PyPiProxyApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *PyPiProxyApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *PyPiProxyApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *PyPiProxyApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *PyPiProxyApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetFormat

`func (o *PyPiProxyApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PyPiProxyApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PyPiProxyApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetHttpClient

`func (o *PyPiProxyApiRepository) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *PyPiProxyApiRepository) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *PyPiProxyApiRepository) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetName

`func (o *PyPiProxyApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PyPiProxyApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PyPiProxyApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetNegativeCache

`func (o *PyPiProxyApiRepository) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *PyPiProxyApiRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *PyPiProxyApiRepository) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetOnline

`func (o *PyPiProxyApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *PyPiProxyApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *PyPiProxyApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetProxy

`func (o *PyPiProxyApiRepository) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PyPiProxyApiRepository) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PyPiProxyApiRepository) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetPypi

`func (o *PyPiProxyApiRepository) GetPypi() PyPiProxyAttributes`

GetPypi returns the Pypi field if non-nil, zero value otherwise.

### GetPypiOk

`func (o *PyPiProxyApiRepository) GetPypiOk() (*PyPiProxyAttributes, bool)`

GetPypiOk returns a tuple with the Pypi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPypi

`func (o *PyPiProxyApiRepository) SetPypi(v PyPiProxyAttributes)`

SetPypi sets Pypi field to given value.


### GetReplication

`func (o *PyPiProxyApiRepository) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *PyPiProxyApiRepository) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *PyPiProxyApiRepository) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *PyPiProxyApiRepository) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetRoutingRuleName

`func (o *PyPiProxyApiRepository) GetRoutingRuleName() string`

GetRoutingRuleName returns the RoutingRuleName field if non-nil, zero value otherwise.

### GetRoutingRuleNameOk

`func (o *PyPiProxyApiRepository) GetRoutingRuleNameOk() (*string, bool)`

GetRoutingRuleNameOk returns a tuple with the RoutingRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRuleName

`func (o *PyPiProxyApiRepository) SetRoutingRuleName(v string)`

SetRoutingRuleName sets RoutingRuleName field to given value.

### HasRoutingRuleName

`func (o *PyPiProxyApiRepository) HasRoutingRuleName() bool`

HasRoutingRuleName returns a boolean if a field has been set.

### GetStorage

`func (o *PyPiProxyApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PyPiProxyApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PyPiProxyApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *PyPiProxyApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PyPiProxyApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PyPiProxyApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *PyPiProxyApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PyPiProxyApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PyPiProxyApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


