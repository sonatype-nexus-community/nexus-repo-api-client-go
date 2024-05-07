# S3BlobStoreApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketConfiguration** | [**S3BlobStoreApiBucketConfiguration**](S3BlobStoreApiBucketConfiguration.md) |  | 
**Name** | **string** | The name of the S3 blob store. | 
**SoftQuota** | Pointer to [**BlobStoreApiSoftQuota**](BlobStoreApiSoftQuota.md) |  | [optional] 
**Type** | Pointer to **string** | The blob store type. | [optional] [readonly] 

## Methods

### NewS3BlobStoreApiModel

`func NewS3BlobStoreApiModel(bucketConfiguration S3BlobStoreApiBucketConfiguration, name string, ) *S3BlobStoreApiModel`

NewS3BlobStoreApiModel instantiates a new S3BlobStoreApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BlobStoreApiModelWithDefaults

`func NewS3BlobStoreApiModelWithDefaults() *S3BlobStoreApiModel`

NewS3BlobStoreApiModelWithDefaults instantiates a new S3BlobStoreApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketConfiguration

`func (o *S3BlobStoreApiModel) GetBucketConfiguration() S3BlobStoreApiBucketConfiguration`

GetBucketConfiguration returns the BucketConfiguration field if non-nil, zero value otherwise.

### GetBucketConfigurationOk

`func (o *S3BlobStoreApiModel) GetBucketConfigurationOk() (*S3BlobStoreApiBucketConfiguration, bool)`

GetBucketConfigurationOk returns a tuple with the BucketConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketConfiguration

`func (o *S3BlobStoreApiModel) SetBucketConfiguration(v S3BlobStoreApiBucketConfiguration)`

SetBucketConfiguration sets BucketConfiguration field to given value.


### GetName

`func (o *S3BlobStoreApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3BlobStoreApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3BlobStoreApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetSoftQuota

`func (o *S3BlobStoreApiModel) GetSoftQuota() BlobStoreApiSoftQuota`

GetSoftQuota returns the SoftQuota field if non-nil, zero value otherwise.

### GetSoftQuotaOk

`func (o *S3BlobStoreApiModel) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool)`

GetSoftQuotaOk returns a tuple with the SoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftQuota

`func (o *S3BlobStoreApiModel) SetSoftQuota(v BlobStoreApiSoftQuota)`

SetSoftQuota sets SoftQuota field to given value.

### HasSoftQuota

`func (o *S3BlobStoreApiModel) HasSoftQuota() bool`

HasSoftQuota returns a boolean if a field has been set.

### GetType

`func (o *S3BlobStoreApiModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3BlobStoreApiModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3BlobStoreApiModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *S3BlobStoreApiModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


