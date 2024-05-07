# ReadLdapServerXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**GroupType** | **string** | Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true. | 
**Host** | **string** | LDAP server connection hostname | 
**Id** | Pointer to **string** | LDAP server ID | [optional] 
**LdapGroupsAsRoles** | Pointer to **bool** | Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles | [optional] 
**MaxIncidentsCount** | **int32** | How many retry attempts | 
**Name** | **string** | LDAP server name | 
**Order** | Pointer to **int32** | Order number in which the server is being used when looking for a user | [optional] 
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

### NewReadLdapServerXo

`func NewReadLdapServerXo(authScheme string, connectionRetryDelaySeconds int32, connectionTimeoutSeconds int32, groupType string, host string, maxIncidentsCount int32, name string, port int32, protocol string, searchBase string, ) *ReadLdapServerXo`

NewReadLdapServerXo instantiates a new ReadLdapServerXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadLdapServerXoWithDefaults

`func NewReadLdapServerXoWithDefaults() *ReadLdapServerXo`

NewReadLdapServerXoWithDefaults instantiates a new ReadLdapServerXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthRealm

`func (o *ReadLdapServerXo) GetAuthRealm() string`

GetAuthRealm returns the AuthRealm field if non-nil, zero value otherwise.

### GetAuthRealmOk

`func (o *ReadLdapServerXo) GetAuthRealmOk() (*string, bool)`

GetAuthRealmOk returns a tuple with the AuthRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRealm

`func (o *ReadLdapServerXo) SetAuthRealm(v string)`

SetAuthRealm sets AuthRealm field to given value.

### HasAuthRealm

`func (o *ReadLdapServerXo) HasAuthRealm() bool`

HasAuthRealm returns a boolean if a field has been set.

### GetAuthScheme

`func (o *ReadLdapServerXo) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ReadLdapServerXo) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ReadLdapServerXo) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.


### GetAuthUsername

`func (o *ReadLdapServerXo) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *ReadLdapServerXo) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *ReadLdapServerXo) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *ReadLdapServerXo) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### GetConnectionRetryDelaySeconds

`func (o *ReadLdapServerXo) GetConnectionRetryDelaySeconds() int32`

GetConnectionRetryDelaySeconds returns the ConnectionRetryDelaySeconds field if non-nil, zero value otherwise.

### GetConnectionRetryDelaySecondsOk

`func (o *ReadLdapServerXo) GetConnectionRetryDelaySecondsOk() (*int32, bool)`

GetConnectionRetryDelaySecondsOk returns a tuple with the ConnectionRetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionRetryDelaySeconds

`func (o *ReadLdapServerXo) SetConnectionRetryDelaySeconds(v int32)`

SetConnectionRetryDelaySeconds sets ConnectionRetryDelaySeconds field to given value.


### GetConnectionTimeoutSeconds

`func (o *ReadLdapServerXo) GetConnectionTimeoutSeconds() int32`

GetConnectionTimeoutSeconds returns the ConnectionTimeoutSeconds field if non-nil, zero value otherwise.

### GetConnectionTimeoutSecondsOk

`func (o *ReadLdapServerXo) GetConnectionTimeoutSecondsOk() (*int32, bool)`

GetConnectionTimeoutSecondsOk returns a tuple with the ConnectionTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutSeconds

`func (o *ReadLdapServerXo) SetConnectionTimeoutSeconds(v int32)`

SetConnectionTimeoutSeconds sets ConnectionTimeoutSeconds field to given value.


### GetGroupBaseDn

`func (o *ReadLdapServerXo) GetGroupBaseDn() string`

GetGroupBaseDn returns the GroupBaseDn field if non-nil, zero value otherwise.

### GetGroupBaseDnOk

`func (o *ReadLdapServerXo) GetGroupBaseDnOk() (*string, bool)`

GetGroupBaseDnOk returns a tuple with the GroupBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBaseDn

`func (o *ReadLdapServerXo) SetGroupBaseDn(v string)`

SetGroupBaseDn sets GroupBaseDn field to given value.

### HasGroupBaseDn

`func (o *ReadLdapServerXo) HasGroupBaseDn() bool`

HasGroupBaseDn returns a boolean if a field has been set.

### GetGroupIdAttribute

`func (o *ReadLdapServerXo) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *ReadLdapServerXo) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *ReadLdapServerXo) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.

### HasGroupIdAttribute

`func (o *ReadLdapServerXo) HasGroupIdAttribute() bool`

HasGroupIdAttribute returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *ReadLdapServerXo) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *ReadLdapServerXo) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *ReadLdapServerXo) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *ReadLdapServerXo) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberFormat

`func (o *ReadLdapServerXo) GetGroupMemberFormat() string`

GetGroupMemberFormat returns the GroupMemberFormat field if non-nil, zero value otherwise.

### GetGroupMemberFormatOk

`func (o *ReadLdapServerXo) GetGroupMemberFormatOk() (*string, bool)`

GetGroupMemberFormatOk returns a tuple with the GroupMemberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberFormat

`func (o *ReadLdapServerXo) SetGroupMemberFormat(v string)`

SetGroupMemberFormat sets GroupMemberFormat field to given value.

### HasGroupMemberFormat

`func (o *ReadLdapServerXo) HasGroupMemberFormat() bool`

HasGroupMemberFormat returns a boolean if a field has been set.

### GetGroupObjectClass

`func (o *ReadLdapServerXo) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *ReadLdapServerXo) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *ReadLdapServerXo) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.

