# OciHostedRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Cosign** | Pointer to [**OciCosignConfiguration**](OciCosignConfiguration.md) |  | [optional] 
**Name** | **string** | A unique identifier for this repository (must be lowercase for OCI repositories) | 
**Oci** | [**OciAttributes**](OciAttributes.md) |  | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**DockerHostedStorageAttributes**](DockerHostedStorageAttributes.md) |  | 

## Methods

### NewOciHostedRepositoryApiRequest

`func NewOciHostedRepositoryApiRequest(name string, oci OciAttributes, online bool, storage DockerHostedStorageAttributes, ) *OciHostedRepositoryApiRequest`

NewOciHostedRepositoryApiRequest instantiates a new OciHostedRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciHostedRepositoryApiRequestWithDefaults

`func NewOciHostedRepositoryApiRequestWithDefaults() *OciHostedRepositoryApiRequest`

NewOciHostedRepositoryApiRequestWithDefaults instantiates a new OciHostedRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *OciHostedRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *OciHostedRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *OciHostedRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *OciHostedRepositoryApiRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *OciHostedRepositoryApiRequest) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *OciHostedRepositoryApiRequest) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *OciHostedRepositoryApiRequest) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *OciHostedRepositoryApiRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCosign

`func (o *OciHostedRepositoryApiRequest) GetCosign() OciCosignConfiguration`

GetCosign returns the Cosign field if non-nil, zero value otherwise.

### GetCosignOk

`func (o *OciHostedRepositoryApiRequest) GetCosignOk() (*OciCosignConfiguration, bool)`

GetCosignOk returns a tuple with the Cosign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosign

`func (o *OciHostedRepositoryApiRequest) SetCosign(v OciCosignConfiguration)`

SetCosign sets Cosign field to given value.

### HasCosign

`func (o *OciHostedRepositoryApiRequest) HasCosign() bool`

HasCosign returns a boolean if a field has been set.

### GetName

`func (o *OciHostedRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OciHostedRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OciHostedRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOci

`func (o *OciHostedRepositoryApiRequest) GetOci() OciAttributes`

GetOci returns the Oci field if non-nil, zero value otherwise.

### GetOciOk

`func (o *OciHostedRepositoryApiRequest) GetOciOk() (*OciAttributes, bool)`

GetOciOk returns a tuple with the Oci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOci

`func (o *OciHostedRepositoryApiRequest) SetOci(v OciAttributes)`

SetOci sets Oci field to given value.


### GetOnline

`func (o *OciHostedRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *OciHostedRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *OciHostedRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *OciHostedRepositoryApiRequest) GetStorage() DockerHostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *OciHostedRepositoryApiRequest) GetStorageOk() (*DockerHostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *OciHostedRepositoryApiRequest) SetStorage(v DockerHostedStorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


