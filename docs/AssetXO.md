# AssetXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobCreated** | Pointer to **time.Time** |  | [optional] 
**Checksum** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int64** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastDownloaded** | Pointer to **time.Time** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
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

### GetChecksum

`func (o *AssetXO) GetChecksum() map[string]map[string]interface{}`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *AssetXO) GetChecksumOk() (*map[string]map[string]interface{}, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *AssetXO) SetChecksum(v map[string]map[string]interface{})`

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


