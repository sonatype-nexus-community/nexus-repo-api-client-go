# LdapTestUserMappingXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | User&#39;s email address from LDAP | [optional] 
**Membership** | Pointer to **[]string** | Group memberships from LDAP | [optional] 
**RealName** | Pointer to **string** | User&#39;s real name from LDAP | [optional] 
**Username** | Pointer to **string** | Username from LDAP | [optional] 

## Methods

### NewLdapTestUserMappingXo

`func NewLdapTestUserMappingXo() *LdapTestUserMappingXo`

NewLdapTestUserMappingXo instantiates a new LdapTestUserMappingXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapTestUserMappingXoWithDefaults

`func NewLdapTestUserMappingXoWithDefaults() *LdapTestUserMappingXo`

NewLdapTestUserMappingXoWithDefaults instantiates a new LdapTestUserMappingXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LdapTestUserMappingXo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LdapTestUserMappingXo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LdapTestUserMappingXo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LdapTestUserMappingXo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMembership

`func (o *LdapTestUserMappingXo) GetMembership() []string`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *LdapTestUserMappingXo) GetMembershipOk() (*[]string, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *LdapTestUserMappingXo) SetMembership(v []string)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *LdapTestUserMappingXo) HasMembership() bool`

HasMembership returns a boolean if a field has been set.

### GetRealName

`func (o *LdapTestUserMappingXo) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *LdapTestUserMappingXo) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *LdapTestUserMappingXo) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *LdapTestUserMappingXo) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetUsername

`func (o *LdapTestUserMappingXo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LdapTestUserMappingXo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LdapTestUserMappingXo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LdapTestUserMappingXo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


