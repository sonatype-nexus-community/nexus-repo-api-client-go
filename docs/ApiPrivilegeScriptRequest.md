# ApiPrivilegeScriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the privilege.  This value cannot be changed. | [optional] 
**ScriptName** | Pointer to **string** | The name of a script to give access to. | [optional] 

## Methods

### NewApiPrivilegeScriptRequest

`func NewApiPrivilegeScriptRequest() *ApiPrivilegeScriptRequest`

NewApiPrivilegeScriptRequest instantiates a new ApiPrivilegeScriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeScriptRequestWithDefaults

`func NewApiPrivilegeScriptRequestWithDefaults() *ApiPrivilegeScriptRequest`

NewApiPrivilegeScriptRequestWithDefaults instantiates a new ApiPrivilegeScriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApiPrivilegeScriptRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPrivilegeScriptRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPrivilegeScriptRequest) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApiPrivilegeScriptRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDescription

`func (o *ApiPrivilegeScriptRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeScriptRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeScriptRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeScriptRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ApiPrivilegeScriptRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeScriptRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeScriptRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPrivilegeScriptRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScriptName

`func (o *ApiPrivilegeScriptRequest) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *ApiPrivilegeScriptRequest) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *ApiPrivilegeScriptRequest) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *ApiPrivilegeScriptRequest) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


