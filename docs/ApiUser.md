# ApiUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | The email address associated with the user. | [optional] 
**ExternalRoles** | Pointer to **[]string** | The roles which the user has been assigned in an external source, e.g. LDAP group. These cannot be changed within the Nexus Repository Manager. | [optional] 
**FirstName** | Pointer to **string** | The first name of the user. | [optional] 
**LastName** | Pointer to **string** | The last name of the user. | [optional] 
**ReadOnly** | Pointer to **bool** | Indicates whether the user&#39;s properties could be modified by the Nexus Repository Manager. When false only roles are considered during update. | [optional] 
**Roles** | Pointer to **[]string** | The roles which the user has been assigned within Nexus. | [optional] 
**Source** | Pointer to **string** | The source of the user. When creating user, if the source is \&quot;default\&quot;, the local Nexus Security User will be created. If the source is anything but \&quot;default\&quot; (i.e., LDAP or SAML), a local LdapUser (or SamlUser) will be created and have the listed roles assigned to that user. | [optional] 
**Status** | **string** | The user&#39;s status, e.g. active or disabled. | 
**UserId** | Pointer to **string** | The userid which is required for login. This value cannot be changed. | [optional] 

## Methods

### NewApiUser

`func NewApiUser(status string, ) *ApiUser`

NewApiUser instantiates a new ApiUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserWithDefaults

`func NewApiUserWithDefaults() *ApiUser`

NewApiUserWithDefaults instantiates a new ApiUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *ApiUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ApiUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ApiUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ApiUser) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetExternalRoles

`func (o *ApiUser) GetExternalRoles() []string`

GetExternalRoles returns the ExternalRoles field if non-nil, zero value otherwise.

### GetExternalRolesOk

`func (o *ApiUser) GetExternalRolesOk() (*[]string, bool)`

GetExternalRolesOk returns a tuple with the ExternalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRoles

`func (o *ApiUser) SetExternalRoles(v []string)`

SetExternalRoles sets ExternalRoles field to given value.

### HasExternalRoles

`func (o *ApiUser) HasExternalRoles() bool`

HasExternalRoles returns a boolean if a field has been set.

### GetFirstName

`func (o *ApiUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ApiUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ApiUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ApiUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ApiUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ApiUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ApiUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ApiUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetReadOnly

`func (o *ApiUser) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApiUser) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApiUser) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApiUser) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRoles

`func (o *ApiUser) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiUser) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiUser) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSource

`func (o *ApiUser) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiUser) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiUser) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiUser) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *ApiUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUserId

`func (o *ApiUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


