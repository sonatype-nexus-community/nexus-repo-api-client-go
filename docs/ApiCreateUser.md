# ApiCreateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | The email address associated with the user. | [optional] 
**FirstName** | Pointer to **string** | The first name of the user. | [optional] 
**LastName** | Pointer to **string** | The last name of the user. | [optional] 
**Password** | Pointer to **string** | The password for the new user. | [optional] 
**Roles** | Pointer to **[]string** | The roles which the user has been assigned within Nexus. | [optional] 
**Status** | **string** | The user&#39;s status, e.g. active or disabled. | 
**UserId** | Pointer to **string** | The userid which is required for login. This value cannot be changed. | [optional] 

## Methods

### NewApiCreateUser

`func NewApiCreateUser(status string, ) *ApiCreateUser`

NewApiCreateUser instantiates a new ApiCreateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCreateUserWithDefaults

`func NewApiCreateUserWithDefaults() *ApiCreateUser`

NewApiCreateUserWithDefaults instantiates a new ApiCreateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *ApiCreateUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ApiCreateUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ApiCreateUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ApiCreateUser) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetFirstName

`func (o *ApiCreateUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ApiCreateUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ApiCreateUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ApiCreateUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ApiCreateUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ApiCreateUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ApiCreateUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ApiCreateUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPassword

`func (o *ApiCreateUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiCreateUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiCreateUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiCreateUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoles

`func (o *ApiCreateUser) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiCreateUser) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiCreateUser) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiCreateUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *ApiCreateUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiCreateUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiCreateUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUserId

`func (o *ApiCreateUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiCreateUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiCreateUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiCreateUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


