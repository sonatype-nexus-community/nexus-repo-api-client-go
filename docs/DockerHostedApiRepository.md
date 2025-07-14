# DockerHostedApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Docker** | [**DockerAttributes**](DockerAttributes.md) |  | 
**Format** | Pointer to **string** |  | [optional] [default to "docker"]
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**DockerHostedStorageAttributes**](DockerHostedStorageAttributes.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "hosted"]
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewDockerHostedApiRepository

`func NewDockerHostedApiRepository(docker DockerAttributes, online bool, storage DockerHostedStorageAttributes, ) *DockerHostedApiRepository`

NewDockerHostedApiRepository instantiates a new DockerHostedApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerHostedApiRepositoryWithDefaults

`func NewDockerHostedApiRepositoryWithDefaults() *DockerHostedApiRepository`

NewDockerHostedApiRepositoryWithDefaults instantiates a new DockerHostedApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *DockerHostedApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *DockerHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *DockerHostedApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *DockerHostedApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *DockerHostedApiRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DockerHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DockerHostedApiRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *DockerHostedApiRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDocker

`func (o *DockerHostedApiRepository) GetDocker() DockerAttributes`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *DockerHostedApiRepository) GetDockerOk() (*DockerAttributes, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *DockerHostedApiRepository) SetDocker(v DockerAttributes)`

SetDocker sets Docker field to given value.


### GetFormat

`func (o *DockerHostedApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DockerHostedApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DockerHostedApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DockerHostedApiRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *DockerHostedApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerHostedApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerHostedApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DockerHostedApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *DockerHostedApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *DockerHostedApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *DockerHostedApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *DockerHostedApiRepository) GetStorage() DockerHostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DockerHostedApiRepository) GetStorageOk() (*DockerHostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DockerHostedApiRepository) SetStorage(v DockerHostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *DockerHostedApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DockerHostedApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DockerHostedApiRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DockerHostedApiRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *DockerHostedApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DockerHostedApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DockerHostedApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DockerHostedApiRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


