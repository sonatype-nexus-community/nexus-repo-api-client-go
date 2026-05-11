# ApiPermissionRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logical** | **string** | How this permission combines with siblings (AND or OR) | 
**Permission** | **string** | Permission string as declared on the resource | 

## Methods

### NewApiPermissionRequirement

`func NewApiPermissionRequirement(logical string, permission string, ) *ApiPermissionRequirement`

NewApiPermissionRequirement instantiates a new ApiPermissionRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPermissionRequirementWithDefaults

`func NewApiPermissionRequirementWithDefaults() *ApiPermissionRequirement`

NewApiPermissionRequirementWithDefaults instantiates a new ApiPermissionRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogical

`func (o *ApiPermissionRequirement) GetLogical() string`

GetLogical returns the Logical field if non-nil, zero value otherwise.

### GetLogicalOk

`func (o *ApiPermissionRequirement) GetLogicalOk() (*string, bool)`

GetLogicalOk returns a tuple with the Logical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogical

`func (o *ApiPermissionRequirement) SetLogical(v string)`

SetLogical sets Logical field to given value.


### GetPermission

`func (o *ApiPermissionRequirement) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ApiPermissionRequirement) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ApiPermissionRequirement) SetPermission(v string)`

SetPermission sets Permission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


