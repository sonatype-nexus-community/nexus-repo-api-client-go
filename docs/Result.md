# Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to [**Throwable**](Throwable.md) |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewResult

`func NewResult() *Result`

NewResult instantiates a new Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultWithDefaults

`func NewResultWithDefaults() *Result`

NewResultWithDefaults instantiates a new Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *Result) GetDetails() map[string]map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Result) GetDetailsOk() (*map[string]map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Result) SetDetails(v map[string]map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Result) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDuration

`func (o *Result) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Result) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Result) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Result) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetError

`func (o *Result) GetError() Throwable`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Result) GetErrorOk() (*Throwable, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Result) SetError(v Throwable)`

SetError sets Error field to given value.

### HasError

`func (o *Result) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHealthy

`func (o *Result) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *Result) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *Result) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *Result) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetMessage

`func (o *Result) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Result) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Result) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Result) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTime

`func (o *Result) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Result) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Result) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *Result) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimestamp

`func (o *Result) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Result) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Result) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Result) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


