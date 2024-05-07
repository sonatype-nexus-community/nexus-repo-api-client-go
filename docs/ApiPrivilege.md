# ApiPrivilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the privilege.  This value cannot be changed. | [optional] 
**ReadOnly** | Pointer to **bool** | Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system. | [optional] 
**Type** | Pointer to **string** | The type of privilege, each type covers different portions of the system. External values supplied to this will be ignored by the system. | [optional] 

## Methods

### NewApiPrivilege

`func NewApiPrivilege() *ApiPrivilege`

NewApiPrivilege instantiates a new ApiPrivilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeWithDefaults

`func NewApiPrivilegeWithDefaults() *ApiPrivilege`

NewApiPrivilegeWithDefaults instantiates a new ApiPrivilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApiPrivilege) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilege) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilege) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilege) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ApiPrivilege) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilege) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilege) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPrivilege) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadOnly

`func (o *ApiPrivilege) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApiPrivilege) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApiPrivilege) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApiPrivilege) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetType

`func (o *ApiPrivilege) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPrivilege) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPrivilege) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiPrivilege) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


