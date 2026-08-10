# ApiPrivilegeRepositoryContentSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | 
**ContentSelector** | **string** | The name of a content selector that will be used to grant access to content via this privilege. | 
**Description** | Pointer to **string** |  | [optional] 
**Format** | **string** | The repository format (i.e &#39;nuget&#39;, &#39;npm&#39;) this privilege will grant access to (or * for all). | 
**Name** | **string** | The name of the privilege.  This value cannot be changed. | 
**ReadOnly** | Pointer to **bool** | Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system. | [optional] 
**Repository** | **string** | The name of the repository this privilege will grant access to (or * for all). | 
**Type** | Pointer to **string** | The type of privilege, each type covers different portions of the system. External values supplied to this will be ignored by the system. | [optional] 

## Methods

### NewApiPrivilegeRepositoryContentSelector

`func NewApiPrivilegeRepositoryContentSelector(actions []string, contentSelector string, format string, name string, repository string, ) *ApiPrivilegeRepositoryContentSelector`

NewApiPrivilegeRepositoryContentSelector instantiates a new ApiPrivilegeRepositoryContentSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeRepositoryContentSelectorWithDefaults

`func NewApiPrivilegeRepositoryContentSelectorWithDefaults() *ApiPrivilegeRepositoryContentSelector`

NewApiPrivilegeRepositoryContentSelectorWithDefaults instantiates a new ApiPrivilegeRepositoryContentSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApiPrivilegeRepositoryContentSelector) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPrivilegeRepositoryContentSelector) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPrivilegeRepositoryContentSelector) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetContentSelector

`func (o *ApiPrivilegeRepositoryContentSelector) GetContentSelector() string`

GetContentSelector returns the ContentSelector field if non-nil, zero value otherwise.

### GetContentSelectorOk

`func (o *ApiPrivilegeRepositoryContentSelector) GetContentSelectorOk() (*string, bool)`

GetContentSelectorOk returns a tuple with the ContentSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSelector

`func (o *ApiPrivilegeRepositoryContentSelector) SetContentSelector(v string)`

SetContentSelector sets ContentSelector field to given value.


### GetDescription

`func (o *ApiPrivilegeRepositoryContentSelector) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeRepositoryContentSelector) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeRepositoryContentSelector) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeRepositoryContentSelector) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *ApiPrivilegeRepositoryContentSelector) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ApiPrivilegeRepositoryContentSelector) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ApiPrivilegeRepositoryContentSelector) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetName

`func (o *ApiPrivilegeRepositoryContentSelector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeRepositoryContentSelector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeRepositoryContentSelector) SetName(v string)`

SetName sets Name field to given value.


### GetReadOnly

`func (o *ApiPrivilegeRepositoryContentSelector) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApiPrivilegeRepositoryContentSelector) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApiPrivilegeRepositoryContentSelector) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApiPrivilegeRepositoryContentSelector) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRepository

`func (o *ApiPrivilegeRepositoryContentSelector) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiPrivilegeRepositoryContentSelector) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiPrivilegeRepositoryContentSelector) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetType

`func (o *ApiPrivilegeRepositoryContentSelector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPrivilegeRepositoryContentSelector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPrivilegeRepositoryContentSelector) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiPrivilegeRepositoryContentSelector) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


