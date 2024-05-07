# BlobStoreApiSoftQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** | The limit in MB. | [optional] 
**Type** | Pointer to **string** | The type to use such as spaceRemainingQuota, or spaceUsedQuota | [optional] 

## Methods

### NewBlobStoreApiSoftQuota

`func NewBlobStoreApiSoftQuota() *BlobStoreApiSoftQuota`

NewBlobStoreApiSoftQuota instantiates a new BlobStoreApiSoftQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobStoreApiSoftQuotaWithDefaults

`func NewBlobStoreApiSoftQuotaWithDefaults() *BlobStoreApiSoftQuota`

NewBlobStoreApiSoftQuotaWithDefaults instantiates a new BlobStoreApiSoftQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *BlobStoreApiSoftQuota) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BlobStoreApiSoftQuota) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BlobStoreApiSoftQuota) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BlobStoreApiSoftQuota) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetType

`func (o *BlobStoreApiSoftQuota) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlobStoreApiSoftQuota) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlobStoreApiSoftQuota) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlobStoreApiSoftQuota) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


