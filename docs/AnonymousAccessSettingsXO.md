# AnonymousAccessSettingsXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not Anonymous Access is enabled | [optional] 
**RealmName** | Pointer to **string** | The name of the authentication realm for the anonymous account | [optional] 
**UserId** | Pointer to **string** | The username of the anonymous account | [optional] 

## Methods

### NewAnonymousAccessSettingsXO

`func NewAnonymousAccessSettingsXO() *AnonymousAccessSettingsXO`

NewAnonymousAccessSettingsXO instantiates a new AnonymousAccessSettingsXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnonymousAccessSettingsXOWithDefaults

`func NewAnonymousAccessSettingsXOWithDefaults() *AnonymousAccessSettingsXO`

NewAnonymousAccessSettingsXOWithDefaults instantiates a new AnonymousAccessSettingsXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AnonymousAccessSettingsXO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AnonymousAccessSettingsXO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AnonymousAccessSettingsXO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AnonymousAccessSettingsXO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRealmName

`func (o *AnonymousAccessSettingsXO) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *AnonymousAccessSettingsXO) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *AnonymousAccessSettingsXO) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *AnonymousAccessSettingsXO) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetUserId

`func (o *AnonymousAccessSettingsXO) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AnonymousAccessSettingsXO) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AnonymousAccessSettingsXO) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AnonymousAccessSettingsXO) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


