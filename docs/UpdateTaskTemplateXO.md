# UpdateTaskTemplateXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmail** | Pointer to **string** | e-mail for task notifications. | [optional] 
**Enabled** | **bool** | Indicates if the task would be enabled. | 
**Frequency** | [**FrequencyXO**](FrequencyXO.md) |  | 
**Name** | **string** | The name of the task template. | 
**NotificationCondition** | **string** | Condition required to notify a task execution. | 
**Properties** | Pointer to **map[string]string** | Additional properties for the task | [optional] 

## Methods

### NewUpdateTaskTemplateXO

`func NewUpdateTaskTemplateXO(enabled bool, frequency FrequencyXO, name string, notificationCondition string, ) *UpdateTaskTemplateXO`

NewUpdateTaskTemplateXO instantiates a new UpdateTaskTemplateXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskTemplateXOWithDefaults

`func NewUpdateTaskTemplateXOWithDefaults() *UpdateTaskTemplateXO`

NewUpdateTaskTemplateXOWithDefaults instantiates a new UpdateTaskTemplateXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEmail

`func (o *UpdateTaskTemplateXO) GetAlertEmail() string`

GetAlertEmail returns the AlertEmail field if non-nil, zero value otherwise.

### GetAlertEmailOk

`func (o *UpdateTaskTemplateXO) GetAlertEmailOk() (*string, bool)`

GetAlertEmailOk returns a tuple with the AlertEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail

`func (o *UpdateTaskTemplateXO) SetAlertEmail(v string)`

SetAlertEmail sets AlertEmail field to given value.

### HasAlertEmail

`func (o *UpdateTaskTemplateXO) HasAlertEmail() bool`

HasAlertEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateTaskTemplateXO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateTaskTemplateXO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateTaskTemplateXO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFrequency

`func (o *UpdateTaskTemplateXO) GetFrequency() FrequencyXO`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *UpdateTaskTemplateXO) GetFrequencyOk() (*FrequencyXO, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *UpdateTaskTemplateXO) SetFrequency(v FrequencyXO)`

SetFrequency sets Frequency field to given value.


### GetName

`func (o *UpdateTaskTemplateXO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTaskTemplateXO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTaskTemplateXO) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationCondition

`func (o *UpdateTaskTemplateXO) GetNotificationCondition() string`

GetNotificationCondition returns the NotificationCondition field if non-nil, zero value otherwise.

### GetNotificationConditionOk

`func (o *UpdateTaskTemplateXO) GetNotificationConditionOk() (*string, bool)`

GetNotificationConditionOk returns a tuple with the NotificationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCondition

`func (o *UpdateTaskTemplateXO) SetNotificationCondition(v string)`

SetNotificationCondition sets NotificationCondition field to given value.


### GetProperties

`func (o *UpdateTaskTemplateXO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpdateTaskTemplateXO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpdateTaskTemplateXO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UpdateTaskTemplateXO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


