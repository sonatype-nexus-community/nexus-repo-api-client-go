# FileBlobStoreApiUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | An absolute path or a path relative to &lt;data-directory&gt;/blobs | [optional] 
**SoftQuota** | Pointer to [**BlobStoreApiSoftQuota**](BlobStoreApiSoftQuota.md) |  | [optional] 

## Methods

### NewFileBlobStoreApiUpdateRequest

`func NewFileBlobStoreApiUpdateRequest() *FileBlobStoreApiUpdateRequest`

NewFileBlobStoreApiUpdateRequest instantiates a new FileBlobStoreApiUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileBlobStoreApiUpdateRequestWithDefaults

`func NewFileBlobStoreApiUpdateRequestWithDefaults() *FileBlobStoreApiUpdateRequest`

NewFileBlobStoreApiUpdateRequestWithDefaults instantiates a new FileBlobStoreApiUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FileBlobStoreApiUpdateRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileBlobStoreApiUpdateRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileBlobStoreApiUpdateRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileBlobStoreApiUpdateRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSoftQuota

`func (o *FileBlobStoreApiUpdateRequest) GetSoftQuota() BlobStoreApiSoftQuota`

GetSoftQuota returns the SoftQuota field if non-nil, zero value otherwise.

### GetSoftQuotaOk

`func (o *FileBlobStoreApiUpdateRequest) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool)`

GetSoftQuotaOk returns a tuple with the SoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftQuota

`func (o *FileBlobStoreApiUpdateRequest) SetSoftQuota(v BlobStoreApiSoftQuota)`

SetSoftQuota sets SoftQuota field to given value.

### HasSoftQuota

`func (o *FileBlobStoreApiUpdateRequest) HasSoftQuota() bool`

HasSoftQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


