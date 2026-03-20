# UserTokenXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**CreatedByRealm** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**IsExpired** | Pointer to **bool** |  | [optional] 
**NameCode** | **string** |  | 
**PassCode** | Pointer to **string** |  | [optional] 
**Realm** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserTokenXO

`func NewUserTokenXO(created time.Time, nameCode string, ) *UserTokenXO`

NewUserTokenXO instantiates a new UserTokenXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTokenXOWithDefaults

`func NewUserTokenXOWithDefaults() *UserTokenXO`

NewUserTokenXOWithDefaults instantiates a new UserTokenXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UserTokenXO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserTokenXO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserTokenXO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedByRealm

`func (o *UserTokenXO) GetCreatedByRealm() string`

GetCreatedByRealm returns the CreatedByRealm field if non-nil, zero value otherwise.

### GetCreatedByRealmOk

`func (o *UserTokenXO) GetCreatedByRealmOk() (*string, bool)`

GetCreatedByRealmOk returns a tuple with the CreatedByRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByRealm

`func (o *UserTokenXO) SetCreatedByRealm(v string)`

SetCreatedByRealm sets CreatedByRealm field to given value.

### HasCreatedByRealm

`func (o *UserTokenXO) HasCreatedByRealm() bool`

HasCreatedByRealm returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *UserTokenXO) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *UserTokenXO) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *UserTokenXO) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *UserTokenXO) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetExpires

`func (o *UserTokenXO) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *UserTokenXO) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *UserTokenXO) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *UserTokenXO) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetIsExpired

`func (o *UserTokenXO) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *UserTokenXO) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *UserTokenXO) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *UserTokenXO) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetNameCode

`func (o *UserTokenXO) GetNameCode() string`

GetNameCode returns the NameCode field if non-nil, zero value otherwise.

### GetNameCodeOk

`func (o *UserTokenXO) GetNameCodeOk() (*string, bool)`

GetNameCodeOk returns a tuple with the NameCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameCode

`func (o *UserTokenXO) SetNameCode(v string)`

SetNameCode sets NameCode field to given value.


### GetPassCode

`func (o *UserTokenXO) GetPassCode() string`

GetPassCode returns the PassCode field if non-nil, zero value otherwise.

### GetPassCodeOk

`func (o *UserTokenXO) GetPassCodeOk() (*string, bool)`

GetPassCodeOk returns a tuple with the PassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCode

`func (o *UserTokenXO) SetPassCode(v string)`

SetPassCode sets PassCode field to given value.

### HasPassCode

`func (o *UserTokenXO) HasPassCode() bool`

HasPassCode returns a boolean if a field has been set.

### GetRealm

`func (o *UserTokenXO) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *UserTokenXO) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *UserTokenXO) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *UserTokenXO) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetUserId

`func (o *UserTokenXO) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserTokenXO) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserTokenXO) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserTokenXO) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


