# IpAllowListSettingsXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveMode** | Pointer to **string** | Effective mode after applying forceDisable and feature flag overrides. May differ from &#39;mode&#39; when the feature is force-disabled or the feature flag is off. | [optional] 
**FeatureEnabled** | Pointer to **bool** | Whether the IP Allow List feature is currently enabled | [optional] 
**Ipv4AddressesCovered** | Pointer to **int64** | Number of individual IPv4 addresses covered by IPv4 and CIDR_IPV4 entries | [optional] 
**Ipv6AddressesCovered** | Pointer to **int64** | Number of individual IPv6 addresses covered by IPv6 and CIDR_IPV6 entries | [optional] 
**MaxEntries** | Pointer to **int32** | Maximum number of entries allowed | [optional] 
**Mode** | Pointer to **string** | Current mode: DISABLED, MONITOR, or ENFORCE | [optional] 
**TotalEntries** | Pointer to **int32** | Total number of entries in the allow list | [optional] 

## Methods

### NewIpAllowListSettingsXO

`func NewIpAllowListSettingsXO() *IpAllowListSettingsXO`

NewIpAllowListSettingsXO instantiates a new IpAllowListSettingsXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAllowListSettingsXOWithDefaults

`func NewIpAllowListSettingsXOWithDefaults() *IpAllowListSettingsXO`

NewIpAllowListSettingsXOWithDefaults instantiates a new IpAllowListSettingsXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveMode

`func (o *IpAllowListSettingsXO) GetEffectiveMode() string`

GetEffectiveMode returns the EffectiveMode field if non-nil, zero value otherwise.

### GetEffectiveModeOk

`func (o *IpAllowListSettingsXO) GetEffectiveModeOk() (*string, bool)`

GetEffectiveModeOk returns a tuple with the EffectiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveMode

`func (o *IpAllowListSettingsXO) SetEffectiveMode(v string)`

SetEffectiveMode sets EffectiveMode field to given value.

### HasEffectiveMode

`func (o *IpAllowListSettingsXO) HasEffectiveMode() bool`

HasEffectiveMode returns a boolean if a field has been set.

### GetFeatureEnabled

`func (o *IpAllowListSettingsXO) GetFeatureEnabled() bool`

GetFeatureEnabled returns the FeatureEnabled field if non-nil, zero value otherwise.

### GetFeatureEnabledOk

`func (o *IpAllowListSettingsXO) GetFeatureEnabledOk() (*bool, bool)`

GetFeatureEnabledOk returns a tuple with the FeatureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureEnabled

`func (o *IpAllowListSettingsXO) SetFeatureEnabled(v bool)`

SetFeatureEnabled sets FeatureEnabled field to given value.

### HasFeatureEnabled

`func (o *IpAllowListSettingsXO) HasFeatureEnabled() bool`

HasFeatureEnabled returns a boolean if a field has been set.

### GetIpv4AddressesCovered

`func (o *IpAllowListSettingsXO) GetIpv4AddressesCovered() int64`

GetIpv4AddressesCovered returns the Ipv4AddressesCovered field if non-nil, zero value otherwise.

### GetIpv4AddressesCoveredOk

`func (o *IpAllowListSettingsXO) GetIpv4AddressesCoveredOk() (*int64, bool)`

GetIpv4AddressesCoveredOk returns a tuple with the Ipv4AddressesCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddressesCovered

`func (o *IpAllowListSettingsXO) SetIpv4AddressesCovered(v int64)`

SetIpv4AddressesCovered sets Ipv4AddressesCovered field to given value.

### HasIpv4AddressesCovered

`func (o *IpAllowListSettingsXO) HasIpv4AddressesCovered() bool`

HasIpv4AddressesCovered returns a boolean if a field has been set.

### GetIpv6AddressesCovered

`func (o *IpAllowListSettingsXO) GetIpv6AddressesCovered() int64`

GetIpv6AddressesCovered returns the Ipv6AddressesCovered field if non-nil, zero value otherwise.

### GetIpv6AddressesCoveredOk

`func (o *IpAllowListSettingsXO) GetIpv6AddressesCoveredOk() (*int64, bool)`

GetIpv6AddressesCoveredOk returns a tuple with the Ipv6AddressesCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6AddressesCovered

`func (o *IpAllowListSettingsXO) SetIpv6AddressesCovered(v int64)`

SetIpv6AddressesCovered sets Ipv6AddressesCovered field to given value.

### HasIpv6AddressesCovered

`func (o *IpAllowListSettingsXO) HasIpv6AddressesCovered() bool`

HasIpv6AddressesCovered returns a boolean if a field has been set.

### GetMaxEntries

`func (o *IpAllowListSettingsXO) GetMaxEntries() int32`

GetMaxEntries returns the MaxEntries field if non-nil, zero value otherwise.

### GetMaxEntriesOk

`func (o *IpAllowListSettingsXO) GetMaxEntriesOk() (*int32, bool)`

GetMaxEntriesOk returns a tuple with the MaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntries

`func (o *IpAllowListSettingsXO) SetMaxEntries(v int32)`

SetMaxEntries sets MaxEntries field to given value.

### HasMaxEntries

`func (o *IpAllowListSettingsXO) HasMaxEntries() bool`

HasMaxEntries returns a boolean if a field has been set.

### GetMode

`func (o *IpAllowListSettingsXO) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IpAllowListSettingsXO) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IpAllowListSettingsXO) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *IpAllowListSettingsXO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTotalEntries

`func (o *IpAllowListSettingsXO) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *IpAllowListSettingsXO) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *IpAllowListSettingsXO) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *IpAllowListSettingsXO) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


