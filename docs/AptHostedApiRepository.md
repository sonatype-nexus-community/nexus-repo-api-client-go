# AptHostedApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apt** | [**AptHostedRepositoriesAttributes**](AptHostedRepositoriesAttributes.md) |  | 
**AptSigning** | [**AptSigningRepositoriesAttributes**](AptSigningRepositoriesAttributes.md) |  | 
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 

## Methods

### NewAptHostedApiRepository

`func NewAptHostedApiRepository(apt AptHostedRepositoriesAttributes, aptSigning AptSigningRepositoriesAttributes, online bool, storage HostedStorageAttributes, ) *AptHostedApiRepository`

NewAptHostedApiRepository instantiates a new AptHostedApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAptHostedApiRepositoryWithDefaults

`func NewAptHostedApiRepositoryWithDefaults() *AptHostedApiRepository`

NewAptHostedApiRepositoryWithDefaults instantiates a new AptHostedApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApt

`func (o *AptHostedApiRepository) GetApt() AptHostedRepositoriesAttributes`

GetApt returns the Apt field if non-nil, zero value otherwise.

### GetAptOk

`func (o *AptHostedApiRepository) GetAptOk() (*AptHostedRepositoriesAttributes, bool)`

GetAptOk returns a tuple with the Apt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApt

`func (o *AptHostedApiRepository) SetApt(v AptHostedRepositoriesAttributes)`

SetApt sets Apt field to given value.


### GetAptSigning

`func (o *AptHostedApiRepository) GetAptSigning() AptSigningRepositoriesAttributes`

GetAptSigning returns the AptSigning field if non-nil, zero value otherwise.

### GetAptSigningOk

`func (o *AptHostedApiRepository) GetAptSigningOk() (*AptSigningRepositoriesAttributes, bool)`

GetAptSigningOk returns a tuple with the AptSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptSigning

`func (o *AptHostedApiRepository) SetAptSigning(v AptSigningRepositoriesAttributes)`

SetAptSigning sets AptSigning field to given value.


### GetCleanup

`func (o *AptHostedApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *AptHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *AptHostedApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *AptHostedApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *AptHostedApiRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AptHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AptHostedApiRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AptHostedApiRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetName

`func (o *AptHostedApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AptHostedApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AptHostedApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AptHostedApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *AptHostedApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AptHostedApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AptHostedApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *AptHostedApiRepository) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AptHostedApiRepository) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AptHostedApiRepository) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


