# ApiPrivilegeApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | 
**Description** | Pointer to **string** |  | [optional] 
**Domain** | **string** | The domain (i.e. &#39;blobstores&#39;, &#39;capabilities&#39; or even &#39;*&#39; for all) that this privilege is granting access to.  Note that creating new privileges with a domain is only necessary when using plugins that define their own domain(s). | 
**Name** | **string** | The name of the privilege.  This value cannot be changed. | 
**ReadOnly** | Pointer to **bool** | Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system. | [optional] 
**Type** | Pointer to **string** | The type of privilege, each type covers different portions of the system. External values supplied to this will be ignored by the system. | [optional] 

## Methods

### NewApiPrivilegeApplication

`func NewApiPrivilegeApplication(actions []string, domain string, name string, ) *ApiPrivilegeApplication`

NewApiPrivilegeApplication instantiates a new ApiPrivilegeApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeApplicationWithDefaults

`func NewApiPrivilegeApplicationWithDefaults() *ApiPrivilegeApplication`

NewApiPrivilegeApplicationWithDefaults instantiates a new ApiPrivilegeApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApiPrivilegeApplication) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPrivilegeApplication) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPrivilegeApplication) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetDescription

`func (o *ApiPrivilegeApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *ApiPrivilegeApplication) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiPrivilegeApplication) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiPrivilegeApplication) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *ApiPrivilegeApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeApplication) SetName(v string)`

SetName sets Name field to given value.


### GetReadOnly

`func (o *ApiPrivilegeApplication) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApiPrivilegeApplication) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApiPrivilegeApplication) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApiPrivilegeApplication) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetType

`func (o *ApiPrivilegeApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPrivilegeApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPrivilegeApplication) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiPrivilegeApplication) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