### HasGroupObjectClass

`func (o *ReadLdapServerXo) HasGroupObjectClass() bool`

HasGroupObjectClass returns a boolean if a field has been set.

### GetGroupSubtree

`func (o *ReadLdapServerXo) GetGroupSubtree() bool`

GetGroupSubtree returns the GroupSubtree field if non-nil, zero value otherwise.

### GetGroupSubtreeOk

`func (o *ReadLdapServerXo) GetGroupSubtreeOk() (*bool, bool)`

GetGroupSubtreeOk returns a tuple with the GroupSubtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSubtree

`func (o *ReadLdapServerXo) SetGroupSubtree(v bool)`

SetGroupSubtree sets GroupSubtree field to given value.

### HasGroupSubtree

`func (o *ReadLdapServerXo) HasGroupSubtree() bool`

HasGroupSubtree returns a boolean if a field has been set.

### GetGroupType

`func (o *ReadLdapServerXo) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *ReadLdapServerXo) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *ReadLdapServerXo) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.


### GetHost

`func (o *ReadLdapServerXo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ReadLdapServerXo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ReadLdapServerXo) SetHost(v string)`

SetHost sets Host field to given value.


### GetId

`func (o *ReadLdapServerXo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReadLdapServerXo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReadLdapServerXo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReadLdapServerXo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLdapGroupsAsRoles

`func (o *ReadLdapServerXo) GetLdapGroupsAsRoles() bool`

GetLdapGroupsAsRoles returns the LdapGroupsAsRoles field if non-nil, zero value otherwise.

### GetLdapGroupsAsRolesOk

`func (o *ReadLdapServerXo) GetLdapGroupsAsRolesOk() (*bool, bool)`

GetLdapGroupsAsRolesOk returns a tuple with the LdapGroupsAsRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupsAsRoles

`func (o *ReadLdapServerXo) SetLdapGroupsAsRoles(v bool)`

SetLdapGroupsAsRoles sets LdapGroupsAsRoles field to given value.

### HasLdapGroupsAsRoles

`func (o *ReadLdapServerXo) HasLdapGroupsAsRoles() bool`

HasLdapGroupsAsRoles returns a boolean if a field has been set.

### GetMaxIncidentsCount

`func (o *ReadLdapServerXo) GetMaxIncidentsCount() int32`

GetMaxIncidentsCount returns the MaxIncidentsCount field if non-nil, zero value otherwise.

### GetMaxIncidentsCountOk

`func (o *ReadLdapServerXo) GetMaxIncidentsCountOk() (*int32, bool)`

GetMaxIncidentsCountOk returns a tuple with the MaxIncidentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIncidentsCount

`func (o *ReadLdapServerXo) SetMaxIncidentsCount(v int32)`

SetMaxIncidentsCount sets MaxIncidentsCount field to given value.


### GetName

`func (o *ReadLdapServerXo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReadLdapServerXo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReadLdapServerXo) SetName(v string)`

SetName sets Name field to given value.


### GetOrder

`func (o *ReadLdapServerXo) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ReadLdapServerXo) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ReadLdapServerXo) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ReadLdapServerXo) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPort

`func (o *ReadLdapServerXo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ReadLdapServerXo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ReadLdapServerXo) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *ReadLdapServerXo) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ReadLdapServerXo) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ReadLdapServerXo) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSearchBase

