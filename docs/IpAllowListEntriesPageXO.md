# IpAllowListEntriesPageXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]IpAllowListEntryXO**](IpAllowListEntryXO.md) | List of entries on the current page | [optional] 
**Page** | Pointer to **int32** | Current page number (0-based) | [optional] 
**PageSize** | Pointer to **int32** | Number of entries per page | [optional] 
**TotalEntries** | Pointer to **int32** | Total number of entries across all pages | [optional] 
**TotalPages** | Pointer to **int32** | Total number of pages | [optional] 

## Methods

### NewIpAllowListEntriesPageXO

`func NewIpAllowListEntriesPageXO() *IpAllowListEntriesPageXO`

NewIpAllowListEntriesPageXO instantiates a new IpAllowListEntriesPageXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAllowListEntriesPageXOWithDefaults

`func NewIpAllowListEntriesPageXOWithDefaults() *IpAllowListEntriesPageXO`

NewIpAllowListEntriesPageXOWithDefaults instantiates a new IpAllowListEntriesPageXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *IpAllowListEntriesPageXO) GetEntries() []IpAllowListEntryXO`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *IpAllowListEntriesPageXO) GetEntriesOk() (*[]IpAllowListEntryXO, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *IpAllowListEntriesPageXO) SetEntries(v []IpAllowListEntryXO)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *IpAllowListEntriesPageXO) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetPage

`func (o *IpAllowListEntriesPageXO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *IpAllowListEntriesPageXO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *IpAllowListEntriesPageXO) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *IpAllowListEntriesPageXO) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *IpAllowListEntriesPageXO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *IpAllowListEntriesPageXO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *IpAllowListEntriesPageXO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *IpAllowListEntriesPageXO) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalEntries

`func (o *IpAllowListEntriesPageXO) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *IpAllowListEntriesPageXO) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *IpAllowListEntriesPageXO) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *IpAllowListEntriesPageXO) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.

### GetTotalPages

`func (o *IpAllowListEntriesPageXO) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *IpAllowListEntriesPageXO) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *IpAllowListEntriesPageXO) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *IpAllowListEntriesPageXO) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


