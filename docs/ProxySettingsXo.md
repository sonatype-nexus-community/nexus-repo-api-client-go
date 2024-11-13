# ProxySettingsXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthInfo** | [**AuthSettingsXo**](AuthSettingsXo.md) |  | 
**Enabled** | **bool** | proxy enabled | 
**Host** | **string** | proxy host | 
**Port** | **string** | proxy port | 

## Methods

### NewProxySettingsXo

`func NewProxySettingsXo(authInfo AuthSettingsXo, enabled bool, host string, port string, ) *ProxySettingsXo`

NewProxySettingsXo instantiates a new ProxySettingsXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxySettingsXoWithDefaults

`func NewProxySettingsXoWithDefaults() *ProxySettingsXo`

NewProxySettingsXoWithDefaults instantiates a new ProxySettingsXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthInfo

`func (o *ProxySettingsXo) GetAuthInfo() AuthSettingsXo`

GetAuthInfo returns the AuthInfo field if non-nil, zero value otherwise.

### GetAuthInfoOk

`func (o *ProxySettingsXo) GetAuthInfoOk() (*AuthSettingsXo, bool)`

GetAuthInfoOk returns a tuple with the AuthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthInfo

`func (o *ProxySettingsXo) SetAuthInfo(v AuthSettingsXo)`

SetAuthInfo sets AuthInfo field to given value.


### GetEnabled

`func (o *ProxySettingsXo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProxySettingsXo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProxySettingsXo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHost

`func (o *ProxySettingsXo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProxySettingsXo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProxySettingsXo) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ProxySettingsXo) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProxySettingsXo) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProxySettingsXo) SetPort(v string)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


