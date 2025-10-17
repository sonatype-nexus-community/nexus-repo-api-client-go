# ConfigurationApplyResponseXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedConfigurations** | Pointer to [**[]ApplyStatusXO**](ApplyStatusXO.md) |  | [optional] 
**FailedConfigurations** | Pointer to [**[]ApplyStatusXO**](ApplyStatusXO.md) |  | [optional] 
**NotAppliedConfigurations** | Pointer to [**[]ApplyStatusXO**](ApplyStatusXO.md) |  | [optional] 

## Methods

### NewConfigurationApplyResponseXO

`func NewConfigurationApplyResponseXO() *ConfigurationApplyResponseXO`

NewConfigurationApplyResponseXO instantiates a new ConfigurationApplyResponseXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationApplyResponseXOWithDefaults

`func NewConfigurationApplyResponseXOWithDefaults() *ConfigurationApplyResponseXO`

NewConfigurationApplyResponseXOWithDefaults instantiates a new ConfigurationApplyResponseXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedConfigurations

`func (o *ConfigurationApplyResponseXO) GetAppliedConfigurations() []ApplyStatusXO`

GetAppliedConfigurations returns the AppliedConfigurations field if non-nil, zero value otherwise.

### GetAppliedConfigurationsOk

`func (o *ConfigurationApplyResponseXO) GetAppliedConfigurationsOk() (*[]ApplyStatusXO, bool)`

GetAppliedConfigurationsOk returns a tuple with the AppliedConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedConfigurations

`func (o *ConfigurationApplyResponseXO) SetAppliedConfigurations(v []ApplyStatusXO)`

SetAppliedConfigurations sets AppliedConfigurations field to given value.

### HasAppliedConfigurations

`func (o *ConfigurationApplyResponseXO) HasAppliedConfigurations() bool`

HasAppliedConfigurations returns a boolean if a field has been set.

### GetFailedConfigurations

`func (o *ConfigurationApplyResponseXO) GetFailedConfigurations() []ApplyStatusXO`

GetFailedConfigurations returns the FailedConfigurations field if non-nil, zero value otherwise.

### GetFailedConfigurationsOk

`func (o *ConfigurationApplyResponseXO) GetFailedConfigurationsOk() (*[]ApplyStatusXO, bool)`

GetFailedConfigurationsOk returns a tuple with the FailedConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedConfigurations

`func (o *ConfigurationApplyResponseXO) SetFailedConfigurations(v []ApplyStatusXO)`

SetFailedConfigurations sets FailedConfigurations field to given value.

### HasFailedConfigurations

`func (o *ConfigurationApplyResponseXO) HasFailedConfigurations() bool`

HasFailedConfigurations returns a boolean if a field has been set.

### GetNotAppliedConfigurations

`func (o *ConfigurationApplyResponseXO) GetNotAppliedConfigurations() []ApplyStatusXO`

GetNotAppliedConfigurations returns the NotAppliedConfigurations field if non-nil, zero value otherwise.

### GetNotAppliedConfigurationsOk

`func (o *ConfigurationApplyResponseXO) GetNotAppliedConfigurationsOk() (*[]ApplyStatusXO, bool)`

GetNotAppliedConfigurationsOk returns a tuple with the NotAppliedConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAppliedConfigurations

`func (o *ConfigurationApplyResponseXO) SetNotAppliedConfigurations(v []ApplyStatusXO)`

SetNotAppliedConfigurations sets NotAppliedConfigurations field to given value.

### HasNotAppliedConfigurations

`func (o *ConfigurationApplyResponseXO) HasNotAppliedConfigurations() bool`

HasNotAppliedConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


