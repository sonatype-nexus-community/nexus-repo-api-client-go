# AzureBlobStoreApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketConfiguration** | [**AzureBlobStoreApiBucketConfiguration**](AzureBlobStoreApiBucketConfiguration.md) |  | 
**Name** | **string** | The name of the Azure blob store. | 
**SoftQuota** | Pointer to [**BlobStoreApiSoftQuota**](BlobStoreApiSoftQuota.md) |  | [optional] 

## Methods

### NewAzureBlobStoreApiModel

`func NewAzureBlobStoreApiModel(bucketConfiguration AzureBlobStoreApiBucketConfiguration, name string, ) *AzureBlobStoreApiModel`

NewAzureBlobStoreApiModel instantiates a new AzureBlobStoreApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStoreApiModelWithDefaults

`func NewAzureBlobStoreApiModelWithDefaults() *AzureBlobStoreApiModel`

NewAzureBlobStoreApiModelWithDefaults instantiates a new AzureBlobStoreApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketConfiguration

`func (o *AzureBlobStoreApiModel) GetBucketConfiguration() AzureBlobStoreApiBucketConfiguration`

GetBucketConfiguration returns the BucketConfiguration field if non-nil, zero value otherwise.

### GetBucketConfigurationOk

`func (o *AzureBlobStoreApiModel) GetBucketConfigurationOk() (*AzureBlobStoreApiBucketConfiguration, bool)`

GetBucketConfigurationOk returns a tuple with the BucketConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketConfiguration

`func (o *AzureBlobStoreApiModel) SetBucketConfiguration(v AzureBlobStoreApiBucketConfiguration)`

SetBucketConfiguration sets BucketConfiguration field to given value.


### GetName

`func (o *AzureBlobStoreApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureBlobStoreApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureBlobStoreApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetSoftQuota

`func (o *AzureBlobStoreApiModel) GetSoftQuota() BlobStoreApiSoftQuota`

GetSoftQuota returns the SoftQuota field if non-nil, zero value otherwise.

### GetSoftQuotaOk

`func (o *AzureBlobStoreApiModel) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool)`

GetSoftQuotaOk returns a tuple with the SoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftQuota

`func (o *AzureBlobStoreApiModel) SetSoftQuota(v BlobStoreApiSoftQuota)`

SetSoftQuota sets SoftQuota field to given value.

### HasSoftQuota

`func (o *AzureBlobStoreApiModel) HasSoftQuota() bool`

HasSoftQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


