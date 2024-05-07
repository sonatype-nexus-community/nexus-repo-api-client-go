# AptHostedRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apt** | [**AptHostedRepositoriesAttributes**](AptHostedRepositoriesAttributes.md) |  | 
**AptSigning** | [**AptSigningRepositoriesAttributes**](AptSigningRepositoriesAttributes.md) |  | 
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 

## Methods

### NewAptHostedRepositoryApiRequest

`func NewAptHostedRepositoryApiRequest(apt AptHostedRepositoriesAttributes, aptSigning AptSigningRepositoriesAttributes, name string, online bool, storage HostedStorageAttributes, ) *AptHostedRepositoryApiRequest`

NewAptHostedRepositoryApiRequest instantiates a new AptHostedRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAptHostedRepositoryApiRequestWithDefaults

`func NewAptHostedRepositoryApiRequestWithDefaults() *AptHostedRepositoryApiRequest`

NewAptHostedRepositoryApiRequestWithDefaults instantiates a new AptHostedRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApt

`func (o *AptHostedRepositoryApiRequest) GetApt() AptHostedRepositoriesAttributes`

GetApt returns the Apt field if non-nil, zero value otherwise.

### GetAptOk

`func (o *AptHostedRepositoryApiRequest) GetAptOk() (*AptHostedRepositoriesAttributes, bool)`

GetAptOk returns a tuple with the Apt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApt

`func (o *AptHostedRepositoryApiRequest) SetApt(v AptHostedRepositoriesAttributes)`

SetApt sets Apt field to given value.


### GetAptSigning

`func (o *AptHostedRepositoryApiRequest) GetAptSigning() AptSigningRepositoriesAttributes`

GetAptSigning returns the AptSigning field if non-nil, zero value otherwise.

### GetAptSigningOk

`func (o *AptHostedRepositoryApiRequest) GetAptSigningOk() (*AptSigningRepositoriesAttributes, bool)`

GetAptSigningOk returns a tuple with the AptSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptSigning

`func (o *AptHostedRepositoryApiRequest) SetAptSigning(v AptSigningRepositoriesAttributes)`

SetAptSigning sets AptSigning field to given value.


### GetCleanup

`func (o *AptHostedRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *AptHostedRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *AptHostedRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *AptHostedRepositoryApiRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *AptHostedRepositoryApiRequest) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AptHostedRepositoryApiRequest) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AptHostedRepositoryApiRequest) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AptHostedRepositoryApiRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetName

`func (o *AptHostedRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AptHostedRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AptHostedRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *AptHostedRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AptHostedRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AptHostedRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *AptHostedRepositoryApiRequest) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AptHostedRepositoryApiRequest) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AptHostedRepositoryApiRequest) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


