# ApiPrivilegeWildcardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the privilege.  This value cannot be changed. | [optional] 
**Pattern** | Pointer to **string** | A colon separated list of parts that create a permission string. | [optional] 

## Methods

### NewApiPrivilegeWildcardRequest

`func NewApiPrivilegeWildcardRequest() *ApiPrivilegeWildcardRequest`

NewApiPrivilegeWildcardRequest instantiates a new ApiPrivilegeWildcardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeWildcardRequestWithDefaults

`func NewApiPrivilegeWildcardRequestWithDefaults() *ApiPrivilegeWildcardRequest`

NewApiPrivilegeWildcardRequestWithDefaults instantiates a new ApiPrivilegeWildcardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApiPrivilegeWildcardRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeWildcardRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeWildcardRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeWildcardRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ApiPrivilegeWildcardRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeWildcardRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeWildcardRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPrivilegeWildcardRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPattern

`func (o *ApiPrivilegeWildcardRequest) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ApiPrivilegeWildcardRequest) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ApiPrivilegeWildcardRequest) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *ApiPrivilegeWildcardRequest) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


