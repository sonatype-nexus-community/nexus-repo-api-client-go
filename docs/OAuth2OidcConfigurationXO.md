# OAuth2OidcConfigurationXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationCustomParams** | Pointer to **map[string]string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**EmailClaim** | Pointer to **string** |  | [optional] 
**ExactMatchClaims** | Pointer to **map[string]string** |  | [optional] 
**FirstNameClaim** | Pointer to **string** |  | [optional] 
**GroupsClaim** | Pointer to **string** |  | [optional] 
**IdpAuthorizationUrl** | Pointer to **string** |  | [optional] 
**IdpJwks** | Pointer to **string** |  | [optional] 
**IdpJwksUrl** | Pointer to **string** |  | [optional] 
**IdpJwsAlgorithm** | **string** |  | 
**IdpLogoutUrl** | Pointer to **string** |  | [optional] 
**IdpTokenUrl** | Pointer to **string** |  | [optional] 
**LastNameClaim** | Pointer to **string** |  | [optional] 
**TokenRequestCustomParams** | Pointer to **map[string]string** |  | [optional] 
**UseTrustStore** | Pointer to **bool** |  | [optional] 
**UsernameClaim** | **string** |  | 

## Methods

### NewOAuth2OidcConfigurationXO

`func NewOAuth2OidcConfigurationXO(idpJwsAlgorithm string, usernameClaim string, ) *OAuth2OidcConfigurationXO`

NewOAuth2OidcConfigurationXO instantiates a new OAuth2OidcConfigurationXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2OidcConfigurationXOWithDefaults

`func NewOAuth2OidcConfigurationXOWithDefaults() *OAuth2OidcConfigurationXO`

NewOAuth2OidcConfigurationXOWithDefaults instantiates a new OAuth2OidcConfigurationXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationCustomParams

`func (o *OAuth2OidcConfigurationXO) GetAuthorizationCustomParams() map[string]string`

GetAuthorizationCustomParams returns the AuthorizationCustomParams field if non-nil, zero value otherwise.

### GetAuthorizationCustomParamsOk

`func (o *OAuth2OidcConfigurationXO) GetAuthorizationCustomParamsOk() (*map[string]string, bool)`

GetAuthorizationCustomParamsOk returns a tuple with the AuthorizationCustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCustomParams

`func (o *OAuth2OidcConfigurationXO) SetAuthorizationCustomParams(v map[string]string)`

SetAuthorizationCustomParams sets AuthorizationCustomParams field to given value.

### HasAuthorizationCustomParams

`func (o *OAuth2OidcConfigurationXO) HasAuthorizationCustomParams() bool`

HasAuthorizationCustomParams returns a boolean if a field has been set.

### GetClientId

`func (o *OAuth2OidcConfigurationXO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2OidcConfigurationXO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2OidcConfigurationXO) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuth2OidcConfigurationXO) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuth2OidcConfigurationXO) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2OidcConfigurationXO) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2OidcConfigurationXO) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuth2OidcConfigurationXO) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetEmailClaim

`func (o *OAuth2OidcConfigurationXO) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *OAuth2OidcConfigurationXO) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *OAuth2OidcConfigurationXO) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.

### HasEmailClaim

`func (o *OAuth2OidcConfigurationXO) HasEmailClaim() bool`

HasEmailClaim returns a boolean if a field has been set.

### GetExactMatchClaims

`func (o *OAuth2OidcConfigurationXO) GetExactMatchClaims() map[string]string`

GetExactMatchClaims returns the ExactMatchClaims field if non-nil, zero value otherwise.

### GetExactMatchClaimsOk

`func (o *OAuth2OidcConfigurationXO) GetExactMatchClaimsOk() (*map[string]string, bool)`

GetExactMatchClaimsOk returns a tuple with the ExactMatchClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatchClaims

`func (o *OAuth2OidcConfigurationXO) SetExactMatchClaims(v map[string]string)`

SetExactMatchClaims sets ExactMatchClaims field to given value.

### HasExactMatchClaims

`func (o *OAuth2OidcConfigurationXO) HasExactMatchClaims() bool`

HasExactMatchClaims returns a boolean if a field has been set.

### GetFirstNameClaim

`func (o *OAuth2OidcConfigurationXO) GetFirstNameClaim() string`

