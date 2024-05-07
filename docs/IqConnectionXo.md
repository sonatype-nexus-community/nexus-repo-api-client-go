# IqConnectionXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **string** | Authentication method | 
**Enabled** | Pointer to **bool** | Whether to use Sonatype Repository Firewall | [optional] 
**FailOpenModeEnabled** | Pointer to **bool** | Allow by default when quarantine is enabled and the IQ connection fails | [optional] 
**Password** | Pointer to **string** | Credentials for the Sonatype Repository Firewall User | [optional] 
**Properties** | Pointer to **string** | Additional properties to configure for Sonatype Repository Firewall | [optional] 
**ShowLink** | Pointer to **bool** | Show Sonatype Repository Firewall link in Browse menu when server is enabled | [optional] 
**TimeoutSeconds** | Pointer to **int32** | Seconds to wait for activity before stopping and retrying the connection. Leave blank to use the globally defined HTTP timeout. | [optional] 
**Url** | Pointer to **string** | The address of your Sonatype Repository Firewall | [optional] 
**UseTrustStoreForUrl** | Pointer to **bool** | Use certificates stored in the Nexus Repository Manager truststore to connect to Sonatype Repository Firewall | [optional] 
**Username** | Pointer to **string** | User with access to Sonatype Repository Firewall | [optional] 

## Methods

### NewIqConnectionXo

`func NewIqConnectionXo(authenticationType string, ) *IqConnectionXo`

NewIqConnectionXo instantiates a new IqConnectionXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqConnectionXoWithDefaults

`func NewIqConnectionXoWithDefaults() *IqConnectionXo`

NewIqConnectionXoWithDefaults instantiates a new IqConnectionXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *IqConnectionXo) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *IqConnectionXo) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *IqConnectionXo) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetEnabled

`func (o *IqConnectionXo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IqConnectionXo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IqConnectionXo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IqConnectionXo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFailOpenModeEnabled

`func (o *IqConnectionXo) GetFailOpenModeEnabled() bool`

GetFailOpenModeEnabled returns the FailOpenModeEnabled field if non-nil, zero value otherwise.

### GetFailOpenModeEnabledOk

`func (o *IqConnectionXo) GetFailOpenModeEnabledOk() (*bool, bool)`

GetFailOpenModeEnabledOk returns a tuple with the FailOpenModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOpenModeEnabled

`func (o *IqConnectionXo) SetFailOpenModeEnabled(v bool)`

SetFailOpenModeEnabled sets FailOpenModeEnabled field to given value.

### HasFailOpenModeEnabled

`func (o *IqConnectionXo) HasFailOpenModeEnabled() bool`

HasFailOpenModeEnabled returns a boolean if a field has been set.

### GetPassword

`func (o *IqConnectionXo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IqConnectionXo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IqConnectionXo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IqConnectionXo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProperties

`func (o *IqConnectionXo) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IqConnectionXo) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IqConnectionXo) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IqConnectionXo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetShowLink

`func (o *IqConnectionXo) GetShowLink() bool`

GetShowLink returns the ShowLink field if non-nil, zero value otherwise.

### GetShowLinkOk

`func (o *IqConnectionXo) GetShowLinkOk() (*bool, bool)`

GetShowLinkOk returns a tuple with the ShowLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLink

`func (o *IqConnectionXo) SetShowLink(v bool)`

SetShowLink sets ShowLink field to given value.

### HasShowLink

`func (o *IqConnectionXo) HasShowLink() bool`

HasShowLink returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *IqConnectionXo) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *IqConnectionXo) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *IqConnectionXo) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *IqConnectionXo) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetUrl

`func (o *IqConnectionXo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IqConnectionXo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IqConnectionXo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IqConnectionXo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUseTrustStoreForUrl

`func (o *IqConnectionXo) GetUseTrustStoreForUrl() bool`

GetUseTrustStoreForUrl returns the UseTrustStoreForUrl field if non-nil, zero value otherwise.

### GetUseTrustStoreForUrlOk

`func (o *IqConnectionXo) GetUseTrustStoreForUrlOk() (*bool, bool)`

GetUseTrustStoreForUrlOk returns a tuple with the UseTrustStoreForUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrustStoreForUrl

`func (o *IqConnectionXo) SetUseTrustStoreForUrl(v bool)`

SetUseTrustStoreForUrl sets UseTrustStoreForUrl field to given value.

### HasUseTrustStoreForUrl

`func (o *IqConnectionXo) HasUseTrustStoreForUrl() bool`

HasUseTrustStoreForUrl returns a boolean if a field has been set.

### GetUsername

`func (o *IqConnectionXo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IqConnectionXo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IqConnectionXo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IqConnectionXo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


