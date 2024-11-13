# GoogleCloudBlobstoreApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketConfiguration** | [**GoogleCloudBlobStoreApiBucketConfiguration**](GoogleCloudBlobStoreApiBucketConfiguration.md) |  | 
**Name** | **string** | The name of the GC Storage blob store. | 
**SoftQuota** | Pointer to [**BlobStoreApiSoftQuota**](BlobStoreApiSoftQuota.md) |  | [optional] 
**Type** | Pointer to **string** | The blob store type. | [optional] [readonly] 

## Methods

### NewGoogleCloudBlobstoreApiModel

`func NewGoogleCloudBlobstoreApiModel(bucketConfiguration GoogleCloudBlobStoreApiBucketConfiguration, name string, ) *GoogleCloudBlobstoreApiModel`

NewGoogleCloudBlobstoreApiModel instantiates a new GoogleCloudBlobstoreApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudBlobstoreApiModelWithDefaults

`func NewGoogleCloudBlobstoreApiModelWithDefaults() *GoogleCloudBlobstoreApiModel`

NewGoogleCloudBlobstoreApiModelWithDefaults instantiates a new GoogleCloudBlobstoreApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketConfiguration

`func (o *GoogleCloudBlobstoreApiModel) GetBucketConfiguration() GoogleCloudBlobStoreApiBucketConfiguration`

GetBucketConfiguration returns the BucketConfiguration field if non-nil, zero value otherwise.

### GetBucketConfigurationOk

`func (o *GoogleCloudBlobstoreApiModel) GetBucketConfigurationOk() (*GoogleCloudBlobStoreApiBucketConfiguration, bool)`

GetBucketConfigurationOk returns a tuple with the BucketConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketConfiguration

`func (o *GoogleCloudBlobstoreApiModel) SetBucketConfiguration(v GoogleCloudBlobStoreApiBucketConfiguration)`

SetBucketConfiguration sets BucketConfiguration field to given value.


### GetName

`func (o *GoogleCloudBlobstoreApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleCloudBlobstoreApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleCloudBlobstoreApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetSoftQuota

`func (o *GoogleCloudBlobstoreApiModel) GetSoftQuota() BlobStoreApiSoftQuota`

GetSoftQuota returns the SoftQuota field if non-nil, zero value otherwise.

### GetSoftQuotaOk

`func (o *GoogleCloudBlobstoreApiModel) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool)`

GetSoftQuotaOk returns a tuple with the SoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftQuota

`func (o *GoogleCloudBlobstoreApiModel) SetSoftQuota(v BlobStoreApiSoftQuota)`

SetSoftQuota sets SoftQuota field to given value.

### HasSoftQuota

`func (o *GoogleCloudBlobstoreApiModel) HasSoftQuota() bool`

HasSoftQuota returns a boolean if a field has been set.

### GetType

`func (o *GoogleCloudBlobstoreApiModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GoogleCloudBlobstoreApiModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GoogleCloudBlobstoreApiModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GoogleCloudBlobstoreApiModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


