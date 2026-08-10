# ApiPrivilegeRepositoryAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | 
**Description** | Pointer to **string** |  | [optional] 
**Format** | **string** | The repository format (i.e &#39;nuget&#39;, &#39;npm&#39;) this privilege will grant access to (or * for all). | 
**Name** | **string** | The name of the privilege.  This value cannot be changed. | 
**ReadOnly** | Pointer to **bool** | Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system. | [optional] 
**Repository** | **string** | The name of the repository this privilege will grant access to (or * for all). | 
**Type** | Pointer to **string** | The type of privilege, each type covers different portions of the system. External values supplied to this will be ignored by the system. | [optional] 

## Methods

### NewApiPrivilegeRepositoryAdmin

`func NewApiPrivilegeRepositoryAdmin(actions []string, format string, name string, repository string, ) *ApiPrivilegeRepositoryAdmin`

NewApiPrivilegeRepositoryAdmin instantiates a new ApiPrivilegeRepositoryAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeRepositoryAdminWithDefaults

`func NewApiPrivilegeRepositoryAdminWithDefaults() *ApiPrivilegeRepositoryAdmin`

NewApiPrivilegeRepositoryAdminWithDefaults instantiates a new ApiPrivilegeRepositoryAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApiPrivilegeRepositoryAdmin) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPrivilegeRepositoryAdmin) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPrivilegeRepositoryAdmin) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetDescription

`func (o *ApiPrivilegeRepositoryAdmin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeRepositoryAdmin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeRepositoryAdmin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeRepositoryAdmin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *ApiPrivilegeRepositoryAdmin) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ApiPrivilegeRepositoryAdmin) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ApiPrivilegeRepositoryAdmin) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetName

`func (o *ApiPrivilegeRepositoryAdmin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeRepositoryAdmin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeRepositoryAdmin) SetName(v string)`

SetName sets Name field to given value.


### GetReadOnly

`func (o *ApiPrivilegeRepositoryAdmin) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApiPrivilegeRepositoryAdmin) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApiPrivilegeRepositoryAdmin) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApiPrivilegeRepositoryAdmin) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRepository

`func (o *ApiPrivilegeRepositoryAdmin) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiPrivilegeRepositoryAdmin) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiPrivilegeRepositoryAdmin) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetType

`func (o *ApiPrivilegeRepositoryAdmin) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPrivilegeRepositoryAdmin) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPrivilegeRepositoryAdmin) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiPrivilegeRepositoryAdmin) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


