# SamlUserXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | The email address of the SAML user | [optional] 
**FirstName** | Pointer to **string** | The first name of the SAML user | [optional] 
**LastName** | Pointer to **string** | The last name of the SAML user | [optional] 
**Roles** | Pointer to **[]string** | The roles assigned to the SAML user | [optional] 
**Status** | **string** | The status of the SAML user (active, disabled) | 
**UserId** | **string** | The userid for the SAML user | 

## Methods

### NewSamlUserXO

`func NewSamlUserXO(status string, userId string, ) *SamlUserXO`

NewSamlUserXO instantiates a new SamlUserXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlUserXOWithDefaults

`func NewSamlUserXOWithDefaults() *SamlUserXO`

NewSamlUserXOWithDefaults instantiates a new SamlUserXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *SamlUserXO) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SamlUserXO) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SamlUserXO) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SamlUserXO) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetFirstName

`func (o *SamlUserXO) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SamlUserXO) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SamlUserXO) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SamlUserXO) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SamlUserXO) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SamlUserXO) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SamlUserXO) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SamlUserXO) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRoles

`func (o *SamlUserXO) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SamlUserXO) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SamlUserXO) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *SamlUserXO) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *SamlUserXO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SamlUserXO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SamlUserXO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUserId

`func (o *SamlUserXO) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SamlUserXO) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SamlUserXO) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


