# CreateLdapServerXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPassword** | **string** | The password to bind with. Required if authScheme other than none. | 
**AuthRealm** | Pointer to **string** | The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5 | [optional] 
**AuthScheme** | **string** | Authentication scheme used for connecting to LDAP server | 
**AuthUsername** | Pointer to **string** | This must be a fully qualified username if simple authentication is used. Required if authScheme other than none. | [optional] 
**ConnectionRetryDelaySeconds** | **int32** | How long to wait before retrying | 
**ConnectionTimeoutSeconds** | **int32** | How long to wait before timeout | 
**GroupBaseDn** | Pointer to **string** | The relative DN where group objects are found (e.g. ou&#x3D;Group). This value will have the Search base DN value appended to form the full Group search base DN. | [optional] 
**GroupIdAttribute** | Pointer to **string** | This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static | [optional] 
**GroupMemberAttribute** | Pointer to **string** | LDAP attribute containing the usernames for the group. Required if groupType is static | [optional] 
**GroupMemberFormat** | Pointer to **string** | The format of user ID stored in the group member attribute. Required if groupType is static | [optional] 
**GroupObjectClass** | Pointer to **string** | LDAP class for group objects. Required if groupType is static | [optional] 
**GroupSubtree** | Pointer to **bool** | Are groups located in structures below the group base DN | [optional] 
**GroupType** | Pointer to **string** | Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true. | [optional] 
**Host** | **string** | LDAP server connection hostname | 
**LdapGroupsAsRoles** | Pointer to **bool** | Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles | [optional] 
**MaxIncidentsCount** | **int32** | How many retry attempts | 
**Name** | **string** | LDAP server name | 
**Port** | **int32** | LDAP server connection port to use | 
**Protocol** | **string** | LDAP server connection Protocol to use | 
**SearchBase** | **string** | LDAP location to be added to the connection URL | 
**UseTrustStore** | Pointer to **bool** | Whether to use certificates stored in Nexus Repository Manager&#39;s truststore | [optional] 
**UserBaseDn** | Pointer to **string** | The relative DN where user objects are found (e.g. ou&#x3D;people). This value will have the Search base DN value appended to form the full User search base DN. | [optional] 
**UserEmailAddressAttribute** | Pointer to **string** | This is used to find an email address given the user ID | [optional] 
**UserIdAttribute** | Pointer to **string** | This is used to find a user given its user ID | [optional] 
**UserLdapFilter** | Pointer to **string** | LDAP search filter to limit user search | [optional] 
**UserMemberOfAttribute** | Pointer to **string** | Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic | [optional] 
**UserObjectClass** | Pointer to **string** | LDAP class for user objects | [optional] 
**UserPasswordAttribute** | Pointer to **string** | If this field is blank the user will be authenticated against a bind with the LDAP server | [optional] 
**UserRealNameAttribute** | Pointer to **string** | This is used to find a real name given the user ID | [optional] 
**UserSubtree** | Pointer to **bool** | Are users located in structures below the user base DN? | [optional] 

## Methods

### NewCreateLdapServerXo

`func NewCreateLdapServerXo(authPassword string, authScheme string, connectionRetryDelaySeconds int32, connectionTimeoutSeconds int32, host string, maxIncidentsCount int32, name string, port int32, protocol string, searchBase string, ) *CreateLdapServerXo`

NewCreateLdapServerXo instantiates a new CreateLdapServerXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLdapServerXoWithDefaults

`func NewCreateLdapServerXoWithDefaults() *CreateLdapServerXo`

NewCreateLdapServerXoWithDefaults instantiates a new CreateLdapServerXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthPassword

`func (o *CreateLdapServerXo) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *CreateLdapServerXo) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *CreateLdapServerXo) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.


### GetAuthRealm

`func (o *CreateLdapServerXo) GetAuthRealm() string`

GetAuthRealm returns the AuthRealm field if non-nil, zero value otherwise.

### GetAuthRealmOk

`func (o *CreateLdapServerXo) GetAuthRealmOk() (*string, bool)`

GetAuthRealmOk returns a tuple with the AuthRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRealm

`func (o *CreateLdapServerXo) SetAuthRealm(v string)`

SetAuthRealm sets AuthRealm field to given value.

### HasAuthRealm

`func (o *CreateLdapServerXo) HasAuthRealm() bool`

HasAuthRealm returns a boolean if a field has been set.

### GetAuthScheme

`func (o *CreateLdapServerXo) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *CreateLdapServerXo) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *CreateLdapServerXo) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.


