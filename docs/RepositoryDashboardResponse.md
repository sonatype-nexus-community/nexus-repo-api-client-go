# RepositoryDashboardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**FilterInfo**](FilterInfo.md) |  | [optional] 
**GlobalConfigAvailable** | Pointer to **bool** |  | [optional] 
**HasSelections** | Pointer to **bool** |  | [optional] 
**Items** | Pointer to [**[]RepositoryDashboard**](RepositoryDashboard.md) |  | [optional] 
**NumberOfMonitoredRepositories** | Pointer to **int32** |  | [optional] 
**Pagination** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 
**Sort** | Pointer to [**SortInfo**](SortInfo.md) |  | [optional] 
**TotalRepositories** | Pointer to **int32** |  | [optional] 

## Methods

### NewRepositoryDashboardResponse

`func NewRepositoryDashboardResponse() *RepositoryDashboardResponse`

NewRepositoryDashboardResponse instantiates a new RepositoryDashboardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryDashboardResponseWithDefaults

`func NewRepositoryDashboardResponseWithDefaults() *RepositoryDashboardResponse`

NewRepositoryDashboardResponseWithDefaults instantiates a new RepositoryDashboardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *RepositoryDashboardResponse) GetFilters() FilterInfo`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RepositoryDashboardResponse) GetFiltersOk() (*FilterInfo, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RepositoryDashboardResponse) SetFilters(v FilterInfo)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RepositoryDashboardResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGlobalConfigAvailable

`func (o *RepositoryDashboardResponse) GetGlobalConfigAvailable() bool`

GetGlobalConfigAvailable returns the GlobalConfigAvailable field if non-nil, zero value otherwise.

### GetGlobalConfigAvailableOk

`func (o *RepositoryDashboardResponse) GetGlobalConfigAvailableOk() (*bool, bool)`

GetGlobalConfigAvailableOk returns a tuple with the GlobalConfigAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalConfigAvailable

`func (o *RepositoryDashboardResponse) SetGlobalConfigAvailable(v bool)`

SetGlobalConfigAvailable sets GlobalConfigAvailable field to given value.

### HasGlobalConfigAvailable

`func (o *RepositoryDashboardResponse) HasGlobalConfigAvailable() bool`

HasGlobalConfigAvailable returns a boolean if a field has been set.

### GetHasSelections

`func (o *RepositoryDashboardResponse) GetHasSelections() bool`

GetHasSelections returns the HasSelections field if non-nil, zero value otherwise.

### GetHasSelectionsOk

`func (o *RepositoryDashboardResponse) GetHasSelectionsOk() (*bool, bool)`

GetHasSelectionsOk returns a tuple with the HasSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSelections

`func (o *RepositoryDashboardResponse) SetHasSelections(v bool)`

SetHasSelections sets HasSelections field to given value.

### HasHasSelections

`func (o *RepositoryDashboardResponse) HasHasSelections() bool`

HasHasSelections returns a boolean if a field has been set.

### GetItems

`func (o *RepositoryDashboardResponse) GetItems() []RepositoryDashboard`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RepositoryDashboardResponse) GetItemsOk() (*[]RepositoryDashboard, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RepositoryDashboardResponse) SetItems(v []RepositoryDashboard)`

SetItems sets Items field to given value.

### HasItems

`func (o *RepositoryDashboardResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNumberOfMonitoredRepositories

`func (o *RepositoryDashboardResponse) GetNumberOfMonitoredRepositories() int32`

GetNumberOfMonitoredRepositories returns the NumberOfMonitoredRepositories field if non-nil, zero value otherwise.

### GetNumberOfMonitoredRepositoriesOk

`func (o *RepositoryDashboardResponse) GetNumberOfMonitoredRepositoriesOk() (*int32, bool)`

GetNumberOfMonitoredRepositoriesOk returns a tuple with the NumberOfMonitoredRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfMonitoredRepositories

`func (o *RepositoryDashboardResponse) SetNumberOfMonitoredRepositories(v int32)`

SetNumberOfMonitoredRepositories sets NumberOfMonitoredRepositories field to given value.

### HasNumberOfMonitoredRepositories

`func (o *RepositoryDashboardResponse) HasNumberOfMonitoredRepositories() bool`

HasNumberOfMonitoredRepositories returns a boolean if a field has been set.

### GetPagination

`func (o *RepositoryDashboardResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RepositoryDashboardResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RepositoryDashboardResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RepositoryDashboardResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *RepositoryDashboardResponse) GetSort() SortInfo`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RepositoryDashboardResponse) GetSortOk() (*SortInfo, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RepositoryDashboardResponse) SetSort(v SortInfo)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RepositoryDashboardResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotalRepositories

`func (o *RepositoryDashboardResponse) GetTotalRepositories() int32`

GetTotalRepositories returns the TotalRepositories field if non-nil, zero value otherwise.

### GetTotalRepositoriesOk

`func (o *RepositoryDashboardResponse) GetTotalRepositoriesOk() (*int32, bool)`

GetTotalRepositoriesOk returns a tuple with the TotalRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepositories

`func (o *RepositoryDashboardResponse) SetTotalRepositories(v int32)`

SetTotalRepositories sets TotalRepositories field to given value.

### HasTotalRepositories

`func (o *RepositoryDashboardResponse) HasTotalRepositories() bool`

HasTotalRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


