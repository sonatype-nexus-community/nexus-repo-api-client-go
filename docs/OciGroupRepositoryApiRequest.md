# OciGroupRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cosign** | Pointer to [**OciCosignConfiguration**](OciCosignConfiguration.md) |  | [optional] 
**Group** | [**GroupDeployAttributes**](GroupDeployAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository (must be lowercase for OCI repositories) | 
**Oci** | [**OciAttributes**](OciAttributes.md) |  | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 

## Methods

### NewOciGroupRepositoryApiRequest

`func NewOciGroupRepositoryApiRequest(group GroupDeployAttributes, name string, oci OciAttributes, online bool, storage StorageAttributes, ) *OciGroupRepositoryApiRequest`

NewOciGroupRepositoryApiRequest instantiates a new OciGroupRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciGroupRepositoryApiRequestWithDefaults

`func NewOciGroupRepositoryApiRequestWithDefaults() *OciGroupRepositoryApiRequest`

NewOciGroupRepositoryApiRequestWithDefaults instantiates a new OciGroupRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCosign

`func (o *OciGroupRepositoryApiRequest) GetCosign() OciCosignConfiguration`

GetCosign returns the Cosign field if non-nil, zero value otherwise.

### GetCosignOk

`func (o *OciGroupRepositoryApiRequest) GetCosignOk() (*OciCosignConfiguration, bool)`

GetCosignOk returns a tuple with the Cosign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosign

`func (o *OciGroupRepositoryApiRequest) SetCosign(v OciCosignConfiguration)`

SetCosign sets Cosign field to given value.

### HasCosign

`func (o *OciGroupRepositoryApiRequest) HasCosign() bool`

HasCosign returns a boolean if a field has been set.

### GetGroup

`func (o *OciGroupRepositoryApiRequest) GetGroup() GroupDeployAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *OciGroupRepositoryApiRequest) GetGroupOk() (*GroupDeployAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *OciGroupRepositoryApiRequest) SetGroup(v GroupDeployAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *OciGroupRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OciGroupRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OciGroupRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOci

`func (o *OciGroupRepositoryApiRequest) GetOci() OciAttributes`

GetOci returns the Oci field if non-nil, zero value otherwise.

### GetOciOk

`func (o *OciGroupRepositoryApiRequest) GetOciOk() (*OciAttributes, bool)`

GetOciOk returns a tuple with the Oci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOci

`func (o *OciGroupRepositoryApiRequest) SetOci(v OciAttributes)`

SetOci sets Oci field to given value.


### GetOnline

`func (o *OciGroupRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *OciGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *OciGroupRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *OciGroupRepositoryApiRequest) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *OciGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *OciGroupRepositoryApiRequest) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


