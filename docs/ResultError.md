# ResultError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalizedMessage** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StackTrace** | Pointer to [**[]ResultErrorStackTraceInner**](ResultErrorStackTraceInner.md) |  | [optional] 

## Methods

### NewResultError

`func NewResultError() *ResultError`

NewResultError instantiates a new ResultError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultErrorWithDefaults

`func NewResultErrorWithDefaults() *ResultError`

NewResultErrorWithDefaults instantiates a new ResultError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalizedMessage

`func (o *ResultError) GetLocalizedMessage() string`

GetLocalizedMessage returns the LocalizedMessage field if non-nil, zero value otherwise.

### GetLocalizedMessageOk

`func (o *ResultError) GetLocalizedMessageOk() (*string, bool)`

GetLocalizedMessageOk returns a tuple with the LocalizedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMessage

`func (o *ResultError) SetLocalizedMessage(v string)`

SetLocalizedMessage sets LocalizedMessage field to given value.

### HasLocalizedMessage

`func (o *ResultError) HasLocalizedMessage() bool`

HasLocalizedMessage returns a boolean if a field has been set.

### GetMessage

`func (o *ResultError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResultError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResultError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResultError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStackTrace

`func (o *ResultError) GetStackTrace() []ResultErrorStackTraceInner`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *ResultError) GetStackTraceOk() (*[]ResultErrorStackTraceInner, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *ResultError) SetStackTrace(v []ResultErrorStackTraceInner)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *ResultError) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


