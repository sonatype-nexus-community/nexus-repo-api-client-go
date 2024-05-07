# CleanupPolicyAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyNames** | Pointer to **[]string** | Components that match any of the applied policies will be deleted | [optional] 

## Methods

### NewCleanupPolicyAttributes

`func NewCleanupPolicyAttributes() *CleanupPolicyAttributes`

NewCleanupPolicyAttributes instantiates a new CleanupPolicyAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanupPolicyAttributesWithDefaults

`func NewCleanupPolicyAttributesWithDefaults() *CleanupPolicyAttributes`

NewCleanupPolicyAttributesWithDefaults instantiates a new CleanupPolicyAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyNames

`func (o *CleanupPolicyAttributes) GetPolicyNames() []string`

GetPolicyNames returns the PolicyNames field if non-nil, zero value otherwise.

### GetPolicyNamesOk

`func (o *CleanupPolicyAttributes) GetPolicyNamesOk() (*[]string, bool)`

GetPolicyNamesOk returns a tuple with the PolicyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNames

`func (o *CleanupPolicyAttributes) SetPolicyNames(v []string)`

SetPolicyNames sets PolicyNames field to given value.

### HasPolicyNames

`func (o *CleanupPolicyAttributes) HasPolicyNames() bool`

HasPolicyNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


