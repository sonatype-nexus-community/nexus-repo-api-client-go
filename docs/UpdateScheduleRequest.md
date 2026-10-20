# UpdateScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateScheduleRequest

`func NewUpdateScheduleRequest() *UpdateScheduleRequest`

NewUpdateScheduleRequest instantiates a new UpdateScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateScheduleRequestWithDefaults

`func NewUpdateScheduleRequestWithDefaults() *UpdateScheduleRequest`

NewUpdateScheduleRequestWithDefaults instantiates a new UpdateScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *UpdateScheduleRequest) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *UpdateScheduleRequest) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *UpdateScheduleRequest) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *UpdateScheduleRequest) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


