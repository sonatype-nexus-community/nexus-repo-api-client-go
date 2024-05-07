# ApiEmailConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**FromAddress** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**NexusTrustStoreEnabled** | Pointer to **bool** | Use the Nexus Repository Manager&#39;s certificate truststore | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Port** | **int32** |  | 
**SslOnConnectEnabled** | Pointer to **bool** | Enable SSL/TLS Encryption upon Connection | [optional] 
**SslServerIdentityCheckEnabled** | Pointer to **bool** | Verify the server certificate when using TLS or SSL | [optional] 
**StartTlsEnabled** | Pointer to **bool** | Enable STARTTLS Support for Insecure Connections | [optional] 
**StartTlsRequired** | Pointer to **bool** | Require STARTTLS Support | [optional] 
**SubjectPrefix** | Pointer to **string** | A prefix to add to all email subjects to aid in identifying automated emails | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewApiEmailConfiguration

`func NewApiEmailConfiguration(port int32, ) *ApiEmailConfiguration`

NewApiEmailConfiguration instantiates a new ApiEmailConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEmailConfigurationWithDefaults

`func NewApiEmailConfigurationWithDefaults() *ApiEmailConfiguration`

NewApiEmailConfigurationWithDefaults instantiates a new ApiEmailConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiEmailConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiEmailConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiEmailConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiEmailConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFromAddress

`func (o *ApiEmailConfiguration) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ApiEmailConfiguration) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ApiEmailConfiguration) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *ApiEmailConfiguration) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetHost

`func (o *ApiEmailConfiguration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ApiEmailConfiguration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ApiEmailConfiguration) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ApiEmailConfiguration) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNexusTrustStoreEnabled

`func (o *ApiEmailConfiguration) GetNexusTrustStoreEnabled() bool`

GetNexusTrustStoreEnabled returns the NexusTrustStoreEnabled field if non-nil, zero value otherwise.

### GetNexusTrustStoreEnabledOk

`func (o *ApiEmailConfiguration) GetNexusTrustStoreEnabledOk() (*bool, bool)`

GetNexusTrustStoreEnabledOk returns a tuple with the NexusTrustStoreEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusTrustStoreEnabled

`func (o *ApiEmailConfiguration) SetNexusTrustStoreEnabled(v bool)`

SetNexusTrustStoreEnabled sets NexusTrustStoreEnabled field to given value.

### HasNexusTrustStoreEnabled

`func (o *ApiEmailConfiguration) HasNexusTrustStoreEnabled() bool`

HasNexusTrustStoreEnabled returns a boolean if a field has been set.

### GetPassword

`func (o *ApiEmailConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiEmailConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiEmailConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiEmailConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *ApiEmailConfiguration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiEmailConfiguration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiEmailConfiguration) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSslOnConnectEnabled

`func (o *ApiEmailConfiguration) GetSslOnConnectEnabled() bool`

GetSslOnConnectEnabled returns the SslOnConnectEnabled field if non-nil, zero value otherwise.

### GetSslOnConnectEnabledOk

`func (o *ApiEmailConfiguration) GetSslOnConnectEnabledOk() (*bool, bool)`

GetSslOnConnectEnabledOk returns a tuple with the SslOnConnectEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslOnConnectEnabled

`func (o *ApiEmailConfiguration) SetSslOnConnectEnabled(v bool)`

SetSslOnConnectEnabled sets SslOnConnectEnabled field to given value.

### HasSslOnConnectEnabled

`func (o *ApiEmailConfiguration) HasSslOnConnectEnabled() bool`

HasSslOnConnectEnabled returns a boolean if a field has been set.

### GetSslServerIdentityCheckEnabled

`func (o *ApiEmailConfiguration) GetSslServerIdentityCheckEnabled() bool`

GetSslServerIdentityCheckEnabled returns the SslServerIdentityCheckEnabled field if non-nil, zero value otherwise.

### GetSslServerIdentityCheckEnabledOk

`func (o *ApiEmailConfiguration) GetSslServerIdentityCheckEnabledOk() (*bool, bool)`

GetSslServerIdentityCheckEnabledOk returns a tuple with the SslServerIdentityCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslServerIdentityCheckEnabled

`func (o *ApiEmailConfiguration) SetSslServerIdentityCheckEnabled(v bool)`

SetSslServerIdentityCheckEnabled sets SslServerIdentityCheckEnabled field to given value.

### HasSslServerIdentityCheckEnabled

`func (o *ApiEmailConfiguration) HasSslServerIdentityCheckEnabled() bool`

HasSslServerIdentityCheckEnabled returns a boolean if a field has been set.

### GetStartTlsEnabled

`func (o *ApiEmailConfiguration) GetStartTlsEnabled() bool`

GetStartTlsEnabled returns the StartTlsEnabled field if non-nil, zero value otherwise.

### GetStartTlsEnabledOk

`func (o *ApiEmailConfiguration) GetStartTlsEnabledOk() (*bool, bool)`

GetStartTlsEnabledOk returns a tuple with the StartTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTlsEnabled

`func (o *ApiEmailConfiguration) SetStartTlsEnabled(v bool)`

SetStartTlsEnabled sets StartTlsEnabled field to given value.

### HasStartTlsEnabled

`func (o *ApiEmailConfiguration) HasStartTlsEnabled() bool`

HasStartTlsEnabled returns a boolean if a field has been set.

### GetStartTlsRequired

`func (o *ApiEmailConfiguration) GetStartTlsRequired() bool`

GetStartTlsRequired returns the StartTlsRequired field if non-nil, zero value otherwise.

### GetStartTlsRequiredOk

`func (o *ApiEmailConfiguration) GetStartTlsRequiredOk() (*bool, bool)`

GetStartTlsRequiredOk returns a tuple with the StartTlsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTlsRequired

`func (o *ApiEmailConfiguration) SetStartTlsRequired(v bool)`

SetStartTlsRequired sets StartTlsRequired field to given value.

### HasStartTlsRequired

`func (o *ApiEmailConfiguration) HasStartTlsRequired() bool`

HasStartTlsRequired returns a boolean if a field has been set.

### GetSubjectPrefix

`func (o *ApiEmailConfiguration) GetSubjectPrefix() string`

GetSubjectPrefix returns the SubjectPrefix field if non-nil, zero value otherwise.

### GetSubjectPrefixOk

`func (o *ApiEmailConfiguration) GetSubjectPrefixOk() (*string, bool)`

GetSubjectPrefixOk returns a tuple with the SubjectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPrefix

`func (o *ApiEmailConfiguration) SetSubjectPrefix(v string)`

SetSubjectPrefix sets SubjectPrefix field to given value.

### HasSubjectPrefix

`func (o *ApiEmailConfiguration) HasSubjectPrefix() bool`

HasSubjectPrefix returns a boolean if a field has been set.

### GetUsername

`func (o *ApiEmailConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiEmailConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiEmailConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiEmailConfiguration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


