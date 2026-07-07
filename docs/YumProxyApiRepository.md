# YumProxyApiRepository

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
**YumSigning** | Pointer to [**YumSigningRepositoriesAttributes**](YumSigningRepositoriesAttributes.md) |  | [optional] 
**Format** | **string** |  | [default to "yum"]
**RoutingRuleName** | Pointer to **string** | The name of the routing rule assigned to this repository | [optional] 
**Type** | **string** |  | [default to "proxy"]
**Url** | **string** |  | 

## Methods

### NewYumProxyApiRepository

`func NewYumProxyApiRepository(httpClient HttpClientAttributes, name string, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes, format string, type_ string, url string, ) *YumProxyApiRepository`

NewYumProxyApiRepository instantiates a new YumProxyApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYumProxyApiRepositoryWithDefaults

`func NewYumProxyApiRepositoryWithDefaults() *YumProxyApiRepository`

NewYumProxyApiRepositoryWithDefaults instantiates a new YumProxyApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *YumProxyApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *YumProxyApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *YumProxyApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *YumProxyApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetHttpClient

`func (o *YumProxyApiRepository) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *YumProxyApiRepository) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *YumProxyApiRepository) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetName

`func (o *YumProxyApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *YumProxyApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *YumProxyApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetNegativeCache

`func (o *YumProxyApiRepository) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *YumProxyApiRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *YumProxyApiRepository) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetOnline

`func (o *YumProxyApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *YumProxyApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *YumProxyApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetProxy

`func (o *YumProxyApiRepository) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *YumProxyApiRepository) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *YumProxyApiRepository) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetReplication

`func (o *YumProxyApiRepository) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *YumProxyApiRepository) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *YumProxyApiRepository) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *YumProxyApiRepository) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetRoutingRule

`func (o *YumProxyApiRepository) GetRoutingRule() string`

GetRoutingRule returns the RoutingRule field if non-nil, zero value otherwise.

### GetRoutingRuleOk

`func (o *YumProxyApiRepository) GetRoutingRuleOk() (*string, bool)`

GetRoutingRuleOk returns a tuple with the RoutingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRule

`func (o *YumProxyApiRepository) SetRoutingRule(v string)`

SetRoutingRule sets RoutingRule field to given value.

### HasRoutingRule

`func (o *YumProxyApiRepository) HasRoutingRule() bool`

HasRoutingRule returns a boolean if a field has been set.

### GetStorage

`func (o *YumProxyApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *YumProxyApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *YumProxyApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetYumSigning

`func (o *YumProxyApiRepository) GetYumSigning() YumSigningRepositoriesAttributes`

GetYumSigning returns the YumSigning field if non-nil, zero value otherwise.

### GetYumSigningOk

`func (o *YumProxyApiRepository) GetYumSigningOk() (*YumSigningRepositoriesAttributes, bool)`

GetYumSigningOk returns a tuple with the YumSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumSigning

`func (o *YumProxyApiRepository) SetYumSigning(v YumSigningRepositoriesAttributes)`

SetYumSigning sets YumSigning field to given value.

### HasYumSigning

`func (o *YumProxyApiRepository) HasYumSigning() bool`

HasYumSigning returns a boolean if a field has been set.

### GetFormat

`func (o *YumProxyApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *YumProxyApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *YumProxyApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetRoutingRuleName

`func (o *YumProxyApiRepository) GetRoutingRuleName() string`

GetRoutingRuleName returns the RoutingRuleName field if non-nil, zero value otherwise.

### GetRoutingRuleNameOk

`func (o *YumProxyApiRepository) GetRoutingRuleNameOk() (*string, bool)`

GetRoutingRuleNameOk returns a tuple with the RoutingRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRuleName

`func (o *YumProxyApiRepository) SetRoutingRuleName(v string)`

SetRoutingRuleName sets RoutingRuleName field to given value.

### HasRoutingRuleName

`func (o *YumProxyApiRepository) HasRoutingRuleName() bool`

HasRoutingRuleName returns a boolean if a field has been set.

### GetType

`func (o *YumProxyApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *YumProxyApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *YumProxyApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *YumProxyApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *YumProxyApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *YumProxyApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


