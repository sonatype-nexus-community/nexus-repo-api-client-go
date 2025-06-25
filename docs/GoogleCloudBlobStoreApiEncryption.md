# GoogleCloudBlobStoreApiEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionKey** | Pointer to **string** | CryptoKey ID for KMS encryption. | [optional] 
**EncryptionType** | Pointer to **string** | The type of GCP server side encryption to use. | [optional] 

## Methods

### NewGoogleCloudBlobStoreApiEncryption

`func NewGoogleCloudBlobStoreApiEncryption() *GoogleCloudBlobStoreApiEncryption`

NewGoogleCloudBlobStoreApiEncryption instantiates a new GoogleCloudBlobStoreApiEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudBlobStoreApiEncryptionWithDefaults

`func NewGoogleCloudBlobStoreApiEncryptionWithDefaults() *GoogleCloudBlobStoreApiEncryption`

NewGoogleCloudBlobStoreApiEncryptionWithDefaults instantiates a new GoogleCloudBlobStoreApiEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionKey

`func (o *GoogleCloudBlobStoreApiEncryption) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *GoogleCloudBlobStoreApiEncryption) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *GoogleCloudBlobStoreApiEncryption) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *GoogleCloudBlobStoreApiEncryption) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetEncryptionType

`func (o *GoogleCloudBlobStoreApiEncryption) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *GoogleCloudBlobStoreApiEncryption) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *GoogleCloudBlobStoreApiEncryption) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *GoogleCloudBlobStoreApiEncryption) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


