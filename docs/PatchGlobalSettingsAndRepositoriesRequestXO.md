# PatchGlobalSettingsAndRepositoriesRequestXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityTimeFrame** | Pointer to **int32** | Activity time frame in days (optional) | [optional] 
**AddRepositoryIds** | Pointer to **[]string** | Repository IDs to add to monitoring (optional) | [optional] 
**ArtifactLatestVersions** | Pointer to **int32** | Number of latest artifact versions to track (optional) | [optional] 
**AutoEnrollNewRepos** | Pointer to **bool** | Enable auto-enrollment for new hosted repositories (optional) | [optional] 
**PolicyEvaluationStage** | Pointer to **string** | Policy evaluation stage (optional) | [optional] 
**RemoveRepositoryIds** | Pointer to **[]string** | Repository IDs to remove from monitoring (optional) | [optional] 

## Methods

### NewPatchGlobalSettingsAndRepositoriesRequestXO

`func NewPatchGlobalSettingsAndRepositoriesRequestXO() *PatchGlobalSettingsAndRepositoriesRequestXO`

NewPatchGlobalSettingsAndRepositoriesRequestXO instantiates a new PatchGlobalSettingsAndRepositoriesRequestXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchGlobalSettingsAndRepositoriesRequestXOWithDefaults

`func NewPatchGlobalSettingsAndRepositoriesRequestXOWithDefaults() *PatchGlobalSettingsAndRepositoriesRequestXO`

NewPatchGlobalSettingsAndRepositoriesRequestXOWithDefaults instantiates a new PatchGlobalSettingsAndRepositoriesRequestXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityTimeFrame

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetActivityTimeFrame() int32`

GetActivityTimeFrame returns the ActivityTimeFrame field if non-nil, zero value otherwise.

### GetActivityTimeFrameOk

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetActivityTimeFrameOk() (*int32, bool)`

GetActivityTimeFrameOk returns a tuple with the ActivityTimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTimeFrame

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) SetActivityTimeFrame(v int32)`

SetActivityTimeFrame sets ActivityTimeFrame field to given value.

### HasActivityTimeFrame

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) HasActivityTimeFrame() bool`

HasActivityTimeFrame returns a boolean if a field has been set.

### GetAddRepositoryIds

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetAddRepositoryIds() []string`

GetAddRepositoryIds returns the AddRepositoryIds field if non-nil, zero value otherwise.

### GetAddRepositoryIdsOk

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetAddRepositoryIdsOk() (*[]string, bool)`

GetAddRepositoryIdsOk returns a tuple with the AddRepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRepositoryIds

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) SetAddRepositoryIds(v []string)`

SetAddRepositoryIds sets AddRepositoryIds field to given value.

### HasAddRepositoryIds

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) HasAddRepositoryIds() bool`

HasAddRepositoryIds returns a boolean if a field has been set.

### GetArtifactLatestVersions

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetArtifactLatestVersions() int32`

GetArtifactLatestVersions returns the ArtifactLatestVersions field if non-nil, zero value otherwise.

### GetArtifactLatestVersionsOk

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetArtifactLatestVersionsOk() (*int32, bool)`

GetArtifactLatestVersionsOk returns a tuple with the ArtifactLatestVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactLatestVersions

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) SetArtifactLatestVersions(v int32)`

SetArtifactLatestVersions sets ArtifactLatestVersions field to given value.

### HasArtifactLatestVersions

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) HasArtifactLatestVersions() bool`

HasArtifactLatestVersions returns a boolean if a field has been set.

### GetAutoEnrollNewRepos

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetAutoEnrollNewRepos() bool`

GetAutoEnrollNewRepos returns the AutoEnrollNewRepos field if non-nil, zero value otherwise.

### GetAutoEnrollNewReposOk

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetAutoEnrollNewReposOk() (*bool, bool)`

GetAutoEnrollNewReposOk returns a tuple with the AutoEnrollNewRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrollNewRepos

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) SetAutoEnrollNewRepos(v bool)`

SetAutoEnrollNewRepos sets AutoEnrollNewRepos field to given value.

### HasAutoEnrollNewRepos

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) HasAutoEnrollNewRepos() bool`

HasAutoEnrollNewRepos returns a boolean if a field has been set.

### GetPolicyEvaluationStage

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetPolicyEvaluationStage() string`

GetPolicyEvaluationStage returns the PolicyEvaluationStage field if non-nil, zero value otherwise.

### GetPolicyEvaluationStageOk

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetPolicyEvaluationStageOk() (*string, bool)`

GetPolicyEvaluationStageOk returns a tuple with the PolicyEvaluationStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEvaluationStage

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) SetPolicyEvaluationStage(v string)`

SetPolicyEvaluationStage sets PolicyEvaluationStage field to given value.

### HasPolicyEvaluationStage

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) HasPolicyEvaluationStage() bool`

HasPolicyEvaluationStage returns a boolean if a field has been set.

### GetRemoveRepositoryIds

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetRemoveRepositoryIds() []string`

GetRemoveRepositoryIds returns the RemoveRepositoryIds field if non-nil, zero value otherwise.

### GetRemoveRepositoryIdsOk

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) GetRemoveRepositoryIdsOk() (*[]string, bool)`

GetRemoveRepositoryIdsOk returns a tuple with the RemoveRepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveRepositoryIds

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) SetRemoveRepositoryIds(v []string)`

SetRemoveRepositoryIds sets RemoveRepositoryIds field to given value.

### HasRemoveRepositoryIds

`func (o *PatchGlobalSettingsAndRepositoriesRequestXO) HasRemoveRepositoryIds() bool`

HasRemoveRepositoryIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


