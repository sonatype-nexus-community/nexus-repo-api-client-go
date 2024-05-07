# FileBlobStoreApiCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** | The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to or it can be a path relative to the sonatype-work directory. | [optional] 
**SoftQuota** | Pointer to [**BlobStoreApiSoftQuota**](BlobStoreApiSoftQuota.md) |  | [optional] 

## Methods

### NewFileBlobStoreApiCreateRequest

`func NewFileBlobStoreApiCreateRequest() *FileBlobStoreApiCreateRequest`

NewFileBlobStoreApiCreateRequest instantiates a new FileBlobStoreApiCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileBlobStoreApiCreateRequestWithDefaults

`func NewFileBlobStoreApiCreateRequestWithDefaults() *FileBlobStoreApiCreateRequest`

NewFileBlobStoreApiCreateRequestWithDefaults instantiates a new FileBlobStoreApiCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileBlobStoreApiCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileBlobStoreApiCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileBlobStoreApiCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileBlobStoreApiCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *FileBlobStoreApiCreateRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileBlobStoreApiCreateRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileBlobStoreApiCreateRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileBlobStoreApiCreateRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSoftQuota

`func (o *FileBlobStoreApiCreateRequest) GetSoftQuota() BlobStoreApiSoftQuota`

GetSoftQuota returns the SoftQuota field if non-nil, zero value otherwise.

### GetSoftQuotaOk

`func (o *FileBlobStoreApiCreateRequest) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool)`

GetSoftQuotaOk returns a tuple with the SoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftQuota

`func (o *FileBlobStoreApiCreateRequest) SetSoftQuota(v BlobStoreApiSoftQuota)`

SetSoftQuota sets SoftQuota field to given value.

### HasSoftQuota

`func (o *FileBlobStoreApiCreateRequest) HasSoftQuota() bool`

HasSoftQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


