# UserTokensApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not User Tokens feature is enabled | [optional] 
**ProtectContent** | Pointer to **bool** | Additionally require user tokens for repository authentication | [optional] 

## Methods

### NewUserTokensApiModel

`func NewUserTokensApiModel() *UserTokensApiModel`

NewUserTokensApiModel instantiates a new UserTokensApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTokensApiModelWithDefaults

`func NewUserTokensApiModelWithDefaults() *UserTokensApiModel`

NewUserTokensApiModelWithDefaults instantiates a new UserTokensApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UserTokensApiModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserTokensApiModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserTokensApiModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserTokensApiModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProtectContent

`func (o *UserTokensApiModel) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *UserTokensApiModel) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *UserTokensApiModel) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *UserTokensApiModel) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


