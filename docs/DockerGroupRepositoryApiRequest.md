# DockerGroupRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Docker** | [**DockerAttributes**](DockerAttributes.md) |  | 
**Group** | [**GroupDeployAttributes**](GroupDeployAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 

## Methods

### NewDockerGroupRepositoryApiRequest

`func NewDockerGroupRepositoryApiRequest(docker DockerAttributes, group GroupDeployAttributes, name string, online bool, storage StorageAttributes, ) *DockerGroupRepositoryApiRequest`

NewDockerGroupRepositoryApiRequest instantiates a new DockerGroupRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerGroupRepositoryApiRequestWithDefaults

`func NewDockerGroupRepositoryApiRequestWithDefaults() *DockerGroupRepositoryApiRequest`

NewDockerGroupRepositoryApiRequestWithDefaults instantiates a new DockerGroupRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocker

`func (o *DockerGroupRepositoryApiRequest) GetDocker() DockerAttributes`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *DockerGroupRepositoryApiRequest) GetDockerOk() (*DockerAttributes, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *DockerGroupRepositoryApiRequest) SetDocker(v DockerAttributes)`

SetDocker sets Docker field to given value.


### GetGroup

`func (o *DockerGroupRepositoryApiRequest) GetGroup() GroupDeployAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DockerGroupRepositoryApiRequest) GetGroupOk() (*GroupDeployAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DockerGroupRepositoryApiRequest) SetGroup(v GroupDeployAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *DockerGroupRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerGroupRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerGroupRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *DockerGroupRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *DockerGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *DockerGroupRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *DockerGroupRepositoryApiRequest) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DockerGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DockerGroupRepositoryApiRequest) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


