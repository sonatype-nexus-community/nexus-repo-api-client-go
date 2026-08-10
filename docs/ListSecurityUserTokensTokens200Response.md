# ListSecurityUserTokensTokens200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationToken** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]UserTokenXO**](UserTokenXO.md) |  | [optional] 
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

### NewListSecurityUserTokensTokens200Response

`func NewListSecurityUserTokensTokens200Response(created time.Time, nameCode string, ) *ListSecurityUserTokensTokens200Response`

NewListSecurityUserTokensTokens200Response instantiates a new ListSecurityUserTokensTokens200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityUserTokensTokens200ResponseWithDefaults

`func NewListSecurityUserTokensTokens200ResponseWithDefaults() *ListSecurityUserTokensTokens200Response`

NewListSecurityUserTokensTokens200ResponseWithDefaults instantiates a new ListSecurityUserTokensTokens200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationToken

`func (o *ListSecurityUserTokensTokens200Response) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ListSecurityUserTokensTokens200Response) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ListSecurityUserTokensTokens200Response) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ListSecurityUserTokensTokens200Response) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetItems

`func (o *ListSecurityUserTokensTokens200Response) GetItems() []UserTokenXO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListSecurityUserTokensTokens200Response) GetItemsOk() (*[]UserTokenXO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListSecurityUserTokensTokens200Response) SetItems(v []UserTokenXO)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListSecurityUserTokensTokens200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCreated

`func (o *ListSecurityUserTokensTokens200Response) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListSecurityUserTokensTokens200Response) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListSecurityUserTokensTokens200Response) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedByRealm

`func (o *ListSecurityUserTokensTokens200Response) GetCreatedByRealm() string`

GetCreatedByRealm returns the CreatedByRealm field if non-nil, zero value otherwise.

### GetCreatedByRealmOk

`func (o *ListSecurityUserTokensTokens200Response) GetCreatedByRealmOk() (*string, bool)`

GetCreatedByRealmOk returns a tuple with the CreatedByRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByRealm

`func (o *ListSecurityUserTokensTokens200Response) SetCreatedByRealm(v string)`

SetCreatedByRealm sets CreatedByRealm field to given value.

### HasCreatedByRealm

`func (o *ListSecurityUserTokensTokens200Response) HasCreatedByRealm() bool`

HasCreatedByRealm returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ListSecurityUserTokensTokens200Response) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ListSecurityUserTokensTokens200Response) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ListSecurityUserTokensTokens200Response) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ListSecurityUserTokensTokens200Response) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetExpires

`func (o *ListSecurityUserTokensTokens200Response) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ListSecurityUserTokensTokens200Response) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ListSecurityUserTokensTokens200Response) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ListSecurityUserTokensTokens200Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetIsExpired

`func (o *ListSecurityUserTokensTokens200Response) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *ListSecurityUserTokensTokens200Response) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *ListSecurityUserTokensTokens200Response) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *ListSecurityUserTokensTokens200Response) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetNameCode

`func (o *ListSecurityUserTokensTokens200Response) GetNameCode() string`

GetNameCode returns the NameCode field if non-nil, zero value otherwise.

### GetNameCodeOk

`func (o *ListSecurityUserTokensTokens200Response) GetNameCodeOk() (*string, bool)`

GetNameCodeOk returns a tuple with the NameCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameCode

`func (o *ListSecurityUserTokensTokens200Response) SetNameCode(v string)`

SetNameCode sets NameCode field to given value.


### GetPassCode

`func (o *ListSecurityUserTokensTokens200Response) GetPassCode() string`

GetPassCode returns the PassCode field if non-nil, zero value otherwise.

### GetPassCodeOk

`func (o *ListSecurityUserTokensTokens200Response) GetPassCodeOk() (*string, bool)`

GetPassCodeOk returns a tuple with the PassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCode

`func (o *ListSecurityUserTokensTokens200Response) SetPassCode(v string)`

SetPassCode sets PassCode field to given value.

### HasPassCode

`func (o *ListSecurityUserTokensTokens200Response) HasPassCode() bool`

HasPassCode returns a boolean if a field has been set.

### GetRealm

`func (o *ListSecurityUserTokensTokens200Response) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *ListSecurityUserTokensTokens200Response) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *ListSecurityUserTokensTokens200Response) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *ListSecurityUserTokensTokens200Response) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetUserId

`func (o *ListSecurityUserTokensTokens200Response) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListSecurityUserTokensTokens200Response) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListSecurityUserTokensTokens200Response) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListSecurityUserTokensTokens200Response) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