### GetAuthUsername

`func (o *CreateLdapServerXo) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *CreateLdapServerXo) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *CreateLdapServerXo) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *CreateLdapServerXo) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### GetConnectionRetryDelaySeconds

`func (o *CreateLdapServerXo) GetConnectionRetryDelaySeconds() int32`

GetConnectionRetryDelaySeconds returns the ConnectionRetryDelaySeconds field if non-nil, zero value otherwise.

### GetConnectionRetryDelaySecondsOk

`func (o *CreateLdapServerXo) GetConnectionRetryDelaySecondsOk() (*int32, bool)`

GetConnectionRetryDelaySecondsOk returns a tuple with the ConnectionRetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionRetryDelaySeconds

`func (o *CreateLdapServerXo) SetConnectionRetryDelaySeconds(v int32)`

SetConnectionRetryDelaySeconds sets ConnectionRetryDelaySeconds field to given value.


### GetConnectionTimeoutSeconds

`func (o *CreateLdapServerXo) GetConnectionTimeoutSeconds() int32`

GetConnectionTimeoutSeconds returns the ConnectionTimeoutSeconds field if non-nil, zero value otherwise.

### GetConnectionTimeoutSecondsOk

`func (o *CreateLdapServerXo) GetConnectionTimeoutSecondsOk() (*int32, bool)`

GetConnectionTimeoutSecondsOk returns a tuple with the ConnectionTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutSeconds

`func (o *CreateLdapServerXo) SetConnectionTimeoutSeconds(v int32)`

SetConnectionTimeoutSeconds sets ConnectionTimeoutSeconds field to given value.


### GetGroupBaseDn

`func (o *CreateLdapServerXo) GetGroupBaseDn() string`

GetGroupBaseDn returns the GroupBaseDn field if non-nil, zero value otherwise.

### GetGroupBaseDnOk

`func (o *CreateLdapServerXo) GetGroupBaseDnOk() (*string, bool)`

GetGroupBaseDnOk returns a tuple with the GroupBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBaseDn

`func (o *CreateLdapServerXo) SetGroupBaseDn(v string)`

SetGroupBaseDn sets GroupBaseDn field to given value.

### HasGroupBaseDn

`func (o *CreateLdapServerXo) HasGroupBaseDn() bool`

HasGroupBaseDn returns a boolean if a field has been set.

### GetGroupIdAttribute

`func (o *CreateLdapServerXo) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *CreateLdapServerXo) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *CreateLdapServerXo) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.

### HasGroupIdAttribute

`func (o *CreateLdapServerXo) HasGroupIdAttribute() bool`

HasGroupIdAttribute returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *CreateLdapServerXo) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *CreateLdapServerXo) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *CreateLdapServerXo) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *CreateLdapServerXo) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberFormat

`func (o *CreateLdapServerXo) GetGroupMemberFormat() string`

GetGroupMemberFormat returns the GroupMemberFormat field if non-nil, zero value otherwise.

### GetGroupMemberFormatOk

`func (o *CreateLdapServerXo) GetGroupMemberFormatOk() (*string, bool)`

GetGroupMemberFormatOk returns a tuple with the GroupMemberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberFormat

`func (o *CreateLdapServerXo) SetGroupMemberFormat(v string)`

SetGroupMemberFormat sets GroupMemberFormat field to given value.

### HasGroupMemberFormat

`func (o *CreateLdapServerXo) HasGroupMemberFormat() bool`

HasGroupMemberFormat returns a boolean if a field has been set.

### GetGroupObjectClass

`func (o *CreateLdapServerXo) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *CreateLdapServerXo) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *CreateLdapServerXo) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.

### HasGroupObjectClass

`func (o *CreateLdapServerXo) HasGroupObjectClass() bool`

HasGroupObjectClass returns a boolean if a field has been set.

### GetGroupSubtree

`func (o *CreateLdapServerXo) GetGroupSubtree() bool`

GetGroupSubtree returns the GroupSubtree field if non-nil, zero value otherwise.

### GetGroupSubtreeOk

`func (o *CreateLdapServerXo) GetGroupSubtreeOk() (*bool, bool)`

GetGroupSubtreeOk returns a tuple with the GroupSubtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSubtree

`func (o *CreateLdapServerXo) SetGroupSubtree(v bool)`

SetGroupSubtree sets GroupSubtree field to given value.

### HasGroupSubtree

`func (o *CreateLdapServerXo) HasGroupSubtree() bool`

HasGroupSubtree returns a boolean if a field has been set.

### GetGroupType

`func (o *CreateLdapServerXo) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *CreateLdapServerXo) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *CreateLdapServerXo) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *CreateLdapServerXo) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetHost

