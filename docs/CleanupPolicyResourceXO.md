# CleanupPolicyResourceXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaAssetRegex** | Pointer to **string** | asset name matcher (Remove components that have at least one asset name matching the following regular expression pattern:) | [optional] 
**CriteriaLastBlobUpdated** | Pointer to **int64** | component age (Components published over “x” days ago (e.g 1-999)) | [optional] 
**CriteriaLastDownloaded** | Pointer to **int64** | component usage (Components downloaded in “x” amount of days (e.g 1-999)) | [optional] 
**CriteriaReleaseType** | Pointer to **string** | release type (Remove components that are of the following release type:) | [optional] 
**Format** | **string** | repository format | 
**Name** | **string** | policy name | 
**Notes** | Pointer to **string** | description | [optional] 
**Retain** | Pointer to **int32** | keep the latest \&quot;x\&quot; number of versions | [optional] 

## Methods

### NewCleanupPolicyResourceXO

`func NewCleanupPolicyResourceXO(format string, name string, ) *CleanupPolicyResourceXO`

NewCleanupPolicyResourceXO instantiates a new CleanupPolicyResourceXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanupPolicyResourceXOWithDefaults

`func NewCleanupPolicyResourceXOWithDefaults() *CleanupPolicyResourceXO`

NewCleanupPolicyResourceXOWithDefaults instantiates a new CleanupPolicyResourceXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaAssetRegex

`func (o *CleanupPolicyResourceXO) GetCriteriaAssetRegex() string`

GetCriteriaAssetRegex returns the CriteriaAssetRegex field if non-nil, zero value otherwise.

### GetCriteriaAssetRegexOk

`func (o *CleanupPolicyResourceXO) GetCriteriaAssetRegexOk() (*string, bool)`

GetCriteriaAssetRegexOk returns a tuple with the CriteriaAssetRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaAssetRegex

`func (o *CleanupPolicyResourceXO) SetCriteriaAssetRegex(v string)`

SetCriteriaAssetRegex sets CriteriaAssetRegex field to given value.

### HasCriteriaAssetRegex

`func (o *CleanupPolicyResourceXO) HasCriteriaAssetRegex() bool`

HasCriteriaAssetRegex returns a boolean if a field has been set.

### GetCriteriaLastBlobUpdated

`func (o *CleanupPolicyResourceXO) GetCriteriaLastBlobUpdated() int64`

GetCriteriaLastBlobUpdated returns the CriteriaLastBlobUpdated field if non-nil, zero value otherwise.

### GetCriteriaLastBlobUpdatedOk

`func (o *CleanupPolicyResourceXO) GetCriteriaLastBlobUpdatedOk() (*int64, bool)`

GetCriteriaLastBlobUpdatedOk returns a tuple with the CriteriaLastBlobUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaLastBlobUpdated

`func (o *CleanupPolicyResourceXO) SetCriteriaLastBlobUpdated(v int64)`

SetCriteriaLastBlobUpdated sets CriteriaLastBlobUpdated field to given value.

### HasCriteriaLastBlobUpdated

`func (o *CleanupPolicyResourceXO) HasCriteriaLastBlobUpdated() bool`

HasCriteriaLastBlobUpdated returns a boolean if a field has been set.

### GetCriteriaLastDownloaded

`func (o *CleanupPolicyResourceXO) GetCriteriaLastDownloaded() int64`

GetCriteriaLastDownloaded returns the CriteriaLastDownloaded field if non-nil, zero value otherwise.

### GetCriteriaLastDownloadedOk

`func (o *CleanupPolicyResourceXO) GetCriteriaLastDownloadedOk() (*int64, bool)`

GetCriteriaLastDownloadedOk returns a tuple with the CriteriaLastDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaLastDownloaded

`func (o *CleanupPolicyResourceXO) SetCriteriaLastDownloaded(v int64)`

SetCriteriaLastDownloaded sets CriteriaLastDownloaded field to given value.

### HasCriteriaLastDownloaded

`func (o *CleanupPolicyResourceXO) HasCriteriaLastDownloaded() bool`

HasCriteriaLastDownloaded returns a boolean if a field has been set.

### GetCriteriaReleaseType

`func (o *CleanupPolicyResourceXO) GetCriteriaReleaseType() string`

GetCriteriaReleaseType returns the CriteriaReleaseType field if non-nil, zero value otherwise.

### GetCriteriaReleaseTypeOk

`func (o *CleanupPolicyResourceXO) GetCriteriaReleaseTypeOk() (*string, bool)`

GetCriteriaReleaseTypeOk returns a tuple with the CriteriaReleaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaReleaseType

`func (o *CleanupPolicyResourceXO) SetCriteriaReleaseType(v string)`

SetCriteriaReleaseType sets CriteriaReleaseType field to given value.

### HasCriteriaReleaseType

`func (o *CleanupPolicyResourceXO) HasCriteriaReleaseType() bool`

HasCriteriaReleaseType returns a boolean if a field has been set.

### GetFormat

`func (o *CleanupPolicyResourceXO) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CleanupPolicyResourceXO) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CleanupPolicyResourceXO) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetName

`func (o *CleanupPolicyResourceXO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CleanupPolicyResourceXO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CleanupPolicyResourceXO) SetName(v string)`

SetName sets Name field to given value.


### GetNotes

`func (o *CleanupPolicyResourceXO) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CleanupPolicyResourceXO) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CleanupPolicyResourceXO) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CleanupPolicyResourceXO) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRetain

`func (o *CleanupPolicyResourceXO) GetRetain() int32`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *CleanupPolicyResourceXO) GetRetainOk() (*int32, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *CleanupPolicyResourceXO) SetRetain(v int32)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *CleanupPolicyResourceXO) HasRetain() bool`

HasRetain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