`func (o *ReadLdapServerXo) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *ReadLdapServerXo) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *ReadLdapServerXo) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.


### GetUseTrustStore

`func (o *ReadLdapServerXo) GetUseTrustStore() bool`

GetUseTrustStore returns the UseTrustStore field if non-nil, zero value otherwise.

### GetUseTrustStoreOk

`func (o *ReadLdapServerXo) GetUseTrustStoreOk() (*bool, bool)`

GetUseTrustStoreOk returns a tuple with the UseTrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrustStore

`func (o *ReadLdapServerXo) SetUseTrustStore(v bool)`

SetUseTrustStore sets UseTrustStore field to given value.

### HasUseTrustStore

`func (o *ReadLdapServerXo) HasUseTrustStore() bool`

HasUseTrustStore returns a boolean if a field has been set.

### GetUserBaseDn

`func (o *ReadLdapServerXo) GetUserBaseDn() string`

GetUserBaseDn returns the UserBaseDn field if non-nil, zero value otherwise.

### GetUserBaseDnOk

`func (o *ReadLdapServerXo) GetUserBaseDnOk() (*string, bool)`

GetUserBaseDnOk returns a tuple with the UserBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDn

`func (o *ReadLdapServerXo) SetUserBaseDn(v string)`

SetUserBaseDn sets UserBaseDn field to given value.

### HasUserBaseDn

`func (o *ReadLdapServerXo) HasUserBaseDn() bool`

HasUserBaseDn returns a boolean if a field has been set.

### GetUserEmailAddressAttribute

`func (o *ReadLdapServerXo) GetUserEmailAddressAttribute() string`

GetUserEmailAddressAttribute returns the UserEmailAddressAttribute field if non-nil, zero value otherwise.

### GetUserEmailAddressAttributeOk

`func (o *ReadLdapServerXo) GetUserEmailAddressAttributeOk() (*string, bool)`

GetUserEmailAddressAttributeOk returns a tuple with the UserEmailAddressAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmailAddressAttribute

`func (o *ReadLdapServerXo) SetUserEmailAddressAttribute(v string)`

SetUserEmailAddressAttribute sets UserEmailAddressAttribute field to given value.

### HasUserEmailAddressAttribute

`func (o *ReadLdapServerXo) HasUserEmailAddressAttribute() bool`

HasUserEmailAddressAttribute returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *ReadLdapServerXo) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *ReadLdapServerXo) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *ReadLdapServerXo) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *ReadLdapServerXo) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserLdapFilter

`func (o *ReadLdapServerXo) GetUserLdapFilter() string`

GetUserLdapFilter returns the UserLdapFilter field if non-nil, zero value otherwise.

### GetUserLdapFilterOk

`func (o *ReadLdapServerXo) GetUserLdapFilterOk() (*string, bool)`

GetUserLdapFilterOk returns a tuple with the UserLdapFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLdapFilter

`func (o *ReadLdapServerXo) SetUserLdapFilter(v string)`

SetUserLdapFilter sets UserLdapFilter field to given value.

### HasUserLdapFilter

`func (o *ReadLdapServerXo) HasUserLdapFilter() bool`

HasUserLdapFilter returns a boolean if a field has been set.

### GetUserMemberOfAttribute

`func (o *ReadLdapServerXo) GetUserMemberOfAttribute() string`

GetUserMemberOfAttribute returns the UserMemberOfAttribute field if non-nil, zero value otherwise.

### GetUserMemberOfAttributeOk

`func (o *ReadLdapServerXo) GetUserMemberOfAttributeOk() (*string, bool)`

GetUserMemberOfAttributeOk returns a tuple with the UserMemberOfAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMemberOfAttribute

`func (o *ReadLdapServerXo) SetUserMemberOfAttribute(v string)`

SetUserMemberOfAttribute sets UserMemberOfAttribute field to given value.

### HasUserMemberOfAttribute

`func (o *ReadLdapServerXo) HasUserMemberOfAttribute() bool`

HasUserMemberOfAttribute returns a boolean if a field has been set.

### GetUserObjectClass

`func (o *ReadLdapServerXo) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *ReadLdapServerXo) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *ReadLdapServerXo) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.

### HasUserObjectClass

`func (o *ReadLdapServerXo) HasUserObjectClass() bool`

HasUserObjectClass returns a boolean if a field has been set.

### GetUserPasswordAttribute

`func (o *ReadLdapServerXo) GetUserPasswordAttribute() string`

GetUserPasswordAttribute returns the UserPasswordAttribute field if non-nil, zero value otherwise.

### GetUserPasswordAttributeOk

`func (o *ReadLdapServerXo) GetUserPasswordAttributeOk() (*string, bool)`

GetUserPasswordAttributeOk returns a tuple with the UserPasswordAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordAttribute

`func (o *ReadLdapServerXo) SetUserPasswordAttribute(v string)`

SetUserPasswordAttribute sets UserPasswordAttribute field to given value.

### HasUserPasswordAttribute

`func (o *ReadLdapServerXo) HasUserPasswordAttribute() bool`

HasUserPasswordAttribute returns a boolean if a field has been set.

### GetUserRealNameAttribute

`func (o *ReadLdapServerXo) GetUserRealNameAttribute() string`

GetUserRealNameAttribute returns the UserRealNameAttribute field if non-nil, zero value otherwise.

### GetUserRealNameAttributeOk

`func (o *ReadLdapServerXo) GetUserRealNameAttributeOk() (*string, bool)`

GetUserRealNameAttributeOk returns a tuple with the UserRealNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRealNameAttribute

`func (o *ReadLdapServerXo) SetUserRealNameAttribute(v string)`

SetUserRealNameAttribute sets UserRealNameAttribute field to given value.

### HasUserRealNameAttribute

`func (o *ReadLdapServerXo) HasUserRealNameAttribute() bool`

HasUserRealNameAttribute returns a boolean if a field has been set.

### GetUserSubtree

`func (o *ReadLdapServerXo) GetUserSubtree() bool`

GetUserSubtree returns the UserSubtree field if non-nil, zero value otherwise.

### GetUserSubtreeOk

`func (o *ReadLdapServerXo) GetUserSubtreeOk() (*bool, bool)`

GetUserSubtreeOk returns a tuple with the UserSubtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSubtree

`func (o *ReadLdapServerXo) SetUserSubtree(v bool)`

SetUserSubtree sets UserSubtree field to given value.

### HasUserSubtree

`func (o *ReadLdapServerXo) HasUserSubtree() bool`

HasUserSubtree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


