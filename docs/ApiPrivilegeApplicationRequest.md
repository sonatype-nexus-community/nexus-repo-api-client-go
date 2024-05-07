# ApiPrivilegeApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** | The domain (i.e. &#39;blobstores&#39;, &#39;capabilities&#39; or even &#39;*&#39; for all) that this privilege is granting access to.  Note that creating new privileges with a domain is only necessary when using plugins that define their own domain(s). | [optional] 
**Name** | Pointer to **string** | The name of the privilege.  This value cannot be changed. | [optional] 

## Methods

### NewApiPrivilegeApplicationRequest

`func NewApiPrivilegeApplicationRequest() *ApiPrivilegeApplicationRequest`

NewApiPrivilegeApplicationRequest instantiates a new ApiPrivilegeApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeApplicationRequestWithDefaults

`func NewApiPrivilegeApplicationRequestWithDefaults() *ApiPrivilegeApplicationRequest`

NewApiPrivilegeApplicationRequestWithDefaults instantiates a new ApiPrivilegeApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApiPrivilegeApplicationRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPrivilegeApplicationRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPrivilegeApplicationRequest) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApiPrivilegeApplicationRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDescription

`func (o *ApiPrivilegeApplicationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeApplicationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeApplicationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeApplicationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *ApiPrivilegeApplicationRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiPrivilegeApplicationRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiPrivilegeApplicationRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiPrivilegeApplicationRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *ApiPrivilegeApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeApplicationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPrivilegeApplicationRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


