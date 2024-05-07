# ApiPrivilegeRepositoryAdminRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** | The repository format (i.e &#39;nuget&#39;, &#39;npm&#39;) this privilege will grant access to (or * for all). | [optional] 
**Name** | Pointer to **string** | The name of the privilege.  This value cannot be changed. | [optional] 
**Repository** | Pointer to **string** | The name of the repository this privilege will grant access to (or * for all). | [optional] 

## Methods

### NewApiPrivilegeRepositoryAdminRequest

`func NewApiPrivilegeRepositoryAdminRequest() *ApiPrivilegeRepositoryAdminRequest`

NewApiPrivilegeRepositoryAdminRequest instantiates a new ApiPrivilegeRepositoryAdminRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeRepositoryAdminRequestWithDefaults

`func NewApiPrivilegeRepositoryAdminRequestWithDefaults() *ApiPrivilegeRepositoryAdminRequest`

NewApiPrivilegeRepositoryAdminRequestWithDefaults instantiates a new ApiPrivilegeRepositoryAdminRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApiPrivilegeRepositoryAdminRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPrivilegeRepositoryAdminRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPrivilegeRepositoryAdminRequest) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApiPrivilegeRepositoryAdminRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDescription

`func (o *ApiPrivilegeRepositoryAdminRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeRepositoryAdminRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeRepositoryAdminRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeRepositoryAdminRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *ApiPrivilegeRepositoryAdminRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ApiPrivilegeRepositoryAdminRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ApiPrivilegeRepositoryAdminRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ApiPrivilegeRepositoryAdminRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *ApiPrivilegeRepositoryAdminRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeRepositoryAdminRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeRepositoryAdminRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPrivilegeRepositoryAdminRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *ApiPrivilegeRepositoryAdminRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiPrivilegeRepositoryAdminRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiPrivilegeRepositoryAdminRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ApiPrivilegeRepositoryAdminRequest) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


