# AuthSettingsXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | auth enabled | 
**NtlmDomain** | **string** | Windows NTLM Domain | 
**NtlmHost** | **string** | Windows NTLM Hostname | 
**Password** | **string** | user password | 
**Username** | **string** | user name | 

## Methods

### NewAuthSettingsXo

`func NewAuthSettingsXo(enabled bool, ntlmDomain string, ntlmHost string, password string, username string, ) *AuthSettingsXo`

NewAuthSettingsXo instantiates a new AuthSettingsXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSettingsXoWithDefaults

`func NewAuthSettingsXoWithDefaults() *AuthSettingsXo`

NewAuthSettingsXoWithDefaults instantiates a new AuthSettingsXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AuthSettingsXo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthSettingsXo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthSettingsXo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetNtlmDomain

`func (o *AuthSettingsXo) GetNtlmDomain() string`

GetNtlmDomain returns the NtlmDomain field if non-nil, zero value otherwise.

### GetNtlmDomainOk

`func (o *AuthSettingsXo) GetNtlmDomainOk() (*string, bool)`

GetNtlmDomainOk returns a tuple with the NtlmDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtlmDomain

`func (o *AuthSettingsXo) SetNtlmDomain(v string)`

SetNtlmDomain sets NtlmDomain field to given value.


### GetNtlmHost

`func (o *AuthSettingsXo) GetNtlmHost() string`

GetNtlmHost returns the NtlmHost field if non-nil, zero value otherwise.

### GetNtlmHostOk

`func (o *AuthSettingsXo) GetNtlmHostOk() (*string, bool)`

GetNtlmHostOk returns a tuple with the NtlmHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtlmHost

`func (o *AuthSettingsXo) SetNtlmHost(v string)`

SetNtlmHost sets NtlmHost field to given value.


### GetPassword

`func (o *AuthSettingsXo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthSettingsXo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthSettingsXo) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *AuthSettingsXo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthSettingsXo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthSettingsXo) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


