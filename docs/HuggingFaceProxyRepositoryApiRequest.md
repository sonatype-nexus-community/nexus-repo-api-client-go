# HuggingFaceProxyRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**HttpClient** | [**HttpClientAttributes**](HttpClientAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**NegativeCache** | [**NegativeCacheAttributes**](NegativeCacheAttributes.md) |  | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Proxy** | [**ProxyAttributes**](ProxyAttributes.md) |  | 
**Replication** | Pointer to [**ReplicationAttributes**](ReplicationAttributes.md) |  | [optional] 
**RoutingRule** | Pointer to **string** |  | [optional] 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 

## Methods

### NewHuggingFaceProxyRepositoryApiRequest

`func NewHuggingFaceProxyRepositoryApiRequest(httpClient HttpClientAttributes, name string, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes, ) *HuggingFaceProxyRepositoryApiRequest`

NewHuggingFaceProxyRepositoryApiRequest instantiates a new HuggingFaceProxyRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceProxyRepositoryApiRequestWithDefaults

`func NewHuggingFaceProxyRepositoryApiRequestWithDefaults() *HuggingFaceProxyRepositoryApiRequest`

NewHuggingFaceProxyRepositoryApiRequestWithDefaults instantiates a new HuggingFaceProxyRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *HuggingFaceProxyRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *HuggingFaceProxyRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *HuggingFaceProxyRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *HuggingFaceProxyRepositoryApiRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetHttpClient

`func (o *HuggingFaceProxyRepositoryApiRequest) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *HuggingFaceProxyRepositoryApiRequest) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *HuggingFaceProxyRepositoryApiRequest) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetName

`func (o *HuggingFaceProxyRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HuggingFaceProxyRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HuggingFaceProxyRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNegativeCache

`func (o *HuggingFaceProxyRepositoryApiRequest) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *HuggingFaceProxyRepositoryApiRequest) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *HuggingFaceProxyRepositoryApiRequest) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetOnline

`func (o *HuggingFaceProxyRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *HuggingFaceProxyRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *HuggingFaceProxyRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetProxy

`func (o *HuggingFaceProxyRepositoryApiRequest) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *HuggingFaceProxyRepositoryApiRequest) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *HuggingFaceProxyRepositoryApiRequest) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetReplication

`func (o *HuggingFaceProxyRepositoryApiRequest) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *HuggingFaceProxyRepositoryApiRequest) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *HuggingFaceProxyRepositoryApiRequest) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *HuggingFaceProxyRepositoryApiRequest) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetRoutingRule

`func (o *HuggingFaceProxyRepositoryApiRequest) GetRoutingRule() string`

GetRoutingRule returns the RoutingRule field if non-nil, zero value otherwise.

### GetRoutingRuleOk

`func (o *HuggingFaceProxyRepositoryApiRequest) GetRoutingRuleOk() (*string, bool)`

GetRoutingRuleOk returns a tuple with the RoutingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRule

`func (o *HuggingFaceProxyRepositoryApiRequest) SetRoutingRule(v string)`

SetRoutingRule sets RoutingRule field to given value.

### HasRoutingRule

`func (o *HuggingFaceProxyRepositoryApiRequest) HasRoutingRule() bool`

HasRoutingRule returns a boolean if a field has been set.

### GetStorage

`func (o *HuggingFaceProxyRepositoryApiRequest) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *HuggingFaceProxyRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *HuggingFaceProxyRepositoryApiRequest) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


