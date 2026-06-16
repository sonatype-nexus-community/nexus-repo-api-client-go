# TaskXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmail** | Pointer to **string** | Email address to send notifications to | [optional] 
**CronExpression** | Pointer to **string** | Cron expression for advanced schedules | [optional] 
**CurrentState** | Pointer to **string** | The current state of the task | [optional] 
**Enabled** | Pointer to **bool** | Whether the task is enabled | [optional] 
**Id** | Pointer to **string** | The unique identifier of the task | [optional] 
**LastRun** | Pointer to **time.Time** | The last run start time (ISO 8601 format) | [optional] 
**LastRunResult** | Pointer to **string** | The result of the last run | [optional] 
**Message** | Pointer to **string** | A human-readable message describing the task | [optional] 
**Name** | Pointer to **string** | The name of the task | [optional] 
**NextRun** | Pointer to **time.Time** | The next scheduled run time (ISO 8601 format) | [optional] 
**NotificationCondition** | Pointer to **string** | Condition for sending notifications | [optional] 
**Properties** | Pointer to **map[string]string** | Task-specific configuration properties | [optional] 
**RecurringDays** | Pointer to **[]int32** | Days of the week/month for recurring schedules (1&#x3D;Sunday, 7&#x3D;Saturday, 999&#x3D;last day of month) | [optional] 
**Schedule** | Pointer to **string** | The schedule type | [optional] 
**StartDate** | Pointer to **time.Time** | The start date for scheduled tasks (ISO 8601 format) | [optional] 
**TimeZoneOffset** | Pointer to **string** | Time zone offset for cron schedules | [optional] 
**Type** | Pointer to **string** | The type identifier of the task | [optional] 

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

### GetAlertEmail

`func (o *TaskXO) GetAlertEmail() string`

GetAlertEmail returns the AlertEmail field if non-nil, zero value otherwise.

### GetAlertEmailOk

`func (o *TaskXO) GetAlertEmailOk() (*string, bool)`

GetAlertEmailOk returns a tuple with the AlertEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail

`func (o *TaskXO) SetAlertEmail(v string)`

SetAlertEmail sets AlertEmail field to given value.

### HasAlertEmail

`func (o *TaskXO) HasAlertEmail() bool`

HasAlertEmail returns a boolean if a field has been set.

### GetCronExpression

`func (o *TaskXO) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *TaskXO) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *TaskXO) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *TaskXO) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

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

### GetEnabled

`func (o *TaskXO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TaskXO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TaskXO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TaskXO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

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

### GetNotificationCondition

`func (o *TaskXO) GetNotificationCondition() string`

GetNotificationCondition returns the NotificationCondition field if non-nil, zero value otherwise.

### GetNotificationConditionOk

`func (o *TaskXO) GetNotificationConditionOk() (*string, bool)`

GetNotificationConditionOk returns a tuple with the NotificationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCondition

`func (o *TaskXO) SetNotificationCondition(v string)`

SetNotificationCondition sets NotificationCondition field to given value.

### HasNotificationCondition

`func (o *TaskXO) HasNotificationCondition() bool`

HasNotificationCondition returns a boolean if a field has been set.

### GetProperties

`func (o *TaskXO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TaskXO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TaskXO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TaskXO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRecurringDays

`func (o *TaskXO) GetRecurringDays() []int32`

GetRecurringDays returns the RecurringDays field if non-nil, zero value otherwise.

### GetRecurringDaysOk

`func (o *TaskXO) GetRecurringDaysOk() (*[]int32, bool)`

GetRecurringDaysOk returns a tuple with the RecurringDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDays

`func (o *TaskXO) SetRecurringDays(v []int32)`

SetRecurringDays sets RecurringDays field to given value.

### HasRecurringDays

`func (o *TaskXO) HasRecurringDays() bool`

HasRecurringDays returns a boolean if a field has been set.

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

### GetStartDate

`func (o *TaskXO) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TaskXO) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TaskXO) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TaskXO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTimeZoneOffset

`func (o *TaskXO) GetTimeZoneOffset() string`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *TaskXO) GetTimeZoneOffsetOk() (*string, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *TaskXO) SetTimeZoneOffset(v string)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.

### HasTimeZoneOffset

`func (o *TaskXO) HasTimeZoneOffset() bool`

HasTimeZoneOffset returns a boolean if a field has been set.

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


