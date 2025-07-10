# GetAllPrivileges200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** | The domain (i.e. &#39;blobstores&#39;, &#39;capabilities&#39; or even &#39;*&#39; for all) that this privilege is granting access to.  Note that creating new privileges with a domain is only necessary when using plugins that define their own domain(s). | [optional] 
**Name** | Pointer to **string** | The name of the privilege.  This value cannot be changed. | [optional] 
**Format** | Pointer to **string** | The repository format (i.e &#39;nuget&#39;, &#39;npm&#39;) this privilege will grant access to (or * for all). | [optional] 
**Repository** | Pointer to **string** | The name of the repository this privilege will grant access to (or * for all). | [optional] 
**ContentSelector** | Pointer to **string** | The name of a content selector that will be used to grant access to content via this privilege. | [optional] 
**ScriptName** | Pointer to **string** | The name of a script to give access to. | [optional] 
**Pattern** | Pointer to **string** | A colon separated list of parts that create a permission string. | [optional] 
**ReadOnly** | Pointer to **bool** | Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system. | [optional] 
**Type** | Pointer to **string** | The type of privilege, each type covers different portions of the system. External values supplied to this will be ignored by the system. | [optional] 

## Methods

### NewGetAllPrivileges200ResponseInner

`func NewGetAllPrivileges200ResponseInner() *GetAllPrivileges200ResponseInner`

NewGetAllPrivileges200ResponseInner instantiates a new GetAllPrivileges200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllPrivileges200ResponseInnerWithDefaults

`func NewGetAllPrivileges200ResponseInnerWithDefaults() *GetAllPrivileges200ResponseInner`

NewGetAllPrivileges200ResponseInnerWithDefaults instantiates a new GetAllPrivileges200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *GetAllPrivileges200ResponseInner) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetAllPrivileges200ResponseInner) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetAllPrivileges200ResponseInner) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GetAllPrivileges200ResponseInner) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDescription

`func (o *GetAllPrivileges200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAllPrivileges200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAllPrivileges200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAllPrivileges200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *GetAllPrivileges200ResponseInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetAllPrivileges200ResponseInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetAllPrivileges200ResponseInner) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetAllPrivileges200ResponseInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *GetAllPrivileges200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAllPrivileges200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAllPrivileges200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAllPrivileges200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *GetAllPrivileges200ResponseInner) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GetAllPrivileges200ResponseInner) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GetAllPrivileges200ResponseInner) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GetAllPrivileges200ResponseInner) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetRepository

`func (o *GetAllPrivileges200ResponseInner) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GetAllPrivileges200ResponseInner) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GetAllPrivileges200ResponseInner) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GetAllPrivileges200ResponseInner) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetContentSelector

`func (o *GetAllPrivileges200ResponseInner) GetContentSelector() string`

GetContentSelector returns the ContentSelector field if non-nil, zero value otherwise.

### GetContentSelectorOk

`func (o *GetAllPrivileges200ResponseInner) GetContentSelectorOk() (*string, bool)`

GetContentSelectorOk returns a tuple with the ContentSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSelector

`func (o *GetAllPrivileges200ResponseInner) SetContentSelector(v string)`

SetContentSelector sets ContentSelector field to given value.

### HasContentSelector

`func (o *GetAllPrivileges200ResponseInner) HasContentSelector() bool`

HasContentSelector returns a boolean if a field has been set.

### GetScriptName

`func (o *GetAllPrivileges200ResponseInner) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *GetAllPrivileges200ResponseInner) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *GetAllPrivileges200ResponseInner) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *GetAllPrivileges200ResponseInner) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetPattern

`func (o *GetAllPrivileges200ResponseInner) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *GetAllPrivileges200ResponseInner) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *GetAllPrivileges200ResponseInner) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *GetAllPrivileges200ResponseInner) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetReadOnly

`func (o *GetAllPrivileges200ResponseInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *GetAllPrivileges200ResponseInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *GetAllPrivileges200ResponseInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *GetAllPrivileges200ResponseInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetType

`func (o *GetAllPrivileges200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAllPrivileges200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAllPrivileges200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAllPrivileges200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


