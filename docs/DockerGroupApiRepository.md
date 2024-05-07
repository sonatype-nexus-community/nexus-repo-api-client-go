# DockerGroupApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Docker** | [**DockerAttributes**](DockerAttributes.md) |  | 
**Group** | [**GroupDeployAttributes**](GroupDeployAttributes.md) |  | 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 

## Methods

### NewDockerGroupApiRepository

`func NewDockerGroupApiRepository(docker DockerAttributes, group GroupDeployAttributes, online bool, storage StorageAttributes, ) *DockerGroupApiRepository`

NewDockerGroupApiRepository instantiates a new DockerGroupApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerGroupApiRepositoryWithDefaults

`func NewDockerGroupApiRepositoryWithDefaults() *DockerGroupApiRepository`

NewDockerGroupApiRepositoryWithDefaults instantiates a new DockerGroupApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocker

`func (o *DockerGroupApiRepository) GetDocker() DockerAttributes`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *DockerGroupApiRepository) GetDockerOk() (*DockerAttributes, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *DockerGroupApiRepository) SetDocker(v DockerAttributes)`

SetDocker sets Docker field to given value.


### GetGroup

`func (o *DockerGroupApiRepository) GetGroup() GroupDeployAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DockerGroupApiRepository) GetGroupOk() (*GroupDeployAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DockerGroupApiRepository) SetGroup(v GroupDeployAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *DockerGroupApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerGroupApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerGroupApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DockerGroupApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *DockerGroupApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *DockerGroupApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *DockerGroupApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *DockerGroupApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DockerGroupApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DockerGroupApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


