# ApiPrivilegeRepositoryContentSelectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | [optional] 
**ContentSelector** | Pointer to **string** | The name of a content selector that will be used to grant access to content via this privilege. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** | The repository format (i.e &#39;nuget&#39;, &#39;npm&#39;) this privilege will grant access to (or * for all). | [optional] 
**Name** | Pointer to **string** | The name of the privilege.  This value cannot be changed. | [optional] 
**Repository** | Pointer to **string** | The name of the repository this privilege will grant access to (or * for all). | [optional] 

## Methods

### NewApiPrivilegeRepositoryContentSelectorRequest

`func NewApiPrivilegeRepositoryContentSelectorRequest() *ApiPrivilegeRepositoryContentSelectorRequest`

NewApiPrivilegeRepositoryContentSelectorRequest instantiates a new ApiPrivilegeRepositoryContentSelectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeRepositoryContentSelectorRequestWithDefaults

`func NewApiPrivilegeRepositoryContentSelectorRequestWithDefaults() *ApiPrivilegeRepositoryContentSelectorRequest`

NewApiPrivilegeRepositoryContentSelectorRequestWithDefaults instantiates a new ApiPrivilegeRepositoryContentSelectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetContentSelector

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetContentSelector() string`

GetContentSelector returns the ContentSelector field if non-nil, zero value otherwise.

### GetContentSelectorOk

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetContentSelectorOk() (*string, bool)`

GetContentSelectorOk returns a tuple with the ContentSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSelector

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetContentSelector(v string)`

SetContentSelector sets ContentSelector field to given value.

### HasContentSelector

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasContentSelector() bool`

HasContentSelector returns a boolean if a field has been set.

### GetDescription

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ApiPrivilegeRepositoryContentSelectorRequest) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


