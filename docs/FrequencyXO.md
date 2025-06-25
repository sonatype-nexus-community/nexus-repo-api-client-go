# FrequencyXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **string** | Cron expression for the task. Only applies for for \&quot;cron\&quot; schedule. | [optional] 
**RecurringDays** | Pointer to **[]int32** | Array with the number of the days the task must run. For \&quot;weekly\&quot; schedule allowed values, 1 to 7. For \&quot;monthly\&quot; schedule allowed values, 1 to 31. | [optional] 
**Schedule** | **string** | Type of schedule (\&quot;manual\&quot;, \&quot;once\&quot;, \&quot;hourly\&quot;, \&quot;daily\&quot;, \&quot;weekly\&quot;, \&quot;monthly\&quot;, \&quot;cron\&quot;) | 
**StartDate** | Pointer to **int64** | Start date of the task represented in unix timestamp. Does not apply for \&quot;manual\&quot; schedule. | [optional] 
**TimeZoneOffset** | Pointer to **string** | The offset time zone of the client. Example:  | [optional] 

## Methods

### NewFrequencyXO

`func NewFrequencyXO(schedule string, ) *FrequencyXO`

NewFrequencyXO instantiates a new FrequencyXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrequencyXOWithDefaults

`func NewFrequencyXOWithDefaults() *FrequencyXO`

NewFrequencyXOWithDefaults instantiates a new FrequencyXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *FrequencyXO) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *FrequencyXO) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *FrequencyXO) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *FrequencyXO) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetRecurringDays

`func (o *FrequencyXO) GetRecurringDays() []int32`

GetRecurringDays returns the RecurringDays field if non-nil, zero value otherwise.

### GetRecurringDaysOk

`func (o *FrequencyXO) GetRecurringDaysOk() (*[]int32, bool)`

GetRecurringDaysOk returns a tuple with the RecurringDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDays

`func (o *FrequencyXO) SetRecurringDays(v []int32)`

SetRecurringDays sets RecurringDays field to given value.

### HasRecurringDays

`func (o *FrequencyXO) HasRecurringDays() bool`

HasRecurringDays returns a boolean if a field has been set.

### GetSchedule

`func (o *FrequencyXO) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *FrequencyXO) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *FrequencyXO) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetStartDate

`func (o *FrequencyXO) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FrequencyXO) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FrequencyXO) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FrequencyXO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTimeZoneOffset

`func (o *FrequencyXO) GetTimeZoneOffset() string`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *FrequencyXO) GetTimeZoneOffsetOk() (*string, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *FrequencyXO) SetTimeZoneOffset(v string)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.

### HasTimeZoneOffset

`func (o *FrequencyXO) HasTimeZoneOffset() bool`

HasTimeZoneOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


