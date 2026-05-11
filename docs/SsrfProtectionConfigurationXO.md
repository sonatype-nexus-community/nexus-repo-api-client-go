# SsrfProtectionConfigurationXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDomains** | Pointer to **[]string** | List of domain names allowed to bypass SSRF protection | [optional] 
**AllowedIPs** | Pointer to **[]string** | List of IP addresses allowed to bypass SSRF protection | [optional] 
**Enabled** | **bool** | Whether SSRF protection is enabled | 

## Methods

### NewSsrfProtectionConfigurationXO

`func NewSsrfProtectionConfigurationXO(enabled bool, ) *SsrfProtectionConfigurationXO`

NewSsrfProtectionConfigurationXO instantiates a new SsrfProtectionConfigurationXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsrfProtectionConfigurationXOWithDefaults

`func NewSsrfProtectionConfigurationXOWithDefaults() *SsrfProtectionConfigurationXO`

NewSsrfProtectionConfigurationXOWithDefaults instantiates a new SsrfProtectionConfigurationXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDomains

`func (o *SsrfProtectionConfigurationXO) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *SsrfProtectionConfigurationXO) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *SsrfProtectionConfigurationXO) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *SsrfProtectionConfigurationXO) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAllowedIPs

`func (o *SsrfProtectionConfigurationXO) GetAllowedIPs() []string`

GetAllowedIPs returns the AllowedIPs field if non-nil, zero value otherwise.

### GetAllowedIPsOk

`func (o *SsrfProtectionConfigurationXO) GetAllowedIPsOk() (*[]string, bool)`

GetAllowedIPsOk returns a tuple with the AllowedIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIPs

`func (o *SsrfProtectionConfigurationXO) SetAllowedIPs(v []string)`

SetAllowedIPs sets AllowedIPs field to given value.

### HasAllowedIPs

`func (o *SsrfProtectionConfigurationXO) HasAllowedIPs() bool`

HasAllowedIPs returns a boolean if a field has been set.

### GetEnabled

`func (o *SsrfProtectionConfigurationXO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SsrfProtectionConfigurationXO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SsrfProtectionConfigurationXO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


