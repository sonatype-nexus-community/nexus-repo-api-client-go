# CrowdApiXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | **string** |  | 
**ApplicationPassword** | **string** |  | 
**Enabled** | **bool** |  | 
**RealmActive** | **bool** |  | 
**Timeout** | Pointer to **int32** |  | [optional] 
**Url** | **string** |  | 
**UseTrustStoreForUrl** | Pointer to **bool** |  | [optional] 

## Methods

### NewCrowdApiXO

`func NewCrowdApiXO(applicationName string, applicationPassword string, enabled bool, realmActive bool, url string, ) *CrowdApiXO`

NewCrowdApiXO instantiates a new CrowdApiXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrowdApiXOWithDefaults

`func NewCrowdApiXOWithDefaults() *CrowdApiXO`

NewCrowdApiXOWithDefaults instantiates a new CrowdApiXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *CrowdApiXO) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *CrowdApiXO) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *CrowdApiXO) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetApplicationPassword

`func (o *CrowdApiXO) GetApplicationPassword() string`

GetApplicationPassword returns the ApplicationPassword field if non-nil, zero value otherwise.

### GetApplicationPasswordOk

`func (o *CrowdApiXO) GetApplicationPasswordOk() (*string, bool)`

GetApplicationPasswordOk returns a tuple with the ApplicationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPassword

`func (o *CrowdApiXO) SetApplicationPassword(v string)`

SetApplicationPassword sets ApplicationPassword field to given value.


### GetEnabled

`func (o *CrowdApiXO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CrowdApiXO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CrowdApiXO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRealmActive

`func (o *CrowdApiXO) GetRealmActive() bool`

GetRealmActive returns the RealmActive field if non-nil, zero value otherwise.

### GetRealmActiveOk

`func (o *CrowdApiXO) GetRealmActiveOk() (*bool, bool)`

GetRealmActiveOk returns a tuple with the RealmActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmActive

`func (o *CrowdApiXO) SetRealmActive(v bool)`

SetRealmActive sets RealmActive field to given value.


### GetTimeout

`func (o *CrowdApiXO) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CrowdApiXO) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CrowdApiXO) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CrowdApiXO) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *CrowdApiXO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CrowdApiXO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CrowdApiXO) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUseTrustStoreForUrl

`func (o *CrowdApiXO) GetUseTrustStoreForUrl() bool`

GetUseTrustStoreForUrl returns the UseTrustStoreForUrl field if non-nil, zero value otherwise.

### GetUseTrustStoreForUrlOk

`func (o *CrowdApiXO) GetUseTrustStoreForUrlOk() (*bool, bool)`

GetUseTrustStoreForUrlOk returns a tuple with the UseTrustStoreForUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrustStoreForUrl

`func (o *CrowdApiXO) SetUseTrustStoreForUrl(v bool)`

SetUseTrustStoreForUrl sets UseTrustStoreForUrl field to given value.

### HasUseTrustStoreForUrl

`func (o *CrowdApiXO) HasUseTrustStoreForUrl() bool`

HasUseTrustStoreForUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


