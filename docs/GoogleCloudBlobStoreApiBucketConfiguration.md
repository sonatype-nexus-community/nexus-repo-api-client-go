# GoogleCloudBlobStoreApiBucketConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | [**GoogleCloudBlobStoreApiBucket**](GoogleCloudBlobStoreApiBucket.md) |  | 
**BucketSecurity** | Pointer to [**GoogleCloudBlobStoreApiBucketAuthentication**](GoogleCloudBlobStoreApiBucketAuthentication.md) |  | [optional] 
**Encryption** | Pointer to [**GoogleCloudBlobStoreApiEncryption**](GoogleCloudBlobStoreApiEncryption.md) |  | [optional] 

## Methods

### NewGoogleCloudBlobStoreApiBucketConfiguration

`func NewGoogleCloudBlobStoreApiBucketConfiguration(bucket GoogleCloudBlobStoreApiBucket, ) *GoogleCloudBlobStoreApiBucketConfiguration`

NewGoogleCloudBlobStoreApiBucketConfiguration instantiates a new GoogleCloudBlobStoreApiBucketConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudBlobStoreApiBucketConfigurationWithDefaults

`func NewGoogleCloudBlobStoreApiBucketConfigurationWithDefaults() *GoogleCloudBlobStoreApiBucketConfiguration`

NewGoogleCloudBlobStoreApiBucketConfigurationWithDefaults instantiates a new GoogleCloudBlobStoreApiBucketConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) GetBucket() GoogleCloudBlobStoreApiBucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) GetBucketOk() (*GoogleCloudBlobStoreApiBucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) SetBucket(v GoogleCloudBlobStoreApiBucket)`

SetBucket sets Bucket field to given value.


### GetBucketSecurity

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) GetBucketSecurity() GoogleCloudBlobStoreApiBucketAuthentication`

GetBucketSecurity returns the BucketSecurity field if non-nil, zero value otherwise.

### GetBucketSecurityOk

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) GetBucketSecurityOk() (*GoogleCloudBlobStoreApiBucketAuthentication, bool)`

GetBucketSecurityOk returns a tuple with the BucketSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSecurity

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) SetBucketSecurity(v GoogleCloudBlobStoreApiBucketAuthentication)`

SetBucketSecurity sets BucketSecurity field to given value.

### HasBucketSecurity

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) HasBucketSecurity() bool`

HasBucketSecurity returns a boolean if a field has been set.

### GetEncryption

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) GetEncryption() GoogleCloudBlobStoreApiEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) GetEncryptionOk() (*GoogleCloudBlobStoreApiEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) SetEncryption(v GoogleCloudBlobStoreApiEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *GoogleCloudBlobStoreApiBucketConfiguration) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


