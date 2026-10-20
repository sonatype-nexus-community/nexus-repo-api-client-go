# PatchGlobalSettingsAndRepositoriesResponseXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | Error message if operation failed | [optional] 
**Failed** | Pointer to **map[string]string** | Map of repository IDs that failed with error messages | [optional] 
**RepositoriesAdded** | Pointer to **int32** | Count of repositories added to monitoring | [optional] 
**RepositoriesRemoved** | Pointer to **int32** | Count of repositories removed from monitoring | [optional] 
**SettingsUpdated** | Pointer to **bool** | Whether global settings were updated | [optional] 
**Success** | Pointer to **bool** | Whether the operation succeeded | [optional] 

## Methods

### NewPatchGlobalSettingsAndRepositoriesResponseXO

`func NewPatchGlobalSettingsAndRepositoriesResponseXO() *PatchGlobalSettingsAndRepositoriesResponseXO`

NewPatchGlobalSettingsAndRepositoriesResponseXO instantiates a new PatchGlobalSettingsAndRepositoriesResponseXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchGlobalSettingsAndRepositoriesResponseXOWithDefaults

`func NewPatchGlobalSettingsAndRepositoriesResponseXOWithDefaults() *PatchGlobalSettingsAndRepositoriesResponseXO`

NewPatchGlobalSettingsAndRepositoriesResponseXOWithDefaults instantiates a new PatchGlobalSettingsAndRepositoriesResponseXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFailed

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetFailed() map[string]string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetFailedOk() (*map[string]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) SetFailed(v map[string]string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetRepositoriesAdded

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetRepositoriesAdded() int32`

GetRepositoriesAdded returns the RepositoriesAdded field if non-nil, zero value otherwise.

### GetRepositoriesAddedOk

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetRepositoriesAddedOk() (*int32, bool)`

GetRepositoriesAddedOk returns a tuple with the RepositoriesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesAdded

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) SetRepositoriesAdded(v int32)`

SetRepositoriesAdded sets RepositoriesAdded field to given value.

### HasRepositoriesAdded

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) HasRepositoriesAdded() bool`

HasRepositoriesAdded returns a boolean if a field has been set.

### GetRepositoriesRemoved

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetRepositoriesRemoved() int32`

GetRepositoriesRemoved returns the RepositoriesRemoved field if non-nil, zero value otherwise.

### GetRepositoriesRemovedOk

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetRepositoriesRemovedOk() (*int32, bool)`

GetRepositoriesRemovedOk returns a tuple with the RepositoriesRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesRemoved

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) SetRepositoriesRemoved(v int32)`

SetRepositoriesRemoved sets RepositoriesRemoved field to given value.

### HasRepositoriesRemoved

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) HasRepositoriesRemoved() bool`

HasRepositoriesRemoved returns a boolean if a field has been set.

### GetSettingsUpdated

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetSettingsUpdated() bool`

GetSettingsUpdated returns the SettingsUpdated field if non-nil, zero value otherwise.

### GetSettingsUpdatedOk

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetSettingsUpdatedOk() (*bool, bool)`

GetSettingsUpdatedOk returns a tuple with the SettingsUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsUpdated

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) SetSettingsUpdated(v bool)`

SetSettingsUpdated sets SettingsUpdated field to given value.

### HasSettingsUpdated

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) HasSettingsUpdated() bool`

HasSettingsUpdated returns a boolean if a field has been set.

### GetSuccess

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PatchGlobalSettingsAndRepositoriesResponseXO) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


