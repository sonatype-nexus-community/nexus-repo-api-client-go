# BulkUploadResultXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedCount** | Pointer to **int32** | Number of entries successfully added | [optional] 
**RejectedCount** | Pointer to **int32** | Number of entries rejected due to validation errors | [optional] 
**RejectedEntries** | Pointer to [**[]RejectedEntryXO**](RejectedEntryXO.md) | Structured list of rejected entries with line number, value, and reason | [optional] 
**SkippedCount** | Pointer to **int32** | Number of entries skipped (duplicates) | [optional] 
**TotalRows** | Pointer to **int32** | Total number of data rows processed (excluding header) | [optional] 

## Methods

### NewBulkUploadResultXO

`func NewBulkUploadResultXO() *BulkUploadResultXO`

NewBulkUploadResultXO instantiates a new BulkUploadResultXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUploadResultXOWithDefaults

`func NewBulkUploadResultXOWithDefaults() *BulkUploadResultXO`

NewBulkUploadResultXOWithDefaults instantiates a new BulkUploadResultXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedCount

`func (o *BulkUploadResultXO) GetAddedCount() int32`

GetAddedCount returns the AddedCount field if non-nil, zero value otherwise.

### GetAddedCountOk

`func (o *BulkUploadResultXO) GetAddedCountOk() (*int32, bool)`

GetAddedCountOk returns a tuple with the AddedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCount

`func (o *BulkUploadResultXO) SetAddedCount(v int32)`

SetAddedCount sets AddedCount field to given value.

### HasAddedCount

`func (o *BulkUploadResultXO) HasAddedCount() bool`

HasAddedCount returns a boolean if a field has been set.

### GetRejectedCount

`func (o *BulkUploadResultXO) GetRejectedCount() int32`

GetRejectedCount returns the RejectedCount field if non-nil, zero value otherwise.

### GetRejectedCountOk

`func (o *BulkUploadResultXO) GetRejectedCountOk() (*int32, bool)`

GetRejectedCountOk returns a tuple with the RejectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedCount

`func (o *BulkUploadResultXO) SetRejectedCount(v int32)`

SetRejectedCount sets RejectedCount field to given value.

### HasRejectedCount

`func (o *BulkUploadResultXO) HasRejectedCount() bool`

HasRejectedCount returns a boolean if a field has been set.

### GetRejectedEntries

`func (o *BulkUploadResultXO) GetRejectedEntries() []RejectedEntryXO`

GetRejectedEntries returns the RejectedEntries field if non-nil, zero value otherwise.

### GetRejectedEntriesOk

`func (o *BulkUploadResultXO) GetRejectedEntriesOk() (*[]RejectedEntryXO, bool)`

GetRejectedEntriesOk returns a tuple with the RejectedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedEntries

`func (o *BulkUploadResultXO) SetRejectedEntries(v []RejectedEntryXO)`

SetRejectedEntries sets RejectedEntries field to given value.

### HasRejectedEntries

`func (o *BulkUploadResultXO) HasRejectedEntries() bool`

HasRejectedEntries returns a boolean if a field has been set.

### GetSkippedCount

`func (o *BulkUploadResultXO) GetSkippedCount() int32`

GetSkippedCount returns the SkippedCount field if non-nil, zero value otherwise.

### GetSkippedCountOk

`func (o *BulkUploadResultXO) GetSkippedCountOk() (*int32, bool)`

GetSkippedCountOk returns a tuple with the SkippedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedCount

`func (o *BulkUploadResultXO) SetSkippedCount(v int32)`

SetSkippedCount sets SkippedCount field to given value.

### HasSkippedCount

`func (o *BulkUploadResultXO) HasSkippedCount() bool`

HasSkippedCount returns a boolean if a field has been set.

### GetTotalRows

`func (o *BulkUploadResultXO) GetTotalRows() int32`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *BulkUploadResultXO) GetTotalRowsOk() (*int32, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *BulkUploadResultXO) SetTotalRows(v int32)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *BulkUploadResultXO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


