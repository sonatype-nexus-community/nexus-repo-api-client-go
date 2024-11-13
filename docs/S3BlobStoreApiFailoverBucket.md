# S3BlobStoreApiFailoverBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | The name of the bucket in the region | [readonly] 
**Region** | **string** | The region containing the bucket | [readonly] 

## Methods

### NewS3BlobStoreApiFailoverBucket

`func NewS3BlobStoreApiFailoverBucket(bucketName string, region string, ) *S3BlobStoreApiFailoverBucket`

NewS3BlobStoreApiFailoverBucket instantiates a new S3BlobStoreApiFailoverBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BlobStoreApiFailoverBucketWithDefaults

`func NewS3BlobStoreApiFailoverBucketWithDefaults() *S3BlobStoreApiFailoverBucket`

NewS3BlobStoreApiFailoverBucketWithDefaults instantiates a new S3BlobStoreApiFailoverBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *S3BlobStoreApiFailoverBucket) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3BlobStoreApiFailoverBucket) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3BlobStoreApiFailoverBucket) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetRegion

`func (o *S3BlobStoreApiFailoverBucket) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3BlobStoreApiFailoverBucket) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3BlobStoreApiFailoverBucket) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


