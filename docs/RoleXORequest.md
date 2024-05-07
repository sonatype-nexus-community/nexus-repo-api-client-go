# RoleXORequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of this role. | [optional] 
**Id** | Pointer to **string** | The id of the role. | [optional] 
**Name** | Pointer to **string** | The name of the role. | [optional] 
**Privileges** | Pointer to **[]string** | The list of privileges assigned to this role. | [optional] 
**Roles** | Pointer to **[]string** | The list of roles assigned to this role. | [optional] 

## Methods

### NewRoleXORequest

`func NewRoleXORequest() *RoleXORequest`

NewRoleXORequest instantiates a new RoleXORequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleXORequestWithDefaults

`func NewRoleXORequestWithDefaults() *RoleXORequest`

NewRoleXORequestWithDefaults instantiates a new RoleXORequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoleXORequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleXORequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleXORequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleXORequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *RoleXORequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleXORequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleXORequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleXORequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoleXORequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleXORequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleXORequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleXORequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivileges

`func (o *RoleXORequest) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *RoleXORequest) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *RoleXORequest) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *RoleXORequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetRoles

`func (o *RoleXORequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleXORequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleXORequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *RoleXORequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