GetFirstNameClaim returns the FirstNameClaim field if non-nil, zero value otherwise.

### GetFirstNameClaimOk

`func (o *OAuth2OidcConfigurationXO) GetFirstNameClaimOk() (*string, bool)`

GetFirstNameClaimOk returns a tuple with the FirstNameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameClaim

`func (o *OAuth2OidcConfigurationXO) SetFirstNameClaim(v string)`

SetFirstNameClaim sets FirstNameClaim field to given value.

### HasFirstNameClaim

`func (o *OAuth2OidcConfigurationXO) HasFirstNameClaim() bool`

HasFirstNameClaim returns a boolean if a field has been set.

### GetGroupsClaim

`func (o *OAuth2OidcConfigurationXO) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *OAuth2OidcConfigurationXO) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *OAuth2OidcConfigurationXO) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *OAuth2OidcConfigurationXO) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.

### GetIdpAuthorizationUrl

`func (o *OAuth2OidcConfigurationXO) GetIdpAuthorizationUrl() string`

GetIdpAuthorizationUrl returns the IdpAuthorizationUrl field if non-nil, zero value otherwise.

### GetIdpAuthorizationUrlOk

`func (o *OAuth2OidcConfigurationXO) GetIdpAuthorizationUrlOk() (*string, bool)`

GetIdpAuthorizationUrlOk returns a tuple with the IdpAuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpAuthorizationUrl

`func (o *OAuth2OidcConfigurationXO) SetIdpAuthorizationUrl(v string)`

SetIdpAuthorizationUrl sets IdpAuthorizationUrl field to given value.

### HasIdpAuthorizationUrl

`func (o *OAuth2OidcConfigurationXO) HasIdpAuthorizationUrl() bool`

HasIdpAuthorizationUrl returns a boolean if a field has been set.

### GetIdpJwks

`func (o *OAuth2OidcConfigurationXO) GetIdpJwks() string`

GetIdpJwks returns the IdpJwks field if non-nil, zero value otherwise.

### GetIdpJwksOk

`func (o *OAuth2OidcConfigurationXO) GetIdpJwksOk() (*string, bool)`

GetIdpJwksOk returns a tuple with the IdpJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpJwks

`func (o *OAuth2OidcConfigurationXO) SetIdpJwks(v string)`

SetIdpJwks sets IdpJwks field to given value.

### HasIdpJwks

`func (o *OAuth2OidcConfigurationXO) HasIdpJwks() bool`

HasIdpJwks returns a boolean if a field has been set.

### GetIdpJwksUrl

`func (o *OAuth2OidcConfigurationXO) GetIdpJwksUrl() string`

GetIdpJwksUrl returns the IdpJwksUrl field if non-nil, zero value otherwise.

### GetIdpJwksUrlOk

`func (o *OAuth2OidcConfigurationXO) GetIdpJwksUrlOk() (*string, bool)`

GetIdpJwksUrlOk returns a tuple with the IdpJwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpJwksUrl

`func (o *OAuth2OidcConfigurationXO) SetIdpJwksUrl(v string)`

SetIdpJwksUrl sets IdpJwksUrl field to given value.

### HasIdpJwksUrl

`func (o *OAuth2OidcConfigurationXO) HasIdpJwksUrl() bool`

HasIdpJwksUrl returns a boolean if a field has been set.

### GetIdpJwsAlgorithm

`func (o *OAuth2OidcConfigurationXO) GetIdpJwsAlgorithm() string`

GetIdpJwsAlgorithm returns the IdpJwsAlgorithm field if non-nil, zero value otherwise.

### GetIdpJwsAlgorithmOk

`func (o *OAuth2OidcConfigurationXO) GetIdpJwsAlgorithmOk() (*string, bool)`

GetIdpJwsAlgorithmOk returns a tuple with the IdpJwsAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpJwsAlgorithm

`func (o *OAuth2OidcConfigurationXO) SetIdpJwsAlgorithm(v string)`

SetIdpJwsAlgorithm sets IdpJwsAlgorithm field to given value.


### GetIdpLogoutUrl

`func (o *OAuth2OidcConfigurationXO) GetIdpLogoutUrl() string`

GetIdpLogoutUrl returns the IdpLogoutUrl field if non-nil, zero value otherwise.

### GetIdpLogoutUrlOk

`func (o *OAuth2OidcConfigurationXO) GetIdpLogoutUrlOk() (*string, bool)`

GetIdpLogoutUrlOk returns a tuple with the IdpLogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpLogoutUrl

`func (o *OAuth2OidcConfigurationXO) SetIdpLogoutUrl(v string)`

SetIdpLogoutUrl sets IdpLogoutUrl field to given value.

### HasIdpLogoutUrl

`func (o *OAuth2OidcConfigurationXO) HasIdpLogoutUrl() bool`

HasIdpLogoutUrl returns a boolean if a field has been set.

### GetIdpTokenUrl

`func (o *OAuth2OidcConfigurationXO) GetIdpTokenUrl() string`

GetIdpTokenUrl returns the IdpTokenUrl field if non-nil, zero value otherwise.

### GetIdpTokenUrlOk

`func (o *OAuth2OidcConfigurationXO) GetIdpTokenUrlOk() (*string, bool)`

GetIdpTokenUrlOk returns a tuple with the IdpTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpTokenUrl

`func (o *OAuth2OidcConfigurationXO) SetIdpTokenUrl(v string)`

SetIdpTokenUrl sets IdpTokenUrl field to given value.

### HasIdpTokenUrl

`func (o *OAuth2OidcConfigurationXO) HasIdpTokenUrl() bool`

HasIdpTokenUrl returns a boolean if a field has been set.

### GetLastNameClaim

`func (o *OAuth2OidcConfigurationXO) GetLastNameClaim() string`

GetLastNameClaim returns the LastNameClaim field if non-nil, zero value otherwise.

### GetLastNameClaimOk

`func (o *OAuth2OidcConfigurationXO) GetLastNameClaimOk() (*string, bool)`

GetLastNameClaimOk returns a tuple with the LastNameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameClaim

`func (o *OAuth2OidcConfigurationXO) SetLastNameClaim(v string)`

SetLastNameClaim sets LastNameClaim field to given value.

### HasLastNameClaim

`func (o *OAuth2OidcConfigurationXO) HasLastNameClaim() bool`

HasLastNameClaim returns a boolean if a field has been set.

### GetTokenRequestCustomParams

`func (o *OAuth2OidcConfigurationXO) GetTokenRequestCustomParams() map[string]string`

GetTokenRequestCustomParams returns the TokenRequestCustomParams field if non-nil, zero value otherwise.

### GetTokenRequestCustomParamsOk

`func (o *OAuth2OidcConfigurationXO) GetTokenRequestCustomParamsOk() (*map[string]string, bool)`

GetTokenRequestCustomParamsOk returns a tuple with the TokenRequestCustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestCustomParams

`func (o *OAuth2OidcConfigurationXO) SetTokenRequestCustomParams(v map[string]string)`

SetTokenRequestCustomParams sets TokenRequestCustomParams field to given value.

### HasTokenRequestCustomParams

`func (o *OAuth2OidcConfigurationXO) HasTokenRequestCustomParams() bool`

HasTokenRequestCustomParams returns a boolean if a field has been set.

### GetUseTrustStore

`func (o *OAuth2OidcConfigurationXO) GetUseTrustStore() bool`

GetUseTrustStore returns the UseTrustStore field if non-nil, zero value otherwise.

### GetUseTrustStoreOk

`func (o *OAuth2OidcConfigurationXO) GetUseTrustStoreOk() (*bool, bool)`

GetUseTrustStoreOk returns a tuple with the UseTrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrustStore

`func (o *OAuth2OidcConfigurationXO) SetUseTrustStore(v bool)`

SetUseTrustStore sets UseTrustStore field to given value.

### HasUseTrustStore

`func (o *OAuth2OidcConfigurationXO) HasUseTrustStore() bool`

HasUseTrustStore returns a boolean if a field has been set.

### GetUsernameClaim

`func (o *OAuth2OidcConfigurationXO) GetUsernameClaim() string`

GetUsernameClaim returns the UsernameClaim field if non-nil, zero value otherwise.

### GetUsernameClaimOk

`func (o *OAuth2OidcConfigurationXO) GetUsernameClaimOk() (*string, bool)`

GetUsernameClaimOk returns a tuple with the UsernameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameClaim

`func (o *OAuth2OidcConfigurationXO) SetUsernameClaim(v string)`

SetUsernameClaim sets UsernameClaim field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