`func (o *CreateLdapServerXo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateLdapServerXo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateLdapServerXo) SetHost(v string)`

SetHost sets Host field to given value.


### GetLdapGroupsAsRoles

`func (o *CreateLdapServerXo) GetLdapGroupsAsRoles() bool`

GetLdapGroupsAsRoles returns the LdapGroupsAsRoles field if non-nil, zero value otherwise.

### GetLdapGroupsAsRolesOk

`func (o *CreateLdapServerXo) GetLdapGroupsAsRolesOk() (*bool, bool)`

GetLdapGroupsAsRolesOk returns a tuple with the LdapGroupsAsRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupsAsRoles

`func (o *CreateLdapServerXo) SetLdapGroupsAsRoles(v bool)`

SetLdapGroupsAsRoles sets LdapGroupsAsRoles field to given value.

### HasLdapGroupsAsRoles

`func (o *CreateLdapServerXo) HasLdapGroupsAsRoles() bool`

HasLdapGroupsAsRoles returns a boolean if a field has been set.

### GetMaxIncidentsCount

`func (o *CreateLdapServerXo) GetMaxIncidentsCount() int32`

GetMaxIncidentsCount returns the MaxIncidentsCount field if non-nil, zero value otherwise.

### GetMaxIncidentsCountOk

`func (o *CreateLdapServerXo) GetMaxIncidentsCountOk() (*int32, bool)`

GetMaxIncidentsCountOk returns a tuple with the MaxIncidentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIncidentsCount

`func (o *CreateLdapServerXo) SetMaxIncidentsCount(v int32)`

SetMaxIncidentsCount sets MaxIncidentsCount field to given value.


### GetName

`func (o *CreateLdapServerXo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLdapServerXo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLdapServerXo) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *CreateLdapServerXo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateLdapServerXo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateLdapServerXo) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *CreateLdapServerXo) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateLdapServerXo) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateLdapServerXo) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSearchBase

`func (o *CreateLdapServerXo) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *CreateLdapServerXo) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *CreateLdapServerXo) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.


### GetUseTrustStore

`func (o *CreateLdapServerXo) GetUseTrustStore() bool`

GetUseTrustStore returns the UseTrustStore field if non-nil, zero value otherwise.

### GetUseTrustStoreOk

`func (o *CreateLdapServerXo) GetUseTrustStoreOk() (*bool, bool)`

GetUseTrustStoreOk returns a tuple with the UseTrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrustStore

`func (o *CreateLdapServerXo) SetUseTrustStore(v bool)`

SetUseTrustStore sets UseTrustStore field to given value.

### HasUseTrustStore

`func (o *CreateLdapServerXo) HasUseTrustStore() bool`

HasUseTrustStore returns a boolean if a field has been set.

### GetUserBaseDn

`func (o *CreateLdapServerXo) GetUserBaseDn() string`

GetUserBaseDn returns the UserBaseDn field if non-nil, zero value otherwise.

### GetUserBaseDnOk

`func (o *CreateLdapServerXo) GetUserBaseDnOk() (*string, bool)`

GetUserBaseDnOk returns a tuple with the UserBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDn

`func (o *CreateLdapServerXo) SetUserBaseDn(v string)`

SetUserBaseDn sets UserBaseDn field to given value.

### HasUserBaseDn

`func (o *CreateLdapServerXo) HasUserBaseDn() bool`

HasUserBaseDn returns a boolean if a field has been set.

### GetUserEmailAddressAttribute

`func (o *CreateLdapServerXo) GetUserEmailAddressAttribute() string`

GetUserEmailAddressAttribute returns the UserEmailAddressAttribute field if non-nil, zero value otherwise.

### GetUserEmailAddressAttributeOk

`func (o *CreateLdapServerXo) GetUserEmailAddressAttributeOk() (*string, bool)`

GetUserEmailAddressAttributeOk returns a tuple with the UserEmailAddressAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmailAddressAttribute

`func (o *CreateLdapServerXo) SetUserEmailAddressAttribute(v string)`

SetUserEmailAddressAttribute sets UserEmailAddressAttribute field to given value.

