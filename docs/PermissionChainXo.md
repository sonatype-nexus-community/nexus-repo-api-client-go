# PermissionChainXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to **string** | The permission string | [optional] 
**Privilege** | Pointer to [**EntityRefXo**](EntityRefXo.md) |  | [optional] 
**Role** | Pointer to [**EntityRefXo**](EntityRefXo.md) |  | [optional] 
**User** | Pointer to [**EntityRefXo**](EntityRefXo.md) |  | [optional] 

## Methods

### NewPermissionChainXo

`func NewPermissionChainXo() *PermissionChainXo`

NewPermissionChainXo instantiates a new PermissionChainXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionChainXoWithDefaults

`func NewPermissionChainXoWithDefaults() *PermissionChainXo`

NewPermissionChainXoWithDefaults instantiates a new PermissionChainXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *PermissionChainXo) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionChainXo) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionChainXo) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *PermissionChainXo) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPrivilege

`func (o *PermissionChainXo) GetPrivilege() EntityRefXo`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *PermissionChainXo) GetPrivilegeOk() (*EntityRefXo, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *PermissionChainXo) SetPrivilege(v EntityRefXo)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *PermissionChainXo) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetRole

`func (o *PermissionChainXo) GetRole() EntityRefXo`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PermissionChainXo) GetRoleOk() (*EntityRefXo, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PermissionChainXo) SetRole(v EntityRefXo)`

SetRole sets Role field to given value.

### HasRole

`func (o *PermissionChainXo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *PermissionChainXo) GetUser() EntityRefXo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PermissionChainXo) GetUserOk() (*EntityRefXo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PermissionChainXo) SetUser(v EntityRefXo)`

SetUser sets User field to given value.

### HasUser

`func (o *PermissionChainXo) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


