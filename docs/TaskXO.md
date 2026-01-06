# TaskXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentState** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastRun** | Pointer to **time.Time** |  | [optional] 
**LastRunResult** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextRun** | Pointer to **time.Time** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskXO

`func NewTaskXO() *TaskXO`

NewTaskXO instantiates a new TaskXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskXOWithDefaults

`func NewTaskXOWithDefaults() *TaskXO`

NewTaskXOWithDefaults instantiates a new TaskXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentState

`func (o *TaskXO) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *TaskXO) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *TaskXO) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *TaskXO) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetId

`func (o *TaskXO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskXO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskXO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskXO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastRun

`func (o *TaskXO) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *TaskXO) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *TaskXO) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *TaskXO) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastRunResult

`func (o *TaskXO) GetLastRunResult() string`

GetLastRunResult returns the LastRunResult field if non-nil, zero value otherwise.

### GetLastRunResultOk

`func (o *TaskXO) GetLastRunResultOk() (*string, bool)`

GetLastRunResultOk returns a tuple with the LastRunResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunResult

`func (o *TaskXO) SetLastRunResult(v string)`

SetLastRunResult sets LastRunResult field to given value.

### HasLastRunResult

`func (o *TaskXO) HasLastRunResult() bool`

HasLastRunResult returns a boolean if a field has been set.

### GetMessage

`func (o *TaskXO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskXO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskXO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TaskXO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *TaskXO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskXO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskXO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskXO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRun

`func (o *TaskXO) GetNextRun() time.Time`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *TaskXO) GetNextRunOk() (*time.Time, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *TaskXO) SetNextRun(v time.Time)`

SetNextRun sets NextRun field to given value.

### HasNextRun

`func (o *TaskXO) HasNextRun() bool`

HasNextRun returns a boolean if a field has been set.

### GetSchedule

`func (o *TaskXO) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *TaskXO) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *TaskXO) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *TaskXO) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetType

`func (o *TaskXO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskXO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskXO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskXO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


