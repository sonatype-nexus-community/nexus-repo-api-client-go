# LdapVerifyLoginXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base64Password** | **string** | Base64-encoded password to test | 
**Base64Username** | **string** | Base64-encoded username to test | 
**LdapServer** | [**CreateLdapServerXo**](CreateLdapServerXo.md) |  | 

## Methods

### NewLdapVerifyLoginXo

`func NewLdapVerifyLoginXo(base64Password string, base64Username string, ldapServer CreateLdapServerXo, ) *LdapVerifyLoginXo`

NewLdapVerifyLoginXo instantiates a new LdapVerifyLoginXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapVerifyLoginXoWithDefaults

`func NewLdapVerifyLoginXoWithDefaults() *LdapVerifyLoginXo`

NewLdapVerifyLoginXoWithDefaults instantiates a new LdapVerifyLoginXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase64Password

`func (o *LdapVerifyLoginXo) GetBase64Password() string`

GetBase64Password returns the Base64Password field if non-nil, zero value otherwise.

### GetBase64PasswordOk

`func (o *LdapVerifyLoginXo) GetBase64PasswordOk() (*string, bool)`

GetBase64PasswordOk returns a tuple with the Base64Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Password

`func (o *LdapVerifyLoginXo) SetBase64Password(v string)`

SetBase64Password sets Base64Password field to given value.


### GetBase64Username

`func (o *LdapVerifyLoginXo) GetBase64Username() string`

GetBase64Username returns the Base64Username field if non-nil, zero value otherwise.

### GetBase64UsernameOk

`func (o *LdapVerifyLoginXo) GetBase64UsernameOk() (*string, bool)`

GetBase64UsernameOk returns a tuple with the Base64Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Username

`func (o *LdapVerifyLoginXo) SetBase64Username(v string)`

SetBase64Username sets Base64Username field to given value.


### GetLdapServer

`func (o *LdapVerifyLoginXo) GetLdapServer() CreateLdapServerXo`

GetLdapServer returns the LdapServer field if non-nil, zero value otherwise.

### GetLdapServerOk

`func (o *LdapVerifyLoginXo) GetLdapServerOk() (*CreateLdapServerXo, bool)`

GetLdapServerOk returns a tuple with the LdapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServer

`func (o *LdapVerifyLoginXo) SetLdapServer(v CreateLdapServerXo)`

SetLdapServer sets LdapServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


