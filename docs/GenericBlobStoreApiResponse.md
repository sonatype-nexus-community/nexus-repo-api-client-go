# GenericBlobStoreApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableSpaceInBytes** | Pointer to **int64** |  | [optional] 
**BlobCount** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SoftQuota** | Pointer to [**BlobStoreApiSoftQuota**](BlobStoreApiSoftQuota.md) |  | [optional] 
**TotalSizeInBytes** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Unavailable** | Pointer to **bool** |  | [optional] 

## Methods

### NewGenericBlobStoreApiResponse

`func NewGenericBlobStoreApiResponse() *GenericBlobStoreApiResponse`

NewGenericBlobStoreApiResponse instantiates a new GenericBlobStoreApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericBlobStoreApiResponseWithDefaults

`func NewGenericBlobStoreApiResponseWithDefaults() *GenericBlobStoreApiResponse`

NewGenericBlobStoreApiResponseWithDefaults instantiates a new GenericBlobStoreApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableSpaceInBytes

`func (o *GenericBlobStoreApiResponse) GetAvailableSpaceInBytes() int64`

GetAvailableSpaceInBytes returns the AvailableSpaceInBytes field if non-nil, zero value otherwise.

### GetAvailableSpaceInBytesOk

`func (o *GenericBlobStoreApiResponse) GetAvailableSpaceInBytesOk() (*int64, bool)`

GetAvailableSpaceInBytesOk returns a tuple with the AvailableSpaceInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSpaceInBytes

`func (o *GenericBlobStoreApiResponse) SetAvailableSpaceInBytes(v int64)`

SetAvailableSpaceInBytes sets AvailableSpaceInBytes field to given value.

### HasAvailableSpaceInBytes

`func (o *GenericBlobStoreApiResponse) HasAvailableSpaceInBytes() bool`

HasAvailableSpaceInBytes returns a boolean if a field has been set.

### GetBlobCount

`func (o *GenericBlobStoreApiResponse) GetBlobCount() int64`

GetBlobCount returns the BlobCount field if non-nil, zero value otherwise.

### GetBlobCountOk

`func (o *GenericBlobStoreApiResponse) GetBlobCountOk() (*int64, bool)`

GetBlobCountOk returns a tuple with the BlobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobCount

`func (o *GenericBlobStoreApiResponse) SetBlobCount(v int64)`

SetBlobCount sets BlobCount field to given value.

### HasBlobCount

`func (o *GenericBlobStoreApiResponse) HasBlobCount() bool`

HasBlobCount returns a boolean if a field has been set.

### GetName

`func (o *GenericBlobStoreApiResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericBlobStoreApiResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericBlobStoreApiResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenericBlobStoreApiResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSoftQuota

`func (o *GenericBlobStoreApiResponse) GetSoftQuota() BlobStoreApiSoftQuota`

GetSoftQuota returns the SoftQuota field if non-nil, zero value otherwise.

### GetSoftQuotaOk

`func (o *GenericBlobStoreApiResponse) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool)`

GetSoftQuotaOk returns a tuple with the SoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftQuota

`func (o *GenericBlobStoreApiResponse) SetSoftQuota(v BlobStoreApiSoftQuota)`

SetSoftQuota sets SoftQuota field to given value.

### HasSoftQuota

`func (o *GenericBlobStoreApiResponse) HasSoftQuota() bool`

HasSoftQuota returns a boolean if a field has been set.

### GetTotalSizeInBytes

`func (o *GenericBlobStoreApiResponse) GetTotalSizeInBytes() int64`

GetTotalSizeInBytes returns the TotalSizeInBytes field if non-nil, zero value otherwise.

### GetTotalSizeInBytesOk

`func (o *GenericBlobStoreApiResponse) GetTotalSizeInBytesOk() (*int64, bool)`

GetTotalSizeInBytesOk returns a tuple with the TotalSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeInBytes

`func (o *GenericBlobStoreApiResponse) SetTotalSizeInBytes(v int64)`

SetTotalSizeInBytes sets TotalSizeInBytes field to given value.

### HasTotalSizeInBytes

`func (o *GenericBlobStoreApiResponse) HasTotalSizeInBytes() bool`

HasTotalSizeInBytes returns a boolean if a field has been set.

### GetType

`func (o *GenericBlobStoreApiResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericBlobStoreApiResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericBlobStoreApiResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GenericBlobStoreApiResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnavailable

`func (o *GenericBlobStoreApiResponse) GetUnavailable() bool`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *GenericBlobStoreApiResponse) GetUnavailableOk() (*bool, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *GenericBlobStoreApiResponse) SetUnavailable(v bool)`

SetUnavailable sets Unavailable field to given value.

### HasUnavailable

`func (o *GenericBlobStoreApiResponse) HasUnavailable() bool`

HasUnavailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


