# S3BlobStoreApiAdvancedBucketConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** | A custom endpoint URL for third party object stores using the S3 API. | [optional] [readonly] 
**ForcePathStyle** | Pointer to **bool** | Setting this flag will result in path-style access being used for all requests. | [optional] [readonly] 
**MaxConnectionPoolSize** | Pointer to **int32** | Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore. | [optional] [readonly] 
**SignerType** | Pointer to **string** | An API signature version which may be required for third party object stores using the S3 API. | [optional] [readonly] 

## Methods

### NewS3BlobStoreApiAdvancedBucketConnection

`func NewS3BlobStoreApiAdvancedBucketConnection() *S3BlobStoreApiAdvancedBucketConnection`

NewS3BlobStoreApiAdvancedBucketConnection instantiates a new S3BlobStoreApiAdvancedBucketConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BlobStoreApiAdvancedBucketConnectionWithDefaults

`func NewS3BlobStoreApiAdvancedBucketConnectionWithDefaults() *S3BlobStoreApiAdvancedBucketConnection`

NewS3BlobStoreApiAdvancedBucketConnectionWithDefaults instantiates a new S3BlobStoreApiAdvancedBucketConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *S3BlobStoreApiAdvancedBucketConnection) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *S3BlobStoreApiAdvancedBucketConnection) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *S3BlobStoreApiAdvancedBucketConnection) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *S3BlobStoreApiAdvancedBucketConnection) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetForcePathStyle

`func (o *S3BlobStoreApiAdvancedBucketConnection) GetForcePathStyle() bool`

GetForcePathStyle returns the ForcePathStyle field if non-nil, zero value otherwise.

### GetForcePathStyleOk

`func (o *S3BlobStoreApiAdvancedBucketConnection) GetForcePathStyleOk() (*bool, bool)`

GetForcePathStyleOk returns a tuple with the ForcePathStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePathStyle

`func (o *S3BlobStoreApiAdvancedBucketConnection) SetForcePathStyle(v bool)`

SetForcePathStyle sets ForcePathStyle field to given value.

### HasForcePathStyle

`func (o *S3BlobStoreApiAdvancedBucketConnection) HasForcePathStyle() bool`

HasForcePathStyle returns a boolean if a field has been set.

### GetMaxConnectionPoolSize

`func (o *S3BlobStoreApiAdvancedBucketConnection) GetMaxConnectionPoolSize() int32`

GetMaxConnectionPoolSize returns the MaxConnectionPoolSize field if non-nil, zero value otherwise.

### GetMaxConnectionPoolSizeOk

`func (o *S3BlobStoreApiAdvancedBucketConnection) GetMaxConnectionPoolSizeOk() (*int32, bool)`

GetMaxConnectionPoolSizeOk returns a tuple with the MaxConnectionPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionPoolSize

`func (o *S3BlobStoreApiAdvancedBucketConnection) SetMaxConnectionPoolSize(v int32)`

SetMaxConnectionPoolSize sets MaxConnectionPoolSize field to given value.

### HasMaxConnectionPoolSize

`func (o *S3BlobStoreApiAdvancedBucketConnection) HasMaxConnectionPoolSize() bool`

HasMaxConnectionPoolSize returns a boolean if a field has been set.

### GetSignerType

`func (o *S3BlobStoreApiAdvancedBucketConnection) GetSignerType() string`

GetSignerType returns the SignerType field if non-nil, zero value otherwise.

### GetSignerTypeOk

`func (o *S3BlobStoreApiAdvancedBucketConnection) GetSignerTypeOk() (*string, bool)`

GetSignerTypeOk returns a tuple with the SignerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerType

`func (o *S3BlobStoreApiAdvancedBucketConnection) SetSignerType(v string)`

SetSignerType sets SignerType field to given value.

### HasSignerType

`func (o *S3BlobStoreApiAdvancedBucketConnection) HasSignerType() bool`

HasSignerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


