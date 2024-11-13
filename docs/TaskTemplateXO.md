# TaskTemplateXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmail** | Pointer to **string** | e-mail for task notifications. | [optional] 
**Enabled** | **bool** | Indicates if the task would be enabled. | 
**Frequency** | [**FrequencyXO**](FrequencyXO.md) |  | 
**Name** | **string** | The name of the task template. | 
**NotificationCondition** | **string** | Condition required to notify a task execution. | 
**Properties** | Pointer to **map[string]string** | Additional properties for the task | [optional] 
**Type** | **string** | The type of task to be created. | 

## Methods

### NewTaskTemplateXO

`func NewTaskTemplateXO(enabled bool, frequency FrequencyXO, name string, notificationCondition string, type_ string, ) *TaskTemplateXO`

NewTaskTemplateXO instantiates a new TaskTemplateXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTemplateXOWithDefaults

`func NewTaskTemplateXOWithDefaults() *TaskTemplateXO`

NewTaskTemplateXOWithDefaults instantiates a new TaskTemplateXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEmail

`func (o *TaskTemplateXO) GetAlertEmail() string`

GetAlertEmail returns the AlertEmail field if non-nil, zero value otherwise.

### GetAlertEmailOk

`func (o *TaskTemplateXO) GetAlertEmailOk() (*string, bool)`

GetAlertEmailOk returns a tuple with the AlertEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail

`func (o *TaskTemplateXO) SetAlertEmail(v string)`

SetAlertEmail sets AlertEmail field to given value.

### HasAlertEmail

`func (o *TaskTemplateXO) HasAlertEmail() bool`

HasAlertEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *TaskTemplateXO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TaskTemplateXO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TaskTemplateXO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFrequency

`func (o *TaskTemplateXO) GetFrequency() FrequencyXO`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *TaskTemplateXO) GetFrequencyOk() (*FrequencyXO, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *TaskTemplateXO) SetFrequency(v FrequencyXO)`

SetFrequency sets Frequency field to given value.


### GetName

`func (o *TaskTemplateXO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskTemplateXO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskTemplateXO) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationCondition

`func (o *TaskTemplateXO) GetNotificationCondition() string`

GetNotificationCondition returns the NotificationCondition field if non-nil, zero value otherwise.

### GetNotificationConditionOk

`func (o *TaskTemplateXO) GetNotificationConditionOk() (*string, bool)`

GetNotificationConditionOk returns a tuple with the NotificationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCondition

`func (o *TaskTemplateXO) SetNotificationCondition(v string)`

SetNotificationCondition sets NotificationCondition field to given value.


### GetProperties

`func (o *TaskTemplateXO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TaskTemplateXO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TaskTemplateXO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TaskTemplateXO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetType

`func (o *TaskTemplateXO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskTemplateXO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskTemplateXO) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


