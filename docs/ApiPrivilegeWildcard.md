# ApiPrivilegeWildcard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** | The name of the privilege.  This value cannot be changed. | 
**Pattern** | **string** | A colon separated list of parts that create a permission string. | 
**ReadOnly** | Pointer to **bool** | Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system. | [optional] 
**Type** | Pointer to **string** | The type of privilege, each type covers different portions of the system. External values supplied to this will be ignored by the system. | [optional] 

## Methods

### NewApiPrivilegeWildcard

`func NewApiPrivilegeWildcard(name string, pattern string, ) *ApiPrivilegeWildcard`

NewApiPrivilegeWildcard instantiates a new ApiPrivilegeWildcard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeWildcardWithDefaults

`func NewApiPrivilegeWildcardWithDefaults() *ApiPrivilegeWildcard`

NewApiPrivilegeWildcardWithDefaults instantiates a new ApiPrivilegeWildcard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApiPrivilegeWildcard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeWildcard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeWildcard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeWildcard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ApiPrivilegeWildcard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeWildcard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeWildcard) SetName(v string)`

SetName sets Name field to given value.


### GetPattern

`func (o *ApiPrivilegeWildcard) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ApiPrivilegeWildcard) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ApiPrivilegeWildcard) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetReadOnly

`func (o *ApiPrivilegeWildcard) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApiPrivilegeWildcard) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApiPrivilegeWildcard) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApiPrivilegeWildcard) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetType

`func (o *ApiPrivilegeWildcard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPrivilegeWildcard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPrivilegeWildcard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiPrivilegeWildcard) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


