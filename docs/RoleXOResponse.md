# RoleXOResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of this role. | [optional] 
**Id** | Pointer to **string** | The id of the role. | [optional] 
**Name** | Pointer to **string** | The name of the role. | [optional] 
**Privileges** | Pointer to **[]string** | The list of privileges assigned to this role. | [optional] 
**ReadOnly** | Pointer to **bool** | Indicates whether the role can be changed. The system will ignore any supplied external values. | [optional] 
**Roles** | Pointer to **[]string** | The list of roles assigned to this role. | [optional] 
**Source** | Pointer to **string** | The user source which is the origin of this role. | [optional] 

## Methods

### NewRoleXOResponse

`func NewRoleXOResponse() *RoleXOResponse`

NewRoleXOResponse instantiates a new RoleXOResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleXOResponseWithDefaults

`func NewRoleXOResponseWithDefaults() *RoleXOResponse`

NewRoleXOResponseWithDefaults instantiates a new RoleXOResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoleXOResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleXOResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleXOResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleXOResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *RoleXOResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleXOResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleXOResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleXOResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoleXOResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleXOResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleXOResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleXOResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivileges

`func (o *RoleXOResponse) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *RoleXOResponse) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *RoleXOResponse) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *RoleXOResponse) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetReadOnly

`func (o *RoleXOResponse) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *RoleXOResponse) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *RoleXOResponse) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *RoleXOResponse) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRoles

`func (o *RoleXOResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleXOResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleXOResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *RoleXOResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSource

`func (o *RoleXOResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RoleXOResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RoleXOResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RoleXOResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


