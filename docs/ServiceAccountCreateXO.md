# ServiceAccountCreateXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ExpirationDays** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**RoleId** | **string** |  | 

## Methods

### NewServiceAccountCreateXO

`func NewServiceAccountCreateXO(name string, roleId string, ) *ServiceAccountCreateXO`

NewServiceAccountCreateXO instantiates a new ServiceAccountCreateXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountCreateXOWithDefaults

`func NewServiceAccountCreateXOWithDefaults() *ServiceAccountCreateXO`

NewServiceAccountCreateXOWithDefaults instantiates a new ServiceAccountCreateXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServiceAccountCreateXO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAccountCreateXO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAccountCreateXO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAccountCreateXO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationDays

`func (o *ServiceAccountCreateXO) GetExpirationDays() int32`

GetExpirationDays returns the ExpirationDays field if non-nil, zero value otherwise.

### GetExpirationDaysOk

`func (o *ServiceAccountCreateXO) GetExpirationDaysOk() (*int32, bool)`

GetExpirationDaysOk returns a tuple with the ExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDays

`func (o *ServiceAccountCreateXO) SetExpirationDays(v int32)`

SetExpirationDays sets ExpirationDays field to given value.

### HasExpirationDays

`func (o *ServiceAccountCreateXO) HasExpirationDays() bool`

HasExpirationDays returns a boolean if a field has been set.

### GetName

`func (o *ServiceAccountCreateXO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountCreateXO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountCreateXO) SetName(v string)`

SetName sets Name field to given value.


### GetRoleId

`func (o *ServiceAccountCreateXO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ServiceAccountCreateXO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ServiceAccountCreateXO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


