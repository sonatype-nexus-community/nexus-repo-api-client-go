# S3BlobStoreApiEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionKey** | Pointer to **string** | The encryption key. | [optional] [readonly] 
**EncryptionType** | Pointer to **string** | The type of S3 server side encryption to use. | [optional] [readonly] 

## Methods

### NewS3BlobStoreApiEncryption

`func NewS3BlobStoreApiEncryption() *S3BlobStoreApiEncryption`

NewS3BlobStoreApiEncryption instantiates a new S3BlobStoreApiEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BlobStoreApiEncryptionWithDefaults

`func NewS3BlobStoreApiEncryptionWithDefaults() *S3BlobStoreApiEncryption`

NewS3BlobStoreApiEncryptionWithDefaults instantiates a new S3BlobStoreApiEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionKey

`func (o *S3BlobStoreApiEncryption) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *S3BlobStoreApiEncryption) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *S3BlobStoreApiEncryption) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *S3BlobStoreApiEncryption) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetEncryptionType

`func (o *S3BlobStoreApiEncryption) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *S3BlobStoreApiEncryption) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *S3BlobStoreApiEncryption) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *S3BlobStoreApiEncryption) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


