# SamlConfigurationXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAttribute** | Pointer to **string** | SAML attribute name for email | [optional] 
**EntityId** | Pointer to **string** | SAML Service Provider&#39;s unique identifying URI | [optional] 
**FirstNameAttribute** | Pointer to **string** | SAML attribute name for the first name | [optional] 
**GroupsAttribute** | Pointer to **string** | SAML attribute name for groups which maps the Identity Provider groups to a Nexus Repository Manager role | [optional] 
**IdpMetadata** | **string** | SAML Identity Provider Metadata XML | 
**LastNameAttribute** | Pointer to **string** | SAML attribute name for the last name | [optional] 
**UsernameAttribute** | **string** | SAML attribute name for the username | 
**ValidateAssertionSignature** | Pointer to **bool** | Validate signatures on assertions from Identity Provider | [optional] 
**ValidateResponseSignature** | Pointer to **bool** | Validate signatures on responses from Identity Provider | [optional] 

## Methods

### NewSamlConfigurationXO

`func NewSamlConfigurationXO(idpMetadata string, usernameAttribute string, ) *SamlConfigurationXO`

NewSamlConfigurationXO instantiates a new SamlConfigurationXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlConfigurationXOWithDefaults

`func NewSamlConfigurationXOWithDefaults() *SamlConfigurationXO`

NewSamlConfigurationXOWithDefaults instantiates a new SamlConfigurationXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAttribute

`func (o *SamlConfigurationXO) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *SamlConfigurationXO) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *SamlConfigurationXO) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *SamlConfigurationXO) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetEntityId

`func (o *SamlConfigurationXO) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SamlConfigurationXO) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SamlConfigurationXO) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SamlConfigurationXO) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetFirstNameAttribute

`func (o *SamlConfigurationXO) GetFirstNameAttribute() string`

GetFirstNameAttribute returns the FirstNameAttribute field if non-nil, zero value otherwise.

### GetFirstNameAttributeOk

`func (o *SamlConfigurationXO) GetFirstNameAttributeOk() (*string, bool)`

GetFirstNameAttributeOk returns a tuple with the FirstNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameAttribute

`func (o *SamlConfigurationXO) SetFirstNameAttribute(v string)`

SetFirstNameAttribute sets FirstNameAttribute field to given value.

### HasFirstNameAttribute

`func (o *SamlConfigurationXO) HasFirstNameAttribute() bool`

HasFirstNameAttribute returns a boolean if a field has been set.

### GetGroupsAttribute

`func (o *SamlConfigurationXO) GetGroupsAttribute() string`

GetGroupsAttribute returns the GroupsAttribute field if non-nil, zero value otherwise.

### GetGroupsAttributeOk

`func (o *SamlConfigurationXO) GetGroupsAttributeOk() (*string, bool)`

GetGroupsAttributeOk returns a tuple with the GroupsAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsAttribute

`func (o *SamlConfigurationXO) SetGroupsAttribute(v string)`

SetGroupsAttribute sets GroupsAttribute field to given value.

### HasGroupsAttribute

`func (o *SamlConfigurationXO) HasGroupsAttribute() bool`

HasGroupsAttribute returns a boolean if a field has been set.

### GetIdpMetadata

`func (o *SamlConfigurationXO) GetIdpMetadata() string`

GetIdpMetadata returns the IdpMetadata field if non-nil, zero value otherwise.

### GetIdpMetadataOk

`func (o *SamlConfigurationXO) GetIdpMetadataOk() (*string, bool)`

GetIdpMetadataOk returns a tuple with the IdpMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadata

`func (o *SamlConfigurationXO) SetIdpMetadata(v string)`

SetIdpMetadata sets IdpMetadata field to given value.


### GetLastNameAttribute

`func (o *SamlConfigurationXO) GetLastNameAttribute() string`

GetLastNameAttribute returns the LastNameAttribute field if non-nil, zero value otherwise.

### GetLastNameAttributeOk

`func (o *SamlConfigurationXO) GetLastNameAttributeOk() (*string, bool)`

GetLastNameAttributeOk returns a tuple with the LastNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameAttribute

`func (o *SamlConfigurationXO) SetLastNameAttribute(v string)`

SetLastNameAttribute sets LastNameAttribute field to given value.

### HasLastNameAttribute

`func (o *SamlConfigurationXO) HasLastNameAttribute() bool`

HasLastNameAttribute returns a boolean if a field has been set.

### GetUsernameAttribute

`func (o *SamlConfigurationXO) GetUsernameAttribute() string`

GetUsernameAttribute returns the UsernameAttribute field if non-nil, zero value otherwise.

### GetUsernameAttributeOk

`func (o *SamlConfigurationXO) GetUsernameAttributeOk() (*string, bool)`

GetUsernameAttributeOk returns a tuple with the UsernameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAttribute

`func (o *SamlConfigurationXO) SetUsernameAttribute(v string)`

SetUsernameAttribute sets UsernameAttribute field to given value.


### GetValidateAssertionSignature

`func (o *SamlConfigurationXO) GetValidateAssertionSignature() bool`

GetValidateAssertionSignature returns the ValidateAssertionSignature field if non-nil, zero value otherwise.

### GetValidateAssertionSignatureOk

`func (o *SamlConfigurationXO) GetValidateAssertionSignatureOk() (*bool, bool)`

GetValidateAssertionSignatureOk returns a tuple with the ValidateAssertionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateAssertionSignature

`func (o *SamlConfigurationXO) SetValidateAssertionSignature(v bool)`

SetValidateAssertionSignature sets ValidateAssertionSignature field to given value.

### HasValidateAssertionSignature

`func (o *SamlConfigurationXO) HasValidateAssertionSignature() bool`

HasValidateAssertionSignature returns a boolean if a field has been set.

### GetValidateResponseSignature

`func (o *SamlConfigurationXO) GetValidateResponseSignature() bool`

GetValidateResponseSignature returns the ValidateResponseSignature field if non-nil, zero value otherwise.

### GetValidateResponseSignatureOk

`func (o *SamlConfigurationXO) GetValidateResponseSignatureOk() (*bool, bool)`

GetValidateResponseSignatureOk returns a tuple with the ValidateResponseSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateResponseSignature

`func (o *SamlConfigurationXO) SetValidateResponseSignature(v bool)`

SetValidateResponseSignature sets ValidateResponseSignature field to given value.

### HasValidateResponseSignature

`func (o *SamlConfigurationXO) HasValidateResponseSignature() bool`

HasValidateResponseSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


