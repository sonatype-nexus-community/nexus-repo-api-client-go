# S3BlobStoreApiBucketConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedBucketConnection** | Pointer to [**S3BlobStoreApiAdvancedBucketConnection**](S3BlobStoreApiAdvancedBucketConnection.md) |  | [optional] 
**Bucket** | [**S3BlobStoreApiBucket**](S3BlobStoreApiBucket.md) |  | 
**BucketSecurity** | Pointer to [**S3BlobStoreApiBucketSecurity**](S3BlobStoreApiBucketSecurity.md) |  | [optional] 
**Encryption** | Pointer to [**S3BlobStoreApiEncryption**](S3BlobStoreApiEncryption.md) |  | [optional] 

## Methods

### NewS3BlobStoreApiBucketConfiguration

`func NewS3BlobStoreApiBucketConfiguration(bucket S3BlobStoreApiBucket, ) *S3BlobStoreApiBucketConfiguration`

NewS3BlobStoreApiBucketConfiguration instantiates a new S3BlobStoreApiBucketConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BlobStoreApiBucketConfigurationWithDefaults

`func NewS3BlobStoreApiBucketConfigurationWithDefaults() *S3BlobStoreApiBucketConfiguration`

NewS3BlobStoreApiBucketConfigurationWithDefaults instantiates a new S3BlobStoreApiBucketConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedBucketConnection

`func (o *S3BlobStoreApiBucketConfiguration) GetAdvancedBucketConnection() S3BlobStoreApiAdvancedBucketConnection`

GetAdvancedBucketConnection returns the AdvancedBucketConnection field if non-nil, zero value otherwise.

### GetAdvancedBucketConnectionOk

`func (o *S3BlobStoreApiBucketConfiguration) GetAdvancedBucketConnectionOk() (*S3BlobStoreApiAdvancedBucketConnection, bool)`

GetAdvancedBucketConnectionOk returns a tuple with the AdvancedBucketConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedBucketConnection

`func (o *S3BlobStoreApiBucketConfiguration) SetAdvancedBucketConnection(v S3BlobStoreApiAdvancedBucketConnection)`

SetAdvancedBucketConnection sets AdvancedBucketConnection field to given value.

### HasAdvancedBucketConnection

`func (o *S3BlobStoreApiBucketConfiguration) HasAdvancedBucketConnection() bool`

HasAdvancedBucketConnection returns a boolean if a field has been set.

### GetBucket

`func (o *S3BlobStoreApiBucketConfiguration) GetBucket() S3BlobStoreApiBucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3BlobStoreApiBucketConfiguration) GetBucketOk() (*S3BlobStoreApiBucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3BlobStoreApiBucketConfiguration) SetBucket(v S3BlobStoreApiBucket)`

SetBucket sets Bucket field to given value.


### GetBucketSecurity

`func (o *S3BlobStoreApiBucketConfiguration) GetBucketSecurity() S3BlobStoreApiBucketSecurity`

GetBucketSecurity returns the BucketSecurity field if non-nil, zero value otherwise.

### GetBucketSecurityOk

`func (o *S3BlobStoreApiBucketConfiguration) GetBucketSecurityOk() (*S3BlobStoreApiBucketSecurity, bool)`

GetBucketSecurityOk returns a tuple with the BucketSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSecurity

`func (o *S3BlobStoreApiBucketConfiguration) SetBucketSecurity(v S3BlobStoreApiBucketSecurity)`

SetBucketSecurity sets BucketSecurity field to given value.

### HasBucketSecurity

`func (o *S3BlobStoreApiBucketConfiguration) HasBucketSecurity() bool`

HasBucketSecurity returns a boolean if a field has been set.

### GetEncryption

`func (o *S3BlobStoreApiBucketConfiguration) GetEncryption() S3BlobStoreApiEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *S3BlobStoreApiBucketConfiguration) GetEncryptionOk() (*S3BlobStoreApiEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *S3BlobStoreApiBucketConfiguration) SetEncryption(v S3BlobStoreApiEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *S3BlobStoreApiBucketConfiguration) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


