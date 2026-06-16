# AlpineHostedRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlpineSigning** | Pointer to [**AlpineSigningRepositoriesAttributes**](AlpineSigningRepositoriesAttributes.md) |  | [optional] 
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 

## Methods

### NewAlpineHostedRepositoryApiRequest

`func NewAlpineHostedRepositoryApiRequest(name string, online bool, storage HostedStorageAttributes, ) *AlpineHostedRepositoryApiRequest`

NewAlpineHostedRepositoryApiRequest instantiates a new AlpineHostedRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlpineHostedRepositoryApiRequestWithDefaults

`func NewAlpineHostedRepositoryApiRequestWithDefaults() *AlpineHostedRepositoryApiRequest`

NewAlpineHostedRepositoryApiRequestWithDefaults instantiates a new AlpineHostedRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpineSigning

`func (o *AlpineHostedRepositoryApiRequest) GetAlpineSigning() AlpineSigningRepositoriesAttributes`

GetAlpineSigning returns the AlpineSigning field if non-nil, zero value otherwise.

### GetAlpineSigningOk

`func (o *AlpineHostedRepositoryApiRequest) GetAlpineSigningOk() (*AlpineSigningRepositoriesAttributes, bool)`

GetAlpineSigningOk returns a tuple with the AlpineSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpineSigning

`func (o *AlpineHostedRepositoryApiRequest) SetAlpineSigning(v AlpineSigningRepositoriesAttributes)`

SetAlpineSigning sets AlpineSigning field to given value.

### HasAlpineSigning

`func (o *AlpineHostedRepositoryApiRequest) HasAlpineSigning() bool`

HasAlpineSigning returns a boolean if a field has been set.

### GetCleanup

`func (o *AlpineHostedRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *AlpineHostedRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *AlpineHostedRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *AlpineHostedRepositoryApiRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *AlpineHostedRepositoryApiRequest) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AlpineHostedRepositoryApiRequest) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AlpineHostedRepositoryApiRequest) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AlpineHostedRepositoryApiRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetName

`func (o *AlpineHostedRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlpineHostedRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlpineHostedRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *AlpineHostedRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AlpineHostedRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AlpineHostedRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *AlpineHostedRepositoryApiRequest) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AlpineHostedRepositoryApiRequest) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AlpineHostedRepositoryApiRequest) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


