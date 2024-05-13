# StorageAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobStoreName** | **string** | Blob store used to store repository contents | 
**StrictContentTypeValidation** | **bool** | Whether to validate uploaded content&#39;s MIME type appropriate for the repository format | 
**WritePolicy** | Pointer to **string** | Controls if deployments of and updates to assets are allowed | [optional] 

## Methods

### NewStorageAttributes

`func NewStorageAttributes(blobStoreName string, strictContentTypeValidation bool, ) *StorageAttributes`

NewStorageAttributes instantiates a new StorageAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageAttributesWithDefaults

`func NewStorageAttributesWithDefaults() *StorageAttributes`

NewStorageAttributesWithDefaults instantiates a new StorageAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobStoreName

`func (o *StorageAttributes) GetBlobStoreName() string`

GetBlobStoreName returns the BlobStoreName field if non-nil, zero value otherwise.

### GetBlobStoreNameOk

`func (o *StorageAttributes) GetBlobStoreNameOk() (*string, bool)`

GetBlobStoreNameOk returns a tuple with the BlobStoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobStoreName

`func (o *StorageAttributes) SetBlobStoreName(v string)`

SetBlobStoreName sets BlobStoreName field to given value.


### GetStrictContentTypeValidation

`func (o *StorageAttributes) GetStrictContentTypeValidation() bool`

GetStrictContentTypeValidation returns the StrictContentTypeValidation field if non-nil, zero value otherwise.

### GetStrictContentTypeValidationOk

`func (o *StorageAttributes) GetStrictContentTypeValidationOk() (*bool, bool)`

GetStrictContentTypeValidationOk returns a tuple with the StrictContentTypeValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictContentTypeValidation

`func (o *StorageAttributes) SetStrictContentTypeValidation(v bool)`

SetStrictContentTypeValidation sets StrictContentTypeValidation field to given value.


### GetWritePolicy

`func (o *StorageAttributes) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *StorageAttributes) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *StorageAttributes) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.

### HasWritePolicy

`func (o *StorageAttributes) HasWritePolicy() bool`

HasWritePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


