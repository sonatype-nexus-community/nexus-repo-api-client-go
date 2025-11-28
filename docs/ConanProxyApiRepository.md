# ConanProxyApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**ConanProxy** | Pointer to [**ConanProxyAttributes**](ConanProxyAttributes.md) |  | [optional] 
**HttpClient** | [**HttpClientAttributes**](HttpClientAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**NegativeCache** | [**NegativeCacheAttributes**](NegativeCacheAttributes.md) |  | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Proxy** | [**ProxyAttributes**](ProxyAttributes.md) |  | 
**Replication** | Pointer to [**ReplicationAttributes**](ReplicationAttributes.md) |  | [optional] 
**RoutingRule** | Pointer to **string** |  | [optional] 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Format** | **string** |  | [default to "conan"]
**Type** | **string** |  | [default to "proxy"]
**Url** | **string** |  | 

## Methods

### NewConanProxyApiRepository

`func NewConanProxyApiRepository(httpClient HttpClientAttributes, name string, negativeCache NegativeCacheAttributes, online bool, proxy ProxyAttributes, storage StorageAttributes, format string, type_ string, url string, ) *ConanProxyApiRepository`

NewConanProxyApiRepository instantiates a new ConanProxyApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConanProxyApiRepositoryWithDefaults

`func NewConanProxyApiRepositoryWithDefaults() *ConanProxyApiRepository`

NewConanProxyApiRepositoryWithDefaults instantiates a new ConanProxyApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *ConanProxyApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *ConanProxyApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *ConanProxyApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *ConanProxyApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetConanProxy

`func (o *ConanProxyApiRepository) GetConanProxy() ConanProxyAttributes`

GetConanProxy returns the ConanProxy field if non-nil, zero value otherwise.

### GetConanProxyOk

`func (o *ConanProxyApiRepository) GetConanProxyOk() (*ConanProxyAttributes, bool)`

GetConanProxyOk returns a tuple with the ConanProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConanProxy

`func (o *ConanProxyApiRepository) SetConanProxy(v ConanProxyAttributes)`

SetConanProxy sets ConanProxy field to given value.

### HasConanProxy

`func (o *ConanProxyApiRepository) HasConanProxy() bool`

HasConanProxy returns a boolean if a field has been set.

### GetHttpClient

`func (o *ConanProxyApiRepository) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *ConanProxyApiRepository) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *ConanProxyApiRepository) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetName

`func (o *ConanProxyApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConanProxyApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConanProxyApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetNegativeCache

`func (o *ConanProxyApiRepository) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *ConanProxyApiRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *ConanProxyApiRepository) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetOnline

`func (o *ConanProxyApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ConanProxyApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ConanProxyApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetProxy

`func (o *ConanProxyApiRepository) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ConanProxyApiRepository) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ConanProxyApiRepository) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetReplication

`func (o *ConanProxyApiRepository) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *ConanProxyApiRepository) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *ConanProxyApiRepository) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *ConanProxyApiRepository) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetRoutingRule

`func (o *ConanProxyApiRepository) GetRoutingRule() string`

GetRoutingRule returns the RoutingRule field if non-nil, zero value otherwise.

### GetRoutingRuleOk

`func (o *ConanProxyApiRepository) GetRoutingRuleOk() (*string, bool)`

GetRoutingRuleOk returns a tuple with the RoutingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRule

`func (o *ConanProxyApiRepository) SetRoutingRule(v string)`

SetRoutingRule sets RoutingRule field to given value.

### HasRoutingRule

`func (o *ConanProxyApiRepository) HasRoutingRule() bool`

HasRoutingRule returns a boolean if a field has been set.

### GetStorage

`func (o *ConanProxyApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ConanProxyApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ConanProxyApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetFormat

`func (o *ConanProxyApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ConanProxyApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ConanProxyApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetType

`func (o *ConanProxyApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConanProxyApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConanProxyApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *ConanProxyApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConanProxyApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConanProxyApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


