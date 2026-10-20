# RepositoryDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | 
**HasCustomConfig** | Pointer to **bool** |  | [optional] 
**IsSelected** | Pointer to **bool** |  | [optional] 
**LastUpdated** | **time.Time** |  | 
**NumberOfComponents** | Pointer to **int32** |  | [optional] 
**RepositoryId** | **string** |  | 
**RepositoryName** | **string** |  | 
**SelectedAt** | Pointer to **time.Time** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### NewRepositoryDashboard

`func NewRepositoryDashboard(format string, lastUpdated time.Time, repositoryId string, repositoryName string, ) *RepositoryDashboard`

NewRepositoryDashboard instantiates a new RepositoryDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryDashboardWithDefaults

`func NewRepositoryDashboardWithDefaults() *RepositoryDashboard`

NewRepositoryDashboardWithDefaults instantiates a new RepositoryDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *RepositoryDashboard) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RepositoryDashboard) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RepositoryDashboard) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetHasCustomConfig

`func (o *RepositoryDashboard) GetHasCustomConfig() bool`

GetHasCustomConfig returns the HasCustomConfig field if non-nil, zero value otherwise.

### GetHasCustomConfigOk

`func (o *RepositoryDashboard) GetHasCustomConfigOk() (*bool, bool)`

GetHasCustomConfigOk returns a tuple with the HasCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomConfig

`func (o *RepositoryDashboard) SetHasCustomConfig(v bool)`

SetHasCustomConfig sets HasCustomConfig field to given value.

### HasHasCustomConfig

`func (o *RepositoryDashboard) HasHasCustomConfig() bool`

HasHasCustomConfig returns a boolean if a field has been set.

### GetIsSelected

`func (o *RepositoryDashboard) GetIsSelected() bool`

GetIsSelected returns the IsSelected field if non-nil, zero value otherwise.

### GetIsSelectedOk

`func (o *RepositoryDashboard) GetIsSelectedOk() (*bool, bool)`

GetIsSelectedOk returns a tuple with the IsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelected

`func (o *RepositoryDashboard) SetIsSelected(v bool)`

SetIsSelected sets IsSelected field to given value.

### HasIsSelected

`func (o *RepositoryDashboard) HasIsSelected() bool`

HasIsSelected returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RepositoryDashboard) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RepositoryDashboard) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RepositoryDashboard) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetNumberOfComponents

`func (o *RepositoryDashboard) GetNumberOfComponents() int32`

GetNumberOfComponents returns the NumberOfComponents field if non-nil, zero value otherwise.

### GetNumberOfComponentsOk

`func (o *RepositoryDashboard) GetNumberOfComponentsOk() (*int32, bool)`

GetNumberOfComponentsOk returns a tuple with the NumberOfComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfComponents

`func (o *RepositoryDashboard) SetNumberOfComponents(v int32)`

SetNumberOfComponents sets NumberOfComponents field to given value.

### HasNumberOfComponents

`func (o *RepositoryDashboard) HasNumberOfComponents() bool`

HasNumberOfComponents returns a boolean if a field has been set.

### GetRepositoryId

`func (o *RepositoryDashboard) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *RepositoryDashboard) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *RepositoryDashboard) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.


### GetRepositoryName

`func (o *RepositoryDashboard) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *RepositoryDashboard) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *RepositoryDashboard) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.


### GetSelectedAt

`func (o *RepositoryDashboard) GetSelectedAt() time.Time`

GetSelectedAt returns the SelectedAt field if non-nil, zero value otherwise.

### GetSelectedAtOk

`func (o *RepositoryDashboard) GetSelectedAtOk() (*time.Time, bool)`

GetSelectedAtOk returns a tuple with the SelectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedAt

`func (o *RepositoryDashboard) SetSelectedAt(v time.Time)`

SetSelectedAt sets SelectedAt field to given value.

### HasSelectedAt

`func (o *RepositoryDashboard) HasSelectedAt() bool`

HasSelectedAt returns a boolean if a field has been set.

### GetSize

`func (o *RepositoryDashboard) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RepositoryDashboard) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RepositoryDashboard) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *RepositoryDashboard) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


