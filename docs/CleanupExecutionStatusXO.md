# CleanupExecutionStatusXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentCount** | Pointer to **int64** |  | [optional] 
**ComponentsDeleted** | Pointer to **int64** |  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 
**DurationMs** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCleanupExecutionStatusXO

`func NewCleanupExecutionStatusXO() *CleanupExecutionStatusXO`

NewCleanupExecutionStatusXO instantiates a new CleanupExecutionStatusXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanupExecutionStatusXOWithDefaults

`func NewCleanupExecutionStatusXOWithDefaults() *CleanupExecutionStatusXO`

NewCleanupExecutionStatusXOWithDefaults instantiates a new CleanupExecutionStatusXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentCount

`func (o *CleanupExecutionStatusXO) GetComponentCount() int64`

GetComponentCount returns the ComponentCount field if non-nil, zero value otherwise.

### GetComponentCountOk

`func (o *CleanupExecutionStatusXO) GetComponentCountOk() (*int64, bool)`

GetComponentCountOk returns a tuple with the ComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCount

`func (o *CleanupExecutionStatusXO) SetComponentCount(v int64)`

SetComponentCount sets ComponentCount field to given value.

### HasComponentCount

`func (o *CleanupExecutionStatusXO) HasComponentCount() bool`

HasComponentCount returns a boolean if a field has been set.

### GetComponentsDeleted

`func (o *CleanupExecutionStatusXO) GetComponentsDeleted() int64`

GetComponentsDeleted returns the ComponentsDeleted field if non-nil, zero value otherwise.

### GetComponentsDeletedOk

`func (o *CleanupExecutionStatusXO) GetComponentsDeletedOk() (*int64, bool)`

GetComponentsDeletedOk returns a tuple with the ComponentsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentsDeleted

`func (o *CleanupExecutionStatusXO) SetComponentsDeleted(v int64)`

SetComponentsDeleted sets ComponentsDeleted field to given value.

### HasComponentsDeleted

`func (o *CleanupExecutionStatusXO) HasComponentsDeleted() bool`

HasComponentsDeleted returns a boolean if a field has been set.

### GetDryRun

`func (o *CleanupExecutionStatusXO) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CleanupExecutionStatusXO) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CleanupExecutionStatusXO) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CleanupExecutionStatusXO) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetDurationMs

`func (o *CleanupExecutionStatusXO) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *CleanupExecutionStatusXO) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *CleanupExecutionStatusXO) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *CleanupExecutionStatusXO) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetError

`func (o *CleanupExecutionStatusXO) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CleanupExecutionStatusXO) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CleanupExecutionStatusXO) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CleanupExecutionStatusXO) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *CleanupExecutionStatusXO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CleanupExecutionStatusXO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CleanupExecutionStatusXO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CleanupExecutionStatusXO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicy

`func (o *CleanupExecutionStatusXO) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CleanupExecutionStatusXO) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CleanupExecutionStatusXO) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CleanupExecutionStatusXO) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRepository

`func (o *CleanupExecutionStatusXO) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CleanupExecutionStatusXO) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CleanupExecutionStatusXO) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CleanupExecutionStatusXO) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetStatus

`func (o *CleanupExecutionStatusXO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CleanupExecutionStatusXO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CleanupExecutionStatusXO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CleanupExecutionStatusXO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


