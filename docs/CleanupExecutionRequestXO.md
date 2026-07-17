# CleanupExecutionRequestXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**Repository** | **string** |  | 

## Methods

### NewCleanupExecutionRequestXO

`func NewCleanupExecutionRequestXO(repository string, ) *CleanupExecutionRequestXO`

NewCleanupExecutionRequestXO instantiates a new CleanupExecutionRequestXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanupExecutionRequestXOWithDefaults

`func NewCleanupExecutionRequestXOWithDefaults() *CleanupExecutionRequestXO`

NewCleanupExecutionRequestXOWithDefaults instantiates a new CleanupExecutionRequestXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CleanupExecutionRequestXO) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CleanupExecutionRequestXO) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CleanupExecutionRequestXO) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CleanupExecutionRequestXO) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPolicy

`func (o *CleanupExecutionRequestXO) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CleanupExecutionRequestXO) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CleanupExecutionRequestXO) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CleanupExecutionRequestXO) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRepository

`func (o *CleanupExecutionRequestXO) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CleanupExecutionRequestXO) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CleanupExecutionRequestXO) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


