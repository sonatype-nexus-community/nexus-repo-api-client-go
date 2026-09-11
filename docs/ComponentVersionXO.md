# ComponentVersionXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdated** | Pointer to **time.Time** | The most recent modification time across the repositories holding this version | [optional] 
**Repositories** | Pointer to **[]string** | Names of the repositories containing this version | [optional] 
**Version** | Pointer to **string** | The component version | [optional] 

## Methods

### NewComponentVersionXO

`func NewComponentVersionXO() *ComponentVersionXO`

NewComponentVersionXO instantiates a new ComponentVersionXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionXOWithDefaults

`func NewComponentVersionXOWithDefaults() *ComponentVersionXO`

NewComponentVersionXOWithDefaults instantiates a new ComponentVersionXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdated

`func (o *ComponentVersionXO) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ComponentVersionXO) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ComponentVersionXO) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ComponentVersionXO) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRepositories

`func (o *ComponentVersionXO) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ComponentVersionXO) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ComponentVersionXO) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *ComponentVersionXO) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetVersion

`func (o *ComponentVersionXO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComponentVersionXO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComponentVersionXO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComponentVersionXO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


