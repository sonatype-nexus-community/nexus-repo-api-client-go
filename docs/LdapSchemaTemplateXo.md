# LdapSchemaTemplateXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupBaseDn** | Pointer to **string** | The relative DN where groups are found | [optional] 
**GroupIdAttribute** | Pointer to **string** | LDAP attribute containing the group ID | [optional] 
**GroupMemberAttribute** | Pointer to **string** | LDAP attribute containing group member references | [optional] 
**GroupMemberFormat** | Pointer to **string** | Format of group member attribute values | [optional] 
**GroupObjectClass** | Pointer to **string** | LDAP object class for groups | [optional] 
**GroupSubtree** | Pointer to **bool** | Whether to include subtrees when searching for groups | [optional] 
**GroupType** | Pointer to **string** | Group type: &#39;static&#39; or &#39;dynamic&#39; | [optional] 
**LdapGroupsAsRoles** | Pointer to **bool** | Whether LDAP groups should be mapped as roles | [optional] 
**Name** | Pointer to **string** | Template name | [optional] 
**UserBaseDn** | Pointer to **string** | The relative DN where users are found | [optional] 
**UserEmailAddressAttribute** | Pointer to **string** | LDAP attribute containing the user&#39;s email address | [optional] 
**UserIdAttribute** | Pointer to **string** | LDAP attribute containing the user ID | [optional] 
**UserLdapFilter** | Pointer to **string** | Optional LDAP filter to limit user search | [optional] 
**UserMemberOfAttribute** | Pointer to **string** | LDAP attribute on user containing group membership (for dynamic groups) | [optional] 
**UserObjectClass** | Pointer to **string** | LDAP object class for users | [optional] 
**UserPasswordAttribute** | Pointer to **string** | LDAP attribute containing the user&#39;s password | [optional] 
**UserRealNameAttribute** | Pointer to **string** | LDAP attribute containing the user&#39;s real name | [optional] 
**UserSubtree** | Pointer to **bool** | Whether to include subtrees when searching for users | [optional] 

## Methods

### NewLdapSchemaTemplateXo

`func NewLdapSchemaTemplateXo() *LdapSchemaTemplateXo`

NewLdapSchemaTemplateXo instantiates a new LdapSchemaTemplateXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapSchemaTemplateXoWithDefaults

`func NewLdapSchemaTemplateXoWithDefaults() *LdapSchemaTemplateXo`

NewLdapSchemaTemplateXoWithDefaults instantiates a new LdapSchemaTemplateXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupBaseDn

`func (o *LdapSchemaTemplateXo) GetGroupBaseDn() string`

GetGroupBaseDn returns the GroupBaseDn field if non-nil, zero value otherwise.

### GetGroupBaseDnOk

`func (o *LdapSchemaTemplateXo) GetGroupBaseDnOk() (*string, bool)`

GetGroupBaseDnOk returns a tuple with the GroupBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBaseDn

`func (o *LdapSchemaTemplateXo) SetGroupBaseDn(v string)`

SetGroupBaseDn sets GroupBaseDn field to given value.

### HasGroupBaseDn

`func (o *LdapSchemaTemplateXo) HasGroupBaseDn() bool`

HasGroupBaseDn returns a boolean if a field has been set.

### GetGroupIdAttribute

`func (o *LdapSchemaTemplateXo) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *LdapSchemaTemplateXo) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *LdapSchemaTemplateXo) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.

### HasGroupIdAttribute

`func (o *LdapSchemaTemplateXo) HasGroupIdAttribute() bool`

HasGroupIdAttribute returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *LdapSchemaTemplateXo) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *LdapSchemaTemplateXo) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *LdapSchemaTemplateXo) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *LdapSchemaTemplateXo) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberFormat

`func (o *LdapSchemaTemplateXo) GetGroupMemberFormat() string`

GetGroupMemberFormat returns the GroupMemberFormat field if non-nil, zero value otherwise.

### GetGroupMemberFormatOk

`func (o *LdapSchemaTemplateXo) GetGroupMemberFormatOk() (*string, bool)`

GetGroupMemberFormatOk returns a tuple with the GroupMemberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberFormat

`func (o *LdapSchemaTemplateXo) SetGroupMemberFormat(v string)`

SetGroupMemberFormat sets GroupMemberFormat field to given value.

### HasGroupMemberFormat

`func (o *LdapSchemaTemplateXo) HasGroupMemberFormat() bool`

HasGroupMemberFormat returns a boolean if a field has been set.

### GetGroupObjectClass

`func (o *LdapSchemaTemplateXo) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *LdapSchemaTemplateXo) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *LdapSchemaTemplateXo) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.