### HasUserEmailAddressAttribute

`func (o *CreateLdapServerXo) HasUserEmailAddressAttribute() bool`

HasUserEmailAddressAttribute returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *CreateLdapServerXo) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *CreateLdapServerXo) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *CreateLdapServerXo) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *CreateLdapServerXo) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserLdapFilter

`func (o *CreateLdapServerXo) GetUserLdapFilter() string`

GetUserLdapFilter returns the UserLdapFilter field if non-nil, zero value otherwise.

### GetUserLdapFilterOk

`func (o *CreateLdapServerXo) GetUserLdapFilterOk() (*string, bool)`

GetUserLdapFilterOk returns a tuple with the UserLdapFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLdapFilter

`func (o *CreateLdapServerXo) SetUserLdapFilter(v string)`

SetUserLdapFilter sets UserLdapFilter field to given value.

### HasUserLdapFilter

`func (o *CreateLdapServerXo) HasUserLdapFilter() bool`

HasUserLdapFilter returns a boolean if a field has been set.

### GetUserMemberOfAttribute

`func (o *CreateLdapServerXo) GetUserMemberOfAttribute() string`

GetUserMemberOfAttribute returns the UserMemberOfAttribute field if non-nil, zero value otherwise.

### GetUserMemberOfAttributeOk

`func (o *CreateLdapServerXo) GetUserMemberOfAttributeOk() (*string, bool)`

GetUserMemberOfAttributeOk returns a tuple with the UserMemberOfAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMemberOfAttribute

`func (o *CreateLdapServerXo) SetUserMemberOfAttribute(v string)`

SetUserMemberOfAttribute sets UserMemberOfAttribute field to given value.

### HasUserMemberOfAttribute

`func (o *CreateLdapServerXo) HasUserMemberOfAttribute() bool`

HasUserMemberOfAttribute returns a boolean if a field has been set.

### GetUserObjectClass

`func (o *CreateLdapServerXo) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *CreateLdapServerXo) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *CreateLdapServerXo) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.

### HasUserObjectClass

`func (o *CreateLdapServerXo) HasUserObjectClass() bool`

HasUserObjectClass returns a boolean if a field has been set.

### GetUserPasswordAttribute

`func (o *CreateLdapServerXo) GetUserPasswordAttribute() string`

GetUserPasswordAttribute returns the UserPasswordAttribute field if non-nil, zero value otherwise.

### GetUserPasswordAttributeOk

`func (o *CreateLdapServerXo) GetUserPasswordAttributeOk() (*string, bool)`

GetUserPasswordAttributeOk returns a tuple with the UserPasswordAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordAttribute

`func (o *CreateLdapServerXo) SetUserPasswordAttribute(v string)`

SetUserPasswordAttribute sets UserPasswordAttribute field to given value.

### HasUserPasswordAttribute

`func (o *CreateLdapServerXo) HasUserPasswordAttribute() bool`

HasUserPasswordAttribute returns a boolean if a field has been set.

### GetUserRealNameAttribute

`func (o *CreateLdapServerXo) GetUserRealNameAttribute() string`

GetUserRealNameAttribute returns the UserRealNameAttribute field if non-nil, zero value otherwise.

### GetUserRealNameAttributeOk

`func (o *CreateLdapServerXo) GetUserRealNameAttributeOk() (*string, bool)`

GetUserRealNameAttributeOk returns a tuple with the UserRealNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRealNameAttribute

`func (o *CreateLdapServerXo) SetUserRealNameAttribute(v string)`

SetUserRealNameAttribute sets UserRealNameAttribute field to given value.

### HasUserRealNameAttribute

`func (o *CreateLdapServerXo) HasUserRealNameAttribute() bool`

HasUserRealNameAttribute returns a boolean if a field has been set.

### GetUserSubtree

`func (o *CreateLdapServerXo) GetUserSubtree() bool`

GetUserSubtree returns the UserSubtree field if non-nil, zero value otherwise.

### GetUserSubtreeOk

`func (o *CreateLdapServerXo) GetUserSubtreeOk() (*bool, bool)`

GetUserSubtreeOk returns a tuple with the UserSubtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSubtree

`func (o *CreateLdapServerXo) SetUserSubtree(v bool)`

SetUserSubtree sets UserSubtree field to given value.

### HasUserSubtree

`func (o *CreateLdapServerXo) HasUserSubtree() bool`

HasUserSubtree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


