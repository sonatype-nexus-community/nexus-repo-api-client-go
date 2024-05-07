# DockerHostedStorageAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobStoreName** | **string** | Blob store used to store repository contents | 
**LatestPolicy** | Pointer to **bool** | Whether to allow redeploying the &#39;latest&#39; tag but defer to the Deployment Policy for all other tags | [optional] 
**StrictContentTypeValidation** | **bool** | Whether to validate uploaded content&#39;s MIME type appropriate for the repository format | 
**WritePolicy** | **string** | Controls if deployments of and updates to assets are allowed | 

## Methods

### NewDockerHostedStorageAttributes

`func NewDockerHostedStorageAttributes(blobStoreName string, strictContentTypeValidation bool, writePolicy string, ) *DockerHostedStorageAttributes`

NewDockerHostedStorageAttributes instantiates a new DockerHostedStorageAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerHostedStorageAttributesWithDefaults

`func NewDockerHostedStorageAttributesWithDefaults() *DockerHostedStorageAttributes`

NewDockerHostedStorageAttributesWithDefaults instantiates a new DockerHostedStorageAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobStoreName

`func (o *DockerHostedStorageAttributes) GetBlobStoreName() string`

GetBlobStoreName returns the BlobStoreName field if non-nil, zero value otherwise.

### GetBlobStoreNameOk

`func (o *DockerHostedStorageAttributes) GetBlobStoreNameOk() (*string, bool)`

GetBlobStoreNameOk returns a tuple with the BlobStoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobStoreName

`func (o *DockerHostedStorageAttributes) SetBlobStoreName(v string)`

SetBlobStoreName sets BlobStoreName field to given value.


### GetLatestPolicy

`func (o *DockerHostedStorageAttributes) GetLatestPolicy() bool`

GetLatestPolicy returns the LatestPolicy field if non-nil, zero value otherwise.

### GetLatestPolicyOk

`func (o *DockerHostedStorageAttributes) GetLatestPolicyOk() (*bool, bool)`

GetLatestPolicyOk returns a tuple with the LatestPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPolicy

`func (o *DockerHostedStorageAttributes) SetLatestPolicy(v bool)`

SetLatestPolicy sets LatestPolicy field to given value.

### HasLatestPolicy

`func (o *DockerHostedStorageAttributes) HasLatestPolicy() bool`

HasLatestPolicy returns a boolean if a field has been set.

### GetStrictContentTypeValidation

`func (o *DockerHostedStorageAttributes) GetStrictContentTypeValidation() bool`

GetStrictContentTypeValidation returns the StrictContentTypeValidation field if non-nil, zero value otherwise.

### GetStrictContentTypeValidationOk

`func (o *DockerHostedStorageAttributes) GetStrictContentTypeValidationOk() (*bool, bool)`

GetStrictContentTypeValidationOk returns a tuple with the StrictContentTypeValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictContentTypeValidation

`func (o *DockerHostedStorageAttributes) SetStrictContentTypeValidation(v bool)`

SetStrictContentTypeValidation sets StrictContentTypeValidation field to given value.


### GetWritePolicy

`func (o *DockerHostedStorageAttributes) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *DockerHostedStorageAttributes) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *DockerHostedStorageAttributes) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


