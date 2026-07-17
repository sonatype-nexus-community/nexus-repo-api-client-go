# CleanupPolicyXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaAssetRegex** | Pointer to **string** |  | [optional] 
**CriteriaLastBlobUpdated** | Pointer to **int64** |  | [optional] 
**CriteriaLastDownloaded** | Pointer to **int64** |  | [optional] 
**CriteriaReleaseType** | Pointer to **string** |  | [optional] 
**Format** | **string** |  | 
**InUseCount** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Notes** | Pointer to **string** |  | [optional] 
**Repositories** | Pointer to **[]string** |  | [optional] 
**Retain** | Pointer to **int32** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 

## Methods

### NewCleanupPolicyXO

`func NewCleanupPolicyXO(format string, name string, ) *CleanupPolicyXO`

NewCleanupPolicyXO instantiates a new CleanupPolicyXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanupPolicyXOWithDefaults

`func NewCleanupPolicyXOWithDefaults() *CleanupPolicyXO`

NewCleanupPolicyXOWithDefaults instantiates a new CleanupPolicyXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaAssetRegex

`func (o *CleanupPolicyXO) GetCriteriaAssetRegex() string`

GetCriteriaAssetRegex returns the CriteriaAssetRegex field if non-nil, zero value otherwise.

### GetCriteriaAssetRegexOk

`func (o *CleanupPolicyXO) GetCriteriaAssetRegexOk() (*string, bool)`

GetCriteriaAssetRegexOk returns a tuple with the CriteriaAssetRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaAssetRegex

`func (o *CleanupPolicyXO) SetCriteriaAssetRegex(v string)`

SetCriteriaAssetRegex sets CriteriaAssetRegex field to given value.

### HasCriteriaAssetRegex

`func (o *CleanupPolicyXO) HasCriteriaAssetRegex() bool`

HasCriteriaAssetRegex returns a boolean if a field has been set.

### GetCriteriaLastBlobUpdated

`func (o *CleanupPolicyXO) GetCriteriaLastBlobUpdated() int64`

GetCriteriaLastBlobUpdated returns the CriteriaLastBlobUpdated field if non-nil, zero value otherwise.

### GetCriteriaLastBlobUpdatedOk

`func (o *CleanupPolicyXO) GetCriteriaLastBlobUpdatedOk() (*int64, bool)`

GetCriteriaLastBlobUpdatedOk returns a tuple with the CriteriaLastBlobUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaLastBlobUpdated

`func (o *CleanupPolicyXO) SetCriteriaLastBlobUpdated(v int64)`

SetCriteriaLastBlobUpdated sets CriteriaLastBlobUpdated field to given value.

### HasCriteriaLastBlobUpdated

`func (o *CleanupPolicyXO) HasCriteriaLastBlobUpdated() bool`

HasCriteriaLastBlobUpdated returns a boolean if a field has been set.

### GetCriteriaLastDownloaded

`func (o *CleanupPolicyXO) GetCriteriaLastDownloaded() int64`

GetCriteriaLastDownloaded returns the CriteriaLastDownloaded field if non-nil, zero value otherwise.

### GetCriteriaLastDownloadedOk

`func (o *CleanupPolicyXO) GetCriteriaLastDownloadedOk() (*int64, bool)`

GetCriteriaLastDownloadedOk returns a tuple with the CriteriaLastDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaLastDownloaded

`func (o *CleanupPolicyXO) SetCriteriaLastDownloaded(v int64)`

SetCriteriaLastDownloaded sets CriteriaLastDownloaded field to given value.

### HasCriteriaLastDownloaded

`func (o *CleanupPolicyXO) HasCriteriaLastDownloaded() bool`

HasCriteriaLastDownloaded returns a boolean if a field has been set.

### GetCriteriaReleaseType

`func (o *CleanupPolicyXO) GetCriteriaReleaseType() string`

GetCriteriaReleaseType returns the CriteriaReleaseType field if non-nil, zero value otherwise.

### GetCriteriaReleaseTypeOk

`func (o *CleanupPolicyXO) GetCriteriaReleaseTypeOk() (*string, bool)`

GetCriteriaReleaseTypeOk returns a tuple with the CriteriaReleaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaReleaseType

`func (o *CleanupPolicyXO) SetCriteriaReleaseType(v string)`

SetCriteriaReleaseType sets CriteriaReleaseType field to given value.

### HasCriteriaReleaseType

`func (o *CleanupPolicyXO) HasCriteriaReleaseType() bool`

HasCriteriaReleaseType returns a boolean if a field has been set.

### GetFormat

`func (o *CleanupPolicyXO) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CleanupPolicyXO) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CleanupPolicyXO) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetInUseCount

`func (o *CleanupPolicyXO) GetInUseCount() int32`

GetInUseCount returns the InUseCount field if non-nil, zero value otherwise.

### GetInUseCountOk

`func (o *CleanupPolicyXO) GetInUseCountOk() (*int32, bool)`

GetInUseCountOk returns a tuple with the InUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUseCount

`func (o *CleanupPolicyXO) SetInUseCount(v int32)`

SetInUseCount sets InUseCount field to given value.

### HasInUseCount

`func (o *CleanupPolicyXO) HasInUseCount() bool`

HasInUseCount returns a boolean if a field has been set.

### GetName

`func (o *CleanupPolicyXO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CleanupPolicyXO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CleanupPolicyXO) SetName(v string)`

SetName sets Name field to given value.


### GetNotes

`func (o *CleanupPolicyXO) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CleanupPolicyXO) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CleanupPolicyXO) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CleanupPolicyXO) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRepositories

`func (o *CleanupPolicyXO) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *CleanupPolicyXO) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *CleanupPolicyXO) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *CleanupPolicyXO) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetRetain

`func (o *CleanupPolicyXO) GetRetain() int32`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *CleanupPolicyXO) GetRetainOk() (*int32, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *CleanupPolicyXO) SetRetain(v int32)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *CleanupPolicyXO) HasRetain() bool`

HasRetain returns a boolean if a field has been set.

### GetSortBy

`func (o *CleanupPolicyXO) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *CleanupPolicyXO) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *CleanupPolicyXO) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *CleanupPolicyXO) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


