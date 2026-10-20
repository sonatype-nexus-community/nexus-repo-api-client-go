# OAuth2UserXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | The email address of the OAuth2 user | [optional] 
**FirstName** | Pointer to **string** | The first name of the OAuth2 user | [optional] 
**LastName** | Pointer to **string** | The last name of the OAuth2 user | [optional] 
**Roles** | Pointer to **[]string** | The roles assigned to the OAuth2 user | [optional] 
**Status** | **string** | The status of the OAuth2 user. One of: active, locked, disabled | 
**UserId** | **string** | The userid for the OAuth2 user | 

## Methods

### NewOAuth2UserXO

`func NewOAuth2UserXO(status string, userId string, ) *OAuth2UserXO`

NewOAuth2UserXO instantiates a new OAuth2UserXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2UserXOWithDefaults

`func NewOAuth2UserXOWithDefaults() *OAuth2UserXO`

NewOAuth2UserXOWithDefaults instantiates a new OAuth2UserXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *OAuth2UserXO) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *OAuth2UserXO) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *OAuth2UserXO) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *OAuth2UserXO) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetFirstName

`func (o *OAuth2UserXO) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OAuth2UserXO) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OAuth2UserXO) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OAuth2UserXO) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *OAuth2UserXO) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OAuth2UserXO) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OAuth2UserXO) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OAuth2UserXO) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRoles

`func (o *OAuth2UserXO) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OAuth2UserXO) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OAuth2UserXO) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OAuth2UserXO) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *OAuth2UserXO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2UserXO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2UserXO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUserId

`func (o *OAuth2UserXO) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OAuth2UserXO) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OAuth2UserXO) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


