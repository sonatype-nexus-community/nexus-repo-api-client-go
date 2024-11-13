# HttpSettingsXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpProxy** | [**ProxySettingsXo**](ProxySettingsXo.md) |  | 
**HttpsProxy** | [**ProxySettingsXo**](ProxySettingsXo.md) |  | 
**NonProxyHosts** | Pointer to **[]string** |  | [optional] 
**Retries** | **int32** | Connection/Socket Retry Attempts | 
**Timeout** | **int32** | Connection/Socket Timeout | 
**UserAgent** | **string** | User-Agent Customization | 

## Methods

### NewHttpSettingsXo

`func NewHttpSettingsXo(httpProxy ProxySettingsXo, httpsProxy ProxySettingsXo, retries int32, timeout int32, userAgent string, ) *HttpSettingsXo`

NewHttpSettingsXo instantiates a new HttpSettingsXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpSettingsXoWithDefaults

`func NewHttpSettingsXoWithDefaults() *HttpSettingsXo`

NewHttpSettingsXoWithDefaults instantiates a new HttpSettingsXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpProxy

`func (o *HttpSettingsXo) GetHttpProxy() ProxySettingsXo`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *HttpSettingsXo) GetHttpProxyOk() (*ProxySettingsXo, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *HttpSettingsXo) SetHttpProxy(v ProxySettingsXo)`

SetHttpProxy sets HttpProxy field to given value.


### GetHttpsProxy

`func (o *HttpSettingsXo) GetHttpsProxy() ProxySettingsXo`

GetHttpsProxy returns the HttpsProxy field if non-nil, zero value otherwise.

### GetHttpsProxyOk

`func (o *HttpSettingsXo) GetHttpsProxyOk() (*ProxySettingsXo, bool)`

GetHttpsProxyOk returns a tuple with the HttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxy

`func (o *HttpSettingsXo) SetHttpsProxy(v ProxySettingsXo)`

SetHttpsProxy sets HttpsProxy field to given value.


### GetNonProxyHosts

`func (o *HttpSettingsXo) GetNonProxyHosts() []string`

GetNonProxyHosts returns the NonProxyHosts field if non-nil, zero value otherwise.

### GetNonProxyHostsOk

`func (o *HttpSettingsXo) GetNonProxyHostsOk() (*[]string, bool)`

GetNonProxyHostsOk returns a tuple with the NonProxyHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonProxyHosts

`func (o *HttpSettingsXo) SetNonProxyHosts(v []string)`

SetNonProxyHosts sets NonProxyHosts field to given value.

### HasNonProxyHosts

`func (o *HttpSettingsXo) HasNonProxyHosts() bool`

HasNonProxyHosts returns a boolean if a field has been set.

### GetRetries

`func (o *HttpSettingsXo) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *HttpSettingsXo) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *HttpSettingsXo) SetRetries(v int32)`

SetRetries sets Retries field to given value.


### GetTimeout

`func (o *HttpSettingsXo) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HttpSettingsXo) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HttpSettingsXo) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetUserAgent

`func (o *HttpSettingsXo) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *HttpSettingsXo) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *HttpSettingsXo) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