### HasGroupObjectClass

`func (o *LdapSchemaTemplateXo) HasGroupObjectClass() bool`

HasGroupObjectClass returns a boolean if a field has been set.

### GetGroupSubtree

`func (o *LdapSchemaTemplateXo) GetGroupSubtree() bool`

GetGroupSubtree returns the GroupSubtree field if non-nil, zero value otherwise.

### GetGroupSubtreeOk

`func (o *LdapSchemaTemplateXo) GetGroupSubtreeOk() (*bool, bool)`

GetGroupSubtreeOk returns a tuple with the GroupSubtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSubtree

`func (o *LdapSchemaTemplateXo) SetGroupSubtree(v bool)`

SetGroupSubtree sets GroupSubtree field to given value.

### HasGroupSubtree

`func (o *LdapSchemaTemplateXo) HasGroupSubtree() bool`

HasGroupSubtree returns a boolean if a field has been set.

### GetGroupType

`func (o *LdapSchemaTemplateXo) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *LdapSchemaTemplateXo) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *LdapSchemaTemplateXo) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *LdapSchemaTemplateXo) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetLdapGroupsAsRoles

`func (o *LdapSchemaTemplateXo) GetLdapGroupsAsRoles() bool`

GetLdapGroupsAsRoles returns the LdapGroupsAsRoles field if non-nil, zero value otherwise.

### GetLdapGroupsAsRolesOk

`func (o *LdapSchemaTemplateXo) GetLdapGroupsAsRolesOk() (*bool, bool)`

GetLdapGroupsAsRolesOk returns a tuple with the LdapGroupsAsRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupsAsRoles

`func (o *LdapSchemaTemplateXo) SetLdapGroupsAsRoles(v bool)`

SetLdapGroupsAsRoles sets LdapGroupsAsRoles field to given value.

### HasLdapGroupsAsRoles

`func (o *LdapSchemaTemplateXo) HasLdapGroupsAsRoles() bool`

HasLdapGroupsAsRoles returns a boolean if a field has been set.

### GetName

`func (o *LdapSchemaTemplateXo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapSchemaTemplateXo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapSchemaTemplateXo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapSchemaTemplateXo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserBaseDn

`func (o *LdapSchemaTemplateXo) GetUserBaseDn() string`

GetUserBaseDn returns the UserBaseDn field if non-nil, zero value otherwise.

### GetUserBaseDnOk

`func (o *LdapSchemaTemplateXo) GetUserBaseDnOk() (*string, bool)`

GetUserBaseDnOk returns a tuple with the UserBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDn

`func (o *LdapSchemaTemplateXo) SetUserBaseDn(v string)`

SetUserBaseDn sets UserBaseDn field to given value.

### HasUserBaseDn

`func (o *LdapSchemaTemplateXo) HasUserBaseDn() bool`

HasUserBaseDn returns a boolean if a field has been set.

### GetUserEmailAddressAttribute

`func (o *LdapSchemaTemplateXo) GetUserEmailAddressAttribute() string`

GetUserEmailAddressAttribute returns the UserEmailAddressAttribute field if non-nil, zero value otherwise.

### GetUserEmailAddressAttributeOk

`func (o *LdapSchemaTemplateXo) GetUserEmailAddressAttributeOk() (*string, bool)`

GetUserEmailAddressAttributeOk returns a tuple with the UserEmailAddressAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmailAddressAttribute

`func (o *LdapSchemaTemplateXo) SetUserEmailAddressAttribute(v string)`

SetUserEmailAddressAttribute sets UserEmailAddressAttribute field to given value.

### HasUserEmailAddressAttribute

`func (o *LdapSchemaTemplateXo) HasUserEmailAddressAttribute() bool`

HasUserEmailAddressAttribute returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *LdapSchemaTemplateXo) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *LdapSchemaTemplateXo) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *LdapSchemaTemplateXo) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *LdapSchemaTemplateXo) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserLdapFilter

`func (o *LdapSchemaTemplateXo) GetUserLdapFilter() string`

GetUserLdapFilter returns the UserLdapFilter field if non-nil, zero value otherwise.

### GetUserLdapFilterOk

`func (o *LdapSchemaTemplateXo) GetUserLdapFilterOk() (*string, bool)`

GetUserLdapFilterOk returns a tuple with the UserLdapFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLdapFilter

`func (o *LdapSchemaTemplateXo) SetUserLdapFilter(v string)`

SetUserLdapFilter sets UserLdapFilter field to given value.

### HasUserLdapFilter

`func (o *LdapSchemaTemplateXo) HasUserLdapFilter() bool`

HasUserLdapFilter returns a boolean if a field has been set.

### GetUserMemberOfAttribute

`func (o *LdapSchemaTemplateXo) GetUserMemberOfAttribute() string`

GetUserMemberOfAttribute returns the UserMemberOfAttribute field if non-nil, zero value otherwise.

### GetUserMemberOfAttributeOk

`func (o *LdapSchemaTemplateXo) GetUserMemberOfAttributeOk() (*string, bool)`

GetUserMemberOfAttributeOk returns a tuple with the UserMemberOfAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMemberOfAttribute

`func (o *LdapSchemaTemplateXo) SetUserMemberOfAttribute(v string)`

SetUserMemberOfAttribute sets UserMemberOfAttribute field to given value.

### HasUserMemberOfAttribute

`func (o *LdapSchemaTemplateXo) HasUserMemberOfAttribute() bool`

HasUserMemberOfAttribute returns a boolean if a field has been set.

### GetUserObjectClass

`func (o *LdapSchemaTemplateXo) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *LdapSchemaTemplateXo) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *LdapSchemaTemplateXo) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.

### HasUserObjectClass

`func (o *LdapSchemaTemplateXo) HasUserObjectClass() bool`

HasUserObjectClass returns a boolean if a field has been set.

### GetUserPasswordAttribute

`func (o *LdapSchemaTemplateXo) GetUserPasswordAttribute() string`

GetUserPasswordAttribute returns the UserPasswordAttribute field if non-nil, zero value otherwise.

### GetUserPasswordAttributeOk

`func (o *LdapSchemaTemplateXo) GetUserPasswordAttributeOk() (*string, bool)`

GetUserPasswordAttributeOk returns a tuple with the UserPasswordAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordAttribute

`func (o *LdapSchemaTemplateXo) SetUserPasswordAttribute(v string)`

SetUserPasswordAttribute sets UserPasswordAttribute field to given value.

### HasUserPasswordAttribute

`func (o *LdapSchemaTemplateXo) HasUserPasswordAttribute() bool`

HasUserPasswordAttribute returns a boolean if a field has been set.

### GetUserRealNameAttribute

`func (o *LdapSchemaTemplateXo) GetUserRealNameAttribute() string`

GetUserRealNameAttribute returns the UserRealNameAttribute field if non-nil, zero value otherwise.

### GetUserRealNameAttributeOk

`func (o *LdapSchemaTemplateXo) GetUserRealNameAttributeOk() (*string, bool)`

GetUserRealNameAttributeOk returns a tuple with the UserRealNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRealNameAttribute

`func (o *LdapSchemaTemplateXo) SetUserRealNameAttribute(v string)`

SetUserRealNameAttribute sets UserRealNameAttribute field to given value.

### HasUserRealNameAttribute

`func (o *LdapSchemaTemplateXo) HasUserRealNameAttribute() bool`

HasUserRealNameAttribute returns a boolean if a field has been set.

### GetUserSubtree

`func (o *LdapSchemaTemplateXo) GetUserSubtree() bool`

GetUserSubtree returns the UserSubtree field if non-nil, zero value otherwise.

### GetUserSubtreeOk

`func (o *LdapSchemaTemplateXo) GetUserSubtreeOk() (*bool, bool)`

GetUserSubtreeOk returns a tuple with the UserSubtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSubtree

`func (o *LdapSchemaTemplateXo) SetUserSubtree(v bool)`

SetUserSubtree sets UserSubtree field to given value.

### HasUserSubtree

`func (o *LdapSchemaTemplateXo) HasUserSubtree() bool`

HasUserSubtree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


