# ApiPrivilegeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges.  Only applicable for Privilege Types: repository-admin, repository-content-selector, repository-view, script | [optional] 
**ContentSelector** | Pointer to **string** | The name of a content selector that will be used to grant access to content via this privilege.  Only applicable for Privilege Types: repository-content-selector | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** | The domain (i.e. &#39;blobstores&#39;, &#39;capabilities&#39; or even &#39;*&#39; for all) that this privilege is granting access to. Note that creating new privileges with a domain is only necessary when using plugins that define their own domain(s).  Only applicable for Privilege Types: application | [optional] 
**Format** | Pointer to **string** | The repository format (i.e &#39;nuget&#39;, &#39;npm&#39;) this privilege will grant access to (or * for all).  Only applicable for Privilege Types: repository-admin, repository-content-selector, repository-view | [optional] 
**Name** | **string** | The name of the privilege. This value cannot be changed. | 
**Pattern** | Pointer to **string** | A colon separated list of parts that create a permission string.  Only applicable for Privilege Types: wildcard | [optional] 
**ReadOnly** | Pointer to **bool** | Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system. | [optional] 
**Repository** | Pointer to **string** | The name of the repository this privilege will grant access to (or * for all). The repository MUST exist.  Only applicable for Privilege Types: repository-admin, repository-content-selector, repository-view | [optional] 
**ScriptName** | Pointer to **string** | The name of a script to give access to.  Only applicable for Privilege Types: script | [optional] 
**Type** | **string** | The type of privilege, each type covers different portion of the system. External values supplied to this will be ignored by the system. | 

## Methods

### NewApiPrivilegeRequest

`func NewApiPrivilegeRequest(name string, type_ string, ) *ApiPrivilegeRequest`

NewApiPrivilegeRequest instantiates a new ApiPrivilegeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeRequestWithDefaults

`func NewApiPrivilegeRequestWithDefaults() *ApiPrivilegeRequest`

NewApiPrivilegeRequestWithDefaults instantiates a new ApiPrivilegeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApiPrivilegeRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPrivilegeRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPrivilegeRequest) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApiPrivilegeRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetContentSelector

`func (o *ApiPrivilegeRequest) GetContentSelector() string`

GetContentSelector returns the ContentSelector field if non-nil, zero value otherwise.

### GetContentSelectorOk

`func (o *ApiPrivilegeRequest) GetContentSelectorOk() (*string, bool)`

GetContentSelectorOk returns a tuple with the ContentSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSelector

`func (o *ApiPrivilegeRequest) SetContentSelector(v string)`

SetContentSelector sets ContentSelector field to given value.

### HasContentSelector

`func (o *ApiPrivilegeRequest) HasContentSelector() bool`

HasContentSelector returns a boolean if a field has been set.

### GetDescription

`func (o *ApiPrivilegeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *ApiPrivilegeRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiPrivilegeRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiPrivilegeRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiPrivilegeRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFormat

`func (o *ApiPrivilegeRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ApiPrivilegeRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ApiPrivilegeRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ApiPrivilegeRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *ApiPrivilegeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPattern

`func (o *ApiPrivilegeRequest) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ApiPrivilegeRequest) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ApiPrivilegeRequest) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *ApiPrivilegeRequest) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetReadOnly

`func (o *ApiPrivilegeRequest) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApiPrivilegeRequest) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApiPrivilegeRequest) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApiPrivilegeRequest) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRepository

`func (o *ApiPrivilegeRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiPrivilegeRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiPrivilegeRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ApiPrivilegeRequest) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetScriptName

`func (o *ApiPrivilegeRequest) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *ApiPrivilegeRequest) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *ApiPrivilegeRequest) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *ApiPrivilegeRequest) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetType

`func (o *ApiPrivilegeRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPrivilegeRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPrivilegeRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


