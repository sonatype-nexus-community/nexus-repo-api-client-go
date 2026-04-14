# AssetXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobCreated** | Pointer to **time.Time** |  | [optional] 
**BlobRef** | Pointer to **string** |  | [optional] 
**BlobStoreName** | Pointer to **string** |  | [optional] 
**BlobUpdated** | Pointer to **time.Time** |  | [optional] 
**Checksum** | Pointer to **map[string]string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int64** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastDownloaded** | Pointer to **time.Time** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**LastVerified** | Pointer to **time.Time** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Uploader** | Pointer to **string** |  | [optional] 
**UploaderIp** | Pointer to **string** |  | [optional] 

## Methods

### NewAssetXO

`func NewAssetXO() *AssetXO`

NewAssetXO instantiates a new AssetXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetXOWithDefaults

`func NewAssetXOWithDefaults() *AssetXO`

NewAssetXOWithDefaults instantiates a new AssetXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobCreated

`func (o *AssetXO) GetBlobCreated() time.Time`

GetBlobCreated returns the BlobCreated field if non-nil, zero value otherwise.

### GetBlobCreatedOk

`func (o *AssetXO) GetBlobCreatedOk() (*time.Time, bool)`

GetBlobCreatedOk returns a tuple with the BlobCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobCreated

`func (o *AssetXO) SetBlobCreated(v time.Time)`

SetBlobCreated sets BlobCreated field to given value.

### HasBlobCreated

`func (o *AssetXO) HasBlobCreated() bool`

HasBlobCreated returns a boolean if a field has been set.

### GetBlobRef

`func (o *AssetXO) GetBlobRef() string`

GetBlobRef returns the BlobRef field if non-nil, zero value otherwise.

### GetBlobRefOk

`func (o *AssetXO) GetBlobRefOk() (*string, bool)`

GetBlobRefOk returns a tuple with the BlobRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobRef

`func (o *AssetXO) SetBlobRef(v string)`

SetBlobRef sets BlobRef field to given value.

### HasBlobRef

`func (o *AssetXO) HasBlobRef() bool`

HasBlobRef returns a boolean if a field has been set.

### GetBlobStoreName

`func (o *AssetXO) GetBlobStoreName() string`

GetBlobStoreName returns the BlobStoreName field if non-nil, zero value otherwise.

### GetBlobStoreNameOk

`func (o *AssetXO) GetBlobStoreNameOk() (*string, bool)`

GetBlobStoreNameOk returns a tuple with the BlobStoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobStoreName

`func (o *AssetXO) SetBlobStoreName(v string)`

SetBlobStoreName sets BlobStoreName field to given value.

### HasBlobStoreName

`func (o *AssetXO) HasBlobStoreName() bool`

HasBlobStoreName returns a boolean if a field has been set.

### GetBlobUpdated

`func (o *AssetXO) GetBlobUpdated() time.Time`

GetBlobUpdated returns the BlobUpdated field if non-nil, zero value otherwise.

### GetBlobUpdatedOk

`func (o *AssetXO) GetBlobUpdatedOk() (*time.Time, bool)`

GetBlobUpdatedOk returns a tuple with the BlobUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobUpdated

`func (o *AssetXO) SetBlobUpdated(v time.Time)`

SetBlobUpdated sets BlobUpdated field to given value.

### HasBlobUpdated

`func (o *AssetXO) HasBlobUpdated() bool`

HasBlobUpdated returns a boolean if a field has been set.

### GetChecksum

`func (o *AssetXO) GetChecksum() map[string]string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *AssetXO) GetChecksumOk() (*map[string]string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *AssetXO) SetChecksum(v map[string]string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *AssetXO) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetContentType

`func (o *AssetXO) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AssetXO) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AssetXO) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AssetXO) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *AssetXO) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *AssetXO) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *AssetXO) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *AssetXO) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetFileSize

`func (o *AssetXO) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AssetXO) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AssetXO) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *AssetXO) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFormat

`func (o *AssetXO) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AssetXO) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AssetXO) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AssetXO) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetId

`func (o *AssetXO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetXO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetXO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetXO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastDownloaded

`func (o *AssetXO) GetLastDownloaded() time.Time`

GetLastDownloaded returns the LastDownloaded field if non-nil, zero value otherwise.

### GetLastDownloadedOk

`func (o *AssetXO) GetLastDownloadedOk() (*time.Time, bool)`

GetLastDownloadedOk returns a tuple with the LastDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDownloaded

`func (o *AssetXO) SetLastDownloaded(v time.Time)`

SetLastDownloaded sets LastDownloaded field to given value.

### HasLastDownloaded

`func (o *AssetXO) HasLastDownloaded() bool`

HasLastDownloaded returns a boolean if a field has been set.

### GetLastModified

`func (o *AssetXO) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AssetXO) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AssetXO) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AssetXO) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastVerified

`func (o *AssetXO) GetLastVerified() time.Time`

GetLastVerified returns the LastVerified field if non-nil, zero value otherwise.

### GetLastVerifiedOk

`func (o *AssetXO) GetLastVerifiedOk() (*time.Time, bool)`

GetLastVerifiedOk returns a tuple with the LastVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerified

`func (o *AssetXO) SetLastVerified(v time.Time)`

SetLastVerified sets LastVerified field to given value.

### HasLastVerified

`func (o *AssetXO) HasLastVerified() bool`

HasLastVerified returns a boolean if a field has been set.

### GetPath

`func (o *AssetXO) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AssetXO) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AssetXO) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AssetXO) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepository

`func (o *AssetXO) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AssetXO) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AssetXO) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AssetXO) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetUploader

`func (o *AssetXO) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *AssetXO) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *AssetXO) SetUploader(v string)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *AssetXO) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### GetUploaderIp

`func (o *AssetXO) GetUploaderIp() string`

GetUploaderIp returns the UploaderIp field if non-nil, zero value otherwise.

### GetUploaderIpOk

`func (o *AssetXO) GetUploaderIpOk() (*string, bool)`

GetUploaderIpOk returns a tuple with the UploaderIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderIp

`func (o *AssetXO) SetUploaderIp(v string)`

SetUploaderIp sets UploaderIp field to given value.

### HasUploaderIp

`func (o *AssetXO) HasUploaderIp() bool`

HasUploaderIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


